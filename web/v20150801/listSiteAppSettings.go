// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// String dictionary resource
func ListSiteAppSettings(ctx *pulumi.Context, args *ListSiteAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListSiteAppSettingsResult, error) {
	var rv ListSiteAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAppSettingsArgs struct {
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource
type ListSiteAppSettingsResult struct {
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Settings
	Properties map[string]string `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

func ListSiteAppSettingsOutput(ctx *pulumi.Context, args ListSiteAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListSiteAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteAppSettingsResult, error) {
			args := v.(ListSiteAppSettingsArgs)
			r, err := ListSiteAppSettings(ctx, &args, opts...)
			var s ListSiteAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteAppSettingsResultOutput)
}

type ListSiteAppSettingsOutputArgs struct {
	// Name of web app
	Name pulumi.StringInput `pulumi:"name"`
	// Name of resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSiteAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAppSettingsArgs)(nil)).Elem()
}

// String dictionary resource
type ListSiteAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListSiteAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAppSettingsResult)(nil)).Elem()
}

func (o ListSiteAppSettingsResultOutput) ToListSiteAppSettingsResultOutput() ListSiteAppSettingsResultOutput {
	return o
}

func (o ListSiteAppSettingsResultOutput) ToListSiteAppSettingsResultOutputWithContext(ctx context.Context) ListSiteAppSettingsResultOutput {
	return o
}

// Resource Id
func (o ListSiteAppSettingsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Kind of resource
func (o ListSiteAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o ListSiteAppSettingsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o ListSiteAppSettingsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Settings
func (o ListSiteAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource tags
func (o ListSiteAppSettingsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ListSiteAppSettingsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteAppSettingsResultOutput{})
}
