// SPDX-FileCopyrightText: 2025 Jon Thor Kristinsson
// SPDX-License-Identifier: Apache-2.0

package v0_0_43

import (
	"context"
	"testing"

	api "github.com/jontk/slurm-client/internal/api/v0_0_43"
	"github.com/jontk/slurm-client/internal/common/types"
	"github.com/stretchr/testify/assert"
)

func TestNewReservationAdapter(t *testing.T) {
	adapter := NewReservationAdapter(&api.ClientWithResponses{})
	assert.NotNil(t, adapter)
	assert.NotNil(t, adapter.BaseManager)
}

func TestReservationAdapter_ValidateContext(t *testing.T) {
	adapter := NewReservationAdapter(&api.ClientWithResponses{})

	// Test nil context
	err := adapter.ValidateContext(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test valid context
	err = adapter.ValidateContext(context.Background())
	assert.NoError(t, err)
}

func TestReservationAdapter_List(t *testing.T) {
	adapter := NewReservationAdapter(nil) // Use nil client for testing validation logic

	// Test nil context validation
	_, err := adapter.List(nil, &types.ReservationListOptions{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	_, err = adapter.List(context.Background(), &types.ReservationListOptions{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestReservationAdapter_Get(t *testing.T) {
	adapter := NewReservationAdapter(nil)

	// Test empty reservation name
	_, err := adapter.Get(context.Background(), "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reservation name is required")

	// Test nil context
	_, err = adapter.Get(nil, "test-reservation")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	_, err = adapter.Get(context.Background(), "test-reservation")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestReservationAdapter_ConvertAPIReservationToCommon(t *testing.T) {
	adapter := NewReservationAdapter(&api.ClientWithResponses{})

	tests := []struct {
		name               string
		apiReservation     api.V0043ReservationInfo
		expectedName       string
		expectedPartition  string
	}{
		{
			name: "full reservation",
			apiReservation: api.V0043ReservationInfo{
				Name:      ptrString("maintenance"),
				Partition: ptrString("compute"),
				Accounts:  ptrString("root"),
				Users:     ptrString("admin"),
			},
			expectedName:      "maintenance",
			expectedPartition: "compute",
		},
		{
			name: "minimal reservation",
			apiReservation: api.V0043ReservationInfo{
				Name: ptrString("test-res"),
			},
			expectedName:      "test-res",
			expectedPartition: "",
		},
		{
			name:               "empty reservation",
			apiReservation:     api.V0043ReservationInfo{},
			expectedName:       "",
			expectedPartition:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := adapter.convertAPIReservationToCommon(tt.apiReservation)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedName, result.Name)
			// Note: Check available fields in the result struct
			// assert.Equal(t, tt.expectedPartition, result.Partition)
		})
	}
}

func TestReservationAdapter_Create(t *testing.T) {
	adapter := NewReservationAdapter(nil)

	// Test nil reservation
	_, err := adapter.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reservation creation data is required")

	// Test missing required fields
	_, err = adapter.Create(context.Background(), &types.ReservationCreate{
		Name: "",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reservation name is required")

	// Test nil context
	_, err = adapter.Create(nil, &types.ReservationCreate{
		Name: "test-reservation",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	_, err = adapter.Create(context.Background(), &types.ReservationCreate{
		Name: "test-reservation",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestReservationAdapter_Update(t *testing.T) {
	adapter := NewReservationAdapter(nil)

	// Test nil update
	err := adapter.Update(context.Background(), "test-reservation", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reservation update data is required")

	// Test empty reservation name
	err = adapter.Update(context.Background(), "", &types.ReservationUpdate{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reservation name is required")

	// Test nil context
	err = adapter.Update(nil, "test-reservation", &types.ReservationUpdate{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	err = adapter.Update(context.Background(), "test-reservation", &types.ReservationUpdate{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestReservationAdapter_Delete(t *testing.T) {
	adapter := NewReservationAdapter(nil)

	// Test empty reservation name
	err := adapter.Delete(context.Background(), "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "reservation name is required")

	// Test nil context
	err = adapter.Delete(nil, "test-reservation")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	err = adapter.Delete(context.Background(), "test-reservation")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestReservationAdapter_ValidateReservationCreate(t *testing.T) {
	adapter := NewReservationAdapter(&api.ClientWithResponses{})

	tests := []struct {
		name          string
		reservation   *types.ReservationCreate
		expectedError bool
		errorContains string
	}{
		{
			name: "valid reservation",
			reservation: &types.ReservationCreate{
				Name: "test-reservation",
			},
			expectedError: false,
		},
		{
			name:          "nil reservation",
			reservation:   nil,
			expectedError: true,
			errorContains: "reservation creation data is required",
		},
		{
			name: "missing name",
			reservation: &types.ReservationCreate{
				Name: "",
			},
			expectedError: true,
			errorContains: "reservation name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := adapter.validateReservationCreate(tt.reservation)

			if tt.expectedError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestReservationAdapter_ValidateReservationUpdate(t *testing.T) {
	adapter := NewReservationAdapter(&api.ClientWithResponses{})

	tests := []struct {
		name          string
		update        *types.ReservationUpdate
		expectedError bool
		errorContains string
	}{
		{
			name: "valid update",
			update: &types.ReservationUpdate{
				// Note: Check ReservationUpdate struct fields
			// Accounts: ptrString("newaccount"),
			},
			expectedError: false,
		},
		{
			name:          "nil update",
			update:        nil,
			expectedError: true,
			errorContains: "reservation update data is required",
		},
		{
			name:          "empty update",
			update:        &types.ReservationUpdate{},
			expectedError: false, // Empty updates are allowed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := adapter.validateReservationUpdate(tt.update)

			if tt.expectedError {
				assert.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}