// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Given a list of volume containers to be failed over from a source device, this method returns the eligibility result, as a failover target, for all devices under that resource.
//
// Uses Azure REST API version 2017-06-01.
func ListDeviceFailoverTars(ctx *pulumi.Context, args *ListDeviceFailoverTarsArgs, opts ...pulumi.InvokeOption) (*ListDeviceFailoverTarsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListDeviceFailoverTarsResult
	err := ctx.Invoke("azure-native:storsimple:listDeviceFailoverTars", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDeviceFailoverTarsArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The source device name on which failover is performed.
	SourceDeviceName string `pulumi:"sourceDeviceName"`
	// The list of path IDs of the volume containers that needs to be failed-over, for which we want to fetch the eligible targets.
	VolumeContainers []string `pulumi:"volumeContainers"`
}

// The list of all devices in a resource and their eligibility status as a failover target device.
type ListDeviceFailoverTarsResult struct {
	// The list of all the failover targets.
	Value []FailoverTargetResponse `pulumi:"value"`
}

func ListDeviceFailoverTarsOutput(ctx *pulumi.Context, args ListDeviceFailoverTarsOutputArgs, opts ...pulumi.InvokeOption) ListDeviceFailoverTarsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListDeviceFailoverTarsResultOutput, error) {
			args := v.(ListDeviceFailoverTarsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:storsimple:listDeviceFailoverTars", args, ListDeviceFailoverTarsResultOutput{}, options).(ListDeviceFailoverTarsResultOutput), nil
		}).(ListDeviceFailoverTarsResultOutput)
}

type ListDeviceFailoverTarsOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The source device name on which failover is performed.
	SourceDeviceName pulumi.StringInput `pulumi:"sourceDeviceName"`
	// The list of path IDs of the volume containers that needs to be failed-over, for which we want to fetch the eligible targets.
	VolumeContainers pulumi.StringArrayInput `pulumi:"volumeContainers"`
}

func (ListDeviceFailoverTarsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceFailoverTarsArgs)(nil)).Elem()
}

// The list of all devices in a resource and their eligibility status as a failover target device.
type ListDeviceFailoverTarsResultOutput struct{ *pulumi.OutputState }

func (ListDeviceFailoverTarsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDeviceFailoverTarsResult)(nil)).Elem()
}

func (o ListDeviceFailoverTarsResultOutput) ToListDeviceFailoverTarsResultOutput() ListDeviceFailoverTarsResultOutput {
	return o
}

func (o ListDeviceFailoverTarsResultOutput) ToListDeviceFailoverTarsResultOutputWithContext(ctx context.Context) ListDeviceFailoverTarsResultOutput {
	return o
}

// The list of all the failover targets.
func (o ListDeviceFailoverTarsResultOutput) Value() FailoverTargetResponseArrayOutput {
	return o.ApplyT(func(v ListDeviceFailoverTarsResult) []FailoverTargetResponse { return v.Value }).(FailoverTargetResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDeviceFailoverTarsResultOutput{})
}
