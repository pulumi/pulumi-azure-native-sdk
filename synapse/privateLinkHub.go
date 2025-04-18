// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A privateLinkHub
//
// Uses Azure REST API version 2021-06-01. In version 2.x of the Azure Native provider, it used API version 2021-06-01.
//
// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type PrivateLinkHub struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of private endpoint connections
	PrivateEndpointConnections PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// PrivateLinkHub provisioning state
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkHub registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkHub(ctx *pulumi.Context,
	name string, args *PrivateLinkHubArgs, opts ...pulumi.ResourceOption) (*PrivateLinkHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:PrivateLinkHub"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:PrivateLinkHub"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkHub
	err := ctx.RegisterResource("azure-native:synapse:PrivateLinkHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkHub gets an existing PrivateLinkHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkHubState, opts ...pulumi.ResourceOption) (*PrivateLinkHub, error) {
	var resource PrivateLinkHub
	err := ctx.ReadResource("azure-native:synapse:PrivateLinkHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkHub resources.
type privateLinkHubState struct {
}

type PrivateLinkHubState struct {
}

func (PrivateLinkHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkHubState)(nil)).Elem()
}

type privateLinkHubArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of the privateLinkHub
	PrivateLinkHubName *string `pulumi:"privateLinkHubName"`
	// PrivateLinkHub provisioning state
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateLinkHub resource.
type PrivateLinkHubArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of the privateLinkHub
	PrivateLinkHubName pulumi.StringPtrInput
	// PrivateLinkHub provisioning state
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PrivateLinkHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkHubArgs)(nil)).Elem()
}

type PrivateLinkHubInput interface {
	pulumi.Input

	ToPrivateLinkHubOutput() PrivateLinkHubOutput
	ToPrivateLinkHubOutputWithContext(ctx context.Context) PrivateLinkHubOutput
}

func (*PrivateLinkHub) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkHub)(nil)).Elem()
}

func (i *PrivateLinkHub) ToPrivateLinkHubOutput() PrivateLinkHubOutput {
	return i.ToPrivateLinkHubOutputWithContext(context.Background())
}

func (i *PrivateLinkHub) ToPrivateLinkHubOutputWithContext(ctx context.Context) PrivateLinkHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkHubOutput)
}

type PrivateLinkHubOutput struct{ *pulumi.OutputState }

func (PrivateLinkHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkHub)(nil)).Elem()
}

func (o PrivateLinkHubOutput) ToPrivateLinkHubOutput() PrivateLinkHubOutput {
	return o
}

func (o PrivateLinkHubOutput) ToPrivateLinkHubOutputWithContext(ctx context.Context) PrivateLinkHubOutput {
	return o
}

// The Azure API version of the resource.
func (o PrivateLinkHubOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o PrivateLinkHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateLinkHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections
func (o PrivateLinkHubOutput) PrivateEndpointConnections() PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
	return o.ApplyT(func(v *PrivateLinkHub) PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionForPrivateLinkHubBasicResponseArrayOutput)
}

// PrivateLinkHub provisioning state
func (o PrivateLinkHubOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o PrivateLinkHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateLinkHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkHubOutput{})
}
