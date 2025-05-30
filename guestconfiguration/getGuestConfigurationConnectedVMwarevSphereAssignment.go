// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package guestconfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a guest configuration assignment
//
// Uses Azure REST API version 2024-04-05.
//
// Other available API versions: 2022-01-25. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native guestconfiguration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupGuestConfigurationConnectedVMwarevSphereAssignment(ctx *pulumi.Context, args *LookupGuestConfigurationConnectedVMwarevSphereAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult
	err := ctx.Invoke("azure-native:guestconfiguration:getGuestConfigurationConnectedVMwarevSphereAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupGuestConfigurationConnectedVMwarevSphereAssignmentArgs struct {
	// The guest configuration assignment name.
	GuestConfigurationAssignmentName string `pulumi:"guestConfigurationAssignmentName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
}

// Guest configuration assignment is an association between a machine and guest configuration.
type LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// ARM resource id of the guest configuration assignment.
	Id string `pulumi:"id"`
	// Region where the VM is located.
	Location *string `pulumi:"location"`
	// Name of the guest configuration assignment.
	Name *string `pulumi:"name"`
	// Properties of the Guest configuration assignment.
	Properties GuestConfigurationAssignmentPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult
func (val *LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) Defaults() *LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupGuestConfigurationConnectedVMwarevSphereAssignmentOutput(ctx *pulumi.Context, args LookupGuestConfigurationConnectedVMwarevSphereAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput, error) {
			args := v.(LookupGuestConfigurationConnectedVMwarevSphereAssignmentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:guestconfiguration:getGuestConfigurationConnectedVMwarevSphereAssignment", args, LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput{}, options).(LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput), nil
		}).(LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput)
}

type LookupGuestConfigurationConnectedVMwarevSphereAssignmentOutputArgs struct {
	// The guest configuration assignment name.
	GuestConfigurationAssignmentName pulumi.StringInput `pulumi:"guestConfigurationAssignmentName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the virtual machine.
	VmName pulumi.StringInput `pulumi:"vmName"`
}

func (LookupGuestConfigurationConnectedVMwarevSphereAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationConnectedVMwarevSphereAssignmentArgs)(nil)).Elem()
}

// Guest configuration assignment is an association between a machine and guest configuration.
type LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult)(nil)).Elem()
}

func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) ToLookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput() LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput {
	return o
}

func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) ToLookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutputWithContext(ctx context.Context) LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) string {
		return v.AzureApiVersion
	}).(pulumi.StringOutput)
}

// ARM resource id of the guest configuration assignment.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Region where the VM is located.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the guest configuration assignment.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Properties of the Guest configuration assignment.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) Properties() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) GuestConfigurationAssignmentPropertiesResponse {
		return v.Properties
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) SystemDataResponse {
		return v.SystemData
	}).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestConfigurationConnectedVMwarevSphereAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestConfigurationConnectedVMwarevSphereAssignmentResultOutput{})
}
