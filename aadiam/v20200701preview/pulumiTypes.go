// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureADMetricsPropertiesFormatResponse struct {
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
}

type AzureADMetricsPropertiesFormatResponseOutput struct{ *pulumi.OutputState }

func (AzureADMetricsPropertiesFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureADMetricsPropertiesFormatResponse)(nil)).Elem()
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ToAzureADMetricsPropertiesFormatResponseOutput() AzureADMetricsPropertiesFormatResponseOutput {
	return o
}

func (o AzureADMetricsPropertiesFormatResponseOutput) ToAzureADMetricsPropertiesFormatResponseOutputWithContext(ctx context.Context) AzureADMetricsPropertiesFormatResponseOutput {
	return o
}

// The provisioning state of the resource.
func (o AzureADMetricsPropertiesFormatResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v AzureADMetricsPropertiesFormatResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureADMetricsPropertiesFormatResponseOutput{})
}