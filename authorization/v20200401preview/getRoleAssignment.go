// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Role Assignments
func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20200401preview:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	RoleAssignmentName string  `pulumi:"roleAssignmentName"`
	Scope              string  `pulumi:"scope"`
	TenantId           *string `pulumi:"tenantId"`
}

// Role Assignments
type LookupRoleAssignmentResult struct {
	CanDelegate                        *bool   `pulumi:"canDelegate"`
	Condition                          *string `pulumi:"condition"`
	ConditionVersion                   *string `pulumi:"conditionVersion"`
	CreatedBy                          *string `pulumi:"createdBy"`
	CreatedOn                          *string `pulumi:"createdOn"`
	DelegatedManagedIdentityResourceId *string `pulumi:"delegatedManagedIdentityResourceId"`
	Description                        *string `pulumi:"description"`
	Id                                 string  `pulumi:"id"`
	Name                               string  `pulumi:"name"`
	PrincipalId                        *string `pulumi:"principalId"`
	PrincipalType                      *string `pulumi:"principalType"`
	RoleDefinitionId                   *string `pulumi:"roleDefinitionId"`
	Scope                              *string `pulumi:"scope"`
	Type                               string  `pulumi:"type"`
	UpdatedBy                          *string `pulumi:"updatedBy"`
	UpdatedOn                          *string `pulumi:"updatedOn"`
}

func LookupRoleAssignmentOutput(ctx *pulumi.Context, args LookupRoleAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupRoleAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoleAssignmentResult, error) {
			args := v.(LookupRoleAssignmentArgs)
			r, err := LookupRoleAssignment(ctx, &args, opts...)
			var s LookupRoleAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoleAssignmentResultOutput)
}

type LookupRoleAssignmentOutputArgs struct {
	RoleAssignmentName pulumi.StringInput    `pulumi:"roleAssignmentName"`
	Scope              pulumi.StringInput    `pulumi:"scope"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (LookupRoleAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentArgs)(nil)).Elem()
}

// Role Assignments
type LookupRoleAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupRoleAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoleAssignmentResult)(nil)).Elem()
}

func (o LookupRoleAssignmentResultOutput) ToLookupRoleAssignmentResultOutput() LookupRoleAssignmentResultOutput {
	return o
}

func (o LookupRoleAssignmentResultOutput) ToLookupRoleAssignmentResultOutputWithContext(ctx context.Context) LookupRoleAssignmentResultOutput {
	return o
}

func (o LookupRoleAssignmentResultOutput) CanDelegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *bool { return v.CanDelegate }).(pulumi.BoolPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) ConditionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.ConditionVersion }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) CreatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.CreatedOn }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) DelegatedManagedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.DelegatedManagedIdentityResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRoleAssignmentResultOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.UpdatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupRoleAssignmentResultOutput) UpdatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoleAssignmentResult) *string { return v.UpdatedOn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoleAssignmentResultOutput{})
}