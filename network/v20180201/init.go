// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk"
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
	case "azure-native:network/v20180201:ApplicationGateway":
		r = &ApplicationGateway{}
	case "azure-native:network/v20180201:ApplicationSecurityGroup":
		r = &ApplicationSecurityGroup{}
	case "azure-native:network/v20180201:ConnectionMonitor":
		r = &ConnectionMonitor{}
	case "azure-native:network/v20180201:DdosProtectionPlan":
		r = &DdosProtectionPlan{}
	case "azure-native:network/v20180201:Endpoint":
		r = &Endpoint{}
	case "azure-native:network/v20180201:ExpressRouteCircuit":
		r = &ExpressRouteCircuit{}
	case "azure-native:network/v20180201:ExpressRouteCircuitAuthorization":
		r = &ExpressRouteCircuitAuthorization{}
	case "azure-native:network/v20180201:ExpressRouteCircuitConnection":
		r = &ExpressRouteCircuitConnection{}
	case "azure-native:network/v20180201:ExpressRouteCircuitPeering":
		r = &ExpressRouteCircuitPeering{}
	case "azure-native:network/v20180201:ExpressRouteCrossConnectionPeering":
		r = &ExpressRouteCrossConnectionPeering{}
	case "azure-native:network/v20180201:InboundNatRule":
		r = &InboundNatRule{}
	case "azure-native:network/v20180201:LoadBalancer":
		r = &LoadBalancer{}
	case "azure-native:network/v20180201:LocalNetworkGateway":
		r = &LocalNetworkGateway{}
	case "azure-native:network/v20180201:NetworkInterface":
		r = &NetworkInterface{}
	case "azure-native:network/v20180201:NetworkSecurityGroup":
		r = &NetworkSecurityGroup{}
	case "azure-native:network/v20180201:NetworkWatcher":
		r = &NetworkWatcher{}
	case "azure-native:network/v20180201:PacketCapture":
		r = &PacketCapture{}
	case "azure-native:network/v20180201:Profile":
		r = &Profile{}
	case "azure-native:network/v20180201:PublicIPAddress":
		r = &PublicIPAddress{}
	case "azure-native:network/v20180201:Route":
		r = &Route{}
	case "azure-native:network/v20180201:RouteFilter":
		r = &RouteFilter{}
	case "azure-native:network/v20180201:RouteFilterRule":
		r = &RouteFilterRule{}
	case "azure-native:network/v20180201:RouteTable":
		r = &RouteTable{}
	case "azure-native:network/v20180201:SecurityRule":
		r = &SecurityRule{}
	case "azure-native:network/v20180201:Subnet":
		r = &Subnet{}
	case "azure-native:network/v20180201:VirtualNetwork":
		r = &VirtualNetwork{}
	case "azure-native:network/v20180201:VirtualNetworkGateway":
		r = &VirtualNetworkGateway{}
	case "azure-native:network/v20180201:VirtualNetworkGatewayConnection":
		r = &VirtualNetworkGatewayConnection{}
	case "azure-native:network/v20180201:VirtualNetworkPeering":
		r = &VirtualNetworkPeering{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := pulumiazurenativesdk.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"network/v20180201",
		&module{version},
	)
}