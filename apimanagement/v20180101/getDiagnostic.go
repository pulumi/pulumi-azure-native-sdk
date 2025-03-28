// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of the Diagnostic specified by its identifier.
func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Diagnostic details.
type LookupDiagnosticResult struct {
	// Indicates whether a diagnostic should receive data or not.
	Enabled bool `pulumi:"enabled"`
	// Resource ID.
	Id string `pulumi:"id"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}

func LookupDiagnosticOutput(ctx *pulumi.Context, args LookupDiagnosticOutputArgs, opts ...pulumi.InvokeOption) LookupDiagnosticResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDiagnosticResultOutput, error) {
			args := v.(LookupDiagnosticArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:apimanagement/v20180101:getDiagnostic", args, LookupDiagnosticResultOutput{}, options).(LookupDiagnosticResultOutput), nil
		}).(LookupDiagnosticResultOutput)
}

type LookupDiagnosticOutputArgs struct {
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput `pulumi:"diagnosticId"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDiagnosticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticArgs)(nil)).Elem()
}

// Diagnostic details.
type LookupDiagnosticResultOutput struct{ *pulumi.OutputState }

func (LookupDiagnosticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticResult)(nil)).Elem()
}

func (o LookupDiagnosticResultOutput) ToLookupDiagnosticResultOutput() LookupDiagnosticResultOutput {
	return o
}

func (o LookupDiagnosticResultOutput) ToLookupDiagnosticResultOutputWithContext(ctx context.Context) LookupDiagnosticResultOutput {
	return o
}

// Indicates whether a diagnostic should receive data or not.
func (o LookupDiagnosticResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Resource ID.
func (o LookupDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type for API Management resource.
func (o LookupDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticResultOutput{})
}
