// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified product.
//
// Uses Azure REST API version 2022-06-01.
//
// Other available API versions: 2020-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestack [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetProductResult
	err := ctx.Invoke("azure-native:azurestack:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProductArgs struct {
	// Name of the product.
	ProductName string `pulumi:"productName"`
	// Name of the Azure Stack registration.
	RegistrationName string `pulumi:"registrationName"`
	// Name of the resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// Product information.
type GetProductResult struct {
	// The part number used for billing purposes.
	BillingPartNumber *string `pulumi:"billingPartNumber"`
	// Product compatibility with current device.
	Compatibility *CompatibilityResponse `pulumi:"compatibility"`
	// The description of the product.
	Description *string `pulumi:"description"`
	// The display name of the product.
	DisplayName *string `pulumi:"displayName"`
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string `pulumi:"etag"`
	// The identifier of the gallery item corresponding to the product.
	GalleryItemIdentity *string `pulumi:"galleryItemIdentity"`
	// Additional links available for this product.
	IconUris *IconUrisResponse `pulumi:"iconUris"`
	// ID of the resource.
	Id string `pulumi:"id"`
	// The legal terms.
	LegalTerms *string `pulumi:"legalTerms"`
	// Additional links available for this product.
	Links []ProductLinkResponse `pulumi:"links"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The offer representing the product.
	Offer *string `pulumi:"offer"`
	// The version of the product offer.
	OfferVersion *string `pulumi:"offerVersion"`
	// The length of product content.
	PayloadLength *float64 `pulumi:"payloadLength"`
	// The privacy policy.
	PrivacyPolicy *string `pulumi:"privacyPolicy"`
	// The kind of the product (virtualMachine or virtualMachineExtension)
	ProductKind *string `pulumi:"productKind"`
	// Additional properties for the product.
	ProductProperties *ProductPropertiesResponse `pulumi:"productProperties"`
	// The user-friendly name of the product publisher.
	PublisherDisplayName *string `pulumi:"publisherDisplayName"`
	// Publisher identifier.
	PublisherIdentifier *string `pulumi:"publisherIdentifier"`
	// The product SKU.
	Sku *string `pulumi:"sku"`
	// Type of Resource.
	Type string `pulumi:"type"`
	// The type of the Virtual Machine Extension.
	VmExtensionType *string `pulumi:"vmExtensionType"`
}

func GetProductOutput(ctx *pulumi.Context, args GetProductOutputArgs, opts ...pulumi.InvokeOption) GetProductResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProductResultOutput, error) {
			args := v.(GetProductArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:azurestack:getProduct", args, GetProductResultOutput{}, options).(GetProductResultOutput), nil
		}).(GetProductResultOutput)
}

type GetProductOutputArgs struct {
	// Name of the product.
	ProductName pulumi.StringInput `pulumi:"productName"`
	// Name of the Azure Stack registration.
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	// Name of the resource group.
	ResourceGroup pulumi.StringInput `pulumi:"resourceGroup"`
}

func (GetProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductArgs)(nil)).Elem()
}

// Product information.
type GetProductResultOutput struct{ *pulumi.OutputState }

func (GetProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductResult)(nil)).Elem()
}

func (o GetProductResultOutput) ToGetProductResultOutput() GetProductResultOutput {
	return o
}

func (o GetProductResultOutput) ToGetProductResultOutputWithContext(ctx context.Context) GetProductResultOutput {
	return o
}

// The part number used for billing purposes.
func (o GetProductResultOutput) BillingPartNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.BillingPartNumber }).(pulumi.StringPtrOutput)
}

// Product compatibility with current device.
func (o GetProductResultOutput) Compatibility() CompatibilityResponsePtrOutput {
	return o.ApplyT(func(v GetProductResult) *CompatibilityResponse { return v.Compatibility }).(CompatibilityResponsePtrOutput)
}

// The description of the product.
func (o GetProductResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the product.
func (o GetProductResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The entity tag used for optimistic concurrency when modifying the resource.
func (o GetProductResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The identifier of the gallery item corresponding to the product.
func (o GetProductResultOutput) GalleryItemIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.GalleryItemIdentity }).(pulumi.StringPtrOutput)
}

// Additional links available for this product.
func (o GetProductResultOutput) IconUris() IconUrisResponsePtrOutput {
	return o.ApplyT(func(v GetProductResult) *IconUrisResponse { return v.IconUris }).(IconUrisResponsePtrOutput)
}

// ID of the resource.
func (o GetProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Id }).(pulumi.StringOutput)
}

// The legal terms.
func (o GetProductResultOutput) LegalTerms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.LegalTerms }).(pulumi.StringPtrOutput)
}

// Additional links available for this product.
func (o GetProductResultOutput) Links() ProductLinkResponseArrayOutput {
	return o.ApplyT(func(v GetProductResult) []ProductLinkResponse { return v.Links }).(ProductLinkResponseArrayOutput)
}

// Name of the resource.
func (o GetProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Name }).(pulumi.StringOutput)
}

// The offer representing the product.
func (o GetProductResultOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

// The version of the product offer.
func (o GetProductResultOutput) OfferVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.OfferVersion }).(pulumi.StringPtrOutput)
}

// The length of product content.
func (o GetProductResultOutput) PayloadLength() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetProductResult) *float64 { return v.PayloadLength }).(pulumi.Float64PtrOutput)
}

// The privacy policy.
func (o GetProductResultOutput) PrivacyPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.PrivacyPolicy }).(pulumi.StringPtrOutput)
}

// The kind of the product (virtualMachine or virtualMachineExtension)
func (o GetProductResultOutput) ProductKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.ProductKind }).(pulumi.StringPtrOutput)
}

// Additional properties for the product.
func (o GetProductResultOutput) ProductProperties() ProductPropertiesResponsePtrOutput {
	return o.ApplyT(func(v GetProductResult) *ProductPropertiesResponse { return v.ProductProperties }).(ProductPropertiesResponsePtrOutput)
}

// The user-friendly name of the product publisher.
func (o GetProductResultOutput) PublisherDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.PublisherDisplayName }).(pulumi.StringPtrOutput)
}

// Publisher identifier.
func (o GetProductResultOutput) PublisherIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.PublisherIdentifier }).(pulumi.StringPtrOutput)
}

// The product SKU.
func (o GetProductResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

// Type of Resource.
func (o GetProductResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Type }).(pulumi.StringOutput)
}

// The type of the Virtual Machine Extension.
func (o GetProductResultOutput) VmExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.VmExtensionType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductResultOutput{})
}
