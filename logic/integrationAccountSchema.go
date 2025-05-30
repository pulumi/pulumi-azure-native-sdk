// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration account schema.
//
// Uses Azure REST API version 2019-05-01. In version 2.x of the Azure Native provider, it used API version 2019-05-01.
//
// Other available API versions: 2015-08-01-preview, 2018-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native logic [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type IntegrationAccountSchema struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The content.
	Content pulumi.StringPtrOutput `pulumi:"content"`
	// The content link.
	ContentLink ContentLinkResponseOutput `pulumi:"contentLink"`
	// The content type.
	ContentType pulumi.StringPtrOutput `pulumi:"contentType"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The document name.
	DocumentName pulumi.StringPtrOutput `pulumi:"documentName"`
	// The file name.
	FileName pulumi.StringPtrOutput `pulumi:"fileName"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schema type.
	SchemaType pulumi.StringOutput `pulumi:"schemaType"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The target namespace of the schema.
	TargetNamespace pulumi.StringPtrOutput `pulumi:"targetNamespace"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountSchema registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountSchema(ctx *pulumi.Context,
	name string, args *IntegrationAccountSchemaArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountSchema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SchemaType == nil {
		return nil, errors.New("invalid value for required argument 'SchemaType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Schema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountSchema"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountSchema"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationAccountSchema
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccountSchema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountSchema gets an existing IntegrationAccountSchema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountSchemaState, opts ...pulumi.ResourceOption) (*IntegrationAccountSchema, error) {
	var resource IntegrationAccountSchema
	err := ctx.ReadResource("azure-native:logic:IntegrationAccountSchema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountSchema resources.
type integrationAccountSchemaState struct {
}

type IntegrationAccountSchemaState struct {
}

func (IntegrationAccountSchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSchemaState)(nil)).Elem()
}

type integrationAccountSchemaArgs struct {
	// The content.
	Content *string `pulumi:"content"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The document name.
	DocumentName *string `pulumi:"documentName"`
	// The file name.
	FileName *string `pulumi:"fileName"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The integration account schema name.
	SchemaName *string `pulumi:"schemaName"`
	// The schema type.
	SchemaType string `pulumi:"schemaType"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The target namespace of the schema.
	TargetNamespace *string `pulumi:"targetNamespace"`
}

// The set of arguments for constructing a IntegrationAccountSchema resource.
type IntegrationAccountSchemaArgs struct {
	// The content.
	Content pulumi.StringPtrInput
	// The content type.
	ContentType pulumi.StringPtrInput
	// The document name.
	DocumentName pulumi.StringPtrInput
	// The file name.
	FileName pulumi.StringPtrInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The integration account schema name.
	SchemaName pulumi.StringPtrInput
	// The schema type.
	SchemaType pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The target namespace of the schema.
	TargetNamespace pulumi.StringPtrInput
}

func (IntegrationAccountSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountSchemaArgs)(nil)).Elem()
}

type IntegrationAccountSchemaInput interface {
	pulumi.Input

	ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput
	ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput
}

func (*IntegrationAccountSchema) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSchema)(nil)).Elem()
}

func (i *IntegrationAccountSchema) ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput {
	return i.ToIntegrationAccountSchemaOutputWithContext(context.Background())
}

func (i *IntegrationAccountSchema) ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountSchemaOutput)
}

type IntegrationAccountSchemaOutput struct{ *pulumi.OutputState }

func (IntegrationAccountSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountSchema)(nil)).Elem()
}

func (o IntegrationAccountSchemaOutput) ToIntegrationAccountSchemaOutput() IntegrationAccountSchemaOutput {
	return o
}

func (o IntegrationAccountSchemaOutput) ToIntegrationAccountSchemaOutputWithContext(ctx context.Context) IntegrationAccountSchemaOutput {
	return o
}

// The Azure API version of the resource.
func (o IntegrationAccountSchemaOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The changed time.
func (o IntegrationAccountSchemaOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// The content.
func (o IntegrationAccountSchemaOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.Content }).(pulumi.StringPtrOutput)
}

// The content link.
func (o IntegrationAccountSchemaOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) ContentLinkResponseOutput { return v.ContentLink }).(ContentLinkResponseOutput)
}

// The content type.
func (o IntegrationAccountSchemaOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The created time.
func (o IntegrationAccountSchemaOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The document name.
func (o IntegrationAccountSchemaOutput) DocumentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.DocumentName }).(pulumi.StringPtrOutput)
}

// The file name.
func (o IntegrationAccountSchemaOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.FileName }).(pulumi.StringPtrOutput)
}

// The resource location.
func (o IntegrationAccountSchemaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o IntegrationAccountSchemaOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Gets the resource name.
func (o IntegrationAccountSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The schema type.
func (o IntegrationAccountSchemaOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.SchemaType }).(pulumi.StringOutput)
}

// The resource tags.
func (o IntegrationAccountSchemaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The target namespace of the schema.
func (o IntegrationAccountSchemaOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringPtrOutput { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

// Gets the resource type.
func (o IntegrationAccountSchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountSchema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountSchemaOutput{})
}
