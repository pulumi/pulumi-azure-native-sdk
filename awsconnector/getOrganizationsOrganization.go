// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a OrganizationsOrganization
//
// Uses Azure REST API version 2024-12-01.
func LookupOrganizationsOrganization(ctx *pulumi.Context, args *LookupOrganizationsOrganizationArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationsOrganizationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupOrganizationsOrganizationResult
	err := ctx.Invoke("azure-native:awsconnector:getOrganizationsOrganization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupOrganizationsOrganizationArgs struct {
	// Name of OrganizationsOrganization
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupOrganizationsOrganizationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties OrganizationsOrganizationPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupOrganizationsOrganizationResult
func (val *LookupOrganizationsOrganizationResult) Defaults() *LookupOrganizationsOrganizationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupOrganizationsOrganizationOutput(ctx *pulumi.Context, args LookupOrganizationsOrganizationOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationsOrganizationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOrganizationsOrganizationResultOutput, error) {
			args := v.(LookupOrganizationsOrganizationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getOrganizationsOrganization", args, LookupOrganizationsOrganizationResultOutput{}, options).(LookupOrganizationsOrganizationResultOutput), nil
		}).(LookupOrganizationsOrganizationResultOutput)
}

type LookupOrganizationsOrganizationOutputArgs struct {
	// Name of OrganizationsOrganization
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupOrganizationsOrganizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationsOrganizationArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupOrganizationsOrganizationResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationsOrganizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationsOrganizationResult)(nil)).Elem()
}

func (o LookupOrganizationsOrganizationResultOutput) ToLookupOrganizationsOrganizationResultOutput() LookupOrganizationsOrganizationResultOutput {
	return o
}

func (o LookupOrganizationsOrganizationResultOutput) ToLookupOrganizationsOrganizationResultOutputWithContext(ctx context.Context) LookupOrganizationsOrganizationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupOrganizationsOrganizationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupOrganizationsOrganizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupOrganizationsOrganizationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupOrganizationsOrganizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupOrganizationsOrganizationResultOutput) Properties() OrganizationsOrganizationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) OrganizationsOrganizationPropertiesResponse {
		return v.Properties
	}).(OrganizationsOrganizationPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupOrganizationsOrganizationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupOrganizationsOrganizationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupOrganizationsOrganizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationsOrganizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationsOrganizationResultOutput{})
}
