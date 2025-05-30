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

// External OAuth authorization server settings.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
//
// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type AuthorizationServer struct {
	pulumi.CustomResourceState

	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint pulumi.StringOutput `pulumi:"authorizationEndpoint"`
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods pulumi.StringArrayOutput `pulumi:"authorizationMethods"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods pulumi.StringArrayOutput `pulumi:"bearerTokenSendingMethods"`
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod pulumi.StringArrayOutput `pulumi:"clientAuthenticationMethod"`
	// Client or app id registered with this authorization server.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint pulumi.StringOutput `pulumi:"clientRegistrationEndpoint"`
	// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope pulumi.StringPtrOutput `pulumi:"defaultScope"`
	// Description of the authorization server. Can contain HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User-friendly authorization server name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes pulumi.StringArrayOutput `pulumi:"grantTypes"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword pulumi.StringPtrOutput `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername pulumi.StringPtrOutput `pulumi:"resourceOwnerUsername"`
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState pulumi.BoolPtrOutput `pulumi:"supportState"`
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters TokenBodyParameterContractResponseArrayOutput `pulumi:"tokenBodyParameters"`
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint pulumi.StringPtrOutput `pulumi:"tokenEndpoint"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// If true, the authorization server will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation pulumi.BoolPtrOutput `pulumi:"useInApiDocumentation"`
	// If true, the authorization server may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole pulumi.BoolPtrOutput `pulumi:"useInTestConsole"`
}

// NewAuthorizationServer registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationServer(ctx *pulumi.Context,
	name string, args *AuthorizationServerArgs, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationEndpoint'")
	}
	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientRegistrationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'ClientRegistrationEndpoint'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.GrantTypes == nil {
		return nil, errors.New("invalid value for required argument 'GrantTypes'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:AuthorizationServer"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:AuthorizationServer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AuthorizationServer
	err := ctx.RegisterResource("azure-native:apimanagement:AuthorizationServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationServer gets an existing AuthorizationServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationServerState, opts ...pulumi.ResourceOption) (*AuthorizationServer, error) {
	var resource AuthorizationServer
	err := ctx.ReadResource("azure-native:apimanagement:AuthorizationServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationServer resources.
type authorizationServerState struct {
}

type AuthorizationServerState struct {
}

func (AuthorizationServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerState)(nil)).Elem()
}

type authorizationServerArgs struct {
	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods []AuthorizationMethod `pulumi:"authorizationMethods"`
	// Identifier of the authorization server.
	Authsid *string `pulumi:"authsid"`
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods []string `pulumi:"bearerTokenSendingMethods"`
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod []string `pulumi:"clientAuthenticationMethod"`
	// Client or app id registered with this authorization server.
	ClientId string `pulumi:"clientId"`
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint string `pulumi:"clientRegistrationEndpoint"`
	// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret *string `pulumi:"clientSecret"`
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope *string `pulumi:"defaultScope"`
	// Description of the authorization server. Can contain HTML formatting tags.
	Description *string `pulumi:"description"`
	// User-friendly authorization server name.
	DisplayName string `pulumi:"displayName"`
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes []string `pulumi:"grantTypes"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword *string `pulumi:"resourceOwnerPassword"`
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername *string `pulumi:"resourceOwnerUsername"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState *bool `pulumi:"supportState"`
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters []TokenBodyParameterContract `pulumi:"tokenBodyParameters"`
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
	// If true, the authorization server will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation *bool `pulumi:"useInApiDocumentation"`
	// If true, the authorization server may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole *bool `pulumi:"useInTestConsole"`
}

// The set of arguments for constructing a AuthorizationServer resource.
type AuthorizationServerArgs struct {
	// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
	AuthorizationEndpoint pulumi.StringInput
	// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
	AuthorizationMethods AuthorizationMethodArrayInput
	// Identifier of the authorization server.
	Authsid pulumi.StringPtrInput
	// Specifies the mechanism by which access token is passed to the API.
	BearerTokenSendingMethods pulumi.StringArrayInput
	// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
	ClientAuthenticationMethod pulumi.StringArrayInput
	// Client or app id registered with this authorization server.
	ClientId pulumi.StringInput
	// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
	ClientRegistrationEndpoint pulumi.StringInput
	// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	ClientSecret pulumi.StringPtrInput
	// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
	DefaultScope pulumi.StringPtrInput
	// Description of the authorization server. Can contain HTML formatting tags.
	Description pulumi.StringPtrInput
	// User-friendly authorization server name.
	DisplayName pulumi.StringInput
	// Form of an authorization grant, which the client uses to request the access token.
	GrantTypes pulumi.StringArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
	ResourceOwnerPassword pulumi.StringPtrInput
	// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
	ResourceOwnerUsername pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
	SupportState pulumi.BoolPtrInput
	// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
	TokenBodyParameters TokenBodyParameterContractArrayInput
	// OAuth token endpoint. Contains absolute URI to entity being referenced.
	TokenEndpoint pulumi.StringPtrInput
	// If true, the authorization server will be used in the API documentation in the developer portal. False by default if no value is provided.
	UseInApiDocumentation pulumi.BoolPtrInput
	// If true, the authorization server may be used in the developer portal test console. True by default if no value is provided.
	UseInTestConsole pulumi.BoolPtrInput
}

func (AuthorizationServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationServerArgs)(nil)).Elem()
}

type AuthorizationServerInput interface {
	pulumi.Input

	ToAuthorizationServerOutput() AuthorizationServerOutput
	ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput
}

func (*AuthorizationServer) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationServer)(nil)).Elem()
}

func (i *AuthorizationServer) ToAuthorizationServerOutput() AuthorizationServerOutput {
	return i.ToAuthorizationServerOutputWithContext(context.Background())
}

func (i *AuthorizationServer) ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationServerOutput)
}

type AuthorizationServerOutput struct{ *pulumi.OutputState }

func (AuthorizationServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationServer)(nil)).Elem()
}

func (o AuthorizationServerOutput) ToAuthorizationServerOutput() AuthorizationServerOutput {
	return o
}

func (o AuthorizationServerOutput) ToAuthorizationServerOutputWithContext(ctx context.Context) AuthorizationServerOutput {
	return o
}

// OAuth authorization endpoint. See http://tools.ietf.org/html/rfc6749#section-3.2.
func (o AuthorizationServerOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

// HTTP verbs supported by the authorization endpoint. GET must be always present. POST is optional.
func (o AuthorizationServerOutput) AuthorizationMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.AuthorizationMethods }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o AuthorizationServerOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Specifies the mechanism by which access token is passed to the API.
func (o AuthorizationServerOutput) BearerTokenSendingMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.BearerTokenSendingMethods }).(pulumi.StringArrayOutput)
}

// Method of authentication supported by the token endpoint of this authorization server. Possible values are Basic and/or Body. When Body is specified, client credentials and other parameters are passed within the request body in the application/x-www-form-urlencoded format.
func (o AuthorizationServerOutput) ClientAuthenticationMethod() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.ClientAuthenticationMethod }).(pulumi.StringArrayOutput)
}

// Client or app id registered with this authorization server.
func (o AuthorizationServerOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// Optional reference to a page where client or app registration for this authorization server is performed. Contains absolute URL to entity being referenced.
func (o AuthorizationServerOutput) ClientRegistrationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.ClientRegistrationEndpoint }).(pulumi.StringOutput)
}

// Client or app secret registered with this authorization server. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
func (o AuthorizationServerOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Access token scope that is going to be requested by default. Can be overridden at the API level. Should be provided in the form of a string containing space-delimited values.
func (o AuthorizationServerOutput) DefaultScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.DefaultScope }).(pulumi.StringPtrOutput)
}

// Description of the authorization server. Can contain HTML formatting tags.
func (o AuthorizationServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User-friendly authorization server name.
func (o AuthorizationServerOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Form of an authorization grant, which the client uses to request the access token.
func (o AuthorizationServerOutput) GrantTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringArrayOutput { return v.GrantTypes }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o AuthorizationServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner password.
func (o AuthorizationServerOutput) ResourceOwnerPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.ResourceOwnerPassword }).(pulumi.StringPtrOutput)
}

// Can be optionally specified when resource owner password grant type is supported by this authorization server. Default resource owner username.
func (o AuthorizationServerOutput) ResourceOwnerUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.ResourceOwnerUsername }).(pulumi.StringPtrOutput)
}

// If true, authorization server will include state parameter from the authorization request to its response. Client may use state parameter to raise protocol security.
func (o AuthorizationServerOutput) SupportState() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.BoolPtrOutput { return v.SupportState }).(pulumi.BoolPtrOutput)
}

// Additional parameters required by the token endpoint of this authorization server represented as an array of JSON objects with name and value string properties, i.e. {"name" : "name value", "value": "a value"}.
func (o AuthorizationServerOutput) TokenBodyParameters() TokenBodyParameterContractResponseArrayOutput {
	return o.ApplyT(func(v *AuthorizationServer) TokenBodyParameterContractResponseArrayOutput {
		return v.TokenBodyParameters
	}).(TokenBodyParameterContractResponseArrayOutput)
}

// OAuth token endpoint. Contains absolute URI to entity being referenced.
func (o AuthorizationServerOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringPtrOutput { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AuthorizationServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// If true, the authorization server will be used in the API documentation in the developer portal. False by default if no value is provided.
func (o AuthorizationServerOutput) UseInApiDocumentation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.BoolPtrOutput { return v.UseInApiDocumentation }).(pulumi.BoolPtrOutput)
}

// If true, the authorization server may be used in the developer portal test console. True by default if no value is provided.
func (o AuthorizationServerOutput) UseInTestConsole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthorizationServer) pulumi.BoolPtrOutput { return v.UseInTestConsole }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthorizationServerOutput{})
}
