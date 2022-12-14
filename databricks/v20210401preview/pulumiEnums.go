// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Keyvault
type EncryptionKeySource string

const (
	EncryptionKeySource_Microsoft_Keyvault = EncryptionKeySource("Microsoft.Keyvault")
)

// The encryption keySource (provider). Possible values (case-insensitive):  Default, Microsoft.Keyvault
type KeySource string

const (
	KeySourceDefault             = KeySource("Default")
	KeySource_Microsoft_Keyvault = KeySource("Microsoft.Keyvault")
)

// The status of a private endpoint connection
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
type RequiredNsgRules string

const (
	RequiredNsgRulesAllRules               = RequiredNsgRules("AllRules")
	RequiredNsgRulesNoAzureDatabricksRules = RequiredNsgRules("NoAzureDatabricksRules")
	RequiredNsgRulesNoAzureServiceRules    = RequiredNsgRules("NoAzureServiceRules")
)

func init() {
}
