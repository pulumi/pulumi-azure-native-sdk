// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Web Job Information.
func LookupWebAppFunction(ctx *pulumi.Context, args *LookupWebAppFunctionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppFunctionResult, error) {
	var rv LookupWebAppFunctionResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppFunction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppFunctionArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Web Job Information.
type LookupWebAppFunctionResult struct {
	Config             interface{}       `pulumi:"config"`
	ConfigHref         *string           `pulumi:"configHref"`
	Files              map[string]string `pulumi:"files"`
	FunctionAppId      *string           `pulumi:"functionAppId"`
	Href               *string           `pulumi:"href"`
	Id                 string            `pulumi:"id"`
	Kind               *string           `pulumi:"kind"`
	Name               string            `pulumi:"name"`
	ScriptHref         *string           `pulumi:"scriptHref"`
	ScriptRootPathHref *string           `pulumi:"scriptRootPathHref"`
	SecretsFileHref    *string           `pulumi:"secretsFileHref"`
	TestData           *string           `pulumi:"testData"`
	Type               string            `pulumi:"type"`
}

func LookupWebAppFunctionOutput(ctx *pulumi.Context, args LookupWebAppFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppFunctionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppFunctionResult, error) {
			args := v.(LookupWebAppFunctionArgs)
			r, err := LookupWebAppFunction(ctx, &args, opts...)
			var s LookupWebAppFunctionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppFunctionResultOutput)
}

type LookupWebAppFunctionOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFunctionArgs)(nil)).Elem()
}

// Web Job Information.
type LookupWebAppFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppFunctionResult)(nil)).Elem()
}

func (o LookupWebAppFunctionResultOutput) ToLookupWebAppFunctionResultOutput() LookupWebAppFunctionResultOutput {
	return o
}

func (o LookupWebAppFunctionResultOutput) ToLookupWebAppFunctionResultOutputWithContext(ctx context.Context) LookupWebAppFunctionResultOutput {
	return o
}

func (o LookupWebAppFunctionResultOutput) Config() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) interface{} { return v.Config }).(pulumi.AnyOutput)
}

func (o LookupWebAppFunctionResultOutput) ConfigHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ConfigHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Files() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) map[string]string { return v.Files }).(pulumi.StringMapOutput)
}

func (o LookupWebAppFunctionResultOutput) FunctionAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.FunctionAppId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Href() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Href }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppFunctionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppFunctionResultOutput) ScriptHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ScriptHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) ScriptRootPathHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.ScriptRootPathHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) SecretsFileHref() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.SecretsFileHref }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) TestData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) *string { return v.TestData }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppFunctionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppFunctionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppFunctionResultOutput{})
}