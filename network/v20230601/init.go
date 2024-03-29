// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:network/v20230601:AdminRule":
		r = &AdminRule{}
	case "azure-native:network/v20230601:AdminRuleCollection":
		r = &AdminRuleCollection{}
	case "azure-native:network/v20230601:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20230601:ApplicationGatewayPrivateEndpointConnection":
		r = &ApplicationGatewayPrivateEndpointConnection{}
	case "azure-native:network/v20230601:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20230601:AzureFirewall":
		r = &AzureFirewall{}
	case "azure-native:network/v20230601:BastionHost":
		r = &BastionHost{}
	case "azure-native:network/v20230601:ConfigurationPolicyGroup":
		r = &ConfigurationPolicyGroup{}
	case "azure-native:network/v20230601:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20230601:ConnectivityConfiguration":
		r = &ConnectivityConfiguration{}
	case "azure-native:network/v20230601:CustomIPPrefix":
		r = &CustomIPPrefix{}
	case "azure-native:network/v20230601:DdosCustomPolicy":
		r = &DdosCustomPolicy{}
	case "azure-native:network/v20230601:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20230601:DefaultAdminRule":
		r = &DefaultAdminRule{}
	case "azure-native:network/v20230601:DscpConfiguration":
		r = &DscpConfiguration{}
	case "azure-native:network/v20230601:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20230601:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20230601:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20230601:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20230601:ExpressRouteConnection":
		r = &ExpressRouteConnection{}
	case "azure-native:network/v20230601:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20230601:ExpressRouteGateway":
		r = &ExpressRouteGateway{}
	case "azure-native:network/v20230601:ExpressRoutePort":
		r = &ExpressRoutePort{}
	case "azure-native:network/v20230601:ExpressRoutePortAuthorization":
		r = &ExpressRoutePortAuthorization{}
	case "azure-native:network/v20230601:FirewallPolicy":
		r = &FirewallPolicy{}
	case "azure-native:network/v20230601:FirewallPolicyRuleCollectionGroup":
		r = &FirewallPolicyRuleCollectionGroup{}
	case "azure-native:network/v20230601:FlowLog":
		r = &FlowLog{}
	case "azure-native:network/v20230601:HubRouteTable":
		r = &HubRouteTable{}
	case "azure-native:network/v20230601:HubVirtualNetworkConnection":
		r = &HubVirtualNetworkConnection{}
	case "azure-native:network/v20230601:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20230601:IpAllocation":
		r = &IpAllocation{}
	case "azure-native:network/v20230601:IpGroup":
		r = &IpGroup{}
	case "azure-native:network/v20230601:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20230601:LoadBalancerBackendAddressPool":
		r = &LoadBalancerBackendAddressPool{}
	case "azure-native:network/v20230601:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20230601:ManagementGroupNetworkManagerConnection":
		r = &ManagementGroupNetworkManagerConnection{}
	case "azure-native:network/v20230601:NatGateway":
		r = &NatGateway{}
	case "azure-native:network/v20230601:NatRule":
		r = &NatRule{}
	case "azure-native:network/v20230601:NetworkGroup":
		r = &NetworkGroup{}
	case "azure-native:network/v20230601:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20230601:NetworkInterfaceTapConfiguration":
		r = &NetworkInterfaceTapConfiguration{}
	case "azure-native:network/v20230601:NetworkManager":
		r = &NetworkManager{}
	case "azure-native:network/v20230601:NetworkProfile":
		r = &NetworkProfile{}
	case "azure-native:network/v20230601:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20230601:NetworkVirtualAppliance":
		r = &NetworkVirtualAppliance{}
	case "azure-native:network/v20230601:NetworkVirtualApplianceConnection":
		r = &NetworkVirtualApplianceConnection{}
	case "azure-native:network/v20230601:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20230601:P2sVpnGateway":
		r = &P2sVpnGateway{}
	case "azure-native:network/v20230601:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20230601:PrivateDnsZoneGroup":
		r = &PrivateDnsZoneGroup{}
	case "azure-native:network/v20230601:PrivateEndpoint":
		r = &PrivateEndpoint{}
	case "azure-native:network/v20230601:PrivateLinkService":
		r = &PrivateLinkService{}
	case "azure-native:network/v20230601:PrivateLinkServicePrivateEndpointConnection":
		r = &PrivateLinkServicePrivateEndpointConnection{}
	case "azure-native:network/v20230601:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20230601:PublicIPPrefix":
		r = &PublicIPPrefix{}
	case "azure-native:network/v20230601:Route":
		r = &Route{}
	case "azure-native:network/v20230601:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20230601:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20230601:RouteMap":
		r = &RouteMap{}
	case "azure-native:network/v20230601:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20230601:RoutingIntent":
		r = &RoutingIntent{}
	case "azure-native:network/v20230601:ScopeConnection":
		r = &ScopeConnection{}
	case "azure-native:network/v20230601:SecurityAdminConfiguration":
		r = &SecurityAdminConfiguration{}
	case "azure-native:network/v20230601:SecurityPartnerProvider":
		r = &SecurityPartnerProvider{}
	case "azure-native:network/v20230601:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20230601:ServiceEndpointPolicy":
		r = &ServiceEndpointPolicy{}
	case "azure-native:network/v20230601:ServiceEndpointPolicyDefinition":
		r = &ServiceEndpointPolicyDefinition{}
	case "azure-native:network/v20230601:StaticMember":
		r = &StaticMember{}
	case "azure-native:network/v20230601:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20230601:SubscriptionNetworkManagerConnection":
		r = &SubscriptionNetworkManagerConnection{}
	case "azure-native:network/v20230601:VirtualApplianceSite":
		r = &VirtualApplianceSite{}
	case "azure-native:network/v20230601:VirtualHub":
		r = &VirtualHub{}
	case "azure-native:network/v20230601:VirtualHubBgpConnection":
		r = &VirtualHubBgpConnection{}
	case "azure-native:network/v20230601:VirtualHubIpConfiguration":
		r = &VirtualHubIpConfiguration{}
	case "azure-native:network/v20230601:VirtualHubRouteTableV2":
		r = &VirtualHubRouteTableV2{}
	case "azure-native:network/v20230601:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20230601:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20230601:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20230601:VirtualNetworkGatewayNatRule":
		r = &VirtualNetworkGatewayNatRule{}
	case "azure-native:network/v20230601:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	case "azure-native:network/v20230601:VirtualNetworkTap":
		r = &VirtualNetworkTap{}
	case "azure-native:network/v20230601:VirtualRouter":
		r = &VirtualRouter{}
	case "azure-native:network/v20230601:VirtualRouterPeering":
		r = &VirtualRouterPeering{}
	case "azure-native:network/v20230601:VirtualWan":
		r = &VirtualWan{}
	case "azure-native:network/v20230601:VpnConnection":
		r = &VpnConnection{}
	case "azure-native:network/v20230601:VpnGateway":
		r = &VpnGateway{}
	case "azure-native:network/v20230601:VpnServerConfiguration":
		r = &VpnServerConfiguration{}
	case "azure-native:network/v20230601:VpnSite":
		r = &VpnSite{}
	case "azure-native:network/v20230601:WebApplicationFirewallPolicy":
		r = &WebApplicationFirewallPolicy{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"network/v20230601",
		&module{version},
	)
}
