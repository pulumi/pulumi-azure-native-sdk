// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcompute

// Indicates whether machines associated with the private link scope can also use public Azure Arc service endpoints.
type PublicNetworkAccessType string

const (
	// Allows Azure Arc agents to communicate with Azure Arc services over both public (internet) and private endpoints.
	PublicNetworkAccessTypeEnabled = PublicNetworkAccessType("Enabled")
	// Does not allow Azure Arc agents to communicate with Azure Arc services over public (internet) endpoints. The agents must use the private link.
	PublicNetworkAccessTypeDisabled = PublicNetworkAccessType("Disabled")
)

func init() {
}
