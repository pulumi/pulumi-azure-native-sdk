// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200215

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing an event hub data connection.
type EventHubDataConnection struct {
	pulumi.CustomResourceState

	// The event hub messages compression type
	Compression pulumi.StringPtrOutput `pulumi:"compression"`
	// The event hub consumer group.
	ConsumerGroup pulumi.StringOutput `pulumi:"consumerGroup"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId pulumi.StringOutput `pulumi:"eventHubResourceId"`
	// System properties of the event hub
	EventSystemProperties pulumi.StringArrayOutput `pulumi:"eventSystemProperties"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventHub'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventHubDataConnection registers a new resource with the given unique name, arguments, and options.
func NewEventHubDataConnection(ctx *pulumi.Context,
	name string, args *EventHubDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventHubDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ConsumerGroup == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerGroup'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.EventHubResourceId == nil {
		return nil, errors.New("invalid value for required argument 'EventHubResourceId'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Kind = pulumi.String("EventHub")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:EventHubDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:EventHubDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHubDataConnection
	err := ctx.RegisterResource("azure-native:kusto/v20200215:EventHubDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubDataConnection gets an existing EventHubDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubDataConnectionState, opts ...pulumi.ResourceOption) (*EventHubDataConnection, error) {
	var resource EventHubDataConnection
	err := ctx.ReadResource("azure-native:kusto/v20200215:EventHubDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubDataConnection resources.
type eventHubDataConnectionState struct {
}

type EventHubDataConnectionState struct {
}

func (EventHubDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubDataConnectionState)(nil)).Elem()
}

type eventHubDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The event hub messages compression type
	Compression *string `pulumi:"compression"`
	// The event hub consumer group.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// The name of the data connection.
	DataConnectionName *string `pulumi:"dataConnectionName"`
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat *string `pulumi:"dataFormat"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId string `pulumi:"eventHubResourceId"`
	// System properties of the event hub
	EventSystemProperties []string `pulumi:"eventSystemProperties"`
	// Kind of the endpoint for the data connection
	// Expected value is 'EventHub'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a EventHubDataConnection resource.
type EventHubDataConnectionArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The event hub messages compression type
	Compression pulumi.StringPtrInput
	// The event hub consumer group.
	ConsumerGroup pulumi.StringInput
	// The name of the data connection.
	DataConnectionName pulumi.StringPtrInput
	// The data format of the message. Optionally the data format can be added to each message.
	DataFormat pulumi.StringPtrInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput
	// The resource ID of the event hub to be used to create a data connection.
	EventHubResourceId pulumi.StringInput
	// System properties of the event hub
	EventSystemProperties pulumi.StringArrayInput
	// Kind of the endpoint for the data connection
	// Expected value is 'EventHub'.
	Kind pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
	MappingRuleName pulumi.StringPtrInput
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput
	// The table where the data should be ingested. Optionally the table information can be added to each message.
	TableName pulumi.StringPtrInput
}

func (EventHubDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubDataConnectionArgs)(nil)).Elem()
}

type EventHubDataConnectionInput interface {
	pulumi.Input

	ToEventHubDataConnectionOutput() EventHubDataConnectionOutput
	ToEventHubDataConnectionOutputWithContext(ctx context.Context) EventHubDataConnectionOutput
}

func (*EventHubDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubDataConnection)(nil)).Elem()
}

func (i *EventHubDataConnection) ToEventHubDataConnectionOutput() EventHubDataConnectionOutput {
	return i.ToEventHubDataConnectionOutputWithContext(context.Background())
}

func (i *EventHubDataConnection) ToEventHubDataConnectionOutputWithContext(ctx context.Context) EventHubDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubDataConnectionOutput)
}

type EventHubDataConnectionOutput struct{ *pulumi.OutputState }

func (EventHubDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubDataConnection)(nil)).Elem()
}

func (o EventHubDataConnectionOutput) ToEventHubDataConnectionOutput() EventHubDataConnectionOutput {
	return o
}

func (o EventHubDataConnectionOutput) ToEventHubDataConnectionOutputWithContext(ctx context.Context) EventHubDataConnectionOutput {
	return o
}

// The event hub messages compression type
func (o EventHubDataConnectionOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.Compression }).(pulumi.StringPtrOutput)
}

// The event hub consumer group.
func (o EventHubDataConnectionOutput) ConsumerGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.ConsumerGroup }).(pulumi.StringOutput)
}

// The data format of the message. Optionally the data format can be added to each message.
func (o EventHubDataConnectionOutput) DataFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.DataFormat }).(pulumi.StringPtrOutput)
}

// The resource ID of the event hub to be used to create a data connection.
func (o EventHubDataConnectionOutput) EventHubResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.EventHubResourceId }).(pulumi.StringOutput)
}

// System properties of the event hub
func (o EventHubDataConnectionOutput) EventSystemProperties() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringArrayOutput { return v.EventSystemProperties }).(pulumi.StringArrayOutput)
}

// Kind of the endpoint for the data connection
// Expected value is 'EventHub'.
func (o EventHubDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o EventHubDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The mapping rule to be used to ingest the data. Optionally the mapping information can be added to each message.
func (o EventHubDataConnectionOutput) MappingRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.MappingRuleName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o EventHubDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The table where the data should be ingested. Optionally the table information can be added to each message.
func (o EventHubDataConnectionOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EventHubDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventHubDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubDataConnectionOutput{})
}