// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The resource containing the Azure Stack activation key.
func GetRegistrationActivationKey(ctx *pulumi.Context, args *GetRegistrationActivationKeyArgs, opts ...pulumi.InvokeOption) (*GetRegistrationActivationKeyResult, error) {
	var rv GetRegistrationActivationKeyResult
	err := ctx.Invoke("azure-native:azurestack/v20200601preview:getRegistrationActivationKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetRegistrationActivationKeyArgs struct {
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}

// The resource containing the Azure Stack activation key.
type GetRegistrationActivationKeyResult struct {
	ActivationKey *string `pulumi:"activationKey"`
}

func GetRegistrationActivationKeyOutput(ctx *pulumi.Context, args GetRegistrationActivationKeyOutputArgs, opts ...pulumi.InvokeOption) GetRegistrationActivationKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegistrationActivationKeyResult, error) {
			args := v.(GetRegistrationActivationKeyArgs)
			r, err := GetRegistrationActivationKey(ctx, &args, opts...)
			var s GetRegistrationActivationKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegistrationActivationKeyResultOutput)
}

type GetRegistrationActivationKeyOutputArgs struct {
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup    pulumi.StringInput `pulumi:"resourceGroup"`
}

func (GetRegistrationActivationKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistrationActivationKeyArgs)(nil)).Elem()
}

// The resource containing the Azure Stack activation key.
type GetRegistrationActivationKeyResultOutput struct{ *pulumi.OutputState }

func (GetRegistrationActivationKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegistrationActivationKeyResult)(nil)).Elem()
}

func (o GetRegistrationActivationKeyResultOutput) ToGetRegistrationActivationKeyResultOutput() GetRegistrationActivationKeyResultOutput {
	return o
}

func (o GetRegistrationActivationKeyResultOutput) ToGetRegistrationActivationKeyResultOutputWithContext(ctx context.Context) GetRegistrationActivationKeyResultOutput {
	return o
}

func (o GetRegistrationActivationKeyResultOutput) ActivationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegistrationActivationKeyResult) *string { return v.ActivationKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegistrationActivationKeyResultOutput{})
}