// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the default rollout details.
//
// Uses Azure REST API version 2021-09-01-preview.
func LookupDefaultRollout(ctx *pulumi.Context, args *LookupDefaultRolloutArgs, opts ...pulumi.InvokeOption) (*LookupDefaultRolloutResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultRolloutResult
	err := ctx.Invoke("azure-native:providerhub:getDefaultRollout", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDefaultRolloutArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The rollout name.
	RolloutName string `pulumi:"rolloutName"`
}

// Default rollout definition.
type LookupDefaultRolloutResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Properties of the rollout.
	Properties DefaultRolloutResponseProperties `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupDefaultRolloutOutput(ctx *pulumi.Context, args LookupDefaultRolloutOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultRolloutResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDefaultRolloutResultOutput, error) {
			args := v.(LookupDefaultRolloutArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:providerhub:getDefaultRollout", args, LookupDefaultRolloutResultOutput{}, options).(LookupDefaultRolloutResultOutput), nil
		}).(LookupDefaultRolloutResultOutput)
}

type LookupDefaultRolloutOutputArgs struct {
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	// The rollout name.
	RolloutName pulumi.StringInput `pulumi:"rolloutName"`
}

func (LookupDefaultRolloutOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultRolloutArgs)(nil)).Elem()
}

// Default rollout definition.
type LookupDefaultRolloutResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultRolloutResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultRolloutResult)(nil)).Elem()
}

func (o LookupDefaultRolloutResultOutput) ToLookupDefaultRolloutResultOutput() LookupDefaultRolloutResultOutput {
	return o
}

func (o LookupDefaultRolloutResultOutput) ToLookupDefaultRolloutResultOutputWithContext(ctx context.Context) LookupDefaultRolloutResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDefaultRolloutResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupDefaultRolloutResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDefaultRolloutResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the rollout.
func (o LookupDefaultRolloutResultOutput) Properties() DefaultRolloutResponsePropertiesOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) DefaultRolloutResponseProperties { return v.Properties }).(DefaultRolloutResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDefaultRolloutResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDefaultRolloutResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultRolloutResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultRolloutResultOutput{})
}
