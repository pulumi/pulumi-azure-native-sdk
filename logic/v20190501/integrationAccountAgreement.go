// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration account agreement.
type IntegrationAccountAgreement struct {
	pulumi.CustomResourceState

	// The agreement type.
	AgreementType pulumi.StringOutput `pulumi:"agreementType"`
	// The changed time.
	ChangedTime pulumi.StringOutput `pulumi:"changedTime"`
	// The agreement content.
	Content AgreementContentResponseOutput `pulumi:"content"`
	// The created time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityResponseOutput `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringOutput `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity BusinessIdentityResponseOutput `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringOutput `pulumi:"hostPartner"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The metadata.
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationAccountAgreement registers a new resource with the given unique name, arguments, and options.
func NewIntegrationAccountAgreement(ctx *pulumi.Context,
	name string, args *IntegrationAccountAgreementArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountAgreement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgreementType == nil {
		return nil, errors.New("invalid value for required argument 'AgreementType'")
	}
	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.GuestIdentity == nil {
		return nil, errors.New("invalid value for required argument 'GuestIdentity'")
	}
	if args.GuestPartner == nil {
		return nil, errors.New("invalid value for required argument 'GuestPartner'")
	}
	if args.HostIdentity == nil {
		return nil, errors.New("invalid value for required argument 'HostIdentity'")
	}
	if args.HostPartner == nil {
		return nil, errors.New("invalid value for required argument 'HostPartner'")
	}
	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountAgreement"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountAgreement"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountAgreement
	err := ctx.RegisterResource("azure-native:logic/v20190501:IntegrationAccountAgreement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationAccountAgreement gets an existing IntegrationAccountAgreement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationAccountAgreement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountAgreementState, opts ...pulumi.ResourceOption) (*IntegrationAccountAgreement, error) {
	var resource IntegrationAccountAgreement
	err := ctx.ReadResource("azure-native:logic/v20190501:IntegrationAccountAgreement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationAccountAgreement resources.
type integrationAccountAgreementState struct {
}

type IntegrationAccountAgreementState struct {
}

func (IntegrationAccountAgreementState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAgreementState)(nil)).Elem()
}

type integrationAccountAgreementArgs struct {
	// The integration account agreement name.
	AgreementName *string `pulumi:"agreementName"`
	// The agreement type.
	AgreementType AgreementType `pulumi:"agreementType"`
	// The agreement content.
	Content AgreementContent `pulumi:"content"`
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentity `pulumi:"guestIdentity"`
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner string `pulumi:"guestPartner"`
	// The business identity of the host partner.
	HostIdentity BusinessIdentity `pulumi:"hostIdentity"`
	// The integration account partner that is set as host partner for this agreement.
	HostPartner string `pulumi:"hostPartner"`
	// The integration account name.
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationAccountAgreement resource.
type IntegrationAccountAgreementArgs struct {
	// The integration account agreement name.
	AgreementName pulumi.StringPtrInput
	// The agreement type.
	AgreementType AgreementTypeInput
	// The agreement content.
	Content AgreementContentInput
	// The business identity of the guest partner.
	GuestIdentity BusinessIdentityInput
	// The integration account partner that is set as guest partner for this agreement.
	GuestPartner pulumi.StringInput
	// The business identity of the host partner.
	HostIdentity BusinessIdentityInput
	// The integration account partner that is set as host partner for this agreement.
	HostPartner pulumi.StringInput
	// The integration account name.
	IntegrationAccountName pulumi.StringInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The metadata.
	Metadata pulumi.Input
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationAccountAgreementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountAgreementArgs)(nil)).Elem()
}

type IntegrationAccountAgreementInput interface {
	pulumi.Input

	ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput
	ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput
}

func (*IntegrationAccountAgreement) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAgreement)(nil)).Elem()
}

func (i *IntegrationAccountAgreement) ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput {
	return i.ToIntegrationAccountAgreementOutputWithContext(context.Background())
}

func (i *IntegrationAccountAgreement) ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountAgreementOutput)
}

type IntegrationAccountAgreementOutput struct{ *pulumi.OutputState }

func (IntegrationAccountAgreementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountAgreement)(nil)).Elem()
}

func (o IntegrationAccountAgreementOutput) ToIntegrationAccountAgreementOutput() IntegrationAccountAgreementOutput {
	return o
}

func (o IntegrationAccountAgreementOutput) ToIntegrationAccountAgreementOutputWithContext(ctx context.Context) IntegrationAccountAgreementOutput {
	return o
}

// The agreement type.
func (o IntegrationAccountAgreementOutput) AgreementType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.AgreementType }).(pulumi.StringOutput)
}

// The changed time.
func (o IntegrationAccountAgreementOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

// The agreement content.
func (o IntegrationAccountAgreementOutput) Content() AgreementContentResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) AgreementContentResponseOutput { return v.Content }).(AgreementContentResponseOutput)
}

// The created time.
func (o IntegrationAccountAgreementOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The business identity of the guest partner.
func (o IntegrationAccountAgreementOutput) GuestIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) BusinessIdentityResponseOutput { return v.GuestIdentity }).(BusinessIdentityResponseOutput)
}

// The integration account partner that is set as guest partner for this agreement.
func (o IntegrationAccountAgreementOutput) GuestPartner() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.GuestPartner }).(pulumi.StringOutput)
}

// The business identity of the host partner.
func (o IntegrationAccountAgreementOutput) HostIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) BusinessIdentityResponseOutput { return v.HostIdentity }).(BusinessIdentityResponseOutput)
}

// The integration account partner that is set as host partner for this agreement.
func (o IntegrationAccountAgreementOutput) HostPartner() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.HostPartner }).(pulumi.StringOutput)
}

// The resource location.
func (o IntegrationAccountAgreementOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The metadata.
func (o IntegrationAccountAgreementOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.AnyOutput { return v.Metadata }).(pulumi.AnyOutput)
}

// Gets the resource name.
func (o IntegrationAccountAgreementOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource tags.
func (o IntegrationAccountAgreementOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o IntegrationAccountAgreementOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationAccountAgreement) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountAgreementOutput{})
}