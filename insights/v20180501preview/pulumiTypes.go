// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Static definitions of the ProactiveDetection configuration rule (same values for all components).
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions struct {
	// The rule description
	Description *string `pulumi:"description"`
	// The rule name as it is displayed in UI
	DisplayName *string `pulumi:"displayName"`
	// URL which displays additional info about the proactive detection rule
	HelpUrl *string `pulumi:"helpUrl"`
	// A flag indicating whether the rule is enabled by default
	IsEnabledByDefault *bool `pulumi:"isEnabledByDefault"`
	// A flag indicating whether the rule is hidden (from the UI)
	IsHidden *bool `pulumi:"isHidden"`
	// A flag indicating whether the rule is in preview
	IsInPreview *bool `pulumi:"isInPreview"`
	// The rule name
	Name *string `pulumi:"name"`
	// A flag indicating whether email notifications are supported for detections for this rule
	SupportsEmailNotifications *bool `pulumi:"supportsEmailNotifications"`
}

// Static definitions of the ProactiveDetection configuration rule (same values for all components).
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput {
	return o
}

// The rule description
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The rule name as it is displayed in UI
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// URL which displays additional info about the proactive detection rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether the rule is enabled by default
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is hidden (from the UI)
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is in preview
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

// The rule name
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether email notifications are supported for detections for this rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) Elem() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions
		return ret
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput)
}

// The rule description
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The rule name as it is displayed in UI
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// URL which displays additional info about the proactive detection rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether the rule is enabled by default
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is hidden (from the UI)
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is in preview
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

// The rule name
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether email notifications are supported for detections for this rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

// Static definitions of the ProactiveDetection configuration rule (same values for all components).
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions struct {
	// The rule description
	Description *string `pulumi:"description"`
	// The rule name as it is displayed in UI
	DisplayName *string `pulumi:"displayName"`
	// URL which displays additional info about the proactive detection rule
	HelpUrl *string `pulumi:"helpUrl"`
	// A flag indicating whether the rule is enabled by default
	IsEnabledByDefault *bool `pulumi:"isEnabledByDefault"`
	// A flag indicating whether the rule is hidden (from the UI)
	IsHidden *bool `pulumi:"isHidden"`
	// A flag indicating whether the rule is in preview
	IsInPreview *bool `pulumi:"isInPreview"`
	// The rule name
	Name *string `pulumi:"name"`
	// A flag indicating whether email notifications are supported for detections for this rule
	SupportsEmailNotifications *bool `pulumi:"supportsEmailNotifications"`
}

// ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsInput is an input type that accepts ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs and ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput values.
// You can construct a concrete instance of `ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsInput` via:
//
//	ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs{...}
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput
}

// Static definitions of the ProactiveDetection configuration rule (same values for all components).
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs struct {
	// The rule description
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The rule name as it is displayed in UI
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// URL which displays additional info about the proactive detection rule
	HelpUrl pulumi.StringPtrInput `pulumi:"helpUrl"`
	// A flag indicating whether the rule is enabled by default
	IsEnabledByDefault pulumi.BoolPtrInput `pulumi:"isEnabledByDefault"`
	// A flag indicating whether the rule is hidden (from the UI)
	IsHidden pulumi.BoolPtrInput `pulumi:"isHidden"`
	// A flag indicating whether the rule is in preview
	IsInPreview pulumi.BoolPtrInput `pulumi:"isInPreview"`
	// The rule name
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A flag indicating whether email notifications are supported for detections for this rule
	SupportsEmailNotifications pulumi.BoolPtrInput `pulumi:"supportsEmailNotifications"`
}

func (ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions)(nil)).Elem()
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput)
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput).ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(ctx)
}

// ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrInput is an input type that accepts ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs, ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtr and ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput values.
// You can construct a concrete instance of `ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrInput` via:
//
//	        ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs{...}
//
//	or:
//
//	        nil
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrInput interface {
	pulumi.Input

	ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput
	ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput
}

type applicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrType ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs

func ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtr(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsArgs) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrInput {
	return (*applicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrType)(v)
}

func (*applicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions)(nil)).Elem()
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return i.ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (i *applicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrType) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput)
}

// Static definitions of the ProactiveDetection configuration rule (same values for all components).
type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return o.ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(context.Background())
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions {
		return &v
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput)
}

// The rule description
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The rule name as it is displayed in UI
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// URL which displays additional info about the proactive detection rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether the rule is enabled by default
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is hidden (from the UI)
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is in preview
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

// The rule name
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether email notifications are supported for detections for this rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

type ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions)(nil)).Elem()
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) ToApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutputWithContext(ctx context.Context) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput {
	return o
}

func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) Elem() ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions {
		if v != nil {
			return *v
		}
		var ret ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions
		return ret
	}).(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput)
}

// The rule description
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The rule name as it is displayed in UI
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// URL which displays additional info about the proactive detection rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) HelpUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.HelpUrl
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether the rule is enabled by default
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) IsEnabledByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabledByDefault
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is hidden (from the UI)
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) IsHidden() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsHidden
	}).(pulumi.BoolPtrOutput)
}

// A flag indicating whether the rule is in preview
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) IsInPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.IsInPreview
	}).(pulumi.BoolPtrOutput)
}

// The rule name
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// A flag indicating whether email notifications are supported for detections for this rule
func (o ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput) SupportsEmailNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitions) *bool {
		if v == nil {
			return nil
		}
		return v.SupportsEmailNotifications
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesResponseRuleDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsOutput{})
	pulumi.RegisterOutputType(ApplicationInsightsComponentProactiveDetectionConfigurationPropertiesRuleDefinitionsPtrOutput{})
}
