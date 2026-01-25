// SPDX-FileCopyrightText: 2025 Jon Thor Kristinsson
// SPDX-License-Identifier: Apache-2.0

package v0_0_43

import (
	"context"
	"strings"

	api "github.com/jontk/slurm-client/internal/api/v0_0_43"
	"github.com/jontk/slurm-client/internal/common"
	"github.com/jontk/slurm-client/internal/common/types"
	"github.com/jontk/slurm-client/internal/managers/base"
	"github.com/jontk/slurm-client/pkg/errors"
)

// PartitionAdapter implements the PartitionAdapter interface for v0.0.43
type PartitionAdapter struct {
	*base.BaseManager
	client  *api.ClientWithResponses
	wrapper *api.WrapperClient
}

// NewPartitionAdapter creates a new Partition adapter for v0.0.43
func NewPartitionAdapter(client *api.ClientWithResponses) *PartitionAdapter {
	return &PartitionAdapter{
		BaseManager: base.NewBaseManager("v0.0.43", "Partition"),
		client:      client,
		wrapper:     nil, // We'll implement this later
	}
}

// List retrieves a list of partitions with optional filtering
func (a *PartitionAdapter) List(ctx context.Context, opts *types.PartitionListOptions) (*types.PartitionList, error) {
	// Use base validation
	if err := a.ValidateContext(ctx); err != nil {
		return nil, err
	}

	// Check client initialization
	if err := a.CheckClientInitialized(a.client); err != nil {
		return nil, err
	}

	// Prepare parameters for the API call
	params := &api.SlurmV0043GetPartitionsParams{}

	// Apply filters from options
	if opts != nil {
		// Note: v0.0.43 doesn't have a PartitionName parameter for filtering
		// We'll have to filter client-side
		if opts.UpdateTime != nil {
			updateTimeStr := opts.UpdateTime.Format("2006-01-02T15:04:05")
			params.UpdateTime = &updateTimeStr
		}
	}

	// Call the generated OpenAPI client
	resp, err := a.client.SlurmV0043GetPartitionsWithResponse(ctx, params)
	if err != nil {
		return nil, a.HandleAPIError(err)
	}

	// Use common response error handling
	var apiErrors *api.V0043OpenapiErrors
	if resp.JSON200 != nil {
		apiErrors = resp.JSON200.Errors
	}

	responseAdapter := api.NewResponseAdapter(resp.StatusCode(), apiErrors)
	if err := common.HandleAPIResponse(responseAdapter, "v0.0.43"); err != nil {
		return nil, err
	}

	// Check for unexpected response format
	if err := a.CheckNilResponse(resp.JSON200, "List Partitions"); err != nil {
		return nil, err
	}
	if err := a.CheckNilResponse(resp.JSON200.Partitions, "List Partitions - partitions field"); err != nil {
		return nil, err
	}

	// Convert the response to common types
	partitionList := make([]types.Partition, 0, len(resp.JSON200.Partitions))
	for _, apiPartition := range resp.JSON200.Partitions {
		partition := a.convertAPIPartitionToCommon(apiPartition)
		partitionList = append(partitionList, *partition)
	}

	// Apply client-side filtering if needed
	if opts != nil {
		partitionList = a.filterPartitionList(partitionList, opts)
	}

	// Apply pagination
	listOpts := base.ListOptions{}
	if opts != nil {
		listOpts.Limit = opts.Limit
		listOpts.Offset = opts.Offset
	}

	// Apply pagination
	start := listOpts.Offset
	if start < 0 {
		start = 0
	}
	if start >= len(partitionList) {
		return &types.PartitionList{
			Partitions: []types.Partition{},
			Total:      len(partitionList),
		}, nil
	}

	end := len(partitionList)
	if listOpts.Limit > 0 {
		end = start + listOpts.Limit
		if end > len(partitionList) {
			end = len(partitionList)
		}
	}

	return &types.PartitionList{
		Partitions: partitionList[start:end],
		Total:      len(partitionList),
	}, nil
}

// Get retrieves a specific partition by name
func (a *PartitionAdapter) Get(ctx context.Context, partitionName string) (*types.Partition, error) {
	// Use base validation
	if err := a.ValidateContext(ctx); err != nil {
		return nil, err
	}
	if err := a.ValidateResourceName(partitionName, "partition name"); err != nil {
		return nil, err
	}
	if err := a.CheckClientInitialized(a.client); err != nil {
		return nil, err
	}

	// Prepare parameters for the API call
	params := &api.SlurmV0043GetPartitionParams{}

	// Call the generated OpenAPI client
	resp, err := a.client.SlurmV0043GetPartitionWithResponse(ctx, partitionName, params)
	if err != nil {
		return nil, a.HandleAPIError(err)
	}

	// Use common response error handling
	var apiErrors *api.V0043OpenapiErrors
	if resp.JSON200 != nil {
		apiErrors = resp.JSON200.Errors
	}

	responseAdapter := api.NewResponseAdapter(resp.StatusCode(), apiErrors)
	if err := common.HandleAPIResponse(responseAdapter, "v0.0.43"); err != nil {
		return nil, err
	}

	// Check for unexpected response format
	if err := a.CheckNilResponse(resp.JSON200, "Get Partition"); err != nil {
		return nil, err
	}
	if err := a.CheckNilResponse(resp.JSON200.Partitions, "Get Partition - partitions field"); err != nil {
		return nil, err
	}

	// Check if we got any partition entries
	if len(resp.JSON200.Partitions) == 0 {
		return nil, common.NewResourceNotFoundError("Partition", partitionName)
	}

	// Convert the first partition (should be the only one)
	partition := a.convertAPIPartitionToCommon(resp.JSON200.Partitions[0])

	return partition, nil
}

// Create creates a new partition
func (a *PartitionAdapter) Create(ctx context.Context, partition *types.PartitionCreate) (*types.PartitionCreateResponse, error) {
	// v0.0.43 doesn't support partition creation through the REST API
	return nil, errors.NewClientError(
		errors.ErrorCodeUnsupportedOperation,
		"partition creation not supported in v0.0.43",
		"Method not allowed (405)")
}

// Update updates an existing partition
func (a *PartitionAdapter) Update(ctx context.Context, partitionName string, update *types.PartitionUpdate) error {
	// v0.0.43 doesn't support partition updates through the REST API
	return errors.NewClientError(
		errors.ErrorCodeUnsupportedOperation,
		"partition updates not supported in v0.0.43",
		"Method not allowed (405)")
}

// Delete deletes a partition
func (a *PartitionAdapter) Delete(ctx context.Context, partitionName string) error {
	// v0.0.43 doesn't support partition deletion through the REST API
	return errors.NewClientError(
		errors.ErrorCodeUnsupportedOperation,
		"partition deletion not supported in v0.0.43",
		"Method not allowed (405)")
}

// filterPartitionList applies client-side filtering to partition list
func (a *PartitionAdapter) filterPartitionList(partitions []types.Partition, opts *types.PartitionListOptions) []types.Partition {
	if opts == nil {
		return partitions
	}

	filtered := make([]types.Partition, 0, len(partitions))
	for _, partition := range partitions {
		if a.matchesPartitionFilters(partition, opts) {
			filtered = append(filtered, partition)
		}
	}

	return filtered
}

// matchesPartitionFilters checks if a partition matches the given filters
func (a *PartitionAdapter) matchesPartitionFilters(partition types.Partition, opts *types.PartitionListOptions) bool {
	// Filter by names (already handled by API, but included for completeness)
	if len(opts.Names) > 0 {
		found := false
		for _, name := range opts.Names {
			if strings.EqualFold(partition.Name, name) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	// Filter by states
	if len(opts.States) > 0 {
		found := false
		for _, state := range opts.States {
			if partition.State == state {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// convertAPIPartitionToCommon converts a v0.0.43 API Partition to common Partition type
// extractPartitionState extracts the partition state from API structure
func (a *PartitionAdapter) extractPartitionState(apiPartition api.V0043PartitionInfo) types.PartitionState {
	if apiPartition.Partition != nil && apiPartition.Partition.State != nil && len(*apiPartition.Partition.State) > 0 {
		return types.PartitionState((*apiPartition.Partition.State)[0])
	}
	return ""
}

// extractPartitionNodes extracts node information from API structure
func (a *PartitionAdapter) extractPartitionNodes(apiPartition api.V0043PartitionInfo) (string, int32) {
	var nodes string
	var totalNodes int32
	if apiPartition.Nodes != nil {
		if apiPartition.Nodes.Configured != nil {
			nodes = *apiPartition.Nodes.Configured
		}
		if apiPartition.Nodes.Total != nil {
			totalNodes = *apiPartition.Nodes.Total
		}
	}
	return nodes, totalNodes
}

// extractPartitionLimits extracts time and node limits from API structure
func (a *PartitionAdapter) extractPartitionLimits(apiPartition api.V0043PartitionInfo) (int32, int32, int32, int32) {
	var maxNodes, maxTime, minNodes, defaultTime int32
	if apiPartition.Maximums != nil {
		if apiPartition.Maximums.Nodes != nil && apiPartition.Maximums.Nodes.Number != nil {
			maxNodes = *apiPartition.Maximums.Nodes.Number
		}
		if apiPartition.Maximums.Time != nil && apiPartition.Maximums.Time.Number != nil {
			maxTime = *apiPartition.Maximums.Time.Number
		}
	}
	if apiPartition.Minimums != nil && apiPartition.Minimums.Nodes != nil {
		minNodes = *apiPartition.Minimums.Nodes
	}
	if apiPartition.Defaults != nil && apiPartition.Defaults.Time != nil && apiPartition.Defaults.Time.Number != nil {
		defaultTime = *apiPartition.Defaults.Time.Number
	}
	return maxNodes, maxTime, minNodes, defaultTime
}

// extractPartitionPriority extracts priority information from API structure
func (a *PartitionAdapter) extractPartitionPriority(apiPartition api.V0043PartitionInfo) int32 {
	if apiPartition.Priority != nil && apiPartition.Priority.JobFactor != nil {
		return *apiPartition.Priority.JobFactor
	}
	return 0
}

// extractPartitionAccounts extracts account allow/deny lists from API structure
func (a *PartitionAdapter) extractPartitionAccounts(apiPartition api.V0043PartitionInfo) ([]string, []string) {
	var allowAccounts, denyAccounts []string
	if apiPartition.Accounts != nil {
		if apiPartition.Accounts.Allowed != nil {
			allowAccounts = strings.Split(*apiPartition.Accounts.Allowed, ",")
		}
		if apiPartition.Accounts.Deny != nil {
			denyAccounts = strings.Split(*apiPartition.Accounts.Deny, ",")
		}
	}
	return allowAccounts, denyAccounts
}

// extractPartitionQoS extracts QoS information from API structure
func (a *PartitionAdapter) extractPartitionQoS(apiPartition api.V0043PartitionInfo) ([]string, []string, string) {
	var allowQoS, denyQoS []string
	var assignedQoS string
	if apiPartition.Qos != nil {
		if apiPartition.Qos.Allowed != nil {
			allowQoS = strings.Split(*apiPartition.Qos.Allowed, ",")
		}
		if apiPartition.Qos.Deny != nil {
			denyQoS = strings.Split(*apiPartition.Qos.Deny, ",")
		}
		if apiPartition.Qos.Assigned != nil {
			assignedQoS = *apiPartition.Qos.Assigned
		}
	}
	return allowQoS, denyQoS, assignedQoS
}

func (a *PartitionAdapter) convertAPIPartitionToCommon(apiPartition api.V0043PartitionInfo) *types.Partition {
	partition := &types.Partition{}

	// Basic fields
	if apiPartition.Name != nil {
		partition.Name = *apiPartition.Name
	}

	// State
	partition.State = a.extractPartitionState(apiPartition)

	// Nodes
	partition.Nodes, partition.TotalNodes = a.extractPartitionNodes(apiPartition)

	// Limits
	partition.MaxNodes, partition.MaxTime, partition.MinNodes, partition.DefaultTime = a.extractPartitionLimits(apiPartition)

	// Priority
	partition.Priority = a.extractPartitionPriority(apiPartition)

	// Accounts
	partition.AllowAccounts, partition.DenyAccounts = a.extractPartitionAccounts(apiPartition)

	// QoS
	partition.AllowQoS, partition.DenyQoS, partition.QoS = a.extractPartitionQoS(apiPartition)

	// Flags - v0.0.43 doesn't have a direct Flags field
	// These might be determined from other fields or defaults
	// For now, we'll leave them as false

	return partition
}
