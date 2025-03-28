// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datamigration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The services resource is the top-level resource that represents the Database Migration Service. The GET method retrieves information about a service instance.
//
// Uses Azure REST API version 2021-06-30.
//
// Other available API versions: 2022-03-30-preview, 2023-07-15-preview.
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:datamigration:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	// Name of the resource group
	GroupName string `pulumi:"groupName"`
	// Name of the service
	ServiceName string `pulumi:"serviceName"`
}

// A Database Migration Service resource
type LookupServiceResult struct {
	// HTTP strong entity tag value. Ignored if submitted
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The resource kind. Only 'vm' (the default) is supported.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The resource's provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// The public key of the service, used to encrypt secrets sent to the service
	PublicKey *string `pulumi:"publicKey"`
	// Service SKU
	Sku *ServiceSkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The ID of the Microsoft.Network/networkInterfaces resource which the service have
	VirtualNicId *string `pulumi:"virtualNicId"`
	// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
	VirtualSubnetId string `pulumi:"virtualSubnetId"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServiceResultOutput, error) {
			args := v.(LookupServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datamigration:getService", args, LookupServiceResultOutput{}, options).(LookupServiceResultOutput), nil
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	// Name of the resource group
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// Name of the service
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

// A Database Migration Service resource
type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// HTTP strong entity tag value. Ignored if submitted
func (o LookupServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Resource ID.
func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource kind. Only 'vm' (the default) is supported.
func (o LookupServiceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location.
func (o LookupServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource's provisioning state
func (o LookupServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The public key of the service, used to encrypt secrets sent to the service
func (o LookupServiceResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

// Service SKU
func (o LookupServiceResultOutput) Sku() ServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceSkuResponse { return v.Sku }).(ServiceSkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The ID of the Microsoft.Network/networkInterfaces resource which the service have
func (o LookupServiceResultOutput) VirtualNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.VirtualNicId }).(pulumi.StringPtrOutput)
}

// The ID of the Microsoft.Network/virtualNetworks/subnets resource to which the service should be joined
func (o LookupServiceResultOutput) VirtualSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.VirtualSubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
