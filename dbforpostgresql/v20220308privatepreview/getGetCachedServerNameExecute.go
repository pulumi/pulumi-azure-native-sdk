// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220308privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get available cached server name for fast provisioning
func GetGetCachedServerNameExecute(ctx *pulumi.Context, args *GetGetCachedServerNameExecuteArgs, opts ...pulumi.InvokeOption) (*GetGetCachedServerNameExecuteResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetGetCachedServerNameExecuteResult
	err := ctx.Invoke("azure-native:dbforpostgresql/v20220308privatepreview:getGetCachedServerNameExecute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGetCachedServerNameExecuteArgs struct {
	// The name of the location.
	LocationName string `pulumi:"locationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU (pricing tier) of the server.
	Sku Sku `pulumi:"sku"`
	// Storage properties of a server.
	Storage Storage `pulumi:"storage"`
	// PostgreSQL Server version.
	Version string `pulumi:"version"`
}

// Represents a resource name of a cached server
type GetGetCachedServerNameExecuteResult struct {
	// The name of available cached server
	Name string `pulumi:"name"`
}

func GetGetCachedServerNameExecuteOutput(ctx *pulumi.Context, args GetGetCachedServerNameExecuteOutputArgs, opts ...pulumi.InvokeOption) GetGetCachedServerNameExecuteResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetGetCachedServerNameExecuteResultOutput, error) {
			args := v.(GetGetCachedServerNameExecuteArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:dbforpostgresql/v20220308privatepreview:getGetCachedServerNameExecute", args, GetGetCachedServerNameExecuteResultOutput{}, options).(GetGetCachedServerNameExecuteResultOutput), nil
		}).(GetGetCachedServerNameExecuteResultOutput)
}

type GetGetCachedServerNameExecuteOutputArgs struct {
	// The name of the location.
	LocationName pulumi.StringInput `pulumi:"locationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The SKU (pricing tier) of the server.
	Sku SkuInput `pulumi:"sku"`
	// Storage properties of a server.
	Storage StorageInput `pulumi:"storage"`
	// PostgreSQL Server version.
	Version pulumi.StringInput `pulumi:"version"`
}

func (GetGetCachedServerNameExecuteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGetCachedServerNameExecuteArgs)(nil)).Elem()
}

// Represents a resource name of a cached server
type GetGetCachedServerNameExecuteResultOutput struct{ *pulumi.OutputState }

func (GetGetCachedServerNameExecuteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGetCachedServerNameExecuteResult)(nil)).Elem()
}

func (o GetGetCachedServerNameExecuteResultOutput) ToGetGetCachedServerNameExecuteResultOutput() GetGetCachedServerNameExecuteResultOutput {
	return o
}

func (o GetGetCachedServerNameExecuteResultOutput) ToGetGetCachedServerNameExecuteResultOutputWithContext(ctx context.Context) GetGetCachedServerNameExecuteResultOutput {
	return o
}

// The name of available cached server
func (o GetGetCachedServerNameExecuteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGetCachedServerNameExecuteResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGetCachedServerNameExecuteResultOutput{})
}
