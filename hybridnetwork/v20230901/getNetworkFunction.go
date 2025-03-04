// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified network function resource.
func LookupNetworkFunction(ctx *pulumi.Context, args *LookupNetworkFunctionArgs, opts ...pulumi.InvokeOption) (*LookupNetworkFunctionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkFunctionResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20230901:getNetworkFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkFunctionArgs struct {
	// The name of the network function resource.
	NetworkFunctionName string `pulumi:"networkFunctionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Network function resource response.
type LookupNetworkFunctionResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The managed identity of the network function.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Network function properties.
	Properties interface{} `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupNetworkFunctionOutput(ctx *pulumi.Context, args LookupNetworkFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkFunctionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkFunctionResultOutput, error) {
			args := v.(LookupNetworkFunctionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridnetwork/v20230901:getNetworkFunction", args, LookupNetworkFunctionResultOutput{}, options).(LookupNetworkFunctionResultOutput), nil
		}).(LookupNetworkFunctionResultOutput)
}

type LookupNetworkFunctionOutputArgs struct {
	// The name of the network function resource.
	NetworkFunctionName pulumi.StringInput `pulumi:"networkFunctionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFunctionArgs)(nil)).Elem()
}

// Network function resource response.
type LookupNetworkFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkFunctionResult)(nil)).Elem()
}

func (o LookupNetworkFunctionResultOutput) ToLookupNetworkFunctionResultOutput() LookupNetworkFunctionResultOutput {
	return o
}

func (o LookupNetworkFunctionResultOutput) ToLookupNetworkFunctionResultOutputWithContext(ctx context.Context) LookupNetworkFunctionResultOutput {
	return o
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupNetworkFunctionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupNetworkFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed identity of the network function.
func (o LookupNetworkFunctionResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupNetworkFunctionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupNetworkFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network function properties.
func (o LookupNetworkFunctionResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupNetworkFunctionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupNetworkFunctionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupNetworkFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkFunctionResultOutput{})
}
