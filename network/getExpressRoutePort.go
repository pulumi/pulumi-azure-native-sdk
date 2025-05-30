// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the requested ExpressRoutePort resource.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupExpressRoutePort(ctx *pulumi.Context, args *LookupExpressRoutePortArgs, opts ...pulumi.InvokeOption) (*LookupExpressRoutePortResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExpressRoutePortResult
	err := ctx.Invoke("azure-native:network:getExpressRoutePort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRoutePortArgs struct {
	// The name of ExpressRoutePort.
	ExpressRoutePortName string `pulumi:"expressRoutePortName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ExpressRoutePort resource definition.
type LookupExpressRoutePortResult struct {
	// Date of the physical port allocation to be used in Letter of Authorization.
	AllocationDate string `pulumi:"allocationDate"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Bandwidth of procured ports in Gbps.
	BandwidthInGbps *int `pulumi:"bandwidthInGbps"`
	// The billing type of the ExpressRoutePort resource.
	BillingType *string `pulumi:"billingType"`
	// Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
	Circuits []SubResourceResponse `pulumi:"circuits"`
	// Encapsulation method on physical ports.
	Encapsulation *string `pulumi:"encapsulation"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Ether type of the physical port.
	EtherType string `pulumi:"etherType"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The identity of ExpressRoutePort, if configured.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The set of physical links of the ExpressRoutePort resource.
	Links []ExpressRouteLinkResponse `pulumi:"links"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Maximum transmission unit of the physical port pair(s).
	Mtu string `pulumi:"mtu"`
	// Resource name.
	Name string `pulumi:"name"`
	// The name of the peering location that the ExpressRoutePort is mapped to physically.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// Aggregate Gbps of associated circuit bandwidths.
	ProvisionedBandwidthInGbps float64 `pulumi:"provisionedBandwidthInGbps"`
	// The provisioning state of the express route port resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The resource GUID property of the express route port resource.
	ResourceGuid string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupExpressRoutePortOutput(ctx *pulumi.Context, args LookupExpressRoutePortOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRoutePortResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExpressRoutePortResultOutput, error) {
			args := v.(LookupExpressRoutePortArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getExpressRoutePort", args, LookupExpressRoutePortResultOutput{}, options).(LookupExpressRoutePortResultOutput), nil
		}).(LookupExpressRoutePortResultOutput)
}

type LookupExpressRoutePortOutputArgs struct {
	// The name of ExpressRoutePort.
	ExpressRoutePortName pulumi.StringInput `pulumi:"expressRoutePortName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRoutePortOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortArgs)(nil)).Elem()
}

// ExpressRoutePort resource definition.
type LookupExpressRoutePortResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRoutePortResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRoutePortResult)(nil)).Elem()
}

func (o LookupExpressRoutePortResultOutput) ToLookupExpressRoutePortResultOutput() LookupExpressRoutePortResultOutput {
	return o
}

func (o LookupExpressRoutePortResultOutput) ToLookupExpressRoutePortResultOutputWithContext(ctx context.Context) LookupExpressRoutePortResultOutput {
	return o
}

// Date of the physical port allocation to be used in Letter of Authorization.
func (o LookupExpressRoutePortResultOutput) AllocationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.AllocationDate }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupExpressRoutePortResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Bandwidth of procured ports in Gbps.
func (o LookupExpressRoutePortResultOutput) BandwidthInGbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *int { return v.BandwidthInGbps }).(pulumi.IntPtrOutput)
}

// The billing type of the ExpressRoutePort resource.
func (o LookupExpressRoutePortResultOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.BillingType }).(pulumi.StringPtrOutput)
}

// Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
func (o LookupExpressRoutePortResultOutput) Circuits() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) []SubResourceResponse { return v.Circuits }).(SubResourceResponseArrayOutput)
}

// Encapsulation method on physical ports.
func (o LookupExpressRoutePortResultOutput) Encapsulation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.Encapsulation }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupExpressRoutePortResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Ether type of the physical port.
func (o LookupExpressRoutePortResultOutput) EtherType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.EtherType }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupExpressRoutePortResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The identity of ExpressRoutePort, if configured.
func (o LookupExpressRoutePortResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The set of physical links of the ExpressRoutePort resource.
func (o LookupExpressRoutePortResultOutput) Links() ExpressRouteLinkResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) []ExpressRouteLinkResponse { return v.Links }).(ExpressRouteLinkResponseArrayOutput)
}

// Resource location.
func (o LookupExpressRoutePortResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Maximum transmission unit of the physical port pair(s).
func (o LookupExpressRoutePortResultOutput) Mtu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Mtu }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupExpressRoutePortResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the peering location that the ExpressRoutePort is mapped to physically.
func (o LookupExpressRoutePortResultOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) *string { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

// Aggregate Gbps of associated circuit bandwidths.
func (o LookupExpressRoutePortResultOutput) ProvisionedBandwidthInGbps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupExpressRoutePortResult) float64 { return v.ProvisionedBandwidthInGbps }).(pulumi.Float64Output)
}

// The provisioning state of the express route port resource.
func (o LookupExpressRoutePortResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The resource GUID property of the express route port resource.
func (o LookupExpressRoutePortResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupExpressRoutePortResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupExpressRoutePortResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRoutePortResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRoutePortResultOutput{})
}
