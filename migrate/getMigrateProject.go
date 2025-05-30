// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Migrate Project REST Resource.
//
// Uses Azure REST API version 2018-09-01-preview.
func LookupMigrateProject(ctx *pulumi.Context, args *LookupMigrateProjectArgs, opts ...pulumi.InvokeOption) (*LookupMigrateProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMigrateProjectResult
	err := ctx.Invoke("azure-native:migrate:getMigrateProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateProjectArgs struct {
	// Name of the Azure Migrate project.
	MigrateProjectName string `pulumi:"migrateProjectName"`
	// Name of the Azure Resource Group that migrate project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Migrate Project REST Resource.
type LookupMigrateProjectResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets or sets the eTag for concurrency control.
	ETag *string `pulumi:"eTag"`
	// Gets the relative URL to get this migrate project.
	Id string `pulumi:"id"`
	// Gets or sets the Azure location in which migrate project is created.
	Location *string `pulumi:"location"`
	// Gets the name of the migrate project.
	Name string `pulumi:"name"`
	// Gets or sets the nested properties.
	Properties MigrateProjectPropertiesResponse `pulumi:"properties"`
	// Gets or sets the tags.
	Tags *MigrateProjectResponseTags `pulumi:"tags"`
	// Handled by resource provider. Type = Microsoft.Migrate/MigrateProject.
	Type string `pulumi:"type"`
}

func LookupMigrateProjectOutput(ctx *pulumi.Context, args LookupMigrateProjectOutputArgs, opts ...pulumi.InvokeOption) LookupMigrateProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMigrateProjectResultOutput, error) {
			args := v.(LookupMigrateProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:migrate:getMigrateProject", args, LookupMigrateProjectResultOutput{}, options).(LookupMigrateProjectResultOutput), nil
		}).(LookupMigrateProjectResultOutput)
}

type LookupMigrateProjectOutputArgs struct {
	// Name of the Azure Migrate project.
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	// Name of the Azure Resource Group that migrate project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMigrateProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectArgs)(nil)).Elem()
}

// Migrate Project REST Resource.
type LookupMigrateProjectResultOutput struct{ *pulumi.OutputState }

func (LookupMigrateProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectResult)(nil)).Elem()
}

func (o LookupMigrateProjectResultOutput) ToLookupMigrateProjectResultOutput() LookupMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectResultOutput) ToLookupMigrateProjectResultOutputWithContext(ctx context.Context) LookupMigrateProjectResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMigrateProjectResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the eTag for concurrency control.
func (o LookupMigrateProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Gets the relative URL to get this migrate project.
func (o LookupMigrateProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Gets or sets the Azure location in which migrate project is created.
func (o LookupMigrateProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the name of the migrate project.
func (o LookupMigrateProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the nested properties.
func (o LookupMigrateProjectResultOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) MigrateProjectPropertiesResponse { return v.Properties }).(MigrateProjectPropertiesResponseOutput)
}

// Gets or sets the tags.
func (o LookupMigrateProjectResultOutput) Tags() MigrateProjectResponseTagsPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) *MigrateProjectResponseTags { return v.Tags }).(MigrateProjectResponseTagsPtrOutput)
}

// Handled by resource provider. Type = Microsoft.Migrate/MigrateProject.
func (o LookupMigrateProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrateProjectResultOutput{})
}
