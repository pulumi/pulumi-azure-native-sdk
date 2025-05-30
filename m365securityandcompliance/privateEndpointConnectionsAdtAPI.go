// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m365securityandcompliance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Private Endpoint Connection resource.
//
// Uses Azure REST API version 2021-03-25-preview. In version 2.x of the Azure Native provider, it used API version 2021-03-25-preview.
type PrivateEndpointConnectionsAdtAPI struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Required property for system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnectionsAdtAPI registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionsAdtAPIArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsAdtAPI, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:m365securityandcompliance/v20210325preview:PrivateEndpointConnectionsAdtAPI"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateEndpointConnectionsAdtAPI
	err := ctx.RegisterResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsAdtAPI", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnectionsAdtAPI gets an existing PrivateEndpointConnectionsAdtAPI resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnectionsAdtAPI(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionsAdtAPIState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionsAdtAPI, error) {
	var resource PrivateEndpointConnectionsAdtAPI
	err := ctx.ReadResource("azure-native:m365securityandcompliance:PrivateEndpointConnectionsAdtAPI", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnectionsAdtAPI resources.
type privateEndpointConnectionsAdtAPIState struct {
}

type PrivateEndpointConnectionsAdtAPIState struct {
}

func (PrivateEndpointConnectionsAdtAPIState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsAdtAPIState)(nil)).Elem()
}

type privateEndpointConnectionsAdtAPIArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a PrivateEndpointConnectionsAdtAPI resource.
type PrivateEndpointConnectionsAdtAPIArgs struct {
	// The name of the private endpoint connection associated with the Azure resource
	PrivateEndpointConnectionName pulumi.StringPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput
	// The name of the service instance.
	ResourceName pulumi.StringInput
}

func (PrivateEndpointConnectionsAdtAPIArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionsAdtAPIArgs)(nil)).Elem()
}

type PrivateEndpointConnectionsAdtAPIInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionsAdtAPIOutput() PrivateEndpointConnectionsAdtAPIOutput
	ToPrivateEndpointConnectionsAdtAPIOutputWithContext(ctx context.Context) PrivateEndpointConnectionsAdtAPIOutput
}

func (*PrivateEndpointConnectionsAdtAPI) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsAdtAPI)(nil)).Elem()
}

func (i *PrivateEndpointConnectionsAdtAPI) ToPrivateEndpointConnectionsAdtAPIOutput() PrivateEndpointConnectionsAdtAPIOutput {
	return i.ToPrivateEndpointConnectionsAdtAPIOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionsAdtAPI) ToPrivateEndpointConnectionsAdtAPIOutputWithContext(ctx context.Context) PrivateEndpointConnectionsAdtAPIOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionsAdtAPIOutput)
}

type PrivateEndpointConnectionsAdtAPIOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionsAdtAPIOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionsAdtAPI)(nil)).Elem()
}

func (o PrivateEndpointConnectionsAdtAPIOutput) ToPrivateEndpointConnectionsAdtAPIOutput() PrivateEndpointConnectionsAdtAPIOutput {
	return o
}

func (o PrivateEndpointConnectionsAdtAPIOutput) ToPrivateEndpointConnectionsAdtAPIOutputWithContext(ctx context.Context) PrivateEndpointConnectionsAdtAPIOutput {
	return o
}

// The Azure API version of the resource.
func (o PrivateEndpointConnectionsAdtAPIOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateEndpointConnectionsAdtAPIOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionsAdtAPIOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionsAdtAPIOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionsAdtAPIOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Required property for system data
func (o PrivateEndpointConnectionsAdtAPIOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionsAdtAPIOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionsAdtAPI) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionsAdtAPIOutput{})
}
