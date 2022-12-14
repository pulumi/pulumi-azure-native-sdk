// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220505preview

// Managed identity type.
type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned = ManagedIdentityTypes("SystemAssigned")
)

// Flag specifying if the resource monitoring is enabled or disabled.
type MonitoringStatus string

const (
	MonitoringStatusEnabled  = MonitoringStatus("Enabled")
	MonitoringStatusDisabled = MonitoringStatus("Disabled")
)

// Provisioning state of the monitoring tag rules.
type ProvisioningState string

const (
	ProvisioningStateAccepted     = ProvisioningState("Accepted")
	ProvisioningStateCreating     = ProvisioningState("Creating")
	ProvisioningStateUpdating     = ProvisioningState("Updating")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
	ProvisioningStateDeleted      = ProvisioningState("Deleted")
	ProvisioningStateNotSpecified = ProvisioningState("NotSpecified")
)

// Valid actions for a filtering tag.
type TagAction string

const (
	TagActionInclude = TagAction("Include")
	TagActionExclude = TagAction("Exclude")
)

func init() {
}
