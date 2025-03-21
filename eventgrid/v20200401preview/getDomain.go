// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of a domain.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:eventgrid/v20200401preview:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainArgs struct {
	// Name of the domain.
	DomainName string `pulumi:"domainName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// EventGrid Domain.
type LookupDomainResult struct {
	// Endpoint for the domain.
	Endpoint string `pulumi:"endpoint"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Identity information for the resource.
	Identity *IdentityInfoResponse `pulumi:"identity"`
	// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleResponse `pulumi:"inboundIpRules"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Metric resource id for the domain.
	MetricResourceId string `pulumi:"metricResourceId"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Provisioning state of the domain.
	ProvisioningState string `pulumi:"provisioningState"`
	// This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The Sku pricing tier for the domain.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupDomainResult
func (val *LookupDomainResult) Defaults() *LookupDomainResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.InputSchema == nil {
		inputSchema_ := "EventGridSchema"
		tmp.InputSchema = &inputSchema_
	}
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}
func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDomainResultOutput, error) {
			args := v.(LookupDomainArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:eventgrid/v20200401preview:getDomain", args, LookupDomainResultOutput{}, options).(LookupDomainResultOutput), nil
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// Name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

// EventGrid Domain.
type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// Endpoint for the domain.
func (o LookupDomainResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Fully qualified identifier of the resource.
func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity information for the resource.
func (o LookupDomainResultOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *IdentityInfoResponse { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

// This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered only if PublicNetworkAccess is enabled.
func (o LookupDomainResultOutput) InboundIpRules() InboundIpRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []InboundIpRuleResponse { return v.InboundIpRules }).(InboundIpRuleResponseArrayOutput)
}

// This determines the format that Event Grid should expect for incoming events published to the domain.
func (o LookupDomainResultOutput) InputSchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.InputSchema }).(pulumi.StringPtrOutput)
}

// Information about the InputSchemaMapping which specified the info about mapping event payload.
func (o LookupDomainResultOutput) InputSchemaMapping() JsonInputSchemaMappingResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *JsonInputSchemaMappingResponse { return v.InputSchemaMapping }).(JsonInputSchemaMappingResponsePtrOutput)
}

// Location of the resource.
func (o LookupDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

// Metric resource id for the domain.
func (o LookupDomainResultOutput) MetricResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.MetricResourceId }).(pulumi.StringOutput)
}

// Name of the resource.
func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of private endpoint connections.
func (o LookupDomainResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Provisioning state of the domain.
func (o LookupDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// This determines if traffic is allowed over public network. By default it is enabled.
// You can further restrict to specific IPs by configuring <seealso cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
func (o LookupDomainResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// The Sku pricing tier for the domain.
func (o LookupDomainResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// Tags of the resource.
func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o LookupDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
