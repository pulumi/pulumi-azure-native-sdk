// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sovereign

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Landing zone registration resource type.
//
// Uses Azure REST API version 2025-02-27-preview. In version 2.x of the Azure Native provider, it used API version 2025-02-27-preview.
type LandingZoneRegistrationOperation struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties LandingZoneRegistrationResourcePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLandingZoneRegistrationOperation registers a new resource with the given unique name, arguments, and options.
func NewLandingZoneRegistrationOperation(ctx *pulumi.Context,
	name string, args *LandingZoneRegistrationOperationArgs, opts ...pulumi.ResourceOption) (*LandingZoneRegistrationOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LandingZoneAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LandingZoneAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sovereign/v20250227preview:LandingZoneRegistrationOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LandingZoneRegistrationOperation
	err := ctx.RegisterResource("azure-native:sovereign:LandingZoneRegistrationOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLandingZoneRegistrationOperation gets an existing LandingZoneRegistrationOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLandingZoneRegistrationOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LandingZoneRegistrationOperationState, opts ...pulumi.ResourceOption) (*LandingZoneRegistrationOperation, error) {
	var resource LandingZoneRegistrationOperation
	err := ctx.ReadResource("azure-native:sovereign:LandingZoneRegistrationOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LandingZoneRegistrationOperation resources.
type landingZoneRegistrationOperationState struct {
}

type LandingZoneRegistrationOperationState struct {
}

func (LandingZoneRegistrationOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*landingZoneRegistrationOperationState)(nil)).Elem()
}

type landingZoneRegistrationOperationArgs struct {
	// The landing zone account.
	LandingZoneAccountName string `pulumi:"landingZoneAccountName"`
	// The name of the landing zone registration resource.
	LandingZoneRegistrationName *string `pulumi:"landingZoneRegistrationName"`
	// The resource-specific properties for this resource.
	Properties *LandingZoneRegistrationResourceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LandingZoneRegistrationOperation resource.
type LandingZoneRegistrationOperationArgs struct {
	// The landing zone account.
	LandingZoneAccountName pulumi.StringInput
	// The name of the landing zone registration resource.
	LandingZoneRegistrationName pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties LandingZoneRegistrationResourcePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (LandingZoneRegistrationOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*landingZoneRegistrationOperationArgs)(nil)).Elem()
}

type LandingZoneRegistrationOperationInput interface {
	pulumi.Input

	ToLandingZoneRegistrationOperationOutput() LandingZoneRegistrationOperationOutput
	ToLandingZoneRegistrationOperationOutputWithContext(ctx context.Context) LandingZoneRegistrationOperationOutput
}

func (*LandingZoneRegistrationOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**LandingZoneRegistrationOperation)(nil)).Elem()
}

func (i *LandingZoneRegistrationOperation) ToLandingZoneRegistrationOperationOutput() LandingZoneRegistrationOperationOutput {
	return i.ToLandingZoneRegistrationOperationOutputWithContext(context.Background())
}

func (i *LandingZoneRegistrationOperation) ToLandingZoneRegistrationOperationOutputWithContext(ctx context.Context) LandingZoneRegistrationOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LandingZoneRegistrationOperationOutput)
}

type LandingZoneRegistrationOperationOutput struct{ *pulumi.OutputState }

func (LandingZoneRegistrationOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LandingZoneRegistrationOperation)(nil)).Elem()
}

func (o LandingZoneRegistrationOperationOutput) ToLandingZoneRegistrationOperationOutput() LandingZoneRegistrationOperationOutput {
	return o
}

func (o LandingZoneRegistrationOperationOutput) ToLandingZoneRegistrationOperationOutputWithContext(ctx context.Context) LandingZoneRegistrationOperationOutput {
	return o
}

// The Azure API version of the resource.
func (o LandingZoneRegistrationOperationOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZoneRegistrationOperation) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o LandingZoneRegistrationOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZoneRegistrationOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LandingZoneRegistrationOperationOutput) Properties() LandingZoneRegistrationResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *LandingZoneRegistrationOperation) LandingZoneRegistrationResourcePropertiesResponseOutput {
		return v.Properties
	}).(LandingZoneRegistrationResourcePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LandingZoneRegistrationOperationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LandingZoneRegistrationOperation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LandingZoneRegistrationOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LandingZoneRegistrationOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LandingZoneRegistrationOperationOutput{})
}
