// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
// API Version: 2022-02-01-preview.
func LookupComponentVersion(ctx *pulumi.Context, args *LookupComponentVersionArgs, opts ...pulumi.InvokeOption) (*LookupComponentVersionResult, error) {
	var rv LookupComponentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices:getComponentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupComponentVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupComponentVersionResult struct {
	ComponentVersionDetails ComponentVersionResponse `pulumi:"componentVersionDetails"`
	Id                      string                   `pulumi:"id"`
	Name                    string                   `pulumi:"name"`
	SystemData              SystemDataResponse       `pulumi:"systemData"`
	Type                    string                   `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupComponentVersionResult
func (val *LookupComponentVersionResult) Defaults() *LookupComponentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ComponentVersionDetails = *tmp.ComponentVersionDetails.Defaults()

	return &tmp
}

func LookupComponentVersionOutput(ctx *pulumi.Context, args LookupComponentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupComponentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentVersionResult, error) {
			args := v.(LookupComponentVersionArgs)
			r, err := LookupComponentVersion(ctx, &args, opts...)
			var s LookupComponentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentVersionResultOutput)
}

type LookupComponentVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupComponentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupComponentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupComponentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentVersionResult)(nil)).Elem()
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutput() LookupComponentVersionResultOutput {
	return o
}

func (o LookupComponentVersionResultOutput) ToLookupComponentVersionResultOutputWithContext(ctx context.Context) LookupComponentVersionResultOutput {
	return o
}

func (o LookupComponentVersionResultOutput) ComponentVersionDetails() ComponentVersionResponseOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) ComponentVersionResponse { return v.ComponentVersionDetails }).(ComponentVersionResponseOutput)
}

func (o LookupComponentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComponentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComponentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupComponentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComponentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentVersionResultOutput{})
}