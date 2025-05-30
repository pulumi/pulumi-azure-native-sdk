// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the module identified by module name.
//
// Uses Azure REST API version 2023-11-01.
//
// Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupModule(ctx *pulumi.Context, args *LookupModuleArgs, opts ...pulumi.InvokeOption) (*LookupModuleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupModuleResult
	err := ctx.Invoke("azure-native:automation:getModule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModuleArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The module name.
	ModuleName string `pulumi:"moduleName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the module type.
type LookupModuleResult struct {
	// Gets the activity count of the module.
	ActivityCount *int `pulumi:"activityCount"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets the error info of the module.
	Error *ModuleErrorInfoResponse `pulumi:"error"`
	// Gets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Gets type of module, if its composite or not.
	IsComposite *bool `pulumi:"isComposite"`
	// Gets the isGlobal flag of the module.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets the provisioning state of the module.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the size in bytes of the module.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Gets the version of the module.
	Version *string `pulumi:"version"`
}

func LookupModuleOutput(ctx *pulumi.Context, args LookupModuleOutputArgs, opts ...pulumi.InvokeOption) LookupModuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupModuleResultOutput, error) {
			args := v.(LookupModuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:automation:getModule", args, LookupModuleResultOutput{}, options).(LookupModuleResultOutput), nil
		}).(LookupModuleResultOutput)
}

type LookupModuleOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The module name.
	ModuleName pulumi.StringInput `pulumi:"moduleName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupModuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModuleArgs)(nil)).Elem()
}

// Definition of the module type.
type LookupModuleResultOutput struct{ *pulumi.OutputState }

func (LookupModuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModuleResult)(nil)).Elem()
}

func (o LookupModuleResultOutput) ToLookupModuleResultOutput() LookupModuleResultOutput {
	return o
}

func (o LookupModuleResultOutput) ToLookupModuleResultOutputWithContext(ctx context.Context) LookupModuleResultOutput {
	return o
}

// Gets the activity count of the module.
func (o LookupModuleResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

// The Azure API version of the resource.
func (o LookupModuleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets the creation time.
func (o LookupModuleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o LookupModuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets the error info of the module.
func (o LookupModuleResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

// Gets the etag of the resource.
func (o LookupModuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupModuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets type of module, if its composite or not.
func (o LookupModuleResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

// Gets the isGlobal flag of the module.
func (o LookupModuleResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

// Gets the last modified time.
func (o LookupModuleResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The Azure Region where the resource lives
func (o LookupModuleResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupModuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the provisioning state of the module.
func (o LookupModuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets the size in bytes of the module.
func (o LookupModuleResultOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

// Resource tags.
func (o LookupModuleResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupModuleResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupModuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Type }).(pulumi.StringOutput)
}

// Gets the version of the module.
func (o LookupModuleResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModuleResultOutput{})
}
