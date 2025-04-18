// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a devcenter plan member.
//
// Uses Azure REST API version 2024-10-01-preview.
//
// Other available API versions: 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPlanMember(ctx *pulumi.Context, args *LookupPlanMemberArgs, opts ...pulumi.InvokeOption) (*LookupPlanMemberResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPlanMemberResult
	err := ctx.Invoke("azure-native:devcenter:getPlanMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlanMemberArgs struct {
	// The name of a devcenter plan member.
	MemberName string `pulumi:"memberName"`
	// The name of the devcenter plan.
	PlanName string `pulumi:"planName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a devcenter plan member resource.
type LookupPlanMemberResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The unique id of the member.
	MemberId *string `pulumi:"memberId"`
	// The type of the member (user, group)
	MemberType *string `pulumi:"memberType"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The sync status of the member.
	SyncStatus PlanMemberSyncStatusResponse `pulumi:"syncStatus"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The tier of the member.
	Tier *string `pulumi:"tier"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupPlanMemberOutput(ctx *pulumi.Context, args LookupPlanMemberOutputArgs, opts ...pulumi.InvokeOption) LookupPlanMemberResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPlanMemberResultOutput, error) {
			args := v.(LookupPlanMemberArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter:getPlanMember", args, LookupPlanMemberResultOutput{}, options).(LookupPlanMemberResultOutput), nil
		}).(LookupPlanMemberResultOutput)
}

type LookupPlanMemberOutputArgs struct {
	// The name of a devcenter plan member.
	MemberName pulumi.StringInput `pulumi:"memberName"`
	// The name of the devcenter plan.
	PlanName pulumi.StringInput `pulumi:"planName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPlanMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlanMemberArgs)(nil)).Elem()
}

// Represents a devcenter plan member resource.
type LookupPlanMemberResultOutput struct{ *pulumi.OutputState }

func (LookupPlanMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlanMemberResult)(nil)).Elem()
}

func (o LookupPlanMemberResultOutput) ToLookupPlanMemberResultOutput() LookupPlanMemberResultOutput {
	return o
}

func (o LookupPlanMemberResultOutput) ToLookupPlanMemberResultOutputWithContext(ctx context.Context) LookupPlanMemberResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupPlanMemberResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupPlanMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// The unique id of the member.
func (o LookupPlanMemberResultOutput) MemberId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) *string { return v.MemberId }).(pulumi.StringPtrOutput)
}

// The type of the member (user, group)
func (o LookupPlanMemberResultOutput) MemberType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) *string { return v.MemberType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupPlanMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupPlanMemberResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The sync status of the member.
func (o LookupPlanMemberResultOutput) SyncStatus() PlanMemberSyncStatusResponseOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) PlanMemberSyncStatusResponse { return v.SyncStatus }).(PlanMemberSyncStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPlanMemberResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPlanMemberResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The tier of the member.
func (o LookupPlanMemberResultOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPlanMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlanMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlanMemberResultOutput{})
}
