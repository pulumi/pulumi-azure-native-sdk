// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180712

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns an Enterprise Channel specified by the parameters.
func LookupEnterpriseChannel(ctx *pulumi.Context, args *LookupEnterpriseChannelArgs, opts ...pulumi.InvokeOption) (*LookupEnterpriseChannelResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEnterpriseChannelResult
	err := ctx.Invoke("azure-native:botservice/v20180712:getEnterpriseChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnterpriseChannelArgs struct {
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
}

// Enterprise Channel resource definition
type LookupEnterpriseChannelResult struct {
	// Entity Tag
	Etag *string `pulumi:"etag"`
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The set of properties specific to an Enterprise Channel resource.
	Properties EnterpriseChannelPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}

func LookupEnterpriseChannelOutput(ctx *pulumi.Context, args LookupEnterpriseChannelOutputArgs, opts ...pulumi.InvokeOption) LookupEnterpriseChannelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEnterpriseChannelResultOutput, error) {
			args := v.(LookupEnterpriseChannelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:botservice/v20180712:getEnterpriseChannel", args, LookupEnterpriseChannelResultOutput{}, options).(LookupEnterpriseChannelResultOutput), nil
		}).(LookupEnterpriseChannelResultOutput)
}

type LookupEnterpriseChannelOutputArgs struct {
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupEnterpriseChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseChannelArgs)(nil)).Elem()
}

// Enterprise Channel resource definition
type LookupEnterpriseChannelResultOutput struct{ *pulumi.OutputState }

func (LookupEnterpriseChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnterpriseChannelResult)(nil)).Elem()
}

func (o LookupEnterpriseChannelResultOutput) ToLookupEnterpriseChannelResultOutput() LookupEnterpriseChannelResultOutput {
	return o
}

func (o LookupEnterpriseChannelResultOutput) ToLookupEnterpriseChannelResultOutputWithContext(ctx context.Context) LookupEnterpriseChannelResultOutput {
	return o
}

// Entity Tag
func (o LookupEnterpriseChannelResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Specifies the resource ID.
func (o LookupEnterpriseChannelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Required. Gets or sets the Kind of the resource.
func (o LookupEnterpriseChannelResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o LookupEnterpriseChannelResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupEnterpriseChannelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to an Enterprise Channel resource.
func (o LookupEnterpriseChannelResultOutput) Properties() EnterpriseChannelPropertiesResponseOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) EnterpriseChannelPropertiesResponse { return v.Properties }).(EnterpriseChannelPropertiesResponseOutput)
}

// Gets or sets the SKU of the resource.
func (o LookupEnterpriseChannelResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupEnterpriseChannelResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupEnterpriseChannelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnterpriseChannelResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnterpriseChannelResultOutput{})
}
