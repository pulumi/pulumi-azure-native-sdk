// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a BusinessCase
//
// Uses Azure REST API version 2024-01-01-preview.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupBusinessCaseOperation(ctx *pulumi.Context, args *LookupBusinessCaseOperationArgs, opts ...pulumi.InvokeOption) (*LookupBusinessCaseOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBusinessCaseOperationResult
	err := ctx.Invoke("azure-native:migrate:getBusinessCaseOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBusinessCaseOperationArgs struct {
	// Business case ARM name
	BusinessCaseName string `pulumi:"businessCaseName"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Business case resource.
type LookupBusinessCaseOperationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState string `pulumi:"provisioningState"`
	// Gets the state of business case reports.
	ReportStatusDetails []ReportDetailsResponse `pulumi:"reportStatusDetails"`
	// Business case settings.
	Settings *SettingsResponse `pulumi:"settings"`
	// Business case state.
	State string `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupBusinessCaseOperationResult
func (val *LookupBusinessCaseOperationResult) Defaults() *LookupBusinessCaseOperationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Settings = tmp.Settings.Defaults()

	return &tmp
}
func LookupBusinessCaseOperationOutput(ctx *pulumi.Context, args LookupBusinessCaseOperationOutputArgs, opts ...pulumi.InvokeOption) LookupBusinessCaseOperationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBusinessCaseOperationResultOutput, error) {
			args := v.(LookupBusinessCaseOperationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:migrate:getBusinessCaseOperation", args, LookupBusinessCaseOperationResultOutput{}, options).(LookupBusinessCaseOperationResultOutput), nil
		}).(LookupBusinessCaseOperationResultOutput)
}

type LookupBusinessCaseOperationOutputArgs struct {
	// Business case ARM name
	BusinessCaseName pulumi.StringInput `pulumi:"businessCaseName"`
	// Assessment Project Name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBusinessCaseOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBusinessCaseOperationArgs)(nil)).Elem()
}

// Business case resource.
type LookupBusinessCaseOperationResultOutput struct{ *pulumi.OutputState }

func (LookupBusinessCaseOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBusinessCaseOperationResult)(nil)).Elem()
}

func (o LookupBusinessCaseOperationResultOutput) ToLookupBusinessCaseOperationResultOutput() LookupBusinessCaseOperationResultOutput {
	return o
}

func (o LookupBusinessCaseOperationResultOutput) ToLookupBusinessCaseOperationResultOutputWithContext(ctx context.Context) LookupBusinessCaseOperationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupBusinessCaseOperationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupBusinessCaseOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBusinessCaseOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o LookupBusinessCaseOperationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the state of business case reports.
func (o LookupBusinessCaseOperationResultOutput) ReportStatusDetails() ReportDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) []ReportDetailsResponse { return v.ReportStatusDetails }).(ReportDetailsResponseArrayOutput)
}

// Business case settings.
func (o LookupBusinessCaseOperationResultOutput) Settings() SettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) *SettingsResponse { return v.Settings }).(SettingsResponsePtrOutput)
}

// Business case state.
func (o LookupBusinessCaseOperationResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) string { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBusinessCaseOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBusinessCaseOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBusinessCaseOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBusinessCaseOperationResultOutput{})
}
