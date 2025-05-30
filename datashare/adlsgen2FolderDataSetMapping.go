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

// An ADLS Gen2 folder data set mapping.
//
// Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.
type ADLSGen2FolderDataSetMapping struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The id of the source data set.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus pulumi.StringOutput `pulumi:"dataSetMappingStatus"`
	// File system to which the folder belongs.
	FileSystem pulumi.StringOutput `pulumi:"fileSystem"`
	// Folder path within the file system.
	FolderPath pulumi.StringOutput `pulumi:"folderPath"`
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2Folder'.
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

// NewADLSGen2FolderDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewADLSGen2FolderDataSetMapping(ctx *pulumi.Context,
	name string, args *ADLSGen2FolderDataSetMappingArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSetMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataSetId == nil {
		return nil, errors.New("invalid value for required argument 'DataSetId'")
	}
	if args.FileSystem == nil {
		return nil, errors.New("invalid value for required argument 'FileSystem'")
	}
	if args.FolderPath == nil {
		return nil, errors.New("invalid value for required argument 'FolderPath'")
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
	args.Kind = pulumi.String("AdlsGen2Folder")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FolderDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2StorageAccountDataSetMapping"),
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
			Type: pulumi.String("azure-native:datashare:BlobContainerDataSetMapping"),
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
	var resource ADLSGen2FolderDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare:ADLSGen2FolderDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADLSGen2FolderDataSetMapping gets an existing ADLSGen2FolderDataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADLSGen2FolderDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FolderDataSetMappingState, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSetMapping, error) {
	var resource ADLSGen2FolderDataSetMapping
	err := ctx.ReadResource("azure-native:datashare:ADLSGen2FolderDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADLSGen2FolderDataSetMapping resources.
type adlsgen2FolderDataSetMappingState struct {
}

type ADLSGen2FolderDataSetMappingState struct {
}

func (ADLSGen2FolderDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetMappingState)(nil)).Elem()
}

type adlsgen2FolderDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// File system to which the folder belongs.
	FileSystem string `pulumi:"fileSystem"`
	// Folder path within the file system.
	FolderPath string `pulumi:"folderPath"`
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2Folder'.
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

// The set of arguments for constructing a ADLSGen2FolderDataSetMapping resource.
type ADLSGen2FolderDataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The id of the source data set.
	DataSetId pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// File system to which the folder belongs.
	FileSystem pulumi.StringInput
	// Folder path within the file system.
	FolderPath pulumi.StringInput
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2Folder'.
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

func (ADLSGen2FolderDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetMappingArgs)(nil)).Elem()
}

type ADLSGen2FolderDataSetMappingInput interface {
	pulumi.Input

	ToADLSGen2FolderDataSetMappingOutput() ADLSGen2FolderDataSetMappingOutput
	ToADLSGen2FolderDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetMappingOutput
}

func (*ADLSGen2FolderDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSetMapping)(nil)).Elem()
}

func (i *ADLSGen2FolderDataSetMapping) ToADLSGen2FolderDataSetMappingOutput() ADLSGen2FolderDataSetMappingOutput {
	return i.ToADLSGen2FolderDataSetMappingOutputWithContext(context.Background())
}

func (i *ADLSGen2FolderDataSetMapping) ToADLSGen2FolderDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FolderDataSetMappingOutput)
}

type ADLSGen2FolderDataSetMappingOutput struct{ *pulumi.OutputState }

func (ADLSGen2FolderDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSetMapping)(nil)).Elem()
}

func (o ADLSGen2FolderDataSetMappingOutput) ToADLSGen2FolderDataSetMappingOutput() ADLSGen2FolderDataSetMappingOutput {
	return o
}

func (o ADLSGen2FolderDataSetMappingOutput) ToADLSGen2FolderDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetMappingOutput {
	return o
}

// The Azure API version of the resource.
func (o ADLSGen2FolderDataSetMappingOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The id of the source data set.
func (o ADLSGen2FolderDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o ADLSGen2FolderDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// File system to which the folder belongs.
func (o ADLSGen2FolderDataSetMappingOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

// Folder path within the file system.
func (o ADLSGen2FolderDataSetMappingOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'AdlsGen2Folder'.
func (o ADLSGen2FolderDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o ADLSGen2FolderDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o ADLSGen2FolderDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource group of storage account.
func (o ADLSGen2FolderDataSetMappingOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set.
func (o ADLSGen2FolderDataSetMappingOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account.
func (o ADLSGen2FolderDataSetMappingOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o ADLSGen2FolderDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o ADLSGen2FolderDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FolderDataSetMappingOutput{})
}
