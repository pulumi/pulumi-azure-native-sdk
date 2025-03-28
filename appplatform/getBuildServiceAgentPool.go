// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get build service agent pool.
//
// Uses Azure REST API version 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-01-01-preview, 2024-05-01-preview.
func LookupBuildServiceAgentPool(ctx *pulumi.Context, args *LookupBuildServiceAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupBuildServiceAgentPoolResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBuildServiceAgentPoolResult
	err := ctx.Invoke("azure-native:appplatform:getBuildServiceAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildServiceAgentPoolArgs struct {
	// The name of the build service agent pool resource.
	AgentPoolName string `pulumi:"agentPoolName"`
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The build service agent pool resource
type LookupBuildServiceAgentPoolResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// build service agent pool properties
	Properties BuildServiceAgentPoolPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupBuildServiceAgentPoolOutput(ctx *pulumi.Context, args LookupBuildServiceAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupBuildServiceAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBuildServiceAgentPoolResultOutput, error) {
			args := v.(LookupBuildServiceAgentPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:appplatform:getBuildServiceAgentPool", args, LookupBuildServiceAgentPoolResultOutput{}, options).(LookupBuildServiceAgentPoolResultOutput), nil
		}).(LookupBuildServiceAgentPoolResultOutput)
}

type LookupBuildServiceAgentPoolOutputArgs struct {
	// The name of the build service agent pool resource.
	AgentPoolName pulumi.StringInput `pulumi:"agentPoolName"`
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBuildServiceAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceAgentPoolArgs)(nil)).Elem()
}

// The build service agent pool resource
type LookupBuildServiceAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupBuildServiceAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceAgentPoolResult)(nil)).Elem()
}

func (o LookupBuildServiceAgentPoolResultOutput) ToLookupBuildServiceAgentPoolResultOutput() LookupBuildServiceAgentPoolResultOutput {
	return o
}

func (o LookupBuildServiceAgentPoolResultOutput) ToLookupBuildServiceAgentPoolResultOutputWithContext(ctx context.Context) LookupBuildServiceAgentPoolResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupBuildServiceAgentPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupBuildServiceAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// build service agent pool properties
func (o LookupBuildServiceAgentPoolResultOutput) Properties() BuildServiceAgentPoolPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) BuildServiceAgentPoolPropertiesResponse { return v.Properties }).(BuildServiceAgentPoolPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBuildServiceAgentPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupBuildServiceAgentPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceAgentPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildServiceAgentPoolResultOutput{})
}
