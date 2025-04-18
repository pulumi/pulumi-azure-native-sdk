// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the detailed information about a linked server of a redis cache (requires Premium SKU).
//
// Uses Azure REST API version 2017-02-01.
func LookupRedisLinkedServer(ctx *pulumi.Context, args *LookupRedisLinkedServerArgs, opts ...pulumi.InvokeOption) (*LookupRedisLinkedServerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRedisLinkedServerResult
	err := ctx.Invoke("azure-native:redis:getRedisLinkedServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisLinkedServerArgs struct {
	// The name of the linked server.
	LinkedServerName string `pulumi:"linkedServerName"`
	// The name of the redis cache.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response to put/get linked server (with properties) for Redis cache.
type LookupRedisLinkedServerResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId string `pulumi:"linkedRedisCacheId"`
	// Location of the linked redis cache.
	LinkedRedisCacheLocation string `pulumi:"linkedRedisCacheLocation"`
	// Resource name.
	Name string `pulumi:"name"`
	// Terminal state of the link between primary and secondary redis cache.
	ProvisioningState string `pulumi:"provisioningState"`
	// Role of the linked server.
	ServerRole string `pulumi:"serverRole"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupRedisLinkedServerOutput(ctx *pulumi.Context, args LookupRedisLinkedServerOutputArgs, opts ...pulumi.InvokeOption) LookupRedisLinkedServerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRedisLinkedServerResultOutput, error) {
			args := v.(LookupRedisLinkedServerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:redis:getRedisLinkedServer", args, LookupRedisLinkedServerResultOutput{}, options).(LookupRedisLinkedServerResultOutput), nil
		}).(LookupRedisLinkedServerResultOutput)
}

type LookupRedisLinkedServerOutputArgs struct {
	// The name of the linked server.
	LinkedServerName pulumi.StringInput `pulumi:"linkedServerName"`
	// The name of the redis cache.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRedisLinkedServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisLinkedServerArgs)(nil)).Elem()
}

// Response to put/get linked server (with properties) for Redis cache.
type LookupRedisLinkedServerResultOutput struct{ *pulumi.OutputState }

func (LookupRedisLinkedServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisLinkedServerResult)(nil)).Elem()
}

func (o LookupRedisLinkedServerResultOutput) ToLookupRedisLinkedServerResultOutput() LookupRedisLinkedServerResultOutput {
	return o
}

func (o LookupRedisLinkedServerResultOutput) ToLookupRedisLinkedServerResultOutputWithContext(ctx context.Context) LookupRedisLinkedServerResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupRedisLinkedServerResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupRedisLinkedServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// Fully qualified resourceId of the linked redis cache.
func (o LookupRedisLinkedServerResultOutput) LinkedRedisCacheId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.LinkedRedisCacheId }).(pulumi.StringOutput)
}

// Location of the linked redis cache.
func (o LookupRedisLinkedServerResultOutput) LinkedRedisCacheLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.LinkedRedisCacheLocation }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupRedisLinkedServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Terminal state of the link between primary and secondary redis cache.
func (o LookupRedisLinkedServerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Role of the linked server.
func (o LookupRedisLinkedServerResultOutput) ServerRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.ServerRole }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupRedisLinkedServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisLinkedServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisLinkedServerResultOutput{})
}
