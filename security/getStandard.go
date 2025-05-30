// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a specific security standard for the requested scope
//
// Uses Azure REST API version 2021-08-01-preview.
func LookupStandard(ctx *pulumi.Context, args *LookupStandardArgs, opts ...pulumi.InvokeOption) (*LookupStandardResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupStandardResult
	err := ctx.Invoke("azure-native:security:getStandard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStandardArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Security Standard key - unique key for the standard type
	StandardId string `pulumi:"standardId"`
}

// Security Standard on a resource
type LookupStandardResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// category of the standard provided
	Category *string `pulumi:"category"`
	// List of component objects containing component unique keys (such as assessment keys) to apply to standard scope.  Currently only supports assessment keys.
	Components []StandardComponentPropertiesResponse `pulumi:"components"`
	// description of the standard
	Description *string `pulumi:"description"`
	// display name of the standard, equivalent to the standardId
	DisplayName *string `pulumi:"displayName"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag *string `pulumi:"etag"`
	// Resource Id
	Id string `pulumi:"id"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// standard type (Custom or BuiltIn only currently)
	StandardType string `pulumi:"standardType"`
	// List of all standard supported clouds.
	SupportedClouds []string `pulumi:"supportedClouds"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupStandardOutput(ctx *pulumi.Context, args LookupStandardOutputArgs, opts ...pulumi.InvokeOption) LookupStandardResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStandardResultOutput, error) {
			args := v.(LookupStandardArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:security:getStandard", args, LookupStandardResultOutput{}, options).(LookupStandardResultOutput), nil
		}).(LookupStandardResultOutput)
}

type LookupStandardOutputArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The Security Standard key - unique key for the standard type
	StandardId pulumi.StringInput `pulumi:"standardId"`
}

func (LookupStandardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardArgs)(nil)).Elem()
}

// Security Standard on a resource
type LookupStandardResultOutput struct{ *pulumi.OutputState }

func (LookupStandardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStandardResult)(nil)).Elem()
}

func (o LookupStandardResultOutput) ToLookupStandardResultOutput() LookupStandardResultOutput {
	return o
}

func (o LookupStandardResultOutput) ToLookupStandardResultOutputWithContext(ctx context.Context) LookupStandardResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupStandardResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// category of the standard provided
func (o LookupStandardResultOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Category }).(pulumi.StringPtrOutput)
}

// List of component objects containing component unique keys (such as assessment keys) to apply to standard scope.  Currently only supports assessment keys.
func (o LookupStandardResultOutput) Components() StandardComponentPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []StandardComponentPropertiesResponse { return v.Components }).(StandardComponentPropertiesResponseArrayOutput)
}

// description of the standard
func (o LookupStandardResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// display name of the standard, equivalent to the standardId
func (o LookupStandardResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o LookupStandardResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupStandardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the resource
func (o LookupStandardResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o LookupStandardResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStandardResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupStandardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Name }).(pulumi.StringOutput)
}

// standard type (Custom or BuiltIn only currently)
func (o LookupStandardResultOutput) StandardType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.StandardType }).(pulumi.StringOutput)
}

// List of all standard supported clouds.
func (o LookupStandardResultOutput) SupportedClouds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStandardResult) []string { return v.SupportedClouds }).(pulumi.StringArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupStandardResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStandardResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of key value pairs that describe the resource.
func (o LookupStandardResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStandardResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupStandardResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStandardResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStandardResultOutput{})
}
