// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domainregistration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Address information for domain registration.
type Address struct {
	// First line of an Address.
	Address1 string `pulumi:"address1"`
	// The second line of the Address. Optional.
	Address2 *string `pulumi:"address2"`
	// The city for the address.
	City string `pulumi:"city"`
	// The country for the address.
	Country string `pulumi:"country"`
	// The postal code for the address.
	PostalCode string `pulumi:"postalCode"`
	// The state or province for the address.
	State string `pulumi:"state"`
}

// AddressInput is an input type that accepts AddressArgs and AddressOutput values.
// You can construct a concrete instance of `AddressInput` via:
//
//	AddressArgs{...}
type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(context.Context) AddressOutput
}

// Address information for domain registration.
type AddressArgs struct {
	// First line of an Address.
	Address1 pulumi.StringInput `pulumi:"address1"`
	// The second line of the Address. Optional.
	Address2 pulumi.StringPtrInput `pulumi:"address2"`
	// The city for the address.
	City pulumi.StringInput `pulumi:"city"`
	// The country for the address.
	Country pulumi.StringInput `pulumi:"country"`
	// The postal code for the address.
	PostalCode pulumi.StringInput `pulumi:"postalCode"`
	// The state or province for the address.
	State pulumi.StringInput `pulumi:"state"`
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (i AddressArgs) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

func (i AddressArgs) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i AddressArgs) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput).ToAddressPtrOutputWithContext(ctx)
}

// AddressPtrInput is an input type that accepts AddressArgs, AddressPtr and AddressPtrOutput values.
// You can construct a concrete instance of `AddressPtrInput` via:
//
//	        AddressArgs{...}
//
//	or:
//
//	        nil
type AddressPtrInput interface {
	pulumi.Input

	ToAddressPtrOutput() AddressPtrOutput
	ToAddressPtrOutputWithContext(context.Context) AddressPtrOutput
}

type addressPtrType AddressArgs

func AddressPtr(v *AddressArgs) AddressPtrInput {
	return (*addressPtrType)(v)
}

func (*addressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *addressPtrType) ToAddressPtrOutput() AddressPtrOutput {
	return i.ToAddressPtrOutputWithContext(context.Background())
}

func (i *addressPtrType) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressPtrOutput)
}

// Address information for domain registration.
type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

func (o AddressOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o.ToAddressPtrOutputWithContext(context.Background())
}

func (o AddressOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Address) *Address {
		return &v
	}).(AddressPtrOutput)
}

// First line of an Address.
func (o AddressOutput) Address1() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.Address1 }).(pulumi.StringOutput)
}

// The second line of the Address. Optional.
func (o AddressOutput) Address2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Address) *string { return v.Address2 }).(pulumi.StringPtrOutput)
}

// The city for the address.
func (o AddressOutput) City() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.City }).(pulumi.StringOutput)
}

// The country for the address.
func (o AddressOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.Country }).(pulumi.StringOutput)
}

// The postal code for the address.
func (o AddressOutput) PostalCode() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.PostalCode }).(pulumi.StringOutput)
}

// The state or province for the address.
func (o AddressOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Address) string { return v.State }).(pulumi.StringOutput)
}

type AddressPtrOutput struct{ *pulumi.OutputState }

func (AddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressPtrOutput) ToAddressPtrOutput() AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) ToAddressPtrOutputWithContext(ctx context.Context) AddressPtrOutput {
	return o
}

func (o AddressPtrOutput) Elem() AddressOutput {
	return o.ApplyT(func(v *Address) Address {
		if v != nil {
			return *v
		}
		var ret Address
		return ret
	}).(AddressOutput)
}

// First line of an Address.
func (o AddressPtrOutput) Address1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.Address1
	}).(pulumi.StringPtrOutput)
}

// The second line of the Address. Optional.
func (o AddressPtrOutput) Address2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return v.Address2
	}).(pulumi.StringPtrOutput)
}

// The city for the address.
func (o AddressPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.City
	}).(pulumi.StringPtrOutput)
}

// The country for the address.
func (o AddressPtrOutput) Country() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.Country
	}).(pulumi.StringPtrOutput)
}

// The postal code for the address.
func (o AddressPtrOutput) PostalCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.PostalCode
	}).(pulumi.StringPtrOutput)
}

// The state or province for the address.
func (o AddressPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type Contact struct {
	// Mailing address.
	AddressMailing *Address `pulumi:"addressMailing"`
	// Email address.
	Email string `pulumi:"email"`
	// Fax number.
	Fax *string `pulumi:"fax"`
	// Job title.
	JobTitle *string `pulumi:"jobTitle"`
	// First name.
	NameFirst string `pulumi:"nameFirst"`
	// Last name.
	NameLast string `pulumi:"nameLast"`
	// Middle name.
	NameMiddle *string `pulumi:"nameMiddle"`
	// Organization contact belongs to.
	Organization *string `pulumi:"organization"`
	// Phone number.
	Phone string `pulumi:"phone"`
}

// ContactInput is an input type that accepts ContactArgs and ContactOutput values.
// You can construct a concrete instance of `ContactInput` via:
//
//	ContactArgs{...}
type ContactInput interface {
	pulumi.Input

	ToContactOutput() ContactOutput
	ToContactOutputWithContext(context.Context) ContactOutput
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type ContactArgs struct {
	// Mailing address.
	AddressMailing AddressPtrInput `pulumi:"addressMailing"`
	// Email address.
	Email pulumi.StringInput `pulumi:"email"`
	// Fax number.
	Fax pulumi.StringPtrInput `pulumi:"fax"`
	// Job title.
	JobTitle pulumi.StringPtrInput `pulumi:"jobTitle"`
	// First name.
	NameFirst pulumi.StringInput `pulumi:"nameFirst"`
	// Last name.
	NameLast pulumi.StringInput `pulumi:"nameLast"`
	// Middle name.
	NameMiddle pulumi.StringPtrInput `pulumi:"nameMiddle"`
	// Organization contact belongs to.
	Organization pulumi.StringPtrInput `pulumi:"organization"`
	// Phone number.
	Phone pulumi.StringInput `pulumi:"phone"`
}

func (ContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Contact)(nil)).Elem()
}

func (i ContactArgs) ToContactOutput() ContactOutput {
	return i.ToContactOutputWithContext(context.Background())
}

func (i ContactArgs) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactOutput)
}

// Contact information for domain registration. If 'Domain Privacy' option is not selected then the contact information is made publicly available through the Whois
// directories as per ICANN requirements.
type ContactOutput struct{ *pulumi.OutputState }

func (ContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Contact)(nil)).Elem()
}

func (o ContactOutput) ToContactOutput() ContactOutput {
	return o
}

func (o ContactOutput) ToContactOutputWithContext(ctx context.Context) ContactOutput {
	return o
}

// Mailing address.
func (o ContactOutput) AddressMailing() AddressPtrOutput {
	return o.ApplyT(func(v Contact) *Address { return v.AddressMailing }).(AddressPtrOutput)
}

// Email address.
func (o ContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.Email }).(pulumi.StringOutput)
}

// Fax number.
func (o ContactOutput) Fax() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.Fax }).(pulumi.StringPtrOutput)
}

// Job title.
func (o ContactOutput) JobTitle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.JobTitle }).(pulumi.StringPtrOutput)
}

// First name.
func (o ContactOutput) NameFirst() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.NameFirst }).(pulumi.StringOutput)
}

// Last name.
func (o ContactOutput) NameLast() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.NameLast }).(pulumi.StringOutput)
}

// Middle name.
func (o ContactOutput) NameMiddle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.NameMiddle }).(pulumi.StringPtrOutput)
}

// Organization contact belongs to.
func (o ContactOutput) Organization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Contact) *string { return v.Organization }).(pulumi.StringPtrOutput)
}

// Phone number.
func (o ContactOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v Contact) string { return v.Phone }).(pulumi.StringOutput)
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsent struct {
	// Timestamp when the agreements were accepted.
	AgreedAt *string `pulumi:"agreedAt"`
	// Client IP address.
	AgreedBy *string `pulumi:"agreedBy"`
	// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
	AgreementKeys []string `pulumi:"agreementKeys"`
}

// DomainPurchaseConsentInput is an input type that accepts DomainPurchaseConsentArgs and DomainPurchaseConsentOutput values.
// You can construct a concrete instance of `DomainPurchaseConsentInput` via:
//
//	DomainPurchaseConsentArgs{...}
type DomainPurchaseConsentInput interface {
	pulumi.Input

	ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput
	ToDomainPurchaseConsentOutputWithContext(context.Context) DomainPurchaseConsentOutput
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsentArgs struct {
	// Timestamp when the agreements were accepted.
	AgreedAt pulumi.StringPtrInput `pulumi:"agreedAt"`
	// Client IP address.
	AgreedBy pulumi.StringPtrInput `pulumi:"agreedBy"`
	// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
	AgreementKeys pulumi.StringArrayInput `pulumi:"agreementKeys"`
}

func (DomainPurchaseConsentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPurchaseConsent)(nil)).Elem()
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput {
	return i.ToDomainPurchaseConsentOutputWithContext(context.Background())
}

func (i DomainPurchaseConsentArgs) ToDomainPurchaseConsentOutputWithContext(ctx context.Context) DomainPurchaseConsentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainPurchaseConsentOutput)
}

// Domain purchase consent object, representing acceptance of applicable legal agreements.
type DomainPurchaseConsentOutput struct{ *pulumi.OutputState }

func (DomainPurchaseConsentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainPurchaseConsent)(nil)).Elem()
}

func (o DomainPurchaseConsentOutput) ToDomainPurchaseConsentOutput() DomainPurchaseConsentOutput {
	return o
}

func (o DomainPurchaseConsentOutput) ToDomainPurchaseConsentOutputWithContext(ctx context.Context) DomainPurchaseConsentOutput {
	return o
}

// Timestamp when the agreements were accepted.
func (o DomainPurchaseConsentOutput) AgreedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) *string { return v.AgreedAt }).(pulumi.StringPtrOutput)
}

// Client IP address.
func (o DomainPurchaseConsentOutput) AgreedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) *string { return v.AgreedBy }).(pulumi.StringPtrOutput)
}

// List of applicable legal agreement keys. This list can be retrieved using ListLegalAgreements API under <code>TopLevelDomain</code> resource.
func (o DomainPurchaseConsentOutput) AgreementKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DomainPurchaseConsent) []string { return v.AgreementKeys }).(pulumi.StringArrayOutput)
}

// Details of a hostname derived from a domain.
type HostNameResponse struct {
	// Name of the Azure resource the hostname is assigned to. If it is assigned to a Traffic Manager then it will be the Traffic Manager name otherwise it will be the app name.
	AzureResourceName *string `pulumi:"azureResourceName"`
	// Type of the Azure resource the hostname is assigned to.
	AzureResourceType *string `pulumi:"azureResourceType"`
	// Type of the DNS record.
	CustomHostNameDnsRecordType *string `pulumi:"customHostNameDnsRecordType"`
	// Type of the hostname.
	HostNameType *string `pulumi:"hostNameType"`
	// Name of the hostname.
	Name *string `pulumi:"name"`
	// List of apps the hostname is assigned to. This list will have more than one app only if the hostname is pointing to a Traffic Manager.
	SiteNames []string `pulumi:"siteNames"`
}

// Details of a hostname derived from a domain.
type HostNameResponseOutput struct{ *pulumi.OutputState }

func (HostNameResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameResponse)(nil)).Elem()
}

func (o HostNameResponseOutput) ToHostNameResponseOutput() HostNameResponseOutput {
	return o
}

func (o HostNameResponseOutput) ToHostNameResponseOutputWithContext(ctx context.Context) HostNameResponseOutput {
	return o
}

// Name of the Azure resource the hostname is assigned to. If it is assigned to a Traffic Manager then it will be the Traffic Manager name otherwise it will be the app name.
func (o HostNameResponseOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

// Type of the Azure resource the hostname is assigned to.
func (o HostNameResponseOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

// Type of the DNS record.
func (o HostNameResponseOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

// Type of the hostname.
func (o HostNameResponseOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.HostNameType }).(pulumi.StringPtrOutput)
}

// Name of the hostname.
func (o HostNameResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// List of apps the hostname is assigned to. This list will have more than one app only if the hostname is pointing to a Traffic Manager.
func (o HostNameResponseOutput) SiteNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HostNameResponse) []string { return v.SiteNames }).(pulumi.StringArrayOutput)
}

type HostNameResponseArrayOutput struct{ *pulumi.OutputState }

func (HostNameResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameResponse)(nil)).Elem()
}

func (o HostNameResponseArrayOutput) ToHostNameResponseArrayOutput() HostNameResponseArrayOutput {
	return o
}

func (o HostNameResponseArrayOutput) ToHostNameResponseArrayOutputWithContext(ctx context.Context) HostNameResponseArrayOutput {
	return o
}

func (o HostNameResponseArrayOutput) Index(i pulumi.IntInput) HostNameResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameResponse {
		return vs[0].([]HostNameResponse)[vs[1].(int)]
	}).(HostNameResponseOutput)
}

// Identifies an object.
type NameIdentifierResponse struct {
	// Name of the object.
	Name *string `pulumi:"name"`
}

// Identifies an object.
type NameIdentifierResponseOutput struct{ *pulumi.OutputState }

func (NameIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameIdentifierResponse)(nil)).Elem()
}

func (o NameIdentifierResponseOutput) ToNameIdentifierResponseOutput() NameIdentifierResponseOutput {
	return o
}

func (o NameIdentifierResponseOutput) ToNameIdentifierResponseOutputWithContext(ctx context.Context) NameIdentifierResponseOutput {
	return o
}

// Name of the object.
func (o NameIdentifierResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameIdentifierResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type NameIdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (NameIdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameIdentifierResponse)(nil)).Elem()
}

func (o NameIdentifierResponseArrayOutput) ToNameIdentifierResponseArrayOutput() NameIdentifierResponseArrayOutput {
	return o
}

func (o NameIdentifierResponseArrayOutput) ToNameIdentifierResponseArrayOutputWithContext(ctx context.Context) NameIdentifierResponseArrayOutput {
	return o
}

func (o NameIdentifierResponseArrayOutput) Index(i pulumi.IntInput) NameIdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameIdentifierResponse {
		return vs[0].([]NameIdentifierResponse)[vs[1].(int)]
	}).(NameIdentifierResponseOutput)
}

// Legal agreement for a top level domain.
type TldLegalAgreementResponse struct {
	// Unique identifier for the agreement.
	AgreementKey string `pulumi:"agreementKey"`
	// Agreement details.
	Content string `pulumi:"content"`
	// Agreement title.
	Title string `pulumi:"title"`
	// URL where a copy of the agreement details is hosted.
	Url *string `pulumi:"url"`
}

// Legal agreement for a top level domain.
type TldLegalAgreementResponseOutput struct{ *pulumi.OutputState }

func (TldLegalAgreementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TldLegalAgreementResponse)(nil)).Elem()
}

func (o TldLegalAgreementResponseOutput) ToTldLegalAgreementResponseOutput() TldLegalAgreementResponseOutput {
	return o
}

func (o TldLegalAgreementResponseOutput) ToTldLegalAgreementResponseOutputWithContext(ctx context.Context) TldLegalAgreementResponseOutput {
	return o
}

// Unique identifier for the agreement.
func (o TldLegalAgreementResponseOutput) AgreementKey() pulumi.StringOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) string { return v.AgreementKey }).(pulumi.StringOutput)
}

// Agreement details.
func (o TldLegalAgreementResponseOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) string { return v.Content }).(pulumi.StringOutput)
}

// Agreement title.
func (o TldLegalAgreementResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) string { return v.Title }).(pulumi.StringOutput)
}

// URL where a copy of the agreement details is hosted.
func (o TldLegalAgreementResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TldLegalAgreementResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type TldLegalAgreementResponseArrayOutput struct{ *pulumi.OutputState }

func (TldLegalAgreementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TldLegalAgreementResponse)(nil)).Elem()
}

func (o TldLegalAgreementResponseArrayOutput) ToTldLegalAgreementResponseArrayOutput() TldLegalAgreementResponseArrayOutput {
	return o
}

func (o TldLegalAgreementResponseArrayOutput) ToTldLegalAgreementResponseArrayOutputWithContext(ctx context.Context) TldLegalAgreementResponseArrayOutput {
	return o
}

func (o TldLegalAgreementResponseArrayOutput) Index(i pulumi.IntInput) TldLegalAgreementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TldLegalAgreementResponse {
		return vs[0].([]TldLegalAgreementResponse)[vs[1].(int)]
	}).(TldLegalAgreementResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressPtrOutput{})
	pulumi.RegisterOutputType(ContactOutput{})
	pulumi.RegisterOutputType(DomainPurchaseConsentOutput{})
	pulumi.RegisterOutputType(HostNameResponseOutput{})
	pulumi.RegisterOutputType(HostNameResponseArrayOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseOutput{})
	pulumi.RegisterOutputType(NameIdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseOutput{})
	pulumi.RegisterOutputType(TldLegalAgreementResponseArrayOutput{})
}
