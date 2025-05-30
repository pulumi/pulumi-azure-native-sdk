// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongocluster

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a mongo cluster.
//
// Uses Azure REST API version 2024-07-01.
//
// Other available API versions: 2024-03-01-preview, 2024-06-01-preview, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mongocluster [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupMongoCluster(ctx *pulumi.Context, args *LookupMongoClusterArgs, opts ...pulumi.InvokeOption) (*LookupMongoClusterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoClusterResult
	err := ctx.Invoke("azure-native:mongocluster:getMongoCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMongoClusterArgs struct {
	// The name of the mongo cluster.
	MongoClusterName string `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a mongo cluster resource.
type LookupMongoClusterResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties MongoClusterPropertiesResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupMongoClusterOutput(ctx *pulumi.Context, args LookupMongoClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMongoClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMongoClusterResultOutput, error) {
			args := v.(LookupMongoClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:mongocluster:getMongoCluster", args, LookupMongoClusterResultOutput{}, options).(LookupMongoClusterResultOutput), nil
		}).(LookupMongoClusterResultOutput)
}

type LookupMongoClusterOutputArgs struct {
	// The name of the mongo cluster.
	MongoClusterName pulumi.StringInput `pulumi:"mongoClusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMongoClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoClusterArgs)(nil)).Elem()
}

// Represents a mongo cluster resource.
type LookupMongoClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMongoClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoClusterResult)(nil)).Elem()
}

func (o LookupMongoClusterResultOutput) ToLookupMongoClusterResultOutput() LookupMongoClusterResultOutput {
	return o
}

func (o LookupMongoClusterResultOutput) ToLookupMongoClusterResultOutputWithContext(ctx context.Context) LookupMongoClusterResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMongoClusterResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupMongoClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupMongoClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupMongoClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o LookupMongoClusterResultOutput) Properties() MongoClusterPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) MongoClusterPropertiesResponse { return v.Properties }).(MongoClusterPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupMongoClusterResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupMongoClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupMongoClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoClusterResultOutput{})
}
