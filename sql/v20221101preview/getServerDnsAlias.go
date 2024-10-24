// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a server DNS alias.
func LookupServerDnsAlias(ctx *pulumi.Context, args *LookupServerDnsAliasArgs, opts ...pulumi.InvokeOption) (*LookupServerDnsAliasResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServerDnsAliasResult
	err := ctx.Invoke("azure-native:sql/v20221101preview:getServerDnsAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerDnsAliasArgs struct {
	// The name of the server dns alias.
	DnsAliasName string `pulumi:"dnsAliasName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server that the alias is pointing to.
	ServerName string `pulumi:"serverName"`
}

// A server DNS alias.
type LookupServerDnsAliasResult struct {
	// The fully qualified DNS record for alias
	AzureDnsRecord string `pulumi:"azureDnsRecord"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupServerDnsAliasOutput(ctx *pulumi.Context, args LookupServerDnsAliasOutputArgs, opts ...pulumi.InvokeOption) LookupServerDnsAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerDnsAliasResultOutput, error) {
			args := v.(LookupServerDnsAliasArgs)
			opts = utilities.PkgInvokeDefaultOpts(opts)
			var rv LookupServerDnsAliasResult
			secret, err := ctx.InvokePackageRaw("azure-native:sql/v20221101preview:getServerDnsAlias", args, &rv, "", opts...)
			if err != nil {
				return LookupServerDnsAliasResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupServerDnsAliasResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupServerDnsAliasResultOutput), nil
			}
			return output, nil
		}).(LookupServerDnsAliasResultOutput)
}

type LookupServerDnsAliasOutputArgs struct {
	// The name of the server dns alias.
	DnsAliasName pulumi.StringInput `pulumi:"dnsAliasName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the server that the alias is pointing to.
	ServerName pulumi.StringInput `pulumi:"serverName"`
}

func (LookupServerDnsAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDnsAliasArgs)(nil)).Elem()
}

// A server DNS alias.
type LookupServerDnsAliasResultOutput struct{ *pulumi.OutputState }

func (LookupServerDnsAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerDnsAliasResult)(nil)).Elem()
}

func (o LookupServerDnsAliasResultOutput) ToLookupServerDnsAliasResultOutput() LookupServerDnsAliasResultOutput {
	return o
}

func (o LookupServerDnsAliasResultOutput) ToLookupServerDnsAliasResultOutputWithContext(ctx context.Context) LookupServerDnsAliasResultOutput {
	return o
}

// The fully qualified DNS record for alias
func (o LookupServerDnsAliasResultOutput) AzureDnsRecord() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.AzureDnsRecord }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupServerDnsAliasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupServerDnsAliasResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupServerDnsAliasResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerDnsAliasResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerDnsAliasResultOutput{})
}
