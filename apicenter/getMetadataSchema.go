// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apicenter

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns details of the metadata schema.
//
// Uses Azure REST API version 2024-03-15-preview.
//
// Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupMetadataSchema(ctx *pulumi.Context, args *LookupMetadataSchemaArgs, opts ...pulumi.InvokeOption) (*LookupMetadataSchemaResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMetadataSchemaResult
	err := ctx.Invoke("azure-native:apicenter:getMetadataSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetadataSchemaArgs struct {
	// The name of the metadata schema.
	MetadataSchemaName string `pulumi:"metadataSchemaName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
}

// Metadata schema entity. Used to define metadata for the entities in API catalog.
type LookupMetadataSchemaResult struct {
	// The assignees
	AssignedTo []MetadataAssignmentResponse `pulumi:"assignedTo"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The schema defining the type.
	Schema string `pulumi:"schema"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMetadataSchemaOutput(ctx *pulumi.Context, args LookupMetadataSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupMetadataSchemaResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMetadataSchemaResultOutput, error) {
			args := v.(LookupMetadataSchemaArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apicenter:getMetadataSchema", args, LookupMetadataSchemaResultOutput{}, options).(LookupMetadataSchemaResultOutput), nil
		}).(LookupMetadataSchemaResultOutput)
}

type LookupMetadataSchemaOutputArgs struct {
	// The name of the metadata schema.
	MetadataSchemaName pulumi.StringInput `pulumi:"metadataSchemaName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupMetadataSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataSchemaArgs)(nil)).Elem()
}

// Metadata schema entity. Used to define metadata for the entities in API catalog.
type LookupMetadataSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupMetadataSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataSchemaResult)(nil)).Elem()
}

func (o LookupMetadataSchemaResultOutput) ToLookupMetadataSchemaResultOutput() LookupMetadataSchemaResultOutput {
	return o
}

func (o LookupMetadataSchemaResultOutput) ToLookupMetadataSchemaResultOutputWithContext(ctx context.Context) LookupMetadataSchemaResultOutput {
	return o
}

// The assignees
func (o LookupMetadataSchemaResultOutput) AssignedTo() MetadataAssignmentResponseArrayOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) []MetadataAssignmentResponse { return v.AssignedTo }).(MetadataAssignmentResponseArrayOutput)
}

// The Azure API version of the resource.
func (o LookupMetadataSchemaResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMetadataSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMetadataSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

// The schema defining the type.
func (o LookupMetadataSchemaResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) string { return v.Schema }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMetadataSchemaResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMetadataSchemaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataSchemaResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetadataSchemaResultOutput{})
}
