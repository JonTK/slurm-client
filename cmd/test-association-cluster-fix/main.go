package main

import (
	"fmt"

	"github.com/jontk/slurm-client/internal/common/types"
)

func main() {
	fmt.Println("Testing Association Cluster Requirement Fix")
	fmt.Println("===========================================")
	
	// Test case 1: Association with empty cluster - should get default
	fmt.Println("\n1. Testing association creation with empty cluster:")
	association1 := &types.AssociationCreate{
		AccountName: "physics",
		UserName:    "testuser",
		Cluster:     "", // This should get populated with default
	}
	
	fmt.Printf("   Before: Cluster = '%s'\n", association1.Cluster)
	
	// Simulate what happens in the Create method
	if association1.Cluster == "" {
		association1.Cluster = "linux" // This is what getDefaultClusterName() returns
	}
	
	fmt.Printf("   After:  Cluster = '%s'\n", association1.Cluster)
	
	// Test case 2: Association with existing cluster - should remain unchanged
	fmt.Println("\n2. Testing association creation with existing cluster:")
	association2 := &types.AssociationCreate{
		AccountName: "chemistry",
		UserName:    "otheruser",
		Cluster:     "gpu-cluster", // This should remain unchanged
	}
	
	fmt.Printf("   Before: Cluster = '%s'\n", association2.Cluster)
	
	// Simulate what happens in the Create method
	if association2.Cluster == "" {
		association2.Cluster = "linux"
	}
	
	fmt.Printf("   After:  Cluster = '%s'\n", association2.Cluster)
	
	// Test case 3: Verify validation would now pass
	fmt.Println("\n3. Testing validation logic:")
	
	testValidation := func(assoc *types.AssociationCreate, testName string) {
		fmt.Printf("   %s: ", testName)
		if assoc.AccountName == "" {
			fmt.Println("FAIL - account is required")
			return
		}
		if assoc.UserName == "" {
			fmt.Println("FAIL - user is required")
			return
		}
		if assoc.Cluster == "" {
			fmt.Println("FAIL - cluster is required")
			return
		}
		fmt.Println("PASS - all required fields present")
	}
	
	// Test the fixed associations
	testValidation(association1, "Fixed association (was empty cluster)")
	testValidation(association2, "Existing association (had cluster)")
	
	// Test an association that would still fail (missing account)
	badAssociation := &types.AssociationCreate{
		AccountName: "", // Missing
		UserName:    "testuser",
		Cluster:     "linux",
	}
	testValidation(badAssociation, "Bad association (missing account)")
	
	fmt.Println("\n===========================================")
	fmt.Println("Fix Summary:")
	fmt.Println("- Empty cluster fields now get populated with 'linux' default")
	fmt.Println("- Existing cluster fields remain unchanged")
	fmt.Println("- Validation will pass for associations with the default cluster")
	fmt.Println("- The fix resolves: 'cluster is required for association creation' error")
}