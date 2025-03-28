// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagemover

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an Agent resource.
//
// Uses Azure REST API version 2023-03-01.
//
// Other available API versions: 2023-07-01-preview, 2023-10-01, 2024-07-01.
func LookupAgent(ctx *pulumi.Context, args *LookupAgentArgs, opts ...pulumi.InvokeOption) (*LookupAgentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAgentResult
	err := ctx.Invoke("azure-native:storagemover:getAgent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentArgs struct {
	// The name of the Agent resource.
	AgentName string `pulumi:"agentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName string `pulumi:"storageMoverName"`
}

// The Agent resource.
type LookupAgentResult struct {
	// The Agent status.
	AgentStatus string `pulumi:"agentStatus"`
	// The Agent version.
	AgentVersion string `pulumi:"agentVersion"`
	// The fully qualified resource ID of the Hybrid Compute resource for the Agent.
	ArcResourceId string `pulumi:"arcResourceId"`
	// The VM UUID of the Hybrid Compute resource for the Agent.
	ArcVmUuid string `pulumi:"arcVmUuid"`
	// A description for the Agent.
	Description  *string                             `pulumi:"description"`
	ErrorDetails AgentPropertiesResponseErrorDetails `pulumi:"errorDetails"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The last updated time of the Agent status.
	LastStatusUpdate string `pulumi:"lastStatusUpdate"`
	// Local IP address reported by the Agent.
	LocalIPAddress string `pulumi:"localIPAddress"`
	// Available memory reported by the Agent, in MB.
	MemoryInMB float64 `pulumi:"memoryInMB"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Available compute cores reported by the Agent.
	NumberOfCores float64 `pulumi:"numberOfCores"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource system metadata.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Uptime of the Agent in seconds.
	UptimeInSeconds float64 `pulumi:"uptimeInSeconds"`
}

func LookupAgentOutput(ctx *pulumi.Context, args LookupAgentOutputArgs, opts ...pulumi.InvokeOption) LookupAgentResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAgentResultOutput, error) {
			args := v.(LookupAgentArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:storagemover:getAgent", args, LookupAgentResultOutput{}, options).(LookupAgentResultOutput), nil
		}).(LookupAgentResultOutput)
}

type LookupAgentOutputArgs struct {
	// The name of the Agent resource.
	AgentName pulumi.StringInput `pulumi:"agentName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Storage Mover resource.
	StorageMoverName pulumi.StringInput `pulumi:"storageMoverName"`
}

func (LookupAgentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentArgs)(nil)).Elem()
}

// The Agent resource.
type LookupAgentResultOutput struct{ *pulumi.OutputState }

func (LookupAgentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentResult)(nil)).Elem()
}

func (o LookupAgentResultOutput) ToLookupAgentResultOutput() LookupAgentResultOutput {
	return o
}

func (o LookupAgentResultOutput) ToLookupAgentResultOutputWithContext(ctx context.Context) LookupAgentResultOutput {
	return o
}

// The Agent status.
func (o LookupAgentResultOutput) AgentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.AgentStatus }).(pulumi.StringOutput)
}

// The Agent version.
func (o LookupAgentResultOutput) AgentVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.AgentVersion }).(pulumi.StringOutput)
}

// The fully qualified resource ID of the Hybrid Compute resource for the Agent.
func (o LookupAgentResultOutput) ArcResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.ArcResourceId }).(pulumi.StringOutput)
}

// The VM UUID of the Hybrid Compute resource for the Agent.
func (o LookupAgentResultOutput) ArcVmUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.ArcVmUuid }).(pulumi.StringOutput)
}

// A description for the Agent.
func (o LookupAgentResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgentResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAgentResultOutput) ErrorDetails() AgentPropertiesResponseErrorDetailsOutput {
	return o.ApplyT(func(v LookupAgentResult) AgentPropertiesResponseErrorDetails { return v.ErrorDetails }).(AgentPropertiesResponseErrorDetailsOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupAgentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Id }).(pulumi.StringOutput)
}

// The last updated time of the Agent status.
func (o LookupAgentResultOutput) LastStatusUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.LastStatusUpdate }).(pulumi.StringOutput)
}

// Local IP address reported by the Agent.
func (o LookupAgentResultOutput) LocalIPAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.LocalIPAddress }).(pulumi.StringOutput)
}

// Available memory reported by the Agent, in MB.
func (o LookupAgentResultOutput) MemoryInMB() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentResult) float64 { return v.MemoryInMB }).(pulumi.Float64Output)
}

// The name of the resource
func (o LookupAgentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Available compute cores reported by the Agent.
func (o LookupAgentResultOutput) NumberOfCores() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentResult) float64 { return v.NumberOfCores }).(pulumi.Float64Output)
}

// The provisioning state of this resource.
func (o LookupAgentResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Resource system metadata.
func (o LookupAgentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAgentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAgentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentResult) string { return v.Type }).(pulumi.StringOutput)
}

// Uptime of the Agent in seconds.
func (o LookupAgentResultOutput) UptimeInSeconds() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAgentResult) float64 { return v.UptimeInSeconds }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentResultOutput{})
}
