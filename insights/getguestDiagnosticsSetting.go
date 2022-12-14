// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Virtual machine guest diagnostics settings resource.
// API Version: 2018-06-01-preview.
func GetguestDiagnosticsSetting(ctx *pulumi.Context, args *GetguestDiagnosticsSettingArgs, opts ...pulumi.InvokeOption) (*GetguestDiagnosticsSettingResult, error) {
	var rv GetguestDiagnosticsSettingResult
	err := ctx.Invoke("azure-native:insights:getguestDiagnosticsSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetguestDiagnosticsSettingArgs struct {
	// The name of the diagnostic setting.
	DiagnosticSettingsName string `pulumi:"diagnosticSettingsName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Virtual machine guest diagnostics settings resource.
type GetguestDiagnosticsSettingResult struct {
	// the array of data source object which are configured to collect and send data
	DataSources []DataSourceResponse `pulumi:"dataSources"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// Operating system type for the configuration
	OsType       *string `pulumi:"osType"`
	ProxySetting *string `pulumi:"proxySetting"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure resource type
	Type string `pulumi:"type"`
}

func GetguestDiagnosticsSettingOutput(ctx *pulumi.Context, args GetguestDiagnosticsSettingOutputArgs, opts ...pulumi.InvokeOption) GetguestDiagnosticsSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetguestDiagnosticsSettingResult, error) {
			args := v.(GetguestDiagnosticsSettingArgs)
			r, err := GetguestDiagnosticsSetting(ctx, &args, opts...)
			var s GetguestDiagnosticsSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetguestDiagnosticsSettingResultOutput)
}

type GetguestDiagnosticsSettingOutputArgs struct {
	// The name of the diagnostic setting.
	DiagnosticSettingsName pulumi.StringInput `pulumi:"diagnosticSettingsName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetguestDiagnosticsSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetguestDiagnosticsSettingArgs)(nil)).Elem()
}

// Virtual machine guest diagnostics settings resource.
type GetguestDiagnosticsSettingResultOutput struct{ *pulumi.OutputState }

func (GetguestDiagnosticsSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetguestDiagnosticsSettingResult)(nil)).Elem()
}

func (o GetguestDiagnosticsSettingResultOutput) ToGetguestDiagnosticsSettingResultOutput() GetguestDiagnosticsSettingResultOutput {
	return o
}

func (o GetguestDiagnosticsSettingResultOutput) ToGetguestDiagnosticsSettingResultOutputWithContext(ctx context.Context) GetguestDiagnosticsSettingResultOutput {
	return o
}

// the array of data source object which are configured to collect and send data
func (o GetguestDiagnosticsSettingResultOutput) DataSources() DataSourceResponseArrayOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) []DataSourceResponse { return v.DataSources }).(DataSourceResponseArrayOutput)
}

// Azure resource Id
func (o GetguestDiagnosticsSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o GetguestDiagnosticsSettingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o GetguestDiagnosticsSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operating system type for the configuration
func (o GetguestDiagnosticsSettingResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GetguestDiagnosticsSettingResultOutput) ProxySetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) *string { return v.ProxySetting }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o GetguestDiagnosticsSettingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type
func (o GetguestDiagnosticsSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetguestDiagnosticsSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetguestDiagnosticsSettingResultOutput{})
}
