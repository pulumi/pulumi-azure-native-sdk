// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network function definition version.
//
// Uses Azure REST API version 2024-04-15. In version 2.x of the Azure Native provider, it used API version 2023-09-01.
//
// Other available API versions: 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridnetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type NetworkFunctionDefinitionVersion struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Network function definition version properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkFunctionDefinitionVersion registers a new resource with the given unique name, arguments, and options.
func NewNetworkFunctionDefinitionVersion(ctx *pulumi.Context,
	name string, args *NetworkFunctionDefinitionVersionArgs, opts ...pulumi.ResourceOption) (*NetworkFunctionDefinitionVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkFunctionDefinitionGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkFunctionDefinitionGroupName'")
	}
	if args.PublisherName == nil {
		return nil, errors.New("invalid value for required argument 'PublisherName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20230901:NetworkFunctionDefinitionVersion"),
		},
		{
			Type: pulumi.String("azure-native:hybridnetwork/v20240415:NetworkFunctionDefinitionVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource NetworkFunctionDefinitionVersion
	err := ctx.RegisterResource("azure-native:hybridnetwork:NetworkFunctionDefinitionVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkFunctionDefinitionVersion gets an existing NetworkFunctionDefinitionVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkFunctionDefinitionVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkFunctionDefinitionVersionState, opts ...pulumi.ResourceOption) (*NetworkFunctionDefinitionVersion, error) {
	var resource NetworkFunctionDefinitionVersion
	err := ctx.ReadResource("azure-native:hybridnetwork:NetworkFunctionDefinitionVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkFunctionDefinitionVersion resources.
type networkFunctionDefinitionVersionState struct {
}

type NetworkFunctionDefinitionVersionState struct {
}

func (NetworkFunctionDefinitionVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFunctionDefinitionVersionState)(nil)).Elem()
}

type networkFunctionDefinitionVersionArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the network function definition group.
	NetworkFunctionDefinitionGroupName string `pulumi:"networkFunctionDefinitionGroupName"`
	// The name of the network function definition version. The name should conform to the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
	NetworkFunctionDefinitionVersionName *string `pulumi:"networkFunctionDefinitionVersionName"`
	// Network function definition version properties.
	Properties interface{} `pulumi:"properties"`
	// The name of the publisher.
	PublisherName string `pulumi:"publisherName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkFunctionDefinitionVersion resource.
type NetworkFunctionDefinitionVersionArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the network function definition group.
	NetworkFunctionDefinitionGroupName pulumi.StringInput
	// The name of the network function definition version. The name should conform to the SemVer 2.0.0 specification: https://semver.org/spec/v2.0.0.html.
	NetworkFunctionDefinitionVersionName pulumi.StringPtrInput
	// Network function definition version properties.
	Properties pulumi.Input
	// The name of the publisher.
	PublisherName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkFunctionDefinitionVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFunctionDefinitionVersionArgs)(nil)).Elem()
}

type NetworkFunctionDefinitionVersionInput interface {
	pulumi.Input

	ToNetworkFunctionDefinitionVersionOutput() NetworkFunctionDefinitionVersionOutput
	ToNetworkFunctionDefinitionVersionOutputWithContext(ctx context.Context) NetworkFunctionDefinitionVersionOutput
}

func (*NetworkFunctionDefinitionVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionDefinitionVersion)(nil)).Elem()
}

func (i *NetworkFunctionDefinitionVersion) ToNetworkFunctionDefinitionVersionOutput() NetworkFunctionDefinitionVersionOutput {
	return i.ToNetworkFunctionDefinitionVersionOutputWithContext(context.Background())
}

func (i *NetworkFunctionDefinitionVersion) ToNetworkFunctionDefinitionVersionOutputWithContext(ctx context.Context) NetworkFunctionDefinitionVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionDefinitionVersionOutput)
}

type NetworkFunctionDefinitionVersionOutput struct{ *pulumi.OutputState }

func (NetworkFunctionDefinitionVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionDefinitionVersion)(nil)).Elem()
}

func (o NetworkFunctionDefinitionVersionOutput) ToNetworkFunctionDefinitionVersionOutput() NetworkFunctionDefinitionVersionOutput {
	return o
}

func (o NetworkFunctionDefinitionVersionOutput) ToNetworkFunctionDefinitionVersionOutputWithContext(ctx context.Context) NetworkFunctionDefinitionVersionOutput {
	return o
}

// The Azure API version of the resource.
func (o NetworkFunctionDefinitionVersionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o NetworkFunctionDefinitionVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o NetworkFunctionDefinitionVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network function definition version properties.
func (o NetworkFunctionDefinitionVersionOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o NetworkFunctionDefinitionVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o NetworkFunctionDefinitionVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o NetworkFunctionDefinitionVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFunctionDefinitionVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkFunctionDefinitionVersionOutput{})
}
