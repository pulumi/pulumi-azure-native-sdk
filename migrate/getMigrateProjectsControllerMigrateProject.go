// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information related to a specific migrate project. Returns a json object of type 'migrateProject' as specified in the models section.
//
// Uses Azure REST API version 2020-05-01.
//
// Other available API versions: 2023-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupMigrateProjectsControllerMigrateProject(ctx *pulumi.Context, args *LookupMigrateProjectsControllerMigrateProjectArgs, opts ...pulumi.InvokeOption) (*LookupMigrateProjectsControllerMigrateProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMigrateProjectsControllerMigrateProjectResult
	err := ctx.Invoke("azure-native:migrate:getMigrateProjectsControllerMigrateProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMigrateProjectsControllerMigrateProjectArgs struct {
	// Migrate project name.
	MigrateProjectName string `pulumi:"migrateProjectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Migrate project.
type LookupMigrateProjectsControllerMigrateProjectResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{projectName}
	Id string `pulumi:"id"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the project.
	Name string `pulumi:"name"`
	// Properties of a migrate project.
	Properties MigrateProjectPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Type of the object = [Microsoft.Migrate/migrateProjects].
	Type string `pulumi:"type"`
}

func LookupMigrateProjectsControllerMigrateProjectOutput(ctx *pulumi.Context, args LookupMigrateProjectsControllerMigrateProjectOutputArgs, opts ...pulumi.InvokeOption) LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMigrateProjectsControllerMigrateProjectResultOutput, error) {
			args := v.(LookupMigrateProjectsControllerMigrateProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:migrate:getMigrateProjectsControllerMigrateProject", args, LookupMigrateProjectsControllerMigrateProjectResultOutput{}, options).(LookupMigrateProjectsControllerMigrateProjectResultOutput), nil
		}).(LookupMigrateProjectsControllerMigrateProjectResultOutput)
}

type LookupMigrateProjectsControllerMigrateProjectOutputArgs struct {
	// Migrate project name.
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMigrateProjectsControllerMigrateProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectsControllerMigrateProjectArgs)(nil)).Elem()
}

// Migrate project.
type LookupMigrateProjectsControllerMigrateProjectResultOutput struct{ *pulumi.OutputState }

func (LookupMigrateProjectsControllerMigrateProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMigrateProjectsControllerMigrateProjectResult)(nil)).Elem()
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ToLookupMigrateProjectsControllerMigrateProjectResultOutput() LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return o
}

func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ToLookupMigrateProjectsControllerMigrateProjectResultOutputWithContext(ctx context.Context) LookupMigrateProjectsControllerMigrateProjectResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// For optimistic concurrency control.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{projectName}
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Azure location in which project is created.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of a migrate project.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) MigrateProjectPropertiesResponse {
		return v.Properties
	}).(MigrateProjectPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the object = [Microsoft.Migrate/migrateProjects].
func (o LookupMigrateProjectsControllerMigrateProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMigrateProjectsControllerMigrateProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMigrateProjectsControllerMigrateProjectResultOutput{})
}
