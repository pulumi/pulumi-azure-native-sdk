// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of hybrid runbook worker group.
type HybridRunbookWorkerGroup struct {
	pulumi.CustomResourceState

	// Sets the credential of a worker group.
	Credential RunAsCredentialAssociationPropertyResponsePtrOutput `pulumi:"credential"`
	// Type of the HybridWorkerGroup.
	GroupType pulumi.StringPtrOutput `pulumi:"groupType"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource system metadata.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHybridRunbookWorkerGroup registers a new resource with the given unique name, arguments, and options.
func NewHybridRunbookWorkerGroup(ctx *pulumi.Context,
	name string, args *HybridRunbookWorkerGroupArgs, opts ...pulumi.ResourceOption) (*HybridRunbookWorkerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:HybridRunbookWorkerGroup"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20210622:HybridRunbookWorkerGroup"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220222:HybridRunbookWorkerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridRunbookWorkerGroup
	err := ctx.RegisterResource("azure-native:automation/v20220808:HybridRunbookWorkerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridRunbookWorkerGroup gets an existing HybridRunbookWorkerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridRunbookWorkerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridRunbookWorkerGroupState, opts ...pulumi.ResourceOption) (*HybridRunbookWorkerGroup, error) {
	var resource HybridRunbookWorkerGroup
	err := ctx.ReadResource("azure-native:automation/v20220808:HybridRunbookWorkerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridRunbookWorkerGroup resources.
type hybridRunbookWorkerGroupState struct {
}

type HybridRunbookWorkerGroupState struct {
}

func (HybridRunbookWorkerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridRunbookWorkerGroupState)(nil)).Elem()
}

type hybridRunbookWorkerGroupArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Sets the credential of a worker group.
	Credential *RunAsCredentialAssociationProperty `pulumi:"credential"`
	// The hybrid runbook worker group name
	HybridRunbookWorkerGroupName *string `pulumi:"hybridRunbookWorkerGroupName"`
	// Gets or sets the name of the resource.
	Name *string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a HybridRunbookWorkerGroup resource.
type HybridRunbookWorkerGroupArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Sets the credential of a worker group.
	Credential RunAsCredentialAssociationPropertyPtrInput
	// The hybrid runbook worker group name
	HybridRunbookWorkerGroupName pulumi.StringPtrInput
	// Gets or sets the name of the resource.
	Name pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
}

func (HybridRunbookWorkerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridRunbookWorkerGroupArgs)(nil)).Elem()
}

type HybridRunbookWorkerGroupInput interface {
	pulumi.Input

	ToHybridRunbookWorkerGroupOutput() HybridRunbookWorkerGroupOutput
	ToHybridRunbookWorkerGroupOutputWithContext(ctx context.Context) HybridRunbookWorkerGroupOutput
}

func (*HybridRunbookWorkerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridRunbookWorkerGroup)(nil)).Elem()
}

func (i *HybridRunbookWorkerGroup) ToHybridRunbookWorkerGroupOutput() HybridRunbookWorkerGroupOutput {
	return i.ToHybridRunbookWorkerGroupOutputWithContext(context.Background())
}

func (i *HybridRunbookWorkerGroup) ToHybridRunbookWorkerGroupOutputWithContext(ctx context.Context) HybridRunbookWorkerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridRunbookWorkerGroupOutput)
}

type HybridRunbookWorkerGroupOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridRunbookWorkerGroup)(nil)).Elem()
}

func (o HybridRunbookWorkerGroupOutput) ToHybridRunbookWorkerGroupOutput() HybridRunbookWorkerGroupOutput {
	return o
}

func (o HybridRunbookWorkerGroupOutput) ToHybridRunbookWorkerGroupOutputWithContext(ctx context.Context) HybridRunbookWorkerGroupOutput {
	return o
}

// Sets the credential of a worker group.
func (o HybridRunbookWorkerGroupOutput) Credential() RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *HybridRunbookWorkerGroup) RunAsCredentialAssociationPropertyResponsePtrOutput {
		return v.Credential
	}).(RunAsCredentialAssociationPropertyResponsePtrOutput)
}

// Type of the HybridWorkerGroup.
func (o HybridRunbookWorkerGroupOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridRunbookWorkerGroup) pulumi.StringPtrOutput { return v.GroupType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o HybridRunbookWorkerGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridRunbookWorkerGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource system metadata.
func (o HybridRunbookWorkerGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridRunbookWorkerGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o HybridRunbookWorkerGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridRunbookWorkerGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridRunbookWorkerGroupOutput{})
}