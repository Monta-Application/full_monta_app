// Copyright (c) 2024 Trenova Technologies, LLC
//
// Licensed under the Business Source License 1.1 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://trenova.app/pricing/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// Key Terms:
// - Non-production use only
// - Change Date: 2026-11-16
// - Change License: GNU General Public License v2 or later
//
// For full license text, see the LICENSE file in the root directory.

package services

import (
	"context"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/emoss08/trenova/internal/server"
	"github.com/gofiber/contrib/websocket"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type TaskStatusUpdate struct {
	TaskID         string `json:"task_id"`
	Status         string `json:"status"`
	Result         string `json:"result,omitempty"`
	Error          string `json:"error,omitempty"`
	ClientID       string `json:"client_id"`
	OrganizationID string `json:"organization_id"`
	BusinessUnitID string `json:"business_unit_id"`
}

// Message represents a message sent over the websocket connection.
type Message struct {
	Type     string `json:"type"`
	Title    string `json:"title,omitempty"`
	Content  string `json:"content"`
	ClientID string `json:"clientId,omitempty"` // optional field
}

var (
	// clients stores the active clients with their IDs as keys
	clients = make(map[string]*websocket.Conn)
	// mutex to synchronize access to the clients map
	mu sync.Mutex
)

type WebsocketService struct {
	db              *bun.DB
	logger          *zerolog.Logger
	heartbeatCancel context.CancelFunc
}

func NewWebsocketService(s *server.Server) *WebsocketService {
	wsService := &WebsocketService{
		db:     s.DB,
		logger: s.Logger,
	}

	ctx, cancel := context.WithCancel(context.Background())
	go wsService.StartHeartbeat(ctx, 30*time.Second)
	wsService.heartbeatCancel = cancel

	return wsService
}

func (s WebsocketService) StartHeartbeat(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info().Msg("Stopping heartbeat")
			return
		case <-ticker.C:
			s.pingClients()
		}
	}
}

func (s WebsocketService) pingClients() {
	mu.Lock()
	defer mu.Unlock()

	for id, conn := range clients {
		if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			s.logger.Error().Err(err).Msgf("Failed to ping client %s, unregistering", id)
			s.UnregisterClient(id)
		}
	}
}

func (s WebsocketService) RegisterClient(id string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	// Check if the clientID is already registered
	if existingConn, ok := clients[id]; ok {
		// Notify the client about the reconnection
		s.notifyClientAboutReconnection(existingConn)

		// Close the existing connection
		_ = existingConn.Close()
		s.logger.Info().Str("client_id", id).Msg("existing connection closed for re-registration")
	}

	// Register the new connection
	clients[id] = conn
	s.logger.Info().Str("client_id", id).Msg("client registered")
}

func (s WebsocketService) notifyClientAboutReconnection(conn *websocket.Conn) {
	message := Message{
		Type:    "reconnect",
		Title:   "Global Websocket",
		Content: "You have been reconnected due to a duplicate registration.",
	}
	jsonData, err := sonic.Marshal(message)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to marshal reconnection message")
		return
	}

	if err = conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		s.logger.Error().Err(err).Msg("failed to send reconnection message")
	}
}

func (s WebsocketService) UnregisterClient(id string) {
	mu.Lock()

	if conn, ok := clients[id]; ok {
		_ = conn.Close() // Attempt to close the websocket connection gracefully
		delete(clients, id)
	}
	mu.Unlock()

	s.logger.Info().Str("client_id", id).Msg("client unregistered")
}

func (s WebsocketService) NotifyClient(clientID string, message Message) {
	mu.Lock()
	conn, ok := clients[clientID]
	mu.Unlock()

	if ok {
		jsonData, err := sonic.Marshal(message)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed to marshal message")
			return
		}

		if err = conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			s.logger.Error().Err(err).Msgf("Failed to send message to client %s", clientID)
		}
	}
}

func (s WebsocketService) NotifyAllClients(msg Message, senderID string) {
	message, err := sonic.Marshal(msg)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to marshal message")
		return
	}

	mu.Lock()
	for id, conn := range clients {
		if id == senderID {
			continue
		}

		mu.Unlock()
		if err = conn.WriteMessage(websocket.TextMessage, message); err != nil {
			s.logger.Error().Err(err).Msgf("Failed to send message to client %s", id)
		}
		mu.Lock()
	}

	mu.Unlock()
}

func (s WebsocketService) Stop() {
	if s.heartbeatCancel != nil {
		s.heartbeatCancel()
	}
	s.logger.Info().Msg("Websocket service stopped")
}
