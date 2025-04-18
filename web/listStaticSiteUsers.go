// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the list of users of a static site.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListStaticSiteUsers(ctx *pulumi.Context, args *ListStaticSiteUsersArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteUsersResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListStaticSiteUsersResult
	err := ctx.Invoke("azure-native:web:listStaticSiteUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteUsersArgs struct {
	// The auth provider for the users.
	Authprovider string `pulumi:"authprovider"`
	// Name of the static site.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Collection of static site custom users.
type ListStaticSiteUsersResult struct {
	// Link to next page of resources.
	NextLink string `pulumi:"nextLink"`
	// Collection of resources.
	Value []StaticSiteUserARMResourceResponse `pulumi:"value"`
}

func ListStaticSiteUsersOutput(ctx *pulumi.Context, args ListStaticSiteUsersOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteUsersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListStaticSiteUsersResultOutput, error) {
			args := v.(ListStaticSiteUsersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web:listStaticSiteUsers", args, ListStaticSiteUsersResultOutput{}, options).(ListStaticSiteUsersResultOutput), nil
		}).(ListStaticSiteUsersResultOutput)
}

type ListStaticSiteUsersOutputArgs struct {
	// The auth provider for the users.
	Authprovider pulumi.StringInput `pulumi:"authprovider"`
	// Name of the static site.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteUsersArgs)(nil)).Elem()
}

// Collection of static site custom users.
type ListStaticSiteUsersResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteUsersResult)(nil)).Elem()
}

func (o ListStaticSiteUsersResultOutput) ToListStaticSiteUsersResultOutput() ListStaticSiteUsersResultOutput {
	return o
}

func (o ListStaticSiteUsersResultOutput) ToListStaticSiteUsersResultOutputWithContext(ctx context.Context) ListStaticSiteUsersResultOutput {
	return o
}

// Link to next page of resources.
func (o ListStaticSiteUsersResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteUsersResult) string { return v.NextLink }).(pulumi.StringOutput)
}

// Collection of resources.
func (o ListStaticSiteUsersResultOutput) Value() StaticSiteUserARMResourceResponseArrayOutput {
	return o.ApplyT(func(v ListStaticSiteUsersResult) []StaticSiteUserARMResourceResponse { return v.Value }).(StaticSiteUserARMResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteUsersResultOutput{})
}
