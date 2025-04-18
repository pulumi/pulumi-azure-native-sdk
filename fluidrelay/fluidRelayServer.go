// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fluidrelay

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A FluidRelay Server.
//
// Uses Azure REST API version 2022-06-01. In version 2.x of the Azure Native provider, it used API version 2022-06-01.
type FluidRelayServer struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// All encryption configuration for a resource.
	Encryption EncryptionPropertiesResponsePtrOutput `pulumi:"encryption"`
	// The Fluid Relay Service endpoints for this server.
	FluidRelayEndpoints FluidRelayEndpointsResponseOutput `pulumi:"fluidRelayEndpoints"`
	// The Fluid tenantId for this server
	FrsTenantId pulumi.StringOutput `pulumi:"frsTenantId"`
	// The type of identity used for the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provision states for FluidRelay RP
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Sku of the storage associated with the resource
	Storagesku pulumi.StringPtrOutput `pulumi:"storagesku"`
	// System meta data for this resource, including creation and modification information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFluidRelayServer registers a new resource with the given unique name, arguments, and options.
func NewFluidRelayServer(ctx *pulumi.Context,
	name string, args *FluidRelayServerArgs, opts ...pulumi.ResourceOption) (*FluidRelayServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210312preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210615preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210830preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20210910preview:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220215:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220421:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220511:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220526:FluidRelayServer"),
		},
		{
			Type: pulumi.String("azure-native:fluidrelay/v20220601:FluidRelayServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FluidRelayServer
	err := ctx.RegisterResource("azure-native:fluidrelay:FluidRelayServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFluidRelayServer gets an existing FluidRelayServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFluidRelayServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FluidRelayServerState, opts ...pulumi.ResourceOption) (*FluidRelayServer, error) {
	var resource FluidRelayServer
	err := ctx.ReadResource("azure-native:fluidrelay:FluidRelayServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FluidRelayServer resources.
type fluidRelayServerState struct {
}

type FluidRelayServerState struct {
}

func (FluidRelayServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fluidRelayServerState)(nil)).Elem()
}

type fluidRelayServerArgs struct {
	// All encryption configuration for a resource.
	Encryption *EncryptionProperties `pulumi:"encryption"`
	// The Fluid Relay server resource name.
	FluidRelayServerName *string `pulumi:"fluidRelayServerName"`
	// The type of identity used for the resource.
	Identity *Identity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Provision states for FluidRelay RP
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource group containing the resource.
	ResourceGroup string `pulumi:"resourceGroup"`
	// Sku of the storage associated with the resource
	Storagesku *string `pulumi:"storagesku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FluidRelayServer resource.
type FluidRelayServerArgs struct {
	// All encryption configuration for a resource.
	Encryption EncryptionPropertiesPtrInput
	// The Fluid Relay server resource name.
	FluidRelayServerName pulumi.StringPtrInput
	// The type of identity used for the resource.
	Identity IdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Provision states for FluidRelay RP
	ProvisioningState pulumi.StringPtrInput
	// The resource group containing the resource.
	ResourceGroup pulumi.StringInput
	// Sku of the storage associated with the resource
	Storagesku pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (FluidRelayServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fluidRelayServerArgs)(nil)).Elem()
}

type FluidRelayServerInput interface {
	pulumi.Input

	ToFluidRelayServerOutput() FluidRelayServerOutput
	ToFluidRelayServerOutputWithContext(ctx context.Context) FluidRelayServerOutput
}

func (*FluidRelayServer) ElementType() reflect.Type {
	return reflect.TypeOf((**FluidRelayServer)(nil)).Elem()
}

func (i *FluidRelayServer) ToFluidRelayServerOutput() FluidRelayServerOutput {
	return i.ToFluidRelayServerOutputWithContext(context.Background())
}

func (i *FluidRelayServer) ToFluidRelayServerOutputWithContext(ctx context.Context) FluidRelayServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FluidRelayServerOutput)
}

type FluidRelayServerOutput struct{ *pulumi.OutputState }

func (FluidRelayServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FluidRelayServer)(nil)).Elem()
}

func (o FluidRelayServerOutput) ToFluidRelayServerOutput() FluidRelayServerOutput {
	return o
}

func (o FluidRelayServerOutput) ToFluidRelayServerOutputWithContext(ctx context.Context) FluidRelayServerOutput {
	return o
}

// The Azure API version of the resource.
func (o FluidRelayServerOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// All encryption configuration for a resource.
func (o FluidRelayServerOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) EncryptionPropertiesResponsePtrOutput { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

// The Fluid Relay Service endpoints for this server.
func (o FluidRelayServerOutput) FluidRelayEndpoints() FluidRelayEndpointsResponseOutput {
	return o.ApplyT(func(v *FluidRelayServer) FluidRelayEndpointsResponseOutput { return v.FluidRelayEndpoints }).(FluidRelayEndpointsResponseOutput)
}

// The Fluid tenantId for this server
func (o FluidRelayServerOutput) FrsTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.FrsTenantId }).(pulumi.StringOutput)
}

// The type of identity used for the resource.
func (o FluidRelayServerOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o FluidRelayServerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o FluidRelayServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provision states for FluidRelay RP
func (o FluidRelayServerOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Sku of the storage associated with the resource
func (o FluidRelayServerOutput) Storagesku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringPtrOutput { return v.Storagesku }).(pulumi.StringPtrOutput)
}

// System meta data for this resource, including creation and modification information.
func (o FluidRelayServerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FluidRelayServer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o FluidRelayServerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FluidRelayServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FluidRelayServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FluidRelayServerOutput{})
}
