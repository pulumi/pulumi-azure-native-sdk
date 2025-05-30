// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the properties of an existing Azure Cosmos DB fleetspace under a fleet
//
// Uses Azure REST API version 2025-05-01-preview.
func LookupFleetspace(ctx *pulumi.Context, args *LookupFleetspaceArgs, opts ...pulumi.InvokeOption) (*LookupFleetspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFleetspaceResult
	err := ctx.Invoke("azure-native:cosmosdb:getFleetspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFleetspaceArgs struct {
	// Cosmos DB fleet name. Needs to be unique under a subscription.
	FleetName string `pulumi:"fleetName"`
	// Cosmos DB fleetspace name. Needs to be unique under a fleet.
	FleetspaceName string `pulumi:"fleetspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure Cosmos DB Fleetspace.
type LookupFleetspaceResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The kind of API this fleetspace belongs to. Acceptable values: 'NoSQL'
	FleetspaceApiKind *string `pulumi:"fleetspaceApiKind"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// A provisioning state of the Fleetspace.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Configuration for throughput pool in the fleetspace.
	ThroughputPoolConfiguration *FleetspacePropertiesResponseThroughputPoolConfiguration `pulumi:"throughputPoolConfiguration"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupFleetspaceOutput(ctx *pulumi.Context, args LookupFleetspaceOutputArgs, opts ...pulumi.InvokeOption) LookupFleetspaceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFleetspaceResultOutput, error) {
			args := v.(LookupFleetspaceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cosmosdb:getFleetspace", args, LookupFleetspaceResultOutput{}, options).(LookupFleetspaceResultOutput), nil
		}).(LookupFleetspaceResultOutput)
}

type LookupFleetspaceOutputArgs struct {
	// Cosmos DB fleet name. Needs to be unique under a subscription.
	FleetName pulumi.StringInput `pulumi:"fleetName"`
	// Cosmos DB fleetspace name. Needs to be unique under a fleet.
	FleetspaceName pulumi.StringInput `pulumi:"fleetspaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFleetspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetspaceArgs)(nil)).Elem()
}

// An Azure Cosmos DB Fleetspace.
type LookupFleetspaceResultOutput struct{ *pulumi.OutputState }

func (LookupFleetspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFleetspaceResult)(nil)).Elem()
}

func (o LookupFleetspaceResultOutput) ToLookupFleetspaceResultOutput() LookupFleetspaceResultOutput {
	return o
}

func (o LookupFleetspaceResultOutput) ToLookupFleetspaceResultOutputWithContext(ctx context.Context) LookupFleetspaceResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupFleetspaceResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The kind of API this fleetspace belongs to. Acceptable values: 'NoSQL'
func (o LookupFleetspaceResultOutput) FleetspaceApiKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) *string { return v.FleetspaceApiKind }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupFleetspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupFleetspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// A provisioning state of the Fleetspace.
func (o LookupFleetspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupFleetspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Configuration for throughput pool in the fleetspace.
func (o LookupFleetspaceResultOutput) ThroughputPoolConfiguration() FleetspacePropertiesResponseThroughputPoolConfigurationPtrOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) *FleetspacePropertiesResponseThroughputPoolConfiguration {
		return v.ThroughputPoolConfiguration
	}).(FleetspacePropertiesResponseThroughputPoolConfigurationPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupFleetspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFleetspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFleetspaceResultOutput{})
}
