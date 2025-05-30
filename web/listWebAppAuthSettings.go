// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the Authentication/Authorization settings of an app.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListWebAppAuthSettings(ctx *pulumi.Context, args *ListWebAppAuthSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppAuthSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppAuthSettingsResult
	err := ctx.Invoke("azure-native:web:listWebAppAuthSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAuthSettingsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Configuration settings for the Azure App Service Authentication / Authorization feature.
type ListWebAppAuthSettingsResult struct {
	// Gets a JSON string containing the Azure AD Acl settings.
	AadClaimsAuthorization *string `pulumi:"aadClaimsAuthorization"`
	// Login parameters to send to the OpenID Connect authorization endpoint when
	// a user logs in. Each parameter must be in the form "key=value".
	AdditionalLoginParams []string `pulumi:"additionalLoginParams"`
	// Allowed audience values to consider when validating JSON Web Tokens issued by
	// Azure Active Directory. Note that the <code>ClientID</code> value is always considered an
	// allowed audience, regardless of this setting.
	AllowedAudiences []string `pulumi:"allowedAudiences"`
	// External URLs that can be redirected to as part of logging in or logging out of the app. Note that the query string part of the URL is ignored.
	// This is an advanced setting typically only needed by Windows Store application backends.
	// Note that URLs within the current domain are always implicitly allowed.
	AllowedExternalRedirectUrls []string `pulumi:"allowedExternalRedirectUrls"`
	// The path of the config file containing auth settings.
	// If the path is relative, base will the site's root directory.
	AuthFilePath *string `pulumi:"authFilePath"`
	// The Client ID of this relying party application, known as the client_id.
	// This setting is required for enabling OpenID Connection authentication with Azure Active Directory or
	// other 3rd party OpenID Connect providers.
	// More information on OpenID Connect: http://openid.net/specs/openid-connect-core-1_0.html
	ClientId *string `pulumi:"clientId"`
	// The Client Secret of this relying party application (in Azure Active Directory, this is also referred to as the Key).
	// This setting is optional. If no client secret is configured, the OpenID Connect implicit auth flow is used to authenticate end users.
	// Otherwise, the OpenID Connect Authorization Code Flow is used to authenticate end users.
	// More information on OpenID Connect: http://openid.net/specs/openid-connect-core-1_0.html
	ClientSecret *string `pulumi:"clientSecret"`
	// An alternative to the client secret, that is the thumbprint of a certificate used for signing purposes. This property acts as
	// a replacement for the Client Secret. It is also optional.
	ClientSecretCertificateThumbprint *string `pulumi:"clientSecretCertificateThumbprint"`
	// The app setting name that contains the client secret of the relying party application.
	ClientSecretSettingName *string `pulumi:"clientSecretSettingName"`
	// The ConfigVersion of the Authentication / Authorization feature in use for the current app.
	// The setting in this value can control the behavior of the control plane for Authentication / Authorization.
	ConfigVersion *string `pulumi:"configVersion"`
	// The default authentication provider to use when multiple providers are configured.
	// This setting is only needed if multiple providers are configured and the unauthenticated client
	// action is set to "RedirectToLoginPage".
	DefaultProvider *string `pulumi:"defaultProvider"`
	// <code>true</code> if the Authentication / Authorization feature is enabled for the current app; otherwise, <code>false</code>.
	Enabled *bool `pulumi:"enabled"`
	// The App ID of the Facebook app used for login.
	// This setting is required for enabling Facebook Login.
	// Facebook Login documentation: https://developers.facebook.com/docs/facebook-login
	FacebookAppId *string `pulumi:"facebookAppId"`
	// The App Secret of the Facebook app used for Facebook Login.
	// This setting is required for enabling Facebook Login.
	// Facebook Login documentation: https://developers.facebook.com/docs/facebook-login
	FacebookAppSecret *string `pulumi:"facebookAppSecret"`
	// The app setting name that contains the app secret used for Facebook Login.
	FacebookAppSecretSettingName *string `pulumi:"facebookAppSecretSettingName"`
	// The OAuth 2.0 scopes that will be requested as part of Facebook Login authentication.
	// This setting is optional.
	// Facebook Login documentation: https://developers.facebook.com/docs/facebook-login
	FacebookOAuthScopes []string `pulumi:"facebookOAuthScopes"`
	// The Client Id of the GitHub app used for login.
	// This setting is required for enabling Github login
	GitHubClientId *string `pulumi:"gitHubClientId"`
	// The Client Secret of the GitHub app used for Github Login.
	// This setting is required for enabling Github login.
	GitHubClientSecret *string `pulumi:"gitHubClientSecret"`
	// The app setting name that contains the client secret of the Github
	// app used for GitHub Login.
	GitHubClientSecretSettingName *string `pulumi:"gitHubClientSecretSettingName"`
	// The OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.
	// This setting is optional
	GitHubOAuthScopes []string `pulumi:"gitHubOAuthScopes"`
	// The OpenID Connect Client ID for the Google web application.
	// This setting is required for enabling Google Sign-In.
	// Google Sign-In documentation: https://developers.google.com/identity/sign-in/web/
	GoogleClientId *string `pulumi:"googleClientId"`
	// The client secret associated with the Google web application.
	// This setting is required for enabling Google Sign-In.
	// Google Sign-In documentation: https://developers.google.com/identity/sign-in/web/
	GoogleClientSecret *string `pulumi:"googleClientSecret"`
	// The app setting name that contains the client secret associated with
	// the Google web application.
	GoogleClientSecretSettingName *string `pulumi:"googleClientSecretSettingName"`
	// The OAuth 2.0 scopes that will be requested as part of Google Sign-In authentication.
	// This setting is optional. If not specified, "openid", "profile", and "email" are used as default scopes.
	// Google Sign-In documentation: https://developers.google.com/identity/sign-in/web/
	GoogleOAuthScopes []string `pulumi:"googleOAuthScopes"`
	// Resource Id.
	Id string `pulumi:"id"`
	// "true" if the auth config settings should be read from a file,
	// "false" otherwise
	IsAuthFromFile *string `pulumi:"isAuthFromFile"`
	// The OpenID Connect Issuer URI that represents the entity which issues access tokens for this application.
	// When using Azure Active Directory, this value is the URI of the directory tenant, e.g. `https://sts.windows.net/{tenant-guid}/`.
	// This URI is a case-sensitive identifier for the token issuer.
	// More information on OpenID Connect Discovery: http://openid.net/specs/openid-connect-discovery-1_0.html
	Issuer *string `pulumi:"issuer"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// The OAuth 2.0 client ID that was created for the app used for authentication.
	// This setting is required for enabling Microsoft Account authentication.
	// Microsoft Account OAuth documentation: https://dev.onedrive.com/auth/msa_oauth.htm
	MicrosoftAccountClientId *string `pulumi:"microsoftAccountClientId"`
	// The OAuth 2.0 client secret that was created for the app used for authentication.
	// This setting is required for enabling Microsoft Account authentication.
	// Microsoft Account OAuth documentation: https://dev.onedrive.com/auth/msa_oauth.htm
	MicrosoftAccountClientSecret *string `pulumi:"microsoftAccountClientSecret"`
	// The app setting name containing the OAuth 2.0 client secret that was created for the
	// app used for authentication.
	MicrosoftAccountClientSecretSettingName *string `pulumi:"microsoftAccountClientSecretSettingName"`
	// The OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication.
	// This setting is optional. If not specified, "wl.basic" is used as the default scope.
	// Microsoft Account Scopes and permissions documentation: https://msdn.microsoft.com/en-us/library/dn631845.aspx
	MicrosoftAccountOAuthScopes []string `pulumi:"microsoftAccountOAuthScopes"`
	// Resource Name.
	Name string `pulumi:"name"`
	// The RuntimeVersion of the Authentication / Authorization feature in use for the current app.
	// The setting in this value can control the behavior of certain features in the Authentication / Authorization module.
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// The number of hours after session token expiration that a session token can be used to
	// call the token refresh API. The default is 72 hours.
	TokenRefreshExtensionHours *float64 `pulumi:"tokenRefreshExtensionHours"`
	// <code>true</code> to durably store platform-specific security tokens that are obtained during login flows; otherwise, <code>false</code>.
	//  The default is <code>false</code>.
	TokenStoreEnabled *bool `pulumi:"tokenStoreEnabled"`
	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	// This setting is required for enabling Twitter Sign-In.
	// Twitter Sign-In documentation: https://dev.twitter.com/web/sign-in
	TwitterConsumerKey *string `pulumi:"twitterConsumerKey"`
	// The OAuth 1.0a consumer secret of the Twitter application used for sign-in.
	// This setting is required for enabling Twitter Sign-In.
	// Twitter Sign-In documentation: https://dev.twitter.com/web/sign-in
	TwitterConsumerSecret *string `pulumi:"twitterConsumerSecret"`
	// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter
	// application used for sign-in.
	TwitterConsumerSecretSettingName *string `pulumi:"twitterConsumerSecretSettingName"`
	// Resource type.
	Type string `pulumi:"type"`
	// The action to take when an unauthenticated client attempts to access the app.
	UnauthenticatedClientAction *string `pulumi:"unauthenticatedClientAction"`
	// Gets a value indicating whether the issuer should be a valid HTTPS url and be validated as such.
	ValidateIssuer *bool `pulumi:"validateIssuer"`
}

func ListWebAppAuthSettingsOutput(ctx *pulumi.Context, args ListWebAppAuthSettingsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppAuthSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWebAppAuthSettingsResultOutput, error) {
			args := v.(ListWebAppAuthSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web:listWebAppAuthSettings", args, ListWebAppAuthSettingsResultOutput{}, options).(ListWebAppAuthSettingsResultOutput), nil
		}).(ListWebAppAuthSettingsResultOutput)
}

type ListWebAppAuthSettingsOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppAuthSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAuthSettingsArgs)(nil)).Elem()
}

// Configuration settings for the Azure App Service Authentication / Authorization feature.
type ListWebAppAuthSettingsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppAuthSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAuthSettingsResult)(nil)).Elem()
}

func (o ListWebAppAuthSettingsResultOutput) ToListWebAppAuthSettingsResultOutput() ListWebAppAuthSettingsResultOutput {
	return o
}

func (o ListWebAppAuthSettingsResultOutput) ToListWebAppAuthSettingsResultOutputWithContext(ctx context.Context) ListWebAppAuthSettingsResultOutput {
	return o
}

// Gets a JSON string containing the Azure AD Acl settings.
func (o ListWebAppAuthSettingsResultOutput) AadClaimsAuthorization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.AadClaimsAuthorization }).(pulumi.StringPtrOutput)
}

// Login parameters to send to the OpenID Connect authorization endpoint when
// a user logs in. Each parameter must be in the form "key=value".
func (o ListWebAppAuthSettingsResultOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

// Allowed audience values to consider when validating JSON Web Tokens issued by
// Azure Active Directory. Note that the <code>ClientID</code> value is always considered an
// allowed audience, regardless of this setting.
func (o ListWebAppAuthSettingsResultOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

// External URLs that can be redirected to as part of logging in or logging out of the app. Note that the query string part of the URL is ignored.
// This is an advanced setting typically only needed by Windows Store application backends.
// Note that URLs within the current domain are always implicitly allowed.
func (o ListWebAppAuthSettingsResultOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

// The path of the config file containing auth settings.
// If the path is relative, base will the site's root directory.
func (o ListWebAppAuthSettingsResultOutput) AuthFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.AuthFilePath }).(pulumi.StringPtrOutput)
}

// The Client ID of this relying party application, known as the client_id.
// This setting is required for enabling OpenID Connection authentication with Azure Active Directory or
// other 3rd party OpenID Connect providers.
// More information on OpenID Connect: http://openid.net/specs/openid-connect-core-1_0.html
func (o ListWebAppAuthSettingsResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The Client Secret of this relying party application (in Azure Active Directory, this is also referred to as the Key).
// This setting is optional. If no client secret is configured, the OpenID Connect implicit auth flow is used to authenticate end users.
// Otherwise, the OpenID Connect Authorization Code Flow is used to authenticate end users.
// More information on OpenID Connect: http://openid.net/specs/openid-connect-core-1_0.html
func (o ListWebAppAuthSettingsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// An alternative to the client secret, that is the thumbprint of a certificate used for signing purposes. This property acts as
// a replacement for the Client Secret. It is also optional.
func (o ListWebAppAuthSettingsResultOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

// The app setting name that contains the client secret of the relying party application.
func (o ListWebAppAuthSettingsResultOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

// The ConfigVersion of the Authentication / Authorization feature in use for the current app.
// The setting in this value can control the behavior of the control plane for Authentication / Authorization.
func (o ListWebAppAuthSettingsResultOutput) ConfigVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ConfigVersion }).(pulumi.StringPtrOutput)
}

// The default authentication provider to use when multiple providers are configured.
// This setting is only needed if multiple providers are configured and the unauthenticated client
// action is set to "RedirectToLoginPage".
func (o ListWebAppAuthSettingsResultOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the Authentication / Authorization feature is enabled for the current app; otherwise, <code>false</code>.
func (o ListWebAppAuthSettingsResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The App ID of the Facebook app used for login.
// This setting is required for enabling Facebook Login.
// Facebook Login documentation: https://developers.facebook.com/docs/facebook-login
func (o ListWebAppAuthSettingsResultOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

// The App Secret of the Facebook app used for Facebook Login.
// This setting is required for enabling Facebook Login.
// Facebook Login documentation: https://developers.facebook.com/docs/facebook-login
func (o ListWebAppAuthSettingsResultOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

// The app setting name that contains the app secret used for Facebook Login.
func (o ListWebAppAuthSettingsResultOutput) FacebookAppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.FacebookAppSecretSettingName }).(pulumi.StringPtrOutput)
}

// The OAuth 2.0 scopes that will be requested as part of Facebook Login authentication.
// This setting is optional.
// Facebook Login documentation: https://developers.facebook.com/docs/facebook-login
func (o ListWebAppAuthSettingsResultOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

// The Client Id of the GitHub app used for login.
// This setting is required for enabling Github login
func (o ListWebAppAuthSettingsResultOutput) GitHubClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GitHubClientId }).(pulumi.StringPtrOutput)
}

// The Client Secret of the GitHub app used for Github Login.
// This setting is required for enabling Github login.
func (o ListWebAppAuthSettingsResultOutput) GitHubClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GitHubClientSecret }).(pulumi.StringPtrOutput)
}

// The app setting name that contains the client secret of the Github
// app used for GitHub Login.
func (o ListWebAppAuthSettingsResultOutput) GitHubClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GitHubClientSecretSettingName }).(pulumi.StringPtrOutput)
}

// The OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.
// This setting is optional
func (o ListWebAppAuthSettingsResultOutput) GitHubOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.GitHubOAuthScopes }).(pulumi.StringArrayOutput)
}

// The OpenID Connect Client ID for the Google web application.
// This setting is required for enabling Google Sign-In.
// Google Sign-In documentation: https://developers.google.com/identity/sign-in/web/
func (o ListWebAppAuthSettingsResultOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

// The client secret associated with the Google web application.
// This setting is required for enabling Google Sign-In.
// Google Sign-In documentation: https://developers.google.com/identity/sign-in/web/
func (o ListWebAppAuthSettingsResultOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

// The app setting name that contains the client secret associated with
// the Google web application.
func (o ListWebAppAuthSettingsResultOutput) GoogleClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GoogleClientSecretSettingName }).(pulumi.StringPtrOutput)
}

// The OAuth 2.0 scopes that will be requested as part of Google Sign-In authentication.
// This setting is optional. If not specified, "openid", "profile", and "email" are used as default scopes.
// Google Sign-In documentation: https://developers.google.com/identity/sign-in/web/
func (o ListWebAppAuthSettingsResultOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

// Resource Id.
func (o ListWebAppAuthSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// "true" if the auth config settings should be read from a file,
// "false" otherwise
func (o ListWebAppAuthSettingsResultOutput) IsAuthFromFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.IsAuthFromFile }).(pulumi.StringPtrOutput)
}

// The OpenID Connect Issuer URI that represents the entity which issues access tokens for this application.
// When using Azure Active Directory, this value is the URI of the directory tenant, e.g. `https://sts.windows.net/{tenant-guid}/`.
// This URI is a case-sensitive identifier for the token issuer.
// More information on OpenID Connect Discovery: http://openid.net/specs/openid-connect-discovery-1_0.html
func (o ListWebAppAuthSettingsResultOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o ListWebAppAuthSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The OAuth 2.0 client ID that was created for the app used for authentication.
// This setting is required for enabling Microsoft Account authentication.
// Microsoft Account OAuth documentation: https://dev.onedrive.com/auth/msa_oauth.htm
func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

// The OAuth 2.0 client secret that was created for the app used for authentication.
// This setting is required for enabling Microsoft Account authentication.
// Microsoft Account OAuth documentation: https://dev.onedrive.com/auth/msa_oauth.htm
func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

// The app setting name containing the OAuth 2.0 client secret that was created for the
// app used for authentication.
func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.MicrosoftAccountClientSecretSettingName }).(pulumi.StringPtrOutput)
}

// The OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication.
// This setting is optional. If not specified, "wl.basic" is used as the default scope.
// Microsoft Account Scopes and permissions documentation: https://msdn.microsoft.com/en-us/library/dn631845.aspx
func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

// Resource Name.
func (o ListWebAppAuthSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The RuntimeVersion of the Authentication / Authorization feature in use for the current app.
// The setting in this value can control the behavior of certain features in the Authentication / Authorization module.
func (o ListWebAppAuthSettingsResultOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

// The number of hours after session token expiration that a session token can be used to
// call the token refresh API. The default is 72 hours.
func (o ListWebAppAuthSettingsResultOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

// <code>true</code> to durably store platform-specific security tokens that are obtained during login flows; otherwise, <code>false</code>.
//
//	The default is <code>false</code>.
func (o ListWebAppAuthSettingsResultOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *bool { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
// This setting is required for enabling Twitter Sign-In.
// Twitter Sign-In documentation: https://dev.twitter.com/web/sign-in
func (o ListWebAppAuthSettingsResultOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

// The OAuth 1.0a consumer secret of the Twitter application used for sign-in.
// This setting is required for enabling Twitter Sign-In.
// Twitter Sign-In documentation: https://dev.twitter.com/web/sign-in
func (o ListWebAppAuthSettingsResultOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

// The app setting name that contains the OAuth 1.0a consumer secret of the Twitter
// application used for sign-in.
func (o ListWebAppAuthSettingsResultOutput) TwitterConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.TwitterConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ListWebAppAuthSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

// The action to take when an unauthenticated client attempts to access the app.
func (o ListWebAppAuthSettingsResultOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

// Gets a value indicating whether the issuer should be a valid HTTPS url and be validated as such.
func (o ListWebAppAuthSettingsResultOutput) ValidateIssuer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *bool { return v.ValidateIssuer }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppAuthSettingsResultOutput{})
}
