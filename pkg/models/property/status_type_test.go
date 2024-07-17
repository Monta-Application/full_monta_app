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

package property_test

import (
	"testing"

	"github.com/emoss08/trenova/pkg/models/property"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatus_String(t *testing.T) {
	status := property.StatusActive
	assert.Equal(t, "Active", status.String())
}

func TestStatus_Values(t *testing.T) {
	values := property.Status("").Values()
	assert.Equal(t, []string{"Active", "Inactive"}, values)
}

func TestStatus_Value(t *testing.T) {
	status := property.StatusActive
	val, err := status.Value()
	require.NoError(t, err)
	assert.Equal(t, "Active", val)
}

func TestStatus_Scan(t *testing.T) {
	var status property.Status

	err := status.Scan("Active")
	require.NoError(t, err)
	assert.Equal(t, property.StatusActive, status)

	err = status.Scan(nil)
	require.Error(t, err)

	err = status.Scan(123)
	assert.Error(t, err)
}
