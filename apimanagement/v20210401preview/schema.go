// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema Contract details.
type Schema struct {
	pulumi.CustomResourceState

	// Free-form schema entity description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Schema Type. Immutable.
	SchemaType pulumi.StringOutput `pulumi:"schemaType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Json-encoded string for non json-based schema.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaType == nil {
		return nil, errors.New("invalid value for required argument 'SchemaType'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Schema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Schema"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Schema"),
		},
	})
	opts = append(opts, aliases)
	var resource Schema
	err := ctx.RegisterResource("azure-native:apimanagement/v20210401preview:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("azure-native:apimanagement/v20210401preview:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
}

type SchemaState struct {
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	// Free-form schema entity description.
	Description *string `pulumi:"description"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId *string `pulumi:"schemaId"`
	// Schema Type. Immutable.
	SchemaType string `pulumi:"schemaType"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Json-encoded string for non json-based schema.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// Free-form schema entity description.
	Description pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Schema id identifier. Must be unique in the current API Management service instance.
	SchemaId pulumi.StringPtrInput
	// Schema Type. Immutable.
	SchemaType pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Json-encoded string for non json-based schema.
	Value pulumi.StringPtrInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

// Free-form schema entity description.
func (o SchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Schema Type. Immutable.
func (o SchemaOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.SchemaType }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Json-encoded string for non json-based schema.
func (o SchemaOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SchemaOutput{})
}