// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the hybrid identity metadata proxy resource.
func LookupHybridIdentityMetadatum(ctx *pulumi.Context, args *LookupHybridIdentityMetadatumArgs, opts ...pulumi.InvokeOption) (*LookupHybridIdentityMetadatumResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupHybridIdentityMetadatumResult
	err := ctx.Invoke("azure-native:hybridcontainerservice/v20231115preview:getHybridIdentityMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridIdentityMetadatumArgs struct {
	// The fully qualified Azure Resource manager identifier of the connected cluster resource.
	ConnectedClusterResourceUri string `pulumi:"connectedClusterResourceUri"`
}

// Defines the hybridIdentityMetadata.
type LookupHybridIdentityMetadatumResult struct {
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource
	ProvisioningState string `pulumi:"provisioningState"`
	// Onboarding public key for provisioning the Managed identity for the HybridAKS cluster.
	PublicKey *string `pulumi:"publicKey"`
	// Unique id of the parent provisioned cluster resource.
	ResourceUid *string `pulumi:"resourceUid"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupHybridIdentityMetadatumOutput(ctx *pulumi.Context, args LookupHybridIdentityMetadatumOutputArgs, opts ...pulumi.InvokeOption) LookupHybridIdentityMetadatumResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupHybridIdentityMetadatumResultOutput, error) {
			args := v.(LookupHybridIdentityMetadatumArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridcontainerservice/v20231115preview:getHybridIdentityMetadatum", args, LookupHybridIdentityMetadatumResultOutput{}, options).(LookupHybridIdentityMetadatumResultOutput), nil
		}).(LookupHybridIdentityMetadatumResultOutput)
}

type LookupHybridIdentityMetadatumOutputArgs struct {
	// The fully qualified Azure Resource manager identifier of the connected cluster resource.
	ConnectedClusterResourceUri pulumi.StringInput `pulumi:"connectedClusterResourceUri"`
}

func (LookupHybridIdentityMetadatumOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridIdentityMetadatumArgs)(nil)).Elem()
}

// Defines the hybridIdentityMetadata.
type LookupHybridIdentityMetadatumResultOutput struct{ *pulumi.OutputState }

func (LookupHybridIdentityMetadatumResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHybridIdentityMetadatumResult)(nil)).Elem()
}

func (o LookupHybridIdentityMetadatumResultOutput) ToLookupHybridIdentityMetadatumResultOutput() LookupHybridIdentityMetadatumResultOutput {
	return o
}

func (o LookupHybridIdentityMetadatumResultOutput) ToLookupHybridIdentityMetadatumResultOutputWithContext(ctx context.Context) LookupHybridIdentityMetadatumResultOutput {
	return o
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupHybridIdentityMetadatumResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupHybridIdentityMetadatumResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource
func (o LookupHybridIdentityMetadatumResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Onboarding public key for provisioning the Managed identity for the HybridAKS cluster.
func (o LookupHybridIdentityMetadatumResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Unique id of the parent provisioned cluster resource.
func (o LookupHybridIdentityMetadatumResultOutput) ResourceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) *string { return v.ResourceUid }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupHybridIdentityMetadatumResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupHybridIdentityMetadatumResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHybridIdentityMetadatumResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHybridIdentityMetadatumResultOutput{})
}
