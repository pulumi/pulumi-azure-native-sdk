// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Frontend Subresource of Traffic Controller.
type FrontendsInterface struct {
	pulumi.CustomResourceState

	// Frontend IP Address Version (Optional).
	IpAddressVersion pulumi.StringPtrOutput `pulumi:"ipAddressVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Frontend Mode (Optional).
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// test doc
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Frontend Public IP Address (Optional).
	PublicIPAddress FrontendPropertiesIPAddressResponsePtrOutput `pulumi:"publicIPAddress"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFrontendsInterface registers a new resource with the given unique name, arguments, and options.
func NewFrontendsInterface(ctx *pulumi.Context,
	name string, args *FrontendsInterfaceArgs, opts ...pulumi.ResourceOption) (*FrontendsInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TrafficControllerName == nil {
		return nil, errors.New("invalid value for required argument 'TrafficControllerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicenetworking/v20230501preview:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20231101:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20240501preview:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20250101:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking/v20250301preview:FrontendsInterface"),
		},
		{
			Type: pulumi.String("azure-native:servicenetworking:FrontendsInterface"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FrontendsInterface
	err := ctx.RegisterResource("azure-native:servicenetworking/v20221001preview:FrontendsInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFrontendsInterface gets an existing FrontendsInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFrontendsInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FrontendsInterfaceState, opts ...pulumi.ResourceOption) (*FrontendsInterface, error) {
	var resource FrontendsInterface
	err := ctx.ReadResource("azure-native:servicenetworking/v20221001preview:FrontendsInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FrontendsInterface resources.
type frontendsInterfaceState struct {
}

type FrontendsInterfaceState struct {
}

func (FrontendsInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendsInterfaceState)(nil)).Elem()
}

type frontendsInterfaceArgs struct {
	// Frontends
	FrontendName *string `pulumi:"frontendName"`
	// Frontend IP Address Version (Optional).
	IpAddressVersion *FrontendIPAddressVersion `pulumi:"ipAddressVersion"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Frontend Mode (Optional).
	Mode *FrontendMode `pulumi:"mode"`
	// Frontend Public IP Address (Optional).
	PublicIPAddress *FrontendPropertiesIPAddress `pulumi:"publicIPAddress"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// traffic controller name for path
	TrafficControllerName string `pulumi:"trafficControllerName"`
}

// The set of arguments for constructing a FrontendsInterface resource.
type FrontendsInterfaceArgs struct {
	// Frontends
	FrontendName pulumi.StringPtrInput
	// Frontend IP Address Version (Optional).
	IpAddressVersion FrontendIPAddressVersionPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Frontend Mode (Optional).
	Mode FrontendModePtrInput
	// Frontend Public IP Address (Optional).
	PublicIPAddress FrontendPropertiesIPAddressPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// traffic controller name for path
	TrafficControllerName pulumi.StringInput
}

func (FrontendsInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*frontendsInterfaceArgs)(nil)).Elem()
}

type FrontendsInterfaceInput interface {
	pulumi.Input

	ToFrontendsInterfaceOutput() FrontendsInterfaceOutput
	ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput
}

func (*FrontendsInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendsInterface)(nil)).Elem()
}

func (i *FrontendsInterface) ToFrontendsInterfaceOutput() FrontendsInterfaceOutput {
	return i.ToFrontendsInterfaceOutputWithContext(context.Background())
}

func (i *FrontendsInterface) ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontendsInterfaceOutput)
}

type FrontendsInterfaceOutput struct{ *pulumi.OutputState }

func (FrontendsInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontendsInterface)(nil)).Elem()
}

func (o FrontendsInterfaceOutput) ToFrontendsInterfaceOutput() FrontendsInterfaceOutput {
	return o
}

func (o FrontendsInterfaceOutput) ToFrontendsInterfaceOutputWithContext(ctx context.Context) FrontendsInterfaceOutput {
	return o
}

// Frontend IP Address Version (Optional).
func (o FrontendsInterfaceOutput) IpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringPtrOutput { return v.IpAddressVersion }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o FrontendsInterfaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Frontend Mode (Optional).
func (o FrontendsInterfaceOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o FrontendsInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// test doc
func (o FrontendsInterfaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Frontend Public IP Address (Optional).
func (o FrontendsInterfaceOutput) PublicIPAddress() FrontendPropertiesIPAddressResponsePtrOutput {
	return o.ApplyT(func(v *FrontendsInterface) FrontendPropertiesIPAddressResponsePtrOutput { return v.PublicIPAddress }).(FrontendPropertiesIPAddressResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FrontendsInterfaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FrontendsInterface) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o FrontendsInterfaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FrontendsInterfaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FrontendsInterface) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FrontendsInterfaceOutput{})
}
