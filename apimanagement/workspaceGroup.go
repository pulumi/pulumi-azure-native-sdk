// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contract details.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WorkspaceGroup struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
	BuiltIn pulumi.BoolOutput `pulumi:"builtIn"`
	// Group description. Can contain HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Group name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceGroup registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceGroup(ctx *pulumi.Context,
	name string, args *WorkspaceGroupArgs, opts ...pulumi.ResourceOption) (*WorkspaceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:WorkspaceGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:WorkspaceGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:WorkspaceGroup"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:WorkspaceGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceGroup
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceGroup gets an existing WorkspaceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceGroupState, opts ...pulumi.ResourceOption) (*WorkspaceGroup, error) {
	var resource WorkspaceGroup
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceGroup resources.
type workspaceGroupState struct {
}

type WorkspaceGroupState struct {
}

func (WorkspaceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceGroupState)(nil)).Elem()
}

type workspaceGroupArgs struct {
	// Group description.
	Description *string `pulumi:"description"`
	// Group name.
	DisplayName string `pulumi:"displayName"`
	// Identifier of the external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
	ExternalId *string `pulumi:"externalId"`
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId *string `pulumi:"groupId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Group type.
	Type *GroupType `pulumi:"type"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceGroup resource.
type WorkspaceGroupArgs struct {
	// Group description.
	Description pulumi.StringPtrInput
	// Group name.
	DisplayName pulumi.StringInput
	// Identifier of the external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
	ExternalId pulumi.StringPtrInput
	// Group identifier. Must be unique in the current API Management service instance.
	GroupId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Group type.
	Type GroupTypePtrInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceGroupArgs)(nil)).Elem()
}

type WorkspaceGroupInput interface {
	pulumi.Input

	ToWorkspaceGroupOutput() WorkspaceGroupOutput
	ToWorkspaceGroupOutputWithContext(ctx context.Context) WorkspaceGroupOutput
}

func (*WorkspaceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceGroup)(nil)).Elem()
}

func (i *WorkspaceGroup) ToWorkspaceGroupOutput() WorkspaceGroupOutput {
	return i.ToWorkspaceGroupOutputWithContext(context.Background())
}

func (i *WorkspaceGroup) ToWorkspaceGroupOutputWithContext(ctx context.Context) WorkspaceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceGroupOutput)
}

type WorkspaceGroupOutput struct{ *pulumi.OutputState }

func (WorkspaceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceGroup)(nil)).Elem()
}

func (o WorkspaceGroupOutput) ToWorkspaceGroupOutput() WorkspaceGroupOutput {
	return o
}

func (o WorkspaceGroupOutput) ToWorkspaceGroupOutputWithContext(ctx context.Context) WorkspaceGroupOutput {
	return o
}

// The Azure API version of the resource.
func (o WorkspaceGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// true if the group is one of the three system groups (Administrators, Developers, or Guests); otherwise false.
func (o WorkspaceGroupOutput) BuiltIn() pulumi.BoolOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.BoolOutput { return v.BuiltIn }).(pulumi.BoolOutput)
}

// Group description. Can contain HTML formatting tags.
func (o WorkspaceGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Group name.
func (o WorkspaceGroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// For external groups, this property contains the id of the group from the external identity provider, e.g. for Azure Active Directory `aad://<tenant>.onmicrosoft.com/groups/<group object id>`; otherwise the value is null.
func (o WorkspaceGroupOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.StringPtrOutput { return v.ExternalId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkspaceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceGroupOutput{})
}
