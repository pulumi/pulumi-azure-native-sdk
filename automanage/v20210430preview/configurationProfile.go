// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210430preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the configuration profile.
type ConfigurationProfile struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile.
	Properties ConfigurationProfilePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfile registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfile(ctx *pulumi.Context,
	name string, args *ConfigurationProfileArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20220504:ConfigurationProfile"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationProfile
	err := ctx.RegisterResource("azure-native:automanage/v20210430preview:ConfigurationProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfile gets an existing ConfigurationProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileState, opts ...pulumi.ResourceOption) (*ConfigurationProfile, error) {
	var resource ConfigurationProfile
	err := ctx.ReadResource("azure-native:automanage/v20210430preview:ConfigurationProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfile resources.
type configurationProfileState struct {
}

type ConfigurationProfileState struct {
}

func (ConfigurationProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileState)(nil)).Elem()
}

type configurationProfileArgs struct {
	// Name of the configuration profile.
	ConfigurationProfileName *string `pulumi:"configurationProfileName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Properties of the configuration profile.
	Properties *ConfigurationProfileProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationProfile resource.
type ConfigurationProfileArgs struct {
	// Name of the configuration profile.
	ConfigurationProfileName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Properties of the configuration profile.
	Properties ConfigurationProfilePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConfigurationProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileArgs)(nil)).Elem()
}

type ConfigurationProfileInput interface {
	pulumi.Input

	ToConfigurationProfileOutput() ConfigurationProfileOutput
	ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput
}

func (*ConfigurationProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (i *ConfigurationProfile) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return i.ToConfigurationProfileOutputWithContext(context.Background())
}

func (i *ConfigurationProfile) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileOutput)
}

type ConfigurationProfileOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfile)(nil)).Elem()
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutput() ConfigurationProfileOutput {
	return o
}

func (o ConfigurationProfileOutput) ToConfigurationProfileOutputWithContext(ctx context.Context) ConfigurationProfileOutput {
	return o
}

// The geo-location where the resource lives
func (o ConfigurationProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the configuration profile.
func (o ConfigurationProfileOutput) Properties() ConfigurationProfilePropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfile) ConfigurationProfilePropertiesResponseOutput { return v.Properties }).(ConfigurationProfilePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ConfigurationProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileOutput{})
}