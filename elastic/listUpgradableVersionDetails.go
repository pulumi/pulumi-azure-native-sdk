// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Stack Versions that this version can upgrade to
// API Version: 2021-10-01-preview.
func ListUpgradableVersionDetails(ctx *pulumi.Context, args *ListUpgradableVersionDetailsArgs, opts ...pulumi.InvokeOption) (*ListUpgradableVersionDetailsResult, error) {
	var rv ListUpgradableVersionDetailsResult
	err := ctx.Invoke("azure-native:elastic:listUpgradableVersionDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListUpgradableVersionDetailsArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Stack Versions that this version can upgrade to
type ListUpgradableVersionDetailsResult struct {
	CurrentVersion     *string  `pulumi:"currentVersion"`
	UpgradableVersions []string `pulumi:"upgradableVersions"`
}

func ListUpgradableVersionDetailsOutput(ctx *pulumi.Context, args ListUpgradableVersionDetailsOutputArgs, opts ...pulumi.InvokeOption) ListUpgradableVersionDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListUpgradableVersionDetailsResult, error) {
			args := v.(ListUpgradableVersionDetailsArgs)
			r, err := ListUpgradableVersionDetails(ctx, &args, opts...)
			var s ListUpgradableVersionDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListUpgradableVersionDetailsResultOutput)
}

type ListUpgradableVersionDetailsOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListUpgradableVersionDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUpgradableVersionDetailsArgs)(nil)).Elem()
}

// Stack Versions that this version can upgrade to
type ListUpgradableVersionDetailsResultOutput struct{ *pulumi.OutputState }

func (ListUpgradableVersionDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListUpgradableVersionDetailsResult)(nil)).Elem()
}

func (o ListUpgradableVersionDetailsResultOutput) ToListUpgradableVersionDetailsResultOutput() ListUpgradableVersionDetailsResultOutput {
	return o
}

func (o ListUpgradableVersionDetailsResultOutput) ToListUpgradableVersionDetailsResultOutputWithContext(ctx context.Context) ListUpgradableVersionDetailsResultOutput {
	return o
}

func (o ListUpgradableVersionDetailsResultOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListUpgradableVersionDetailsResult) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o ListUpgradableVersionDetailsResultOutput) UpgradableVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListUpgradableVersionDetailsResult) []string { return v.UpgradableVersions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListUpgradableVersionDetailsResultOutput{})
}