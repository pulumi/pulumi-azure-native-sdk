// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aadiam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AzureADMetrics resource.
//
// Uses Azure REST API version 2020-07-01-preview. In version 1.x of the Azure Native provider, it used API version 2020-07-01-preview.
type AzureADMetric struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name       pulumi.StringOutput                          `pulumi:"name"`
	Properties AzureADMetricsPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAzureADMetric registers a new resource with the given unique name, arguments, and options.
func NewAzureADMetric(ctx *pulumi.Context,
	name string, args *AzureADMetricArgs, opts ...pulumi.ResourceOption) (*AzureADMetric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aadiam/v20200701preview:AzureADMetric"),
		},
		{
			Type: pulumi.String("azure-native:aadiam/v20200701preview:azureADMetric"),
		},
		{
			Type: pulumi.String("azure-native:aadiam:azureADMetric"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AzureADMetric
	err := ctx.RegisterResource("azure-native:aadiam:AzureADMetric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureADMetric gets an existing AzureADMetric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureADMetric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureADMetricState, opts ...pulumi.ResourceOption) (*AzureADMetric, error) {
	var resource AzureADMetric
	err := ctx.ReadResource("azure-native:aadiam:AzureADMetric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureADMetric resources.
type azureADMetricState struct {
}

type AzureADMetricState struct {
}

func (AzureADMetricState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADMetricState)(nil)).Elem()
}

type azureADMetricArgs struct {
	// Name of the azureADMetrics instance.
	AzureADMetricsName *string `pulumi:"azureADMetricsName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AzureADMetric resource.
type AzureADMetricArgs struct {
	// Name of the azureADMetrics instance.
	AzureADMetricsName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (AzureADMetricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureADMetricArgs)(nil)).Elem()
}

type AzureADMetricInput interface {
	pulumi.Input

	ToAzureADMetricOutput() AzureADMetricOutput
	ToAzureADMetricOutputWithContext(ctx context.Context) AzureADMetricOutput
}

func (*AzureADMetric) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADMetric)(nil)).Elem()
}

func (i *AzureADMetric) ToAzureADMetricOutput() AzureADMetricOutput {
	return i.ToAzureADMetricOutputWithContext(context.Background())
}

func (i *AzureADMetric) ToAzureADMetricOutputWithContext(ctx context.Context) AzureADMetricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureADMetricOutput)
}

type AzureADMetricOutput struct{ *pulumi.OutputState }

func (AzureADMetricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureADMetric)(nil)).Elem()
}

func (o AzureADMetricOutput) ToAzureADMetricOutput() AzureADMetricOutput {
	return o
}

func (o AzureADMetricOutput) ToAzureADMetricOutputWithContext(ctx context.Context) AzureADMetricOutput {
	return o
}

// The geo-location where the resource lives
func (o AzureADMetricOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o AzureADMetricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AzureADMetricOutput) Properties() AzureADMetricsPropertiesFormatResponseOutput {
	return o.ApplyT(func(v *AzureADMetric) AzureADMetricsPropertiesFormatResponseOutput { return v.Properties }).(AzureADMetricsPropertiesFormatResponseOutput)
}

// Resource tags.
func (o AzureADMetricOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AzureADMetricOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureADMetric) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADMetricOutput{})
}
