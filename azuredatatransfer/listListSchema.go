// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredatatransfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the schemas for the specified connection in a pipeline.
//
// Uses Azure REST API version 2024-09-27.
//
// Other available API versions: 2023-10-11-preview, 2024-01-25, 2024-05-07, 2024-09-11, 2025-03-01-preview, 2025-04-11-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azuredatatransfer [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListListSchema(ctx *pulumi.Context, args *ListListSchemaArgs, opts ...pulumi.InvokeOption) (*ListListSchemaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListSchemaResult
	err := ctx.Invoke("azure-native:azuredatatransfer:listListSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListSchemaArgs struct {
	// Connection ID associated with this schema
	ConnectionId *string `pulumi:"connectionId"`
	// Content of the schema
	Content *string `pulumi:"content"`
	// The direction of the schema.
	Direction *string `pulumi:"direction"`
	// ID associated with this schema
	Id *string `pulumi:"id"`
	// Name of the schema
	Name *string `pulumi:"name"`
	// The name for the pipeline that is to be requested.
	PipelineName string `pulumi:"pipelineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Schema Type
	SchemaType *string `pulumi:"schemaType"`
	// Uri containing SAS token for the zipped schema
	SchemaUri *string `pulumi:"schemaUri"`
	// Status of the schema
	Status *string `pulumi:"status"`
}

// The schemas list result.
type ListListSchemaResult struct {
	// Schemas array.
	Value []SchemaResponse `pulumi:"value"`
}

func ListListSchemaOutput(ctx *pulumi.Context, args ListListSchemaOutputArgs, opts ...pulumi.InvokeOption) ListListSchemaResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListListSchemaResultOutput, error) {
			args := v.(ListListSchemaArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azuredatatransfer:listListSchema", args, ListListSchemaResultOutput{}, options).(ListListSchemaResultOutput), nil
		}).(ListListSchemaResultOutput)
}

type ListListSchemaOutputArgs struct {
	// Connection ID associated with this schema
	ConnectionId pulumi.StringPtrInput `pulumi:"connectionId"`
	// Content of the schema
	Content pulumi.StringPtrInput `pulumi:"content"`
	// The direction of the schema.
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	// ID associated with this schema
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the schema
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The name for the pipeline that is to be requested.
	PipelineName pulumi.StringInput `pulumi:"pipelineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Schema Type
	SchemaType pulumi.StringPtrInput `pulumi:"schemaType"`
	// Uri containing SAS token for the zipped schema
	SchemaUri pulumi.StringPtrInput `pulumi:"schemaUri"`
	// Status of the schema
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (ListListSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListSchemaArgs)(nil)).Elem()
}

// The schemas list result.
type ListListSchemaResultOutput struct{ *pulumi.OutputState }

func (ListListSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListSchemaResult)(nil)).Elem()
}

func (o ListListSchemaResultOutput) ToListListSchemaResultOutput() ListListSchemaResultOutput {
	return o
}

func (o ListListSchemaResultOutput) ToListListSchemaResultOutputWithContext(ctx context.Context) ListListSchemaResultOutput {
	return o
}

// Schemas array.
func (o ListListSchemaResultOutput) Value() SchemaResponseArrayOutput {
	return o.ApplyT(func(v ListListSchemaResult) []SchemaResponse { return v.Value }).(SchemaResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListSchemaResultOutput{})
}
