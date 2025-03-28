// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240808preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get an app attach package.
func LookupAppAttachPackage(ctx *pulumi.Context, args *LookupAppAttachPackageArgs, opts ...pulumi.InvokeOption) (*LookupAppAttachPackageResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAppAttachPackageResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20240808preview:getAppAttachPackage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppAttachPackageArgs struct {
	// The name of the App Attach package
	AppAttachPackageName string `pulumi:"appAttachPackageName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Schema for App Attach Package properties.
type LookupAppAttachPackageResult struct {
	// Field that can be populated with custom data and filtered on in list GET calls
	CustomData *string `pulumi:"customData"`
	// Parameter indicating how the health check should behave if this package fails staging
	FailHealthCheckOnStagingFailure *string `pulumi:"failHealthCheckOnStagingFailure"`
	// List of Hostpool resource Ids.
	HostPoolReferences []string `pulumi:"hostPoolReferences"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// Detailed properties for App Attach Package
	Image *AppAttachPackageInfoPropertiesResponse `pulumi:"image"`
	// URL path to certificate name located in keyVault
	KeyVaultURL *string `pulumi:"keyVaultURL"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Lookback url to third party control plane, is null for native app attach packages
	PackageLookbackUrl *string `pulumi:"packageLookbackUrl"`
	// Specific name of package owner, is "AppAttach" for native app attach packages
	PackageOwnerName *string `pulumi:"packageOwnerName"`
	// The provisioning state of the App Attach Package.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupAppAttachPackageOutput(ctx *pulumi.Context, args LookupAppAttachPackageOutputArgs, opts ...pulumi.InvokeOption) LookupAppAttachPackageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAppAttachPackageResultOutput, error) {
			args := v.(LookupAppAttachPackageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:desktopvirtualization/v20240808preview:getAppAttachPackage", args, LookupAppAttachPackageResultOutput{}, options).(LookupAppAttachPackageResultOutput), nil
		}).(LookupAppAttachPackageResultOutput)
}

type LookupAppAttachPackageOutputArgs struct {
	// The name of the App Attach package
	AppAttachPackageName pulumi.StringInput `pulumi:"appAttachPackageName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAppAttachPackageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppAttachPackageArgs)(nil)).Elem()
}

// Schema for App Attach Package properties.
type LookupAppAttachPackageResultOutput struct{ *pulumi.OutputState }

func (LookupAppAttachPackageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAppAttachPackageResult)(nil)).Elem()
}

func (o LookupAppAttachPackageResultOutput) ToLookupAppAttachPackageResultOutput() LookupAppAttachPackageResultOutput {
	return o
}

func (o LookupAppAttachPackageResultOutput) ToLookupAppAttachPackageResultOutputWithContext(ctx context.Context) LookupAppAttachPackageResultOutput {
	return o
}

// Field that can be populated with custom data and filtered on in list GET calls
func (o LookupAppAttachPackageResultOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

// Parameter indicating how the health check should behave if this package fails staging
func (o LookupAppAttachPackageResultOutput) FailHealthCheckOnStagingFailure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) *string { return v.FailHealthCheckOnStagingFailure }).(pulumi.StringPtrOutput)
}

// List of Hostpool resource Ids.
func (o LookupAppAttachPackageResultOutput) HostPoolReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) []string { return v.HostPoolReferences }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupAppAttachPackageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) string { return v.Id }).(pulumi.StringOutput)
}

// Detailed properties for App Attach Package
func (o LookupAppAttachPackageResultOutput) Image() AppAttachPackageInfoPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) *AppAttachPackageInfoPropertiesResponse { return v.Image }).(AppAttachPackageInfoPropertiesResponsePtrOutput)
}

// URL path to certificate name located in keyVault
func (o LookupAppAttachPackageResultOutput) KeyVaultURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) *string { return v.KeyVaultURL }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o LookupAppAttachPackageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupAppAttachPackageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Lookback url to third party control plane, is null for native app attach packages
func (o LookupAppAttachPackageResultOutput) PackageLookbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) *string { return v.PackageLookbackUrl }).(pulumi.StringPtrOutput)
}

// Specific name of package owner, is "AppAttach" for native app attach packages
func (o LookupAppAttachPackageResultOutput) PackageOwnerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) *string { return v.PackageOwnerName }).(pulumi.StringPtrOutput)
}

// The provisioning state of the App Attach Package.
func (o LookupAppAttachPackageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupAppAttachPackageResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAppAttachPackageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupAppAttachPackageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAppAttachPackageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAppAttachPackageResultOutput{})
}
