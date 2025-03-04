// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure REST API version: 2022-01-01-preview.
func LookupSubAccount(ctx *pulumi.Context, args *LookupSubAccountArgs, opts ...pulumi.InvokeOption) (*LookupSubAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSubAccountResult
	err := ctx.Invoke("azure-native:logz:getSubAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubAccountArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName string `pulumi:"subAccountName"`
}

type LookupSubAccountResult struct {
	// ARM id of the monitor resource.
	Id       string                      `pulumi:"id"`
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	Location string                      `pulumi:"location"`
	// Name of the monitor resource.
	Name string `pulumi:"name"`
	// Properties specific to the monitor resource.
	Properties MonitorPropertiesResponse `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	// The type of the monitor resource.
	Type string `pulumi:"type"`
}

func LookupSubAccountOutput(ctx *pulumi.Context, args LookupSubAccountOutputArgs, opts ...pulumi.InvokeOption) LookupSubAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSubAccountResultOutput, error) {
			args := v.(LookupSubAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:logz:getSubAccount", args, LookupSubAccountResultOutput{}, options).(LookupSubAccountResultOutput), nil
		}).(LookupSubAccountResultOutput)
}

type LookupSubAccountOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName pulumi.StringInput `pulumi:"subAccountName"`
}

func (LookupSubAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubAccountArgs)(nil)).Elem()
}

type LookupSubAccountResultOutput struct{ *pulumi.OutputState }

func (LookupSubAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubAccountResult)(nil)).Elem()
}

func (o LookupSubAccountResultOutput) ToLookupSubAccountResultOutput() LookupSubAccountResultOutput {
	return o
}

func (o LookupSubAccountResultOutput) ToLookupSubAccountResultOutputWithContext(ctx context.Context) LookupSubAccountResultOutput {
	return o
}

// ARM id of the monitor resource.
func (o LookupSubAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubAccountResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSubAccountResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupSubAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o LookupSubAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the monitor resource.
func (o LookupSubAccountResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupSubAccountResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupSubAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSubAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSubAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the monitor resource.
func (o LookupSubAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubAccountResultOutput{})
}
