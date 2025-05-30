// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Instance of StaticCidr resource.
//
// Uses Azure REST API version 2024-05-01.
//
// Other available API versions: 2024-01-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupStaticCidr(ctx *pulumi.Context, args *LookupStaticCidrArgs, opts ...pulumi.InvokeOption) (*LookupStaticCidrResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStaticCidrResult
	err := ctx.Invoke("azure-native:network:getStaticCidr", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticCidrArgs struct {
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// Pool resource name.
	PoolName string `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// StaticCidr resource name to retrieve.
	StaticCidrName string `pulumi:"staticCidrName"`
}

// Instance of StaticCidr resource.
type LookupStaticCidrResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of static CIDR resource.
	Properties StaticCidrPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupStaticCidrOutput(ctx *pulumi.Context, args LookupStaticCidrOutputArgs, opts ...pulumi.InvokeOption) LookupStaticCidrResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStaticCidrResultOutput, error) {
			args := v.(LookupStaticCidrArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:getStaticCidr", args, LookupStaticCidrResultOutput{}, options).(LookupStaticCidrResultOutput), nil
		}).(LookupStaticCidrResultOutput)
}

type LookupStaticCidrOutputArgs struct {
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// Pool resource name.
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// StaticCidr resource name to retrieve.
	StaticCidrName pulumi.StringInput `pulumi:"staticCidrName"`
}

func (LookupStaticCidrOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticCidrArgs)(nil)).Elem()
}

// Instance of StaticCidr resource.
type LookupStaticCidrResultOutput struct{ *pulumi.OutputState }

func (LookupStaticCidrResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticCidrResult)(nil)).Elem()
}

func (o LookupStaticCidrResultOutput) ToLookupStaticCidrResultOutput() LookupStaticCidrResultOutput {
	return o
}

func (o LookupStaticCidrResultOutput) ToLookupStaticCidrResultOutputWithContext(ctx context.Context) LookupStaticCidrResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupStaticCidrResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticCidrResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupStaticCidrResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticCidrResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStaticCidrResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticCidrResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of static CIDR resource.
func (o LookupStaticCidrResultOutput) Properties() StaticCidrPropertiesResponseOutput {
	return o.ApplyT(func(v LookupStaticCidrResult) StaticCidrPropertiesResponse { return v.Properties }).(StaticCidrPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStaticCidrResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStaticCidrResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStaticCidrResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticCidrResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticCidrResultOutput{})
}
