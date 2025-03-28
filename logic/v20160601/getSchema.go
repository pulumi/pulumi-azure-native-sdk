// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an integration account schema.
func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSchemaResult
	err := ctx.Invoke("azure-native:logic/v20160601:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaArgs struct {
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The integration account schema name.
	SchemaName string `pulumi:"schemaName"`
}

// The integration account schema.
type LookupSchemaResult struct {
	// The changed time.
	ChangedTime string `pulumi:"changedTime"`
	// The content.
	Content *string `pulumi:"content"`
	// The content link.
	ContentLink ContentLinkResponse `pulumi:"contentLink"`
	// The content type.
	ContentType *string `pulumi:"contentType"`
	// The created time.
	CreatedTime string `pulumi:"createdTime"`
	// The document name.
	DocumentName *string `pulumi:"documentName"`
	// The file name.
	FileName *string `pulumi:"fileName"`
	// The resource id.
	Id string `pulumi:"id"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// The schema type.
	SchemaType string `pulumi:"schemaType"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The target namespace of the schema.
	TargetNamespace *string `pulumi:"targetNamespace"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

func LookupSchemaOutput(ctx *pulumi.Context, args LookupSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSchemaResultOutput, error) {
			args := v.(LookupSchemaArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:logic/v20160601:getSchema", args, LookupSchemaResultOutput{}, options).(LookupSchemaResultOutput), nil
		}).(LookupSchemaResultOutput)
}

type LookupSchemaOutputArgs struct {
	// The integration account name.
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The integration account schema name.
	SchemaName pulumi.StringInput `pulumi:"schemaName"`
}

func (LookupSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaArgs)(nil)).Elem()
}

// The integration account schema.
type LookupSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaResult)(nil)).Elem()
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutput() LookupSchemaResultOutput {
	return o
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutputWithContext(ctx context.Context) LookupSchemaResultOutput {
	return o
}

// The changed time.
func (o LookupSchemaResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

// The content.
func (o LookupSchemaResultOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// The content link.
func (o LookupSchemaResultOutput) ContentLink() ContentLinkResponseOutput {
	return o.ApplyT(func(v LookupSchemaResult) ContentLinkResponse { return v.ContentLink }).(ContentLinkResponseOutput)
}

// The content type.
func (o LookupSchemaResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// The created time.
func (o LookupSchemaResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// The document name.
func (o LookupSchemaResultOutput) DocumentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.DocumentName }).(pulumi.StringPtrOutput)
}

// The file name.
func (o LookupSchemaResultOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.FileName }).(pulumi.StringPtrOutput)
}

// The resource id.
func (o LookupSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource location.
func (o LookupSchemaResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o LookupSchemaResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSchemaResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// Gets the resource name.
func (o LookupSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

// The schema type.
func (o LookupSchemaResultOutput) SchemaType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.SchemaType }).(pulumi.StringOutput)
}

// The resource tags.
func (o LookupSchemaResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSchemaResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The target namespace of the schema.
func (o LookupSchemaResultOutput) TargetNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchemaResult) *string { return v.TargetNamespace }).(pulumi.StringPtrOutput)
}

// Gets the resource type.
func (o LookupSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaResultOutput{})
}
