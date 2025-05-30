// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The workspace manager member
//
// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WorkspaceManagerMember struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceResourceId pulumi.StringOutput `pulumi:"targetWorkspaceResourceId"`
	// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceTenantId pulumi.StringOutput `pulumi:"targetWorkspaceTenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceManagerMember registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceManagerMember(ctx *pulumi.Context,
	name string, args *WorkspaceManagerMemberArgs, opts ...pulumi.ResourceOption) (*WorkspaceManagerMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetWorkspaceResourceId == nil {
		return nil, errors.New("invalid value for required argument 'TargetWorkspaceResourceId'")
	}
	if args.TargetWorkspaceTenantId == nil {
		return nil, errors.New("invalid value for required argument 'TargetWorkspaceTenantId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:WorkspaceManagerMember"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250401preview:WorkspaceManagerMember"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceManagerMember
	err := ctx.RegisterResource("azure-native:securityinsights:WorkspaceManagerMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceManagerMember gets an existing WorkspaceManagerMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceManagerMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceManagerMemberState, opts ...pulumi.ResourceOption) (*WorkspaceManagerMember, error) {
	var resource WorkspaceManagerMember
	err := ctx.ReadResource("azure-native:securityinsights:WorkspaceManagerMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceManagerMember resources.
type workspaceManagerMemberState struct {
}

type WorkspaceManagerMemberState struct {
}

func (WorkspaceManagerMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerMemberState)(nil)).Elem()
}

type workspaceManagerMemberArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceResourceId string `pulumi:"targetWorkspaceResourceId"`
	// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceTenantId string `pulumi:"targetWorkspaceTenantId"`
	// The name of the workspace manager member
	WorkspaceManagerMemberName *string `pulumi:"workspaceManagerMemberName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceManagerMember resource.
type WorkspaceManagerMemberArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceResourceId pulumi.StringInput
	// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceTenantId pulumi.StringInput
	// The name of the workspace manager member
	WorkspaceManagerMemberName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceManagerMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerMemberArgs)(nil)).Elem()
}

type WorkspaceManagerMemberInput interface {
	pulumi.Input

	ToWorkspaceManagerMemberOutput() WorkspaceManagerMemberOutput
	ToWorkspaceManagerMemberOutputWithContext(ctx context.Context) WorkspaceManagerMemberOutput
}

func (*WorkspaceManagerMember) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerMember)(nil)).Elem()
}

func (i *WorkspaceManagerMember) ToWorkspaceManagerMemberOutput() WorkspaceManagerMemberOutput {
	return i.ToWorkspaceManagerMemberOutputWithContext(context.Background())
}

func (i *WorkspaceManagerMember) ToWorkspaceManagerMemberOutputWithContext(ctx context.Context) WorkspaceManagerMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceManagerMemberOutput)
}

type WorkspaceManagerMemberOutput struct{ *pulumi.OutputState }

func (WorkspaceManagerMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerMember)(nil)).Elem()
}

func (o WorkspaceManagerMemberOutput) ToWorkspaceManagerMemberOutput() WorkspaceManagerMemberOutput {
	return o
}

func (o WorkspaceManagerMemberOutput) ToWorkspaceManagerMemberOutputWithContext(ctx context.Context) WorkspaceManagerMemberOutput {
	return o
}

// The Azure API version of the resource.
func (o WorkspaceManagerMemberOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource Etag.
func (o WorkspaceManagerMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceManagerMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkspaceManagerMemberOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
func (o WorkspaceManagerMemberOutput) TargetWorkspaceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) pulumi.StringOutput { return v.TargetWorkspaceResourceId }).(pulumi.StringOutput)
}

// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
func (o WorkspaceManagerMemberOutput) TargetWorkspaceTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) pulumi.StringOutput { return v.TargetWorkspaceTenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceManagerMemberOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerMember) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceManagerMemberOutput{})
}
