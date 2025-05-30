// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get details of the specified capacity pool
//
// Uses Azure REST API version 2024-09-01.
//
// Other available API versions: 2022-11-01, 2022-11-01-preview, 2023-05-01, 2023-05-01-preview, 2023-07-01, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01-preview, 2025-01-01, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native netapp [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupCapacityPool(ctx *pulumi.Context, args *LookupCapacityPoolArgs, opts ...pulumi.InvokeOption) (*LookupCapacityPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCapacityPoolResult
	err := ctx.Invoke("azure-native:netapp:getCapacityPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCapacityPoolArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Capacity pool resource
type LookupCapacityPoolResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// If enabled (true) the pool can contain cool Access enabled volumes.
	CoolAccess *bool `pulumi:"coolAccess"`
	// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
	EncryptionType *string `pulumi:"encryptionType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// UUID v4 used to identify the Pool
	PoolId string `pulumi:"poolId"`
	// Azure lifecycle management
	ProvisioningState string `pulumi:"provisioningState"`
	// The qos type of the pool
	QosType *string `pulumi:"qosType"`
	// The service level of the file system
	ServiceLevel string `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 1TiB chunks (value must be multiple of 1099511627776).
	Size float64 `pulumi:"size"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Total throughput of pool in MiB/s
	TotalThroughputMibps float64 `pulumi:"totalThroughputMibps"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Utilized throughput of pool in MiB/s
	UtilizedThroughputMibps float64 `pulumi:"utilizedThroughputMibps"`
}

// Defaults sets the appropriate defaults for LookupCapacityPoolResult
func (val *LookupCapacityPoolResult) Defaults() *LookupCapacityPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.CoolAccess == nil {
		coolAccess_ := false
		tmp.CoolAccess = &coolAccess_
	}
	if tmp.EncryptionType == nil {
		encryptionType_ := "Single"
		tmp.EncryptionType = &encryptionType_
	}
	if tmp.QosType == nil {
		qosType_ := "Auto"
		tmp.QosType = &qosType_
	}
	if utilities.IsZero(tmp.Size) {
		tmp.Size = 4398046511104.0
	}
	return &tmp
}
func LookupCapacityPoolOutput(ctx *pulumi.Context, args LookupCapacityPoolOutputArgs, opts ...pulumi.InvokeOption) LookupCapacityPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCapacityPoolResultOutput, error) {
			args := v.(LookupCapacityPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:netapp:getCapacityPool", args, LookupCapacityPoolResultOutput{}, options).(LookupCapacityPoolResultOutput), nil
		}).(LookupCapacityPoolResultOutput)
}

type LookupCapacityPoolOutputArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the capacity pool
	PoolName pulumi.StringInput `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCapacityPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityPoolArgs)(nil)).Elem()
}

// Capacity pool resource
type LookupCapacityPoolResultOutput struct{ *pulumi.OutputState }

func (LookupCapacityPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapacityPoolResult)(nil)).Elem()
}

func (o LookupCapacityPoolResultOutput) ToLookupCapacityPoolResultOutput() LookupCapacityPoolResultOutput {
	return o
}

func (o LookupCapacityPoolResultOutput) ToLookupCapacityPoolResultOutputWithContext(ctx context.Context) LookupCapacityPoolResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupCapacityPoolResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// If enabled (true) the pool can contain cool Access enabled volumes.
func (o LookupCapacityPoolResultOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) *bool { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
func (o LookupCapacityPoolResultOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) *string { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupCapacityPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupCapacityPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupCapacityPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCapacityPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// UUID v4 used to identify the Pool
func (o LookupCapacityPoolResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.PoolId }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o LookupCapacityPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The qos type of the pool
func (o LookupCapacityPoolResultOutput) QosType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) *string { return v.QosType }).(pulumi.StringPtrOutput)
}

// The service level of the file system
func (o LookupCapacityPoolResultOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.ServiceLevel }).(pulumi.StringOutput)
}

// Provisioned size of the pool (in bytes). Allowed values are in 1TiB chunks (value must be multiple of 1099511627776).
func (o LookupCapacityPoolResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCapacityPoolResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCapacityPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCapacityPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Total throughput of pool in MiB/s
func (o LookupCapacityPoolResultOutput) TotalThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCapacityPoolResult) float64 { return v.TotalThroughputMibps }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCapacityPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapacityPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

// Utilized throughput of pool in MiB/s
func (o LookupCapacityPoolResultOutput) UtilizedThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCapacityPoolResult) float64 { return v.UtilizedThroughputMibps }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupCapacityPoolResultOutput{})
}
