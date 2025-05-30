// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Ec2VPCPeeringConnection
//
// Uses Azure REST API version 2024-12-01.
func LookupEc2VpcPeeringConnection(ctx *pulumi.Context, args *LookupEc2VpcPeeringConnectionArgs, opts ...pulumi.InvokeOption) (*LookupEc2VpcPeeringConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEc2VpcPeeringConnectionResult
	err := ctx.Invoke("azure-native:awsconnector:getEc2VpcPeeringConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEc2VpcPeeringConnectionArgs struct {
	// Name of Ec2VPCPeeringConnection
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Microsoft.AwsConnector resource
type LookupEc2VpcPeeringConnectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties Ec2VPCPeeringConnectionPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupEc2VpcPeeringConnectionOutput(ctx *pulumi.Context, args LookupEc2VpcPeeringConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupEc2VpcPeeringConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEc2VpcPeeringConnectionResultOutput, error) {
			args := v.(LookupEc2VpcPeeringConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getEc2VpcPeeringConnection", args, LookupEc2VpcPeeringConnectionResultOutput{}, options).(LookupEc2VpcPeeringConnectionResultOutput), nil
		}).(LookupEc2VpcPeeringConnectionResultOutput)
}

type LookupEc2VpcPeeringConnectionOutputArgs struct {
	// Name of Ec2VPCPeeringConnection
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEc2VpcPeeringConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2VpcPeeringConnectionArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupEc2VpcPeeringConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupEc2VpcPeeringConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2VpcPeeringConnectionResult)(nil)).Elem()
}

func (o LookupEc2VpcPeeringConnectionResultOutput) ToLookupEc2VpcPeeringConnectionResultOutput() LookupEc2VpcPeeringConnectionResultOutput {
	return o
}

func (o LookupEc2VpcPeeringConnectionResultOutput) ToLookupEc2VpcPeeringConnectionResultOutputWithContext(ctx context.Context) LookupEc2VpcPeeringConnectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupEc2VpcPeeringConnectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupEc2VpcPeeringConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupEc2VpcPeeringConnectionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEc2VpcPeeringConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupEc2VpcPeeringConnectionResultOutput) Properties() Ec2VPCPeeringConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) Ec2VPCPeeringConnectionPropertiesResponse {
		return v.Properties
	}).(Ec2VPCPeeringConnectionPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEc2VpcPeeringConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupEc2VpcPeeringConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEc2VpcPeeringConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2VpcPeeringConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEc2VpcPeeringConnectionResultOutput{})
}
