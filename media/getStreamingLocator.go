// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the details of a Streaming Locator in the Media Services account
//
// Uses Azure REST API version 2023-01-01.
//
// Other available API versions: 2018-03-30-preview, 2018-06-01-preview, 2018-07-01, 2020-05-01, 2021-06-01, 2021-11-01, 2022-08-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native media [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupStreamingLocator(ctx *pulumi.Context, args *LookupStreamingLocatorArgs, opts ...pulumi.InvokeOption) (*LookupStreamingLocatorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStreamingLocatorResult
	err := ctx.Invoke("azure-native:media:getStreamingLocator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamingLocatorArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName string `pulumi:"streamingLocatorName"`
}

// A Streaming Locator resource
type LookupStreamingLocatorResult struct {
	// Alternative Media ID of this Streaming Locator
	AlternativeMediaId *string `pulumi:"alternativeMediaId"`
	// Asset Name
	AssetName string `pulumi:"assetName"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The ContentKeys used by this Streaming Locator.
	ContentKeys []StreamingLocatorContentKeyResponse `pulumi:"contentKeys"`
	// The creation time of the Streaming Locator.
	Created string `pulumi:"created"`
	// Name of the default ContentKeyPolicy used by this Streaming Locator.
	DefaultContentKeyPolicyName *string `pulumi:"defaultContentKeyPolicyName"`
	// The end time of the Streaming Locator.
	EndTime *string `pulumi:"endTime"`
	// A list of asset or account filters which apply to this streaming locator
	Filters []string `pulumi:"filters"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The start time of the Streaming Locator.
	StartTime *string `pulumi:"startTime"`
	// The StreamingLocatorId of the Streaming Locator.
	StreamingLocatorId *string `pulumi:"streamingLocatorId"`
	// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
	StreamingPolicyName string `pulumi:"streamingPolicyName"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupStreamingLocatorOutput(ctx *pulumi.Context, args LookupStreamingLocatorOutputArgs, opts ...pulumi.InvokeOption) LookupStreamingLocatorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStreamingLocatorResultOutput, error) {
			args := v.(LookupStreamingLocatorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:media:getStreamingLocator", args, LookupStreamingLocatorResultOutput{}, options).(LookupStreamingLocatorResultOutput), nil
		}).(LookupStreamingLocatorResultOutput)
}

type LookupStreamingLocatorOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Streaming Locator name.
	StreamingLocatorName pulumi.StringInput `pulumi:"streamingLocatorName"`
}

func (LookupStreamingLocatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingLocatorArgs)(nil)).Elem()
}

// A Streaming Locator resource
type LookupStreamingLocatorResultOutput struct{ *pulumi.OutputState }

func (LookupStreamingLocatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamingLocatorResult)(nil)).Elem()
}

func (o LookupStreamingLocatorResultOutput) ToLookupStreamingLocatorResultOutput() LookupStreamingLocatorResultOutput {
	return o
}

func (o LookupStreamingLocatorResultOutput) ToLookupStreamingLocatorResultOutputWithContext(ctx context.Context) LookupStreamingLocatorResultOutput {
	return o
}

// Alternative Media ID of this Streaming Locator
func (o LookupStreamingLocatorResultOutput) AlternativeMediaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.AlternativeMediaId }).(pulumi.StringPtrOutput)
}

// Asset Name
func (o LookupStreamingLocatorResultOutput) AssetName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.AssetName }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupStreamingLocatorResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The ContentKeys used by this Streaming Locator.
func (o LookupStreamingLocatorResultOutput) ContentKeys() StreamingLocatorContentKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) []StreamingLocatorContentKeyResponse { return v.ContentKeys }).(StreamingLocatorContentKeyResponseArrayOutput)
}

// The creation time of the Streaming Locator.
func (o LookupStreamingLocatorResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Created }).(pulumi.StringOutput)
}

// Name of the default ContentKeyPolicy used by this Streaming Locator.
func (o LookupStreamingLocatorResultOutput) DefaultContentKeyPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.DefaultContentKeyPolicyName }).(pulumi.StringPtrOutput)
}

// The end time of the Streaming Locator.
func (o LookupStreamingLocatorResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// A list of asset or account filters which apply to this streaming locator
func (o LookupStreamingLocatorResultOutput) Filters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) []string { return v.Filters }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupStreamingLocatorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStreamingLocatorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Name }).(pulumi.StringOutput)
}

// The start time of the Streaming Locator.
func (o LookupStreamingLocatorResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// The StreamingLocatorId of the Streaming Locator.
func (o LookupStreamingLocatorResultOutput) StreamingLocatorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) *string { return v.StreamingLocatorId }).(pulumi.StringPtrOutput)
}

// Name of the Streaming Policy used by this Streaming Locator. Either specify the name of Streaming Policy you created or use one of the predefined Streaming Policies. The predefined Streaming Policies available are: 'Predefined_DownloadOnly', 'Predefined_ClearStreamingOnly', 'Predefined_DownloadAndClearStreaming', 'Predefined_ClearKey', 'Predefined_MultiDrmCencStreaming' and 'Predefined_MultiDrmStreaming'
func (o LookupStreamingLocatorResultOutput) StreamingPolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.StreamingPolicyName }).(pulumi.StringOutput)
}

// The system metadata relating to this resource.
func (o LookupStreamingLocatorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStreamingLocatorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamingLocatorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamingLocatorResultOutput{})
}
