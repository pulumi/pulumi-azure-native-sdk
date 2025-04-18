// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists all effective virtual networks by specified network group.
//
// Uses Azure REST API version 2022-04-01-preview.
//
// Other available API versions: 2022-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListListEffectiveVirtualNetworkByNetworkGroup(ctx *pulumi.Context, args *ListListEffectiveVirtualNetworkByNetworkGroupArgs, opts ...pulumi.InvokeOption) (*ListListEffectiveVirtualNetworkByNetworkGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListEffectiveVirtualNetworkByNetworkGroupResult
	err := ctx.Invoke("azure-native:network:listListEffectiveVirtualNetworkByNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListEffectiveVirtualNetworkByNetworkGroupArgs struct {
	// The name of the network group.
	NetworkGroupName string `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
}

// Result of the request to list Effective Virtual Network. It contains a list of groups and a URL link to get the next set of results.
type ListListEffectiveVirtualNetworkByNetworkGroupResult struct {
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken *string `pulumi:"skipToken"`
	// Gets a page of EffectiveVirtualNetwork
	Value []EffectiveVirtualNetworkResponse `pulumi:"value"`
}

func ListListEffectiveVirtualNetworkByNetworkGroupOutput(ctx *pulumi.Context, args ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListListEffectiveVirtualNetworkByNetworkGroupResultOutput, error) {
			args := v.(ListListEffectiveVirtualNetworkByNetworkGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:network:listListEffectiveVirtualNetworkByNetworkGroup", args, ListListEffectiveVirtualNetworkByNetworkGroupResultOutput{}, options).(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput), nil
		}).(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput)
}

type ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs struct {
	// The name of the network group.
	NetworkGroupName pulumi.StringInput `pulumi:"networkGroupName"`
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
	SkipToken pulumi.StringPtrInput `pulumi:"skipToken"`
}

func (ListListEffectiveVirtualNetworkByNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListEffectiveVirtualNetworkByNetworkGroupArgs)(nil)).Elem()
}

// Result of the request to list Effective Virtual Network. It contains a list of groups and a URL link to get the next set of results.
type ListListEffectiveVirtualNetworkByNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListEffectiveVirtualNetworkByNetworkGroupResult)(nil)).Elem()
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListListEffectiveVirtualNetworkByNetworkGroupResultOutput() ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) ToListListEffectiveVirtualNetworkByNetworkGroupResultOutputWithContext(ctx context.Context) ListListEffectiveVirtualNetworkByNetworkGroupResultOutput {
	return o
}

// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current request) to retrieve the next page of data.
func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) SkipToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListListEffectiveVirtualNetworkByNetworkGroupResult) *string { return v.SkipToken }).(pulumi.StringPtrOutput)
}

// Gets a page of EffectiveVirtualNetwork
func (o ListListEffectiveVirtualNetworkByNetworkGroupResultOutput) Value() EffectiveVirtualNetworkResponseArrayOutput {
	return o.ApplyT(func(v ListListEffectiveVirtualNetworkByNetworkGroupResult) []EffectiveVirtualNetworkResponse {
		return v.Value
	}).(EffectiveVirtualNetworkResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListEffectiveVirtualNetworkByNetworkGroupResultOutput{})
}
