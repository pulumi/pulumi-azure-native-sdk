// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220130preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Database Migration Service resource
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azure-native:datamigration/v20220130preview:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	GroupName   string `pulumi:"groupName"`
	ServiceName string `pulumi:"serviceName"`
}

// A Database Migration Service resource
type LookupServiceResult struct {
	AutoStopDelay         *string             `pulumi:"autoStopDelay"`
	DeleteResourcesOnStop *bool               `pulumi:"deleteResourcesOnStop"`
	Etag                  *string             `pulumi:"etag"`
	Id                    string              `pulumi:"id"`
	Kind                  *string             `pulumi:"kind"`
	Location              *string             `pulumi:"location"`
	Name                  string              `pulumi:"name"`
	ProvisioningState     string              `pulumi:"provisioningState"`
	PublicKey             *string             `pulumi:"publicKey"`
	Sku                   *ServiceSkuResponse `pulumi:"sku"`
	SystemData            SystemDataResponse  `pulumi:"systemData"`
	Tags                  map[string]string   `pulumi:"tags"`
	Type                  string              `pulumi:"type"`
	VirtualNicId          *string             `pulumi:"virtualNicId"`
	VirtualSubnetId       *string             `pulumi:"virtualSubnetId"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	GroupName   pulumi.StringInput `pulumi:"groupName"`
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

func (o LookupServiceResultOutput) AutoStopDelay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.AutoStopDelay }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) DeleteResourcesOnStop() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *bool { return v.DeleteResourcesOnStop }).(pulumi.BoolPtrOutput)
}

func (o LookupServiceResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) Sku() ServiceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *ServiceSkuResponse { return v.Sku }).(ServiceSkuResponsePtrOutput)
}

func (o LookupServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupServiceResultOutput) VirtualNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.VirtualNicId }).(pulumi.StringPtrOutput)
}

func (o LookupServiceResultOutput) VirtualSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceResult) *string { return v.VirtualSubnetId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}