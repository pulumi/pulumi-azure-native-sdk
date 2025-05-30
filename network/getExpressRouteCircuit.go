// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified express route circuit.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupExpressRouteCircuit(ctx *pulumi.Context, args *LookupExpressRouteCircuitArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRouteCircuitResult
	err := ctx.Invoke("azure-native:network:getExpressRouteCircuit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitArgs struct {
	// The name of express route circuit.
	CircuitName string `pulumi:"circuitName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ExpressRouteCircuit resource.
type LookupExpressRouteCircuitResult struct {
	// Allow classic operations.
	AllowClassicOperations *bool `pulumi:"allowClassicOperations"`
	// The authorizationKey.
	AuthorizationKey *string `pulumi:"authorizationKey"`
	// The authorization status of the Circuit.
	AuthorizationStatus string `pulumi:"authorizationStatus"`
	// The list of authorizations.
	Authorizations []ExpressRouteCircuitAuthorizationResponse `pulumi:"authorizations"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
	BandwidthInGbps *float64 `pulumi:"bandwidthInGbps"`
	// The CircuitProvisioningState state of the resource.
	CircuitProvisioningState *string `pulumi:"circuitProvisioningState"`
	// Flag denoting rate-limiting status of the ExpressRoute direct-port circuit.
	EnableDirectPortRateLimit *bool `pulumi:"enableDirectPortRateLimit"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
	ExpressRoutePort *SubResourceResponse `pulumi:"expressRoutePort"`
	// The GatewayManager Etag.
	GatewayManagerEtag *string `pulumi:"gatewayManagerEtag"`
	// Flag denoting global reach status.
	GlobalReachEnabled *bool `pulumi:"globalReachEnabled"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The list of peerings.
	Peerings []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	// The provisioning state of the express route circuit resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The ServiceKey.
	ServiceKey *string `pulumi:"serviceKey"`
	// The ServiceProviderNotes.
	ServiceProviderNotes *string `pulumi:"serviceProviderNotes"`
	// The ServiceProviderProperties.
	ServiceProviderProperties *ExpressRouteCircuitServiceProviderPropertiesResponse `pulumi:"serviceProviderProperties"`
	// The ServiceProviderProvisioningState state of the resource.
	ServiceProviderProvisioningState *string `pulumi:"serviceProviderProvisioningState"`
	// The SKU.
	Sku *ExpressRouteCircuitSkuResponse `pulumi:"sku"`
	// The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
	Stag int `pulumi:"stag"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupExpressRouteCircuitOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitResultOutput, error) {
			args := v.(LookupExpressRouteCircuitArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getExpressRouteCircuit", args, LookupExpressRouteCircuitResultOutput{}, options).(LookupExpressRouteCircuitResultOutput), nil
		}).(LookupExpressRouteCircuitResultOutput)
}

type LookupExpressRouteCircuitOutputArgs struct {
	// The name of express route circuit.
	CircuitName pulumi.StringInput `pulumi:"circuitName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitArgs)(nil)).Elem()
}

// ExpressRouteCircuit resource.
type LookupExpressRouteCircuitResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitResultOutput) ToLookupExpressRouteCircuitResultOutput() LookupExpressRouteCircuitResultOutput {
	return o
}

func (o LookupExpressRouteCircuitResultOutput) ToLookupExpressRouteCircuitResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitResultOutput {
	return o
}

// Allow classic operations.
func (o LookupExpressRouteCircuitResultOutput) AllowClassicOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *bool { return v.AllowClassicOperations }).(pulumi.BoolPtrOutput)
}

// The authorizationKey.
func (o LookupExpressRouteCircuitResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

// The authorization status of the Circuit.
func (o LookupExpressRouteCircuitResultOutput) AuthorizationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.AuthorizationStatus }).(pulumi.StringOutput)
}

// The list of authorizations.
func (o LookupExpressRouteCircuitResultOutput) Authorizations() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) []ExpressRouteCircuitAuthorizationResponse {
		return v.Authorizations
	}).(ExpressRouteCircuitAuthorizationResponseArrayOutput)
}

// The Azure API version of the resource.
func (o LookupExpressRouteCircuitResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
func (o LookupExpressRouteCircuitResultOutput) BandwidthInGbps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *float64 { return v.BandwidthInGbps }).(pulumi.Float64PtrOutput)
}

// The CircuitProvisioningState state of the resource.
func (o LookupExpressRouteCircuitResultOutput) CircuitProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.CircuitProvisioningState }).(pulumi.StringPtrOutput)
}

// Flag denoting rate-limiting status of the ExpressRoute direct-port circuit.
func (o LookupExpressRouteCircuitResultOutput) EnableDirectPortRateLimit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *bool { return v.EnableDirectPortRateLimit }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRouteCircuitResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
func (o LookupExpressRouteCircuitResultOutput) ExpressRoutePort() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *SubResourceResponse { return v.ExpressRoutePort }).(SubResourceResponsePtrOutput)
}

// The GatewayManager Etag.
func (o LookupExpressRouteCircuitResultOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

// Flag denoting global reach status.
func (o LookupExpressRouteCircuitResultOutput) GlobalReachEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *bool { return v.GlobalReachEnabled }).(pulumi.BoolPtrOutput)
}

// Resource ID.
func (o LookupExpressRouteCircuitResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupExpressRouteCircuitResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o LookupExpressRouteCircuitResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of peerings.
func (o LookupExpressRouteCircuitResultOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) []ExpressRouteCircuitPeeringResponse { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

// The provisioning state of the express route circuit resource.
func (o LookupExpressRouteCircuitResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The ServiceKey.
func (o LookupExpressRouteCircuitResultOutput) ServiceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.ServiceKey }).(pulumi.StringPtrOutput)
}

// The ServiceProviderNotes.
func (o LookupExpressRouteCircuitResultOutput) ServiceProviderNotes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.ServiceProviderNotes }).(pulumi.StringPtrOutput)
}

// The ServiceProviderProperties.
func (o LookupExpressRouteCircuitResultOutput) ServiceProviderProperties() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *ExpressRouteCircuitServiceProviderPropertiesResponse {
		return v.ServiceProviderProperties
	}).(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput)
}

// The ServiceProviderProvisioningState state of the resource.
func (o LookupExpressRouteCircuitResultOutput) ServiceProviderProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.ServiceProviderProvisioningState }).(pulumi.StringPtrOutput)
}

// The SKU.
func (o LookupExpressRouteCircuitResultOutput) Sku() ExpressRouteCircuitSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *ExpressRouteCircuitSkuResponse { return v.Sku }).(ExpressRouteCircuitSkuResponsePtrOutput)
}

// The identifier of the circuit traffic. Outer tag for QinQ encapsulation.
func (o LookupExpressRouteCircuitResultOutput) Stag() pulumi.IntOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) int { return v.Stag }).(pulumi.IntOutput)
}

// Resource tags.
func (o LookupExpressRouteCircuitResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupExpressRouteCircuitResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitResultOutput{})
}
