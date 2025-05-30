// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package powerbi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an existing Power BI Workspace Collection.
//
// Uses Azure REST API version 2016-01-29.
func LookupWorkspaceCollection(ctx *pulumi.Context, args *LookupWorkspaceCollectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceCollectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceCollectionResult
	err := ctx.Invoke("azure-native:powerbi:getWorkspaceCollection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceCollectionArgs struct {
	// Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type LookupWorkspaceCollectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Resource id
	Id *string `pulumi:"id"`
	// Azure location
	Location *string `pulumi:"location"`
	// Workspace collection name
	Name *string `pulumi:"name"`
	// Properties
	Properties interface{}       `pulumi:"properties"`
	Sku        *AzureSkuResponse `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

func LookupWorkspaceCollectionOutput(ctx *pulumi.Context, args LookupWorkspaceCollectionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceCollectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceCollectionResultOutput, error) {
			args := v.(LookupWorkspaceCollectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:powerbi:getWorkspaceCollection", args, LookupWorkspaceCollectionResultOutput{}, options).(LookupWorkspaceCollectionResultOutput), nil
		}).(LookupWorkspaceCollectionResultOutput)
}

type LookupWorkspaceCollectionOutputArgs struct {
	// Azure resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName pulumi.StringInput `pulumi:"workspaceCollectionName"`
}

func (LookupWorkspaceCollectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceCollectionArgs)(nil)).Elem()
}

type LookupWorkspaceCollectionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceCollectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceCollectionResult)(nil)).Elem()
}

func (o LookupWorkspaceCollectionResultOutput) ToLookupWorkspaceCollectionResultOutput() LookupWorkspaceCollectionResultOutput {
	return o
}

func (o LookupWorkspaceCollectionResultOutput) ToLookupWorkspaceCollectionResultOutputWithContext(ctx context.Context) LookupWorkspaceCollectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupWorkspaceCollectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource id
func (o LookupWorkspaceCollectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Azure location
func (o LookupWorkspaceCollectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Workspace collection name
func (o LookupWorkspaceCollectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Properties
func (o LookupWorkspaceCollectionResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Sku() AzureSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *AzureSkuResponse { return v.Sku }).(AzureSkuResponsePtrOutput)
}

func (o LookupWorkspaceCollectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupWorkspaceCollectionResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceCollectionResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceCollectionResultOutput{})
}
