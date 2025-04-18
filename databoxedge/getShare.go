// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a share on the  Data Box Edge/Gateway device.
//
// Uses Azure REST API version 2023-07-01.
//
// Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupShareResult
	err := ctx.Invoke("azure-native:databoxedge:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The share name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Represents a share on the  Data Box Edge/Gateway device.
type LookupShareResult struct {
	// Access protocol to be used by the share.
	AccessProtocol string `pulumi:"accessProtocol"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Azure container mapping for the share.
	AzureContainerInfo *AzureContainerInfoResponse `pulumi:"azureContainerInfo"`
	// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
	ClientAccessRights []ClientAccessRightResponse `pulumi:"clientAccessRights"`
	// Data policy of the share.
	DataPolicy *string `pulumi:"dataPolicy"`
	// Description for the share.
	Description *string `pulumi:"description"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// Current monitoring status of the share.
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The object name.
	Name string `pulumi:"name"`
	// Details of the refresh job on this share.
	RefreshDetails *RefreshDetailsResponse `pulumi:"refreshDetails"`
	// Share mount point to the role.
	ShareMappings []MountPointMapResponse `pulumi:"shareMappings"`
	// Current status of the share.
	ShareStatus string `pulumi:"shareStatus"`
	// Metadata pertaining to creation and last modification of Share
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// Mapping of users and corresponding access rights on the share (required for SMB protocol).
	UserAccessRights []UserAccessRightResponse `pulumi:"userAccessRights"`
}

func LookupShareOutput(ctx *pulumi.Context, args LookupShareOutputArgs, opts ...pulumi.InvokeOption) LookupShareResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupShareResultOutput, error) {
			args := v.(LookupShareArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:databoxedge:getShare", args, LookupShareResultOutput{}, options).(LookupShareResultOutput), nil
		}).(LookupShareResultOutput)
}

type LookupShareOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The share name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}

// Represents a share on the  Data Box Edge/Gateway device.
type LookupShareResultOutput struct{ *pulumi.OutputState }

func (LookupShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareResult)(nil)).Elem()
}

func (o LookupShareResultOutput) ToLookupShareResultOutput() LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToLookupShareResultOutputWithContext(ctx context.Context) LookupShareResultOutput {
	return o
}

// Access protocol to be used by the share.
func (o LookupShareResultOutput) AccessProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.AccessProtocol }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupShareResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Azure container mapping for the share.
func (o LookupShareResultOutput) AzureContainerInfo() AzureContainerInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupShareResult) *AzureContainerInfoResponse { return v.AzureContainerInfo }).(AzureContainerInfoResponsePtrOutput)
}

// List of IP addresses and corresponding access rights on the share(required for NFS protocol).
func (o LookupShareResultOutput) ClientAccessRights() ClientAccessRightResponseArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []ClientAccessRightResponse { return v.ClientAccessRights }).(ClientAccessRightResponseArrayOutput)
}

// Data policy of the share.
func (o LookupShareResultOutput) DataPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.DataPolicy }).(pulumi.StringPtrOutput)
}

// Description for the share.
func (o LookupShareResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Id }).(pulumi.StringOutput)
}

// Current monitoring status of the share.
func (o LookupShareResultOutput) MonitoringStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.MonitoringStatus }).(pulumi.StringOutput)
}

// The object name.
func (o LookupShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Name }).(pulumi.StringOutput)
}

// Details of the refresh job on this share.
func (o LookupShareResultOutput) RefreshDetails() RefreshDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupShareResult) *RefreshDetailsResponse { return v.RefreshDetails }).(RefreshDetailsResponsePtrOutput)
}

// Share mount point to the role.
func (o LookupShareResultOutput) ShareMappings() MountPointMapResponseArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []MountPointMapResponse { return v.ShareMappings }).(MountPointMapResponseArrayOutput)
}

// Current status of the share.
func (o LookupShareResultOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ShareStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Share
func (o LookupShareResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupShareResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Type }).(pulumi.StringOutput)
}

// Mapping of users and corresponding access rights on the share (required for SMB protocol).
func (o LookupShareResultOutput) UserAccessRights() UserAccessRightResponseArrayOutput {
	return o.ApplyT(func(v LookupShareResult) []UserAccessRightResponse { return v.UserAccessRights }).(UserAccessRightResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareResultOutput{})
}
