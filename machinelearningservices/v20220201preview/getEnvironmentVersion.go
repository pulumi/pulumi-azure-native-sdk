// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Azure Resource Manager resource envelope.
func LookupEnvironmentVersion(ctx *pulumi.Context, args *LookupEnvironmentVersionArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentVersionResult, error) {
	var rv LookupEnvironmentVersionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220201preview:getEnvironmentVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupEnvironmentVersionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Version           string `pulumi:"version"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentVersionResult struct {
	EnvironmentVersionDetails EnvironmentVersionResponse `pulumi:"environmentVersionDetails"`
	Id                        string                     `pulumi:"id"`
	Name                      string                     `pulumi:"name"`
	SystemData                SystemDataResponse         `pulumi:"systemData"`
	Type                      string                     `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupEnvironmentVersionResult
func (val *LookupEnvironmentVersionResult) Defaults() *LookupEnvironmentVersionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.EnvironmentVersionDetails = *tmp.EnvironmentVersionDetails.Defaults()

	return &tmp
}

func LookupEnvironmentVersionOutput(ctx *pulumi.Context, args LookupEnvironmentVersionOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentVersionResult, error) {
			args := v.(LookupEnvironmentVersionArgs)
			r, err := LookupEnvironmentVersion(ctx, &args, opts...)
			var s LookupEnvironmentVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentVersionResultOutput)
}

type LookupEnvironmentVersionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Version           pulumi.StringInput `pulumi:"version"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEnvironmentVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVersionArgs)(nil)).Elem()
}

// Azure Resource Manager resource envelope.
type LookupEnvironmentVersionResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentVersionResult)(nil)).Elem()
}

func (o LookupEnvironmentVersionResultOutput) ToLookupEnvironmentVersionResultOutput() LookupEnvironmentVersionResultOutput {
	return o
}

func (o LookupEnvironmentVersionResultOutput) ToLookupEnvironmentVersionResultOutputWithContext(ctx context.Context) LookupEnvironmentVersionResultOutput {
	return o
}

func (o LookupEnvironmentVersionResultOutput) EnvironmentVersionDetails() EnvironmentVersionResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) EnvironmentVersionResponse { return v.EnvironmentVersionDetails }).(EnvironmentVersionResponseOutput)
}

func (o LookupEnvironmentVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEnvironmentVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEnvironmentVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEnvironmentVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentVersionResultOutput{})
}