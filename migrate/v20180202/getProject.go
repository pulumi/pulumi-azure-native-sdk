// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180202

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the project with the specified name.
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Azure Migrate Project.
type LookupProjectResult struct {
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp string `pulumi:"createdTimestamp"`
	// ARM ID of the Service Map workspace created by user.
	CustomerWorkspaceId *string `pulumi:"customerWorkspaceId"`
	// Location of the Service Map workspace created by user.
	CustomerWorkspaceLocation *string `pulumi:"customerWorkspaceLocation"`
	// Reports whether project is under discovery.
	DiscoveryStatus string `pulumi:"discoveryStatus"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}
	Id string `pulumi:"id"`
	// Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
	LastAssessmentTimestamp string `pulumi:"lastAssessmentTimestamp"`
	// Session id of the last discovery.
	LastDiscoverySessionId string `pulumi:"lastDiscoverySessionId"`
	// Time when this project was created. Date-Time represented in ISO-8601 format. This value will be null until discovery is complete.
	LastDiscoveryTimestamp string `pulumi:"lastDiscoveryTimestamp"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Name of the project.
	Name string `pulumi:"name"`
	// Number of assessments created in the project.
	NumberOfAssessments int `pulumi:"numberOfAssessments"`
	// Number of groups created in the project.
	NumberOfGroups int `pulumi:"numberOfGroups"`
	// Number of machines in the project.
	NumberOfMachines int `pulumi:"numberOfMachines"`
	// Provisioning state of the project.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Tags provided by Azure Tagging service.
	Tags interface{} `pulumi:"tags"`
	// Type of the object = [Microsoft.Migrate/projects].
	Type string `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp string `pulumi:"updatedTimestamp"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectResultOutput, error) {
			args := v.(LookupProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:migrate/v20180202:getProject", args, LookupProjectResultOutput{}, options).(LookupProjectResultOutput), nil
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// Azure Migrate Project.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// Time when this project was created. Date-Time represented in ISO-8601 format.
func (o LookupProjectResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

// ARM ID of the Service Map workspace created by user.
func (o LookupProjectResultOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

// Location of the Service Map workspace created by user.
func (o LookupProjectResultOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.CustomerWorkspaceLocation }).(pulumi.StringPtrOutput)
}

// Reports whether project is under discovery.
func (o LookupProjectResultOutput) DiscoveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.DiscoveryStatus }).(pulumi.StringOutput)
}

// For optimistic concurrency control.
func (o LookupProjectResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Path reference to this project /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/projects/{projectName}
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
func (o LookupProjectResultOutput) LastAssessmentTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.LastAssessmentTimestamp }).(pulumi.StringOutput)
}

// Session id of the last discovery.
func (o LookupProjectResultOutput) LastDiscoverySessionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.LastDiscoverySessionId }).(pulumi.StringOutput)
}

// Time when this project was created. Date-Time represented in ISO-8601 format. This value will be null until discovery is complete.
func (o LookupProjectResultOutput) LastDiscoveryTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.LastDiscoveryTimestamp }).(pulumi.StringOutput)
}

// Azure location in which project is created.
func (o LookupProjectResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of assessments created in the project.
func (o LookupProjectResultOutput) NumberOfAssessments() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NumberOfAssessments }).(pulumi.IntOutput)
}

// Number of groups created in the project.
func (o LookupProjectResultOutput) NumberOfGroups() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NumberOfGroups }).(pulumi.IntOutput)
}

// Number of machines in the project.
func (o LookupProjectResultOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.NumberOfMachines }).(pulumi.IntOutput)
}

// Provisioning state of the project.
func (o LookupProjectResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Tags provided by Azure Tagging service.
func (o LookupProjectResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupProjectResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

// Type of the object = [Microsoft.Migrate/projects].
func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

// Time when this project was last updated. Date-Time represented in ISO-8601 format.
func (o LookupProjectResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
