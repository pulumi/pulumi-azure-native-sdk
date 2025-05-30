// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package app

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Workflow properties definition.
//
// Uses Azure REST API version 2024-10-02-preview.
//
// Other available API versions: 2024-02-02-preview, 2024-08-02-preview, 2025-02-02-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native app [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListLogicAppWorkflowsConnections(ctx *pulumi.Context, args *ListLogicAppWorkflowsConnectionsArgs, opts ...pulumi.InvokeOption) (*ListLogicAppWorkflowsConnectionsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListLogicAppWorkflowsConnectionsResult
	err := ctx.Invoke("azure-native:app:listLogicAppWorkflowsConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLogicAppWorkflowsConnectionsArgs struct {
	// Name of the Container App.
	ContainerAppName string `pulumi:"containerAppName"`
	// Name of the Logic App, the extension resource.
	LogicAppName string `pulumi:"logicAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Workflow properties definition.
type ListLogicAppWorkflowsConnectionsResult struct {
	// The resource id.
	Id string `pulumi:"id"`
	// The resource kind.
	Kind *string `pulumi:"kind"`
	// The resource location.
	Location *string `pulumi:"location"`
	// Gets the resource name.
	Name string `pulumi:"name"`
	// Additional workflow properties.
	Properties WorkflowEnvelopeResponseProperties `pulumi:"properties"`
	// Gets the resource type.
	Type string `pulumi:"type"`
}

func ListLogicAppWorkflowsConnectionsOutput(ctx *pulumi.Context, args ListLogicAppWorkflowsConnectionsOutputArgs, opts ...pulumi.InvokeOption) ListLogicAppWorkflowsConnectionsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListLogicAppWorkflowsConnectionsResultOutput, error) {
			args := v.(ListLogicAppWorkflowsConnectionsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:app:listLogicAppWorkflowsConnections", args, ListLogicAppWorkflowsConnectionsResultOutput{}, options).(ListLogicAppWorkflowsConnectionsResultOutput), nil
		}).(ListLogicAppWorkflowsConnectionsResultOutput)
}

type ListLogicAppWorkflowsConnectionsOutputArgs struct {
	// Name of the Container App.
	ContainerAppName pulumi.StringInput `pulumi:"containerAppName"`
	// Name of the Logic App, the extension resource.
	LogicAppName pulumi.StringInput `pulumi:"logicAppName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListLogicAppWorkflowsConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLogicAppWorkflowsConnectionsArgs)(nil)).Elem()
}

// Workflow properties definition.
type ListLogicAppWorkflowsConnectionsResultOutput struct{ *pulumi.OutputState }

func (ListLogicAppWorkflowsConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListLogicAppWorkflowsConnectionsResult)(nil)).Elem()
}

func (o ListLogicAppWorkflowsConnectionsResultOutput) ToListLogicAppWorkflowsConnectionsResultOutput() ListLogicAppWorkflowsConnectionsResultOutput {
	return o
}

func (o ListLogicAppWorkflowsConnectionsResultOutput) ToListLogicAppWorkflowsConnectionsResultOutputWithContext(ctx context.Context) ListLogicAppWorkflowsConnectionsResultOutput {
	return o
}

// The resource id.
func (o ListLogicAppWorkflowsConnectionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListLogicAppWorkflowsConnectionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource kind.
func (o ListLogicAppWorkflowsConnectionsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListLogicAppWorkflowsConnectionsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The resource location.
func (o ListLogicAppWorkflowsConnectionsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListLogicAppWorkflowsConnectionsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o ListLogicAppWorkflowsConnectionsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListLogicAppWorkflowsConnectionsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Additional workflow properties.
func (o ListLogicAppWorkflowsConnectionsResultOutput) Properties() WorkflowEnvelopeResponsePropertiesOutput {
	return o.ApplyT(func(v ListLogicAppWorkflowsConnectionsResult) WorkflowEnvelopeResponseProperties { return v.Properties }).(WorkflowEnvelopeResponsePropertiesOutput)
}

// Gets the resource type.
func (o ListLogicAppWorkflowsConnectionsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListLogicAppWorkflowsConnectionsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListLogicAppWorkflowsConnectionsResultOutput{})
}
