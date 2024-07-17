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

package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AuditLog struct {
	bun.BaseModel `bun:"table:audit_logs,alias:al"`
	ID            uuid.UUID       `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	EntityType    string          `bun:"type:varchar(50),notnull"`
	EntityID      uuid.UUID       `bun:"type:uuid,notnull"`
	Action        string          `bun:"type:varchar(50),notnull"`
	ChangedFields json.RawMessage `bun:"type:jsonb"`
	UserID        uuid.UUID       `bun:"type:uuid,notnull"`
	Timestamp     time.Time       `bun:",nullzero,notnull,default:current_timestamp"`
}
