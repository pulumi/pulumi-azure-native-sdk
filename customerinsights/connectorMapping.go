// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The connector mapping resource format.
// API Version: 2017-04-26.
type ConnectorMapping struct {
	pulumi.CustomResourceState

	// The connector mapping name
	ConnectorMappingName pulumi.StringOutput `pulumi:"connectorMappingName"`
	// The connector name.
	ConnectorName pulumi.StringOutput `pulumi:"connectorName"`
	// Type of connector.
	ConnectorType pulumi.StringPtrOutput `pulumi:"connectorType"`
	// The created time.
	Created pulumi.StringOutput `pulumi:"created"`
	// The DataFormat ID.
	DataFormatId pulumi.StringOutput `pulumi:"dataFormatId"`
	// The description of the connector mapping.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name for the connector mapping.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Defines which entity type the file should map to.
	EntityType pulumi.StringOutput `pulumi:"entityType"`
	// The mapping entity name.
	EntityTypeName pulumi.StringOutput `pulumi:"entityTypeName"`
	// The last modified time.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The properties of the mapping.
	MappingProperties ConnectorMappingPropertiesResponseOutput `pulumi:"mappingProperties"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The next run time based on customer's settings.
	NextRunTime pulumi.StringOutput `pulumi:"nextRunTime"`
	// The RunId.
	RunId pulumi.StringOutput `pulumi:"runId"`
	// State of connector mapping.
	State pulumi.StringOutput `pulumi:"state"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectorMapping registers a new resource with the given unique name, arguments, and options.
func NewConnectorMapping(ctx *pulumi.Context,
	name string, args *ConnectorMappingArgs, opts ...pulumi.ResourceOption) (*ConnectorMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorName'")
	}
	if args.EntityType == nil {
		return nil, errors.New("invalid value for required argument 'EntityType'")
	}
	if args.EntityTypeName == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeName'")
	}
	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.MappingProperties == nil {
		return nil, errors.New("invalid value for required argument 'MappingProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:ConnectorMapping"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:ConnectorMapping"),
		},
	})
	opts = append(opts, aliases)
	var resource ConnectorMapping
	err := ctx.RegisterResource("azure-native:customerinsights:ConnectorMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectorMapping gets an existing ConnectorMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectorMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorMappingState, opts ...pulumi.ResourceOption) (*ConnectorMapping, error) {
	var resource ConnectorMapping
	err := ctx.ReadResource("azure-native:customerinsights:ConnectorMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectorMapping resources.
type connectorMappingState struct {
}

type ConnectorMappingState struct {
}

func (ConnectorMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorMappingState)(nil)).Elem()
}

type connectorMappingArgs struct {
	// The name of the connector.
	ConnectorName string `pulumi:"connectorName"`
	// Type of connector.
	ConnectorType *string `pulumi:"connectorType"`
	// The description of the connector mapping.
	Description *string `pulumi:"description"`
	// Display name for the connector mapping.
	DisplayName *string `pulumi:"displayName"`
	// Defines which entity type the file should map to.
	EntityType EntityTypes `pulumi:"entityType"`
	// The mapping entity name.
	EntityTypeName string `pulumi:"entityTypeName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the connector mapping.
	MappingName *string `pulumi:"mappingName"`
	// The properties of the mapping.
	MappingProperties ConnectorMappingProperties `pulumi:"mappingProperties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConnectorMapping resource.
type ConnectorMappingArgs struct {
	// The name of the connector.
	ConnectorName pulumi.StringInput
	// Type of connector.
	ConnectorType pulumi.StringPtrInput
	// The description of the connector mapping.
	Description pulumi.StringPtrInput
	// Display name for the connector mapping.
	DisplayName pulumi.StringPtrInput
	// Defines which entity type the file should map to.
	EntityType EntityTypesInput
	// The mapping entity name.
	EntityTypeName pulumi.StringInput
	// The name of the hub.
	HubName pulumi.StringInput
	// The name of the connector mapping.
	MappingName pulumi.StringPtrInput
	// The properties of the mapping.
	MappingProperties ConnectorMappingPropertiesInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ConnectorMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorMappingArgs)(nil)).Elem()
}

type ConnectorMappingInput interface {
	pulumi.Input

	ToConnectorMappingOutput() ConnectorMappingOutput
	ToConnectorMappingOutputWithContext(ctx context.Context) ConnectorMappingOutput
}

func (*ConnectorMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMapping)(nil)).Elem()
}

func (i *ConnectorMapping) ToConnectorMappingOutput() ConnectorMappingOutput {
	return i.ToConnectorMappingOutputWithContext(context.Background())
}

func (i *ConnectorMapping) ToConnectorMappingOutputWithContext(ctx context.Context) ConnectorMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectorMappingOutput)
}

type ConnectorMappingOutput struct{ *pulumi.OutputState }

func (ConnectorMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorMapping)(nil)).Elem()
}

func (o ConnectorMappingOutput) ToConnectorMappingOutput() ConnectorMappingOutput {
	return o
}

func (o ConnectorMappingOutput) ToConnectorMappingOutputWithContext(ctx context.Context) ConnectorMappingOutput {
	return o
}

// The connector mapping name
func (o ConnectorMappingOutput) ConnectorMappingName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.ConnectorMappingName }).(pulumi.StringOutput)
}

// The connector name.
func (o ConnectorMappingOutput) ConnectorName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.ConnectorName }).(pulumi.StringOutput)
}

// Type of connector.
func (o ConnectorMappingOutput) ConnectorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringPtrOutput { return v.ConnectorType }).(pulumi.StringPtrOutput)
}

// The created time.
func (o ConnectorMappingOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The DataFormat ID.
func (o ConnectorMappingOutput) DataFormatId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.DataFormatId }).(pulumi.StringOutput)
}

// The description of the connector mapping.
func (o ConnectorMappingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Display name for the connector mapping.
func (o ConnectorMappingOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Defines which entity type the file should map to.
func (o ConnectorMappingOutput) EntityType() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.EntityType }).(pulumi.StringOutput)
}

// The mapping entity name.
func (o ConnectorMappingOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.EntityTypeName }).(pulumi.StringOutput)
}

// The last modified time.
func (o ConnectorMappingOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The properties of the mapping.
func (o ConnectorMappingOutput) MappingProperties() ConnectorMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectorMapping) ConnectorMappingPropertiesResponseOutput { return v.MappingProperties }).(ConnectorMappingPropertiesResponseOutput)
}

// Resource name.
func (o ConnectorMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The next run time based on customer's settings.
func (o ConnectorMappingOutput) NextRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.NextRunTime }).(pulumi.StringOutput)
}

// The RunId.
func (o ConnectorMappingOutput) RunId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.RunId }).(pulumi.StringOutput)
}

// State of connector mapping.
func (o ConnectorMappingOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The hub name.
func (o ConnectorMappingOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o ConnectorMappingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectorMapping) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectorMappingOutput{})
}