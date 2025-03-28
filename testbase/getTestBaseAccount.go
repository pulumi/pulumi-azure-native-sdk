// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Test Base Account.
//
// Uses Azure REST API version 2022-04-01-preview.
//
// Other available API versions: 2023-11-01-preview.
func LookupTestBaseAccount(ctx *pulumi.Context, args *LookupTestBaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupTestBaseAccountResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupTestBaseAccountResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestBaseAccountArgs struct {
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}

// The Test Base Account resource.
type LookupTestBaseAccountResult struct {
	// The access level of the Test Base Account.
	AccessLevel string `pulumi:"accessLevel"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The SKU of the Test Base Account.
	Sku TestBaseAccountSKUResponse `pulumi:"sku"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}

func LookupTestBaseAccountOutput(ctx *pulumi.Context, args LookupTestBaseAccountOutputArgs, opts ...pulumi.InvokeOption) LookupTestBaseAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTestBaseAccountResultOutput, error) {
			args := v.(LookupTestBaseAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:testbase:getTestBaseAccount", args, LookupTestBaseAccountResultOutput{}, options).(LookupTestBaseAccountResultOutput), nil
		}).(LookupTestBaseAccountResultOutput)
}

type LookupTestBaseAccountOutputArgs struct {
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupTestBaseAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestBaseAccountArgs)(nil)).Elem()
}

// The Test Base Account resource.
type LookupTestBaseAccountResultOutput struct{ *pulumi.OutputState }

func (LookupTestBaseAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestBaseAccountResult)(nil)).Elem()
}

func (o LookupTestBaseAccountResultOutput) ToLookupTestBaseAccountResultOutput() LookupTestBaseAccountResultOutput {
	return o
}

func (o LookupTestBaseAccountResultOutput) ToLookupTestBaseAccountResultOutputWithContext(ctx context.Context) LookupTestBaseAccountResultOutput {
	return o
}

// The access level of the Test Base Account.
func (o LookupTestBaseAccountResultOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.AccessLevel }).(pulumi.StringOutput)
}

// Resource Etag.
func (o LookupTestBaseAccountResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Resource ID.
func (o LookupTestBaseAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupTestBaseAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupTestBaseAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o LookupTestBaseAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the Test Base Account.
func (o LookupTestBaseAccountResultOutput) Sku() TestBaseAccountSKUResponseOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) TestBaseAccountSKUResponse { return v.Sku }).(TestBaseAccountSKUResponseOutput)
}

// The system metadata relating to this resource
func (o LookupTestBaseAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o LookupTestBaseAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o LookupTestBaseAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestBaseAccountResultOutput{})
}
