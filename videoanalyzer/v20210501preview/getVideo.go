// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an existing video resource within an account with a given name.
func LookupVideo(ctx *pulumi.Context, args *LookupVideoArgs, opts ...pulumi.InvokeOption) (*LookupVideoResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVideoResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getVideo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVideoArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the video to retrieve.
	VideoName string `pulumi:"videoName"`
}

// The representation of a single video in a Video Analyzer account.
type LookupVideoResult struct {
	// Optional video description provided by the user. Value can be up to 2048 characters long.
	Description *string `pulumi:"description"`
	// Video flags contain information about the available video actions and its dynamic properties based on the current video state.
	Flags VideoFlagsResponse `pulumi:"flags"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Contains information about the video and audio content.
	MediaInfo VideoMediaInfoResponse `pulumi:"mediaInfo"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Video streaming holds information about video streaming URLs.
	Streaming VideoStreamingResponse `pulumi:"streaming"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Optional video title provided by the user. Value can be up to 256 characters long.
	Title *string `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVideoOutput(ctx *pulumi.Context, args LookupVideoOutputArgs, opts ...pulumi.InvokeOption) LookupVideoResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVideoResultOutput, error) {
			args := v.(LookupVideoArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:videoanalyzer/v20210501preview:getVideo", args, LookupVideoResultOutput{}, options).(LookupVideoResultOutput), nil
		}).(LookupVideoResultOutput)
}

type LookupVideoOutputArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the video to retrieve.
	VideoName pulumi.StringInput `pulumi:"videoName"`
}

func (LookupVideoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoArgs)(nil)).Elem()
}

// The representation of a single video in a Video Analyzer account.
type LookupVideoResultOutput struct{ *pulumi.OutputState }

func (LookupVideoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoResult)(nil)).Elem()
}

func (o LookupVideoResultOutput) ToLookupVideoResultOutput() LookupVideoResultOutput {
	return o
}

func (o LookupVideoResultOutput) ToLookupVideoResultOutputWithContext(ctx context.Context) LookupVideoResultOutput {
	return o
}

// Optional video description provided by the user. Value can be up to 2048 characters long.
func (o LookupVideoResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Video flags contain information about the available video actions and its dynamic properties based on the current video state.
func (o LookupVideoResultOutput) Flags() VideoFlagsResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoFlagsResponse { return v.Flags }).(VideoFlagsResponseOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVideoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Id }).(pulumi.StringOutput)
}

// Contains information about the video and audio content.
func (o LookupVideoResultOutput) MediaInfo() VideoMediaInfoResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoMediaInfoResponse { return v.MediaInfo }).(VideoMediaInfoResponseOutput)
}

// The name of the resource
func (o LookupVideoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Name }).(pulumi.StringOutput)
}

// Video streaming holds information about video streaming URLs.
func (o LookupVideoResultOutput) Streaming() VideoStreamingResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) VideoStreamingResponse { return v.Streaming }).(VideoStreamingResponseOutput)
}

// The system metadata relating to this resource.
func (o LookupVideoResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVideoResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Optional video title provided by the user. Value can be up to 256 characters long.
func (o LookupVideoResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVideoResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVideoResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVideoResultOutput{})
}
