// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Namespace/EventHub Connection String
func ListEventHubKeys(ctx *pulumi.Context, args *ListEventHubKeysArgs, opts ...pulumi.InvokeOption) (*ListEventHubKeysResult, error) {
	var rv ListEventHubKeysResult
	err := ctx.Invoke("azure-native:eventhub/v20170401:listEventHubKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEventHubKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	EventHubName          string `pulumi:"eventHubName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}

// Namespace/EventHub Connection String
type ListEventHubKeysResult struct {
	AliasPrimaryConnectionString   string `pulumi:"aliasPrimaryConnectionString"`
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	KeyName                        string `pulumi:"keyName"`
	PrimaryConnectionString        string `pulumi:"primaryConnectionString"`
	PrimaryKey                     string `pulumi:"primaryKey"`
	SecondaryConnectionString      string `pulumi:"secondaryConnectionString"`
	SecondaryKey                   string `pulumi:"secondaryKey"`
}

func ListEventHubKeysOutput(ctx *pulumi.Context, args ListEventHubKeysOutputArgs, opts ...pulumi.InvokeOption) ListEventHubKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListEventHubKeysResult, error) {
			args := v.(ListEventHubKeysArgs)
			r, err := ListEventHubKeys(ctx, &args, opts...)
			var s ListEventHubKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListEventHubKeysResultOutput)
}

type ListEventHubKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	EventHubName          pulumi.StringInput `pulumi:"eventHubName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListEventHubKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEventHubKeysArgs)(nil)).Elem()
}

// Namespace/EventHub Connection String
type ListEventHubKeysResultOutput struct{ *pulumi.OutputState }

func (ListEventHubKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListEventHubKeysResult)(nil)).Elem()
}

func (o ListEventHubKeysResultOutput) ToListEventHubKeysResultOutput() ListEventHubKeysResultOutput {
	return o
}

func (o ListEventHubKeysResultOutput) ToListEventHubKeysResultOutputWithContext(ctx context.Context) ListEventHubKeysResultOutput {
	return o
}

func (o ListEventHubKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ListEventHubKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

func (o ListEventHubKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ListEventHubKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ListEventHubKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListEventHubKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

func (o ListEventHubKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListEventHubKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListEventHubKeysResultOutput{})
}