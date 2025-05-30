// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package peering

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an existing registered prefix with the specified name under the given subscription, resource group and peering.
//
// Uses Azure REST API version 2022-10-01.
func LookupRegisteredPrefix(ctx *pulumi.Context, args *LookupRegisteredPrefixArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredPrefixResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRegisteredPrefixResult
	err := ctx.Invoke("azure-native:peering:getRegisteredPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredPrefixArgs struct {
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the registered prefix.
	RegisteredPrefixName string `pulumi:"registeredPrefixName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The customer's prefix that is registered by the peering service provider.
type LookupRegisteredPrefixResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The error message associated with the validation state, if any.
	ErrorMessage string `pulumi:"errorMessage"`
	// The ID of the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey string `pulumi:"peeringServicePrefixKey"`
	// The customer's prefix from which traffic originates.
	Prefix *string `pulumi:"prefix"`
	// The prefix validation state.
	PrefixValidationState string `pulumi:"prefixValidationState"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupRegisteredPrefixOutput(ctx *pulumi.Context, args LookupRegisteredPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupRegisteredPrefixResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRegisteredPrefixResultOutput, error) {
			args := v.(LookupRegisteredPrefixArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:peering:getRegisteredPrefix", args, LookupRegisteredPrefixResultOutput{}, options).(LookupRegisteredPrefixResultOutput), nil
		}).(LookupRegisteredPrefixResultOutput)
}

type LookupRegisteredPrefixOutputArgs struct {
	// The name of the peering.
	PeeringName pulumi.StringInput `pulumi:"peeringName"`
	// The name of the registered prefix.
	RegisteredPrefixName pulumi.StringInput `pulumi:"registeredPrefixName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRegisteredPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredPrefixArgs)(nil)).Elem()
}

// The customer's prefix that is registered by the peering service provider.
type LookupRegisteredPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupRegisteredPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegisteredPrefixResult)(nil)).Elem()
}

func (o LookupRegisteredPrefixResultOutput) ToLookupRegisteredPrefixResultOutput() LookupRegisteredPrefixResultOutput {
	return o
}

func (o LookupRegisteredPrefixResultOutput) ToLookupRegisteredPrefixResultOutputWithContext(ctx context.Context) LookupRegisteredPrefixResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupRegisteredPrefixResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The error message associated with the validation state, if any.
func (o LookupRegisteredPrefixResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// The ID of the resource.
func (o LookupRegisteredPrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupRegisteredPrefixResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.Name }).(pulumi.StringOutput)
}

// The peering service prefix key that is to be shared with the customer.
func (o LookupRegisteredPrefixResultOutput) PeeringServicePrefixKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.PeeringServicePrefixKey }).(pulumi.StringOutput)
}

// The customer's prefix from which traffic originates.
func (o LookupRegisteredPrefixResultOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) *string { return v.Prefix }).(pulumi.StringPtrOutput)
}

// The prefix validation state.
func (o LookupRegisteredPrefixResultOutput) PrefixValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.PrefixValidationState }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupRegisteredPrefixResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource.
func (o LookupRegisteredPrefixResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegisteredPrefixResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegisteredPrefixResultOutput{})
}
