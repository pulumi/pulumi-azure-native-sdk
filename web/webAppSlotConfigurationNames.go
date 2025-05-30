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

// Slot Config names azure resource.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
//
// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WebAppSlotConfigurationNames struct {
	pulumi.CustomResourceState

	// List of application settings names.
	AppSettingNames pulumi.StringArrayOutput `pulumi:"appSettingNames"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames pulumi.StringArrayOutput `pulumi:"azureStorageConfigNames"`
	// List of connection string names.
	ConnectionStringNames pulumi.StringArrayOutput `pulumi:"connectionStringNames"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppSlotConfigurationNames registers a new resource with the given unique name, arguments, and options.
func NewWebAppSlotConfigurationNames(ctx *pulumi.Context,
	name string, args *WebAppSlotConfigurationNamesArgs, opts ...pulumi.ResourceOption) (*WebAppSlotConfigurationNames, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:WebAppSlotConfigurationNames"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:WebAppSlotConfigurationNames"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppSlotConfigurationNames
	err := ctx.RegisterResource("azure-native:web:WebAppSlotConfigurationNames", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppSlotConfigurationNames gets an existing WebAppSlotConfigurationNames resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppSlotConfigurationNames(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppSlotConfigurationNamesState, opts ...pulumi.ResourceOption) (*WebAppSlotConfigurationNames, error) {
	var resource WebAppSlotConfigurationNames
	err := ctx.ReadResource("azure-native:web:WebAppSlotConfigurationNames", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppSlotConfigurationNames resources.
type webAppSlotConfigurationNamesState struct {
}

type WebAppSlotConfigurationNamesState struct {
}

func (WebAppSlotConfigurationNamesState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotConfigurationNamesState)(nil)).Elem()
}

type webAppSlotConfigurationNamesArgs struct {
	// List of application settings names.
	AppSettingNames []string `pulumi:"appSettingNames"`
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames []string `pulumi:"azureStorageConfigNames"`
	// List of connection string names.
	ConnectionStringNames []string `pulumi:"connectionStringNames"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppSlotConfigurationNames resource.
type WebAppSlotConfigurationNamesArgs struct {
	// List of application settings names.
	AppSettingNames pulumi.StringArrayInput
	// List of external Azure storage account identifiers.
	AzureStorageConfigNames pulumi.StringArrayInput
	// List of connection string names.
	ConnectionStringNames pulumi.StringArrayInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppSlotConfigurationNamesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppSlotConfigurationNamesArgs)(nil)).Elem()
}

type WebAppSlotConfigurationNamesInput interface {
	pulumi.Input

	ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput
	ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput
}

func (*WebAppSlotConfigurationNames) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSlotConfigurationNames)(nil)).Elem()
}

func (i *WebAppSlotConfigurationNames) ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput {
	return i.ToWebAppSlotConfigurationNamesOutputWithContext(context.Background())
}

func (i *WebAppSlotConfigurationNames) ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppSlotConfigurationNamesOutput)
}

type WebAppSlotConfigurationNamesOutput struct{ *pulumi.OutputState }

func (WebAppSlotConfigurationNamesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppSlotConfigurationNames)(nil)).Elem()
}

func (o WebAppSlotConfigurationNamesOutput) ToWebAppSlotConfigurationNamesOutput() WebAppSlotConfigurationNamesOutput {
	return o
}

func (o WebAppSlotConfigurationNamesOutput) ToWebAppSlotConfigurationNamesOutputWithContext(ctx context.Context) WebAppSlotConfigurationNamesOutput {
	return o
}

// List of application settings names.
func (o WebAppSlotConfigurationNamesOutput) AppSettingNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringArrayOutput { return v.AppSettingNames }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o WebAppSlotConfigurationNamesOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// List of external Azure storage account identifiers.
func (o WebAppSlotConfigurationNamesOutput) AzureStorageConfigNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringArrayOutput { return v.AzureStorageConfigNames }).(pulumi.StringArrayOutput)
}

// List of connection string names.
func (o WebAppSlotConfigurationNamesOutput) ConnectionStringNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringArrayOutput { return v.ConnectionStringNames }).(pulumi.StringArrayOutput)
}

// Kind of resource.
func (o WebAppSlotConfigurationNamesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppSlotConfigurationNamesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o WebAppSlotConfigurationNamesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppSlotConfigurationNames) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppSlotConfigurationNamesOutput{})
}
