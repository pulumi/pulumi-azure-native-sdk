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

// The workspace manager assignment
//
// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WorkspaceManagerAssignment struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// List of resources included in this workspace manager assignment
	Items AssignmentItemResponseArrayOutput `pulumi:"items"`
	// The time the last job associated to this assignment ended at
	LastJobEndTime pulumi.StringOutput `pulumi:"lastJobEndTime"`
	// State of the last job associated to this assignment
	LastJobProvisioningState pulumi.StringOutput `pulumi:"lastJobProvisioningState"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The resource name of the workspace manager group targeted by the workspace manager assignment
	TargetResourceName pulumi.StringOutput `pulumi:"targetResourceName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceManagerAssignment registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceManagerAssignment(ctx *pulumi.Context,
	name string, args *WorkspaceManagerAssignmentArgs, opts ...pulumi.ResourceOption) (*WorkspaceManagerAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetResourceName == nil {
		return nil, errors.New("invalid value for required argument 'TargetResourceName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:WorkspaceManagerAssignment"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250401preview:WorkspaceManagerAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceManagerAssignment
	err := ctx.RegisterResource("azure-native:securityinsights:WorkspaceManagerAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceManagerAssignment gets an existing WorkspaceManagerAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceManagerAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceManagerAssignmentState, opts ...pulumi.ResourceOption) (*WorkspaceManagerAssignment, error) {
	var resource WorkspaceManagerAssignment
	err := ctx.ReadResource("azure-native:securityinsights:WorkspaceManagerAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceManagerAssignment resources.
type workspaceManagerAssignmentState struct {
}

type WorkspaceManagerAssignmentState struct {
}

func (WorkspaceManagerAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerAssignmentState)(nil)).Elem()
}

type workspaceManagerAssignmentArgs struct {
	// List of resources included in this workspace manager assignment
	Items []AssignmentItem `pulumi:"items"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the workspace manager group targeted by the workspace manager assignment
	TargetResourceName string `pulumi:"targetResourceName"`
	// The name of the workspace manager assignment
	WorkspaceManagerAssignmentName *string `pulumi:"workspaceManagerAssignmentName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceManagerAssignment resource.
type WorkspaceManagerAssignmentArgs struct {
	// List of resources included in this workspace manager assignment
	Items AssignmentItemArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource name of the workspace manager group targeted by the workspace manager assignment
	TargetResourceName pulumi.StringInput
	// The name of the workspace manager assignment
	WorkspaceManagerAssignmentName pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceManagerAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceManagerAssignmentArgs)(nil)).Elem()
}

type WorkspaceManagerAssignmentInput interface {
	pulumi.Input

	ToWorkspaceManagerAssignmentOutput() WorkspaceManagerAssignmentOutput
	ToWorkspaceManagerAssignmentOutputWithContext(ctx context.Context) WorkspaceManagerAssignmentOutput
}

func (*WorkspaceManagerAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerAssignment)(nil)).Elem()
}

func (i *WorkspaceManagerAssignment) ToWorkspaceManagerAssignmentOutput() WorkspaceManagerAssignmentOutput {
	return i.ToWorkspaceManagerAssignmentOutputWithContext(context.Background())
}

func (i *WorkspaceManagerAssignment) ToWorkspaceManagerAssignmentOutputWithContext(ctx context.Context) WorkspaceManagerAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceManagerAssignmentOutput)
}

type WorkspaceManagerAssignmentOutput struct{ *pulumi.OutputState }

func (WorkspaceManagerAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceManagerAssignment)(nil)).Elem()
}

func (o WorkspaceManagerAssignmentOutput) ToWorkspaceManagerAssignmentOutput() WorkspaceManagerAssignmentOutput {
	return o
}

func (o WorkspaceManagerAssignmentOutput) ToWorkspaceManagerAssignmentOutputWithContext(ctx context.Context) WorkspaceManagerAssignmentOutput {
	return o
}

// The Azure API version of the resource.
func (o WorkspaceManagerAssignmentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource Etag.
func (o WorkspaceManagerAssignmentOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// List of resources included in this workspace manager assignment
func (o WorkspaceManagerAssignmentOutput) Items() AssignmentItemResponseArrayOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) AssignmentItemResponseArrayOutput { return v.Items }).(AssignmentItemResponseArrayOutput)
}

// The time the last job associated to this assignment ended at
func (o WorkspaceManagerAssignmentOutput) LastJobEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.LastJobEndTime }).(pulumi.StringOutput)
}

// State of the last job associated to this assignment
func (o WorkspaceManagerAssignmentOutput) LastJobProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.LastJobProvisioningState }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceManagerAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkspaceManagerAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource name of the workspace manager group targeted by the workspace manager assignment
func (o WorkspaceManagerAssignmentOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.TargetResourceName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceManagerAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceManagerAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceManagerAssignmentOutput{})
}
