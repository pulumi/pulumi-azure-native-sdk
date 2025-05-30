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

// Hybrid Connection for an App Service app.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
//
// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WebAppRelayServiceConnection struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion        pulumi.StringOutput    `pulumi:"azureApiVersion"`
	BiztalkUri             pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName             pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname               pulumi.StringPtrOutput `pulumi:"hostname"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name                     pulumi.StringOutput    `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppRelayServiceConnection registers a new resource with the given unique name, arguments, and options.
func NewWebAppRelayServiceConnection(ctx *pulumi.Context,
	name string, args *WebAppRelayServiceConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnection, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:WebAppRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:WebAppRelayServiceConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppRelayServiceConnection
	err := ctx.RegisterResource("azure-native:web:WebAppRelayServiceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppRelayServiceConnection gets an existing WebAppRelayServiceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppRelayServiceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppRelayServiceConnectionState, opts ...pulumi.ResourceOption) (*WebAppRelayServiceConnection, error) {
	var resource WebAppRelayServiceConnection
	err := ctx.ReadResource("azure-native:web:WebAppRelayServiceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppRelayServiceConnection resources.
type webAppRelayServiceConnectionState struct {
}

type WebAppRelayServiceConnectionState struct {
}

func (WebAppRelayServiceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionState)(nil)).Elem()
}

type webAppRelayServiceConnectionArgs struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             *string `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceType      *string `pulumi:"resourceType"`
}

// The set of arguments for constructing a WebAppRelayServiceConnection resource.
type WebAppRelayServiceConnectionArgs struct {
	BiztalkUri             pulumi.StringPtrInput
	EntityConnectionString pulumi.StringPtrInput
	EntityName             pulumi.StringPtrInput
	Hostname               pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	ResourceType      pulumi.StringPtrInput
}

func (WebAppRelayServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppRelayServiceConnectionArgs)(nil)).Elem()
}

type WebAppRelayServiceConnectionInput interface {
	pulumi.Input

	ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput
	ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput
}

func (*WebAppRelayServiceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppRelayServiceConnection)(nil)).Elem()
}

func (i *WebAppRelayServiceConnection) ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput {
	return i.ToWebAppRelayServiceConnectionOutputWithContext(context.Background())
}

func (i *WebAppRelayServiceConnection) ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppRelayServiceConnectionOutput)
}

type WebAppRelayServiceConnectionOutput struct{ *pulumi.OutputState }

func (WebAppRelayServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppRelayServiceConnection)(nil)).Elem()
}

func (o WebAppRelayServiceConnectionOutput) ToWebAppRelayServiceConnectionOutput() WebAppRelayServiceConnectionOutput {
	return o
}

func (o WebAppRelayServiceConnectionOutput) ToWebAppRelayServiceConnectionOutputWithContext(ctx context.Context) WebAppRelayServiceConnectionOutput {
	return o
}

// The Azure API version of the resource.
func (o WebAppRelayServiceConnectionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

func (o WebAppRelayServiceConnectionOutput) BiztalkUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.BiztalkUri }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) EntityConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.EntityConnectionString }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) EntityName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.EntityName }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.Hostname }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppRelayServiceConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppRelayServiceConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppRelayServiceConnectionOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) ResourceConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.ResourceConnectionString }).(pulumi.StringPtrOutput)
}

func (o WebAppRelayServiceConnectionOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppRelayServiceConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppRelayServiceConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppRelayServiceConnectionOutput{})
}
