// SPDX-FileCopyrightText: 2025 Jon Thor Kristinsson
// SPDX-License-Identifier: Apache-2.0

package v0_0_44

import (
	"context"
	"testing"

	api "github.com/jontk/slurm-client/internal/api/v0_0_44"
	"github.com/jontk/slurm-client/internal/common/types"
	"github.com/stretchr/testify/assert"
)

func TestNewClusterAdapter(t *testing.T) {
	adapter := NewClusterAdapter(&api.ClientWithResponses{})
	assert.NotNil(t, adapter)
	assert.NotNil(t, adapter.BaseManager)
}

func TestClusterAdapter_ValidateContext(t *testing.T) {
	adapter := NewClusterAdapter(&api.ClientWithResponses{})

	// Test nil context
	err := adapter.ValidateContext(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test valid context
	err = adapter.ValidateContext(context.Background())
	assert.NoError(t, err)
}

func TestClusterAdapter_List(t *testing.T) {
	adapter := NewClusterAdapter(nil) // Use nil client for testing validation logic

	// Test nil context validation
	_, err := adapter.List(nil, &types.ClusterListOptions{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	_, err = adapter.List(context.Background(), &types.ClusterListOptions{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestClusterAdapter_Get(t *testing.T) {
	adapter := NewClusterAdapter(nil)

	// Test empty cluster name
	_, err := adapter.Get(context.Background(), "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cluster name is required")

	// Test nil context
	_, err = adapter.Get(nil, "test-cluster")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	_, err = adapter.Get(context.Background(), "test-cluster")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}

func TestClusterAdapter_ConvertAPIClusterToCommon(t *testing.T) {
	adapter := NewClusterAdapter(&api.ClientWithResponses{})

	tests := []struct {
		name         string
		apiCluster   api.V0044ClusterRec
		expectedName string
	}{
		{
			name: "full cluster",
			apiCluster: api.V0044ClusterRec{
				Name: ptrString("cluster1"),
				Controller: &struct {
					Host *string `json:"host,omitempty"`
					Port *int32  `json:"port,omitempty"`
				}{
					Host: ptrString("controller1"),
					Port: ptrInt32(6817),
				},
			},
			expectedName: "cluster1",
		},
		{
			name: "minimal cluster",
			apiCluster: api.V0044ClusterRec{
				Name: ptrString("cluster2"),
			},
			expectedName: "cluster2",
		},
		{
			name:         "empty cluster",
			apiCluster:   api.V0044ClusterRec{},
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := adapter.convertAPIClusterToCommon(tt.apiCluster)
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedName, result.Name)
		})
	}
}

func TestClusterAdapter_Create(t *testing.T) {
	adapter := NewClusterAdapter(nil)

	// Test nil cluster
	_, err := adapter.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cluster creation data is required")

	// Test missing required fields
	_, err = adapter.Create(context.Background(), &types.ClusterCreate{
		Name: "",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cluster name is required")

	// Test nil context
	_, err = adapter.Create(nil, &types.ClusterCreate{
		Name: "test-cluster",
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")
}

// Note: v0.0.44 API doesn't support cluster updates
// The Update method doesn't exist on ClusterAdapter for this version

func TestClusterAdapter_Delete(t *testing.T) {
	adapter := NewClusterAdapter(nil)

	// Test empty cluster name
	err := adapter.Delete(context.Background(), "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cluster name is required")

	// Test nil context
	err = adapter.Delete(nil, "test-cluster")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context is required")

	// Test client initialization check
	err = adapter.Delete(context.Background(), "test-cluster")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "client not initialized")
}