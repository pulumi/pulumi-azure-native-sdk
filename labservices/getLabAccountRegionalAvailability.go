// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The response model from the GetRegionalAvailability action
// API Version: 2018-10-15.
func GetLabAccountRegionalAvailability(ctx *pulumi.Context, args *GetLabAccountRegionalAvailabilityArgs, opts ...pulumi.InvokeOption) (*GetLabAccountRegionalAvailabilityResult, error) {
	var rv GetLabAccountRegionalAvailabilityResult
	err := ctx.Invoke("azure-native:labservices:getLabAccountRegionalAvailability", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLabAccountRegionalAvailabilityArgs struct {
	LabAccountName    string `pulumi:"labAccountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The response model from the GetRegionalAvailability action
type GetLabAccountRegionalAvailabilityResult struct {
	RegionalAvailability []RegionalAvailabilityResponse `pulumi:"regionalAvailability"`
}

func GetLabAccountRegionalAvailabilityOutput(ctx *pulumi.Context, args GetLabAccountRegionalAvailabilityOutputArgs, opts ...pulumi.InvokeOption) GetLabAccountRegionalAvailabilityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLabAccountRegionalAvailabilityResult, error) {
			args := v.(GetLabAccountRegionalAvailabilityArgs)
			r, err := GetLabAccountRegionalAvailability(ctx, &args, opts...)
			var s GetLabAccountRegionalAvailabilityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLabAccountRegionalAvailabilityResultOutput)
}

type GetLabAccountRegionalAvailabilityOutputArgs struct {
	LabAccountName    pulumi.StringInput `pulumi:"labAccountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetLabAccountRegionalAvailabilityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLabAccountRegionalAvailabilityArgs)(nil)).Elem()
}

// The response model from the GetRegionalAvailability action
type GetLabAccountRegionalAvailabilityResultOutput struct{ *pulumi.OutputState }

func (GetLabAccountRegionalAvailabilityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLabAccountRegionalAvailabilityResult)(nil)).Elem()
}

func (o GetLabAccountRegionalAvailabilityResultOutput) ToGetLabAccountRegionalAvailabilityResultOutput() GetLabAccountRegionalAvailabilityResultOutput {
	return o
}

func (o GetLabAccountRegionalAvailabilityResultOutput) ToGetLabAccountRegionalAvailabilityResultOutputWithContext(ctx context.Context) GetLabAccountRegionalAvailabilityResultOutput {
	return o
}

func (o GetLabAccountRegionalAvailabilityResultOutput) RegionalAvailability() RegionalAvailabilityResponseArrayOutput {
	return o.ApplyT(func(v GetLabAccountRegionalAvailabilityResult) []RegionalAvailabilityResponse {
		return v.RegionalAvailability
	}).(RegionalAvailabilityResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLabAccountRegionalAvailabilityResultOutput{})
}