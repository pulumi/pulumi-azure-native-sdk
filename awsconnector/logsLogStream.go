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
type LogsLogStream struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties LogsLogStreamPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLogsLogStream registers a new resource with the given unique name, arguments, and options.
func NewLogsLogStream(ctx *pulumi.Context,
	name string, args *LogsLogStreamArgs, opts ...pulumi.ResourceOption) (*LogsLogStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:LogsLogStream"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LogsLogStream
	err := ctx.RegisterResource("azure-native:awsconnector:LogsLogStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogsLogStream gets an existing LogsLogStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogsLogStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogsLogStreamState, opts ...pulumi.ResourceOption) (*LogsLogStream, error) {
	var resource LogsLogStream
	err := ctx.ReadResource("azure-native:awsconnector:LogsLogStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogsLogStream resources.
type logsLogStreamState struct {
}

type LogsLogStreamState struct {
}

func (LogsLogStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*logsLogStreamState)(nil)).Elem()
}

type logsLogStreamArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of LogsLogStream
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *LogsLogStreamProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LogsLogStream resource.
type LogsLogStreamArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of LogsLogStream
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties LogsLogStreamPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LogsLogStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logsLogStreamArgs)(nil)).Elem()
}

type LogsLogStreamInput interface {
	pulumi.Input

	ToLogsLogStreamOutput() LogsLogStreamOutput
	ToLogsLogStreamOutputWithContext(ctx context.Context) LogsLogStreamOutput
}

func (*LogsLogStream) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsLogStream)(nil)).Elem()
}

func (i *LogsLogStream) ToLogsLogStreamOutput() LogsLogStreamOutput {
	return i.ToLogsLogStreamOutputWithContext(context.Background())
}

func (i *LogsLogStream) ToLogsLogStreamOutputWithContext(ctx context.Context) LogsLogStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsLogStreamOutput)
}

type LogsLogStreamOutput struct{ *pulumi.OutputState }

func (LogsLogStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsLogStream)(nil)).Elem()
}

func (o LogsLogStreamOutput) ToLogsLogStreamOutput() LogsLogStreamOutput {
	return o
}

func (o LogsLogStreamOutput) ToLogsLogStreamOutputWithContext(ctx context.Context) LogsLogStreamOutput {
	return o
}

// The Azure API version of the resource.
func (o LogsLogStreamOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsLogStream) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LogsLogStreamOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsLogStream) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LogsLogStreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsLogStream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LogsLogStreamOutput) Properties() LogsLogStreamPropertiesResponseOutput {
	return o.ApplyT(func(v *LogsLogStream) LogsLogStreamPropertiesResponseOutput { return v.Properties }).(LogsLogStreamPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LogsLogStreamOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LogsLogStream) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LogsLogStreamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LogsLogStream) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LogsLogStreamOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsLogStream) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LogsLogStreamOutput{})
}
