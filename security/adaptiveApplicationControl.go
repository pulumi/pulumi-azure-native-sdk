// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2020-01-01.
type AdaptiveApplicationControl struct {
	pulumi.CustomResourceState

	// The configuration status of the machines group or machine or rule
	ConfigurationStatus pulumi.StringOutput `pulumi:"configurationStatus"`
	// The application control policy enforcement/protection mode of the machine group
	EnforcementMode pulumi.StringPtrOutput                                    `pulumi:"enforcementMode"`
	Issues          AdaptiveApplicationControlIssueSummaryResponseArrayOutput `pulumi:"issues"`
	// Location where the resource is stored
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name                pulumi.StringOutput                   `pulumi:"name"`
	PathRecommendations PathRecommendationResponseArrayOutput `pulumi:"pathRecommendations"`
	// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
	ProtectionMode ProtectionModeResponsePtrOutput `pulumi:"protectionMode"`
	// The initial recommendation status of the machine group or machine
	RecommendationStatus pulumi.StringOutput `pulumi:"recommendationStatus"`
	// The source type of the machine group
	SourceSystem pulumi.StringOutput `pulumi:"sourceSystem"`
	// Resource type
	Type              pulumi.StringOutput                 `pulumi:"type"`
	VmRecommendations VmRecommendationResponseArrayOutput `pulumi:"vmRecommendations"`
}

// NewAdaptiveApplicationControl registers a new resource with the given unique name, arguments, and options.
func NewAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, args *AdaptiveApplicationControlArgs, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AscLocation == nil {
		return nil, errors.New("invalid value for required argument 'AscLocation'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security/v20150601preview:AdaptiveApplicationControl"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101:AdaptiveApplicationControl"),
		},
	})
	opts = append(opts, aliases)
	var resource AdaptiveApplicationControl
	err := ctx.RegisterResource("azure-native:security:AdaptiveApplicationControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdaptiveApplicationControl gets an existing AdaptiveApplicationControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdaptiveApplicationControlState, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	var resource AdaptiveApplicationControl
	err := ctx.ReadResource("azure-native:security:AdaptiveApplicationControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdaptiveApplicationControl resources.
type adaptiveApplicationControlState struct {
}

type AdaptiveApplicationControlState struct {
}

func (AdaptiveApplicationControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveApplicationControlState)(nil)).Elem()
}

type adaptiveApplicationControlArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// The application control policy enforcement/protection mode of the machine group
	EnforcementMode *string `pulumi:"enforcementMode"`
	// Name of an application control machine group
	GroupName           *string              `pulumi:"groupName"`
	PathRecommendations []PathRecommendation `pulumi:"pathRecommendations"`
	// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
	ProtectionMode    *ProtectionMode    `pulumi:"protectionMode"`
	VmRecommendations []VmRecommendation `pulumi:"vmRecommendations"`
}

// The set of arguments for constructing a AdaptiveApplicationControl resource.
type AdaptiveApplicationControlArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput
	// The application control policy enforcement/protection mode of the machine group
	EnforcementMode pulumi.StringPtrInput
	// Name of an application control machine group
	GroupName           pulumi.StringPtrInput
	PathRecommendations PathRecommendationArrayInput
	// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
	ProtectionMode    ProtectionModePtrInput
	VmRecommendations VmRecommendationArrayInput
}

func (AdaptiveApplicationControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveApplicationControlArgs)(nil)).Elem()
}

type AdaptiveApplicationControlInput interface {
	pulumi.Input

	ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput
	ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput
}

func (*AdaptiveApplicationControl) ElementType() reflect.Type {
	return reflect.TypeOf((**AdaptiveApplicationControl)(nil)).Elem()
}

func (i *AdaptiveApplicationControl) ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput {
	return i.ToAdaptiveApplicationControlOutputWithContext(context.Background())
}

func (i *AdaptiveApplicationControl) ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdaptiveApplicationControlOutput)
}

type AdaptiveApplicationControlOutput struct{ *pulumi.OutputState }

func (AdaptiveApplicationControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdaptiveApplicationControl)(nil)).Elem()
}

func (o AdaptiveApplicationControlOutput) ToAdaptiveApplicationControlOutput() AdaptiveApplicationControlOutput {
	return o
}

func (o AdaptiveApplicationControlOutput) ToAdaptiveApplicationControlOutputWithContext(ctx context.Context) AdaptiveApplicationControlOutput {
	return o
}

// The configuration status of the machines group or machine or rule
func (o AdaptiveApplicationControlOutput) ConfigurationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.ConfigurationStatus }).(pulumi.StringOutput)
}

// The application control policy enforcement/protection mode of the machine group
func (o AdaptiveApplicationControlOutput) EnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringPtrOutput { return v.EnforcementMode }).(pulumi.StringPtrOutput)
}

func (o AdaptiveApplicationControlOutput) Issues() AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) AdaptiveApplicationControlIssueSummaryResponseArrayOutput {
		return v.Issues
	}).(AdaptiveApplicationControlIssueSummaryResponseArrayOutput)
}

// Location where the resource is stored
func (o AdaptiveApplicationControlOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name
func (o AdaptiveApplicationControlOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) PathRecommendations() PathRecommendationResponseArrayOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) PathRecommendationResponseArrayOutput {
		return v.PathRecommendations
	}).(PathRecommendationResponseArrayOutput)
}

// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
func (o AdaptiveApplicationControlOutput) ProtectionMode() ProtectionModeResponsePtrOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) ProtectionModeResponsePtrOutput { return v.ProtectionMode }).(ProtectionModeResponsePtrOutput)
}

// The initial recommendation status of the machine group or machine
func (o AdaptiveApplicationControlOutput) RecommendationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.RecommendationStatus }).(pulumi.StringOutput)
}

// The source type of the machine group
func (o AdaptiveApplicationControlOutput) SourceSystem() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.SourceSystem }).(pulumi.StringOutput)
}

// Resource type
func (o AdaptiveApplicationControlOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AdaptiveApplicationControlOutput) VmRecommendations() VmRecommendationResponseArrayOutput {
	return o.ApplyT(func(v *AdaptiveApplicationControl) VmRecommendationResponseArrayOutput { return v.VmRecommendations }).(VmRecommendationResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AdaptiveApplicationControlOutput{})
}