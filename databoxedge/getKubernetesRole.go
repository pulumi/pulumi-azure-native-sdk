// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a specific role by name.
//
// Uses Azure REST API version 2023-07-01.
func LookupKubernetesRole(ctx *pulumi.Context, args *LookupKubernetesRoleArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesRoleResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupKubernetesRoleResult
	err := ctx.Invoke("azure-native:databoxedge:getKubernetesRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubernetesRoleArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The role name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The limited preview of Kubernetes Cluster Management from the Azure supports:
//  1. Using a simple turn-key option in Azure Portal, deploy a Kubernetes cluster on your Azure Stack Edge device.
//  2. Configure Kubernetes cluster running on your device with Arc enabled Kubernetes with a click of a button in the Azure Portal.
//     Azure Arc enables organizations to view, manage, and govern their on-premises Kubernetes clusters using the Azure Portal, command line tools, and APIs.
//  3. Easily configure Persistent Volumes using SMB and NFS shares for storing container data.
//     For more information, refer to the document here: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8-Cloud-Management-20210323.pdf
//     Or Demo: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8S-Cloud-Management-20210323.mp4
//     By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/
type LookupKubernetesRoleResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Host OS supported by the Kubernetes role.
	HostPlatform string `pulumi:"hostPlatform"`
	// Platform where the runtime is hosted.
	HostPlatformType string `pulumi:"hostPlatformType"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// Role type.
	// Expected value is 'Kubernetes'.
	Kind string `pulumi:"kind"`
	// Kubernetes cluster configuration
	KubernetesClusterInfo KubernetesClusterInfoResponse `pulumi:"kubernetesClusterInfo"`
	// Kubernetes role resources
	KubernetesRoleResources KubernetesRoleResourcesResponse `pulumi:"kubernetesRoleResources"`
	// The object name.
	Name string `pulumi:"name"`
	// State of Kubernetes deployment
	ProvisioningState string `pulumi:"provisioningState"`
	// Role status.
	RoleStatus string `pulumi:"roleStatus"`
	// Metadata pertaining to creation and last modification of Role
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
}

func LookupKubernetesRoleOutput(ctx *pulumi.Context, args LookupKubernetesRoleOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesRoleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKubernetesRoleResultOutput, error) {
			args := v.(LookupKubernetesRoleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:databoxedge:getKubernetesRole", args, LookupKubernetesRoleResultOutput{}, options).(LookupKubernetesRoleResultOutput), nil
		}).(LookupKubernetesRoleResultOutput)
}

type LookupKubernetesRoleOutputArgs struct {
	// The device name.
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	// The role name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKubernetesRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesRoleArgs)(nil)).Elem()
}

// The limited preview of Kubernetes Cluster Management from the Azure supports:
//  1. Using a simple turn-key option in Azure Portal, deploy a Kubernetes cluster on your Azure Stack Edge device.
//  2. Configure Kubernetes cluster running on your device with Arc enabled Kubernetes with a click of a button in the Azure Portal.
//     Azure Arc enables organizations to view, manage, and govern their on-premises Kubernetes clusters using the Azure Portal, command line tools, and APIs.
//  3. Easily configure Persistent Volumes using SMB and NFS shares for storing container data.
//     For more information, refer to the document here: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8-Cloud-Management-20210323.pdf
//     Or Demo: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8S-Cloud-Management-20210323.mp4
//     By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/
type LookupKubernetesRoleResultOutput struct{ *pulumi.OutputState }

func (LookupKubernetesRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesRoleResult)(nil)).Elem()
}

func (o LookupKubernetesRoleResultOutput) ToLookupKubernetesRoleResultOutput() LookupKubernetesRoleResultOutput {
	return o
}

func (o LookupKubernetesRoleResultOutput) ToLookupKubernetesRoleResultOutputWithContext(ctx context.Context) LookupKubernetesRoleResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupKubernetesRoleResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Host OS supported by the Kubernetes role.
func (o LookupKubernetesRoleResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

// Platform where the runtime is hosted.
func (o LookupKubernetesRoleResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupKubernetesRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Role type.
// Expected value is 'Kubernetes'.
func (o LookupKubernetesRoleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Kubernetes cluster configuration
func (o LookupKubernetesRoleResultOutput) KubernetesClusterInfo() KubernetesClusterInfoResponseOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) KubernetesClusterInfoResponse { return v.KubernetesClusterInfo }).(KubernetesClusterInfoResponseOutput)
}

// Kubernetes role resources
func (o LookupKubernetesRoleResultOutput) KubernetesRoleResources() KubernetesRoleResourcesResponseOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) KubernetesRoleResourcesResponse { return v.KubernetesRoleResources }).(KubernetesRoleResourcesResponseOutput)
}

// The object name.
func (o LookupKubernetesRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// State of Kubernetes deployment
func (o LookupKubernetesRoleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Role status.
func (o LookupKubernetesRoleResultOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.RoleStatus }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of Role
func (o LookupKubernetesRoleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o LookupKubernetesRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesRoleResultOutput{})
}
