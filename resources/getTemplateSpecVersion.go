// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Template Spec version from a specific Template Spec.
//
// Uses Azure REST API version 2022-02-01.
//
// Other available API versions: 2021-03-01-preview, 2021-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupTemplateSpecVersion(ctx *pulumi.Context, args *LookupTemplateSpecVersionArgs, opts ...pulumi.InvokeOption) (*LookupTemplateSpecVersionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTemplateSpecVersionResult
	err := ctx.Invoke("azure-native:resources:getTemplateSpecVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateSpecVersionArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Template Spec.
	TemplateSpecName string `pulumi:"templateSpecName"`
	// The version of the Template Spec.
	TemplateSpecVersion string `pulumi:"templateSpecVersion"`
}

// Template Spec Version object.
type LookupTemplateSpecVersionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Template Spec version description.
	Description *string `pulumi:"description"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// An array of linked template artifacts.
	LinkedTemplates []LinkedTemplateArtifactResponse `pulumi:"linkedTemplates"`
	// The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location string `pulumi:"location"`
	// The main Azure Resource Manager template content.
	MainTemplate interface{} `pulumi:"mainTemplate"`
	// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata interface{} `pulumi:"metadata"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Type of this resource.
	Type string `pulumi:"type"`
	// The Azure Resource Manager template UI definition content.
	UiFormDefinition interface{} `pulumi:"uiFormDefinition"`
}

func LookupTemplateSpecVersionOutput(ctx *pulumi.Context, args LookupTemplateSpecVersionOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateSpecVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTemplateSpecVersionResultOutput, error) {
			args := v.(LookupTemplateSpecVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:resources:getTemplateSpecVersion", args, LookupTemplateSpecVersionResultOutput{}, options).(LookupTemplateSpecVersionResultOutput), nil
		}).(LookupTemplateSpecVersionResultOutput)
}

type LookupTemplateSpecVersionOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Template Spec.
	TemplateSpecName pulumi.StringInput `pulumi:"templateSpecName"`
	// The version of the Template Spec.
	TemplateSpecVersion pulumi.StringInput `pulumi:"templateSpecVersion"`
}

func (LookupTemplateSpecVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateSpecVersionArgs)(nil)).Elem()
}

// Template Spec Version object.
type LookupTemplateSpecVersionResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateSpecVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateSpecVersionResult)(nil)).Elem()
}

func (o LookupTemplateSpecVersionResultOutput) ToLookupTemplateSpecVersionResultOutput() LookupTemplateSpecVersionResultOutput {
	return o
}

func (o LookupTemplateSpecVersionResultOutput) ToLookupTemplateSpecVersionResultOutputWithContext(ctx context.Context) LookupTemplateSpecVersionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupTemplateSpecVersionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Template Spec version description.
func (o LookupTemplateSpecVersionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupTemplateSpecVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

// An array of linked template artifacts.
func (o LookupTemplateSpecVersionResultOutput) LinkedTemplates() LinkedTemplateArtifactResponseArrayOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) []LinkedTemplateArtifactResponse { return v.LinkedTemplates }).(LinkedTemplateArtifactResponseArrayOutput)
}

// The location of the Template Spec Version. It must match the location of the parent Template Spec.
func (o LookupTemplateSpecVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

// The main Azure Resource Manager template content.
func (o LookupTemplateSpecVersionResultOutput) MainTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.MainTemplate }).(pulumi.AnyOutput)
}

// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
func (o LookupTemplateSpecVersionResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// Name of this resource.
func (o LookupTemplateSpecVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupTemplateSpecVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupTemplateSpecVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of this resource.
func (o LookupTemplateSpecVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

// The Azure Resource Manager template UI definition content.
func (o LookupTemplateSpecVersionResultOutput) UiFormDefinition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateSpecVersionResult) interface{} { return v.UiFormDefinition }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateSpecVersionResultOutput{})
}
