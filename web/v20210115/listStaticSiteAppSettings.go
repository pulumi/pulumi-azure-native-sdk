// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// String dictionary resource.
func ListStaticSiteAppSettings(ctx *pulumi.Context, args *ListStaticSiteAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteAppSettingsResult, error) {
	var rv ListStaticSiteAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20210115:listStaticSiteAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteAppSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListStaticSiteAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListStaticSiteAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteAppSettingsResult, error) {
			args := v.(ListStaticSiteAppSettingsArgs)
			r, err := ListStaticSiteAppSettings(ctx, &args, opts...)
			var s ListStaticSiteAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteAppSettingsResultOutput)
}

type ListStaticSiteAppSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteAppSettingsArgs)(nil)).Elem()
}

// String dictionary resource.
type ListStaticSiteAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteAppSettingsResultOutput) ToListStaticSiteAppSettingsResultOutput() ListStaticSiteAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteAppSettingsResultOutput) ToListStaticSiteAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListStaticSiteAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListStaticSiteAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListStaticSiteAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListStaticSiteAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteAppSettingsResultOutput{})
}