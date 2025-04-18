// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Ec2Instance
//
// Uses Azure REST API version 2024-12-01.
func LookupEc2Instance(ctx *pulumi.Context, args *LookupEc2InstanceArgs, opts ...pulumi.InvokeOption) (*LookupEc2InstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEc2InstanceResult
	err := ctx.Invoke("azure-native:awsconnector:getEc2Instance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEc2InstanceArgs struct {
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// A Microsoft.AwsConnector resource
type LookupEc2InstanceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties Ec2InstancePropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupEc2InstanceResult
func (val *LookupEc2InstanceResult) Defaults() *LookupEc2InstanceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupEc2InstanceOutput(ctx *pulumi.Context, args LookupEc2InstanceOutputArgs, opts ...pulumi.InvokeOption) LookupEc2InstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEc2InstanceResultOutput, error) {
			args := v.(LookupEc2InstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:awsconnector:getEc2Instance", args, LookupEc2InstanceResultOutput{}, options).(LookupEc2InstanceResultOutput), nil
		}).(LookupEc2InstanceResultOutput)
}

type LookupEc2InstanceOutputArgs struct {
	// The fully qualified Azure Resource manager identifier of the resource.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupEc2InstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2InstanceArgs)(nil)).Elem()
}

// A Microsoft.AwsConnector resource
type LookupEc2InstanceResultOutput struct{ *pulumi.OutputState }

func (LookupEc2InstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEc2InstanceResult)(nil)).Elem()
}

func (o LookupEc2InstanceResultOutput) ToLookupEc2InstanceResultOutput() LookupEc2InstanceResultOutput {
	return o
}

func (o LookupEc2InstanceResultOutput) ToLookupEc2InstanceResultOutputWithContext(ctx context.Context) LookupEc2InstanceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupEc2InstanceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2InstanceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupEc2InstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2InstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupEc2InstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2InstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupEc2InstanceResultOutput) Properties() Ec2InstancePropertiesResponseOutput {
	return o.ApplyT(func(v LookupEc2InstanceResult) Ec2InstancePropertiesResponse { return v.Properties }).(Ec2InstancePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupEc2InstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEc2InstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupEc2InstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEc2InstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEc2InstanceResultOutput{})
}
