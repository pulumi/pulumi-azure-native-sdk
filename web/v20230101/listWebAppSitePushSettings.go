// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the Push settings associated with web app.
func ListWebAppSitePushSettings(ctx *pulumi.Context, args *ListWebAppSitePushSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppSitePushSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppSitePushSettingsResult
	err := ctx.Invoke("azure-native:web/v20230101:listWebAppSitePushSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSitePushSettingsArgs struct {
	// Name of web app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Push settings for the App.
type ListWebAppSitePushSettingsResult struct {
	// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
	DynamicTagsJson *string `pulumi:"dynamicTagsJson"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Gets or sets a flag indicating whether the Push endpoint is enabled.
	IsPushEnabled bool `pulumi:"isPushEnabled"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
	TagWhitelistJson *string `pulumi:"tagWhitelistJson"`
	// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
	// Tags can consist of alphanumeric characters and the following:
	// '_', '@', '#', '.', ':', '-'.
	// Validation should be performed at the PushRequestHandler.
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppSitePushSettingsOutput(ctx *pulumi.Context, args ListWebAppSitePushSettingsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSitePushSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWebAppSitePushSettingsResultOutput, error) {
			args := v.(ListWebAppSitePushSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web/v20230101:listWebAppSitePushSettings", args, ListWebAppSitePushSettingsResultOutput{}, options).(ListWebAppSitePushSettingsResultOutput), nil
		}).(ListWebAppSitePushSettingsResultOutput)
}

type ListWebAppSitePushSettingsOutputArgs struct {
	// Name of web app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppSitePushSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSitePushSettingsArgs)(nil)).Elem()
}

// Push settings for the App.
type ListWebAppSitePushSettingsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSitePushSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSitePushSettingsResult)(nil)).Elem()
}

func (o ListWebAppSitePushSettingsResultOutput) ToListWebAppSitePushSettingsResultOutput() ListWebAppSitePushSettingsResultOutput {
	return o
}

func (o ListWebAppSitePushSettingsResultOutput) ToListWebAppSitePushSettingsResultOutputWithContext(ctx context.Context) ListWebAppSitePushSettingsResultOutput {
	return o
}

// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
func (o ListWebAppSitePushSettingsResultOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o ListWebAppSitePushSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets a flag indicating whether the Push endpoint is enabled.
func (o ListWebAppSitePushSettingsResultOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

// Kind of resource.
func (o ListWebAppSitePushSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppSitePushSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
func (o ListWebAppSitePushSettingsResultOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
// Tags can consist of alphanumeric characters and the following:
// '_', '@', '#', '.', ':', '-'.
// Validation should be performed at the PushRequestHandler.
func (o ListWebAppSitePushSettingsResultOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o ListWebAppSitePushSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSitePushSettingsResultOutput{})
}
