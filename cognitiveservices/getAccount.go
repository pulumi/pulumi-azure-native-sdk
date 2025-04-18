// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognitiveservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a Cognitive Services account specified by the parameters.
//
// Uses Azure REST API version 2024-10-01.
//
// Other available API versions: 2023-05-01, 2023-10-01-preview, 2024-04-01-preview, 2024-06-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cognitiveservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:cognitiveservices:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	// The name of Cognitive Services account.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
type LookupAccountResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity for the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The Kind of the resource.
	Kind *string `pulumi:"kind"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of Cognitive Services account.
	Properties AccountPropertiesResponse `pulumi:"properties"`
	// The resource model definition representing SKU
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAccountResult
func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAccountResultOutput, error) {
			args := v.(LookupAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cognitiveservices:getAccount", args, LookupAccountResultOutput{}, options).(LookupAccountResultOutput), nil
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	// The name of Cognitive Services account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

// Cognitive Services account is an Azure resource representing the provisioned account, it's type, location and SKU.
type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAccountResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource Etag.
func (o LookupAccountResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity for the resource.
func (o LookupAccountResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The Kind of the resource.
func (o LookupAccountResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupAccountResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of Cognitive Services account.
func (o LookupAccountResultOutput) Properties() AccountPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) AccountPropertiesResponse { return v.Properties }).(AccountPropertiesResponseOutput)
}

// The resource model definition representing SKU
func (o LookupAccountResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
