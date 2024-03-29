// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

type UserApiKeyResponsePropertiesResponse struct {
	// The User Api Key Generated based on GenerateApiKey flag. This is applicable for non-Portal clients only.
	ApiKey *string `pulumi:"apiKey"`
}

type UserApiKeyResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserApiKeyResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserApiKeyResponsePropertiesResponse)(nil)).Elem()
}

func (o UserApiKeyResponsePropertiesResponseOutput) ToUserApiKeyResponsePropertiesResponseOutput() UserApiKeyResponsePropertiesResponseOutput {
	return o
}

func (o UserApiKeyResponsePropertiesResponseOutput) ToUserApiKeyResponsePropertiesResponseOutputWithContext(ctx context.Context) UserApiKeyResponsePropertiesResponseOutput {
	return o
}

// The User Api Key Generated based on GenerateApiKey flag. This is applicable for non-Portal clients only.
func (o UserApiKeyResponsePropertiesResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserApiKeyResponsePropertiesResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UserApiKeyResponsePropertiesResponseOutput{})
}
