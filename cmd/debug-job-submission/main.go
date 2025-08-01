package main

import (
	"context"
	"encoding/json"
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

	// Create configuration with debug enabled
	cfg := config.NewClientConfig(
		config.WithBaseURL("http://rocky9.ar.jontk.com:6820"),
		config.WithJWTToken(jwtToken),
		config.WithTimeout(30*time.Second),
		config.WithAPIVersion("v0.0.43"),
		config.WithUseAdapters(true),
		config.WithDebug(true), // Enable debug mode
	)

	// Create client
	c, err := client.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== SLURM Job Submission Debug ===")
	fmt.Println("Server:", cfg.BaseURL)
	fmt.Println("API Version:", cfg.APIVersion)
	fmt.Println("JWT Token: [PROVIDED]")
	fmt.Println()

	// First, list available partitions to ensure we're using a valid one
	fmt.Println("1. Listing available partitions:")
	partitions, err := c.Partitions().List(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to list partitions: %v\n", err)
	} else {
		fmt.Printf("Found %d partitions:\n", partitions.Total)
		for _, p := range partitions.Items {
			fmt.Printf("  - %s (State: %s, TotalNodes: %d)\n", p.Name, p.State, p.TotalNodes)
		}
	}
	fmt.Println()

	// List accounts to ensure we're using a valid account
	fmt.Println("2. Listing available accounts:")
	accounts, err := c.Accounts().List(ctx, nil)
	if err != nil {
		fmt.Printf("Failed to list accounts: %v\n", err)
	} else {
		fmt.Printf("Found %d accounts:\n", accounts.Total)
		for _, a := range accounts.Items {
			fmt.Printf("  - %s (Org: %s)\n", a.Name, a.Organization)
		}
	}
	fmt.Println()

	// Test different job submission configurations
	testConfigs := []struct {
		name string
		job  *types.JobCreate
	}{
		{
			name: "Minimal job with script",
			job: &types.JobCreate{
				Name:    "test-job-minimal",
				Account: "root",
				Script:  "#!/bin/bash\nhostname",
			},
		},
		{
			name: "Job with partition",
			job: &types.JobCreate{
				Name:      "test-job-partition",
				Account:   "root",
				Partition: "normal",
				Script:    "#!/bin/bash\nhostname",
			},
		},
		{
			name: "Job with all basic fields",
			job: &types.JobCreate{
				Name:             "test-job-full",
				Account:          "root",
				Partition:        "normal",
				Script:           "#!/bin/bash\nhostname\ndate\necho 'Test job'",
				TimeLimit:        60, // 1 minute
				Nodes:            1,
				WorkingDirectory: "/tmp",
				Environment: map[string]string{
					"PATH": "/usr/bin:/bin",
					"USER": "root",
					"HOME": "/tmp",
				},
			},
		},
		{
			name: "Job with command instead of script",
			job: &types.JobCreate{
				Name:      "test-job-command",
				Account:   "root",
				Partition: "normal",
				Command:   "/bin/hostname",
				TimeLimit: 60,
			},
		},
	}

	fmt.Println("3. Testing job submission configurations:")
	for i, tc := range testConfigs {
		fmt.Printf("\nTest %d: %s\n", i+1, tc.name)
		fmt.Println("-" * 40)
		
		// Print job configuration
		jobJSON, _ := json.MarshalIndent(tc.job, "  ", "  ")
		fmt.Printf("Job Config:\n%s\n", jobJSON)
		
		// Try to submit the job
		resp, err := c.Jobs().Submit(ctx, tc.job)
		if err != nil {
			fmt.Printf("❌ FAILED: %v\n", err)
			// Try to get more details about the error
			if errStr := fmt.Sprintf("%v", err); errStr != "" {
				fmt.Printf("Error Details: %s\n", errStr)
			}
		} else {
			fmt.Printf("✅ SUCCESS: Job submitted with ID %d\n", resp.JobID)
			
			// Try to get job details
			jobInfo, err := c.Jobs().Get(ctx, resp.JobID)
			if err != nil {
				fmt.Printf("Failed to get job info: %v\n", err)
			} else {
				fmt.Printf("Job State: %s, Name: %s\n", jobInfo.State, jobInfo.Name)
			}
		}
	}

	// Test with different account names
	fmt.Println("\n4. Testing with different account names:")
	accountTests := []string{"root", "default", ""}
	for _, account := range accountTests {
		job := &types.JobCreate{
			Name:      fmt.Sprintf("test-job-account-%s", account),
			Account:   account,
			Partition: "normal",
			Script:    "#!/bin/bash\nhostname",
		}
		
		fmt.Printf("\nTesting with Account='%s'\n", account)
		resp, err := c.Jobs().Submit(ctx, job)
		if err != nil {
			fmt.Printf("❌ FAILED: %v\n", err)
		} else {
			fmt.Printf("✅ SUCCESS: Job ID %d\n", resp.JobID)
		}
	}
}