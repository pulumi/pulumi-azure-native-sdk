// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231010preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// support info for firewall.
func GetFirewallSupportInfo(ctx *pulumi.Context, args *GetFirewallSupportInfoArgs, opts ...pulumi.InvokeOption) (*GetFirewallSupportInfoResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetFirewallSupportInfoResult
	err := ctx.Invoke("azure-native:cloudngfw/v20231010preview:getFirewallSupportInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFirewallSupportInfoArgs struct {
	// email address on behalf of which this API called
	Email *string `pulumi:"email"`
	// Firewall resource name
	FirewallName string `pulumi:"firewallName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Support information for the resource
type GetFirewallSupportInfoResult struct {
	// Support account associated with given resource
	AccountId *string `pulumi:"accountId"`
	// account registered in Customer Support Portal
	AccountRegistered *string `pulumi:"accountRegistered"`
	// Product usage is in free trial period
	FreeTrial *string `pulumi:"freeTrial"`
	// Free trial credit remaining
	FreeTrialCreditLeft *int `pulumi:"freeTrialCreditLeft"`
	// Free trial days remaining
	FreeTrialDaysLeft *int `pulumi:"freeTrialDaysLeft"`
	// URL for paloaltonetworks live community
	HelpURL *string `pulumi:"helpURL"`
	// product Serial associated with given resource
	ProductSerial *string `pulumi:"productSerial"`
	// product SKU associated with given resource
	ProductSku *string `pulumi:"productSku"`
	// URL for registering product in paloaltonetworks Customer Service Portal
	RegisterURL *string `pulumi:"registerURL"`
	// URL for paloaltonetworks Customer Service Portal
	SupportURL *string `pulumi:"supportURL"`
	// user domain is supported in Customer Support Portal
	UserDomainSupported *string `pulumi:"userDomainSupported"`
	// user registered in Customer Support Portal
	UserRegistered *string `pulumi:"userRegistered"`
}

func GetFirewallSupportInfoOutput(ctx *pulumi.Context, args GetFirewallSupportInfoOutputArgs, opts ...pulumi.InvokeOption) GetFirewallSupportInfoResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetFirewallSupportInfoResultOutput, error) {
			args := v.(GetFirewallSupportInfoArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:cloudngfw/v20231010preview:getFirewallSupportInfo", args, GetFirewallSupportInfoResultOutput{}, options).(GetFirewallSupportInfoResultOutput), nil
		}).(GetFirewallSupportInfoResultOutput)
}

type GetFirewallSupportInfoOutputArgs struct {
	// email address on behalf of which this API called
	Email pulumi.StringPtrInput `pulumi:"email"`
	// Firewall resource name
	FirewallName pulumi.StringInput `pulumi:"firewallName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetFirewallSupportInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallSupportInfoArgs)(nil)).Elem()
}

// Support information for the resource
type GetFirewallSupportInfoResultOutput struct{ *pulumi.OutputState }

func (GetFirewallSupportInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFirewallSupportInfoResult)(nil)).Elem()
}

func (o GetFirewallSupportInfoResultOutput) ToGetFirewallSupportInfoResultOutput() GetFirewallSupportInfoResultOutput {
	return o
}

func (o GetFirewallSupportInfoResultOutput) ToGetFirewallSupportInfoResultOutputWithContext(ctx context.Context) GetFirewallSupportInfoResultOutput {
	return o
}

// Support account associated with given resource
func (o GetFirewallSupportInfoResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

// account registered in Customer Support Portal
func (o GetFirewallSupportInfoResultOutput) AccountRegistered() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.AccountRegistered }).(pulumi.StringPtrOutput)
}

// Product usage is in free trial period
func (o GetFirewallSupportInfoResultOutput) FreeTrial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.FreeTrial }).(pulumi.StringPtrOutput)
}

// Free trial credit remaining
func (o GetFirewallSupportInfoResultOutput) FreeTrialCreditLeft() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *int { return v.FreeTrialCreditLeft }).(pulumi.IntPtrOutput)
}

// Free trial days remaining
func (o GetFirewallSupportInfoResultOutput) FreeTrialDaysLeft() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *int { return v.FreeTrialDaysLeft }).(pulumi.IntPtrOutput)
}

// URL for paloaltonetworks live community
func (o GetFirewallSupportInfoResultOutput) HelpURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.HelpURL }).(pulumi.StringPtrOutput)
}

// product Serial associated with given resource
func (o GetFirewallSupportInfoResultOutput) ProductSerial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.ProductSerial }).(pulumi.StringPtrOutput)
}

// product SKU associated with given resource
func (o GetFirewallSupportInfoResultOutput) ProductSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.ProductSku }).(pulumi.StringPtrOutput)
}

// URL for registering product in paloaltonetworks Customer Service Portal
func (o GetFirewallSupportInfoResultOutput) RegisterURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.RegisterURL }).(pulumi.StringPtrOutput)
}

// URL for paloaltonetworks Customer Service Portal
func (o GetFirewallSupportInfoResultOutput) SupportURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.SupportURL }).(pulumi.StringPtrOutput)
}

// user domain is supported in Customer Support Portal
func (o GetFirewallSupportInfoResultOutput) UserDomainSupported() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.UserDomainSupported }).(pulumi.StringPtrOutput)
}

// user registered in Customer Support Portal
func (o GetFirewallSupportInfoResultOutput) UserRegistered() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFirewallSupportInfoResult) *string { return v.UserRegistered }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFirewallSupportInfoResultOutput{})
}
