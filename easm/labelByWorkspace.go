// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package easm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Label details
//
// Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.
type LabelByWorkspace struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Label color.
	Color pulumi.StringPtrOutput `pulumi:"color"`
	// Label display name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLabelByWorkspace registers a new resource with the given unique name, arguments, and options.
func NewLabelByWorkspace(ctx *pulumi.Context,
	name string, args *LabelByWorkspaceArgs, opts ...pulumi.ResourceOption) (*LabelByWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:easm/v20220401preview:LabelByWorkspace"),
		},
		{
			Type: pulumi.String("azure-native:easm/v20230401preview:LabelByWorkspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LabelByWorkspace
	err := ctx.RegisterResource("azure-native:easm:LabelByWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabelByWorkspace gets an existing LabelByWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabelByWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelByWorkspaceState, opts ...pulumi.ResourceOption) (*LabelByWorkspace, error) {
	var resource LabelByWorkspace
	err := ctx.ReadResource("azure-native:easm:LabelByWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LabelByWorkspace resources.
type labelByWorkspaceState struct {
}

type LabelByWorkspaceState struct {
}

func (LabelByWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelByWorkspaceState)(nil)).Elem()
}

type labelByWorkspaceArgs struct {
	// Label color.
	Color *string `pulumi:"color"`
	// Label display name.
	DisplayName *string `pulumi:"displayName"`
	// The name of the Label.
	LabelName *string `pulumi:"labelName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LabelByWorkspace resource.
type LabelByWorkspaceArgs struct {
	// Label color.
	Color pulumi.StringPtrInput
	// Label display name.
	DisplayName pulumi.StringPtrInput
	// The name of the Label.
	LabelName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Workspace.
	WorkspaceName pulumi.StringInput
}

func (LabelByWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelByWorkspaceArgs)(nil)).Elem()
}

type LabelByWorkspaceInput interface {
	pulumi.Input

	ToLabelByWorkspaceOutput() LabelByWorkspaceOutput
	ToLabelByWorkspaceOutputWithContext(ctx context.Context) LabelByWorkspaceOutput
}

func (*LabelByWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelByWorkspace)(nil)).Elem()
}

func (i *LabelByWorkspace) ToLabelByWorkspaceOutput() LabelByWorkspaceOutput {
	return i.ToLabelByWorkspaceOutputWithContext(context.Background())
}

func (i *LabelByWorkspace) ToLabelByWorkspaceOutputWithContext(ctx context.Context) LabelByWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelByWorkspaceOutput)
}

type LabelByWorkspaceOutput struct{ *pulumi.OutputState }

func (LabelByWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelByWorkspace)(nil)).Elem()
}

func (o LabelByWorkspaceOutput) ToLabelByWorkspaceOutput() LabelByWorkspaceOutput {
	return o
}

func (o LabelByWorkspaceOutput) ToLabelByWorkspaceOutputWithContext(ctx context.Context) LabelByWorkspaceOutput {
	return o
}

// The Azure API version of the resource.
func (o LabelByWorkspaceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Label color.
func (o LabelByWorkspaceOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringPtrOutput { return v.Color }).(pulumi.StringPtrOutput)
}

// Label display name.
func (o LabelByWorkspaceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LabelByWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource provisioning state.
func (o LabelByWorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LabelByWorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LabelByWorkspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LabelByWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelByWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabelByWorkspaceOutput{})
}
