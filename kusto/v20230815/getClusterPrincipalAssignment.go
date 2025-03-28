// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Kusto cluster principalAssignment.
func LookupClusterPrincipalAssignment(ctx *pulumi.Context, args *LookupClusterPrincipalAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupClusterPrincipalAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupClusterPrincipalAssignmentResult
	err := ctx.Invoke("azure-native:kusto/v20230815:getClusterPrincipalAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterPrincipalAssignmentArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName string `pulumi:"principalAssignmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Class representing a cluster principal assignment.
type LookupClusterPrincipalAssignmentResult struct {
	// The service principal object id in AAD (Azure active directory)
	AadObjectId string `pulumi:"aadObjectId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// The principal name
	PrincipalName string `pulumi:"principalName"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Cluster principal role.
	Role string `pulumi:"role"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName string `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupClusterPrincipalAssignmentOutput(ctx *pulumi.Context, args LookupClusterPrincipalAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupClusterPrincipalAssignmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupClusterPrincipalAssignmentResultOutput, error) {
			args := v.(LookupClusterPrincipalAssignmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:kusto/v20230815:getClusterPrincipalAssignment", args, LookupClusterPrincipalAssignmentResultOutput{}, options).(LookupClusterPrincipalAssignmentResultOutput), nil
		}).(LookupClusterPrincipalAssignmentResultOutput)
}

type LookupClusterPrincipalAssignmentOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringInput `pulumi:"principalAssignmentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupClusterPrincipalAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterPrincipalAssignmentArgs)(nil)).Elem()
}

// Class representing a cluster principal assignment.
type LookupClusterPrincipalAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupClusterPrincipalAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterPrincipalAssignmentResult)(nil)).Elem()
}

func (o LookupClusterPrincipalAssignmentResultOutput) ToLookupClusterPrincipalAssignmentResultOutput() LookupClusterPrincipalAssignmentResultOutput {
	return o
}

func (o LookupClusterPrincipalAssignmentResultOutput) ToLookupClusterPrincipalAssignmentResultOutputWithContext(ctx context.Context) LookupClusterPrincipalAssignmentResultOutput {
	return o
}

// The service principal object id in AAD (Azure active directory)
func (o LookupClusterPrincipalAssignmentResultOutput) AadObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.AadObjectId }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupClusterPrincipalAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupClusterPrincipalAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
func (o LookupClusterPrincipalAssignmentResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o LookupClusterPrincipalAssignmentResultOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o LookupClusterPrincipalAssignmentResultOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o LookupClusterPrincipalAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Cluster principal role.
func (o LookupClusterPrincipalAssignmentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Role }).(pulumi.StringOutput)
}

// The tenant id of the principal
func (o LookupClusterPrincipalAssignmentResultOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o LookupClusterPrincipalAssignmentResultOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupClusterPrincipalAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterPrincipalAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterPrincipalAssignmentResultOutput{})
}
