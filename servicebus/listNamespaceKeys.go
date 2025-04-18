// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the primary and secondary connection strings for the namespace.
//
// Uses Azure REST API version 2024-01-01.
//
// Other available API versions: 2018-01-01-preview, 2021-01-01-preview, 2021-06-01-preview, 2021-11-01, 2022-01-01-preview, 2022-10-01-preview, 2023-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicebus [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListNamespaceKeys(ctx *pulumi.Context, args *ListNamespaceKeysArgs, opts ...pulumi.InvokeOption) (*ListNamespaceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListNamespaceKeysResult
	err := ctx.Invoke("azure-native:servicebus:listNamespaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNamespaceKeysArgs struct {
	// The authorization rule name.
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Namespace/ServiceBus Connection String
type ListNamespaceKeysResult struct {
	// Primary connection string of the alias if GEO DR is enabled
	AliasPrimaryConnectionString string `pulumi:"aliasPrimaryConnectionString"`
	// Secondary  connection string of the alias if GEO DR is enabled
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	// A string that describes the authorization rule.
	KeyName string `pulumi:"keyName"`
	// Primary connection string of the created namespace authorization rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey string `pulumi:"primaryKey"`
	// Secondary connection string of the created namespace authorization rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListNamespaceKeysOutput(ctx *pulumi.Context, args ListNamespaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListNamespaceKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListNamespaceKeysResultOutput, error) {
			args := v.(ListNamespaceKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:servicebus:listNamespaceKeys", args, ListNamespaceKeysResultOutput{}, options).(ListNamespaceKeysResultOutput), nil
		}).(ListNamespaceKeysResultOutput)
}

type ListNamespaceKeysOutputArgs struct {
	// The authorization rule name.
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	// The namespace name
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNamespaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysArgs)(nil)).Elem()
}

// Namespace/ServiceBus Connection String
type ListNamespaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListNamespaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNamespaceKeysResult)(nil)).Elem()
}

func (o ListNamespaceKeysResultOutput) ToListNamespaceKeysResultOutput() ListNamespaceKeysResultOutput {
	return o
}

func (o ListNamespaceKeysResultOutput) ToListNamespaceKeysResultOutputWithContext(ctx context.Context) ListNamespaceKeysResultOutput {
	return o
}

// Primary connection string of the alias if GEO DR is enabled
func (o ListNamespaceKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

// Secondary  connection string of the alias if GEO DR is enabled
func (o ListNamespaceKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

// A string that describes the authorization rule.
func (o ListNamespaceKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

// Primary connection string of the created namespace authorization rule.
func (o ListNamespaceKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListNamespaceKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

// Secondary connection string of the created namespace authorization rule.
func (o ListNamespaceKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

// A base64-encoded 256-bit primary key for signing and validating the SAS token.
func (o ListNamespaceKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListNamespaceKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNamespaceKeysResultOutput{})
}
