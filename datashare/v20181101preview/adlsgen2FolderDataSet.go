// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An ADLS Gen 2 folder data set.
type ADLSGen2FolderDataSet struct {
	pulumi.CustomResourceState

	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// File system to which the folder belongs.
	FileSystem pulumi.StringOutput `pulumi:"fileSystem"`
	// Folder path within the file system.
	FolderPath pulumi.StringOutput `pulumi:"folderPath"`
	// Kind of data set.
	// Expected value is 'AdlsGen2Folder'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource group of storage account
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Storage account name of the source data set
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
	// Subscription id of storage account
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewADLSGen2FolderDataSet registers a new resource with the given unique name, arguments, and options.
func NewADLSGen2FolderDataSet(ctx *pulumi.Context,
	name string, args *ADLSGen2FolderDataSetArgs, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
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
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
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
			Type: pulumi.String("azure-native:datashare:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20201001preview:ADLSGen2FolderDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:ADLSGen2FolderDataSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ADLSGen2FolderDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20181101preview:ADLSGen2FolderDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADLSGen2FolderDataSet gets an existing ADLSGen2FolderDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADLSGen2FolderDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADLSGen2FolderDataSetState, opts ...pulumi.ResourceOption) (*ADLSGen2FolderDataSet, error) {
	var resource ADLSGen2FolderDataSet
	err := ctx.ReadResource("azure-native:datashare/v20181101preview:ADLSGen2FolderDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADLSGen2FolderDataSet resources.
type adlsgen2FolderDataSetState struct {
}

type ADLSGen2FolderDataSetState struct {
}

func (ADLSGen2FolderDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetState)(nil)).Elem()
}

type adlsgen2FolderDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// File system to which the folder belongs.
	FileSystem string `pulumi:"fileSystem"`
	// Folder path within the file system.
	FolderPath string `pulumi:"folderPath"`
	// Kind of data set.
	// Expected value is 'AdlsGen2Folder'.
	Kind string `pulumi:"kind"`
	// Resource group of storage account
	ResourceGroup string `pulumi:"resourceGroup"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
	// Storage account name of the source data set
	StorageAccountName string `pulumi:"storageAccountName"`
	// Subscription id of storage account
	SubscriptionId string `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a ADLSGen2FolderDataSet resource.
type ADLSGen2FolderDataSetArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// File system to which the folder belongs.
	FileSystem pulumi.StringInput
	// Folder path within the file system.
	FolderPath pulumi.StringInput
	// Kind of data set.
	// Expected value is 'AdlsGen2Folder'.
	Kind pulumi.StringInput
	// Resource group of storage account
	ResourceGroup pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
	// Storage account name of the source data set
	StorageAccountName pulumi.StringInput
	// Subscription id of storage account
	SubscriptionId pulumi.StringInput
}

func (ADLSGen2FolderDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adlsgen2FolderDataSetArgs)(nil)).Elem()
}

type ADLSGen2FolderDataSetInput interface {
	pulumi.Input

	ToADLSGen2FolderDataSetOutput() ADLSGen2FolderDataSetOutput
	ToADLSGen2FolderDataSetOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetOutput
}

func (*ADLSGen2FolderDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSet)(nil)).Elem()
}

func (i *ADLSGen2FolderDataSet) ToADLSGen2FolderDataSetOutput() ADLSGen2FolderDataSetOutput {
	return i.ToADLSGen2FolderDataSetOutputWithContext(context.Background())
}

func (i *ADLSGen2FolderDataSet) ToADLSGen2FolderDataSetOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2FolderDataSetOutput)
}

type ADLSGen2FolderDataSetOutput struct{ *pulumi.OutputState }

func (ADLSGen2FolderDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ADLSGen2FolderDataSet)(nil)).Elem()
}

func (o ADLSGen2FolderDataSetOutput) ToADLSGen2FolderDataSetOutput() ADLSGen2FolderDataSetOutput {
	return o
}

func (o ADLSGen2FolderDataSetOutput) ToADLSGen2FolderDataSetOutputWithContext(ctx context.Context) ADLSGen2FolderDataSetOutput {
	return o
}

// Unique id for identifying a data set resource
func (o ADLSGen2FolderDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// File system to which the folder belongs.
func (o ADLSGen2FolderDataSetOutput) FileSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.FileSystem }).(pulumi.StringOutput)
}

// Folder path within the file system.
func (o ADLSGen2FolderDataSetOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'AdlsGen2Folder'.
func (o ADLSGen2FolderDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o ADLSGen2FolderDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource group of storage account
func (o ADLSGen2FolderDataSetOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// Storage account name of the source data set
func (o ADLSGen2FolderDataSetOutput) StorageAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.StorageAccountName }).(pulumi.StringOutput)
}

// Subscription id of storage account
func (o ADLSGen2FolderDataSetOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// Type of the azure resource
func (o ADLSGen2FolderDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ADLSGen2FolderDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ADLSGen2FolderDataSetOutput{})
}