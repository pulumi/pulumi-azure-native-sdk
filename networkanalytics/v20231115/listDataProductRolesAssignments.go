// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List user roles associated with the data product.
func ListDataProductRolesAssignments(ctx *pulumi.Context, args *ListDataProductRolesAssignmentsArgs, opts ...pulumi.InvokeOption) (*ListDataProductRolesAssignmentsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDataProductRolesAssignmentsResult
	err := ctx.Invoke("azure-native:networkanalytics/v20231115:listDataProductRolesAssignments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDataProductRolesAssignmentsArgs struct {
	// The data product resource name
	DataProductName string `pulumi:"dataProductName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// list role assignments.
type ListDataProductRolesAssignmentsResult struct {
	// Count of role assignments.
	Count int `pulumi:"count"`
	// list of role assignments
	RoleAssignmentResponse []RoleAssignmentDetailResponse `pulumi:"roleAssignmentResponse"`
}

func ListDataProductRolesAssignmentsOutput(ctx *pulumi.Context, args ListDataProductRolesAssignmentsOutputArgs, opts ...pulumi.InvokeOption) ListDataProductRolesAssignmentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDataProductRolesAssignmentsResultOutput, error) {
			args := v.(ListDataProductRolesAssignmentsArgs)
			opts = utilities.PkgInvokeDefaultOpts(opts)
			var rv ListDataProductRolesAssignmentsResult
			secret, err := ctx.InvokePackageRaw("azure-native:networkanalytics/v20231115:listDataProductRolesAssignments", args, &rv, "", opts...)
			if err != nil {
				return ListDataProductRolesAssignmentsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(ListDataProductRolesAssignmentsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(ListDataProductRolesAssignmentsResultOutput), nil
			}
			return output, nil
		}).(ListDataProductRolesAssignmentsResultOutput)
}

type ListDataProductRolesAssignmentsOutputArgs struct {
	// The data product resource name
	DataProductName pulumi.StringInput `pulumi:"dataProductName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDataProductRolesAssignmentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDataProductRolesAssignmentsArgs)(nil)).Elem()
}

// list role assignments.
type ListDataProductRolesAssignmentsResultOutput struct{ *pulumi.OutputState }

func (ListDataProductRolesAssignmentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDataProductRolesAssignmentsResult)(nil)).Elem()
}

func (o ListDataProductRolesAssignmentsResultOutput) ToListDataProductRolesAssignmentsResultOutput() ListDataProductRolesAssignmentsResultOutput {
	return o
}

func (o ListDataProductRolesAssignmentsResultOutput) ToListDataProductRolesAssignmentsResultOutputWithContext(ctx context.Context) ListDataProductRolesAssignmentsResultOutput {
	return o
}

// Count of role assignments.
func (o ListDataProductRolesAssignmentsResultOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v ListDataProductRolesAssignmentsResult) int { return v.Count }).(pulumi.IntOutput)
}

// list of role assignments
func (o ListDataProductRolesAssignmentsResultOutput) RoleAssignmentResponse() RoleAssignmentDetailResponseArrayOutput {
	return o.ApplyT(func(v ListDataProductRolesAssignmentsResult) []RoleAssignmentDetailResponse {
		return v.RoleAssignmentResponse
	}).(RoleAssignmentDetailResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDataProductRolesAssignmentsResultOutput{})
}
