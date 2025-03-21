// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a DataSetMapping in a shareSubscription
func LookupBlobStorageAccountDataSetMapping(ctx *pulumi.Context, args *LookupBlobStorageAccountDataSetMappingArgs, opts ...pulumi.InvokeOption) (*LookupBlobStorageAccountDataSetMappingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBlobStorageAccountDataSetMappingResult
	err := ctx.Invoke("azure-native:datashare/v20201001preview:getBlobStorageAccountDataSetMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBlobStorageAccountDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName string `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
}

// A blob storage account data set mapping.
type LookupBlobStorageAccountDataSetMappingResult struct {
	// Gets or sets the container name.
	ContainerName string `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus string `pulumi:"dataSetMappingStatus"`
	// Gets or sets the path to folder within the container.
	Folder string `pulumi:"folder"`
	// The resource id of the azure resource
	Id string `pulumi:"id"`
	// Kind of data set mapping.
	// Expected value is 'BlobStorageAccount'.
	Kind string `pulumi:"kind"`
	// Location of the sink storage account.
	Location string `pulumi:"location"`
	// Gets or sets the mount path on the consumer side where the dataset is to be mapped.
	MountPath *string `pulumi:"mountPath"`
	// Name of the azure resource
	Name string `pulumi:"name"`
	// Provisioning state of the data set mapping.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource id of the sink storage account
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the azure resource
	Type string `pulumi:"type"`
}

func LookupBlobStorageAccountDataSetMappingOutput(ctx *pulumi.Context, args LookupBlobStorageAccountDataSetMappingOutputArgs, opts ...pulumi.InvokeOption) LookupBlobStorageAccountDataSetMappingResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBlobStorageAccountDataSetMappingResultOutput, error) {
			args := v.(LookupBlobStorageAccountDataSetMappingArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datashare/v20201001preview:getBlobStorageAccountDataSetMapping", args, LookupBlobStorageAccountDataSetMappingResultOutput{}, options).(LookupBlobStorageAccountDataSetMappingResultOutput), nil
		}).(LookupBlobStorageAccountDataSetMappingResultOutput)
}

type LookupBlobStorageAccountDataSetMappingOutputArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the dataSetMapping.
	DataSetMappingName pulumi.StringInput `pulumi:"dataSetMappingName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the shareSubscription.
	ShareSubscriptionName pulumi.StringInput `pulumi:"shareSubscriptionName"`
}

func (LookupBlobStorageAccountDataSetMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobStorageAccountDataSetMappingArgs)(nil)).Elem()
}

// A blob storage account data set mapping.
type LookupBlobStorageAccountDataSetMappingResultOutput struct{ *pulumi.OutputState }

func (LookupBlobStorageAccountDataSetMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlobStorageAccountDataSetMappingResult)(nil)).Elem()
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) ToLookupBlobStorageAccountDataSetMappingResultOutput() LookupBlobStorageAccountDataSetMappingResultOutput {
	return o
}

func (o LookupBlobStorageAccountDataSetMappingResultOutput) ToLookupBlobStorageAccountDataSetMappingResultOutputWithContext(ctx context.Context) LookupBlobStorageAccountDataSetMappingResultOutput {
	return o
}

// Gets or sets the container name.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.ContainerName }).(pulumi.StringOutput)
}

// The id of the source data set.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// Gets or sets the path to folder within the container.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Folder }).(pulumi.StringOutput)
}

// The resource id of the azure resource
func (o LookupBlobStorageAccountDataSetMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'BlobStorageAccount'.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Location of the sink storage account.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Location }).(pulumi.StringOutput)
}

// Gets or sets the mount path on the consumer side where the dataset is to be mapped.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

// Name of the azure resource
func (o LookupBlobStorageAccountDataSetMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource id of the sink storage account
func (o LookupBlobStorageAccountDataSetMappingResultOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o LookupBlobStorageAccountDataSetMappingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o LookupBlobStorageAccountDataSetMappingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlobStorageAccountDataSetMappingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlobStorageAccountDataSetMappingResultOutput{})
}
