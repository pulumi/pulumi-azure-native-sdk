// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The batch configuration resource definition.
//
// Uses Azure REST API version 2019-05-01. In version 2.x of the Azure Native provider, it used API version 2019-05-01.
//
// Other available API versions: 2016-06-01, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type IntegrationAccountBatchConfiguration struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The batch configuration properties.
	Properties BatchConfigurationPropertiesResponseOutput `pulumi:"properties"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountBatchConfiguration registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, args *IntegrationAccountBatchConfigurationArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountBatchConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountBatchConfiguration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationAccountBatchConfiguration
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccountBatchConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountBatchConfiguration gets an existing IntegrationAccountBatchConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountBatchConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountBatchConfigurationState, opts ...pulumi.ResourceOption) (*IntegrationAccountBatchConfiguration, error) {
	var resource IntegrationAccountBatchConfiguration
	err := ctx.ReadResource("azure-native:logic:IntegrationAccountBatchConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountBatchConfiguration resources.
type integrationAccountBatchConfigurationState struct {
}

type IntegrationAccountBatchConfigurationState struct {
}

func (IntegrationAccountBatchConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountBatchConfigurationState)(nil)).Elem()
}

type integrationAccountBatchConfigurationArgs struct {
	// The batch configuration name.
	BatchConfigurationName *string `pulumi:"batchConfigurationName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The batch configuration properties.
	Properties BatchConfigurationProperties `pulumi:"properties"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountBatchConfiguration resource.
type IntegrationAccountBatchConfigurationArgs struct {
	// The batch configuration name.
	BatchConfigurationName pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The batch configuration properties.
	Properties BatchConfigurationPropertiesInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountBatchConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountBatchConfigurationArgs)(nil)).Elem()
}

type IntegrationAccountBatchConfigurationInput interface {
	pulumi.Input

	ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput
	ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput
}

func (*IntegrationAccountBatchConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountBatchConfiguration)(nil)).Elem()
}

func (i *IntegrationAccountBatchConfiguration) ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput {
	return i.ToIntegrationAccountBatchConfigurationOutputWithContext(context.Background())
}

func (i *IntegrationAccountBatchConfiguration) ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountBatchConfigurationOutput)
}

type IntegrationAccountBatchConfigurationOutput struct{ *pulumi.OutputState }

func (IntegrationAccountBatchConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountBatchConfiguration)(nil)).Elem()
}

func (o IntegrationAccountBatchConfigurationOutput) ToIntegrationAccountBatchConfigurationOutput() IntegrationAccountBatchConfigurationOutput {
	return o
}

func (o IntegrationAccountBatchConfigurationOutput) ToIntegrationAccountBatchConfigurationOutputWithContext(ctx context.Context) IntegrationAccountBatchConfigurationOutput {
	return o
}

// The Azure API version of the resource.
func (o IntegrationAccountBatchConfigurationOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The resource location.
func (o IntegrationAccountBatchConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o IntegrationAccountBatchConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The batch configuration properties.
func (o IntegrationAccountBatchConfigurationOutput) Properties() BatchConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) BatchConfigurationPropertiesResponseOutput {
		return v.Properties
	}).(BatchConfigurationPropertiesResponseOutput)
}

// The resource tags.
func (o IntegrationAccountBatchConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o IntegrationAccountBatchConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountBatchConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountBatchConfigurationOutput{})
}
