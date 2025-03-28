// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20241101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Experiment resource.
func LookupExperiment(ctx *pulumi.Context, args *LookupExperimentArgs, opts ...pulumi.InvokeOption) (*LookupExperimentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupExperimentResult
	err := ctx.Invoke("azure-native:chaos/v20241101preview:getExperiment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExperimentArgs struct {
	// String that represents a Experiment resource name.
	ExperimentName string `pulumi:"experimentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Model that represents a Experiment resource.
type LookupExperimentResult struct {
	// Optional customer-managed Storage account where Experiment schema will be stored.
	CustomerDataStorage *CustomerDataStoragePropertiesResponse `pulumi:"customerDataStorage"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Most recent provisioning state for the given experiment resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// List of selectors.
	Selectors []interface{} `pulumi:"selectors"`
	// List of steps.
	Steps []ChaosExperimentStepResponse `pulumi:"steps"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupExperimentOutput(ctx *pulumi.Context, args LookupExperimentOutputArgs, opts ...pulumi.InvokeOption) LookupExperimentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupExperimentResultOutput, error) {
			args := v.(LookupExperimentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:chaos/v20241101preview:getExperiment", args, LookupExperimentResultOutput{}, options).(LookupExperimentResultOutput), nil
		}).(LookupExperimentResultOutput)
}

type LookupExperimentOutputArgs struct {
	// String that represents a Experiment resource name.
	ExperimentName pulumi.StringInput `pulumi:"experimentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExperimentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentArgs)(nil)).Elem()
}

// Model that represents a Experiment resource.
type LookupExperimentResultOutput struct{ *pulumi.OutputState }

func (LookupExperimentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExperimentResult)(nil)).Elem()
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutput() LookupExperimentResultOutput {
	return o
}

func (o LookupExperimentResultOutput) ToLookupExperimentResultOutputWithContext(ctx context.Context) LookupExperimentResultOutput {
	return o
}

// Optional customer-managed Storage account where Experiment schema will be stored.
func (o LookupExperimentResultOutput) CustomerDataStorage() CustomerDataStoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *CustomerDataStoragePropertiesResponse { return v.CustomerDataStorage }).(CustomerDataStoragePropertiesResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupExperimentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o LookupExperimentResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupExperimentResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupExperimentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupExperimentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Most recent provisioning state for the given experiment resource.
func (o LookupExperimentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// List of selectors.
func (o LookupExperimentResultOutput) Selectors() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []interface{} { return v.Selectors }).(pulumi.ArrayOutput)
}

// List of steps.
func (o LookupExperimentResultOutput) Steps() ChaosExperimentStepResponseArrayOutput {
	return o.ApplyT(func(v LookupExperimentResult) []ChaosExperimentStepResponse { return v.Steps }).(ChaosExperimentStepResponseArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupExperimentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupExperimentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupExperimentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExperimentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupExperimentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExperimentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExperimentResultOutput{})
}
