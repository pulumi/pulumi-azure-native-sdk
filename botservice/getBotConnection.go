// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package botservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Connection Setting registration for a Bot Service
//
// Uses Azure REST API version 2023-09-15-preview.
//
// Other available API versions: 2022-09-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native botservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupBotConnection(ctx *pulumi.Context, args *LookupBotConnectionArgs, opts ...pulumi.InvokeOption) (*LookupBotConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBotConnectionResult
	err := ctx.Invoke("azure-native:botservice:getBotConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBotConnectionArgs struct {
	// The name of the Bot Service Connection Setting resource.
	ConnectionName string `pulumi:"connectionName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
}

// Bot channel resource definition
type LookupBotConnectionResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Entity Tag.
	Etag *string `pulumi:"etag"`
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// Required. Gets or sets the Kind of the resource.
	Kind *string `pulumi:"kind"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The set of properties specific to bot channel resource
	Properties ConnectionSettingPropertiesResponse `pulumi:"properties"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
	// Entity zones
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for LookupBotConnectionResult
func (val *LookupBotConnectionResult) Defaults() *LookupBotConnectionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupBotConnectionOutput(ctx *pulumi.Context, args LookupBotConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupBotConnectionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBotConnectionResultOutput, error) {
			args := v.(LookupBotConnectionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:botservice:getBotConnection", args, LookupBotConnectionResultOutput{}, options).(LookupBotConnectionResultOutput), nil
		}).(LookupBotConnectionResultOutput)
}

type LookupBotConnectionOutputArgs struct {
	// The name of the Bot Service Connection Setting resource.
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupBotConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotConnectionArgs)(nil)).Elem()
}

// Bot channel resource definition
type LookupBotConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupBotConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotConnectionResult)(nil)).Elem()
}

func (o LookupBotConnectionResultOutput) ToLookupBotConnectionResultOutput() LookupBotConnectionResultOutput {
	return o
}

func (o LookupBotConnectionResultOutput) ToLookupBotConnectionResultOutputWithContext(ctx context.Context) LookupBotConnectionResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupBotConnectionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Entity Tag.
func (o LookupBotConnectionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Specifies the resource ID.
func (o LookupBotConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Required. Gets or sets the Kind of the resource.
func (o LookupBotConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o LookupBotConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupBotConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to bot channel resource
func (o LookupBotConnectionResultOutput) Properties() ConnectionSettingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) ConnectionSettingPropertiesResponse { return v.Properties }).(ConnectionSettingPropertiesResponseOutput)
}

// Gets or sets the SKU of the resource.
func (o LookupBotConnectionResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupBotConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupBotConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Entity zones
func (o LookupBotConnectionResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBotConnectionResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBotConnectionResultOutput{})
}
