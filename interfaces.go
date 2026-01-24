// SPDX-FileCopyrightText: 2025 Jon Thor Kristinsson
// SPDX-License-Identifier: Apache-2.0

package slurm

import (
	"github.com/jontk/slurm-client/interfaces"
)

// SlurmClient represents a version-agnostic Slurm REST API client
// This is a type alias to the internal interface to avoid import cycles
type SlurmClient = interfaces.SlurmClient

// JobManager provides version-agnostic job operations
type JobManager = interfaces.JobManager

// NodeManager provides version-agnostic node operations
type NodeManager = interfaces.NodeManager

// PartitionManager provides version-agnostic partition operations
type PartitionManager = interfaces.PartitionManager

// InfoManager provides version-agnostic cluster information operations
type InfoManager = interfaces.InfoManager

// ReservationManager provides version-agnostic reservation operations
type ReservationManager = interfaces.ReservationManager

// QoSManager provides version-agnostic QoS operations
type QoSManager = interfaces.QoSManager

// AccountManager provides version-agnostic account operations
type AccountManager = interfaces.AccountManager

// UserManager provides version-agnostic user operations
type UserManager = interfaces.UserManager

// Job data structures
// Job represents a SLURM job.
type Job = interfaces.Job

// JobList represents a list of jobs.
type JobList = interfaces.JobList

// JobSubmission contains parameters for submitting a job.
type JobSubmission = interfaces.JobSubmission

// JobSubmitResponse contains the response from a job submission.
type JobSubmitResponse = interfaces.JobSubmitResponse

// JobUpdate contains parameters for updating a job.
type JobUpdate = interfaces.JobUpdate

// JobStep represents a job step.
type JobStep = interfaces.JobStep

// JobStepList represents a list of job steps.
type JobStepList = interfaces.JobStepList

// JobEvent represents an event related to a job.
type JobEvent = interfaces.JobEvent

// Node data structures
// Node represents a compute node.
type Node = interfaces.Node

// NodeList represents a list of nodes.
type NodeList = interfaces.NodeList

// NodeUpdate contains parameters for updating a node.
type NodeUpdate = interfaces.NodeUpdate

// NodeEvent represents an event related to a node.
type NodeEvent = interfaces.NodeEvent

// Partition data structures
// Partition represents a queue partition.
type Partition = interfaces.Partition

// PartitionList represents a list of partitions.
type PartitionList = interfaces.PartitionList

// PartitionUpdate contains parameters for updating a partition.
type PartitionUpdate = interfaces.PartitionUpdate

// PartitionEvent represents an event related to a partition.
type PartitionEvent = interfaces.PartitionEvent

// Cluster information structures
// ClusterInfo contains information about the cluster.
type ClusterInfo = interfaces.ClusterInfo

// ClusterStats contains statistics about the cluster.
type ClusterStats = interfaces.ClusterStats

// APIVersion contains API version information.
type APIVersion = interfaces.APIVersion

// Reservation data structures
// Reservation represents a resource reservation.
type Reservation = interfaces.Reservation

// ReservationList represents a list of reservations.
type ReservationList = interfaces.ReservationList

// ReservationCreate contains parameters for creating a reservation.
type ReservationCreate = interfaces.ReservationCreate

// ReservationCreateResponse contains the response from creating a reservation.
type ReservationCreateResponse = interfaces.ReservationCreateResponse

// ReservationUpdate contains parameters for updating a reservation.
type ReservationUpdate = interfaces.ReservationUpdate

// QoS data structures
// QoS represents a Quality of Service configuration.
type QoS = interfaces.QoS

// QoSList represents a list of QoS configurations.
type QoSList = interfaces.QoSList

// QoSCreate contains parameters for creating a QoS.
type QoSCreate = interfaces.QoSCreate

// QoSCreateResponse contains the response from creating a QoS.
type QoSCreateResponse = interfaces.QoSCreateResponse

// QoSUpdate contains parameters for updating a QoS.
type QoSUpdate = interfaces.QoSUpdate

// Account data structures
// Account represents a user account.
type Account = interfaces.Account

// AccountList represents a list of accounts.
type AccountList = interfaces.AccountList

// AccountCreate contains parameters for creating an account.
type AccountCreate = interfaces.AccountCreate

// AccountCreateResponse contains the response from creating an account.
type AccountCreateResponse = interfaces.AccountCreateResponse

// AccountUpdate contains parameters for updating an account.
type AccountUpdate = interfaces.AccountUpdate

// AccountQuota represents quota information for an account.
type AccountQuota = interfaces.AccountQuota

// AccountUsage represents usage information for an account.
type AccountUsage = interfaces.AccountUsage

// AccountHierarchy represents the hierarchical structure of accounts.
type AccountHierarchy = interfaces.AccountHierarchy

// User data structures
// User represents a SLURM user.
type User = interfaces.User

// UserList represents a list of users.
type UserList = interfaces.UserList

// UserAccount represents a user's association with an account.
type UserAccount = interfaces.UserAccount

// UserAssociation represents a user's association with account and QoS.
type UserAssociation = interfaces.UserAssociation

// UserQuota represents quota information for a user.
type UserQuota = interfaces.UserQuota

// UserAccountQuota represents quota information for a user-account pair.
type UserAccountQuota = interfaces.UserAccountQuota

// UserUsage represents usage information for a user.
type UserUsage = interfaces.UserUsage

// AccountUsageStats represents usage statistics for an account.
type AccountUsageStats = interfaces.AccountUsageStats

// UserFairShare represents fair-share information for a user.
type UserFairShare = interfaces.UserFairShare

// FairShareNode represents a node in the fair-share hierarchy.
type FairShareNode = interfaces.FairShareNode

// JobPriorityFactors represents the factors contributing to job priority.
type JobPriorityFactors = interfaces.JobPriorityFactors

// PriorityWeights represents the weights used in priority calculations.
type PriorityWeights = interfaces.PriorityWeights

// JobPriorityInfo represents priority information for a job.
type JobPriorityInfo = interfaces.JobPriorityInfo

// AssociationUsage represents usage information for an association.
type AssociationUsage = interfaces.AssociationUsage

// QoSLimits represents limit information for a QoS.
type QoSLimits = interfaces.QoSLimits

// UserAccountAssociation represents a user's association with an account.
type UserAccountAssociation = interfaces.UserAccountAssociation

// UserAccessValidation represents the result of user access validation.
type UserAccessValidation = interfaces.UserAccessValidation

// AccountFairShare represents fair-share information for an account.
type AccountFairShare = interfaces.AccountFairShare

// FairShareHierarchy represents the hierarchical fair-share structure.
type FairShareHierarchy = interfaces.FairShareHierarchy

// List options
// ListJobsOptions contains options for listing jobs.
type ListJobsOptions = interfaces.ListJobsOptions

// ListNodesOptions contains options for listing nodes.
type ListNodesOptions = interfaces.ListNodesOptions

// ListPartitionsOptions contains options for listing partitions.
type ListPartitionsOptions = interfaces.ListPartitionsOptions

// ListReservationsOptions contains options for listing reservations.
type ListReservationsOptions = interfaces.ListReservationsOptions

// ListQoSOptions contains options for listing QoS configurations.
type ListQoSOptions = interfaces.ListQoSOptions

// ListAccountsOptions contains options for listing accounts.
type ListAccountsOptions = interfaces.ListAccountsOptions

// ListUsersOptions contains options for listing users.
type ListUsersOptions = interfaces.ListUsersOptions

// ListAccountUsersOptions contains options for listing users in an account.
type ListAccountUsersOptions = interfaces.ListAccountUsersOptions

// ListUserAccountAssociationsOptions contains options for listing user-account associations.
type ListUserAccountAssociationsOptions = interfaces.ListUserAccountAssociationsOptions

// Watch options
// WatchJobsOptions contains options for watching job events.
type WatchJobsOptions = interfaces.WatchJobsOptions

// WatchNodesOptions contains options for watching node events.
type WatchNodesOptions = interfaces.WatchNodesOptions

// WatchPartitionsOptions contains options for watching partition events.
type WatchPartitionsOptions = interfaces.WatchPartitionsOptions
