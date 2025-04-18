// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration settings for the Azure App Service Authentication / Authorization V2 feature.
//
// Uses Azure REST API version 2021-02-01. In version 2.x of the Azure Native provider, it used API version 2021-02-01.
//
// Other available API versions: 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WebAppAuthSettingsV2 struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The configuration settings that determines the validation flow of users using App Service Authentication/Authorization.
	GlobalValidation GlobalValidationResponsePtrOutput `pulumi:"globalValidation"`
	// The configuration settings of the HTTP requests for authentication and authorization requests made against App Service Authentication/Authorization.
	HttpSettings HttpSettingsResponsePtrOutput `pulumi:"httpSettings"`
	// The configuration settings of each of the identity providers used to configure App Service Authentication/Authorization.
	IdentityProviders IdentityProvidersResponsePtrOutput `pulumi:"identityProviders"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The configuration settings of the login flow of users using App Service Authentication/Authorization.
	Login LoginResponsePtrOutput `pulumi:"login"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration settings of the platform of App Service Authentication/Authorization.
	Platform AuthPlatformResponsePtrOutput `pulumi:"platform"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppAuthSettingsV2 registers a new resource with the given unique name, arguments, and options.
func NewWebAppAuthSettingsV2(ctx *pulumi.Context,
	name string, args *WebAppAuthSettingsV2Args, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAuthSettingsV2"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAuthSettingsV2"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppAuthSettingsV2"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAuthSettingsV2"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAuthSettingsV2"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAuthSettingsV2"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAuthSettingsV2"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppAuthSettingsV2
	err := ctx.RegisterResource("azure-native:web:WebAppAuthSettingsV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppAuthSettingsV2 gets an existing WebAppAuthSettingsV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppAuthSettingsV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAuthSettingsV2State, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsV2, error) {
	var resource WebAppAuthSettingsV2
	err := ctx.ReadResource("azure-native:web:WebAppAuthSettingsV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppAuthSettingsV2 resources.
type webAppAuthSettingsV2State struct {
}

type WebAppAuthSettingsV2State struct {
}

func (WebAppAuthSettingsV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsV2State)(nil)).Elem()
}

type webAppAuthSettingsV2Args struct {
	// The configuration settings that determines the validation flow of users using App Service Authentication/Authorization.
	GlobalValidation *GlobalValidation `pulumi:"globalValidation"`
	// The configuration settings of the HTTP requests for authentication and authorization requests made against App Service Authentication/Authorization.
	HttpSettings *HttpSettings `pulumi:"httpSettings"`
	// The configuration settings of each of the identity providers used to configure App Service Authentication/Authorization.
	IdentityProviders *IdentityProviders `pulumi:"identityProviders"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The configuration settings of the login flow of users using App Service Authentication/Authorization.
	Login *Login `pulumi:"login"`
	// Name of web app.
	Name string `pulumi:"name"`
	// The configuration settings of the platform of App Service Authentication/Authorization.
	Platform *AuthPlatform `pulumi:"platform"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppAuthSettingsV2 resource.
type WebAppAuthSettingsV2Args struct {
	// The configuration settings that determines the validation flow of users using App Service Authentication/Authorization.
	GlobalValidation GlobalValidationPtrInput
	// The configuration settings of the HTTP requests for authentication and authorization requests made against App Service Authentication/Authorization.
	HttpSettings HttpSettingsPtrInput
	// The configuration settings of each of the identity providers used to configure App Service Authentication/Authorization.
	IdentityProviders IdentityProvidersPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// The configuration settings of the login flow of users using App Service Authentication/Authorization.
	Login LoginPtrInput
	// Name of web app.
	Name pulumi.StringInput
	// The configuration settings of the platform of App Service Authentication/Authorization.
	Platform AuthPlatformPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppAuthSettingsV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsV2Args)(nil)).Elem()
}

type WebAppAuthSettingsV2Input interface {
	pulumi.Input

	ToWebAppAuthSettingsV2Output() WebAppAuthSettingsV2Output
	ToWebAppAuthSettingsV2OutputWithContext(ctx context.Context) WebAppAuthSettingsV2Output
}

func (*WebAppAuthSettingsV2) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsV2)(nil)).Elem()
}

func (i *WebAppAuthSettingsV2) ToWebAppAuthSettingsV2Output() WebAppAuthSettingsV2Output {
	return i.ToWebAppAuthSettingsV2OutputWithContext(context.Background())
}

func (i *WebAppAuthSettingsV2) ToWebAppAuthSettingsV2OutputWithContext(ctx context.Context) WebAppAuthSettingsV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAuthSettingsV2Output)
}

type WebAppAuthSettingsV2Output struct{ *pulumi.OutputState }

func (WebAppAuthSettingsV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsV2)(nil)).Elem()
}

func (o WebAppAuthSettingsV2Output) ToWebAppAuthSettingsV2Output() WebAppAuthSettingsV2Output {
	return o
}

func (o WebAppAuthSettingsV2Output) ToWebAppAuthSettingsV2OutputWithContext(ctx context.Context) WebAppAuthSettingsV2Output {
	return o
}

// The Azure API version of the resource.
func (o WebAppAuthSettingsV2Output) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The configuration settings that determines the validation flow of users using App Service Authentication/Authorization.
func (o WebAppAuthSettingsV2Output) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) GlobalValidationResponsePtrOutput { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

// The configuration settings of the HTTP requests for authentication and authorization requests made against App Service Authentication/Authorization.
func (o WebAppAuthSettingsV2Output) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) HttpSettingsResponsePtrOutput { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

// The configuration settings of each of the identity providers used to configure App Service Authentication/Authorization.
func (o WebAppAuthSettingsV2Output) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) IdentityProvidersResponsePtrOutput { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

// Kind of resource.
func (o WebAppAuthSettingsV2Output) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The configuration settings of the login flow of users using App Service Authentication/Authorization.
func (o WebAppAuthSettingsV2Output) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) LoginResponsePtrOutput { return v.Login }).(LoginResponsePtrOutput)
}

// Resource Name.
func (o WebAppAuthSettingsV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration settings of the platform of App Service Authentication/Authorization.
func (o WebAppAuthSettingsV2Output) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) AuthPlatformResponsePtrOutput { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

// Resource type.
func (o WebAppAuthSettingsV2Output) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAuthSettingsV2Output{})
}
