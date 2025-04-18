// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Microsoft.AwsConnector resource
//
// Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2024-12-01.
type SageMakerApp struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties SageMakerAppPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSageMakerApp registers a new resource with the given unique name, arguments, and options.
func NewSageMakerApp(ctx *pulumi.Context,
	name string, args *SageMakerAppArgs, opts ...pulumi.ResourceOption) (*SageMakerApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:SageMakerApp"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SageMakerApp
	err := ctx.RegisterResource("azure-native:awsconnector:SageMakerApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSageMakerApp gets an existing SageMakerApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSageMakerApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SageMakerAppState, opts ...pulumi.ResourceOption) (*SageMakerApp, error) {
	var resource SageMakerApp
	err := ctx.ReadResource("azure-native:awsconnector:SageMakerApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SageMakerApp resources.
type sageMakerAppState struct {
}

type SageMakerAppState struct {
}

func (SageMakerAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*sageMakerAppState)(nil)).Elem()
}

type sageMakerAppArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of SageMakerApp
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *SageMakerAppProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SageMakerApp resource.
type SageMakerAppArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of SageMakerApp
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties SageMakerAppPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SageMakerAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sageMakerAppArgs)(nil)).Elem()
}

type SageMakerAppInput interface {
	pulumi.Input

	ToSageMakerAppOutput() SageMakerAppOutput
	ToSageMakerAppOutputWithContext(ctx context.Context) SageMakerAppOutput
}

func (*SageMakerApp) ElementType() reflect.Type {
	return reflect.TypeOf((**SageMakerApp)(nil)).Elem()
}

func (i *SageMakerApp) ToSageMakerAppOutput() SageMakerAppOutput {
	return i.ToSageMakerAppOutputWithContext(context.Background())
}

func (i *SageMakerApp) ToSageMakerAppOutputWithContext(ctx context.Context) SageMakerAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SageMakerAppOutput)
}

type SageMakerAppOutput struct{ *pulumi.OutputState }

func (SageMakerAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SageMakerApp)(nil)).Elem()
}

func (o SageMakerAppOutput) ToSageMakerAppOutput() SageMakerAppOutput {
	return o
}

func (o SageMakerAppOutput) ToSageMakerAppOutputWithContext(ctx context.Context) SageMakerAppOutput {
	return o
}

// The Azure API version of the resource.
func (o SageMakerAppOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SageMakerApp) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o SageMakerAppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SageMakerApp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SageMakerAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SageMakerApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o SageMakerAppOutput) Properties() SageMakerAppPropertiesResponseOutput {
	return o.ApplyT(func(v *SageMakerApp) SageMakerAppPropertiesResponseOutput { return v.Properties }).(SageMakerAppPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SageMakerAppOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SageMakerApp) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SageMakerAppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SageMakerApp) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SageMakerAppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SageMakerApp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SageMakerAppOutput{})
}
