// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a workload classifier
func LookupWorkloadClassifier(ctx *pulumi.Context, args *LookupWorkloadClassifierArgs, opts ...pulumi.InvokeOption) (*LookupWorkloadClassifierResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkloadClassifierResult
	err := ctx.Invoke("azure-native:sql/v20230801:getWorkloadClassifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkloadClassifierArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the workload classifier.
	WorkloadClassifierName string `pulumi:"workloadClassifierName"`
	// The name of the workload group from which to receive the classifier from.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
}

// Workload classifier operations for a data warehouse
type LookupWorkloadClassifierResult struct {
	// The workload classifier context.
	Context *string `pulumi:"context"`
	// The workload classifier end time for classification.
	EndTime *string `pulumi:"endTime"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The workload classifier importance.
	Importance *string `pulumi:"importance"`
	// The workload classifier label.
	Label *string `pulumi:"label"`
	// The workload classifier member name.
	MemberName string `pulumi:"memberName"`
	// Resource name.
	Name string `pulumi:"name"`
	// The workload classifier start time for classification.
	StartTime *string `pulumi:"startTime"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWorkloadClassifierOutput(ctx *pulumi.Context, args LookupWorkloadClassifierOutputArgs, opts ...pulumi.InvokeOption) LookupWorkloadClassifierResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWorkloadClassifierResultOutput, error) {
			args := v.(LookupWorkloadClassifierArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:sql/v20230801:getWorkloadClassifier", args, LookupWorkloadClassifierResultOutput{}, options).(LookupWorkloadClassifierResultOutput), nil
		}).(LookupWorkloadClassifierResultOutput)
}

type LookupWorkloadClassifierOutputArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName pulumi.StringInput `pulumi:"serverName"`
	// The name of the workload classifier.
	WorkloadClassifierName pulumi.StringInput `pulumi:"workloadClassifierName"`
	// The name of the workload group from which to receive the classifier from.
	WorkloadGroupName pulumi.StringInput `pulumi:"workloadGroupName"`
}

func (LookupWorkloadClassifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadClassifierArgs)(nil)).Elem()
}

// Workload classifier operations for a data warehouse
type LookupWorkloadClassifierResultOutput struct{ *pulumi.OutputState }

func (LookupWorkloadClassifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkloadClassifierResult)(nil)).Elem()
}

func (o LookupWorkloadClassifierResultOutput) ToLookupWorkloadClassifierResultOutput() LookupWorkloadClassifierResultOutput {
	return o
}

func (o LookupWorkloadClassifierResultOutput) ToLookupWorkloadClassifierResultOutputWithContext(ctx context.Context) LookupWorkloadClassifierResultOutput {
	return o
}

// The workload classifier context.
func (o LookupWorkloadClassifierResultOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.Context }).(pulumi.StringPtrOutput)
}

// The workload classifier end time for classification.
func (o LookupWorkloadClassifierResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupWorkloadClassifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.Id }).(pulumi.StringOutput)
}

// The workload classifier importance.
func (o LookupWorkloadClassifierResultOutput) Importance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.Importance }).(pulumi.StringPtrOutput)
}

// The workload classifier label.
func (o LookupWorkloadClassifierResultOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// The workload classifier member name.
func (o LookupWorkloadClassifierResultOutput) MemberName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.MemberName }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupWorkloadClassifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.Name }).(pulumi.StringOutput)
}

// The workload classifier start time for classification.
func (o LookupWorkloadClassifierResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWorkloadClassifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkloadClassifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkloadClassifierResultOutput{})
}
