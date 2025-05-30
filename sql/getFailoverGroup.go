// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a failover group.
//
// Uses Azure REST API version 2023-08-01.
//
// Other available API versions: 2015-05-01-preview, 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupFailoverGroup(ctx *pulumi.Context, args *LookupFailoverGroupArgs, opts ...pulumi.InvokeOption) (*LookupFailoverGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupFailoverGroupResult
	err := ctx.Invoke("azure-native:sql:getFailoverGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFailoverGroupArgs struct {
	// The name of the failover group.
	FailoverGroupName string `pulumi:"failoverGroupName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server containing the failover group.
	ServerName string `pulumi:"serverName"`
}

// A failover group.
type LookupFailoverGroupResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// List of databases in the failover group.
	Databases []string `pulumi:"databases"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of partner server information for the failover group.
	PartnerServers []PartnerInfoResponse `pulumi:"partnerServers"`
	// Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint *FailoverGroupReadOnlyEndpointResponse `pulumi:"readOnlyEndpoint"`
	// Read-write endpoint of the failover group instance.
	ReadWriteEndpoint FailoverGroupReadWriteEndpointResponse `pulumi:"readWriteEndpoint"`
	// Local replication role of the failover group instance.
	ReplicationRole string `pulumi:"replicationRole"`
	// Replication state of the failover group instance.
	ReplicationState string `pulumi:"replicationState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupFailoverGroupOutput(ctx *pulumi.Context, args LookupFailoverGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFailoverGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFailoverGroupResultOutput, error) {
			args := v.(LookupFailoverGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:sql:getFailoverGroup", args, LookupFailoverGroupResultOutput{}, options).(LookupFailoverGroupResultOutput), nil
		}).(LookupFailoverGroupResultOutput)
}

type LookupFailoverGroupOutputArgs struct {
	// The name of the failover group.
	FailoverGroupName pulumi.StringInput `pulumi:"failoverGroupName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server containing the failover group.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupFailoverGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverGroupArgs)(nil)).Elem()
}

// A failover group.
type LookupFailoverGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFailoverGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverGroupResult)(nil)).Elem()
}

func (o LookupFailoverGroupResultOutput) ToLookupFailoverGroupResultOutput() LookupFailoverGroupResultOutput {
	return o
}

func (o LookupFailoverGroupResultOutput) ToLookupFailoverGroupResultOutputWithContext(ctx context.Context) LookupFailoverGroupResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupFailoverGroupResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// List of databases in the failover group.
func (o LookupFailoverGroupResultOutput) Databases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) []string { return v.Databases }).(pulumi.StringArrayOutput)
}

// Resource ID.
func (o LookupFailoverGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupFailoverGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupFailoverGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of partner server information for the failover group.
func (o LookupFailoverGroupResultOutput) PartnerServers() PartnerInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) []PartnerInfoResponse { return v.PartnerServers }).(PartnerInfoResponseArrayOutput)
}

// Read-only endpoint of the failover group instance.
func (o LookupFailoverGroupResultOutput) ReadOnlyEndpoint() FailoverGroupReadOnlyEndpointResponsePtrOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) *FailoverGroupReadOnlyEndpointResponse { return v.ReadOnlyEndpoint }).(FailoverGroupReadOnlyEndpointResponsePtrOutput)
}

// Read-write endpoint of the failover group instance.
func (o LookupFailoverGroupResultOutput) ReadWriteEndpoint() FailoverGroupReadWriteEndpointResponseOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) FailoverGroupReadWriteEndpointResponse { return v.ReadWriteEndpoint }).(FailoverGroupReadWriteEndpointResponseOutput)
}

// Local replication role of the failover group instance.
func (o LookupFailoverGroupResultOutput) ReplicationRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.ReplicationRole }).(pulumi.StringOutput)
}

// Replication state of the failover group instance.
func (o LookupFailoverGroupResultOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.ReplicationState }).(pulumi.StringOutput)
}

// Resource tags.
func (o LookupFailoverGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupFailoverGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFailoverGroupResultOutput{})
}
