// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package portal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get current user settings for current signed in user. This operation returns settings for the user's cloud shell preferences including preferred location, storage profile, shell type, font and size settings.
//
// Uses Azure REST API version 2018-10-01.
func LookupUserSettingsWithLocation(ctx *pulumi.Context, args *LookupUserSettingsWithLocationArgs, opts ...pulumi.InvokeOption) (*LookupUserSettingsWithLocationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupUserSettingsWithLocationResult
	err := ctx.Invoke("azure-native:portal:getUserSettingsWithLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserSettingsWithLocationArgs struct {
	// The provider location
	Location string `pulumi:"location"`
	// The name of the user settings
	UserSettingsName string `pulumi:"userSettingsName"`
}

// Response to get user settings
type LookupUserSettingsWithLocationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The cloud shell user settings properties.
	Properties UserPropertiesResponse `pulumi:"properties"`
}

func LookupUserSettingsWithLocationOutput(ctx *pulumi.Context, args LookupUserSettingsWithLocationOutputArgs, opts ...pulumi.InvokeOption) LookupUserSettingsWithLocationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupUserSettingsWithLocationResultOutput, error) {
			args := v.(LookupUserSettingsWithLocationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:portal:getUserSettingsWithLocation", args, LookupUserSettingsWithLocationResultOutput{}, options).(LookupUserSettingsWithLocationResultOutput), nil
		}).(LookupUserSettingsWithLocationResultOutput)
}

type LookupUserSettingsWithLocationOutputArgs struct {
	// The provider location
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the user settings
	UserSettingsName pulumi.StringInput `pulumi:"userSettingsName"`
}

func (LookupUserSettingsWithLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsWithLocationArgs)(nil)).Elem()
}

// Response to get user settings
type LookupUserSettingsWithLocationResultOutput struct{ *pulumi.OutputState }

func (LookupUserSettingsWithLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsWithLocationResult)(nil)).Elem()
}

func (o LookupUserSettingsWithLocationResultOutput) ToLookupUserSettingsWithLocationResultOutput() LookupUserSettingsWithLocationResultOutput {
	return o
}

func (o LookupUserSettingsWithLocationResultOutput) ToLookupUserSettingsWithLocationResultOutputWithContext(ctx context.Context) LookupUserSettingsWithLocationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupUserSettingsWithLocationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserSettingsWithLocationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The cloud shell user settings properties.
func (o LookupUserSettingsWithLocationResultOutput) Properties() UserPropertiesResponseOutput {
	return o.ApplyT(func(v LookupUserSettingsWithLocationResult) UserPropertiesResponse { return v.Properties }).(UserPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserSettingsWithLocationResultOutput{})
}
