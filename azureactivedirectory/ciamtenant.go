// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azureactivedirectory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Azure AD for customers resource.
//
// Uses Azure REST API version 2023-05-17-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-17-preview.
type CIAMTenant struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The type of billing. Will be MAU for all new customers. Cannot be changed if value is 'MAU'. Learn more about Azure AD for customers billing at [aka.ms/b2cBilling](https://aka.ms/b2cbilling).
	BillingType pulumi.StringOutput `pulumi:"billingType"`
	// These properties are used to create the Azure AD for customers tenant. These properties are not part of the Azure resource.
	CreateTenantProperties CreateCIAMTenantPropertiesResponseOutput `pulumi:"createTenantProperties"`
	// The domain name of the tenant
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The data from which the billing type took effect
	EffectiveStartDateUtc pulumi.StringOutput `pulumi:"effectiveStartDateUtc"`
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/ciam-data-location) for more information.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Azure AD for customers tenant resource.
	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SKU properties of the Azure AD for customers tenant. Learn more about Azure AD for customers billing at [https://aka.ms/ciambilling](https://aka.ms/ciambilling).
	Sku CIAMResourceSKUResponseOutput `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource Tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// An identifier of the Azure AD for customers tenant.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the Azure AD for customers tenant resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCIAMTenant registers a new resource with the given unique name, arguments, and options.
func NewCIAMTenant(ctx *pulumi.Context,
	name string, args *CIAMTenantArgs, opts ...pulumi.ResourceOption) (*CIAMTenant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreateTenantProperties == nil {
		return nil, errors.New("invalid value for required argument 'CreateTenantProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azureactivedirectory/v20230517preview:CIAMTenant"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CIAMTenant
	err := ctx.RegisterResource("azure-native:azureactivedirectory:CIAMTenant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCIAMTenant gets an existing CIAMTenant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCIAMTenant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CIAMTenantState, opts ...pulumi.ResourceOption) (*CIAMTenant, error) {
	var resource CIAMTenant
	err := ctx.ReadResource("azure-native:azureactivedirectory:CIAMTenant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CIAMTenant resources.
type ciamtenantState struct {
}

type CIAMTenantState struct {
}

func (CIAMTenantState) ElementType() reflect.Type {
	return reflect.TypeOf((*ciamtenantState)(nil)).Elem()
}

type ciamtenantArgs struct {
	// These properties are used to create the Azure AD for customers tenant. These properties are not part of the Azure resource.
	CreateTenantProperties CreateCIAMTenantProperties `pulumi:"createTenantProperties"`
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/ciam-data-location) for more information.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The initial sub domain of the tenant.
	ResourceName *string `pulumi:"resourceName"`
	// SKU properties of the Azure AD for customers tenant. Learn more about Azure AD for customers billing at [https://aka.ms/ciambilling](https://aka.ms/ciambilling).
	Sku CIAMResourceSKU `pulumi:"sku"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
	// An identifier of the Azure AD for customers tenant.
	TenantId *string `pulumi:"tenantId"`
}

// The set of arguments for constructing a CIAMTenant resource.
type CIAMTenantArgs struct {
	// These properties are used to create the Azure AD for customers tenant. These properties are not part of the Azure resource.
	CreateTenantProperties CreateCIAMTenantPropertiesInput
	// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/ciam-data-location) for more information.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The initial sub domain of the tenant.
	ResourceName pulumi.StringPtrInput
	// SKU properties of the Azure AD for customers tenant. Learn more about Azure AD for customers billing at [https://aka.ms/ciambilling](https://aka.ms/ciambilling).
	Sku CIAMResourceSKUInput
	// Resource Tags
	Tags pulumi.StringMapInput
	// An identifier of the Azure AD for customers tenant.
	TenantId pulumi.StringPtrInput
}

func (CIAMTenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ciamtenantArgs)(nil)).Elem()
}

type CIAMTenantInput interface {
	pulumi.Input

	ToCIAMTenantOutput() CIAMTenantOutput
	ToCIAMTenantOutputWithContext(ctx context.Context) CIAMTenantOutput
}

func (*CIAMTenant) ElementType() reflect.Type {
	return reflect.TypeOf((**CIAMTenant)(nil)).Elem()
}

func (i *CIAMTenant) ToCIAMTenantOutput() CIAMTenantOutput {
	return i.ToCIAMTenantOutputWithContext(context.Background())
}

func (i *CIAMTenant) ToCIAMTenantOutputWithContext(ctx context.Context) CIAMTenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CIAMTenantOutput)
}

type CIAMTenantOutput struct{ *pulumi.OutputState }

func (CIAMTenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CIAMTenant)(nil)).Elem()
}

func (o CIAMTenantOutput) ToCIAMTenantOutput() CIAMTenantOutput {
	return o
}

func (o CIAMTenantOutput) ToCIAMTenantOutputWithContext(ctx context.Context) CIAMTenantOutput {
	return o
}

// The Azure API version of the resource.
func (o CIAMTenantOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The type of billing. Will be MAU for all new customers. Cannot be changed if value is 'MAU'. Learn more about Azure AD for customers billing at [aka.ms/b2cBilling](https://aka.ms/b2cbilling).
func (o CIAMTenantOutput) BillingType() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.BillingType }).(pulumi.StringOutput)
}

// These properties are used to create the Azure AD for customers tenant. These properties are not part of the Azure resource.
func (o CIAMTenantOutput) CreateTenantProperties() CreateCIAMTenantPropertiesResponseOutput {
	return o.ApplyT(func(v *CIAMTenant) CreateCIAMTenantPropertiesResponseOutput { return v.CreateTenantProperties }).(CreateCIAMTenantPropertiesResponseOutput)
}

// The domain name of the tenant
func (o CIAMTenantOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The data from which the billing type took effect
func (o CIAMTenantOutput) EffectiveStartDateUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.EffectiveStartDateUtc }).(pulumi.StringOutput)
}

// The location in which the resource is hosted and data resides. Can be one of 'United States', 'Europe', 'Asia Pacific', or 'Australia'. Refer to [this documentation](https://aka.ms/ciam-data-location) for more information.
func (o CIAMTenantOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the Azure AD for customers tenant resource.
func (o CIAMTenantOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CIAMTenantOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SKU properties of the Azure AD for customers tenant. Learn more about Azure AD for customers billing at [https://aka.ms/ciambilling](https://aka.ms/ciambilling).
func (o CIAMTenantOutput) Sku() CIAMResourceSKUResponseOutput {
	return o.ApplyT(func(v *CIAMTenant) CIAMResourceSKUResponseOutput { return v.Sku }).(CIAMResourceSKUResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CIAMTenantOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CIAMTenant) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource Tags
func (o CIAMTenantOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// An identifier of the Azure AD for customers tenant.
func (o CIAMTenantOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the Azure AD for customers tenant resource.
func (o CIAMTenantOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CIAMTenant) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CIAMTenantOutput{})
}
