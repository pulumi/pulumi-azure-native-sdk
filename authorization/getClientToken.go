// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this function to get an Azure authentication token for the current login context.
func GetClientToken(ctx *pulumi.Context, args *GetClientTokenArgs, opts ...pulumi.InvokeOption) (*GetClientTokenResult, error) {
	var rv GetClientTokenResult
	err := ctx.Invoke("azure-native:authorization:getClientToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetClientTokenArgs struct {
	Endpoint *string `pulumi:"endpoint"`
}

// Configuration values returned by getClientToken.
type GetClientTokenResult struct {
	Token string `pulumi:"token"`
}

func GetClientTokenOutput(ctx *pulumi.Context, args GetClientTokenOutputArgs, opts ...pulumi.InvokeOption) GetClientTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetClientTokenResult, error) {
			args := v.(GetClientTokenArgs)
			r, err := GetClientToken(ctx, &args, opts...)
			var s GetClientTokenResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetClientTokenResultOutput)
}

type GetClientTokenOutputArgs struct {
	Endpoint pulumi.StringPtrInput `pulumi:"endpoint"`
}

func (GetClientTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientTokenArgs)(nil)).Elem()
}

// Configuration values returned by getClientToken.
type GetClientTokenResultOutput struct{ *pulumi.OutputState }

func (GetClientTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientTokenResult)(nil)).Elem()
}

func (o GetClientTokenResultOutput) ToGetClientTokenResultOutput() GetClientTokenResultOutput {
	return o
}

func (o GetClientTokenResultOutput) ToGetClientTokenResultOutputWithContext(ctx context.Context) GetClientTokenResultOutput {
	return o
}

func (o GetClientTokenResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientTokenResult) string { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientTokenResultOutput{})
}