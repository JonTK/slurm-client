package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jontk/slurm-client/internal/client"
	"github.com/jontk/slurm-client/internal/config"
	"github.com/jontk/slurm-client/internal/common/types"
)

func main() {
	// JWT token provided by the user
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI2NTQwMDQ0ODIsImlhdCI6MTc1NDAwNDQ4Miwic3VuIjoicm9vdCJ9.2QNEDMv8t3VzDQ-VsoUb8Gmkc9rMVPzqqz3sU5vy8NY"
	
	// Override with environment variable if provided
	if envToken := os.Getenv("SLURM_JWT"); envToken != "" {
		jwtToken = envToken
	}

	// Create configuration
	cfg := config.NewClientConfig(
		config.WithBaseURL("http://rocky9.ar.jontk.com:6820"),
		config.WithJWTToken(jwtToken),
		config.WithTimeout(30*time.Second),
		config.WithAPIVersion("v0.0.43"),
		config.WithUseAdapters(true),
	)

	// Create client
	c, err := client.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== SLURM Client Critical Fixes Test ===")
	fmt.Println("Server:", cfg.BaseURL)
	fmt.Println("API Version:", cfg.APIVersion)
	fmt.Println("JWT Token: [PROVIDED]")
	fmt.Println()

	// Test 1: Job Submission with Account field
	fmt.Println("TEST 1: Job Submission with Account Field")
	fmt.Println("-" * 40)
	testJobSubmission(ctx, c)
	fmt.Println()

	// Test 2: Account Creation with all required fields
	fmt.Println("TEST 2: Account Creation with Required Fields")
	fmt.Println("-" * 40)
	testAccountCreation(ctx, c)
	fmt.Println()

	// Test 3: Reservation Creation with proper time format
	fmt.Println("TEST 3: Reservation Creation with Time Format")
	fmt.Println("-" * 40)
	testReservationCreation(ctx, c)
	fmt.Println()

	// Test 4: Association Creation with cluster field
	fmt.Println("TEST 4: Association Creation with Cluster Field")
	fmt.Println("-" * 40)
	testAssociationCreation(ctx, c)
	fmt.Println()

	fmt.Println("=== All Critical Fixes Tests Completed ===")
}

func testJobSubmission(ctx context.Context, c *client.Client) {
	job := &types.JobCreate{
		Name:      "test-job-" + time.Now().Format("20060102-150405"),
		Account:   "root", // Account field is now required
		Partition: "normal",
		Script:    "#!/bin/bash\nhostname\ndate\necho 'Job submission test successful'",
		TimeLimit: 60, // 1 minute
		Nodes:     1,
		Environment: map[string]string{
			"TEST_VAR": "test_value",
		},
	}

	fmt.Printf("Submitting job with Account=%s, Partition=%s\n", job.Account, job.Partition)
	
	resp, err := c.Jobs().Submit(ctx, job)
	if err != nil {
		fmt.Printf("❌ FAILED: %v\n", err)
		return
	}

	fmt.Printf("✅ SUCCESS: Job submitted with ID %d\n", resp.JobID)
	
	// Verify job was created
	jobInfo, err := c.Jobs().Get(ctx, resp.JobID)
	if err != nil {
		fmt.Printf("❌ Failed to get job info: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Job verified: Name=%s, State=%s, Account=%s\n", 
		jobInfo.Name, jobInfo.State, jobInfo.Account)
}

func testAccountCreation(ctx context.Context, c *client.Client) {
	account := &types.AccountCreate{
		Name:         "test-account-" + time.Now().Format("20060102-150405"),
		Description:  "Test account for critical fixes validation",
		Organization: "Test Organization",
		Coordinators: []string{"root"},
	}

	fmt.Printf("Creating account: Name=%s, Org=%s\n", account.Name, account.Organization)
	
	resp, err := c.Accounts().Create(ctx, account)
	if err != nil {
		fmt.Printf("❌ FAILED: %v\n", err)
		// This might fail due to permissions, which is expected
		fmt.Println("Note: Account creation may require admin privileges")
		return
	}

	fmt.Printf("✅ SUCCESS: Account created: %s\n", resp.Name)
	
	// Try to retrieve the account
	accountInfo, err := c.Accounts().Get(ctx, account.Name)
	if err != nil {
		fmt.Printf("❌ Failed to get account info: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Account verified: Name=%s, Desc=%s, Org=%s\n",
		accountInfo.Name, accountInfo.Description, accountInfo.Organization)
}

func testReservationCreation(ctx context.Context, c *client.Client) {
	startTime := time.Now().Add(5 * time.Minute)
	endTime := startTime.Add(30 * time.Minute)
	
	reservation := &types.ReservationCreate{
		Name:      "test-res-" + time.Now().Format("20060102-150405"),
		StartTime: startTime,
		EndTime:   &endTime,
		NodeCount: 1,
		Users:     []string{"root"},
		Accounts:  []string{"root"},
		Partition: "normal",
	}

	fmt.Printf("Creating reservation: Name=%s, Start=%s, End=%s\n", 
		reservation.Name, 
		reservation.StartTime.Format("2006-01-02 15:04:05"),
		reservation.EndTime.Format("2006-01-02 15:04:05"))
	
	resp, err := c.Reservations().Create(ctx, reservation)
	if err != nil {
		fmt.Printf("❌ FAILED: %v\n", err)
		// This might fail due to permissions or resource availability
		fmt.Println("Note: Reservation creation may require admin privileges or available nodes")
		return
	}

	fmt.Printf("✅ SUCCESS: Reservation created: %s\n", resp.ReservationName)
	
	// Try to retrieve the reservation
	resInfo, err := c.Reservations().Get(ctx, reservation.Name)
	if err != nil {
		fmt.Printf("❌ Failed to get reservation info: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Reservation verified: Name=%s, Nodes=%d, Users=%v\n",
		resInfo.Name, resInfo.NodeCount, resInfo.Users)
}

func testAssociationCreation(ctx context.Context, c *client.Client) {
	association := &types.AssociationCreate{
		AccountName: "root",
		UserName:    "root",
		Cluster:     "linux", // Cluster field is auto-populated if not provided
		Partition:   "normal",
	}

	fmt.Printf("Creating association: Account=%s, User=%s, Cluster=%s\n", 
		association.AccountName, association.UserName, association.Cluster)
	
	resp, err := c.Associations().Create(ctx, association)
	if err != nil {
		fmt.Printf("❌ FAILED: %v\n", err)
		// This might fail if association already exists
		fmt.Println("Note: Association may already exist")
		return
	}

	fmt.Printf("✅ SUCCESS: Association created: ID=%d\n", resp.ID)
	
	// List associations to verify
	opts := &types.AssociationListOptions{
		Users:    []string{association.UserName},
		Accounts: []string{association.AccountName},
		Clusters: []string{association.Cluster},
	}
	
	assocList, err := c.Associations().List(ctx, opts)
	if err != nil {
		fmt.Printf("❌ Failed to list associations: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Found %d associations for user=%s, account=%s\n", 
		assocList.Total, association.UserName, association.AccountName)
}