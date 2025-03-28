// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240815preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a BrokerListenerResource
func LookupBrokerListener(ctx *pulumi.Context, args *LookupBrokerListenerArgs, opts ...pulumi.InvokeOption) (*LookupBrokerListenerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBrokerListenerResult
	err := ctx.Invoke("azure-native:iotoperations/v20240815preview:getBrokerListener", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBrokerListenerArgs struct {
	// Name of broker.
	BrokerName string `pulumi:"brokerName"`
	// Name of instance.
	InstanceName string `pulumi:"instanceName"`
	// Name of Instance broker listener resource
	ListenerName string `pulumi:"listenerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Instance broker resource
type LookupBrokerListenerResult struct {
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties BrokerListenerPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupBrokerListenerResult
func (val *LookupBrokerListenerResult) Defaults() *LookupBrokerListenerResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupBrokerListenerOutput(ctx *pulumi.Context, args LookupBrokerListenerOutputArgs, opts ...pulumi.InvokeOption) LookupBrokerListenerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBrokerListenerResultOutput, error) {
			args := v.(LookupBrokerListenerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:iotoperations/v20240815preview:getBrokerListener", args, LookupBrokerListenerResultOutput{}, options).(LookupBrokerListenerResultOutput), nil
		}).(LookupBrokerListenerResultOutput)
}

type LookupBrokerListenerOutputArgs struct {
	// Name of broker.
	BrokerName pulumi.StringInput `pulumi:"brokerName"`
	// Name of instance.
	InstanceName pulumi.StringInput `pulumi:"instanceName"`
	// Name of Instance broker listener resource
	ListenerName pulumi.StringInput `pulumi:"listenerName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBrokerListenerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBrokerListenerArgs)(nil)).Elem()
}

// Instance broker resource
type LookupBrokerListenerResultOutput struct{ *pulumi.OutputState }

func (LookupBrokerListenerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBrokerListenerResult)(nil)).Elem()
}

func (o LookupBrokerListenerResultOutput) ToLookupBrokerListenerResultOutput() LookupBrokerListenerResultOutput {
	return o
}

func (o LookupBrokerListenerResultOutput) ToLookupBrokerListenerResultOutputWithContext(ctx context.Context) LookupBrokerListenerResultOutput {
	return o
}

// Edge location of the resource.
func (o LookupBrokerListenerResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupBrokerListenerResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupBrokerListenerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBrokerListenerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBrokerListenerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBrokerListenerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupBrokerListenerResultOutput) Properties() BrokerListenerPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBrokerListenerResult) BrokerListenerPropertiesResponse { return v.Properties }).(BrokerListenerPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBrokerListenerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBrokerListenerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBrokerListenerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBrokerListenerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBrokerListenerResultOutput{})
}
