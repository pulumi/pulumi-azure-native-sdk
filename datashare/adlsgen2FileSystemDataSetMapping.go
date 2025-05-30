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

// An ADLS Gen2 file system data set mapping.
//
// Uses Azure REST API version 2021-08-01. In version 2.x of the Azure Native provider, it used API version 2021-08-01.
type ADLSGen2FileSystemDataSetMapping struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The id of the source data set.
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Gets the status of the data set mapping.
	DataSetMappingStatus pulumi.StringOutput `pulumi:"dataSetMappingStatus"`
	// The file system name.
	FileSystem pulumi.StringOutput `pulumi:"fileSystem"`
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2FileSystem'.
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

// NewADLSGen2FileSystemDataSetMapping registers a new resource with the given unique name, arguments, and options.
func NewADLSGen2FileSystemDataSetMapping(ctx *pulumi.Context,
	name string, args *ADLSGen2FileSystemDataSetMappingArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FileSystemDataSetMapping, error) {
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
	args.Kind = pulumi.String("AdlsGen2FileSystem")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FileSystemDataSetMapping"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FileSystemDataSetMapping"),
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
			Type: pulumi.String("azure-native:datashare:ADLSGen2FolderDataSetMapping"),
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
	var resource ADLSGen2FileSystemDataSetMapping
	err := ctx.RegisterResource("azure-native:datashare:ADLSGen2FileSystemDataSetMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADLSGen2FileSystemDataSetMapping gets an existing ADLSGen2FileSystemDataSetMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADLSGen2FileSystemDataSetMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FileSystemDataSetMappingState, opts ...pulumi.ResourceOption) (*ADLSGen2FileSystemDataSetMapping, error) {
	var resource ADLSGen2FileSystemDataSetMapping
	err := ctx.ReadResource("azure-native:datashare:ADLSGen2FileSystemDataSetMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADLSGen2FileSystemDataSetMapping resources.
type adlsgen2FileSystemDataSetMappingState struct {
}

type ADLSGen2FileSystemDataSetMappingState struct {
}

func (ADLSGen2FileSystemDataSetMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileSystemDataSetMappingState)(nil)).Elem()
}

type adlsgen2FileSystemDataSetMappingArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The id of the source data set.
	DataSetId string `pulumi:"dataSetId"`
	// The name of the data set mapping to be created.
	DataSetMappingName *string `pulumi:"dataSetMappingName"`
	// The file system name.
	FileSystem string `pulumi:"fileSystem"`
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2FileSystem'.
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

// The set of arguments for constructing a ADLSGen2FileSystemDataSetMapping resource.
type ADLSGen2FileSystemDataSetMappingArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The id of the source data set.
	DataSetId pulumi.StringInput
	// The name of the data set mapping to be created.
	DataSetMappingName pulumi.StringPtrInput
	// The file system name.
	FileSystem pulumi.StringInput
	// Kind of data set mapping.
	// Expected value is 'AdlsGen2FileSystem'.
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

func (ADLSGen2FileSystemDataSetMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FileSystemDataSetMappingArgs)(nil)).Elem()
}

type ADLSGen2FileSystemDataSetMappingInput interface {
	pulumi.Input

	ToADLSGen2FileSystemDataSetMappingOutput() ADLSGen2FileSystemDataSetMappingOutput
	ToADLSGen2FileSystemDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetMappingOutput
}

func (*ADLSGen2FileSystemDataSetMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileSystemDataSetMapping)(nil)).Elem()
}

func (i *ADLSGen2FileSystemDataSetMapping) ToADLSGen2FileSystemDataSetMappingOutput() ADLSGen2FileSystemDataSetMappingOutput {
	return i.ToADLSGen2FileSystemDataSetMappingOutputWithContext(context.Background())
}

func (i *ADLSGen2FileSystemDataSetMapping) ToADLSGen2FileSystemDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FileSystemDataSetMappingOutput)
}

type ADLSGen2FileSystemDataSetMappingOutput struct{ *pulumi.OutputState }

func (ADLSGen2FileSystemDataSetMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FileSystemDataSetMapping)(nil)).Elem()
}

func (o ADLSGen2FileSystemDataSetMappingOutput) ToADLSGen2FileSystemDataSetMappingOutput() ADLSGen2FileSystemDataSetMappingOutput {
	return o
}

func (o ADLSGen2FileSystemDataSetMappingOutput) ToADLSGen2FileSystemDataSetMappingOutputWithContext(ctx context.Context) ADLSGen2FileSystemDataSetMappingOutput {
	return o
}

// The Azure API version of the resource.
func (o ADLSGen2FileSystemDataSetMappingOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The id of the source data set.
func (o ADLSGen2FileSystemDataSetMappingOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Gets the status of the data set mapping.
func (o ADLSGen2FileSystemDataSetMappingOutput) DataSetMappingStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.DataSetMappingStatus }).(pulumi.StringOutput)
}

// The file system name.
func (o ADLSGen2FileSystemDataSetMappingOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

// Kind of data set mapping.
// Expected value is 'AdlsGen2FileSystem'.
func (o ADLSGen2FileSystemDataSetMappingOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o ADLSGen2FileSystemDataSetMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the data set mapping.
func (o ADLSGen2FileSystemDataSetMappingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource group of storage account.
func (o ADLSGen2FileSystemDataSetMappingOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set.
func (o ADLSGen2FileSystemDataSetMappingOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account.
func (o ADLSGen2FileSystemDataSetMappingOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o ADLSGen2FileSystemDataSetMappingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o ADLSGen2FileSystemDataSetMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FileSystemDataSetMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FileSystemDataSetMappingOutput{})
}
