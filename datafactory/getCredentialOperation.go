// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a credential.
//
// Uses Azure REST API version 2018-06-01.
func LookupCredentialOperation(ctx *pulumi.Context, args *LookupCredentialOperationArgs, opts ...pulumi.InvokeOption) (*LookupCredentialOperationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCredentialOperationResult
	err := ctx.Invoke("azure-native:datafactory:getCredentialOperation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCredentialOperationArgs struct {
	// Credential name
	CredentialName string `pulumi:"credentialName"`
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Credential resource type.
type LookupCredentialOperationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// Properties of credentials.
	Properties interface{} `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupCredentialOperationOutput(ctx *pulumi.Context, args LookupCredentialOperationOutputArgs, opts ...pulumi.InvokeOption) LookupCredentialOperationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCredentialOperationResultOutput, error) {
			args := v.(LookupCredentialOperationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datafactory:getCredentialOperation", args, LookupCredentialOperationResultOutput{}, options).(LookupCredentialOperationResultOutput), nil
		}).(LookupCredentialOperationResultOutput)
}

type LookupCredentialOperationOutputArgs struct {
	// Credential name
	CredentialName pulumi.StringInput `pulumi:"credentialName"`
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCredentialOperationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialOperationArgs)(nil)).Elem()
}

// Credential resource type.
type LookupCredentialOperationResultOutput struct{ *pulumi.OutputState }

func (LookupCredentialOperationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCredentialOperationResult)(nil)).Elem()
}

func (o LookupCredentialOperationResultOutput) ToLookupCredentialOperationResultOutput() LookupCredentialOperationResultOutput {
	return o
}

func (o LookupCredentialOperationResultOutput) ToLookupCredentialOperationResultOutputWithContext(ctx context.Context) LookupCredentialOperationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupCredentialOperationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Etag identifies change in the resource.
func (o LookupCredentialOperationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupCredentialOperationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupCredentialOperationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of credentials.
func (o LookupCredentialOperationResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// The resource type.
func (o LookupCredentialOperationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCredentialOperationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCredentialOperationResultOutput{})
}
