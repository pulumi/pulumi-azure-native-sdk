// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200602preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// **Gets an access token for live metrics stream data.**
func GetLiveToken(ctx *pulumi.Context, args *GetLiveTokenArgs, opts ...pulumi.InvokeOption) (*GetLiveTokenResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetLiveTokenResult
	err := ctx.Invoke("azure-native:insights/v20200602preview:getLiveToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLiveTokenArgs struct {
	// The identifier of the resource.
	ResourceUri string `pulumi:"resourceUri"`
}

// The response to a live token query.
type GetLiveTokenResult struct {
	// JWT token for accessing live metrics stream data.
	LiveToken string `pulumi:"liveToken"`
}

func GetLiveTokenOutput(ctx *pulumi.Context, args GetLiveTokenOutputArgs, opts ...pulumi.InvokeOption) GetLiveTokenResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLiveTokenResultOutput, error) {
			args := v.(GetLiveTokenArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:insights/v20200602preview:getLiveToken", args, GetLiveTokenResultOutput{}, options).(GetLiveTokenResultOutput), nil
		}).(GetLiveTokenResultOutput)
}

type GetLiveTokenOutputArgs struct {
	// The identifier of the resource.
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (GetLiveTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveTokenArgs)(nil)).Elem()
}

// The response to a live token query.
type GetLiveTokenResultOutput struct{ *pulumi.OutputState }

func (GetLiveTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveTokenResult)(nil)).Elem()
}

func (o GetLiveTokenResultOutput) ToGetLiveTokenResultOutput() GetLiveTokenResultOutput {
	return o
}

func (o GetLiveTokenResultOutput) ToGetLiveTokenResultOutputWithContext(ctx context.Context) GetLiveTokenResultOutput {
	return o
}

// JWT token for accessing live metrics stream data.
func (o GetLiveTokenResultOutput) LiveToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetLiveTokenResult) string { return v.LiveToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLiveTokenResultOutput{})
}
