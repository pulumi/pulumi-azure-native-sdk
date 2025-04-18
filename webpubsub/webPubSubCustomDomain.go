// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package webpubsub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A custom domain
//
// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
//
// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native webpubsub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WebPubSubCustomDomain struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Reference to a resource.
	CustomCertificate ResourceReferenceResponseOutput `pulumi:"customCertificate"`
	// The custom domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebPubSubCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewWebPubSubCustomDomain(ctx *pulumi.Context,
	name string, args *WebPubSubCustomDomainArgs, opts ...pulumi.ResourceOption) (*WebPubSubCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomCertificate == nil {
		return nil, errors.New("invalid value for required argument 'CustomCertificate'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:webpubsub/v20220801preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230201:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230301preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230601preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20230801preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20240101preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20240301:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20240401preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20240801preview:WebPubSubCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:webpubsub/v20241001preview:WebPubSubCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebPubSubCustomDomain
	err := ctx.RegisterResource("azure-native:webpubsub:WebPubSubCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebPubSubCustomDomain gets an existing WebPubSubCustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebPubSubCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebPubSubCustomDomainState, opts ...pulumi.ResourceOption) (*WebPubSubCustomDomain, error) {
	var resource WebPubSubCustomDomain
	err := ctx.ReadResource("azure-native:webpubsub:WebPubSubCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebPubSubCustomDomain resources.
type webPubSubCustomDomainState struct {
}

type WebPubSubCustomDomainState struct {
}

func (WebPubSubCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomDomainState)(nil)).Elem()
}

type webPubSubCustomDomainArgs struct {
	// Reference to a resource.
	CustomCertificate ResourceReference `pulumi:"customCertificate"`
	// The custom domain name.
	DomainName string `pulumi:"domainName"`
	// Custom domain name.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a WebPubSubCustomDomain resource.
type WebPubSubCustomDomainArgs struct {
	// Reference to a resource.
	CustomCertificate ResourceReferenceInput
	// The custom domain name.
	DomainName pulumi.StringInput
	// Custom domain name.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
}

func (WebPubSubCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webPubSubCustomDomainArgs)(nil)).Elem()
}

type WebPubSubCustomDomainInput interface {
	pulumi.Input

	ToWebPubSubCustomDomainOutput() WebPubSubCustomDomainOutput
	ToWebPubSubCustomDomainOutputWithContext(ctx context.Context) WebPubSubCustomDomainOutput
}

func (*WebPubSubCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomDomain)(nil)).Elem()
}

func (i *WebPubSubCustomDomain) ToWebPubSubCustomDomainOutput() WebPubSubCustomDomainOutput {
	return i.ToWebPubSubCustomDomainOutputWithContext(context.Background())
}

func (i *WebPubSubCustomDomain) ToWebPubSubCustomDomainOutputWithContext(ctx context.Context) WebPubSubCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebPubSubCustomDomainOutput)
}

type WebPubSubCustomDomainOutput struct{ *pulumi.OutputState }

func (WebPubSubCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebPubSubCustomDomain)(nil)).Elem()
}

func (o WebPubSubCustomDomainOutput) ToWebPubSubCustomDomainOutput() WebPubSubCustomDomainOutput {
	return o
}

func (o WebPubSubCustomDomainOutput) ToWebPubSubCustomDomainOutputWithContext(ctx context.Context) WebPubSubCustomDomainOutput {
	return o
}

// The Azure API version of the resource.
func (o WebPubSubCustomDomainOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Reference to a resource.
func (o WebPubSubCustomDomainOutput) CustomCertificate() ResourceReferenceResponseOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) ResourceReferenceResponseOutput { return v.CustomCertificate }).(ResourceReferenceResponseOutput)
}

// The custom domain name.
func (o WebPubSubCustomDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The name of the resource
func (o WebPubSubCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o WebPubSubCustomDomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WebPubSubCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WebPubSubCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebPubSubCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebPubSubCustomDomainOutput{})
}
