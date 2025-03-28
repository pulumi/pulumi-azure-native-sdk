// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets applicable inherited settings for this project.
func GetProjectInheritedSettings(ctx *pulumi.Context, args *GetProjectInheritedSettingsArgs, opts ...pulumi.InvokeOption) (*GetProjectInheritedSettingsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetProjectInheritedSettingsResult
	err := ctx.Invoke("azure-native:devcenter/v20250201:getProjectInheritedSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProjectInheritedSettingsArgs struct {
	// The name of the project.
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Applicable inherited settings for a project.
type GetProjectInheritedSettingsResult struct {
	// Network settings that will be enforced on this project.
	NetworkSettings ProjectNetworkSettingsResponse `pulumi:"networkSettings"`
	// Dev Center settings to be used when associating a project with a catalog.
	ProjectCatalogSettings DevCenterProjectCatalogSettingsResponse `pulumi:"projectCatalogSettings"`
}

func GetProjectInheritedSettingsOutput(ctx *pulumi.Context, args GetProjectInheritedSettingsOutputArgs, opts ...pulumi.InvokeOption) GetProjectInheritedSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectInheritedSettingsResultOutput, error) {
			args := v.(GetProjectInheritedSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devcenter/v20250201:getProjectInheritedSettings", args, GetProjectInheritedSettingsResultOutput{}, options).(GetProjectInheritedSettingsResultOutput), nil
		}).(GetProjectInheritedSettingsResultOutput)
}

type GetProjectInheritedSettingsOutputArgs struct {
	// The name of the project.
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetProjectInheritedSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectInheritedSettingsArgs)(nil)).Elem()
}

// Applicable inherited settings for a project.
type GetProjectInheritedSettingsResultOutput struct{ *pulumi.OutputState }

func (GetProjectInheritedSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectInheritedSettingsResult)(nil)).Elem()
}

func (o GetProjectInheritedSettingsResultOutput) ToGetProjectInheritedSettingsResultOutput() GetProjectInheritedSettingsResultOutput {
	return o
}

func (o GetProjectInheritedSettingsResultOutput) ToGetProjectInheritedSettingsResultOutputWithContext(ctx context.Context) GetProjectInheritedSettingsResultOutput {
	return o
}

// Network settings that will be enforced on this project.
func (o GetProjectInheritedSettingsResultOutput) NetworkSettings() ProjectNetworkSettingsResponseOutput {
	return o.ApplyT(func(v GetProjectInheritedSettingsResult) ProjectNetworkSettingsResponse { return v.NetworkSettings }).(ProjectNetworkSettingsResponseOutput)
}

// Dev Center settings to be used when associating a project with a catalog.
func (o GetProjectInheritedSettingsResultOutput) ProjectCatalogSettings() DevCenterProjectCatalogSettingsResponseOutput {
	return o.ApplyT(func(v GetProjectInheritedSettingsResult) DevCenterProjectCatalogSettingsResponse {
		return v.ProjectCatalogSettings
	}).(DevCenterProjectCatalogSettingsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectInheritedSettingsResultOutput{})
}
