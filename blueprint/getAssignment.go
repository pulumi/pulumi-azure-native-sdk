// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package blueprint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a blueprint assignment.
//
// Uses Azure REST API version 2018-11-01-preview.
func LookupAssignment(ctx *pulumi.Context, args *LookupAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAssignmentResult
	err := ctx.Invoke("azure-native:blueprint:getAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssignmentArgs struct {
	// Name of the blueprint assignment.
	AssignmentName string `pulumi:"assignmentName"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope string `pulumi:"resourceScope"`
}

// Represents a blueprint assignment.
type LookupAssignmentResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// ID of the published version of a blueprint definition.
	BlueprintId *string `pulumi:"blueprintId"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Managed identity for this blueprint assignment.
	Identity ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of this blueprint assignment.
	Location string `pulumi:"location"`
	// Defines how resources deployed by a blueprint assignment are locked.
	Locks *AssignmentLockSettingsResponse `pulumi:"locks"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Blueprint assignment parameter values.
	Parameters map[string]ParameterValueResponse `pulumi:"parameters"`
	// State of the blueprint assignment.
	ProvisioningState string `pulumi:"provisioningState"`
	// Names and locations of resource group placeholders.
	ResourceGroups map[string]ResourceGroupValueResponse `pulumi:"resourceGroups"`
	// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
	Scope *string `pulumi:"scope"`
	// Status of blueprint assignment. This field is readonly.
	Status AssignmentStatusResponse `pulumi:"status"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

func LookupAssignmentOutput(ctx *pulumi.Context, args LookupAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupAssignmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAssignmentResultOutput, error) {
			args := v.(LookupAssignmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:blueprint:getAssignment", args, LookupAssignmentResultOutput{}, options).(LookupAssignmentResultOutput), nil
		}).(LookupAssignmentResultOutput)
}

type LookupAssignmentOutputArgs struct {
	// Name of the blueprint assignment.
	AssignmentName pulumi.StringInput `pulumi:"assignmentName"`
	// The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
	ResourceScope pulumi.StringInput `pulumi:"resourceScope"`
}

func (LookupAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssignmentArgs)(nil)).Elem()
}

// Represents a blueprint assignment.
type LookupAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssignmentResult)(nil)).Elem()
}

func (o LookupAssignmentResultOutput) ToLookupAssignmentResultOutput() LookupAssignmentResultOutput {
	return o
}

func (o LookupAssignmentResultOutput) ToLookupAssignmentResultOutputWithContext(ctx context.Context) LookupAssignmentResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAssignmentResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// ID of the published version of a blueprint definition.
func (o LookupAssignmentResultOutput) BlueprintId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.BlueprintId }).(pulumi.StringPtrOutput)
}

// Multi-line explain this resource.
func (o LookupAssignmentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// One-liner string explain this resource.
func (o LookupAssignmentResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed identity for this blueprint assignment.
func (o LookupAssignmentResultOutput) Identity() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v LookupAssignmentResult) ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponseOutput)
}

// The location of this blueprint assignment.
func (o LookupAssignmentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Location }).(pulumi.StringOutput)
}

// Defines how resources deployed by a blueprint assignment are locked.
func (o LookupAssignmentResultOutput) Locks() AssignmentLockSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *AssignmentLockSettingsResponse { return v.Locks }).(AssignmentLockSettingsResponsePtrOutput)
}

// Name of this resource.
func (o LookupAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Blueprint assignment parameter values.
func (o LookupAssignmentResultOutput) Parameters() ParameterValueResponseMapOutput {
	return o.ApplyT(func(v LookupAssignmentResult) map[string]ParameterValueResponse { return v.Parameters }).(ParameterValueResponseMapOutput)
}

// State of the blueprint assignment.
func (o LookupAssignmentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Names and locations of resource group placeholders.
func (o LookupAssignmentResultOutput) ResourceGroups() ResourceGroupValueResponseMapOutput {
	return o.ApplyT(func(v LookupAssignmentResult) map[string]ResourceGroupValueResponse { return v.ResourceGroups }).(ResourceGroupValueResponseMapOutput)
}

// The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
func (o LookupAssignmentResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssignmentResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

// Status of blueprint assignment. This field is readonly.
func (o LookupAssignmentResultOutput) Status() AssignmentStatusResponseOutput {
	return o.ApplyT(func(v LookupAssignmentResult) AssignmentStatusResponse { return v.Status }).(AssignmentStatusResponseOutput)
}

// Type of this resource.
func (o LookupAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssignmentResultOutput{})
}
