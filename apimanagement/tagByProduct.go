// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tag Contract details.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
//
// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type TagByProduct struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Tag name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagByProduct registers a new resource with the given unique name, arguments, and options.
func NewTagByProduct(ctx *pulumi.Context,
	name string, args *TagByProductArgs, opts ...pulumi.ResourceOption) (*TagByProduct, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:TagByProduct"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:TagByProduct"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagByProduct
	err := ctx.RegisterResource("azure-native:apimanagement:TagByProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagByProduct gets an existing TagByProduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagByProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagByProductState, opts ...pulumi.ResourceOption) (*TagByProduct, error) {
	var resource TagByProduct
	err := ctx.ReadResource("azure-native:apimanagement:TagByProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagByProduct resources.
type tagByProductState struct {
}

type TagByProductState struct {
}

func (TagByProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByProductState)(nil)).Elem()
}

type tagByProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId *string `pulumi:"tagId"`
}

// The set of arguments for constructing a TagByProduct resource.
type TagByProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringPtrInput
}

func (TagByProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagByProductArgs)(nil)).Elem()
}

type TagByProductInput interface {
	pulumi.Input

	ToTagByProductOutput() TagByProductOutput
	ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput
}

func (*TagByProduct) ElementType() reflect.Type {
	return reflect.TypeOf((**TagByProduct)(nil)).Elem()
}

func (i *TagByProduct) ToTagByProductOutput() TagByProductOutput {
	return i.ToTagByProductOutputWithContext(context.Background())
}

func (i *TagByProduct) ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagByProductOutput)
}

type TagByProductOutput struct{ *pulumi.OutputState }

func (TagByProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagByProduct)(nil)).Elem()
}

func (o TagByProductOutput) ToTagByProductOutput() TagByProductOutput {
	return o
}

func (o TagByProductOutput) ToTagByProductOutputWithContext(ctx context.Context) TagByProductOutput {
	return o
}

// The Azure API version of the resource.
func (o TagByProductOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByProduct) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Tag name.
func (o TagByProductOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByProduct) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the resource
func (o TagByProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByProduct) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o TagByProductOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagByProduct) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagByProductOutput{})
}
