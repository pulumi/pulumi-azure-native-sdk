// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001

// The type of identity that last modified the resource.
type CreatedByType string

const (
	CreatedByTypeUser            = CreatedByType("User")
	CreatedByTypeApplication     = CreatedByType("Application")
	CreatedByTypeManagedIdentity = CreatedByType("ManagedIdentity")
	CreatedByTypeKey             = CreatedByType("Key")
)

// Desired level of diagnostic data emitted by the cluster.
type DiagnosticLevel string

const (
	DiagnosticLevelOff      = DiagnosticLevel("Off")
	DiagnosticLevelBasic    = DiagnosticLevel("Basic")
	DiagnosticLevelEnhanced = DiagnosticLevel("Enhanced")
)

// Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                         = ManagedServiceIdentityType("None")
	ManagedServiceIdentityTypeSystemAssigned               = ManagedServiceIdentityType("SystemAssigned")
	ManagedServiceIdentityTypeUserAssigned                 = ManagedServiceIdentityType("UserAssigned")
	ManagedServiceIdentityType_SystemAssigned_UserAssigned = ManagedServiceIdentityType("SystemAssigned, UserAssigned")
)

// Customer Intent for Software Assurance Benefit.
type SoftwareAssuranceIntent string

const (
	SoftwareAssuranceIntentEnable  = SoftwareAssuranceIntent("Enable")
	SoftwareAssuranceIntentDisable = SoftwareAssuranceIntent("Disable")
)

// Status of the Software Assurance for the cluster.
type SoftwareAssuranceStatus string

const (
	SoftwareAssuranceStatusEnabled  = SoftwareAssuranceStatus("Enabled")
	SoftwareAssuranceStatusDisabled = SoftwareAssuranceStatus("Disabled")
)

// Desired state of Windows Server Subscription.
type WindowsServerSubscription string

const (
	WindowsServerSubscriptionDisabled = WindowsServerSubscription("Disabled")
	WindowsServerSubscriptionEnabled  = WindowsServerSubscription("Enabled")
)

func init() {
}
