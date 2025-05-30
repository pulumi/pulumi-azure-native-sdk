// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbforpostgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a virtual endpoint for a server.
//
// Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type VirtualEndpoint struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The endpoint type for the virtual endpoint.
	EndpointType pulumi.StringPtrOutput `pulumi:"endpointType"`
	// List of members for a virtual endpoint
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// List of virtual endpoints for a server
	VirtualEndpoints pulumi.StringArrayOutput `pulumi:"virtualEndpoints"`
}

// NewVirtualEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVirtualEndpoint(ctx *pulumi.Context,
	name string, args *VirtualEndpointArgs, opts ...pulumi.ResourceOption) (*VirtualEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20230601preview:VirtualEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20231201preview:VirtualEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20240301preview:VirtualEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20240801:VirtualEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20241101preview:VirtualEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20250101preview:VirtualEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VirtualEndpoint
	err := ctx.RegisterResource("azure-native:dbforpostgresql:VirtualEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualEndpoint gets an existing VirtualEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualEndpointState, opts ...pulumi.ResourceOption) (*VirtualEndpoint, error) {
	var resource VirtualEndpoint
	err := ctx.ReadResource("azure-native:dbforpostgresql:VirtualEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualEndpoint resources.
type virtualEndpointState struct {
}

type VirtualEndpointState struct {
}

func (VirtualEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualEndpointState)(nil)).Elem()
}

type virtualEndpointArgs struct {
	// The endpoint type for the virtual endpoint.
	EndpointType *string `pulumi:"endpointType"`
	// List of members for a virtual endpoint
	Members []string `pulumi:"members"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the virtual endpoint.
	VirtualEndpointName *string `pulumi:"virtualEndpointName"`
}

// The set of arguments for constructing a VirtualEndpoint resource.
type VirtualEndpointArgs struct {
	// The endpoint type for the virtual endpoint.
	EndpointType pulumi.StringPtrInput
	// List of members for a virtual endpoint
	Members pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the virtual endpoint.
	VirtualEndpointName pulumi.StringPtrInput
}

func (VirtualEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualEndpointArgs)(nil)).Elem()
}

type VirtualEndpointInput interface {
	pulumi.Input

	ToVirtualEndpointOutput() VirtualEndpointOutput
	ToVirtualEndpointOutputWithContext(ctx context.Context) VirtualEndpointOutput
}

func (*VirtualEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualEndpoint)(nil)).Elem()
}

func (i *VirtualEndpoint) ToVirtualEndpointOutput() VirtualEndpointOutput {
	return i.ToVirtualEndpointOutputWithContext(context.Background())
}

func (i *VirtualEndpoint) ToVirtualEndpointOutputWithContext(ctx context.Context) VirtualEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualEndpointOutput)
}

type VirtualEndpointOutput struct{ *pulumi.OutputState }

func (VirtualEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualEndpoint)(nil)).Elem()
}

func (o VirtualEndpointOutput) ToVirtualEndpointOutput() VirtualEndpointOutput {
	return o
}

func (o VirtualEndpointOutput) ToVirtualEndpointOutputWithContext(ctx context.Context) VirtualEndpointOutput {
	return o
}

// The Azure API version of the resource.
func (o VirtualEndpointOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualEndpoint) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The endpoint type for the virtual endpoint.
func (o VirtualEndpointOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualEndpoint) pulumi.StringPtrOutput { return v.EndpointType }).(pulumi.StringPtrOutput)
}

// List of members for a virtual endpoint
func (o VirtualEndpointOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualEndpoint) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o VirtualEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VirtualEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VirtualEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// List of virtual endpoints for a server
func (o VirtualEndpointOutput) VirtualEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualEndpoint) pulumi.StringArrayOutput { return v.VirtualEndpoints }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualEndpointOutput{})
}
