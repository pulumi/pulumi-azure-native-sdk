// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The registration definition.
func LookupRegistrationDefinition(ctx *pulumi.Context, args *LookupRegistrationDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationDefinitionResult, error) {
	var rv LookupRegistrationDefinitionResult
	err := ctx.Invoke("azure-native:managedservices/v20220101preview:getRegistrationDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationDefinitionArgs struct {
	RegistrationDefinitionId string `pulumi:"registrationDefinitionId"`
	Scope                    string `pulumi:"scope"`
}

// The registration definition.
type LookupRegistrationDefinitionResult struct {
	Id         string                                   `pulumi:"id"`
	Name       string                                   `pulumi:"name"`
	Plan       *PlanResponse                            `pulumi:"plan"`
	Properties RegistrationDefinitionPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                       `pulumi:"systemData"`
	Type       string                                   `pulumi:"type"`
}

func LookupRegistrationDefinitionOutput(ctx *pulumi.Context, args LookupRegistrationDefinitionOutputArgs, opts ...pulumi.InvokeOption) LookupRegistrationDefinitionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistrationDefinitionResult, error) {
			args := v.(LookupRegistrationDefinitionArgs)
			r, err := LookupRegistrationDefinition(ctx, &args, opts...)
			var s LookupRegistrationDefinitionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistrationDefinitionResultOutput)
}

type LookupRegistrationDefinitionOutputArgs struct {
	RegistrationDefinitionId pulumi.StringInput `pulumi:"registrationDefinitionId"`
	Scope                    pulumi.StringInput `pulumi:"scope"`
}

func (LookupRegistrationDefinitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationDefinitionArgs)(nil)).Elem()
}

// The registration definition.
type LookupRegistrationDefinitionResultOutput struct{ *pulumi.OutputState }

func (LookupRegistrationDefinitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationDefinitionResult)(nil)).Elem()
}

func (o LookupRegistrationDefinitionResultOutput) ToLookupRegistrationDefinitionResultOutput() LookupRegistrationDefinitionResultOutput {
	return o
}

func (o LookupRegistrationDefinitionResultOutput) ToLookupRegistrationDefinitionResultOutputWithContext(ctx context.Context) LookupRegistrationDefinitionResultOutput {
	return o
}

func (o LookupRegistrationDefinitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Properties() RegistrationDefinitionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) RegistrationDefinitionPropertiesResponse {
		return v.Properties
	}).(RegistrationDefinitionPropertiesResponseOutput)
}

func (o LookupRegistrationDefinitionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRegistrationDefinitionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationDefinitionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistrationDefinitionResultOutput{})
}