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
type ApiGatewayRestApi struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties ApiGatewayRestApiPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApiGatewayRestApi registers a new resource with the given unique name, arguments, and options.
func NewApiGatewayRestApi(ctx *pulumi.Context,
	name string, args *ApiGatewayRestApiArgs, opts ...pulumi.ResourceOption) (*ApiGatewayRestApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:ApiGatewayRestApi"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiGatewayRestApi
	err := ctx.RegisterResource("azure-native:awsconnector:ApiGatewayRestApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiGatewayRestApi gets an existing ApiGatewayRestApi resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiGatewayRestApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiGatewayRestApiState, opts ...pulumi.ResourceOption) (*ApiGatewayRestApi, error) {
	var resource ApiGatewayRestApi
	err := ctx.ReadResource("azure-native:awsconnector:ApiGatewayRestApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiGatewayRestApi resources.
type apiGatewayRestApiState struct {
}

type ApiGatewayRestApiState struct {
}

func (ApiGatewayRestApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayRestApiState)(nil)).Elem()
}

type apiGatewayRestApiArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of ApiGatewayRestApi
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *ApiGatewayRestApiProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApiGatewayRestApi resource.
type ApiGatewayRestApiArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of ApiGatewayRestApi
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties ApiGatewayRestApiPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ApiGatewayRestApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayRestApiArgs)(nil)).Elem()
}

type ApiGatewayRestApiInput interface {
	pulumi.Input

	ToApiGatewayRestApiOutput() ApiGatewayRestApiOutput
	ToApiGatewayRestApiOutputWithContext(ctx context.Context) ApiGatewayRestApiOutput
}

func (*ApiGatewayRestApi) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayRestApi)(nil)).Elem()
}

func (i *ApiGatewayRestApi) ToApiGatewayRestApiOutput() ApiGatewayRestApiOutput {
	return i.ToApiGatewayRestApiOutputWithContext(context.Background())
}

func (i *ApiGatewayRestApi) ToApiGatewayRestApiOutputWithContext(ctx context.Context) ApiGatewayRestApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayRestApiOutput)
}

type ApiGatewayRestApiOutput struct{ *pulumi.OutputState }

func (ApiGatewayRestApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayRestApi)(nil)).Elem()
}

func (o ApiGatewayRestApiOutput) ToApiGatewayRestApiOutput() ApiGatewayRestApiOutput {
	return o
}

func (o ApiGatewayRestApiOutput) ToApiGatewayRestApiOutputWithContext(ctx context.Context) ApiGatewayRestApiOutput {
	return o
}

// The Azure API version of the resource.
func (o ApiGatewayRestApiOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ApiGatewayRestApiOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApiGatewayRestApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o ApiGatewayRestApiOutput) Properties() ApiGatewayRestApiPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) ApiGatewayRestApiPropertiesResponseOutput { return v.Properties }).(ApiGatewayRestApiPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ApiGatewayRestApiOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ApiGatewayRestApiOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiGatewayRestApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayRestApi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiGatewayRestApiOutput{})
}
