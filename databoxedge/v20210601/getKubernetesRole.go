// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Kubernetes role.
func LookupKubernetesRole(ctx *pulumi.Context, args *LookupKubernetesRoleArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesRoleResult, error) {
	var rv LookupKubernetesRoleResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601:getKubernetesRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKubernetesRoleArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Kubernetes role.
type LookupKubernetesRoleResult struct {
	HostPlatform            string                          `pulumi:"hostPlatform"`
	HostPlatformType        string                          `pulumi:"hostPlatformType"`
	Id                      string                          `pulumi:"id"`
	Kind                    string                          `pulumi:"kind"`
	KubernetesClusterInfo   KubernetesClusterInfoResponse   `pulumi:"kubernetesClusterInfo"`
	KubernetesRoleResources KubernetesRoleResourcesResponse `pulumi:"kubernetesRoleResources"`
	Name                    string                          `pulumi:"name"`
	ProvisioningState       string                          `pulumi:"provisioningState"`
	RoleStatus              string                          `pulumi:"roleStatus"`
	SystemData              SystemDataResponse              `pulumi:"systemData"`
	Type                    string                          `pulumi:"type"`
}

func LookupKubernetesRoleOutput(ctx *pulumi.Context, args LookupKubernetesRoleOutputArgs, opts ...pulumi.InvokeOption) LookupKubernetesRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubernetesRoleResult, error) {
			args := v.(LookupKubernetesRoleArgs)
			r, err := LookupKubernetesRole(ctx, &args, opts...)
			var s LookupKubernetesRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubernetesRoleResultOutput)
}

type LookupKubernetesRoleOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKubernetesRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubernetesRoleArgs)(nil)).Elem()
}

// Kubernetes role.
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

func (o LookupKubernetesRoleResultOutput) HostPlatform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.HostPlatform }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) HostPlatformType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.HostPlatformType }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) KubernetesClusterInfo() KubernetesClusterInfoResponseOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) KubernetesClusterInfoResponse { return v.KubernetesClusterInfo }).(KubernetesClusterInfoResponseOutput)
}

func (o LookupKubernetesRoleResultOutput) KubernetesRoleResources() KubernetesRoleResourcesResponseOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) KubernetesRoleResourcesResponse { return v.KubernetesRoleResources }).(KubernetesRoleResourcesResponseOutput)
}

func (o LookupKubernetesRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) RoleStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.RoleStatus }).(pulumi.StringOutput)
}

func (o LookupKubernetesRoleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupKubernetesRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubernetesRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubernetesRoleResultOutput{})
}