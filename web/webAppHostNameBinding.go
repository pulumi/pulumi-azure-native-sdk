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

// A hostname binding object.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
//
// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WebAppHostNameBinding struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Azure resource name.
	AzureResourceName pulumi.StringPtrOutput `pulumi:"azureResourceName"`
	// Azure resource type.
	AzureResourceType pulumi.StringPtrOutput `pulumi:"azureResourceType"`
	// Custom DNS record type.
	CustomHostNameDnsRecordType pulumi.StringPtrOutput `pulumi:"customHostNameDnsRecordType"`
	// Fully qualified ARM domain resource URI.
	DomainId pulumi.StringPtrOutput `pulumi:"domainId"`
	// Hostname type.
	HostNameType pulumi.StringPtrOutput `pulumi:"hostNameType"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// App Service app name.
	SiteName pulumi.StringPtrOutput `pulumi:"siteName"`
	// SSL type
	SslState pulumi.StringPtrOutput `pulumi:"sslState"`
	// SSL certificate thumbprint
	Thumbprint pulumi.StringPtrOutput `pulumi:"thumbprint"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Virtual IP address assigned to the hostname if IP based SSL is enabled.
	VirtualIP pulumi.StringOutput `pulumi:"virtualIP"`
}

// NewWebAppHostNameBinding registers a new resource with the given unique name, arguments, and options.
func NewWebAppHostNameBinding(ctx *pulumi.Context,
	name string, args *WebAppHostNameBindingArgs, opts ...pulumi.ResourceOption) (*WebAppHostNameBinding, error) {
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
			Type: pulumi.String("azure-native:web/v20150801:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:WebAppHostNameBinding"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebAppHostNameBinding
	err := ctx.RegisterResource("azure-native:web:WebAppHostNameBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppHostNameBinding gets an existing WebAppHostNameBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppHostNameBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHostNameBindingState, opts ...pulumi.ResourceOption) (*WebAppHostNameBinding, error) {
	var resource WebAppHostNameBinding
	err := ctx.ReadResource("azure-native:web:WebAppHostNameBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppHostNameBinding resources.
type webAppHostNameBindingState struct {
}

type WebAppHostNameBindingState struct {
}

func (WebAppHostNameBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingState)(nil)).Elem()
}

type webAppHostNameBindingArgs struct {
	// Azure resource name.
	AzureResourceName *string `pulumi:"azureResourceName"`
	// Azure resource type.
	AzureResourceType *AzureResourceType `pulumi:"azureResourceType"`
	// Custom DNS record type.
	CustomHostNameDnsRecordType *CustomHostNameDnsRecordType `pulumi:"customHostNameDnsRecordType"`
	// Fully qualified ARM domain resource URI.
	DomainId *string `pulumi:"domainId"`
	// Hostname in the hostname binding.
	HostName *string `pulumi:"hostName"`
	// Hostname type.
	HostNameType *HostNameType `pulumi:"hostNameType"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// App Service app name.
	SiteName *string `pulumi:"siteName"`
	// SSL type
	SslState *SslState `pulumi:"sslState"`
	// SSL certificate thumbprint
	Thumbprint *string `pulumi:"thumbprint"`
}

// The set of arguments for constructing a WebAppHostNameBinding resource.
type WebAppHostNameBindingArgs struct {
	// Azure resource name.
	AzureResourceName pulumi.StringPtrInput
	// Azure resource type.
	AzureResourceType AzureResourceTypePtrInput
	// Custom DNS record type.
	CustomHostNameDnsRecordType CustomHostNameDnsRecordTypePtrInput
	// Fully qualified ARM domain resource URI.
	DomainId pulumi.StringPtrInput
	// Hostname in the hostname binding.
	HostName pulumi.StringPtrInput
	// Hostname type.
	HostNameType HostNameTypePtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// App Service app name.
	SiteName pulumi.StringPtrInput
	// SSL type
	SslState SslStatePtrInput
	// SSL certificate thumbprint
	Thumbprint pulumi.StringPtrInput
}

func (WebAppHostNameBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingArgs)(nil)).Elem()
}

type WebAppHostNameBindingInput interface {
	pulumi.Input

	ToWebAppHostNameBindingOutput() WebAppHostNameBindingOutput
	ToWebAppHostNameBindingOutputWithContext(ctx context.Context) WebAppHostNameBindingOutput
}

func (*WebAppHostNameBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHostNameBinding)(nil)).Elem()
}

func (i *WebAppHostNameBinding) ToWebAppHostNameBindingOutput() WebAppHostNameBindingOutput {
	return i.ToWebAppHostNameBindingOutputWithContext(context.Background())
}

func (i *WebAppHostNameBinding) ToWebAppHostNameBindingOutputWithContext(ctx context.Context) WebAppHostNameBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHostNameBindingOutput)
}

type WebAppHostNameBindingOutput struct{ *pulumi.OutputState }

func (WebAppHostNameBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHostNameBinding)(nil)).Elem()
}

func (o WebAppHostNameBindingOutput) ToWebAppHostNameBindingOutput() WebAppHostNameBindingOutput {
	return o
}

func (o WebAppHostNameBindingOutput) ToWebAppHostNameBindingOutputWithContext(ctx context.Context) WebAppHostNameBindingOutput {
	return o
}

// The Azure API version of the resource.
func (o WebAppHostNameBindingOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Azure resource name.
func (o WebAppHostNameBindingOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

// Azure resource type.
func (o WebAppHostNameBindingOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

// Custom DNS record type.
func (o WebAppHostNameBindingOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

// Fully qualified ARM domain resource URI.
func (o WebAppHostNameBindingOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

// Hostname type.
func (o WebAppHostNameBindingOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.HostNameType }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o WebAppHostNameBindingOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o WebAppHostNameBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// App Service app name.
func (o WebAppHostNameBindingOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.SiteName }).(pulumi.StringPtrOutput)
}

// SSL type
func (o WebAppHostNameBindingOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.SslState }).(pulumi.StringPtrOutput)
}

// SSL certificate thumbprint
func (o WebAppHostNameBindingOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringPtrOutput { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o WebAppHostNameBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Virtual IP address assigned to the hostname if IP based SSL is enabled.
func (o WebAppHostNameBindingOutput) VirtualIP() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBinding) pulumi.StringOutput { return v.VirtualIP }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppHostNameBindingOutput{})
}
