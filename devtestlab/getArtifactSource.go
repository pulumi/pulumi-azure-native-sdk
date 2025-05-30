// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get artifact source.
//
// Uses Azure REST API version 2018-09-15.
func LookupArtifactSource(ctx *pulumi.Context, args *LookupArtifactSourceArgs, opts ...pulumi.InvokeOption) (*LookupArtifactSourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupArtifactSourceResult
	err := ctx.Invoke("azure-native:devtestlab:getArtifactSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactSourceArgs struct {
	// Specify the $expand query. Example: 'properties($select=displayName)'
	Expand *string `pulumi:"expand"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the artifact source.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Properties of an artifact source.
type LookupArtifactSourceResult struct {
	// The folder containing Azure Resource Manager templates.
	ArmTemplateFolderPath *string `pulumi:"armTemplateFolderPath"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The artifact source's branch reference.
	BranchRef *string `pulumi:"branchRef"`
	// The artifact source's creation date.
	CreatedDate string `pulumi:"createdDate"`
	// The artifact source's display name.
	DisplayName *string `pulumi:"displayName"`
	// The folder containing artifacts.
	FolderPath *string `pulumi:"folderPath"`
	// The identifier of the resource.
	Id string `pulumi:"id"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The security token to authenticate to the artifact source.
	SecurityToken *string `pulumi:"securityToken"`
	// The artifact source's type.
	SourceType *string `pulumi:"sourceType"`
	// Indicates if the artifact source is enabled (values: Enabled, Disabled).
	Status *string `pulumi:"status"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
	// The artifact source's URI.
	Uri *string `pulumi:"uri"`
}

func LookupArtifactSourceOutput(ctx *pulumi.Context, args LookupArtifactSourceOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactSourceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupArtifactSourceResultOutput, error) {
			args := v.(LookupArtifactSourceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devtestlab:getArtifactSource", args, LookupArtifactSourceResultOutput{}, options).(LookupArtifactSourceResultOutput), nil
		}).(LookupArtifactSourceResultOutput)
}

type LookupArtifactSourceOutputArgs struct {
	// Specify the $expand query. Example: 'properties($select=displayName)'
	Expand pulumi.StringPtrInput `pulumi:"expand"`
	// The name of the lab.
	LabName pulumi.StringInput `pulumi:"labName"`
	// The name of the artifact source.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArtifactSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceArgs)(nil)).Elem()
}

// Properties of an artifact source.
type LookupArtifactSourceResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactSourceResult)(nil)).Elem()
}

func (o LookupArtifactSourceResultOutput) ToLookupArtifactSourceResultOutput() LookupArtifactSourceResultOutput {
	return o
}

func (o LookupArtifactSourceResultOutput) ToLookupArtifactSourceResultOutputWithContext(ctx context.Context) LookupArtifactSourceResultOutput {
	return o
}

// The folder containing Azure Resource Manager templates.
func (o LookupArtifactSourceResultOutput) ArmTemplateFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.ArmTemplateFolderPath }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupArtifactSourceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The artifact source's branch reference.
func (o LookupArtifactSourceResultOutput) BranchRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.BranchRef }).(pulumi.StringPtrOutput)
}

// The artifact source's creation date.
func (o LookupArtifactSourceResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

// The artifact source's display name.
func (o LookupArtifactSourceResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The folder containing artifacts.
func (o LookupArtifactSourceResultOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.FolderPath }).(pulumi.StringPtrOutput)
}

// The identifier of the resource.
func (o LookupArtifactSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The location of the resource.
func (o LookupArtifactSourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o LookupArtifactSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o LookupArtifactSourceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The security token to authenticate to the artifact source.
func (o LookupArtifactSourceResultOutput) SecurityToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.SecurityToken }).(pulumi.StringPtrOutput)
}

// The artifact source's type.
func (o LookupArtifactSourceResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

// Indicates if the artifact source is enabled (values: Enabled, Disabled).
func (o LookupArtifactSourceResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The tags of the resource.
func (o LookupArtifactSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupArtifactSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o LookupArtifactSourceResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// The artifact source's URI.
func (o LookupArtifactSourceResultOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupArtifactSourceResult) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactSourceResultOutput{})
}
