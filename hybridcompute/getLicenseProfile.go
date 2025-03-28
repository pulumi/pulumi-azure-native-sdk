// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcompute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about the view of a license profile.
//
// Uses Azure REST API version 2023-06-20-preview.
//
// Other available API versions: 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-10, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13.
func LookupLicenseProfile(ctx *pulumi.Context, args *LookupLicenseProfileArgs, opts ...pulumi.InvokeOption) (*LookupLicenseProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupLicenseProfileResult
	err := ctx.Invoke("azure-native:hybridcompute:getLicenseProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLicenseProfileArgs struct {
	// The name of the license profile.
	LicenseProfileName string `pulumi:"licenseProfileName"`
	// The name of the hybrid machine.
	MachineName string `pulumi:"machineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes a license profile in a hybrid machine.
type LookupLicenseProfileResult struct {
	// The resource id of the license.
	AssignedLicense *string `pulumi:"assignedLicense"`
	// The guid id of the license.
	AssignedLicenseImmutableId string `pulumi:"assignedLicenseImmutableId"`
	// Indicates the eligibility state of Esu.
	EsuEligibility string `pulumi:"esuEligibility"`
	// Indicates whether there is an ESU Key currently active for the machine.
	EsuKeyState string `pulumi:"esuKeyState"`
	// The list of ESU keys.
	EsuKeys []EsuKeyResponse `pulumi:"esuKeys"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the Esu servers.
	ServerType string `pulumi:"serverType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupLicenseProfileOutput(ctx *pulumi.Context, args LookupLicenseProfileOutputArgs, opts ...pulumi.InvokeOption) LookupLicenseProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLicenseProfileResultOutput, error) {
			args := v.(LookupLicenseProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:hybridcompute:getLicenseProfile", args, LookupLicenseProfileResultOutput{}, options).(LookupLicenseProfileResultOutput), nil
		}).(LookupLicenseProfileResultOutput)
}

type LookupLicenseProfileOutputArgs struct {
	// The name of the license profile.
	LicenseProfileName pulumi.StringInput `pulumi:"licenseProfileName"`
	// The name of the hybrid machine.
	MachineName pulumi.StringInput `pulumi:"machineName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLicenseProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLicenseProfileArgs)(nil)).Elem()
}

// Describes a license profile in a hybrid machine.
type LookupLicenseProfileResultOutput struct{ *pulumi.OutputState }

func (LookupLicenseProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLicenseProfileResult)(nil)).Elem()
}

func (o LookupLicenseProfileResultOutput) ToLookupLicenseProfileResultOutput() LookupLicenseProfileResultOutput {
	return o
}

func (o LookupLicenseProfileResultOutput) ToLookupLicenseProfileResultOutputWithContext(ctx context.Context) LookupLicenseProfileResultOutput {
	return o
}

// The resource id of the license.
func (o LookupLicenseProfileResultOutput) AssignedLicense() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) *string { return v.AssignedLicense }).(pulumi.StringPtrOutput)
}

// The guid id of the license.
func (o LookupLicenseProfileResultOutput) AssignedLicenseImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.AssignedLicenseImmutableId }).(pulumi.StringOutput)
}

// Indicates the eligibility state of Esu.
func (o LookupLicenseProfileResultOutput) EsuEligibility() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.EsuEligibility }).(pulumi.StringOutput)
}

// Indicates whether there is an ESU Key currently active for the machine.
func (o LookupLicenseProfileResultOutput) EsuKeyState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.EsuKeyState }).(pulumi.StringOutput)
}

// The list of ESU keys.
func (o LookupLicenseProfileResultOutput) EsuKeys() EsuKeyResponseArrayOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) []EsuKeyResponse { return v.EsuKeys }).(EsuKeyResponseArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupLicenseProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupLicenseProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupLicenseProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state, which only appears in the response.
func (o LookupLicenseProfileResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the Esu servers.
func (o LookupLicenseProfileResultOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.ServerType }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupLicenseProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupLicenseProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupLicenseProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLicenseProfileResultOutput{})
}
