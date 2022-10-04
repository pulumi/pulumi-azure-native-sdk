// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220808

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the module type.
func LookupModule(ctx *pulumi.Context, args *LookupModuleArgs, opts ...pulumi.InvokeOption) (*LookupModuleResult, error) {
	var rv LookupModuleResult
	err := ctx.Invoke("azure-native:automation/v20220808:getModule", args, &rv, opts...)
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
	// Gets or sets the activity count of the module.
	ActivityCount *int `pulumi:"activityCount"`
	// Gets or sets the contentLink of the module.
	ContentLink *ContentLinkResponse `pulumi:"contentLink"`
	// Gets or sets the creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Gets or sets the error info of the module.
	Error *ModuleErrorInfoResponse `pulumi:"error"`
	// Gets or sets the etag of the resource.
	Etag *string `pulumi:"etag"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// Gets or sets type of module, if its composite or not.
	IsComposite *bool `pulumi:"isComposite"`
	// Gets or sets the isGlobal flag of the module.
	IsGlobal *bool `pulumi:"isGlobal"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Azure Region where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the provisioning state of the module.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets or sets the size in bytes of the module.
	SizeInBytes *float64 `pulumi:"sizeInBytes"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// Gets or sets the version of the module.
	Version *string `pulumi:"version"`
}

func LookupModuleOutput(ctx *pulumi.Context, args LookupModuleOutputArgs, opts ...pulumi.InvokeOption) LookupModuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModuleResult, error) {
			args := v.(LookupModuleArgs)
			r, err := LookupModule(ctx, &args, opts...)
			var s LookupModuleResult
			if r != nil {
				s = *r
			}
			return s, err
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

// Gets or sets the activity count of the module.
func (o LookupModuleResultOutput) ActivityCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *int { return v.ActivityCount }).(pulumi.IntPtrOutput)
}

// Gets or sets the contentLink of the module.
func (o LookupModuleResultOutput) ContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *ContentLinkResponse { return v.ContentLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets the creation time.
func (o LookupModuleResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o LookupModuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the error info of the module.
func (o LookupModuleResultOutput) Error() ModuleErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *ModuleErrorInfoResponse { return v.Error }).(ModuleErrorInfoResponsePtrOutput)
}

// Gets or sets the etag of the resource.
func (o LookupModuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupModuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets type of module, if its composite or not.
func (o LookupModuleResultOutput) IsComposite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *bool { return v.IsComposite }).(pulumi.BoolPtrOutput)
}

// Gets or sets the isGlobal flag of the module.
func (o LookupModuleResultOutput) IsGlobal() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *bool { return v.IsGlobal }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
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

// Gets or sets the provisioning state of the module.
func (o LookupModuleResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Gets or sets the size in bytes of the module.
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

// Gets or sets the version of the module.
func (o LookupModuleResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupModuleResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModuleResultOutput{})
}