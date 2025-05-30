// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get host secrets for a function app.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2018-02-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListWebAppHostKeys(ctx *pulumi.Context, args *ListWebAppHostKeysArgs, opts ...pulumi.InvokeOption) (*ListWebAppHostKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppHostKeysResult
	err := ctx.Invoke("azure-native:web:listWebAppHostKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppHostKeysArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Functions host level keys.
type ListWebAppHostKeysResult struct {
	// Host level function keys.
	FunctionKeys map[string]string `pulumi:"functionKeys"`
	// Secret key.
	MasterKey *string `pulumi:"masterKey"`
	// System keys.
	SystemKeys map[string]string `pulumi:"systemKeys"`
}

func ListWebAppHostKeysOutput(ctx *pulumi.Context, args ListWebAppHostKeysOutputArgs, opts ...pulumi.InvokeOption) ListWebAppHostKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWebAppHostKeysResultOutput, error) {
			args := v.(ListWebAppHostKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web:listWebAppHostKeys", args, ListWebAppHostKeysResultOutput{}, options).(ListWebAppHostKeysResultOutput), nil
		}).(ListWebAppHostKeysResultOutput)
}

type ListWebAppHostKeysOutputArgs struct {
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppHostKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysArgs)(nil)).Elem()
}

// Functions host level keys.
type ListWebAppHostKeysResultOutput struct{ *pulumi.OutputState }

func (ListWebAppHostKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppHostKeysResult)(nil)).Elem()
}

func (o ListWebAppHostKeysResultOutput) ToListWebAppHostKeysResultOutput() ListWebAppHostKeysResultOutput {
	return o
}

func (o ListWebAppHostKeysResultOutput) ToListWebAppHostKeysResultOutputWithContext(ctx context.Context) ListWebAppHostKeysResultOutput {
	return o
}

// Host level function keys.
func (o ListWebAppHostKeysResultOutput) FunctionKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysResult) map[string]string { return v.FunctionKeys }).(pulumi.StringMapOutput)
}

// Secret key.
func (o ListWebAppHostKeysResultOutput) MasterKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppHostKeysResult) *string { return v.MasterKey }).(pulumi.StringPtrOutput)
}

// System keys.
func (o ListWebAppHostKeysResultOutput) SystemKeys() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppHostKeysResult) map[string]string { return v.SystemKeys }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppHostKeysResultOutput{})
}
