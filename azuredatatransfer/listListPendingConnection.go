// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredatatransfer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists all pending connections for a connection.
// Azure REST API version: 2023-10-11-preview.
//
// Other available API versions: 2024-01-25, 2024-05-07.
func ListListPendingConnection(ctx *pulumi.Context, args *ListListPendingConnectionArgs, opts ...pulumi.InvokeOption) (*ListListPendingConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListListPendingConnectionResult
	err := ctx.Invoke("azure-native:azuredatatransfer:listListPendingConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListListPendingConnectionArgs struct {
	// The name for the connection that is to be requested.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The connections list result.
type ListListPendingConnectionResult struct {
	// Link to next results
	NextLink *string `pulumi:"nextLink"`
	// Connections array.
	Value []PendingConnectionResponse `pulumi:"value"`
}

func ListListPendingConnectionOutput(ctx *pulumi.Context, args ListListPendingConnectionOutputArgs, opts ...pulumi.InvokeOption) ListListPendingConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListListPendingConnectionResult, error) {
			args := v.(ListListPendingConnectionArgs)
			r, err := ListListPendingConnection(ctx, &args, opts...)
			var s ListListPendingConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListListPendingConnectionResultOutput)
}

type ListListPendingConnectionOutputArgs struct {
	// The name for the connection that is to be requested.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListListPendingConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListPendingConnectionArgs)(nil)).Elem()
}

// The connections list result.
type ListListPendingConnectionResultOutput struct{ *pulumi.OutputState }

func (ListListPendingConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListListPendingConnectionResult)(nil)).Elem()
}

func (o ListListPendingConnectionResultOutput) ToListListPendingConnectionResultOutput() ListListPendingConnectionResultOutput {
	return o
}

func (o ListListPendingConnectionResultOutput) ToListListPendingConnectionResultOutputWithContext(ctx context.Context) ListListPendingConnectionResultOutput {
	return o
}

// Link to next results
func (o ListListPendingConnectionResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListListPendingConnectionResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

// Connections array.
func (o ListListPendingConnectionResultOutput) Value() PendingConnectionResponseArrayOutput {
	return o.ApplyT(func(v ListListPendingConnectionResult) []PendingConnectionResponse { return v.Value }).(PendingConnectionResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListListPendingConnectionResultOutput{})
}
