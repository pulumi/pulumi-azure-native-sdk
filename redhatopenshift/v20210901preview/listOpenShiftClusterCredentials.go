// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OpenShiftClusterCredentials represents an OpenShift cluster's credentials.
func ListOpenShiftClusterCredentials(ctx *pulumi.Context, args *ListOpenShiftClusterCredentialsArgs, opts ...pulumi.InvokeOption) (*ListOpenShiftClusterCredentialsResult, error) {
	var rv ListOpenShiftClusterCredentialsResult
	err := ctx.Invoke("azure-native:redhatopenshift/v20210901preview:listOpenShiftClusterCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOpenShiftClusterCredentialsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}

// OpenShiftClusterCredentials represents an OpenShift cluster's credentials.
type ListOpenShiftClusterCredentialsResult struct {
	KubeadminPassword *string `pulumi:"kubeadminPassword"`
	KubeadminUsername *string `pulumi:"kubeadminUsername"`
}

func ListOpenShiftClusterCredentialsOutput(ctx *pulumi.Context, args ListOpenShiftClusterCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListOpenShiftClusterCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOpenShiftClusterCredentialsResult, error) {
			args := v.(ListOpenShiftClusterCredentialsArgs)
			r, err := ListOpenShiftClusterCredentials(ctx, &args, opts...)
			var s ListOpenShiftClusterCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOpenShiftClusterCredentialsResultOutput)
}

type ListOpenShiftClusterCredentialsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (ListOpenShiftClusterCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterCredentialsArgs)(nil)).Elem()
}

// OpenShiftClusterCredentials represents an OpenShift cluster's credentials.
type ListOpenShiftClusterCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListOpenShiftClusterCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOpenShiftClusterCredentialsResult)(nil)).Elem()
}

func (o ListOpenShiftClusterCredentialsResultOutput) ToListOpenShiftClusterCredentialsResultOutput() ListOpenShiftClusterCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterCredentialsResultOutput) ToListOpenShiftClusterCredentialsResultOutputWithContext(ctx context.Context) ListOpenShiftClusterCredentialsResultOutput {
	return o
}

func (o ListOpenShiftClusterCredentialsResultOutput) KubeadminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenShiftClusterCredentialsResult) *string { return v.KubeadminPassword }).(pulumi.StringPtrOutput)
}

func (o ListOpenShiftClusterCredentialsResultOutput) KubeadminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOpenShiftClusterCredentialsResult) *string { return v.KubeadminUsername }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOpenShiftClusterCredentialsResultOutput{})
}