// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Returns whether FTP is allowed on the site or not.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupWebAppFtpAllowedSlot(ctx *pulumi.Context, args *LookupWebAppFtpAllowedSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppFtpAllowedSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppFtpAllowedSlotResult
	err := ctx.Invoke("azure-native:web:getWebAppFtpAllowedSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppFtpAllowedSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}

// Publishing Credentials Policies parameters.
type LookupWebAppFtpAllowedSlotResult struct {
	// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
	Allow bool `pulumi:"allow"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupWebAppFtpAllowedSlotOutput(ctx *pulumi.Context, args LookupWebAppFtpAllowedSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppFtpAllowedSlotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWebAppFtpAllowedSlotResultOutput, error) {
			args := v.(LookupWebAppFtpAllowedSlotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web:getWebAppFtpAllowedSlot", args, LookupWebAppFtpAllowedSlotResultOutput{}, options).(LookupWebAppFtpAllowedSlotResultOutput), nil
		}).(LookupWebAppFtpAllowedSlotResultOutput)
}

type LookupWebAppFtpAllowedSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppFtpAllowedSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFtpAllowedSlotArgs)(nil)).Elem()
}

// Publishing Credentials Policies parameters.
type LookupWebAppFtpAllowedSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppFtpAllowedSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFtpAllowedSlotResult)(nil)).Elem()
}

func (o LookupWebAppFtpAllowedSlotResultOutput) ToLookupWebAppFtpAllowedSlotResultOutput() LookupWebAppFtpAllowedSlotResultOutput {
	return o
}

func (o LookupWebAppFtpAllowedSlotResultOutput) ToLookupWebAppFtpAllowedSlotResultOutputWithContext(ctx context.Context) LookupWebAppFtpAllowedSlotResultOutput {
	return o
}

// <code>true</code> to allow access to a publishing method; otherwise, <code>false</code>.
func (o LookupWebAppFtpAllowedSlotResultOutput) Allow() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedSlotResult) bool { return v.Allow }).(pulumi.BoolOutput)
}

// The Azure API version of the resource.
func (o LookupWebAppFtpAllowedSlotResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedSlotResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource Id.
func (o LookupWebAppFtpAllowedSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o LookupWebAppFtpAllowedSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppFtpAllowedSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o LookupWebAppFtpAllowedSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFtpAllowedSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppFtpAllowedSlotResultOutput{})
}
