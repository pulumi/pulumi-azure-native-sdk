// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Response to get user settings
type UserSettings struct {
	pulumi.CustomResourceState

	// The cloud shell user settings properties.
	Properties UserPropertiesResponseOutput `pulumi:"properties"`
}

// NewUserSettings registers a new resource with the given unique name, arguments, and options.
func NewUserSettings(ctx *pulumi.Context,
	name string, args *UserSettingsArgs, opts ...pulumi.ResourceOption) (*UserSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:portal:UserSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource UserSettings
	err := ctx.RegisterResource("azure-native:portal/v20181001:UserSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserSettings gets an existing UserSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserSettingsState, opts ...pulumi.ResourceOption) (*UserSettings, error) {
	var resource UserSettings
	err := ctx.ReadResource("azure-native:portal/v20181001:UserSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserSettings resources.
type userSettingsState struct {
}

type UserSettingsState struct {
}

func (UserSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsState)(nil)).Elem()
}

type userSettingsArgs struct {
	// The cloud shell user settings properties.
	Properties UserProperties `pulumi:"properties"`
	// The name of the user settings
	UserSettingsName *string `pulumi:"userSettingsName"`
}

// The set of arguments for constructing a UserSettings resource.
type UserSettingsArgs struct {
	// The cloud shell user settings properties.
	Properties UserPropertiesInput
	// The name of the user settings
	UserSettingsName pulumi.StringPtrInput
}

func (UserSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userSettingsArgs)(nil)).Elem()
}

type UserSettingsInput interface {
	pulumi.Input

	ToUserSettingsOutput() UserSettingsOutput
	ToUserSettingsOutputWithContext(ctx context.Context) UserSettingsOutput
}

func (*UserSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSettings)(nil)).Elem()
}

func (i *UserSettings) ToUserSettingsOutput() UserSettingsOutput {
	return i.ToUserSettingsOutputWithContext(context.Background())
}

func (i *UserSettings) ToUserSettingsOutputWithContext(ctx context.Context) UserSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserSettingsOutput)
}

type UserSettingsOutput struct{ *pulumi.OutputState }

func (UserSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserSettings)(nil)).Elem()
}

func (o UserSettingsOutput) ToUserSettingsOutput() UserSettingsOutput {
	return o
}

func (o UserSettingsOutput) ToUserSettingsOutputWithContext(ctx context.Context) UserSettingsOutput {
	return o
}

// The cloud shell user settings properties.
func (o UserSettingsOutput) Properties() UserPropertiesResponseOutput {
	return o.ApplyT(func(v *UserSettings) UserPropertiesResponseOutput { return v.Properties }).(UserPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(UserSettingsOutput{})
}