// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the certificate specified by its identifier.
//
// Uses Azure REST API version 2022-09-01-preview.
//
// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateResult
	err := ctx.Invoke("azure-native:apimanagement:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId string `pulumi:"certificateId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Certificate details.
type LookupCertificateResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate string `pulumi:"expirationDate"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// KeyVault location details of the certificate.
	KeyVault *KeyVaultContractPropertiesResponse `pulumi:"keyVault"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Subject attribute of the certificate.
	Subject string `pulumi:"subject"`
	// Thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCertificateResultOutput, error) {
			args := v.(LookupCertificateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement:getCertificate", args, LookupCertificateResultOutput{}, options).(LookupCertificateResultOutput), nil
		}).(LookupCertificateResultOutput)
}

type LookupCertificateOutputArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId pulumi.StringInput `pulumi:"certificateId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// Certificate details.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupCertificateResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
func (o LookupCertificateResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

// KeyVault location details of the certificate.
func (o LookupCertificateResultOutput) KeyVault() KeyVaultContractPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *KeyVaultContractPropertiesResponse { return v.KeyVault }).(KeyVaultContractPropertiesResponsePtrOutput)
}

// The name of the resource
func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Subject attribute of the certificate.
func (o LookupCertificateResultOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Subject }).(pulumi.StringOutput)
}

// Thumbprint of the certificate.
func (o LookupCertificateResultOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Thumbprint }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCertificateResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}
