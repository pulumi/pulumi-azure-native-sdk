// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents settings of an environment, from which environment instances would be created
type EnvironmentSetting struct {
	pulumi.CustomResourceState

	// Describes the user's progress in configuring their environment setting
	ConfigurationState pulumi.StringPtrOutput `pulumi:"configurationState"`
	// Describes the environment and its resource settings
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Time when the template VM was last changed.
	LastChanged pulumi.StringOutput `pulumi:"lastChanged"`
	// Time when the template VM was last sent for publishing.
	LastPublished pulumi.StringOutput `pulumi:"lastPublished"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Describes the readiness of this environment setting
	PublishingState pulumi.StringOutput `pulumi:"publishingState"`
	// The resource specific settings
	ResourceSettings ResourceSettingsResponseOutput `pulumi:"resourceSettings"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Brief title describing the environment and its resource settings
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
}

// NewEnvironmentSetting registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentSetting(ctx *pulumi.Context,
	name string, args *EnvironmentSettingArgs, opts ...pulumi.ResourceOption) (*EnvironmentSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabAccountName == nil {
		return nil, errors.New("invalid value for required argument 'LabAccountName'")
	}
	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceSettings == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSettings'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices:EnvironmentSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentSetting
	err := ctx.RegisterResource("azure-native:labservices/v20181015:EnvironmentSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentSetting gets an existing EnvironmentSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentSettingState, opts ...pulumi.ResourceOption) (*EnvironmentSetting, error) {
	var resource EnvironmentSetting
	err := ctx.ReadResource("azure-native:labservices/v20181015:EnvironmentSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentSetting resources.
type environmentSettingState struct {
}

type EnvironmentSettingState struct {
}

func (EnvironmentSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSettingState)(nil)).Elem()
}

type environmentSettingArgs struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState *string `pulumi:"configurationState"`
	// Describes the environment and its resource settings
	Description *string `pulumi:"description"`
	// The name of the environment Setting.
	EnvironmentSettingName *string `pulumi:"environmentSettingName"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource specific settings
	ResourceSettings ResourceSettings `pulumi:"resourceSettings"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Brief title describing the environment and its resource settings
	Title *string `pulumi:"title"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

// The set of arguments for constructing a EnvironmentSetting resource.
type EnvironmentSettingArgs struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState pulumi.StringPtrInput
	// Describes the environment and its resource settings
	Description pulumi.StringPtrInput
	// The name of the environment Setting.
	EnvironmentSettingName pulumi.StringPtrInput
	// The name of the lab Account.
	LabAccountName pulumi.StringInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource specific settings
	ResourceSettings ResourceSettingsInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// Brief title describing the environment and its resource settings
	Title pulumi.StringPtrInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (EnvironmentSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSettingArgs)(nil)).Elem()
}

type EnvironmentSettingInput interface {
	pulumi.Input

	ToEnvironmentSettingOutput() EnvironmentSettingOutput
	ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput
}

func (*EnvironmentSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSetting)(nil)).Elem()
}

func (i *EnvironmentSetting) ToEnvironmentSettingOutput() EnvironmentSettingOutput {
	return i.ToEnvironmentSettingOutputWithContext(context.Background())
}

func (i *EnvironmentSetting) ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSettingOutput)
}

type EnvironmentSettingOutput struct{ *pulumi.OutputState }

func (EnvironmentSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSetting)(nil)).Elem()
}

func (o EnvironmentSettingOutput) ToEnvironmentSettingOutput() EnvironmentSettingOutput {
	return o
}

func (o EnvironmentSettingOutput) ToEnvironmentSettingOutputWithContext(ctx context.Context) EnvironmentSettingOutput {
	return o
}

// Describes the user's progress in configuring their environment setting
func (o EnvironmentSettingOutput) ConfigurationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.ConfigurationState }).(pulumi.StringPtrOutput)
}

// Describes the environment and its resource settings
func (o EnvironmentSettingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Time when the template VM was last changed.
func (o EnvironmentSettingOutput) LastChanged() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.LastChanged }).(pulumi.StringOutput)
}

// Time when the template VM was last sent for publishing.
func (o EnvironmentSettingOutput) LastPublished() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.LastPublished }).(pulumi.StringOutput)
}

// The details of the latest operation. ex: status, error
func (o EnvironmentSettingOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *EnvironmentSetting) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

// The location of the resource.
func (o EnvironmentSettingOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o EnvironmentSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning status of the resource.
func (o EnvironmentSettingOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Describes the readiness of this environment setting
func (o EnvironmentSettingOutput) PublishingState() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.PublishingState }).(pulumi.StringOutput)
}

// The resource specific settings
func (o EnvironmentSettingOutput) ResourceSettings() ResourceSettingsResponseOutput {
	return o.ApplyT(func(v *EnvironmentSetting) ResourceSettingsResponseOutput { return v.ResourceSettings }).(ResourceSettingsResponseOutput)
}

// The tags of the resource.
func (o EnvironmentSettingOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Brief title describing the environment and its resource settings
func (o EnvironmentSettingOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.Title }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o EnvironmentSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o EnvironmentSettingOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSetting) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentSettingOutput{})
}