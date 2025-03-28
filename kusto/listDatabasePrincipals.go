// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of database principals of the given Kusto cluster and database.
//
// Uses Azure REST API version 2022-12-29.
//
// Other available API versions: 2022-07-07, 2023-05-02, 2023-08-15, 2024-04-13.
func ListDatabasePrincipals(ctx *pulumi.Context, args *ListDatabasePrincipalsArgs, opts ...pulumi.InvokeOption) (*ListDatabasePrincipalsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDatabasePrincipalsResult
	err := ctx.Invoke("azure-native:kusto:listDatabasePrincipals", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDatabasePrincipalsArgs struct {
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName string `pulumi:"databaseName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The list Kusto database principals operation response.
type ListDatabasePrincipalsResult struct {
	// The list of Kusto database principals.
	Value []DatabasePrincipalResponse `pulumi:"value"`
}

func ListDatabasePrincipalsOutput(ctx *pulumi.Context, args ListDatabasePrincipalsOutputArgs, opts ...pulumi.InvokeOption) ListDatabasePrincipalsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDatabasePrincipalsResultOutput, error) {
			args := v.(ListDatabasePrincipalsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:kusto:listDatabasePrincipals", args, ListDatabasePrincipalsResultOutput{}, options).(ListDatabasePrincipalsResultOutput), nil
		}).(ListDatabasePrincipalsResultOutput)
}

type ListDatabasePrincipalsOutputArgs struct {
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The name of the resource group containing the Kusto cluster.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDatabasePrincipalsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasePrincipalsArgs)(nil)).Elem()
}

// The list Kusto database principals operation response.
type ListDatabasePrincipalsResultOutput struct{ *pulumi.OutputState }

func (ListDatabasePrincipalsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDatabasePrincipalsResult)(nil)).Elem()
}

func (o ListDatabasePrincipalsResultOutput) ToListDatabasePrincipalsResultOutput() ListDatabasePrincipalsResultOutput {
	return o
}

func (o ListDatabasePrincipalsResultOutput) ToListDatabasePrincipalsResultOutputWithContext(ctx context.Context) ListDatabasePrincipalsResultOutput {
	return o
}

// The list of Kusto database principals.
func (o ListDatabasePrincipalsResultOutput) Value() DatabasePrincipalResponseArrayOutput {
	return o.ApplyT(func(v ListDatabasePrincipalsResult) []DatabasePrincipalResponse { return v.Value }).(DatabasePrincipalResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDatabasePrincipalsResultOutput{})
}
