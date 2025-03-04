// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOnlineEndpoint(ctx *pulumi.Context, args *LookupOnlineEndpointArgs, opts ...pulumi.InvokeOption) (*LookupOnlineEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOnlineEndpointResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getOnlineEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOnlineEndpointArgs struct {
	// Online Endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupOnlineEndpointResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Service identity associated with a resource.
	Identity *ResourceIdentityResponse `pulumi:"identity"`
	// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// [Required] Additional attributes of the entity.
	Properties OnlineEndpointResponse `pulumi:"properties"`
	// System data associated with resource provider
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupOnlineEndpointOutput(ctx *pulumi.Context, args LookupOnlineEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupOnlineEndpointResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOnlineEndpointResultOutput, error) {
			args := v.(LookupOnlineEndpointArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices/v20210301preview:getOnlineEndpoint", args, LookupOnlineEndpointResultOutput{}, options).(LookupOnlineEndpointResultOutput), nil
		}).(LookupOnlineEndpointResultOutput)
}

type LookupOnlineEndpointOutputArgs struct {
	// Online Endpoint name.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupOnlineEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineEndpointArgs)(nil)).Elem()
}

type LookupOnlineEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupOnlineEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOnlineEndpointResult)(nil)).Elem()
}

func (o LookupOnlineEndpointResultOutput) ToLookupOnlineEndpointResultOutput() LookupOnlineEndpointResultOutput {
	return o
}

func (o LookupOnlineEndpointResultOutput) ToLookupOnlineEndpointResultOutputWithContext(ctx context.Context) LookupOnlineEndpointResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupOnlineEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service identity associated with a resource.
func (o LookupOnlineEndpointResultOutput) Identity() ResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) *ResourceIdentityResponse { return v.Identity }).(ResourceIdentityResponsePtrOutput)
}

// Metadata used by portal/tooling/etc to render different UX experiences for resources of the same type.
func (o LookupOnlineEndpointResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupOnlineEndpointResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOnlineEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o LookupOnlineEndpointResultOutput) Properties() OnlineEndpointResponseOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) OnlineEndpointResponse { return v.Properties }).(OnlineEndpointResponseOutput)
}

// System data associated with resource provider
func (o LookupOnlineEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupOnlineEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOnlineEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOnlineEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOnlineEndpointResultOutput{})
}
