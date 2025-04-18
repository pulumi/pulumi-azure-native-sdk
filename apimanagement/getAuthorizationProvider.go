// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the authorization provider specified by its identifier.
//
// Uses Azure REST API version 2022-09-01-preview.
//
// Other available API versions: 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupAuthorizationProvider(ctx *pulumi.Context, args *LookupAuthorizationProviderArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationProviderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationProviderResult
	err := ctx.Invoke("azure-native:apimanagement:getAuthorizationProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizationProviderArgs struct {
	// Identifier of the authorization provider.
	AuthorizationProviderId string `pulumi:"authorizationProviderId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Authorization Provider contract.
type LookupAuthorizationProviderResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Authorization Provider name. Must be 1 to 300 characters long.
	DisplayName *string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity provider name. Must be 1 to 300 characters long.
	IdentityProvider *string `pulumi:"identityProvider"`
	// The name of the resource
	Name string `pulumi:"name"`
	// OAuth2 settings
	Oauth2 *AuthorizationProviderOAuth2SettingsResponse `pulumi:"oauth2"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAuthorizationProviderOutput(ctx *pulumi.Context, args LookupAuthorizationProviderOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationProviderResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationProviderResultOutput, error) {
			args := v.(LookupAuthorizationProviderArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getAuthorizationProvider", args, LookupAuthorizationProviderResultOutput{}, options).(LookupAuthorizationProviderResultOutput), nil
		}).(LookupAuthorizationProviderResultOutput)
}

type LookupAuthorizationProviderOutputArgs struct {
	// Identifier of the authorization provider.
	AuthorizationProviderId pulumi.StringInput `pulumi:"authorizationProviderId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupAuthorizationProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProviderArgs)(nil)).Elem()
}

// Authorization Provider contract.
type LookupAuthorizationProviderResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProviderResult)(nil)).Elem()
}

func (o LookupAuthorizationProviderResultOutput) ToLookupAuthorizationProviderResultOutput() LookupAuthorizationProviderResultOutput {
	return o
}

func (o LookupAuthorizationProviderResultOutput) ToLookupAuthorizationProviderResultOutputWithContext(ctx context.Context) LookupAuthorizationProviderResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAuthorizationProviderResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Authorization Provider name. Must be 1 to 300 characters long.
func (o LookupAuthorizationProviderResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAuthorizationProviderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity provider name. Must be 1 to 300 characters long.
func (o LookupAuthorizationProviderResultOutput) IdentityProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *string { return v.IdentityProvider }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupAuthorizationProviderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Name }).(pulumi.StringOutput)
}

// OAuth2 settings
func (o LookupAuthorizationProviderResultOutput) Oauth2() AuthorizationProviderOAuth2SettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) *AuthorizationProviderOAuth2SettingsResponse {
		return v.Oauth2
	}).(AuthorizationProviderOAuth2SettingsResponsePtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAuthorizationProviderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProviderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationProviderResultOutput{})
}
