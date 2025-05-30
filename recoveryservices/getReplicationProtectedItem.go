// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of an ASR replication protected item.
//
// Uses Azure REST API version 2024-10-01.
//
// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2025-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupReplicationProtectedItem(ctx *pulumi.Context, args *LookupReplicationProtectedItemArgs, opts ...pulumi.InvokeOption) (*LookupReplicationProtectedItemResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationProtectedItemResult
	err := ctx.Invoke("azure-native:recoveryservices:getReplicationProtectedItem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationProtectedItemArgs struct {
	// Fabric unique name.
	FabricName string `pulumi:"fabricName"`
	// Protection container name.
	ProtectionContainerName string `pulumi:"protectionContainerName"`
	// Replication protected item name.
	ReplicatedProtectedItemName string `pulumi:"replicatedProtectedItemName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName string `pulumi:"resourceName"`
}

// Replication protected item.
type LookupReplicationProtectedItemResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name string `pulumi:"name"`
	// The custom data.
	Properties ReplicationProtectedItemPropertiesResponse `pulumi:"properties"`
	// Resource Type
	Type string `pulumi:"type"`
}

func LookupReplicationProtectedItemOutput(ctx *pulumi.Context, args LookupReplicationProtectedItemOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationProtectedItemResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupReplicationProtectedItemResultOutput, error) {
			args := v.(LookupReplicationProtectedItemArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:recoveryservices:getReplicationProtectedItem", args, LookupReplicationProtectedItemResultOutput{}, options).(LookupReplicationProtectedItemResultOutput), nil
		}).(LookupReplicationProtectedItemResultOutput)
}

type LookupReplicationProtectedItemOutputArgs struct {
	// Fabric unique name.
	FabricName pulumi.StringInput `pulumi:"fabricName"`
	// Protection container name.
	ProtectionContainerName pulumi.StringInput `pulumi:"protectionContainerName"`
	// Replication protected item name.
	ReplicatedProtectedItemName pulumi.StringInput `pulumi:"replicatedProtectedItemName"`
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the recovery services vault.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupReplicationProtectedItemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectedItemArgs)(nil)).Elem()
}

// Replication protected item.
type LookupReplicationProtectedItemResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationProtectedItemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationProtectedItemResult)(nil)).Elem()
}

func (o LookupReplicationProtectedItemResultOutput) ToLookupReplicationProtectedItemResultOutput() LookupReplicationProtectedItemResultOutput {
	return o
}

func (o LookupReplicationProtectedItemResultOutput) ToLookupReplicationProtectedItemResultOutputWithContext(ctx context.Context) LookupReplicationProtectedItemResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupReplicationProtectedItemResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupReplicationProtectedItemResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource Location
func (o LookupReplicationProtectedItemResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource Name
func (o LookupReplicationProtectedItemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.Name }).(pulumi.StringOutput)
}

// The custom data.
func (o LookupReplicationProtectedItemResultOutput) Properties() ReplicationProtectedItemPropertiesResponseOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) ReplicationProtectedItemPropertiesResponse {
		return v.Properties
	}).(ReplicationProtectedItemPropertiesResponseOutput)
}

// Resource Type
func (o LookupReplicationProtectedItemResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationProtectedItemResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationProtectedItemResultOutput{})
}
