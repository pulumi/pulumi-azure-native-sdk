// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A FluidRelay Server.
func LookupFluidRelayServer(ctx *pulumi.Context, args *LookupFluidRelayServerArgs, opts ...pulumi.InvokeOption) (*LookupFluidRelayServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFluidRelayServerResult
	err := ctx.Invoke("azure-native:fluidrelay/v20210615preview:getFluidRelayServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFluidRelayServerArgs struct {
	// The resource name.
	Name string `pulumi:"name"`
	// The resource group containing the resource.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// A FluidRelay Server.
type LookupFluidRelayServerResult struct {
	// The Fluid Relay Service endpoints for this server.
	FluidRelayEndpoints FluidRelayEndpointsResponse `pulumi:"fluidRelayEndpoints"`
	// The Fluid tenantId for this server
	FrsTenantId string `pulumi:"frsTenantId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The type of identity used for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provision states for FluidRelay RP
	ProvisioningState *string `pulumi:"provisioningState"`
	// System meta data for this resource, including creation and modification information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFluidRelayServerOutput(ctx *pulumi.Context, args LookupFluidRelayServerOutputArgs, opts ...pulumi.InvokeOption) LookupFluidRelayServerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFluidRelayServerResultOutput, error) {
			args := v.(LookupFluidRelayServerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:fluidrelay/v20210615preview:getFluidRelayServer", args, LookupFluidRelayServerResultOutput{}, options).(LookupFluidRelayServerResultOutput), nil
		}).(LookupFluidRelayServerResultOutput)
}

type LookupFluidRelayServerOutputArgs struct {
	// The resource name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group containing the resource.
	ResourceGroup pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupFluidRelayServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluidRelayServerArgs)(nil)).Elem()
}

// A FluidRelay Server.
type LookupFluidRelayServerResultOutput struct{ *pulumi.OutputState }

func (LookupFluidRelayServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFluidRelayServerResult)(nil)).Elem()
}

func (o LookupFluidRelayServerResultOutput) ToLookupFluidRelayServerResultOutput() LookupFluidRelayServerResultOutput {
	return o
}

func (o LookupFluidRelayServerResultOutput) ToLookupFluidRelayServerResultOutputWithContext(ctx context.Context) LookupFluidRelayServerResultOutput {
	return o
}

// The Fluid Relay Service endpoints for this server.
func (o LookupFluidRelayServerResultOutput) FluidRelayEndpoints() FluidRelayEndpointsResponseOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) FluidRelayEndpointsResponse { return v.FluidRelayEndpoints }).(FluidRelayEndpointsResponseOutput)
}

// The Fluid tenantId for this server
func (o LookupFluidRelayServerResultOutput) FrsTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.FrsTenantId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupFluidRelayServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The type of identity used for the resource.
func (o LookupFluidRelayServerResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupFluidRelayServerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFluidRelayServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provision states for FluidRelay RP
func (o LookupFluidRelayServerResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// System meta data for this resource, including creation and modification information.
func (o LookupFluidRelayServerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupFluidRelayServerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFluidRelayServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFluidRelayServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFluidRelayServerResultOutput{})
}
