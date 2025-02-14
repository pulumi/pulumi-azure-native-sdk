// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the supported kubernetes versions for the specified custom location
func LookupKubernetesVersions(ctx *pulumi.Context, args *LookupKubernetesVersionsArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesVersionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesVersionsResult
	err := ctx.Invoke("azure-native:hybridcontainerservice/v20240101:getKubernetesVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubernetesVersionsArgs struct {
	// The fully qualified Azure Resource Manager identifier of the custom location resource.
	CustomLocationResourceUri string `pulumi:"customLocationResourceUri"`
}

// The supported kubernetes versions.
type LookupKubernetesVersionsResult struct {
	// Extended location pointing to the underlying infrastructure
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name       string                                     `pulumi:"name"`
	Properties KubernetesVersionProfileResponseProperties `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupKubernetesVersionsOutput(ctx *pulumi.Context, args LookupKubernetesVersionsOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesVersionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKubernetesVersionsResultOutput, error) {
			args := v.(LookupKubernetesVersionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridcontainerservice/v20240101:getKubernetesVersions", args, LookupKubernetesVersionsResultOutput{}, options).(LookupKubernetesVersionsResultOutput), nil
		}).(LookupKubernetesVersionsResultOutput)
}

type LookupKubernetesVersionsOutputArgs struct {
	// The fully qualified Azure Resource Manager identifier of the custom location resource.
	CustomLocationResourceUri pulumi.StringInput `pulumi:"customLocationResourceUri"`
}

func (LookupKubernetesVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesVersionsArgs)(nil)).Elem()
}

// The supported kubernetes versions.
type LookupKubernetesVersionsResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesVersionsResult)(nil)).Elem()
}

func (o LookupKubernetesVersionsResultOutput) ToLookupKubernetesVersionsResultOutput() LookupKubernetesVersionsResultOutput {
	return o
}

func (o LookupKubernetesVersionsResultOutput) ToLookupKubernetesVersionsResultOutputWithContext(ctx context.Context) LookupKubernetesVersionsResultOutput {
	return o
}

// Extended location pointing to the underlying infrastructure
func (o LookupKubernetesVersionsResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupKubernetesVersionsResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupKubernetesVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupKubernetesVersionsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesVersionsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKubernetesVersionsResultOutput) Properties() KubernetesVersionProfileResponsePropertiesOutput {
	return o.ApplyT(func(v LookupKubernetesVersionsResult) KubernetesVersionProfileResponseProperties { return v.Properties }).(KubernetesVersionProfileResponsePropertiesOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupKubernetesVersionsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKubernetesVersionsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupKubernetesVersionsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesVersionsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesVersionsResultOutput{})
}
