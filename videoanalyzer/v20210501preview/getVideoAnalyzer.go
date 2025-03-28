// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the details of the specified Video Analyzer account
func LookupVideoAnalyzer(ctx *pulumi.Context, args *LookupVideoAnalyzerArgs, opts ...pulumi.InvokeOption) (*LookupVideoAnalyzerResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupVideoAnalyzerResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getVideoAnalyzer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVideoAnalyzerArgs struct {
	// The Video Analyzer account name.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A Video Analyzer account.
type LookupVideoAnalyzerResult struct {
	// The account encryption properties.
	Encryption AccountEncryptionResponse `pulumi:"encryption"`
	// The list of endpoints associated with this resource.
	Endpoints []EndpointResponse `pulumi:"endpoints"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The set of managed identities associated with the Video Analyzer resource.
	Identity *VideoAnalyzerIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The storage accounts for this resource.
	StorageAccounts []StorageAccountResponse `pulumi:"storageAccounts"`
	// The system data of the Video Analyzer account.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupVideoAnalyzerOutput(ctx *pulumi.Context, args LookupVideoAnalyzerOutputArgs, opts ...pulumi.InvokeOption) LookupVideoAnalyzerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVideoAnalyzerResultOutput, error) {
			args := v.(LookupVideoAnalyzerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:videoanalyzer/v20210501preview:getVideoAnalyzer", args, LookupVideoAnalyzerResultOutput{}, options).(LookupVideoAnalyzerResultOutput), nil
		}).(LookupVideoAnalyzerResultOutput)
}

type LookupVideoAnalyzerOutputArgs struct {
	// The Video Analyzer account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVideoAnalyzerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoAnalyzerArgs)(nil)).Elem()
}

// A Video Analyzer account.
type LookupVideoAnalyzerResultOutput struct{ *pulumi.OutputState }

func (LookupVideoAnalyzerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoAnalyzerResult)(nil)).Elem()
}

func (o LookupVideoAnalyzerResultOutput) ToLookupVideoAnalyzerResultOutput() LookupVideoAnalyzerResultOutput {
	return o
}

func (o LookupVideoAnalyzerResultOutput) ToLookupVideoAnalyzerResultOutputWithContext(ctx context.Context) LookupVideoAnalyzerResultOutput {
	return o
}

// The account encryption properties.
func (o LookupVideoAnalyzerResultOutput) Encryption() AccountEncryptionResponseOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) AccountEncryptionResponse { return v.Encryption }).(AccountEncryptionResponseOutput)
}

// The list of endpoints associated with this resource.
func (o LookupVideoAnalyzerResultOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) []EndpointResponse { return v.Endpoints }).(EndpointResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupVideoAnalyzerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Id }).(pulumi.StringOutput)
}

// The set of managed identities associated with the Video Analyzer resource.
func (o LookupVideoAnalyzerResultOutput) Identity() VideoAnalyzerIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) *VideoAnalyzerIdentityResponse { return v.Identity }).(VideoAnalyzerIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupVideoAnalyzerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupVideoAnalyzerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Name }).(pulumi.StringOutput)
}

// The storage accounts for this resource.
func (o LookupVideoAnalyzerResultOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) []StorageAccountResponse { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

// The system data of the Video Analyzer account.
func (o LookupVideoAnalyzerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupVideoAnalyzerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupVideoAnalyzerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVideoAnalyzerResultOutput{})
}
