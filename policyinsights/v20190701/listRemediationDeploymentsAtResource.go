// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List of deployments for a remediation.
func ListRemediationDeploymentsAtResource(ctx *pulumi.Context, args *ListRemediationDeploymentsAtResourceArgs, opts ...pulumi.InvokeOption) (*ListRemediationDeploymentsAtResourceResult, error) {
	var rv ListRemediationDeploymentsAtResourceResult
	err := ctx.Invoke("azure-native:policyinsights/v20190701:listRemediationDeploymentsAtResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemediationDeploymentsAtResourceArgs struct {
	RemediationName string `pulumi:"remediationName"`
	ResourceId      string `pulumi:"resourceId"`
	Top             *int   `pulumi:"top"`
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtResourceResult struct {
	NextLink string                          `pulumi:"nextLink"`
	Value    []RemediationDeploymentResponse `pulumi:"value"`
}

func ListRemediationDeploymentsAtResourceOutput(ctx *pulumi.Context, args ListRemediationDeploymentsAtResourceOutputArgs, opts ...pulumi.InvokeOption) ListRemediationDeploymentsAtResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemediationDeploymentsAtResourceResult, error) {
			args := v.(ListRemediationDeploymentsAtResourceArgs)
			r, err := ListRemediationDeploymentsAtResource(ctx, &args, opts...)
			var s ListRemediationDeploymentsAtResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemediationDeploymentsAtResourceResultOutput)
}

type ListRemediationDeploymentsAtResourceOutputArgs struct {
	RemediationName pulumi.StringInput `pulumi:"remediationName"`
	ResourceId      pulumi.StringInput `pulumi:"resourceId"`
	Top             pulumi.IntPtrInput `pulumi:"top"`
}

func (ListRemediationDeploymentsAtResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtResourceArgs)(nil)).Elem()
}

// List of deployments for a remediation.
type ListRemediationDeploymentsAtResourceResultOutput struct{ *pulumi.OutputState }

func (ListRemediationDeploymentsAtResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemediationDeploymentsAtResourceResult)(nil)).Elem()
}

func (o ListRemediationDeploymentsAtResourceResultOutput) ToListRemediationDeploymentsAtResourceResultOutput() ListRemediationDeploymentsAtResourceResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtResourceResultOutput) ToListRemediationDeploymentsAtResourceResultOutputWithContext(ctx context.Context) ListRemediationDeploymentsAtResourceResultOutput {
	return o
}

func (o ListRemediationDeploymentsAtResourceResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtResourceResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListRemediationDeploymentsAtResourceResultOutput) Value() RemediationDeploymentResponseArrayOutput {
	return o.ApplyT(func(v ListRemediationDeploymentsAtResourceResult) []RemediationDeploymentResponse { return v.Value }).(RemediationDeploymentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemediationDeploymentsAtResourceResultOutput{})
}