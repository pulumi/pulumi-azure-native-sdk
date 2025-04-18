// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurelargeinstance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an Azure Large Storage instance for the specified subscription, resource
// group, and instance name.
//
// Uses Azure REST API version 2024-08-01-preview.
func LookupAzureLargeStorageInstance(ctx *pulumi.Context, args *LookupAzureLargeStorageInstanceArgs, opts ...pulumi.InvokeOption) (*LookupAzureLargeStorageInstanceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAzureLargeStorageInstanceResult
	err := ctx.Invoke("azure-native:azurelargeinstance:getAzureLargeStorageInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAzureLargeStorageInstanceArgs struct {
	// Name of the AzureLargeStorageInstance.
	AzureLargeStorageInstanceName string `pulumi:"azureLargeStorageInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// AzureLargeStorageInstance info on Azure (ARM properties and
// AzureLargeStorageInstance properties)
type LookupAzureLargeStorageInstanceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Specifies the AzureLargeStorageInstance unique ID.
	AzureLargeStorageInstanceUniqueIdentifier *string `pulumi:"azureLargeStorageInstanceUniqueIdentifier"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Specifies the storage properties for the AzureLargeStorage instance.
	StorageProperties *StoragePropertiesResponse `pulumi:"storageProperties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAzureLargeStorageInstanceOutput(ctx *pulumi.Context, args LookupAzureLargeStorageInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupAzureLargeStorageInstanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAzureLargeStorageInstanceResultOutput, error) {
			args := v.(LookupAzureLargeStorageInstanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azurelargeinstance:getAzureLargeStorageInstance", args, LookupAzureLargeStorageInstanceResultOutput{}, options).(LookupAzureLargeStorageInstanceResultOutput), nil
		}).(LookupAzureLargeStorageInstanceResultOutput)
}

type LookupAzureLargeStorageInstanceOutputArgs struct {
	// Name of the AzureLargeStorageInstance.
	AzureLargeStorageInstanceName pulumi.StringInput `pulumi:"azureLargeStorageInstanceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAzureLargeStorageInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureLargeStorageInstanceArgs)(nil)).Elem()
}

// AzureLargeStorageInstance info on Azure (ARM properties and
// AzureLargeStorageInstance properties)
type LookupAzureLargeStorageInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupAzureLargeStorageInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureLargeStorageInstanceResult)(nil)).Elem()
}

func (o LookupAzureLargeStorageInstanceResultOutput) ToLookupAzureLargeStorageInstanceResultOutput() LookupAzureLargeStorageInstanceResultOutput {
	return o
}

func (o LookupAzureLargeStorageInstanceResultOutput) ToLookupAzureLargeStorageInstanceResultOutputWithContext(ctx context.Context) LookupAzureLargeStorageInstanceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupAzureLargeStorageInstanceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Specifies the AzureLargeStorageInstance unique ID.
func (o LookupAzureLargeStorageInstanceResultOutput) AzureLargeStorageInstanceUniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) *string {
		return v.AzureLargeStorageInstanceUniqueIdentifier
	}).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupAzureLargeStorageInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o LookupAzureLargeStorageInstanceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupAzureLargeStorageInstanceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAzureLargeStorageInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the storage properties for the AzureLargeStorage instance.
func (o LookupAzureLargeStorageInstanceResultOutput) StorageProperties() StoragePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) *StoragePropertiesResponse { return v.StorageProperties }).(StoragePropertiesResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAzureLargeStorageInstanceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAzureLargeStorageInstanceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAzureLargeStorageInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureLargeStorageInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureLargeStorageInstanceResultOutput{})
}
