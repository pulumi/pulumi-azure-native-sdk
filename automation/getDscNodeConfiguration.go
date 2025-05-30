// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve the Dsc node configurations by node configuration.
//
// Uses Azure REST API version 2023-11-01.
//
// Other available API versions: 2015-10-31, 2018-01-15, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupDscNodeConfiguration(ctx *pulumi.Context, args *LookupDscNodeConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDscNodeConfigurationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDscNodeConfigurationResult
	err := ctx.Invoke("azure-native:automation:getDscNodeConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDscNodeConfigurationArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The Dsc node configuration name.
	NodeConfigurationName string `pulumi:"nodeConfigurationName"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Definition of the dsc node configuration.
type LookupDscNodeConfigurationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Gets or sets the configuration of the node.
	Configuration *DscConfigurationAssociationPropertyResponse `pulumi:"configuration"`
	// Gets or sets creation time.
	CreationTime *string `pulumi:"creationTime"`
	// Fully qualified resource Id for the resource
	Id string `pulumi:"id"`
	// If a new build version of NodeConfiguration is required.
	IncrementNodeConfigurationBuild *bool `pulumi:"incrementNodeConfigurationBuild"`
	// Gets or sets the last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Number of nodes with this node configuration assigned
	NodeCount *float64 `pulumi:"nodeCount"`
	// Source of node configuration.
	Source *string `pulumi:"source"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupDscNodeConfigurationOutput(ctx *pulumi.Context, args LookupDscNodeConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDscNodeConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDscNodeConfigurationResultOutput, error) {
			args := v.(LookupDscNodeConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:automation:getDscNodeConfiguration", args, LookupDscNodeConfigurationResultOutput{}, options).(LookupDscNodeConfigurationResultOutput), nil
		}).(LookupDscNodeConfigurationResultOutput)
}

type LookupDscNodeConfigurationOutputArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput `pulumi:"automationAccountName"`
	// The Dsc node configuration name.
	NodeConfigurationName pulumi.StringInput `pulumi:"nodeConfigurationName"`
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDscNodeConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscNodeConfigurationArgs)(nil)).Elem()
}

// Definition of the dsc node configuration.
type LookupDscNodeConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDscNodeConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDscNodeConfigurationResult)(nil)).Elem()
}

func (o LookupDscNodeConfigurationResultOutput) ToLookupDscNodeConfigurationResultOutput() LookupDscNodeConfigurationResultOutput {
	return o
}

func (o LookupDscNodeConfigurationResultOutput) ToLookupDscNodeConfigurationResultOutputWithContext(ctx context.Context) LookupDscNodeConfigurationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDscNodeConfigurationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the configuration of the node.
func (o LookupDscNodeConfigurationResultOutput) Configuration() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *DscConfigurationAssociationPropertyResponse {
		return v.Configuration
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

// Gets or sets creation time.
func (o LookupDscNodeConfigurationResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Fully qualified resource Id for the resource
func (o LookupDscNodeConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// If a new build version of NodeConfiguration is required.
func (o LookupDscNodeConfigurationResultOutput) IncrementNodeConfigurationBuild() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *bool { return v.IncrementNodeConfigurationBuild }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o LookupDscNodeConfigurationResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o LookupDscNodeConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of nodes with this node configuration assigned
func (o LookupDscNodeConfigurationResultOutput) NodeCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *float64 { return v.NodeCount }).(pulumi.Float64PtrOutput)
}

// Source of node configuration.
func (o LookupDscNodeConfigurationResultOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) *string { return v.Source }).(pulumi.StringPtrOutput)
}

// The type of the resource.
func (o LookupDscNodeConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDscNodeConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDscNodeConfigurationResultOutput{})
}
