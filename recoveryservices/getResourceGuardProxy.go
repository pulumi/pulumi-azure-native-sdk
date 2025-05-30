// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns ResourceGuardProxy under vault and with the name referenced in request
//
// Uses Azure REST API version 2024-10-01.
//
// Other available API versions: 2023-02-01, 2023-04-01, 2023-06-01, 2023-08-01, 2024-01-01, 2024-02-01, 2024-04-01, 2024-04-30-preview, 2024-07-30-preview, 2024-11-01-preview, 2025-01-01, 2025-02-01, 2025-02-28-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native recoveryservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupResourceGuardProxy(ctx *pulumi.Context, args *LookupResourceGuardProxyArgs, opts ...pulumi.InvokeOption) (*LookupResourceGuardProxyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceGuardProxyResult
	err := ctx.Invoke("azure-native:recoveryservices:getResourceGuardProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGuardProxyArgs struct {
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	ResourceGuardProxyName string `pulumi:"resourceGuardProxyName"`
	// The name of the recovery services vault.
	VaultName string `pulumi:"vaultName"`
}

type LookupResourceGuardProxyResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Optional ETag.
	ETag *string `pulumi:"eTag"`
	// Resource Id represents the complete path to the resource.
	Id string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name associated with the resource.
	Name string `pulumi:"name"`
	// ResourceGuardProxyBaseResource properties
	Properties ResourceGuardProxyBaseResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type string `pulumi:"type"`
}

func LookupResourceGuardProxyOutput(ctx *pulumi.Context, args LookupResourceGuardProxyOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGuardProxyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupResourceGuardProxyResultOutput, error) {
			args := v.(LookupResourceGuardProxyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:recoveryservices:getResourceGuardProxy", args, LookupResourceGuardProxyResultOutput{}, options).(LookupResourceGuardProxyResultOutput), nil
		}).(LookupResourceGuardProxyResultOutput)
}

type LookupResourceGuardProxyOutputArgs struct {
	// The name of the resource group where the recovery services vault is present.
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceGuardProxyName pulumi.StringInput `pulumi:"resourceGuardProxyName"`
	// The name of the recovery services vault.
	VaultName pulumi.StringInput `pulumi:"vaultName"`
}

func (LookupResourceGuardProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardProxyArgs)(nil)).Elem()
}

type LookupResourceGuardProxyResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGuardProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGuardProxyResult)(nil)).Elem()
}

func (o LookupResourceGuardProxyResultOutput) ToLookupResourceGuardProxyResultOutput() LookupResourceGuardProxyResultOutput {
	return o
}

func (o LookupResourceGuardProxyResultOutput) ToLookupResourceGuardProxyResultOutputWithContext(ctx context.Context) LookupResourceGuardProxyResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupResourceGuardProxyResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Optional ETag.
func (o LookupResourceGuardProxyResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Resource Id represents the complete path to the resource.
func (o LookupResourceGuardProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location.
func (o LookupResourceGuardProxyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name associated with the resource.
func (o LookupResourceGuardProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// ResourceGuardProxyBaseResource properties
func (o LookupResourceGuardProxyResultOutput) Properties() ResourceGuardProxyBaseResponseOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) ResourceGuardProxyBaseResponse { return v.Properties }).(ResourceGuardProxyBaseResponseOutput)
}

// Resource tags.
func (o LookupResourceGuardProxyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
func (o LookupResourceGuardProxyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceGuardProxyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGuardProxyResultOutput{})
}
