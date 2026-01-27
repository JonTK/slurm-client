// SPDX-FileCopyrightText: 2025 Jon Thor Kristinsson
// SPDX-License-Identifier: Apache-2.0

// Test v0.0.42 adapter methods for Shares, TRES, WCKeys, Accounts, Users, Associations
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jontk/slurm-client/interfaces"
	"github.com/jontk/slurm-client/internal/factory"
	"github.com/jontk/slurm-client/pkg/auth"
	"github.com/jontk/slurm-client/pkg/config"
)

func main() {
	ctx := context.Background()

	// Get JWT token from environment
	jwtToken := os.Getenv("SLURM_JWT")
	if jwtToken == "" {
		log.Fatal("SLURM_JWT environment variable is required")
	}

	// Create configuration
	cfg := config.NewDefault()
	cfg.BaseURL = "http://rocky9.ar.jontk.com:6820"
	cfg.Debug = false

	// Create JWT authentication provider
	authProvider := auth.NewTokenAuth(jwtToken)

	// Create factory with adapters enabled
	clientFactory, err := factory.NewClientFactory(
		factory.WithConfig(cfg),
		factory.WithAuth(authProvider),
		factory.WithBaseURL(cfg.BaseURL),
		factory.WithUseAdapters(true), // Enable adapter pattern
	)
	if err != nil {
		log.Fatalf("Failed to create factory: %v", err)
	}

	// Create client for v0.0.42 using adapters
	c, err := clientFactory.NewClientWithVersion(ctx, "v0.0.42")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=== Testing v0.0.42 Adapter Methods ===")

	// Test GetShares (standalone method)
	fmt.Println("1. Testing GetShares()...")
	shares, err := c.GetShares(ctx, &interfaces.GetSharesOptions{})
	if err != nil {
		fmt.Printf("   ❌ ERROR: %v\n", err)
	} else {
		fmt.Printf("   ✅ SUCCESS: Retrieved %d shares\n", len(shares.Shares))
		if len(shares.Shares) > 0 {
			if jsonData, jsonErr := json.MarshalIndent(shares.Shares[0], "      ", "  "); jsonErr == nil {
				fmt.Printf("      Sample: %s\n", string(jsonData))
			}
		}
	}
	fmt.Println()

	// Test GetTRES (standalone method)
	fmt.Println("2. Testing GetTRES()...")
	tres, err := c.GetTRES(ctx)
	if err != nil {
		fmt.Printf("   ❌ ERROR: %v\n", err)
	} else {
		fmt.Printf("   ✅ SUCCESS: Retrieved %d TRES\n", len(tres.TRES))
		if len(tres.TRES) > 0 {
			if jsonData, jsonErr := json.MarshalIndent(tres.TRES[0], "      ", "  "); jsonErr == nil {
				fmt.Printf("      Sample: %s\n", string(jsonData))
			}
		}
	}
	fmt.Println()

	// Test WCKeys
	fmt.Println("3. Testing WCKeys.List()...")
	wckeys, err := c.WCKeys().List(ctx, nil)
	if err != nil {
		fmt.Printf("   ❌ ERROR: %v\n", err)
	} else {
		fmt.Printf("   ✅ SUCCESS: Retrieved %d WCKeys\n", len(wckeys.WCKeys))
		if len(wckeys.WCKeys) > 0 {
			if jsonData, jsonErr := json.MarshalIndent(wckeys.WCKeys[0], "      ", "  "); jsonErr == nil {
				fmt.Printf("      Sample: %s\n", string(jsonData))
			}
		}
	}
	fmt.Println()

	// Test Accounts
	fmt.Println("4. Testing Accounts.List()...")
	accounts, err := c.Accounts().List(ctx, nil)
	if err != nil {
		fmt.Printf("   ❌ ERROR: %v\n", err)
	} else {
		fmt.Printf("   ✅ SUCCESS: Retrieved %d accounts\n", len(accounts.Accounts))
		if len(accounts.Accounts) > 0 {
			if jsonData, jsonErr := json.MarshalIndent(accounts.Accounts[0], "      ", "  "); jsonErr == nil {
				fmt.Printf("      Sample: %s\n", string(jsonData))
			}
		}
	}
	fmt.Println()

	// Test Users
	fmt.Println("5. Testing Users.List()...")
	users, err := c.Users().List(ctx, nil)
	if err != nil {
		fmt.Printf("   ❌ ERROR: %v\n", err)
	} else {
		fmt.Printf("   ✅ SUCCESS: Retrieved %d users\n", len(users.Users))
		if len(users.Users) > 0 {
			if jsonData, jsonErr := json.MarshalIndent(users.Users[0], "      ", "  "); jsonErr == nil {
				fmt.Printf("      Sample: %s\n", string(jsonData))
			}
		}
	}
	fmt.Println()

	// Test Associations
	fmt.Println("6. Testing Associations.List()...")
	assocs, err := c.Associations().List(ctx, nil)
	if err != nil {
		fmt.Printf("   ❌ ERROR: %v\n", err)
	} else {
		fmt.Printf("   ✅ SUCCESS: Retrieved %d associations\n", len(assocs.Associations))
		if len(assocs.Associations) > 0 {
			if jsonData, jsonErr := json.MarshalIndent(assocs.Associations[0], "      ", "  "); jsonErr == nil {
				fmt.Printf("      Sample: %s\n", string(jsonData))
			}
		}
	}
	fmt.Println()

	fmt.Println("=== Test Complete ===")
}
