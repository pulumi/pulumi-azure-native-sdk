// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package delegatednetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an instance of a orchestrator.
//
// Uses Azure REST API version 2023-06-27-preview. In version 2.x of the Azure Native provider, it used API version 2021-03-15.
//
// Other available API versions: 2021-03-15, 2023-05-18-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native delegatednetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type DelegatedSubnetServiceDetails struct {
	pulumi.CustomResourceState

	// Defines prefix size of CIDR blocks allocated to nodes in VnetBlock Mode.
	// Delegated subnet's prefix size should be smaller than this by a minimum of 3.
	AllocationBlockPrefixSize pulumi.IntPtrOutput `pulumi:"allocationBlockPrefixSize"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Properties of the controller.
	ControllerDetails ControllerDetailsResponsePtrOutput `pulumi:"controllerDetails"`
	// Location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current state of dnc delegated subnet resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource guid.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// subnet details
	SubnetDetails SubnetDetailsResponsePtrOutput `pulumi:"subnetDetails"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDelegatedSubnetServiceDetails registers a new resource with the given unique name, arguments, and options.
func NewDelegatedSubnetServiceDetails(ctx *pulumi.Context,
	name string, args *DelegatedSubnetServiceDetailsArgs, opts ...pulumi.ResourceOption) (*DelegatedSubnetServiceDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20200808preview:DelegatedSubnetServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20210315:DelegatedSubnetServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20230518preview:DelegatedSubnetServiceDetails"),
		},
		{
			Type: pulumi.String("azure-native:delegatednetwork/v20230627preview:DelegatedSubnetServiceDetails"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DelegatedSubnetServiceDetails
	err := ctx.RegisterResource("azure-native:delegatednetwork:DelegatedSubnetServiceDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDelegatedSubnetServiceDetails gets an existing DelegatedSubnetServiceDetails resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDelegatedSubnetServiceDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DelegatedSubnetServiceDetailsState, opts ...pulumi.ResourceOption) (*DelegatedSubnetServiceDetails, error) {
	var resource DelegatedSubnetServiceDetails
	err := ctx.ReadResource("azure-native:delegatednetwork:DelegatedSubnetServiceDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DelegatedSubnetServiceDetails resources.
type delegatedSubnetServiceDetailsState struct {
}

type DelegatedSubnetServiceDetailsState struct {
}

func (DelegatedSubnetServiceDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*delegatedSubnetServiceDetailsState)(nil)).Elem()
}

type delegatedSubnetServiceDetailsArgs struct {
	// Defines prefix size of CIDR blocks allocated to nodes in VnetBlock Mode.
	// Delegated subnet's prefix size should be smaller than this by a minimum of 3.
	AllocationBlockPrefixSize *int `pulumi:"allocationBlockPrefixSize"`
	// Properties of the controller.
	ControllerDetails *ControllerDetailsType `pulumi:"controllerDetails"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
	ResourceName *string `pulumi:"resourceName"`
	// subnet details
	SubnetDetails *SubnetDetails `pulumi:"subnetDetails"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DelegatedSubnetServiceDetails resource.
type DelegatedSubnetServiceDetailsArgs struct {
	// Defines prefix size of CIDR blocks allocated to nodes in VnetBlock Mode.
	// Delegated subnet's prefix size should be smaller than this by a minimum of 3.
	AllocationBlockPrefixSize pulumi.IntPtrInput
	// Properties of the controller.
	ControllerDetails ControllerDetailsTypePtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
	ResourceName pulumi.StringPtrInput
	// subnet details
	SubnetDetails SubnetDetailsPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (DelegatedSubnetServiceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*delegatedSubnetServiceDetailsArgs)(nil)).Elem()
}

type DelegatedSubnetServiceDetailsInput interface {
	pulumi.Input

	ToDelegatedSubnetServiceDetailsOutput() DelegatedSubnetServiceDetailsOutput
	ToDelegatedSubnetServiceDetailsOutputWithContext(ctx context.Context) DelegatedSubnetServiceDetailsOutput
}

func (*DelegatedSubnetServiceDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetServiceDetails)(nil)).Elem()
}

func (i *DelegatedSubnetServiceDetails) ToDelegatedSubnetServiceDetailsOutput() DelegatedSubnetServiceDetailsOutput {
	return i.ToDelegatedSubnetServiceDetailsOutputWithContext(context.Background())
}

func (i *DelegatedSubnetServiceDetails) ToDelegatedSubnetServiceDetailsOutputWithContext(ctx context.Context) DelegatedSubnetServiceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DelegatedSubnetServiceDetailsOutput)
}

type DelegatedSubnetServiceDetailsOutput struct{ *pulumi.OutputState }

func (DelegatedSubnetServiceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DelegatedSubnetServiceDetails)(nil)).Elem()
}

func (o DelegatedSubnetServiceDetailsOutput) ToDelegatedSubnetServiceDetailsOutput() DelegatedSubnetServiceDetailsOutput {
	return o
}

func (o DelegatedSubnetServiceDetailsOutput) ToDelegatedSubnetServiceDetailsOutputWithContext(ctx context.Context) DelegatedSubnetServiceDetailsOutput {
	return o
}

// Defines prefix size of CIDR blocks allocated to nodes in VnetBlock Mode.
// Delegated subnet's prefix size should be smaller than this by a minimum of 3.
func (o DelegatedSubnetServiceDetailsOutput) AllocationBlockPrefixSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.IntPtrOutput { return v.AllocationBlockPrefixSize }).(pulumi.IntPtrOutput)
}

// The Azure API version of the resource.
func (o DelegatedSubnetServiceDetailsOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Properties of the controller.
func (o DelegatedSubnetServiceDetailsOutput) ControllerDetails() ControllerDetailsResponsePtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) ControllerDetailsResponsePtrOutput { return v.ControllerDetails }).(ControllerDetailsResponsePtrOutput)
}

// Location of the resource.
func (o DelegatedSubnetServiceDetailsOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o DelegatedSubnetServiceDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current state of dnc delegated subnet resource.
func (o DelegatedSubnetServiceDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource guid.
func (o DelegatedSubnetServiceDetailsOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// subnet details
func (o DelegatedSubnetServiceDetailsOutput) SubnetDetails() SubnetDetailsResponsePtrOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) SubnetDetailsResponsePtrOutput { return v.SubnetDetails }).(SubnetDetailsResponsePtrOutput)
}

// The resource tags.
func (o DelegatedSubnetServiceDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of resource.
func (o DelegatedSubnetServiceDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DelegatedSubnetServiceDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DelegatedSubnetServiceDetailsOutput{})
}
