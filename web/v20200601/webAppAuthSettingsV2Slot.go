// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAuthSettingsV2Slot struct {
	pulumi.CustomResourceState

	GlobalValidation  GlobalValidationResponsePtrOutput  `pulumi:"globalValidation"`
	HttpSettings      HttpSettingsResponsePtrOutput      `pulumi:"httpSettings"`
	IdentityProviders IdentityProvidersResponsePtrOutput `pulumi:"identityProviders"`
	// Kind of resource.
	Kind  pulumi.StringPtrOutput `pulumi:"kind"`
	Login LoginResponsePtrOutput `pulumi:"login"`
	// Resource Name.
	Name     pulumi.StringOutput           `pulumi:"name"`
	Platform AuthPlatformResponsePtrOutput `pulumi:"platform"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppAuthSettingsV2Slot registers a new resource with the given unique name, arguments, and options.
func NewWebAppAuthSettingsV2Slot(ctx *pulumi.Context,
	name string, args *WebAppAuthSettingsV2SlotArgs, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsV2Slot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAuthSettingsV2Slot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAuthSettingsV2Slot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppAuthSettingsV2Slot
	err := ctx.RegisterResource("azure-native:web/v20200601:WebAppAuthSettingsV2Slot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppAuthSettingsV2Slot gets an existing WebAppAuthSettingsV2Slot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppAuthSettingsV2Slot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAuthSettingsV2SlotState, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsV2Slot, error) {
	var resource WebAppAuthSettingsV2Slot
	err := ctx.ReadResource("azure-native:web/v20200601:WebAppAuthSettingsV2Slot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppAuthSettingsV2Slot resources.
type webAppAuthSettingsV2SlotState struct {
}

type WebAppAuthSettingsV2SlotState struct {
}

func (WebAppAuthSettingsV2SlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsV2SlotState)(nil)).Elem()
}

type webAppAuthSettingsV2SlotArgs struct {
	GlobalValidation  *GlobalValidation  `pulumi:"globalValidation"`
	HttpSettings      *HttpSettings      `pulumi:"httpSettings"`
	IdentityProviders *IdentityProviders `pulumi:"identityProviders"`
	// Kind of resource.
	Kind  *string `pulumi:"kind"`
	Login *Login  `pulumi:"login"`
	// Name of web app.
	Name     string        `pulumi:"name"`
	Platform *AuthPlatform `pulumi:"platform"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppAuthSettingsV2Slot resource.
type WebAppAuthSettingsV2SlotArgs struct {
	GlobalValidation  GlobalValidationPtrInput
	HttpSettings      HttpSettingsPtrInput
	IdentityProviders IdentityProvidersPtrInput
	// Kind of resource.
	Kind  pulumi.StringPtrInput
	Login LoginPtrInput
	// Name of web app.
	Name     pulumi.StringInput
	Platform AuthPlatformPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput
}

func (WebAppAuthSettingsV2SlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsV2SlotArgs)(nil)).Elem()
}

type WebAppAuthSettingsV2SlotInput interface {
	pulumi.Input

	ToWebAppAuthSettingsV2SlotOutput() WebAppAuthSettingsV2SlotOutput
	ToWebAppAuthSettingsV2SlotOutputWithContext(ctx context.Context) WebAppAuthSettingsV2SlotOutput
}

func (*WebAppAuthSettingsV2Slot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsV2Slot)(nil)).Elem()
}

func (i *WebAppAuthSettingsV2Slot) ToWebAppAuthSettingsV2SlotOutput() WebAppAuthSettingsV2SlotOutput {
	return i.ToWebAppAuthSettingsV2SlotOutputWithContext(context.Background())
}

func (i *WebAppAuthSettingsV2Slot) ToWebAppAuthSettingsV2SlotOutputWithContext(ctx context.Context) WebAppAuthSettingsV2SlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAuthSettingsV2SlotOutput)
}

type WebAppAuthSettingsV2SlotOutput struct{ *pulumi.OutputState }

func (WebAppAuthSettingsV2SlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsV2Slot)(nil)).Elem()
}

func (o WebAppAuthSettingsV2SlotOutput) ToWebAppAuthSettingsV2SlotOutput() WebAppAuthSettingsV2SlotOutput {
	return o
}

func (o WebAppAuthSettingsV2SlotOutput) ToWebAppAuthSettingsV2SlotOutputWithContext(ctx context.Context) WebAppAuthSettingsV2SlotOutput {
	return o
}

func (o WebAppAuthSettingsV2SlotOutput) GlobalValidation() GlobalValidationResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) GlobalValidationResponsePtrOutput { return v.GlobalValidation }).(GlobalValidationResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) HttpSettings() HttpSettingsResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) HttpSettingsResponsePtrOutput { return v.HttpSettings }).(HttpSettingsResponsePtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) IdentityProviders() IdentityProvidersResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) IdentityProvidersResponsePtrOutput { return v.IdentityProviders }).(IdentityProvidersResponsePtrOutput)
}

// Kind of resource.
func (o WebAppAuthSettingsV2SlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Login() LoginResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) LoginResponsePtrOutput { return v.Login }).(LoginResponsePtrOutput)
}

// Resource Name.
func (o WebAppAuthSettingsV2SlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppAuthSettingsV2SlotOutput) Platform() AuthPlatformResponsePtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) AuthPlatformResponsePtrOutput { return v.Platform }).(AuthPlatformResponsePtrOutput)
}

// Resource type.
func (o WebAppAuthSettingsV2SlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsV2Slot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAuthSettingsV2SlotOutput{})
}