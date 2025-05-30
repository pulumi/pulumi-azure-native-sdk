// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
//
// Uses Azure REST API version 2024-10-02-preview. In version 2.x of the Azure Native provider, it used API version 2024-02-02-preview.
//
// Other available API versions: 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ManagedEnvironmentPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The group ids for the private endpoint resource.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedEnvironmentPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewManagedEnvironmentPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *ManagedEnvironmentPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*ManagedEnvironmentPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:app/v20240202preview:ManagedEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:app/v20240802preview:ManagedEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:app/v20241002preview:ManagedEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:app/v20250202preview:ManagedEnvironmentPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedEnvironmentPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:app:ManagedEnvironmentPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedEnvironmentPrivateEndpointConnection gets an existing ManagedEnvironmentPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedEnvironmentPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedEnvironmentPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*ManagedEnvironmentPrivateEndpointConnection, error) {
	var resource ManagedEnvironmentPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:app:ManagedEnvironmentPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedEnvironmentPrivateEndpointConnection resources.
type managedEnvironmentPrivateEndpointConnectionState struct {
}

type ManagedEnvironmentPrivateEndpointConnectionState struct {
}

func (ManagedEnvironmentPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentPrivateEndpointConnectionState)(nil)).Elem()
}

type managedEnvironmentPrivateEndpointConnectionArgs struct {
	// Name of the Managed Environment.
	EnvironmentName string `pulumi:"environmentName"`
	// The name of the private endpoint connection associated with the Azure resource.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagedEnvironmentPrivateEndpointConnection resource.
type ManagedEnvironmentPrivateEndpointConnectionArgs struct {
	// Name of the Managed Environment.
	EnvironmentName pulumi.StringInput
	// The name of the private endpoint connection associated with the Azure resource.
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ManagedEnvironmentPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedEnvironmentPrivateEndpointConnectionArgs)(nil)).Elem()
}

type ManagedEnvironmentPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToManagedEnvironmentPrivateEndpointConnectionOutput() ManagedEnvironmentPrivateEndpointConnectionOutput
	ToManagedEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedEnvironmentPrivateEndpointConnectionOutput
}

func (*ManagedEnvironmentPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironmentPrivateEndpointConnection)(nil)).Elem()
}

func (i *ManagedEnvironmentPrivateEndpointConnection) ToManagedEnvironmentPrivateEndpointConnectionOutput() ManagedEnvironmentPrivateEndpointConnectionOutput {
	return i.ToManagedEnvironmentPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *ManagedEnvironmentPrivateEndpointConnection) ToManagedEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedEnvironmentPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedEnvironmentPrivateEndpointConnectionOutput)
}

type ManagedEnvironmentPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (ManagedEnvironmentPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedEnvironmentPrivateEndpointConnection)(nil)).Elem()
}

func (o ManagedEnvironmentPrivateEndpointConnectionOutput) ToManagedEnvironmentPrivateEndpointConnectionOutput() ManagedEnvironmentPrivateEndpointConnectionOutput {
	return o
}

func (o ManagedEnvironmentPrivateEndpointConnectionOutput) ToManagedEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedEnvironmentPrivateEndpointConnectionOutput {
	return o
}

// The Azure API version of the resource.
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The group ids for the private endpoint resource.
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedEnvironmentPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedEnvironmentPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedEnvironmentPrivateEndpointConnectionOutput{})
}
