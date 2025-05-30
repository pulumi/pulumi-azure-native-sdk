// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get bare metal machine key set of the provided cluster.
//
// Uses Azure REST API version 2025-02-01.
//
// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupBareMetalMachineKeySet(ctx *pulumi.Context, args *LookupBareMetalMachineKeySetArgs, opts ...pulumi.InvokeOption) (*LookupBareMetalMachineKeySetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBareMetalMachineKeySetResult
	err := ctx.Invoke("azure-native:networkcloud:getBareMetalMachineKeySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBareMetalMachineKeySetArgs struct {
	// The name of the bare metal machine key set.
	BareMetalMachineKeySetName string `pulumi:"bareMetalMachineKeySetName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupBareMetalMachineKeySetResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
	AzureGroupId string `pulumi:"azureGroupId"`
	// The more detailed status of the key set.
	DetailedStatus string `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage string `pulumi:"detailedStatusMessage"`
	// Resource ETag.
	Etag string `pulumi:"etag"`
	// The date and time after which the users in this key set will be removed from the bare metal machines.
	Expiration string `pulumi:"expiration"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The list of IP addresses of jump hosts with management network access from which a login will be allowed for the users.
	JumpHostsAllowed []string `pulumi:"jumpHostsAllowed"`
	// The last time this key set was validated.
	LastValidation string `pulumi:"lastValidation"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The name of the group that users will be assigned to on the operating system of the machines.
	OsGroupName *string `pulumi:"osGroupName"`
	// The access level allowed for the users in this key set.
	PrivilegeLevel string `pulumi:"privilegeLevel"`
	// The provisioning state of the bare metal machine key set.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// The unique list of permitted users.
	UserList []KeySetUserResponse `pulumi:"userList"`
	// The status evaluation of each user.
	UserListStatus []KeySetUserStatusResponse `pulumi:"userListStatus"`
}

func LookupBareMetalMachineKeySetOutput(ctx *pulumi.Context, args LookupBareMetalMachineKeySetOutputArgs, opts ...pulumi.InvokeOption) LookupBareMetalMachineKeySetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBareMetalMachineKeySetResultOutput, error) {
			args := v.(LookupBareMetalMachineKeySetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:networkcloud:getBareMetalMachineKeySet", args, LookupBareMetalMachineKeySetResultOutput{}, options).(LookupBareMetalMachineKeySetResultOutput), nil
		}).(LookupBareMetalMachineKeySetResultOutput)
}

type LookupBareMetalMachineKeySetOutputArgs struct {
	// The name of the bare metal machine key set.
	BareMetalMachineKeySetName pulumi.StringInput `pulumi:"bareMetalMachineKeySetName"`
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBareMetalMachineKeySetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBareMetalMachineKeySetArgs)(nil)).Elem()
}

type LookupBareMetalMachineKeySetResultOutput struct{ *pulumi.OutputState }

func (LookupBareMetalMachineKeySetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBareMetalMachineKeySetResult)(nil)).Elem()
}

func (o LookupBareMetalMachineKeySetResultOutput) ToLookupBareMetalMachineKeySetResultOutput() LookupBareMetalMachineKeySetResultOutput {
	return o
}

func (o LookupBareMetalMachineKeySetResultOutput) ToLookupBareMetalMachineKeySetResultOutputWithContext(ctx context.Context) LookupBareMetalMachineKeySetResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupBareMetalMachineKeySetResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The object ID of Azure Active Directory group that all users in the list must be in for access to be granted. Users that are not in the group will not have access.
func (o LookupBareMetalMachineKeySetResultOutput) AzureGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.AzureGroupId }).(pulumi.StringOutput)
}

// The more detailed status of the key set.
func (o LookupBareMetalMachineKeySetResultOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o LookupBareMetalMachineKeySetResultOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// Resource ETag.
func (o LookupBareMetalMachineKeySetResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The date and time after which the users in this key set will be removed from the bare metal machines.
func (o LookupBareMetalMachineKeySetResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o LookupBareMetalMachineKeySetResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupBareMetalMachineKeySetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of IP addresses of jump hosts with management network access from which a login will be allowed for the users.
func (o LookupBareMetalMachineKeySetResultOutput) JumpHostsAllowed() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) []string { return v.JumpHostsAllowed }).(pulumi.StringArrayOutput)
}

// The last time this key set was validated.
func (o LookupBareMetalMachineKeySetResultOutput) LastValidation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.LastValidation }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupBareMetalMachineKeySetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupBareMetalMachineKeySetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the group that users will be assigned to on the operating system of the machines.
func (o LookupBareMetalMachineKeySetResultOutput) OsGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) *string { return v.OsGroupName }).(pulumi.StringPtrOutput)
}

// The access level allowed for the users in this key set.
func (o LookupBareMetalMachineKeySetResultOutput) PrivilegeLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.PrivilegeLevel }).(pulumi.StringOutput)
}

// The provisioning state of the bare metal machine key set.
func (o LookupBareMetalMachineKeySetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupBareMetalMachineKeySetResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupBareMetalMachineKeySetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupBareMetalMachineKeySetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique list of permitted users.
func (o LookupBareMetalMachineKeySetResultOutput) UserList() KeySetUserResponseArrayOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) []KeySetUserResponse { return v.UserList }).(KeySetUserResponseArrayOutput)
}

// The status evaluation of each user.
func (o LookupBareMetalMachineKeySetResultOutput) UserListStatus() KeySetUserStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupBareMetalMachineKeySetResult) []KeySetUserStatusResponse { return v.UserListStatus }).(KeySetUserStatusResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBareMetalMachineKeySetResultOutput{})
}
