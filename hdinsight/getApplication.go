// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets properties of the specified application.
//
// Uses Azure REST API version 2024-08-01-preview.
//
// Other available API versions: 2021-06-01, 2023-04-15-preview, 2023-08-15-preview, 2025-01-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationResult
	err := ctx.Invoke("azure-native:hdinsight:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationArgs struct {
	// The constant value for the application name.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The HDInsight cluster application
type LookupApplicationResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The ETag for the application
	Etag *string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The properties of the application.
	Properties ApplicationPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags for the application.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupApplicationOutput(ctx *pulumi.Context, args LookupApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupApplicationResultOutput, error) {
			args := v.(LookupApplicationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hdinsight:getApplication", args, LookupApplicationResultOutput{}, options).(LookupApplicationResultOutput), nil
		}).(LookupApplicationResultOutput)
}

type LookupApplicationOutputArgs struct {
	// The constant value for the application name.
	ApplicationName pulumi.StringInput `pulumi:"applicationName"`
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

// The HDInsight cluster application
type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupApplicationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The ETag for the application
func (o LookupApplicationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The properties of the application.
func (o LookupApplicationResultOutput) Properties() ApplicationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) ApplicationPropertiesResponse { return v.Properties }).(ApplicationPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApplicationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags for the application.
func (o LookupApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
