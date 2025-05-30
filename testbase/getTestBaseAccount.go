// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Test Base Account.
//
// Uses Azure REST API version 2023-11-01-preview.
//
// Other available API versions: 2022-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native testbase [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupTestBaseAccount(ctx *pulumi.Context, args *LookupTestBaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupTestBaseAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTestBaseAccountResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestBaseAccountArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

// The Test Base Account resource.
type LookupTestBaseAccountResult struct {
	// The access level of the Test Base Account.
	AccessLevel string `pulumi:"accessLevel"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The identity of the testBaseAccount.
	Identity *SystemAssignedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the Test Base Account.
	Sku TestBaseAccountSKUResponse `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupTestBaseAccountOutput(ctx *pulumi.Context, args LookupTestBaseAccountOutputArgs, opts ...pulumi.InvokeOption) LookupTestBaseAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTestBaseAccountResultOutput, error) {
			args := v.(LookupTestBaseAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:testbase:getTestBaseAccount", args, LookupTestBaseAccountResultOutput{}, options).(LookupTestBaseAccountResultOutput), nil
		}).(LookupTestBaseAccountResultOutput)
}

type LookupTestBaseAccountOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupTestBaseAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestBaseAccountArgs)(nil)).Elem()
}

// The Test Base Account resource.
type LookupTestBaseAccountResultOutput struct{ *pulumi.OutputState }

func (LookupTestBaseAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestBaseAccountResult)(nil)).Elem()
}

func (o LookupTestBaseAccountResultOutput) ToLookupTestBaseAccountResultOutput() LookupTestBaseAccountResultOutput {
	return o
}

func (o LookupTestBaseAccountResultOutput) ToLookupTestBaseAccountResultOutputWithContext(ctx context.Context) LookupTestBaseAccountResultOutput {
	return o
}

// The access level of the Test Base Account.
func (o LookupTestBaseAccountResultOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.AccessLevel }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupTestBaseAccountResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupTestBaseAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the testBaseAccount.
func (o LookupTestBaseAccountResultOutput) Identity() SystemAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) *SystemAssignedServiceIdentityResponse { return v.Identity }).(SystemAssignedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupTestBaseAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupTestBaseAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupTestBaseAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the Test Base Account.
func (o LookupTestBaseAccountResultOutput) Sku() TestBaseAccountSKUResponseOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) TestBaseAccountSKUResponse { return v.Sku }).(TestBaseAccountSKUResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTestBaseAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupTestBaseAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupTestBaseAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestBaseAccountResultOutput{})
}
