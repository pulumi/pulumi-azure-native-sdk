// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Shared Access Authorization Token for the User.
//
// Uses Azure REST API version 2022-09-01-preview.
//
// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func GetUserSharedAccessToken(ctx *pulumi.Context, args *GetUserSharedAccessTokenArgs, opts ...pulumi.InvokeOption) (*GetUserSharedAccessTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetUserSharedAccessTokenResult
	err := ctx.Invoke("azure-native:apimanagement:getUserSharedAccessToken", args.Defaults(), &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetUserSharedAccessTokenArgs struct {
	// The Expiry time of the Token. Maximum token expiry time is set to 30 days. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Expiry string `pulumi:"expiry"`
	// The Key to be used to generate token for user.
	KeyType KeyType `pulumi:"keyType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId string `pulumi:"userId"`
}

// Defaults sets the appropriate defaults for GetUserSharedAccessTokenArgs
func (val *GetUserSharedAccessTokenArgs) Defaults() *GetUserSharedAccessTokenArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if utilities.IsZero(tmp.KeyType) {
		tmp.KeyType = KeyType("primary")
	}
	return &tmp
}

// Get User Token response details.
type GetUserSharedAccessTokenResult struct {
	// Shared Access Authorization token for the User.
	Value *string `pulumi:"value"`
}

func GetUserSharedAccessTokenOutput(ctx *pulumi.Context, args GetUserSharedAccessTokenOutputArgs, opts ...pulumi.InvokeOption) GetUserSharedAccessTokenResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetUserSharedAccessTokenResultOutput, error) {
			args := v.(GetUserSharedAccessTokenArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getUserSharedAccessToken", args.Defaults(), GetUserSharedAccessTokenResultOutput{}, options).(GetUserSharedAccessTokenResultOutput), nil
		}).(GetUserSharedAccessTokenResultOutput)
}

type GetUserSharedAccessTokenOutputArgs struct {
	// The Expiry time of the Token. Maximum token expiry time is set to 30 days. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Expiry pulumi.StringInput `pulumi:"expiry"`
	// The Key to be used to generate token for user.
	KeyType KeyTypeInput `pulumi:"keyType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// User identifier. Must be unique in the current API Management service instance.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (GetUserSharedAccessTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSharedAccessTokenArgs)(nil)).Elem()
}

// Get User Token response details.
type GetUserSharedAccessTokenResultOutput struct{ *pulumi.OutputState }

func (GetUserSharedAccessTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSharedAccessTokenResult)(nil)).Elem()
}

func (o GetUserSharedAccessTokenResultOutput) ToGetUserSharedAccessTokenResultOutput() GetUserSharedAccessTokenResultOutput {
	return o
}

func (o GetUserSharedAccessTokenResultOutput) ToGetUserSharedAccessTokenResultOutputWithContext(ctx context.Context) GetUserSharedAccessTokenResultOutput {
	return o
}

// Shared Access Authorization token for the User.
func (o GetUserSharedAccessTokenResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserSharedAccessTokenResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserSharedAccessTokenResultOutput{})
}
