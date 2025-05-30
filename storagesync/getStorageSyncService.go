// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagesync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a given StorageSyncService.
//
// Uses Azure REST API version 2022-09-01.
//
// Other available API versions: 2022-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagesync [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupStorageSyncService(ctx *pulumi.Context, args *LookupStorageSyncServiceArgs, opts ...pulumi.InvokeOption) (*LookupStorageSyncServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageSyncServiceResult
	err := ctx.Invoke("azure-native:storagesync:getStorageSyncService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageSyncServiceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName string `pulumi:"storageSyncServiceName"`
}

// Storage Sync Service object.
type LookupStorageSyncServiceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// managed identities for the Storage Sync service to interact with other Azure services without maintaining any secrets or credentials in code.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Incoming Traffic Policy
	IncomingTrafficPolicy *string `pulumi:"incomingTrafficPolicy"`
	// Resource Last Operation Name
	LastOperationName string `pulumi:"lastOperationName"`
	// StorageSyncService lastWorkflowId
	LastWorkflowId string `pulumi:"lastWorkflowId"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of private endpoint connection associated with the specified storage sync service
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// StorageSyncService Provisioning State
	ProvisioningState string `pulumi:"provisioningState"`
	// Storage Sync service status.
	StorageSyncServiceStatus int `pulumi:"storageSyncServiceStatus"`
	// Storage Sync service Uid
	StorageSyncServiceUid string `pulumi:"storageSyncServiceUid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Use Identity authorization when customer have finished setup RBAC permissions.
	UseIdentity bool `pulumi:"useIdentity"`
}

func LookupStorageSyncServiceOutput(ctx *pulumi.Context, args LookupStorageSyncServiceOutputArgs, opts ...pulumi.InvokeOption) LookupStorageSyncServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStorageSyncServiceResultOutput, error) {
			args := v.(LookupStorageSyncServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:storagesync:getStorageSyncService", args, LookupStorageSyncServiceResultOutput{}, options).(LookupStorageSyncServiceResultOutput), nil
		}).(LookupStorageSyncServiceResultOutput)
}

type LookupStorageSyncServiceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Storage Sync Service resource.
	StorageSyncServiceName pulumi.StringInput `pulumi:"storageSyncServiceName"`
}

func (LookupStorageSyncServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageSyncServiceArgs)(nil)).Elem()
}

// Storage Sync Service object.
type LookupStorageSyncServiceResultOutput struct{ *pulumi.OutputState }

func (LookupStorageSyncServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageSyncServiceResult)(nil)).Elem()
}

func (o LookupStorageSyncServiceResultOutput) ToLookupStorageSyncServiceResultOutput() LookupStorageSyncServiceResultOutput {
	return o
}

func (o LookupStorageSyncServiceResultOutput) ToLookupStorageSyncServiceResultOutputWithContext(ctx context.Context) LookupStorageSyncServiceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupStorageSyncServiceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupStorageSyncServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// managed identities for the Storage Sync service to interact with other Azure services without maintaining any secrets or credentials in code.
func (o LookupStorageSyncServiceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Incoming Traffic Policy
func (o LookupStorageSyncServiceResultOutput) IncomingTrafficPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) *string { return v.IncomingTrafficPolicy }).(pulumi.StringPtrOutput)
}

// Resource Last Operation Name
func (o LookupStorageSyncServiceResultOutput) LastOperationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.LastOperationName }).(pulumi.StringOutput)
}

// StorageSyncService lastWorkflowId
func (o LookupStorageSyncServiceResultOutput) LastWorkflowId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.LastWorkflowId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupStorageSyncServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupStorageSyncServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connection associated with the specified storage sync service
func (o LookupStorageSyncServiceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

// StorageSyncService Provisioning State
func (o LookupStorageSyncServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Storage Sync service status.
func (o LookupStorageSyncServiceResultOutput) StorageSyncServiceStatus() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) int { return v.StorageSyncServiceStatus }).(pulumi.IntOutput)
}

// Storage Sync service Uid
func (o LookupStorageSyncServiceResultOutput) StorageSyncServiceUid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.StorageSyncServiceUid }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStorageSyncServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupStorageSyncServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupStorageSyncServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Use Identity authorization when customer have finished setup RBAC permissions.
func (o LookupStorageSyncServiceResultOutput) UseIdentity() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageSyncServiceResult) bool { return v.UseIdentity }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageSyncServiceResultOutput{})
}
