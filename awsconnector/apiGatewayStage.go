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
type ApiGatewayStage struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties ApiGatewayStagePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiGatewayStage registers a new resource with the given unique name, arguments, and options.
func NewApiGatewayStage(ctx *pulumi.Context,
	name string, args *ApiGatewayStageArgs, opts ...pulumi.ResourceOption) (*ApiGatewayStage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:ApiGatewayStage"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiGatewayStage
	err := ctx.RegisterResource("azure-native:awsconnector:ApiGatewayStage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiGatewayStage gets an existing ApiGatewayStage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiGatewayStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiGatewayStageState, opts ...pulumi.ResourceOption) (*ApiGatewayStage, error) {
	var resource ApiGatewayStage
	err := ctx.ReadResource("azure-native:awsconnector:ApiGatewayStage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiGatewayStage resources.
type apiGatewayStageState struct {
}

type ApiGatewayStageState struct {
}

func (ApiGatewayStageState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayStageState)(nil)).Elem()
}

type apiGatewayStageArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of ApiGatewayStage
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *ApiGatewayStageProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApiGatewayStage resource.
type ApiGatewayStageArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of ApiGatewayStage
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties ApiGatewayStagePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ApiGatewayStageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayStageArgs)(nil)).Elem()
}

type ApiGatewayStageInput interface {
	pulumi.Input

	ToApiGatewayStageOutput() ApiGatewayStageOutput
	ToApiGatewayStageOutputWithContext(ctx context.Context) ApiGatewayStageOutput
}

func (*ApiGatewayStage) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayStage)(nil)).Elem()
}

func (i *ApiGatewayStage) ToApiGatewayStageOutput() ApiGatewayStageOutput {
	return i.ToApiGatewayStageOutputWithContext(context.Background())
}

func (i *ApiGatewayStage) ToApiGatewayStageOutputWithContext(ctx context.Context) ApiGatewayStageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayStageOutput)
}

type ApiGatewayStageOutput struct{ *pulumi.OutputState }

func (ApiGatewayStageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayStage)(nil)).Elem()
}

func (o ApiGatewayStageOutput) ToApiGatewayStageOutput() ApiGatewayStageOutput {
	return o
}

func (o ApiGatewayStageOutput) ToApiGatewayStageOutputWithContext(ctx context.Context) ApiGatewayStageOutput {
	return o
}

// The Azure API version of the resource.
func (o ApiGatewayStageOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayStage) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ApiGatewayStageOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayStage) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApiGatewayStageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayStage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o ApiGatewayStageOutput) Properties() ApiGatewayStagePropertiesResponseOutput {
	return o.ApplyT(func(v *ApiGatewayStage) ApiGatewayStagePropertiesResponseOutput { return v.Properties }).(ApiGatewayStagePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ApiGatewayStageOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiGatewayStage) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ApiGatewayStageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiGatewayStage) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiGatewayStageOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayStage) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiGatewayStageOutput{})
}
