// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package videoanalyzer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a video resource within Azure Video Analyzer. Videos can be ingested from RTSP cameras through live pipelines or can be created by exporting sequences from existing captured video through a pipeline job. Videos ingested through live pipelines can be streamed through Azure Video Analyzer Player Widget or compatible players. Exported videos can be downloaded as MP4 files.
//
// Uses Azure REST API version 2021-11-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-11-01-preview.
type Video struct {
	pulumi.CustomResourceState

	// Video archival properties.
	Archival VideoArchivalResponsePtrOutput `pulumi:"archival"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Set of URLs to the video content.
	ContentUrls VideoContentUrlsResponseOutput `pulumi:"contentUrls"`
	// Optional video description provided by the user. Value can be up to 2048 characters long.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Video flags contain information about the available video actions and its dynamic properties based on the current video state.
	Flags VideoFlagsResponseOutput `pulumi:"flags"`
	// Contains information about the video and audio content.
	MediaInfo VideoMediaInfoResponsePtrOutput `pulumi:"mediaInfo"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Optional video title provided by the user. Value can be up to 256 characters long.
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVideo registers a new resource with the given unique name, arguments, and options.
func NewVideo(ctx *pulumi.Context,
	name string, args *VideoArgs, opts ...pulumi.ResourceOption) (*Video, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20210501preview:Video"),
		},
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:Video"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Video
	err := ctx.RegisterResource("azure-native:videoanalyzer:Video", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVideo gets an existing Video resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVideo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VideoState, opts ...pulumi.ResourceOption) (*Video, error) {
	var resource Video
	err := ctx.ReadResource("azure-native:videoanalyzer:Video", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Video resources.
type videoState struct {
}

type VideoState struct {
}

func (VideoState) ElementType() reflect.Type {
	return reflect.TypeOf((*videoState)(nil)).Elem()
}

type videoArgs struct {
	// The Azure Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// Video archival properties.
	Archival *VideoArchival `pulumi:"archival"`
	// Optional video description provided by the user. Value can be up to 2048 characters long.
	Description *string `pulumi:"description"`
	// Contains information about the video and audio content.
	MediaInfo *VideoMediaInfo `pulumi:"mediaInfo"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Optional video title provided by the user. Value can be up to 256 characters long.
	Title *string `pulumi:"title"`
	// The Video name.
	VideoName *string `pulumi:"videoName"`
}

// The set of arguments for constructing a Video resource.
type VideoArgs struct {
	// The Azure Video Analyzer account name.
	AccountName pulumi.StringInput
	// Video archival properties.
	Archival VideoArchivalPtrInput
	// Optional video description provided by the user. Value can be up to 2048 characters long.
	Description pulumi.StringPtrInput
	// Contains information about the video and audio content.
	MediaInfo VideoMediaInfoPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Optional video title provided by the user. Value can be up to 256 characters long.
	Title pulumi.StringPtrInput
	// The Video name.
	VideoName pulumi.StringPtrInput
}

func (VideoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*videoArgs)(nil)).Elem()
}

type VideoInput interface {
	pulumi.Input

	ToVideoOutput() VideoOutput
	ToVideoOutputWithContext(ctx context.Context) VideoOutput
}

func (*Video) ElementType() reflect.Type {
	return reflect.TypeOf((**Video)(nil)).Elem()
}

func (i *Video) ToVideoOutput() VideoOutput {
	return i.ToVideoOutputWithContext(context.Background())
}

func (i *Video) ToVideoOutputWithContext(ctx context.Context) VideoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoOutput)
}

type VideoOutput struct{ *pulumi.OutputState }

func (VideoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Video)(nil)).Elem()
}

func (o VideoOutput) ToVideoOutput() VideoOutput {
	return o
}

func (o VideoOutput) ToVideoOutputWithContext(ctx context.Context) VideoOutput {
	return o
}

// Video archival properties.
func (o VideoOutput) Archival() VideoArchivalResponsePtrOutput {
	return o.ApplyT(func(v *Video) VideoArchivalResponsePtrOutput { return v.Archival }).(VideoArchivalResponsePtrOutput)
}

// The Azure API version of the resource.
func (o VideoOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Video) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Set of URLs to the video content.
func (o VideoOutput) ContentUrls() VideoContentUrlsResponseOutput {
	return o.ApplyT(func(v *Video) VideoContentUrlsResponseOutput { return v.ContentUrls }).(VideoContentUrlsResponseOutput)
}

// Optional video description provided by the user. Value can be up to 2048 characters long.
func (o VideoOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Video) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Video flags contain information about the available video actions and its dynamic properties based on the current video state.
func (o VideoOutput) Flags() VideoFlagsResponseOutput {
	return o.ApplyT(func(v *Video) VideoFlagsResponseOutput { return v.Flags }).(VideoFlagsResponseOutput)
}

// Contains information about the video and audio content.
func (o VideoOutput) MediaInfo() VideoMediaInfoResponsePtrOutput {
	return o.ApplyT(func(v *Video) VideoMediaInfoResponsePtrOutput { return v.MediaInfo }).(VideoMediaInfoResponsePtrOutput)
}

// The name of the resource
func (o VideoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Video) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VideoOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Video) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Optional video title provided by the user. Value can be up to 256 characters long.
func (o VideoOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Video) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VideoOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Video) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VideoOutput{})
}
