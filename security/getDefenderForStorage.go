// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Defender for Storage settings for the specified storage account.
//
// Uses Azure REST API version 2024-10-01-preview.
//
// Other available API versions: 2022-12-01-preview, 2024-08-01-preview, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupDefenderForStorage(ctx *pulumi.Context, args *LookupDefenderForStorageArgs, opts ...pulumi.InvokeOption) (*LookupDefenderForStorageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDefenderForStorageResult
	err := ctx.Invoke("azure-native:security:getDefenderForStorage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefenderForStorageArgs struct {
	// The identifier of the resource.
	ResourceId string `pulumi:"resourceId"`
	// Defender for Storage setting name.
	SettingName string `pulumi:"settingName"`
}

// The Defender for Storage resource.
type LookupDefenderForStorageResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// Defender for Storage resource properties.
	Properties DefenderForStorageSettingPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupDefenderForStorageOutput(ctx *pulumi.Context, args LookupDefenderForStorageOutputArgs, opts ...pulumi.InvokeOption) LookupDefenderForStorageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDefenderForStorageResultOutput, error) {
			args := v.(LookupDefenderForStorageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:security:getDefenderForStorage", args, LookupDefenderForStorageResultOutput{}, options).(LookupDefenderForStorageResultOutput), nil
		}).(LookupDefenderForStorageResultOutput)
}

type LookupDefenderForStorageOutputArgs struct {
	// The identifier of the resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Defender for Storage setting name.
	SettingName pulumi.StringInput `pulumi:"settingName"`
}

func (LookupDefenderForStorageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefenderForStorageArgs)(nil)).Elem()
}

// The Defender for Storage resource.
type LookupDefenderForStorageResultOutput struct{ *pulumi.OutputState }

func (LookupDefenderForStorageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefenderForStorageResult)(nil)).Elem()
}

func (o LookupDefenderForStorageResultOutput) ToLookupDefenderForStorageResultOutput() LookupDefenderForStorageResultOutput {
	return o
}

func (o LookupDefenderForStorageResultOutput) ToLookupDefenderForStorageResultOutputWithContext(ctx context.Context) LookupDefenderForStorageResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDefenderForStorageResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefenderForStorageResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupDefenderForStorageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefenderForStorageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupDefenderForStorageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefenderForStorageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Defender for Storage resource properties.
func (o LookupDefenderForStorageResultOutput) Properties() DefenderForStorageSettingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupDefenderForStorageResult) DefenderForStorageSettingPropertiesResponse {
		return v.Properties
	}).(DefenderForStorageSettingPropertiesResponseOutput)
}

// Resource type
func (o LookupDefenderForStorageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefenderForStorageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefenderForStorageResultOutput{})
}
