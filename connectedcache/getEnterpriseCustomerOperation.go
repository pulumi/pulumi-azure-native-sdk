// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectedcache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the properties of a Enterprise customer
//
// Uses Azure REST API version 2023-05-01-preview.
func LookupEnterpriseCustomerOperation(ctx *pulumi.Context, args *LookupEnterpriseCustomerOperationArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseCustomerOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnterpriseCustomerOperationResult
	err := ctx.Invoke("azure-native:connectedcache:getEnterpriseCustomerOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseCustomerOperationArgs struct {
	// Name of the Customer resource
	CustomerResourceName string `pulumi:"customerResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// ConnectedCache Resource. Represents the high level Nodes needed to provision CacheNode and customer resources used in private preview
type LookupEnterpriseCustomerOperationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties CacheNodeOldResponseResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEnterpriseCustomerOperationOutput(ctx *pulumi.Context, args LookupEnterpriseCustomerOperationOutputArgs, opts ...pulumi.InvokeOption) LookupEnterpriseCustomerOperationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEnterpriseCustomerOperationResultOutput, error) {
			args := v.(LookupEnterpriseCustomerOperationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:connectedcache:getEnterpriseCustomerOperation", args, LookupEnterpriseCustomerOperationResultOutput{}, options).(LookupEnterpriseCustomerOperationResultOutput), nil
		}).(LookupEnterpriseCustomerOperationResultOutput)
}

type LookupEnterpriseCustomerOperationOutputArgs struct {
	// Name of the Customer resource
	CustomerResourceName pulumi.StringInput `pulumi:"customerResourceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEnterpriseCustomerOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseCustomerOperationArgs)(nil)).Elem()
}

// ConnectedCache Resource. Represents the high level Nodes needed to provision CacheNode and customer resources used in private preview
type LookupEnterpriseCustomerOperationResultOutput struct{ *pulumi.OutputState }

func (LookupEnterpriseCustomerOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseCustomerOperationResult)(nil)).Elem()
}

func (o LookupEnterpriseCustomerOperationResultOutput) ToLookupEnterpriseCustomerOperationResultOutput() LookupEnterpriseCustomerOperationResultOutput {
	return o
}

func (o LookupEnterpriseCustomerOperationResultOutput) ToLookupEnterpriseCustomerOperationResultOutputWithContext(ctx context.Context) LookupEnterpriseCustomerOperationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupEnterpriseCustomerOperationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupEnterpriseCustomerOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupEnterpriseCustomerOperationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEnterpriseCustomerOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupEnterpriseCustomerOperationResultOutput) Properties() CacheNodeOldResponseResponseOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) CacheNodeOldResponseResponse { return v.Properties }).(CacheNodeOldResponseResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEnterpriseCustomerOperationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupEnterpriseCustomerOperationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEnterpriseCustomerOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseCustomerOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterpriseCustomerOperationResultOutput{})
}
