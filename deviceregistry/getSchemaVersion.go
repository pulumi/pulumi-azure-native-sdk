// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a SchemaVersion
//
// Uses Azure REST API version 2024-09-01-preview.
func LookupSchemaVersion(ctx *pulumi.Context, args *LookupSchemaVersionArgs, opts ...pulumi.InvokeOption) (*LookupSchemaVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSchemaVersionResult
	err := ctx.Invoke("azure-native:deviceregistry:getSchemaVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaVersionArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema name parameter.
	SchemaName string `pulumi:"schemaName"`
	// Schema registry name parameter.
	SchemaRegistryName string `pulumi:"schemaRegistryName"`
	// Schema version name parameter.
	SchemaVersionName string `pulumi:"schemaVersionName"`
}

// Schema version's definition.
type LookupSchemaVersionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Human-readable description of the schema.
	Description *string `pulumi:"description"`
	// Hash of the schema content.
	Hash string `pulumi:"hash"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Schema content.
	SchemaContent string `pulumi:"schemaContent"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Globally unique, immutable, non-reusable id.
	Uuid string `pulumi:"uuid"`
}

func LookupSchemaVersionOutput(ctx *pulumi.Context, args LookupSchemaVersionOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSchemaVersionResultOutput, error) {
			args := v.(LookupSchemaVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:deviceregistry:getSchemaVersion", args, LookupSchemaVersionResultOutput{}, options).(LookupSchemaVersionResultOutput), nil
		}).(LookupSchemaVersionResultOutput)
}

type LookupSchemaVersionOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Schema name parameter.
	SchemaName pulumi.StringInput `pulumi:"schemaName"`
	// Schema registry name parameter.
	SchemaRegistryName pulumi.StringInput `pulumi:"schemaRegistryName"`
	// Schema version name parameter.
	SchemaVersionName pulumi.StringInput `pulumi:"schemaVersionName"`
}

func (LookupSchemaVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaVersionArgs)(nil)).Elem()
}

// Schema version's definition.
type LookupSchemaVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaVersionResult)(nil)).Elem()
}

func (o LookupSchemaVersionResultOutput) ToLookupSchemaVersionResultOutput() LookupSchemaVersionResultOutput {
	return o
}

func (o LookupSchemaVersionResultOutput) ToLookupSchemaVersionResultOutputWithContext(ctx context.Context) LookupSchemaVersionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSchemaVersionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Human-readable description of the schema.
func (o LookupSchemaVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Hash of the schema content.
func (o LookupSchemaVersionResultOutput) Hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.Hash }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupSchemaVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSchemaVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupSchemaVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Schema content.
func (o LookupSchemaVersionResultOutput) SchemaContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.SchemaContent }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSchemaVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSchemaVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Globally unique, immutable, non-reusable id.
func (o LookupSchemaVersionResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaVersionResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaVersionResultOutput{})
}
