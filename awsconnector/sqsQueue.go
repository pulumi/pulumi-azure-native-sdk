// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Microsoft.AwsConnector resource
//
// Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2024-12-01.
type SqsQueue struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties SqsQueuePropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqsQueue registers a new resource with the given unique name, arguments, and options.
func NewSqsQueue(ctx *pulumi.Context,
	name string, args *SqsQueueArgs, opts ...pulumi.ResourceOption) (*SqsQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:SqsQueue"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SqsQueue
	err := ctx.RegisterResource("azure-native:awsconnector:SqsQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqsQueue gets an existing SqsQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqsQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqsQueueState, opts ...pulumi.ResourceOption) (*SqsQueue, error) {
	var resource SqsQueue
	err := ctx.ReadResource("azure-native:awsconnector:SqsQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqsQueue resources.
type sqsQueueState struct {
}

type SqsQueueState struct {
}

func (SqsQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqsQueueState)(nil)).Elem()
}

type sqsQueueArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of SqsQueue
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *SqsQueueProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SqsQueue resource.
type SqsQueueArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of SqsQueue
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties SqsQueuePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SqsQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqsQueueArgs)(nil)).Elem()
}

type SqsQueueInput interface {
	pulumi.Input

	ToSqsQueueOutput() SqsQueueOutput
	ToSqsQueueOutputWithContext(ctx context.Context) SqsQueueOutput
}

func (*SqsQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**SqsQueue)(nil)).Elem()
}

func (i *SqsQueue) ToSqsQueueOutput() SqsQueueOutput {
	return i.ToSqsQueueOutputWithContext(context.Background())
}

func (i *SqsQueue) ToSqsQueueOutputWithContext(ctx context.Context) SqsQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqsQueueOutput)
}

type SqsQueueOutput struct{ *pulumi.OutputState }

func (SqsQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqsQueue)(nil)).Elem()
}

func (o SqsQueueOutput) ToSqsQueueOutput() SqsQueueOutput {
	return o
}

func (o SqsQueueOutput) ToSqsQueueOutputWithContext(ctx context.Context) SqsQueueOutput {
	return o
}

// The Azure API version of the resource.
func (o SqsQueueOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o SqsQueueOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SqsQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o SqsQueueOutput) Properties() SqsQueuePropertiesResponseOutput {
	return o.ApplyT(func(v *SqsQueue) SqsQueuePropertiesResponseOutput { return v.Properties }).(SqsQueuePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SqsQueueOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SqsQueue) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SqsQueueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SqsQueueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqsQueue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqsQueueOutput{})
}
