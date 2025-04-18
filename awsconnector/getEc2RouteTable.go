// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Ec2RouteTable
//
// Uses Azure REST API version 2024-12-01.
func LookupEc2RouteTable(ctx *pulumi.Context, args *LookupEc2RouteTableArgs, opts ...pulumi.InvokeOption) (*LookupEc2RouteTableResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEc2RouteTableResult
	err := ctx.Invoke("azure-native:awsconnector:getEc2RouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEc2RouteTableArgs struct {
	// Name of Ec2RouteTable
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupEc2RouteTableResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties Ec2RouteTablePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEc2RouteTableOutput(ctx *pulumi.Context, args LookupEc2RouteTableOutputArgs, opts ...pulumi.InvokeOption) LookupEc2RouteTableResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEc2RouteTableResultOutput, error) {
			args := v.(LookupEc2RouteTableArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getEc2RouteTable", args, LookupEc2RouteTableResultOutput{}, options).(LookupEc2RouteTableResultOutput), nil
		}).(LookupEc2RouteTableResultOutput)
}

type LookupEc2RouteTableOutputArgs struct {
	// Name of Ec2RouteTable
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEc2RouteTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2RouteTableArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupEc2RouteTableResultOutput struct{ *pulumi.OutputState }

func (LookupEc2RouteTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2RouteTableResult)(nil)).Elem()
}

func (o LookupEc2RouteTableResultOutput) ToLookupEc2RouteTableResultOutput() LookupEc2RouteTableResultOutput {
	return o
}

func (o LookupEc2RouteTableResultOutput) ToLookupEc2RouteTableResultOutputWithContext(ctx context.Context) LookupEc2RouteTableResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupEc2RouteTableResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupEc2RouteTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupEc2RouteTableResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEc2RouteTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupEc2RouteTableResultOutput) Properties() Ec2RouteTablePropertiesResponseOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) Ec2RouteTablePropertiesResponse { return v.Properties }).(Ec2RouteTablePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEc2RouteTableResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupEc2RouteTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEc2RouteTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2RouteTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEc2RouteTableResultOutput{})
}
