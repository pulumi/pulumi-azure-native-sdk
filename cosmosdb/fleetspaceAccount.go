// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure Cosmos DB Fleetspace Account
//
// Uses Azure REST API version 2025-05-01-preview.
type FleetspaceAccount struct {
	pulumi.CustomResourceState

	// The location of  global database account in the Fleetspace Account.
	AccountLocation pulumi.StringPtrOutput `pulumi:"accountLocation"`
	// The resource identifier of global database account in the Fleetspace Account.
	AccountResourceIdentifier pulumi.StringPtrOutput `pulumi:"accountResourceIdentifier"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A provisioning state of the Fleetspace Account.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFleetspaceAccount registers a new resource with the given unique name, arguments, and options.
func NewFleetspaceAccount(ctx *pulumi.Context,
	name string, args *FleetspaceAccountArgs, opts ...pulumi.ResourceOption) (*FleetspaceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FleetName == nil {
		return nil, errors.New("invalid value for required argument 'FleetName'")
	}
	if args.FleetspaceName == nil {
		return nil, errors.New("invalid value for required argument 'FleetspaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250501preview:FleetspaceAccount"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FleetspaceAccount
	err := ctx.RegisterResource("azure-native:cosmosdb:FleetspaceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleetspaceAccount gets an existing FleetspaceAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleetspaceAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetspaceAccountState, opts ...pulumi.ResourceOption) (*FleetspaceAccount, error) {
	var resource FleetspaceAccount
	err := ctx.ReadResource("azure-native:cosmosdb:FleetspaceAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FleetspaceAccount resources.
type fleetspaceAccountState struct {
}

type FleetspaceAccountState struct {
}

func (FleetspaceAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetspaceAccountState)(nil)).Elem()
}

type fleetspaceAccountArgs struct {
	// The location of  global database account in the Fleetspace Account.
	AccountLocation *string `pulumi:"accountLocation"`
	// The resource identifier of global database account in the Fleetspace Account.
	AccountResourceIdentifier *string `pulumi:"accountResourceIdentifier"`
	// Cosmos DB fleet name. Needs to be unique under a subscription.
	FleetName string `pulumi:"fleetName"`
	// Cosmos DB fleetspace account name.
	FleetspaceAccountName *string `pulumi:"fleetspaceAccountName"`
	// Cosmos DB fleetspace name. Needs to be unique under a fleet.
	FleetspaceName string `pulumi:"fleetspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a FleetspaceAccount resource.
type FleetspaceAccountArgs struct {
	// The location of  global database account in the Fleetspace Account.
	AccountLocation pulumi.StringPtrInput
	// The resource identifier of global database account in the Fleetspace Account.
	AccountResourceIdentifier pulumi.StringPtrInput
	// Cosmos DB fleet name. Needs to be unique under a subscription.
	FleetName pulumi.StringInput
	// Cosmos DB fleetspace account name.
	FleetspaceAccountName pulumi.StringPtrInput
	// Cosmos DB fleetspace name. Needs to be unique under a fleet.
	FleetspaceName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (FleetspaceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetspaceAccountArgs)(nil)).Elem()
}

type FleetspaceAccountInput interface {
	pulumi.Input

	ToFleetspaceAccountOutput() FleetspaceAccountOutput
	ToFleetspaceAccountOutputWithContext(ctx context.Context) FleetspaceAccountOutput
}

func (*FleetspaceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetspaceAccount)(nil)).Elem()
}

func (i *FleetspaceAccount) ToFleetspaceAccountOutput() FleetspaceAccountOutput {
	return i.ToFleetspaceAccountOutputWithContext(context.Background())
}

func (i *FleetspaceAccount) ToFleetspaceAccountOutputWithContext(ctx context.Context) FleetspaceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetspaceAccountOutput)
}

type FleetspaceAccountOutput struct{ *pulumi.OutputState }

func (FleetspaceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetspaceAccount)(nil)).Elem()
}

func (o FleetspaceAccountOutput) ToFleetspaceAccountOutput() FleetspaceAccountOutput {
	return o
}

func (o FleetspaceAccountOutput) ToFleetspaceAccountOutputWithContext(ctx context.Context) FleetspaceAccountOutput {
	return o
}

// The location of  global database account in the Fleetspace Account.
func (o FleetspaceAccountOutput) AccountLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetspaceAccount) pulumi.StringPtrOutput { return v.AccountLocation }).(pulumi.StringPtrOutput)
}

// The resource identifier of global database account in the Fleetspace Account.
func (o FleetspaceAccountOutput) AccountResourceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FleetspaceAccount) pulumi.StringPtrOutput { return v.AccountResourceIdentifier }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o FleetspaceAccountOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetspaceAccount) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o FleetspaceAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetspaceAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A provisioning state of the Fleetspace Account.
func (o FleetspaceAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetspaceAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FleetspaceAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FleetspaceAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FleetspaceAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetspaceAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FleetspaceAccountOutput{})
}
