// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the specified artifact store.
//
// Uses Azure REST API version 2024-04-15.
//
// Other available API versions: 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridnetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupArtifactStore(ctx *pulumi.Context, args *LookupArtifactStoreArgs, opts ...pulumi.InvokeOption) (*LookupArtifactStoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupArtifactStoreResult
	err := ctx.Invoke("azure-native:hybridnetwork:getArtifactStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupArtifactStoreArgs struct {
	// The name of the artifact store.
	ArtifactStoreName string `pulumi:"artifactStoreName"`
	// The name of the publisher.
	PublisherName string `pulumi:"publisherName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Artifact store properties.
type LookupArtifactStoreResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// ArtifactStores properties.
	Properties ArtifactStorePropertiesFormatResponse `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupArtifactStoreOutput(ctx *pulumi.Context, args LookupArtifactStoreOutputArgs, opts ...pulumi.InvokeOption) LookupArtifactStoreResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupArtifactStoreResultOutput, error) {
			args := v.(LookupArtifactStoreArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridnetwork:getArtifactStore", args, LookupArtifactStoreResultOutput{}, options).(LookupArtifactStoreResultOutput), nil
		}).(LookupArtifactStoreResultOutput)
}

type LookupArtifactStoreOutputArgs struct {
	// The name of the artifact store.
	ArtifactStoreName pulumi.StringInput `pulumi:"artifactStoreName"`
	// The name of the publisher.
	PublisherName pulumi.StringInput `pulumi:"publisherName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupArtifactStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactStoreArgs)(nil)).Elem()
}

// Artifact store properties.
type LookupArtifactStoreResultOutput struct{ *pulumi.OutputState }

func (LookupArtifactStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupArtifactStoreResult)(nil)).Elem()
}

func (o LookupArtifactStoreResultOutput) ToLookupArtifactStoreResultOutput() LookupArtifactStoreResultOutput {
	return o
}

func (o LookupArtifactStoreResultOutput) ToLookupArtifactStoreResultOutputWithContext(ctx context.Context) LookupArtifactStoreResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupArtifactStoreResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupArtifactStoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupArtifactStoreResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupArtifactStoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// ArtifactStores properties.
func (o LookupArtifactStoreResultOutput) Properties() ArtifactStorePropertiesFormatResponseOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) ArtifactStorePropertiesFormatResponse { return v.Properties }).(ArtifactStorePropertiesFormatResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupArtifactStoreResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupArtifactStoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupArtifactStoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupArtifactStoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupArtifactStoreResultOutput{})
}
