// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties for an Azure Dev Spaces Controller.
//
// Uses Azure REST API version 2019-04-01.
func LookupController(ctx *pulumi.Context, args *LookupControllerArgs, opts ...pulumi.InvokeOption) (*LookupControllerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupControllerResult
	err := ctx.Invoke("azure-native:devspaces:getController", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControllerArgs struct {
	// Name of the resource.
	Name string `pulumi:"name"`
	// Resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupControllerResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// DNS name for accessing DataPlane services
	DataPlaneFqdn string `pulumi:"dataPlaneFqdn"`
	// DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
	HostSuffix string `pulumi:"hostSuffix"`
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// Region where the Azure resource is located.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Provisioning state of the Azure Dev Spaces Controller.
	ProvisioningState string `pulumi:"provisioningState"`
	// Model representing SKU for Azure Dev Spaces Controller.
	Sku SkuResponse `pulumi:"sku"`
	// Tags for the Azure resource.
	Tags map[string]string `pulumi:"tags"`
	// DNS of the target container host's API server
	TargetContainerHostApiServerFqdn string `pulumi:"targetContainerHostApiServerFqdn"`
	// Resource ID of the target container host
	TargetContainerHostResourceId string `pulumi:"targetContainerHostResourceId"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupControllerOutput(ctx *pulumi.Context, args LookupControllerOutputArgs, opts ...pulumi.InvokeOption) LookupControllerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupControllerResultOutput, error) {
			args := v.(LookupControllerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devspaces:getController", args, LookupControllerResultOutput{}, options).(LookupControllerResultOutput), nil
		}).(LookupControllerResultOutput)
}

type LookupControllerOutputArgs struct {
	// Name of the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// Resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupControllerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerArgs)(nil)).Elem()
}

type LookupControllerResultOutput struct{ *pulumi.OutputState }

func (LookupControllerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControllerResult)(nil)).Elem()
}

func (o LookupControllerResultOutput) ToLookupControllerResultOutput() LookupControllerResultOutput {
	return o
}

func (o LookupControllerResultOutput) ToLookupControllerResultOutputWithContext(ctx context.Context) LookupControllerResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupControllerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// DNS name for accessing DataPlane services
func (o LookupControllerResultOutput) DataPlaneFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.DataPlaneFqdn }).(pulumi.StringOutput)
}

// DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
func (o LookupControllerResultOutput) HostSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.HostSuffix }).(pulumi.StringOutput)
}

// Fully qualified resource Id for the resource.
func (o LookupControllerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Region where the Azure resource is located.
func (o LookupControllerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupControllerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the Azure Dev Spaces Controller.
func (o LookupControllerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Model representing SKU for Azure Dev Spaces Controller.
func (o LookupControllerResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupControllerResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

// Tags for the Azure resource.
func (o LookupControllerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupControllerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// DNS of the target container host's API server
func (o LookupControllerResultOutput) TargetContainerHostApiServerFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.TargetContainerHostApiServerFqdn }).(pulumi.StringOutput)
}

// Resource ID of the target container host
func (o LookupControllerResultOutput) TargetContainerHostResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.TargetContainerHostResourceId }).(pulumi.StringOutput)
}

// The type of the resource.
func (o LookupControllerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControllerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControllerResultOutput{})
}
