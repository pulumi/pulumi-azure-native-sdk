// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines Private endpoint proxy resource.
//
// Uses Azure REST API version 2023-01-01.
type PrivateEndpointConnectionProxyController struct {
	pulumi.CustomResourceState

	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	Name pulumi.StringOutput    `pulumi:"name"`
	// Properties of a private endpoint connection proxy.
	Properties PrivateEndpointConnectionProxyPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}

// NewPrivateEndpointConnectionProxyController registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionProxyController(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionProxyControllerArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxyController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrateProjectName == nil {
		return nil, errors.New("invalid value for required argument 'MigrateProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20230101:PrivateEndpointConnectionProxyController"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionProxyController
	err := ctx.RegisterResource("azure-native:migrate:PrivateEndpointConnectionProxyController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionProxyController gets an existing PrivateEndpointConnectionProxyController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionProxyController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionProxyControllerState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionProxyController, error) {
	var resource PrivateEndpointConnectionProxyController
	err := ctx.ReadResource("azure-native:migrate:PrivateEndpointConnectionProxyController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionProxyController resources.
type privateEndpointConnectionProxyControllerState struct {
}

type PrivateEndpointConnectionProxyControllerState struct {
}

func (PrivateEndpointConnectionProxyControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyControllerState)(nil)).Elem()
}

type privateEndpointConnectionProxyControllerArgs struct {
	ETag *string `pulumi:"eTag"`
	// Name of the Azure Migrate project.
	MigrateProjectName string `pulumi:"migrateProjectName"`
	// Private endpoint proxy name.
	PecProxyName *string `pulumi:"pecProxyName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionProxyController resource.
type PrivateEndpointConnectionProxyControllerArgs struct {
	ETag pulumi.StringPtrInput
	// Name of the Azure Migrate project.
	MigrateProjectName pulumi.StringInput
	// Private endpoint proxy name.
	PecProxyName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (PrivateEndpointConnectionProxyControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionProxyControllerArgs)(nil)).Elem()
}

type PrivateEndpointConnectionProxyControllerInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionProxyControllerOutput() PrivateEndpointConnectionProxyControllerOutput
	ToPrivateEndpointConnectionProxyControllerOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyControllerOutput
}

func (*PrivateEndpointConnectionProxyController) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProxyController)(nil)).Elem()
}

func (i *PrivateEndpointConnectionProxyController) ToPrivateEndpointConnectionProxyControllerOutput() PrivateEndpointConnectionProxyControllerOutput {
	return i.ToPrivateEndpointConnectionProxyControllerOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionProxyController) ToPrivateEndpointConnectionProxyControllerOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionProxyControllerOutput)
}

type PrivateEndpointConnectionProxyControllerOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionProxyControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProxyController)(nil)).Elem()
}

func (o PrivateEndpointConnectionProxyControllerOutput) ToPrivateEndpointConnectionProxyControllerOutput() PrivateEndpointConnectionProxyControllerOutput {
	return o
}

func (o PrivateEndpointConnectionProxyControllerOutput) ToPrivateEndpointConnectionProxyControllerOutputWithContext(ctx context.Context) PrivateEndpointConnectionProxyControllerOutput {
	return o
}

func (o PrivateEndpointConnectionProxyControllerOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxyController) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionProxyControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxyController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of a private endpoint connection proxy.
func (o PrivateEndpointConnectionProxyControllerOutput) Properties() PrivateEndpointConnectionProxyPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxyController) PrivateEndpointConnectionProxyPropertiesResponseOutput {
		return v.Properties
	}).(PrivateEndpointConnectionProxyPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o PrivateEndpointConnectionProxyControllerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxyController) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionProxyControllerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProxyController) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionProxyControllerOutput{})
}
