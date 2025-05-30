// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Blob container data set mapping.
//
// Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.
type BlobContainerDataSetMapping struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// BLOB Container name.
	ContainerName pulumi.StringOutput `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus pulumi.StringOutput `pulumi:"dataSetMappingStatus"`
	// Kind of data set mapping.
	// Expected value is 'Container'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the data set mapping.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource group of storage account.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Storage account name of the source data set.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// Subscription id of storage account.
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlobContainerDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewBlobContainerDataSetMapping(ctx *pulumi.Context,
	name string, args *BlobContainerDataSetMappingArgs, opts ...pulumi.ResourceOption) (*BlobContainerDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ContainerName == nil {
		return nil, errors.New("invalid value for required argument 'ContainerName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareSubscriptionName == nil {
		return nil, errors.New("invalid value for required argument 'ShareSubscriptionName'")
	}
	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	args.Kind = pulumi.String("Container")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:BlobStorageAccountDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FileDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobContainerDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobFolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:KustoTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDBTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SqlDWTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FileDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:BlobDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:BlobFolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:KustoClusterDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:KustoDatabaseDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:KustoTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:SqlDBTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:SqlDWTableDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare:SynapseWorkspaceSqlPoolTableDataSetMapping"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BlobContainerDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare:BlobContainerDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobContainerDataSetMapping gets an existing BlobContainerDataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobContainerDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerDataSetMappingState, opts ...pulumi.ResourceOption) (*BlobContainerDataSetMapping, error) {
	var resource BlobContainerDataSetMapping
	err := ctx.ReadResource("azure-native:datashare:BlobContainerDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobContainerDataSetMapping resources.
type blobContainerDataSetMappingState struct {
}

type BlobContainerDataSetMappingState struct {
}

func (BlobContainerDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetMappingState)(nil)).Elem()
}

type blobContainerDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// BLOB Container name.
	ContainerName string `pulumi:"containerName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// Kind of data set mapping.
	// Expected value is 'Container'.
	Kind string `pulumi:"kind"`
	// Resource group of storage account.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName string `pulumi:"shareSubscriptionName"`
	// Storage account name of the source data set.
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account.
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a BlobContainerDataSetMapping resource.
type BlobContainerDataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// BLOB Container name.
	ContainerName pulumi.StringInput
	// The id of the source data set.
	DataSetId pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// Kind of data set mapping.
	// Expected value is 'Container'.
	Kind pulumi.StringInput
	// Resource group of storage account.
	ResourceGroup pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share subscription which will hold the data set sink.
	ShareSubscriptionName pulumi.StringInput
	// Storage account name of the source data set.
	StorageAccountName pulumi.StringInput
	// Subscription id of storage account.
	SubscriptionId pulumi.StringInput
}

func (BlobContainerDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerDataSetMappingArgs)(nil)).Elem()
}

type BlobContainerDataSetMappingInput interface {
	pulumi.Input

	ToBlobContainerDataSetMappingOutput() BlobContainerDataSetMappingOutput
	ToBlobContainerDataSetMappingOutputWithContext(ctx context.Context) BlobContainerDataSetMappingOutput
}

func (*BlobContainerDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerDataSetMapping)(nil)).Elem()
}

func (i *BlobContainerDataSetMapping) ToBlobContainerDataSetMappingOutput() BlobContainerDataSetMappingOutput {
	return i.ToBlobContainerDataSetMappingOutputWithContext(context.Background())
}

func (i *BlobContainerDataSetMapping) ToBlobContainerDataSetMappingOutputWithContext(ctx context.Context) BlobContainerDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerDataSetMappingOutput)
}

type BlobContainerDataSetMappingOutput struct{ *pulumi.OutputState }

func (BlobContainerDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainerDataSetMapping)(nil)).Elem()
}

func (o BlobContainerDataSetMappingOutput) ToBlobContainerDataSetMappingOutput() BlobContainerDataSetMappingOutput {
	return o
}

func (o BlobContainerDataSetMappingOutput) ToBlobContainerDataSetMappingOutputWithContext(ctx context.Context) BlobContainerDataSetMappingOutput {
	return o
}

// The Azure API version of the resource.
func (o BlobContainerDataSetMappingOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// BLOB Container name.
func (o BlobContainerDataSetMappingOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.ContainerName }).(pulumi.StringOutput)
}

// The id of the source data set.
func (o BlobContainerDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o BlobContainerDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'Container'.
func (o BlobContainerDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o BlobContainerDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o BlobContainerDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource group of storage account.
func (o BlobContainerDataSetMappingOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set.
func (o BlobContainerDataSetMappingOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account.
func (o BlobContainerDataSetMappingOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o BlobContainerDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o BlobContainerDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainerDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobContainerDataSetMappingOutput{})
}
