// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the Security ML Analytics Settings.
//
// Uses Azure REST API version 2024-09-01.
func LookupAnomalySecurityMLAnalyticsSettings(ctx *pulumi.Context, args *LookupAnomalySecurityMLAnalyticsSettingsArgs, opts ...pulumi.InvokeOption) (*LookupAnomalySecurityMLAnalyticsSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAnomalySecurityMLAnalyticsSettingsResult
	err := ctx.Invoke("azure-native:securityinsights:getAnomalySecurityMLAnalyticsSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAnomalySecurityMLAnalyticsSettingsArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Security ML Analytics Settings resource name
	SettingsResourceName string `pulumi:"settingsResourceName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Represents Anomaly Security ML Analytics Settings
type LookupAnomalySecurityMLAnalyticsSettingsResult struct {
	// The anomaly settings version of the Anomaly security ml analytics settings that dictates whether job version gets updated or not.
	AnomalySettingsVersion *int `pulumi:"anomalySettingsVersion"`
	// The anomaly version of the AnomalySecurityMLAnalyticsSettings.
	AnomalyVersion string `pulumi:"anomalyVersion"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The customizable observations of the AnomalySecurityMLAnalyticsSettings.
	CustomizableObservations interface{} `pulumi:"customizableObservations"`
	// The description of the SecurityMLAnalyticsSettings.
	Description *string `pulumi:"description"`
	// The display name for settings created by this SecurityMLAnalyticsSettings.
	DisplayName string `pulumi:"displayName"`
	// Determines whether this settings is enabled or disabled.
	Enabled bool `pulumi:"enabled"`
	// Etag of the azure resource
	Etag *string `pulumi:"etag"`
	// The frequency that this SecurityMLAnalyticsSettings will be run.
	Frequency string `pulumi:"frequency"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Determines whether this anomaly security ml analytics settings is a default settings
	IsDefaultSettings bool `pulumi:"isDefaultSettings"`
	// The kind of security ML analytics settings
	// Expected value is 'Anomaly'.
	Kind string `pulumi:"kind"`
	// The last time that this SecurityMLAnalyticsSettings has been modified.
	LastModifiedUtc string `pulumi:"lastModifiedUtc"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The required data sources for this SecurityMLAnalyticsSettings
	RequiredDataConnectors []SecurityMLAnalyticsSettingsDataSourceResponse `pulumi:"requiredDataConnectors"`
	// The anomaly settings definition Id
	SettingsDefinitionId *string `pulumi:"settingsDefinitionId"`
	// The anomaly SecurityMLAnalyticsSettings status
	SettingsStatus string `pulumi:"settingsStatus"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tactics of the SecurityMLAnalyticsSettings
	Tactics []string `pulumi:"tactics"`
	// The techniques of the SecurityMLAnalyticsSettings
	Techniques []string `pulumi:"techniques"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAnomalySecurityMLAnalyticsSettingsOutput(ctx *pulumi.Context, args LookupAnomalySecurityMLAnalyticsSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupAnomalySecurityMLAnalyticsSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAnomalySecurityMLAnalyticsSettingsResultOutput, error) {
			args := v.(LookupAnomalySecurityMLAnalyticsSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:securityinsights:getAnomalySecurityMLAnalyticsSettings", args, LookupAnomalySecurityMLAnalyticsSettingsResultOutput{}, options).(LookupAnomalySecurityMLAnalyticsSettingsResultOutput), nil
		}).(LookupAnomalySecurityMLAnalyticsSettingsResultOutput)
}

type LookupAnomalySecurityMLAnalyticsSettingsOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Security ML Analytics Settings resource name
	SettingsResourceName pulumi.StringInput `pulumi:"settingsResourceName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAnomalySecurityMLAnalyticsSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomalySecurityMLAnalyticsSettingsArgs)(nil)).Elem()
}

// Represents Anomaly Security ML Analytics Settings
type LookupAnomalySecurityMLAnalyticsSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupAnomalySecurityMLAnalyticsSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAnomalySecurityMLAnalyticsSettingsResult)(nil)).Elem()
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) ToLookupAnomalySecurityMLAnalyticsSettingsResultOutput() LookupAnomalySecurityMLAnalyticsSettingsResultOutput {
	return o
}

func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) ToLookupAnomalySecurityMLAnalyticsSettingsResultOutputWithContext(ctx context.Context) LookupAnomalySecurityMLAnalyticsSettingsResultOutput {
	return o
}

// The anomaly settings version of the Anomaly security ml analytics settings that dictates whether job version gets updated or not.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) AnomalySettingsVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *int { return v.AnomalySettingsVersion }).(pulumi.IntPtrOutput)
}

// The anomaly version of the AnomalySecurityMLAnalyticsSettings.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) AnomalyVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.AnomalyVersion }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The customizable observations of the AnomalySecurityMLAnalyticsSettings.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) CustomizableObservations() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) interface{} { return v.CustomizableObservations }).(pulumi.AnyOutput)
}

// The description of the SecurityMLAnalyticsSettings.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name for settings created by this SecurityMLAnalyticsSettings.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Determines whether this settings is enabled or disabled.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Etag of the azure resource
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The frequency that this SecurityMLAnalyticsSettings will be run.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Frequency }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Determines whether this anomaly security ml analytics settings is a default settings
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) IsDefaultSettings() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) bool { return v.IsDefaultSettings }).(pulumi.BoolOutput)
}

// The kind of security ML analytics settings
// Expected value is 'Anomaly'.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The last time that this SecurityMLAnalyticsSettings has been modified.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The required data sources for this SecurityMLAnalyticsSettings
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) RequiredDataConnectors() SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) []SecurityMLAnalyticsSettingsDataSourceResponse {
		return v.RequiredDataConnectors
	}).(SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput)
}

// The anomaly settings definition Id
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) SettingsDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) *string { return v.SettingsDefinitionId }).(pulumi.StringPtrOutput)
}

// The anomaly SecurityMLAnalyticsSettings status
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) SettingsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.SettingsStatus }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tactics of the SecurityMLAnalyticsSettings
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

// The techniques of the SecurityMLAnalyticsSettings
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAnomalySecurityMLAnalyticsSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAnomalySecurityMLAnalyticsSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAnomalySecurityMLAnalyticsSettingsResultOutput{})
}
