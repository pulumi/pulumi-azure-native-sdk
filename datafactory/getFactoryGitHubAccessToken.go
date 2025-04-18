// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get GitHub Access Token.
//
// Uses Azure REST API version 2018-06-01.
func GetFactoryGitHubAccessToken(ctx *pulumi.Context, args *GetFactoryGitHubAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetFactoryGitHubAccessTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetFactoryGitHubAccessTokenResult
	err := ctx.Invoke("azure-native:datafactory:getFactoryGitHubAccessToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFactoryGitHubAccessTokenArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// GitHub access code.
	GitHubAccessCode string `pulumi:"gitHubAccessCode"`
	// GitHub access token base URL.
	GitHubAccessTokenBaseUrl string `pulumi:"gitHubAccessTokenBaseUrl"`
	// GitHub application client ID.
	GitHubClientId *string `pulumi:"gitHubClientId"`
	// GitHub bring your own app client secret information.
	GitHubClientSecret *GitHubClientSecret `pulumi:"gitHubClientSecret"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Get GitHub access token response definition.
type GetFactoryGitHubAccessTokenResult struct {
	// GitHub access token.
	GitHubAccessToken *string `pulumi:"gitHubAccessToken"`
}

func GetFactoryGitHubAccessTokenOutput(ctx *pulumi.Context, args GetFactoryGitHubAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetFactoryGitHubAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFactoryGitHubAccessTokenResultOutput, error) {
			args := v.(GetFactoryGitHubAccessTokenArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datafactory:getFactoryGitHubAccessToken", args, GetFactoryGitHubAccessTokenResultOutput{}, options).(GetFactoryGitHubAccessTokenResultOutput), nil
		}).(GetFactoryGitHubAccessTokenResultOutput)
}

type GetFactoryGitHubAccessTokenOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// GitHub access code.
	GitHubAccessCode pulumi.StringInput `pulumi:"gitHubAccessCode"`
	// GitHub access token base URL.
	GitHubAccessTokenBaseUrl pulumi.StringInput `pulumi:"gitHubAccessTokenBaseUrl"`
	// GitHub application client ID.
	GitHubClientId pulumi.StringPtrInput `pulumi:"gitHubClientId"`
	// GitHub bring your own app client secret information.
	GitHubClientSecret GitHubClientSecretPtrInput `pulumi:"gitHubClientSecret"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetFactoryGitHubAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFactoryGitHubAccessTokenArgs)(nil)).Elem()
}

// Get GitHub access token response definition.
type GetFactoryGitHubAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetFactoryGitHubAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFactoryGitHubAccessTokenResult)(nil)).Elem()
}

func (o GetFactoryGitHubAccessTokenResultOutput) ToGetFactoryGitHubAccessTokenResultOutput() GetFactoryGitHubAccessTokenResultOutput {
	return o
}

func (o GetFactoryGitHubAccessTokenResultOutput) ToGetFactoryGitHubAccessTokenResultOutputWithContext(ctx context.Context) GetFactoryGitHubAccessTokenResultOutput {
	return o
}

// GitHub access token.
func (o GetFactoryGitHubAccessTokenResultOutput) GitHubAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFactoryGitHubAccessTokenResult) *string { return v.GitHubAccessToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFactoryGitHubAccessTokenResultOutput{})
}
