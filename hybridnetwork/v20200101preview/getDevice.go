// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Device resource.
func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	var rv LookupDeviceResult
	err := ctx.Invoke("azure-native:hybridnetwork/v20200101preview:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Device resource.
type LookupDeviceResult struct {
	DeviceType        string                `pulumi:"deviceType"`
	Id                string                `pulumi:"id"`
	Location          string                `pulumi:"location"`
	Name              string                `pulumi:"name"`
	NetworkFunctions  []SubResourceResponse `pulumi:"networkFunctions"`
	ProvisioningState string                `pulumi:"provisioningState"`
	Status            string                `pulumi:"status"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
}

func LookupDeviceOutput(ctx *pulumi.Context, args LookupDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceResult, error) {
			args := v.(LookupDeviceArgs)
			r, err := LookupDevice(ctx, &args, opts...)
			var s LookupDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeviceResultOutput)
}

type LookupDeviceOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceArgs)(nil)).Elem()
}

// Device resource.
type LookupDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceResult)(nil)).Elem()
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutput() LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutputWithContext(ctx context.Context) LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) DeviceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.DeviceType }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) NetworkFunctions() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []SubResourceResponse { return v.NetworkFunctions }).(SubResourceResponseArrayOutput)
}

func (o LookupDeviceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupDeviceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDeviceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDeviceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeviceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceResultOutput{})
}