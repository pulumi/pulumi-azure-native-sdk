// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AzureStorageInfo dictionary resource.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
//
// Other available API versions: 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WebAppAzureStorageAccounts struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure storage accounts.
	Properties AzureStorageInfoValueResponseMapOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppAzureStorageAccounts registers a new resource with the given unique name, arguments, and options.
func NewWebAppAzureStorageAccounts(ctx *pulumi.Context,
	name string, args *WebAppAzureStorageAccountsArgs, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccounts, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:WebAppAzureStorageAccounts"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:WebAppAzureStorageAccounts"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppAzureStorageAccounts
	err := ctx.RegisterResource("azure-native:web:WebAppAzureStorageAccounts", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppAzureStorageAccounts gets an existing WebAppAzureStorageAccounts resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppAzureStorageAccounts(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAzureStorageAccountsState, opts ...pulumi.ResourceOption) (*WebAppAzureStorageAccounts, error) {
	var resource WebAppAzureStorageAccounts
	err := ctx.ReadResource("azure-native:web:WebAppAzureStorageAccounts", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppAzureStorageAccounts resources.
type webAppAzureStorageAccountsState struct {
}

type WebAppAzureStorageAccountsState struct {
}

func (WebAppAzureStorageAccountsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsState)(nil)).Elem()
}

type webAppAzureStorageAccountsArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Azure storage accounts.
	Properties map[string]AzureStorageInfoValue `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppAzureStorageAccounts resource.
type WebAppAzureStorageAccountsArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Azure storage accounts.
	Properties AzureStorageInfoValueMapInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppAzureStorageAccountsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAzureStorageAccountsArgs)(nil)).Elem()
}

type WebAppAzureStorageAccountsInput interface {
	pulumi.Input

	ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput
	ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput
}

func (*WebAppAzureStorageAccounts) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccounts)(nil)).Elem()
}

func (i *WebAppAzureStorageAccounts) ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput {
	return i.ToWebAppAzureStorageAccountsOutputWithContext(context.Background())
}

func (i *WebAppAzureStorageAccounts) ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAzureStorageAccountsOutput)
}

type WebAppAzureStorageAccountsOutput struct{ *pulumi.OutputState }

func (WebAppAzureStorageAccountsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAzureStorageAccounts)(nil)).Elem()
}

func (o WebAppAzureStorageAccountsOutput) ToWebAppAzureStorageAccountsOutput() WebAppAzureStorageAccountsOutput {
	return o
}

func (o WebAppAzureStorageAccountsOutput) ToWebAppAzureStorageAccountsOutputWithContext(ctx context.Context) WebAppAzureStorageAccountsOutput {
	return o
}

// The Azure API version of the resource.
func (o WebAppAzureStorageAccountsOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Kind of resource.
func (o WebAppAzureStorageAccountsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppAzureStorageAccountsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure storage accounts.
func (o WebAppAzureStorageAccountsOutput) Properties() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) AzureStorageInfoValueResponseMapOutput { return v.Properties }).(AzureStorageInfoValueResponseMapOutput)
}

// Resource type.
func (o WebAppAzureStorageAccountsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAzureStorageAccounts) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAzureStorageAccountsOutput{})
}
