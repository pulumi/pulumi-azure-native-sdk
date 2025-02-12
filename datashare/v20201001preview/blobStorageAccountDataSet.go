// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Azure blob storage account data set.
type BlobStorageAccountDataSet struct {
	pulumi.CustomResourceState

	// Unique id for identifying a data set resource
	DataSetId pulumi.StringOutput `pulumi:"dataSetId"`
	// Kind of data set.
	// Expected value is 'BlobStorageAccount'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Location of the storage account.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the azure resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of storage account paths.
	Paths BlobStorageAccountPathResponseArrayOutput `pulumi:"paths"`
	// Resource id of the storage account.
	StorageAccountResourceId pulumi.StringOutput `pulumi:"storageAccountResourceId"`
	// System Data of the Azure resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the azure resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlobStorageAccountDataSet registers a new resource with the given unique name, arguments, and options.
func NewBlobStorageAccountDataSet(ctx *pulumi.Context,
	name string, args *BlobStorageAccountDataSetArgs, opts ...pulumi.ResourceOption) (*BlobStorageAccountDataSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ShareName == nil {
		return nil, errors.New("invalid value for required argument 'ShareName'")
	}
	if args.StorageAccountResourceId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountResourceId'")
	}
	args.Kind = pulumi.String("BlobStorageAccount")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datashare/v20181101preview:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20191101:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20200901:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare/v20210801:BlobStorageAccountDataSet"),
		},
		{
			Type: pulumi.String("azure-native:datashare:BlobStorageAccountDataSet"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BlobStorageAccountDataSet
	err := ctx.RegisterResource("azure-native:datashare/v20201001preview:BlobStorageAccountDataSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobStorageAccountDataSet gets an existing BlobStorageAccountDataSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobStorageAccountDataSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobStorageAccountDataSetState, opts ...pulumi.ResourceOption) (*BlobStorageAccountDataSet, error) {
	var resource BlobStorageAccountDataSet
	err := ctx.ReadResource("azure-native:datashare/v20201001preview:BlobStorageAccountDataSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobStorageAccountDataSet resources.
type blobStorageAccountDataSetState struct {
}

type BlobStorageAccountDataSetState struct {
}

func (BlobStorageAccountDataSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobStorageAccountDataSetState)(nil)).Elem()
}

type blobStorageAccountDataSetArgs struct {
	// The name of the share account.
	AccountName string `pulumi:"accountName"`
	// The name of the dataSet.
	DataSetName *string `pulumi:"dataSetName"`
	// Kind of data set.
	// Expected value is 'BlobStorageAccount'.
	Kind string `pulumi:"kind"`
	// A list of storage account paths.
	Paths []BlobStorageAccountPath `pulumi:"paths"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the share to add the data set to.
	ShareName string `pulumi:"shareName"`
	// Resource id of the storage account.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
}

// The set of arguments for constructing a BlobStorageAccountDataSet resource.
type BlobStorageAccountDataSetArgs struct {
	// The name of the share account.
	AccountName pulumi.StringInput
	// The name of the dataSet.
	DataSetName pulumi.StringPtrInput
	// Kind of data set.
	// Expected value is 'BlobStorageAccount'.
	Kind pulumi.StringInput
	// A list of storage account paths.
	Paths BlobStorageAccountPathArrayInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The name of the share to add the data set to.
	ShareName pulumi.StringInput
	// Resource id of the storage account.
	StorageAccountResourceId pulumi.StringInput
}

func (BlobStorageAccountDataSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobStorageAccountDataSetArgs)(nil)).Elem()
}

type BlobStorageAccountDataSetInput interface {
	pulumi.Input

	ToBlobStorageAccountDataSetOutput() BlobStorageAccountDataSetOutput
	ToBlobStorageAccountDataSetOutputWithContext(ctx context.Context) BlobStorageAccountDataSetOutput
}

func (*BlobStorageAccountDataSet) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageAccountDataSet)(nil)).Elem()
}

func (i *BlobStorageAccountDataSet) ToBlobStorageAccountDataSetOutput() BlobStorageAccountDataSetOutput {
	return i.ToBlobStorageAccountDataSetOutputWithContext(context.Background())
}

func (i *BlobStorageAccountDataSet) ToBlobStorageAccountDataSetOutputWithContext(ctx context.Context) BlobStorageAccountDataSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageAccountDataSetOutput)
}

type BlobStorageAccountDataSetOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountDataSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageAccountDataSet)(nil)).Elem()
}

func (o BlobStorageAccountDataSetOutput) ToBlobStorageAccountDataSetOutput() BlobStorageAccountDataSetOutput {
	return o
}

func (o BlobStorageAccountDataSetOutput) ToBlobStorageAccountDataSetOutputWithContext(ctx context.Context) BlobStorageAccountDataSetOutput {
	return o
}

// Unique id for identifying a data set resource
func (o BlobStorageAccountDataSetOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) pulumi.StringOutput { return v.DataSetId }).(pulumi.StringOutput)
}

// Kind of data set.
// Expected value is 'BlobStorageAccount'.
func (o BlobStorageAccountDataSetOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Location of the storage account.
func (o BlobStorageAccountDataSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the azure resource
func (o BlobStorageAccountDataSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of storage account paths.
func (o BlobStorageAccountDataSetOutput) Paths() BlobStorageAccountPathResponseArrayOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) BlobStorageAccountPathResponseArrayOutput { return v.Paths }).(BlobStorageAccountPathResponseArrayOutput)
}

// Resource id of the storage account.
func (o BlobStorageAccountDataSetOutput) StorageAccountResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) pulumi.StringOutput { return v.StorageAccountResourceId }).(pulumi.StringOutput)
}

// System Data of the Azure resource.
func (o BlobStorageAccountDataSetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the azure resource
func (o BlobStorageAccountDataSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobStorageAccountDataSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobStorageAccountDataSetOutput{})
}
