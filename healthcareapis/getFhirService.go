// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the properties of the specified FHIR Service.
//
// Uses Azure REST API version 2024-03-31.
//
// Other available API versions: 2022-10-01-preview, 2022-12-01, 2023-02-28, 2023-09-06, 2023-11-01, 2023-12-01, 2024-03-01, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthcareapis [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupFhirService(ctx *pulumi.Context, args *LookupFhirServiceArgs, opts ...pulumi.InvokeOption) (*LookupFhirServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFhirServiceResult
	err := ctx.Invoke("azure-native:healthcareapis:getFhirService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFhirServiceArgs struct {
	// The name of FHIR Service resource.
	FhirServiceName string `pulumi:"fhirServiceName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The description of Fhir Service
type LookupFhirServiceResult struct {
	// Fhir Service Azure container registry configuration.
	AcrConfiguration *FhirServiceAcrConfigurationResponse `pulumi:"acrConfiguration"`
	// Fhir Service authentication configuration.
	AuthenticationConfiguration *FhirServiceAuthenticationConfigurationResponse `pulumi:"authenticationConfiguration"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fhir Service Cors configuration.
	CorsConfiguration *FhirServiceCorsConfigurationResponse `pulumi:"corsConfiguration"`
	// The encryption settings of the FHIR service
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// Fhir Service event support status.
	EventState string `pulumi:"eventState"`
	// Fhir Service export configuration.
	ExportConfiguration *FhirServiceExportConfigurationResponse `pulumi:"exportConfiguration"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityResponseIdentity `pulumi:"identity"`
	// Implementation Guides configuration.
	ImplementationGuidesConfiguration *ImplementationGuidesConfigurationResponse `pulumi:"implementationGuidesConfiguration"`
	// Fhir Service import configuration.
	ImportConfiguration *FhirServiceImportConfigurationResponse `pulumi:"importConfiguration"`
	// The kind of the service.
	Kind *string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess string `pulumi:"publicNetworkAccess"`
	// Determines tracking of history for resources.
	ResourceVersionPolicyConfiguration *ResourceVersionPolicyConfigurationResponse `pulumi:"resourceVersionPolicyConfiguration"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupFhirServiceOutput(ctx *pulumi.Context, args LookupFhirServiceOutputArgs, opts ...pulumi.InvokeOption) LookupFhirServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFhirServiceResultOutput, error) {
			args := v.(LookupFhirServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:healthcareapis:getFhirService", args, LookupFhirServiceResultOutput{}, options).(LookupFhirServiceResultOutput), nil
		}).(LookupFhirServiceResultOutput)
}

type LookupFhirServiceOutputArgs struct {
	// The name of FHIR Service resource.
	FhirServiceName pulumi.StringInput `pulumi:"fhirServiceName"`
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFhirServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirServiceArgs)(nil)).Elem()
}

// The description of Fhir Service
type LookupFhirServiceResultOutput struct{ *pulumi.OutputState }

func (LookupFhirServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFhirServiceResult)(nil)).Elem()
}

func (o LookupFhirServiceResultOutput) ToLookupFhirServiceResultOutput() LookupFhirServiceResultOutput {
	return o
}

func (o LookupFhirServiceResultOutput) ToLookupFhirServiceResultOutputWithContext(ctx context.Context) LookupFhirServiceResultOutput {
	return o
}

// Fhir Service Azure container registry configuration.
func (o LookupFhirServiceResultOutput) AcrConfiguration() FhirServiceAcrConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceAcrConfigurationResponse { return v.AcrConfiguration }).(FhirServiceAcrConfigurationResponsePtrOutput)
}

// Fhir Service authentication configuration.
func (o LookupFhirServiceResultOutput) AuthenticationConfiguration() FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceAuthenticationConfigurationResponse {
		return v.AuthenticationConfiguration
	}).(FhirServiceAuthenticationConfigurationResponsePtrOutput)
}

// The Azure API version of the resource.
func (o LookupFhirServiceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fhir Service Cors configuration.
func (o LookupFhirServiceResultOutput) CorsConfiguration() FhirServiceCorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceCorsConfigurationResponse { return v.CorsConfiguration }).(FhirServiceCorsConfigurationResponsePtrOutput)
}

// The encryption settings of the FHIR service
func (o LookupFhirServiceResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o LookupFhirServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fhir Service event support status.
func (o LookupFhirServiceResultOutput) EventState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.EventState }).(pulumi.StringOutput)
}

// Fhir Service export configuration.
func (o LookupFhirServiceResultOutput) ExportConfiguration() FhirServiceExportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceExportConfigurationResponse { return v.ExportConfiguration }).(FhirServiceExportConfigurationResponsePtrOutput)
}

// The resource identifier.
func (o LookupFhirServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o LookupFhirServiceResultOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ServiceManagedIdentityResponseIdentity { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

// Implementation Guides configuration.
func (o LookupFhirServiceResultOutput) ImplementationGuidesConfiguration() ImplementationGuidesConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ImplementationGuidesConfigurationResponse {
		return v.ImplementationGuidesConfiguration
	}).(ImplementationGuidesConfigurationResponsePtrOutput)
}

// Fhir Service import configuration.
func (o LookupFhirServiceResultOutput) ImportConfiguration() FhirServiceImportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *FhirServiceImportConfigurationResponse { return v.ImportConfiguration }).(FhirServiceImportConfigurationResponsePtrOutput)
}

// The kind of the service.
func (o LookupFhirServiceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The resource location.
func (o LookupFhirServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The resource name.
func (o LookupFhirServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of private endpoint connections that are set up for this resource.
func (o LookupFhirServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// The provisioning state.
func (o LookupFhirServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
func (o LookupFhirServiceResultOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

// Determines tracking of history for resources.
func (o LookupFhirServiceResultOutput) ResourceVersionPolicyConfiguration() ResourceVersionPolicyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) *ResourceVersionPolicyConfigurationResponse {
		return v.ResourceVersionPolicyConfiguration
	}).(ResourceVersionPolicyConfigurationResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupFhirServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupFhirServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o LookupFhirServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFhirServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFhirServiceResultOutput{})
}
