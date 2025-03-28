// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230822

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Organization resource.
type Organization struct {
	pulumi.CustomResourceState

	// The creation time of the resource.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Location of Organization resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Confluent offer detail
	OfferDetail OfferDetailResponseOutput `pulumi:"offerDetail"`
	// Id of the Confluent organization.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Provision states for confluent RP
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SSO url for the Confluent organization.
	SsoUrl pulumi.StringOutput `pulumi:"ssoUrl"`
	// Metadata pertaining to creation and last modification of the resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Organization resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Subscriber detail
	UserDetail UserDetailResponseOutput `pulumi:"userDetail"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OfferDetail == nil {
		return nil, errors.New("invalid value for required argument 'OfferDetail'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserDetail == nil {
		return nil, errors.New("invalid value for required argument 'UserDetail'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:confluent/v20200301:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20200301preview:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20210301preview:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20210901preview:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20211201:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20240213:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent/v20240701:Organization"),
		},
		{
			Type: pulumi.String("azure-native:confluent:Organization"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Organization
	err := ctx.RegisterResource("azure-native:confluent/v20230822:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("azure-native:confluent/v20230822:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
}

type OrganizationState struct {
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// Link an existing Confluent organization
	LinkOrganization *LinkOrganization `pulumi:"linkOrganization"`
	// Location of Organization resource
	Location *string `pulumi:"location"`
	// Confluent offer detail
	OfferDetail OfferDetail `pulumi:"offerDetail"`
	// Organization resource name
	OrganizationName *string `pulumi:"organizationName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Organization resource tags
	Tags map[string]string `pulumi:"tags"`
	// Subscriber detail
	UserDetail UserDetail `pulumi:"userDetail"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Link an existing Confluent organization
	LinkOrganization LinkOrganizationPtrInput
	// Location of Organization resource
	Location pulumi.StringPtrInput
	// Confluent offer detail
	OfferDetail OfferDetailInput
	// Organization resource name
	OrganizationName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Organization resource tags
	Tags pulumi.StringMapInput
	// Subscriber detail
	UserDetail UserDetailInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (*Organization) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (i *Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

// The creation time of the resource.
func (o OrganizationOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Location of Organization resource
func (o OrganizationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o OrganizationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Confluent offer detail
func (o OrganizationOutput) OfferDetail() OfferDetailResponseOutput {
	return o.ApplyT(func(v *Organization) OfferDetailResponseOutput { return v.OfferDetail }).(OfferDetailResponseOutput)
}

// Id of the Confluent organization.
func (o OrganizationOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Provision states for confluent RP
func (o OrganizationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SSO url for the Confluent organization.
func (o OrganizationOutput) SsoUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.SsoUrl }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource
func (o OrganizationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Organization) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Organization resource tags
func (o OrganizationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o OrganizationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Organization) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Subscriber detail
func (o OrganizationOutput) UserDetail() UserDetailResponseOutput {
	return o.ApplyT(func(v *Organization) UserDetailResponseOutput { return v.UserDetail }).(UserDetailResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationOutput{})
}
