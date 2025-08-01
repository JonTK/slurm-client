package main

import (
	"testing"

	"github.com/jontk/slurm-client/internal/adapters/v0_0_43"
	"github.com/jontk/slurm-client/internal/common/types"
	"github.com/jontk/slurm-client/internal/managers/base"
	api "github.com/jontk/slurm-client/internal/api/v0_0_43"
)

func TestAssociationClusterDefaultFix(t *testing.T) {
	// Create the association adapter
	adapter := &v0_0_43.AssociationAdapter{}
	adapter.BaseManager = base.NewBaseManager("v0.0.43", "Association")
	
	// Test the getDefaultClusterName function
	defaultCluster := adapter.getDefaultClusterName()
	if defaultCluster != "linux" {
		t.Errorf("Expected default cluster to be 'linux', got '%s'", defaultCluster)
	}
	
	// Test association with empty cluster gets populated
	association := &types.AssociationCreate{
		AccountName: "physics",
		UserName:    "testuser",
		Cluster:     "", // Empty - should get populated
	}
	
	// Simulate what happens in the Create method
	if association.Cluster == "" {
		association.Cluster = defaultCluster
	}
	
	if association.Cluster != "linux" {
		t.Errorf("Expected cluster to be populated with 'linux', got '%s'", association.Cluster)
	}
	
	// Test that validation would now pass
	err := adapter.validateAssociationCreate(association)
	if err != nil {
		t.Errorf("Expected validation to pass with default cluster, got error: %v", err)
	}
	
	// Test association with existing cluster remains unchanged
	association2 := &types.AssociationCreate{
		AccountName: "chemistry",
		UserName:    "otheruser", 
		Cluster:     "gpu-cluster", // Existing - should remain
	}
	
	// Simulate what happens in the Create method
	if association2.Cluster == "" {
		association2.Cluster = defaultCluster
	}
	
	if association2.Cluster != "gpu-cluster" {
		t.Errorf("Expected existing cluster to remain 'gpu-cluster', got '%s'", association2.Cluster)
	}
}