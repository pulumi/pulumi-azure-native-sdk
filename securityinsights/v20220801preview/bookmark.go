// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a bookmark in Azure Security Insights.
type Bookmark struct {
	pulumi.CustomResourceState

	// The time the bookmark was created
	Created pulumi.StringPtrOutput `pulumi:"created"`
	// Describes a user that created the bookmark
	CreatedBy UserInfoResponsePtrOutput `pulumi:"createdBy"`
	// The display name of the bookmark
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Describes the entity mappings of the bookmark
	EntityMappings BookmarkEntityMappingsResponseArrayOutput `pulumi:"entityMappings"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The bookmark event time
	EventTime pulumi.StringPtrOutput `pulumi:"eventTime"`
	// Describes an incident that relates to bookmark
	IncidentInfo IncidentInfoResponsePtrOutput `pulumi:"incidentInfo"`
	// List of labels relevant to this bookmark
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The notes of the bookmark
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The query of the bookmark.
	Query pulumi.StringOutput `pulumi:"query"`
	// The end time for the query
	QueryEndTime pulumi.StringPtrOutput `pulumi:"queryEndTime"`
	// The query result of the bookmark.
	QueryResult pulumi.StringPtrOutput `pulumi:"queryResult"`
	// The start time for the query
	QueryStartTime pulumi.StringPtrOutput `pulumi:"queryStartTime"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// A list of relevant mitre attacks
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// A list of relevant mitre techniques
	Techniques pulumi.StringArrayOutput `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The last time the bookmark was updated
	Updated pulumi.StringPtrOutput `pulumi:"updated"`
	// Describes a user that updated the bookmark
	UpdatedBy UserInfoResponsePtrOutput `pulumi:"updatedBy"`
}

// NewBookmark registers a new resource with the given unique name, arguments, and options.
func NewBookmark(ctx *pulumi.Context,
	name string, args *BookmarkArgs, opts ...pulumi.ResourceOption) (*Bookmark, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Bookmark"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Bookmark"),
		},
	})
	opts = append(opts, aliases)
	var resource Bookmark
	err := ctx.RegisterResource("azure-native:securityinsights/v20220801preview:Bookmark", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBookmark gets an existing Bookmark resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBookmark(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkState, opts ...pulumi.ResourceOption) (*Bookmark, error) {
	var resource Bookmark
	err := ctx.ReadResource("azure-native:securityinsights/v20220801preview:Bookmark", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bookmark resources.
type bookmarkState struct {
}

type BookmarkState struct {
}

func (BookmarkState) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkState)(nil)).Elem()
}

type bookmarkArgs struct {
	// Bookmark ID
	BookmarkId *string `pulumi:"bookmarkId"`
	// The time the bookmark was created
	Created *string `pulumi:"created"`
	// Describes a user that created the bookmark
	CreatedBy *UserInfo `pulumi:"createdBy"`
	// The display name of the bookmark
	DisplayName string `pulumi:"displayName"`
	// Describes the entity mappings of the bookmark
	EntityMappings []BookmarkEntityMappings `pulumi:"entityMappings"`
	// The bookmark event time
	EventTime *string `pulumi:"eventTime"`
	// Describes an incident that relates to bookmark
	IncidentInfo *IncidentInfo `pulumi:"incidentInfo"`
	// List of labels relevant to this bookmark
	Labels []string `pulumi:"labels"`
	// The notes of the bookmark
	Notes *string `pulumi:"notes"`
	// The query of the bookmark.
	Query string `pulumi:"query"`
	// The end time for the query
	QueryEndTime *string `pulumi:"queryEndTime"`
	// The query result of the bookmark.
	QueryResult *string `pulumi:"queryResult"`
	// The start time for the query
	QueryStartTime *string `pulumi:"queryStartTime"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A list of relevant mitre attacks
	Tactics []string `pulumi:"tactics"`
	// A list of relevant mitre techniques
	Techniques []string `pulumi:"techniques"`
	// The last time the bookmark was updated
	Updated *string `pulumi:"updated"`
	// Describes a user that updated the bookmark
	UpdatedBy *UserInfo `pulumi:"updatedBy"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Bookmark resource.
type BookmarkArgs struct {
	// Bookmark ID
	BookmarkId pulumi.StringPtrInput
	// The time the bookmark was created
	Created pulumi.StringPtrInput
	// Describes a user that created the bookmark
	CreatedBy UserInfoPtrInput
	// The display name of the bookmark
	DisplayName pulumi.StringInput
	// Describes the entity mappings of the bookmark
	EntityMappings BookmarkEntityMappingsArrayInput
	// The bookmark event time
	EventTime pulumi.StringPtrInput
	// Describes an incident that relates to bookmark
	IncidentInfo IncidentInfoPtrInput
	// List of labels relevant to this bookmark
	Labels pulumi.StringArrayInput
	// The notes of the bookmark
	Notes pulumi.StringPtrInput
	// The query of the bookmark.
	Query pulumi.StringInput
	// The end time for the query
	QueryEndTime pulumi.StringPtrInput
	// The query result of the bookmark.
	QueryResult pulumi.StringPtrInput
	// The start time for the query
	QueryStartTime pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// A list of relevant mitre attacks
	Tactics pulumi.StringArrayInput
	// A list of relevant mitre techniques
	Techniques pulumi.StringArrayInput
	// The last time the bookmark was updated
	Updated pulumi.StringPtrInput
	// Describes a user that updated the bookmark
	UpdatedBy UserInfoPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (BookmarkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkArgs)(nil)).Elem()
}

type BookmarkInput interface {
	pulumi.Input

	ToBookmarkOutput() BookmarkOutput
	ToBookmarkOutputWithContext(ctx context.Context) BookmarkOutput
}

func (*Bookmark) ElementType() reflect.Type {
	return reflect.TypeOf((**Bookmark)(nil)).Elem()
}

func (i *Bookmark) ToBookmarkOutput() BookmarkOutput {
	return i.ToBookmarkOutputWithContext(context.Background())
}

func (i *Bookmark) ToBookmarkOutputWithContext(ctx context.Context) BookmarkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkOutput)
}

type BookmarkOutput struct{ *pulumi.OutputState }

func (BookmarkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bookmark)(nil)).Elem()
}

func (o BookmarkOutput) ToBookmarkOutput() BookmarkOutput {
	return o
}

func (o BookmarkOutput) ToBookmarkOutputWithContext(ctx context.Context) BookmarkOutput {
	return o
}

// The time the bookmark was created
func (o BookmarkOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.Created }).(pulumi.StringPtrOutput)
}

// Describes a user that created the bookmark
func (o BookmarkOutput) CreatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v *Bookmark) UserInfoResponsePtrOutput { return v.CreatedBy }).(UserInfoResponsePtrOutput)
}

// The display name of the bookmark
func (o BookmarkOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Describes the entity mappings of the bookmark
func (o BookmarkOutput) EntityMappings() BookmarkEntityMappingsResponseArrayOutput {
	return o.ApplyT(func(v *Bookmark) BookmarkEntityMappingsResponseArrayOutput { return v.EntityMappings }).(BookmarkEntityMappingsResponseArrayOutput)
}

// Etag of the azure resource
func (o BookmarkOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The bookmark event time
func (o BookmarkOutput) EventTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.EventTime }).(pulumi.StringPtrOutput)
}

// Describes an incident that relates to bookmark
func (o BookmarkOutput) IncidentInfo() IncidentInfoResponsePtrOutput {
	return o.ApplyT(func(v *Bookmark) IncidentInfoResponsePtrOutput { return v.IncidentInfo }).(IncidentInfoResponsePtrOutput)
}

// List of labels relevant to this bookmark
func (o BookmarkOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o BookmarkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notes of the bookmark
func (o BookmarkOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The query of the bookmark.
func (o BookmarkOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// The end time for the query
func (o BookmarkOutput) QueryEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.QueryEndTime }).(pulumi.StringPtrOutput)
}

// The query result of the bookmark.
func (o BookmarkOutput) QueryResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.QueryResult }).(pulumi.StringPtrOutput)
}

// The start time for the query
func (o BookmarkOutput) QueryStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.QueryStartTime }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BookmarkOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Bookmark) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of relevant mitre attacks
func (o BookmarkOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

// A list of relevant mitre techniques
func (o BookmarkOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BookmarkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The last time the bookmark was updated
func (o BookmarkOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bookmark) pulumi.StringPtrOutput { return v.Updated }).(pulumi.StringPtrOutput)
}

// Describes a user that updated the bookmark
func (o BookmarkOutput) UpdatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v *Bookmark) UserInfoResponsePtrOutput { return v.UpdatedBy }).(UserInfoResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BookmarkOutput{})
}