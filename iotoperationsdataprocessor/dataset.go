// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotoperationsdataprocessor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Dataset resource belonging to an Instance resource.
//
// Uses Azure REST API version 2023-10-04-preview. In version 2.x of the Azure Native provider, it used API version 2023-10-04-preview.
type Dataset struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Detailed description of the Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationResponseOutput `pulumi:"extendedLocation"`
	// List of keys that can be used for joining on enrich.
	Keys DatasetPropertyKeyResponseMapOutput `pulumi:"keys"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to the payload in the message. Enrich will add only the payload to the enriched message, other fields will not be kept except for in the indexes.
	Payload pulumi.StringPtrOutput `pulumi:"payload"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Path to an RFC3339 timestamp in the message. If no path is provided, the ingestion time of the record is used for time-based joins.
	Timestamp pulumi.StringPtrOutput `pulumi:"timestamp"`
	// Time to live for individual records.
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataset registers a new resource with the given unique name, arguments, and options.
func NewDataset(ctx *pulumi.Context,
	name string, args *DatasetArgs, opts ...pulumi.ResourceOption) (*Dataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:iotoperationsdataprocessor/v20231004preview:Dataset"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Dataset
	err := ctx.RegisterResource("azure-native:iotoperationsdataprocessor:Dataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataset gets an existing Dataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetState, opts ...pulumi.ResourceOption) (*Dataset, error) {
	var resource Dataset
	err := ctx.ReadResource("azure-native:iotoperationsdataprocessor:Dataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dataset resources.
type datasetState struct {
}

type DatasetState struct {
}

func (DatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetState)(nil)).Elem()
}

type datasetArgs struct {
	// Name of dataset.
	DatasetName *string `pulumi:"datasetName"`
	// Detailed description of the Dataset.
	Description *string `pulumi:"description"`
	// Edge location of the resource.
	ExtendedLocation ExtendedLocation `pulumi:"extendedLocation"`
	// Name of instance.
	InstanceName string `pulumi:"instanceName"`
	// List of keys that can be used for joining on enrich.
	Keys map[string]DatasetPropertyKey `pulumi:"keys"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Path to the payload in the message. Enrich will add only the payload to the enriched message, other fields will not be kept except for in the indexes.
	Payload *string `pulumi:"payload"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Path to an RFC3339 timestamp in the message. If no path is provided, the ingestion time of the record is used for time-based joins.
	Timestamp *string `pulumi:"timestamp"`
	// Time to live for individual records.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a Dataset resource.
type DatasetArgs struct {
	// Name of dataset.
	DatasetName pulumi.StringPtrInput
	// Detailed description of the Dataset.
	Description pulumi.StringPtrInput
	// Edge location of the resource.
	ExtendedLocation ExtendedLocationInput
	// Name of instance.
	InstanceName pulumi.StringInput
	// List of keys that can be used for joining on enrich.
	Keys DatasetPropertyKeyMapInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Path to the payload in the message. Enrich will add only the payload to the enriched message, other fields will not be kept except for in the indexes.
	Payload pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Path to an RFC3339 timestamp in the message. If no path is provided, the ingestion time of the record is used for time-based joins.
	Timestamp pulumi.StringPtrInput
	// Time to live for individual records.
	Ttl pulumi.StringPtrInput
}

func (DatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetArgs)(nil)).Elem()
}

type DatasetInput interface {
	pulumi.Input

	ToDatasetOutput() DatasetOutput
	ToDatasetOutputWithContext(ctx context.Context) DatasetOutput
}

func (*Dataset) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (i *Dataset) ToDatasetOutput() DatasetOutput {
	return i.ToDatasetOutputWithContext(context.Background())
}

func (i *Dataset) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetOutput)
}

type DatasetOutput struct{ *pulumi.OutputState }

func (DatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dataset)(nil)).Elem()
}

func (o DatasetOutput) ToDatasetOutput() DatasetOutput {
	return o
}

func (o DatasetOutput) ToDatasetOutputWithContext(ctx context.Context) DatasetOutput {
	return o
}

// The Azure API version of the resource.
func (o DatasetOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Detailed description of the Dataset.
func (o DatasetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Edge location of the resource.
func (o DatasetOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *Dataset) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// List of keys that can be used for joining on enrich.
func (o DatasetOutput) Keys() DatasetPropertyKeyResponseMapOutput {
	return o.ApplyT(func(v *Dataset) DatasetPropertyKeyResponseMapOutput { return v.Keys }).(DatasetPropertyKeyResponseMapOutput)
}

// The geo-location where the resource lives
func (o DatasetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path to the payload in the message. Enrich will add only the payload to the enriched message, other fields will not be kept except for in the indexes.
func (o DatasetOutput) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.Payload }).(pulumi.StringPtrOutput)
}

// The status of the last operation.
func (o DatasetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DatasetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Dataset) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DatasetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Path to an RFC3339 timestamp in the message. If no path is provided, the ingestion time of the record is used for time-based joins.
func (o DatasetOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.Timestamp }).(pulumi.StringPtrOutput)
}

// Time to live for individual records.
func (o DatasetOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DatasetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dataset) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DatasetOutput{})
}
