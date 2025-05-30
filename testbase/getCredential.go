// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a test base credential Resource
//
// Uses Azure REST API version 2023-11-01-preview.
func LookupCredential(ctx *pulumi.Context, args *LookupCredentialArgs, opts ...pulumi.InvokeOption) (*LookupCredentialResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCredentialResult
	err := ctx.Invoke("azure-native:testbase:getCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialArgs struct {
	// The credential resource name.
	CredentialName string `pulumi:"credentialName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

// The test base credential resource.
type LookupCredentialResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Credential type.
	CredentialType string `pulumi:"credentialType"`
	// Credential display name.
	DisplayName string `pulumi:"displayName"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCredentialOutput(ctx *pulumi.Context, args LookupCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupCredentialResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCredentialResultOutput, error) {
			args := v.(LookupCredentialArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:testbase:getCredential", args, LookupCredentialResultOutput{}, options).(LookupCredentialResultOutput), nil
		}).(LookupCredentialResultOutput)
}

type LookupCredentialOutputArgs struct {
	// The credential resource name.
	CredentialName pulumi.StringInput `pulumi:"credentialName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialArgs)(nil)).Elem()
}

// The test base credential resource.
type LookupCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialResult)(nil)).Elem()
}

func (o LookupCredentialResultOutput) ToLookupCredentialResultOutput() LookupCredentialResultOutput {
	return o
}

func (o LookupCredentialResultOutput) ToLookupCredentialResultOutputWithContext(ctx context.Context) LookupCredentialResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupCredentialResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Credential type.
func (o LookupCredentialResultOutput) CredentialType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.CredentialType }).(pulumi.StringOutput)
}

// Credential display name.
func (o LookupCredentialResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupCredentialResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCredentialResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCredentialResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCredentialResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCredentialResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCredentialResultOutput{})
}
