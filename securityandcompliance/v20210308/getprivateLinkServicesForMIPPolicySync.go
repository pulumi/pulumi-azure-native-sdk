// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210308

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The description of the service.
func GetprivateLinkServicesForMIPPolicySync(ctx *pulumi.Context, args *GetprivateLinkServicesForMIPPolicySyncArgs, opts ...pulumi.InvokeOption) (*GetprivateLinkServicesForMIPPolicySyncResult, error) {
	var rv GetprivateLinkServicesForMIPPolicySyncResult
	err := ctx.Invoke("azure-native:securityandcompliance/v20210308:getprivateLinkServicesForMIPPolicySync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetprivateLinkServicesForMIPPolicySyncArgs struct {
	// The name of the resource group that contains the service instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName string `pulumi:"resourceName"`
}

// The description of the service.
type GetprivateLinkServicesForMIPPolicySyncResult struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceResponseIdentity `pulumi:"identity"`
	// The kind of the service.
	Kind string `pulumi:"kind"`
	// The resource location.
	Location string `pulumi:"location"`
	// The resource name.
	Name string `pulumi:"name"`
	// The common properties of a service.
	Properties ServicesPropertiesResponse `pulumi:"properties"`
	// Required property for system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}

func GetprivateLinkServicesForMIPPolicySyncOutput(ctx *pulumi.Context, args GetprivateLinkServicesForMIPPolicySyncOutputArgs, opts ...pulumi.InvokeOption) GetprivateLinkServicesForMIPPolicySyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetprivateLinkServicesForMIPPolicySyncResult, error) {
			args := v.(GetprivateLinkServicesForMIPPolicySyncArgs)
			r, err := GetprivateLinkServicesForMIPPolicySync(ctx, &args, opts...)
			var s GetprivateLinkServicesForMIPPolicySyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetprivateLinkServicesForMIPPolicySyncResultOutput)
}

type GetprivateLinkServicesForMIPPolicySyncOutputArgs struct {
	// The name of the resource group that contains the service instance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the service instance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (GetprivateLinkServicesForMIPPolicySyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForMIPPolicySyncArgs)(nil)).Elem()
}

// The description of the service.
type GetprivateLinkServicesForMIPPolicySyncResultOutput struct{ *pulumi.OutputState }

func (GetprivateLinkServicesForMIPPolicySyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetprivateLinkServicesForMIPPolicySyncResult)(nil)).Elem()
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) ToGetprivateLinkServicesForMIPPolicySyncResultOutput() GetprivateLinkServicesForMIPPolicySyncResultOutput {
	return o
}

func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) ToGetprivateLinkServicesForMIPPolicySyncResultOutputWithContext(ctx context.Context) GetprivateLinkServicesForMIPPolicySyncResultOutput {
	return o
}

// An etag associated with the resource, used for optimistic concurrency when editing it.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The resource identifier.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Id }).(pulumi.StringOutput)
}

// Setting indicating whether the service has a managed identity associated with it.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Identity() ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) *ServicesResourceResponseIdentity {
		return v.Identity
	}).(ServicesResourceResponseIdentityPtrOutput)
}

// The kind of the service.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The resource location.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Location }).(pulumi.StringOutput)
}

// The resource name.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Name }).(pulumi.StringOutput)
}

// The common properties of a service.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Properties() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) ServicesPropertiesResponse { return v.Properties }).(ServicesPropertiesResponseOutput)
}

// Required property for system data
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource tags.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The resource type.
func (o GetprivateLinkServicesForMIPPolicySyncResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetprivateLinkServicesForMIPPolicySyncResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetprivateLinkServicesForMIPPolicySyncResultOutput{})
}
