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

// OpenId Connect Provider details.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
//
// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type OpenIdConnectProvider struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Client ID of developer console which is the client application.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client Secret of developer console which is the client application.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// User-friendly description of OpenID Connect Provider.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User-friendly OpenID Connect Provider name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Metadata endpoint URI.
	MetadataEndpoint pulumi.StringOutput `pulumi:"metadataEndpoint"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation pulumi.BoolPtrOutput `pulumi:"useInApiDocumentation"`
	// If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole pulumi.BoolPtrOutput `pulumi:"useInTestConsole"`
}

// NewOpenIdConnectProvider registers a new resource with the given unique name, arguments, and options.
func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MetadataEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'MetadataEndpoint'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:OpenIdConnectProvider"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:OpenIdConnectProvider"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OpenIdConnectProvider
	err := ctx.RegisterResource("azure-native:apimanagement:OpenIdConnectProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenIdConnectProvider gets an existing OpenIdConnectProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	var resource OpenIdConnectProvider
	err := ctx.ReadResource("azure-native:apimanagement:OpenIdConnectProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenIdConnectProvider resources.
type openIdConnectProviderState struct {
}

type OpenIdConnectProviderState struct {
}

func (OpenIdConnectProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderState)(nil)).Elem()
}

type openIdConnectProviderArgs struct {
	// Client ID of developer console which is the client application.
	ClientId string `pulumi:"clientId"`
	// Client Secret of developer console which is the client application.
	ClientSecret *string `pulumi:"clientSecret"`
	// User-friendly description of OpenID Connect Provider.
	Description *string `pulumi:"description"`
	// User-friendly OpenID Connect Provider name.
	DisplayName string `pulumi:"displayName"`
	// Metadata endpoint URI.
	MetadataEndpoint string `pulumi:"metadataEndpoint"`
	// Identifier of the OpenID Connect Provider.
	Opid *string `pulumi:"opid"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation *bool `pulumi:"useInApiDocumentation"`
	// If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole *bool `pulumi:"useInTestConsole"`
}

// The set of arguments for constructing a OpenIdConnectProvider resource.
type OpenIdConnectProviderArgs struct {
	// Client ID of developer console which is the client application.
	ClientId pulumi.StringInput
	// Client Secret of developer console which is the client application.
	ClientSecret pulumi.StringPtrInput
	// User-friendly description of OpenID Connect Provider.
	Description pulumi.StringPtrInput
	// User-friendly OpenID Connect Provider name.
	DisplayName pulumi.StringInput
	// Metadata endpoint URI.
	MetadataEndpoint pulumi.StringInput
	// Identifier of the OpenID Connect Provider.
	Opid pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation pulumi.BoolPtrInput
	// If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole pulumi.BoolPtrInput
}

func (OpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderArgs)(nil)).Elem()
}

type OpenIdConnectProviderInput interface {
	pulumi.Input

	ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput
	ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput
}

func (*OpenIdConnectProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectProvider)(nil)).Elem()
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return i.ToOpenIdConnectProviderOutputWithContext(context.Background())
}

func (i *OpenIdConnectProvider) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OpenIdConnectProviderOutput)
}

type OpenIdConnectProviderOutput struct{ *pulumi.OutputState }

func (OpenIdConnectProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OpenIdConnectProvider)(nil)).Elem()
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutput() OpenIdConnectProviderOutput {
	return o
}

func (o OpenIdConnectProviderOutput) ToOpenIdConnectProviderOutputWithContext(ctx context.Context) OpenIdConnectProviderOutput {
	return o
}

// The Azure API version of the resource.
func (o OpenIdConnectProviderOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Client ID of developer console which is the client application.
func (o OpenIdConnectProviderOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// Client Secret of developer console which is the client application.
func (o OpenIdConnectProviderOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// User-friendly description of OpenID Connect Provider.
func (o OpenIdConnectProviderOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User-friendly OpenID Connect Provider name.
func (o OpenIdConnectProviderOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Metadata endpoint URI.
func (o OpenIdConnectProviderOutput) MetadataEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringOutput { return v.MetadataEndpoint }).(pulumi.StringOutput)
}

// The name of the resource
func (o OpenIdConnectProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OpenIdConnectProviderOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// If true, the Open ID Connect provider will be used in the API documentation in the developer portal. False by default if no value is provided.
func (o OpenIdConnectProviderOutput) UseInApiDocumentation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.BoolPtrOutput { return v.UseInApiDocumentation }).(pulumi.BoolPtrOutput)
}

// If true, the Open ID Connect provider may be used in the developer portal test console. True by default if no value is provided.
func (o OpenIdConnectProviderOutput) UseInTestConsole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OpenIdConnectProvider) pulumi.BoolPtrOutput { return v.UseInTestConsole }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OpenIdConnectProviderOutput{})
}
