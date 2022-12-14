// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

// Gets or sets the cookie affinity
type ApplicationGatewayCookieBasedAffinity string

const (
	ApplicationGatewayCookieBasedAffinityEnabled  = ApplicationGatewayCookieBasedAffinity("Enabled")
	ApplicationGatewayCookieBasedAffinityDisabled = ApplicationGatewayCookieBasedAffinity("Disabled")
)

// Gets or sets the protocol
type ApplicationGatewayProtocol string

const (
	ApplicationGatewayProtocolHttp  = ApplicationGatewayProtocol("Http")
	ApplicationGatewayProtocolHttps = ApplicationGatewayProtocol("Https")
)

// Gets or sets the rule type
type ApplicationGatewayRequestRoutingRuleType string

const (
	ApplicationGatewayRequestRoutingRuleTypeBasic = ApplicationGatewayRequestRoutingRuleType("Basic")
)

// Gets or sets name of application gateway SKU
type ApplicationGatewaySkuName string

const (
	ApplicationGatewaySkuName_Standard_Small  = ApplicationGatewaySkuName("Standard_Small")
	ApplicationGatewaySkuName_Standard_Medium = ApplicationGatewaySkuName("Standard_Medium")
	ApplicationGatewaySkuName_Standard_Large  = ApplicationGatewaySkuName("Standard_Large")
)

// Gets or sets tier of application gateway
type ApplicationGatewayTier string

const (
	ApplicationGatewayTierStandard = ApplicationGatewayTier("Standard")
)

// Gets or sets AuthorizationUseStatus
type AuthorizationUseStatus string

const (
	AuthorizationUseStatusAvailable = AuthorizationUseStatus("Available")
	AuthorizationUseStatusInUse     = AuthorizationUseStatus("InUse")
)

// Gets or sets AdvertisedPublicPrefixState of the Peering resource
type ExpressRouteCircuitPeeringAdvertisedPublicPrefixState string

const (
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateNotConfigured    = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("NotConfigured")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateConfiguring      = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("Configuring")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateConfigured       = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("Configured")
	ExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValidationNeeded = ExpressRouteCircuitPeeringAdvertisedPublicPrefixState("ValidationNeeded")
)

// Gets or sets state of Peering
type ExpressRouteCircuitPeeringStateEnum string

const (
	ExpressRouteCircuitPeeringStateEnumDisabled = ExpressRouteCircuitPeeringStateEnum("Disabled")
	ExpressRouteCircuitPeeringStateEnumEnabled  = ExpressRouteCircuitPeeringStateEnum("Enabled")
)

// Gets or sets PeeringType
type ExpressRouteCircuitPeeringTypeEnum string

const (
	ExpressRouteCircuitPeeringTypeEnumAzurePublicPeering  = ExpressRouteCircuitPeeringTypeEnum("AzurePublicPeering")
	ExpressRouteCircuitPeeringTypeEnumAzurePrivatePeering = ExpressRouteCircuitPeeringTypeEnum("AzurePrivatePeering")
	ExpressRouteCircuitPeeringTypeEnumMicrosoftPeering    = ExpressRouteCircuitPeeringTypeEnum("MicrosoftPeering")
)

// Gets or sets family of the sku.
type ExpressRouteCircuitSkuFamily string

const (
	ExpressRouteCircuitSkuFamilyUnlimitedData = ExpressRouteCircuitSkuFamily("UnlimitedData")
	ExpressRouteCircuitSkuFamilyMeteredData   = ExpressRouteCircuitSkuFamily("MeteredData")
)

// Gets or sets tier of the sku.
type ExpressRouteCircuitSkuTier string

const (
	ExpressRouteCircuitSkuTierStandard = ExpressRouteCircuitSkuTier("Standard")
	ExpressRouteCircuitSkuTierPremium  = ExpressRouteCircuitSkuTier("Premium")
)

// Gets or sets PublicIP allocation method (Static/Dynamic)
type IpAllocationMethod string

const (
	IpAllocationMethodStatic  = IpAllocationMethod("Static")
	IpAllocationMethodDynamic = IpAllocationMethod("Dynamic")
)

// Gets or sets the load distribution policy for this rule
type LoadDistribution string

const (
	LoadDistributionDefault          = LoadDistribution("Default")
	LoadDistributionSourceIP         = LoadDistribution("SourceIP")
	LoadDistributionSourceIPProtocol = LoadDistribution("SourceIPProtocol")
)

// Gets or sets the protocol of the end point. Possible values are http pr Tcp. If Tcp is specified, a received ACK is required for the probe to be successful. If http is specified,a 200 OK response from the specifies URI is required for the probe to be successful
type ProbeProtocol string

const (
	ProbeProtocolHttp = ProbeProtocol("Http")
	ProbeProtocolTcp  = ProbeProtocol("Tcp")
)

// Gets or sets the type of Azure hop the packet should be sent to.
type RouteNextHopType string

const (
	RouteNextHopTypeVirtualNetworkGateway = RouteNextHopType("VirtualNetworkGateway")
	RouteNextHopTypeVnetLocal             = RouteNextHopType("VnetLocal")
	RouteNextHopTypeInternet              = RouteNextHopType("Internet")
	RouteNextHopTypeVirtualAppliance      = RouteNextHopType("VirtualAppliance")
	RouteNextHopTypeNone                  = RouteNextHopType("None")
)

// Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
type SecurityRuleAccess string

const (
	SecurityRuleAccessAllow = SecurityRuleAccess("Allow")
	SecurityRuleAccessDeny  = SecurityRuleAccess("Deny")
)

// Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
type SecurityRuleDirection string

const (
	SecurityRuleDirectionInbound  = SecurityRuleDirection("Inbound")
	SecurityRuleDirectionOutbound = SecurityRuleDirection("Outbound")
)

// Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
type SecurityRuleProtocol string

const (
	SecurityRuleProtocolTcp      = SecurityRuleProtocol("Tcp")
	SecurityRuleProtocolUdp      = SecurityRuleProtocol("Udp")
	SecurityRuleProtocolAsterisk = SecurityRuleProtocol("*")
)

// Gets or sets ServiceProviderProvisioningState state of the resource
type ServiceProviderProvisioningState string

const (
	ServiceProviderProvisioningStateNotProvisioned = ServiceProviderProvisioningState("NotProvisioned")
	ServiceProviderProvisioningStateProvisioning   = ServiceProviderProvisioningState("Provisioning")
	ServiceProviderProvisioningStateProvisioned    = ServiceProviderProvisioningState("Provisioned")
	ServiceProviderProvisioningStateDeprovisioning = ServiceProviderProvisioningState("Deprovisioning")
)

// Gets or sets the transport protocol for the external endpoint. Possible values are Udp or Tcp
type TransportProtocol string

const (
	TransportProtocolUdp = TransportProtocol("Udp")
	TransportProtocolTcp = TransportProtocol("Tcp")
)

func init() {
}
