// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devcenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a machine pool
//
// Uses Azure REST API version 2024-02-01.
//
// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:devcenter:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPoolArgs struct {
	// Name of the pool.
	PoolName string `pulumi:"poolName"`
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A pool of Virtual Machines.
type LookupPoolResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Indicates the number of provisioned Dev Boxes in this pool.
	DevBoxCount int `pulumi:"devBoxCount"`
	// Name of a Dev Box definition in parent Project of this Pool
	DevBoxDefinitionName string `pulumi:"devBoxDefinitionName"`
	// The display name of the pool.
	DisplayName *string `pulumi:"displayName"`
	// Overall health status of the Pool. Indicates whether or not the Pool is available to create Dev Boxes.
	HealthStatus string `pulumi:"healthStatus"`
	// Details on the Pool health status to help diagnose issues. This is only populated when the pool status indicates the pool is in a non-healthy state
	HealthStatusDetails []HealthStatusDetailResponse `pulumi:"healthStatusDetails"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
	LicenseType string `pulumi:"licenseType"`
	// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
	LocalAdministrator string `pulumi:"localAdministrator"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The regions of the managed virtual network (required when managedNetworkType is Managed).
	ManagedVirtualNetworkRegions []string `pulumi:"managedVirtualNetworkRegions"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Name of a Network Connection in parent Project of this Pool
	NetworkConnectionName string `pulumi:"networkConnectionName"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
	SingleSignOnStatus *string `pulumi:"singleSignOnStatus"`
	// Stop on disconnect configuration settings for Dev Boxes created in this pool.
	StopOnDisconnect *StopOnDisconnectConfigurationResponse `pulumi:"stopOnDisconnect"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
	VirtualNetworkType *string `pulumi:"virtualNetworkType"`
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPoolResultOutput, error) {
			args := v.(LookupPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter:getPool", args, LookupPoolResultOutput{}, options).(LookupPoolResultOutput), nil
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	// Name of the pool.
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}

// A pool of Virtual Machines.
type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupPoolResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Indicates the number of provisioned Dev Boxes in this pool.
func (o LookupPoolResultOutput) DevBoxCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPoolResult) int { return v.DevBoxCount }).(pulumi.IntOutput)
}

// Name of a Dev Box definition in parent Project of this Pool
func (o LookupPoolResultOutput) DevBoxDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.DevBoxDefinitionName }).(pulumi.StringOutput)
}

// The display name of the pool.
func (o LookupPoolResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Overall health status of the Pool. Indicates whether or not the Pool is available to create Dev Boxes.
func (o LookupPoolResultOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.HealthStatus }).(pulumi.StringOutput)
}

// Details on the Pool health status to help diagnose issues. This is only populated when the pool status indicates the pool is in a non-healthy state
func (o LookupPoolResultOutput) HealthStatusDetails() HealthStatusDetailResponseArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []HealthStatusDetailResponse { return v.HealthStatusDetails }).(HealthStatusDetailResponseArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies the license type indicating the caller has already acquired licenses for the Dev Boxes that will be created.
func (o LookupPoolResultOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.LicenseType }).(pulumi.StringOutput)
}

// Indicates whether owners of Dev Boxes in this pool are added as local administrators on the Dev Box.
func (o LookupPoolResultOutput) LocalAdministrator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.LocalAdministrator }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// The regions of the managed virtual network (required when managedNetworkType is Managed).
func (o LookupPoolResultOutput) ManagedVirtualNetworkRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPoolResult) []string { return v.ManagedVirtualNetworkRegions }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name of a Network Connection in parent Project of this Pool
func (o LookupPoolResultOutput) NetworkConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.NetworkConnectionName }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Indicates whether Dev Boxes in this pool are created with single sign on enabled. The also requires that single sign on be enabled on the tenant.
func (o LookupPoolResultOutput) SingleSignOnStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.SingleSignOnStatus }).(pulumi.StringPtrOutput)
}

// Stop on disconnect configuration settings for Dev Boxes created in this pool.
func (o LookupPoolResultOutput) StopOnDisconnect() StopOnDisconnectConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *StopOnDisconnectConfigurationResponse { return v.StopOnDisconnect }).(StopOnDisconnectConfigurationResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Indicates whether the pool uses a Virtual Network managed by Microsoft or a customer provided network.
func (o LookupPoolResultOutput) VirtualNetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.VirtualNetworkType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
