// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a single workbook by its resourceName.
func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20210308:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkbookArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
}

// An Application Insights workbook definition.
type LookupWorkbookResult struct {
	// Workbook category, as defined by the user at creation time.
	Category string `pulumi:"category"`
	// The description of the workbook.
	Description *string `pulumi:"description"`
	// The user-defined name (display name) of the workbook.
	DisplayName string `pulumi:"displayName"`
	// Resource etag
	Etag map[string]string `pulumi:"etag"`
	// Azure resource Id
	Id *string `pulumi:"id"`
	// Identity used for BYOS
	Identity *WorkbookManagedIdentityResponse `pulumi:"identity"`
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// The unique revision id for this workbook definition
	Revision *string `pulumi:"revision"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData string `pulumi:"serializedData"`
	// ResourceId for a source resource.
	SourceId *string `pulumi:"sourceId"`
	// BYOS Storage Account URI
	StorageUri *string `pulumi:"storageUri"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this workbook definition.
	TimeModified string `pulumi:"timeModified"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// Unique user id of the specific user that owns this workbook.
	UserId string `pulumi:"userId"`
	// Workbook version
	Version *string `pulumi:"version"`
}

func LookupWorkbookOutput(ctx *pulumi.Context, args LookupWorkbookOutputArgs, opts ...pulumi.InvokeOption) LookupWorkbookResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkbookResultOutput, error) {
			args := v.(LookupWorkbookArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:insights/v20210308:getWorkbook", args, LookupWorkbookResultOutput{}, options).(LookupWorkbookResultOutput), nil
		}).(LookupWorkbookResultOutput)
}

type LookupWorkbookOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWorkbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookArgs)(nil)).Elem()
}

// An Application Insights workbook definition.
type LookupWorkbookResultOutput struct{ *pulumi.OutputState }

func (LookupWorkbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookResult)(nil)).Elem()
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutput() LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutputWithContext(ctx context.Context) LookupWorkbookResultOutput {
	return o
}

// Workbook category, as defined by the user at creation time.
func (o LookupWorkbookResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Category }).(pulumi.StringOutput)
}

// The description of the workbook.
func (o LookupWorkbookResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The user-defined name (display name) of the workbook.
func (o LookupWorkbookResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource etag
func (o LookupWorkbookResultOutput) Etag() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookResult) map[string]string { return v.Etag }).(pulumi.StringMapOutput)
}

// Azure resource Id
func (o LookupWorkbookResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Identity used for BYOS
func (o LookupWorkbookResultOutput) Identity() WorkbookManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *WorkbookManagedIdentityResponse { return v.Identity }).(WorkbookManagedIdentityResponsePtrOutput)
}

// The kind of workbook. Choices are user and shared.
func (o LookupWorkbookResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location
func (o LookupWorkbookResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name
func (o LookupWorkbookResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The unique revision id for this workbook definition
func (o LookupWorkbookResultOutput) Revision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Revision }).(pulumi.StringPtrOutput)
}

// Configuration of this particular workbook. Configuration data is a string containing valid JSON
func (o LookupWorkbookResultOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.SerializedData }).(pulumi.StringOutput)
}

// ResourceId for a source resource.
func (o LookupWorkbookResultOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

// BYOS Storage Account URI
func (o LookupWorkbookResultOutput) StorageUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.StorageUri }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupWorkbookResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkbookResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags
func (o LookupWorkbookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time in UTC of the last modification that was made to this workbook definition.
func (o LookupWorkbookResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o LookupWorkbookResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Unique user id of the specific user that owns this workbook.
func (o LookupWorkbookResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.UserId }).(pulumi.StringOutput)
}

// Workbook version
func (o LookupWorkbookResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkbookResultOutput{})
}
