// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240919preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the properties of a Confidential Ledger.
func LookupLedger(ctx *pulumi.Context, args *LookupLedgerArgs, opts ...pulumi.InvokeOption) (*LookupLedgerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLedgerResult
	err := ctx.Invoke("azure-native:confidentialledger/v20240919preview:getLedger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLedgerArgs struct {
	// Name of the Confidential Ledger
	LedgerName string `pulumi:"ledgerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Confidential Ledger. Contains the properties of Confidential Ledger Resource.
type LookupLedgerResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of Confidential Ledger Resource.
	Properties LedgerPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLedgerOutput(ctx *pulumi.Context, args LookupLedgerOutputArgs, opts ...pulumi.InvokeOption) LookupLedgerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLedgerResultOutput, error) {
			args := v.(LookupLedgerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:confidentialledger/v20240919preview:getLedger", args, LookupLedgerResultOutput{}, options).(LookupLedgerResultOutput), nil
		}).(LookupLedgerResultOutput)
}

type LookupLedgerOutputArgs struct {
	// Name of the Confidential Ledger
	LedgerName pulumi.StringInput `pulumi:"ledgerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLedgerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLedgerArgs)(nil)).Elem()
}

// Confidential Ledger. Contains the properties of Confidential Ledger Resource.
type LookupLedgerResultOutput struct{ *pulumi.OutputState }

func (LookupLedgerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLedgerResult)(nil)).Elem()
}

func (o LookupLedgerResultOutput) ToLookupLedgerResultOutput() LookupLedgerResultOutput {
	return o
}

func (o LookupLedgerResultOutput) ToLookupLedgerResultOutputWithContext(ctx context.Context) LookupLedgerResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLedgerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupLedgerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLedgerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of Confidential Ledger Resource.
func (o LookupLedgerResultOutput) Properties() LedgerPropertiesResponseOutput {
	return o.ApplyT(func(v LookupLedgerResult) LedgerPropertiesResponse { return v.Properties }).(LedgerPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLedgerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLedgerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupLedgerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLedgerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLedgerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLedgerResultOutput{})
}
