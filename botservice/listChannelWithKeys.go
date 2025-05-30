// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package botservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists a Channel registration for a Bot Service including secrets
//
// Uses Azure REST API version 2023-09-15-preview.
//
// Other available API versions: 2022-09-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native botservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func ListChannelWithKeys(ctx *pulumi.Context, args *ListChannelWithKeysArgs, opts ...pulumi.InvokeOption) (*ListChannelWithKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListChannelWithKeysResult
	err := ctx.Invoke("azure-native:botservice:listChannelWithKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type ListChannelWithKeysArgs struct {
	// The name of the Channel resource.
	ChannelName string `pulumi:"channelName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName string `pulumi:"resourceName"`
}

// The ARM channel of list channel with keys operation response.
type ListChannelWithKeysResult struct {
	// Changed time of the resource
	ChangedTime *string `pulumi:"changedTime"`
	// Entity tag of the resource
	EntityTag *string `pulumi:"entityTag"`
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
	Properties interface{} `pulumi:"properties"`
	// Provisioning state of the resource
	ProvisioningState *string `pulumi:"provisioningState"`
	// The set of properties specific to bot channel resource
	Resource interface{} `pulumi:"resource"`
	// Channel settings
	Setting *ChannelSettingsResponse `pulumi:"setting"`
	// Gets or sets the SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
	// Entity zones
	Zones []string `pulumi:"zones"`
}

// Defaults sets the appropriate defaults for ListChannelWithKeysResult
func (val *ListChannelWithKeysResult) Defaults() *ListChannelWithKeysResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Setting = tmp.Setting.Defaults()

	return &tmp
}
func ListChannelWithKeysOutput(ctx *pulumi.Context, args ListChannelWithKeysOutputArgs, opts ...pulumi.InvokeOption) ListChannelWithKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListChannelWithKeysResultOutput, error) {
			args := v.(ListChannelWithKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:botservice:listChannelWithKeys", args, ListChannelWithKeysResultOutput{}, options).(ListChannelWithKeysResultOutput), nil
		}).(ListChannelWithKeysResultOutput)
}

type ListChannelWithKeysOutputArgs struct {
	// The name of the Channel resource.
	ChannelName pulumi.StringInput `pulumi:"channelName"`
	// The name of the Bot resource group in the user subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Bot resource.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (ListChannelWithKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListChannelWithKeysArgs)(nil)).Elem()
}

// The ARM channel of list channel with keys operation response.
type ListChannelWithKeysResultOutput struct{ *pulumi.OutputState }

func (ListChannelWithKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListChannelWithKeysResult)(nil)).Elem()
}

func (o ListChannelWithKeysResultOutput) ToListChannelWithKeysResultOutput() ListChannelWithKeysResultOutput {
	return o
}

func (o ListChannelWithKeysResultOutput) ToListChannelWithKeysResultOutputWithContext(ctx context.Context) ListChannelWithKeysResultOutput {
	return o
}

// Changed time of the resource
func (o ListChannelWithKeysResultOutput) ChangedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.ChangedTime }).(pulumi.StringPtrOutput)
}

// Entity tag of the resource
func (o ListChannelWithKeysResultOutput) EntityTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.EntityTag }).(pulumi.StringPtrOutput)
}

// Entity Tag.
func (o ListChannelWithKeysResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Specifies the resource ID.
func (o ListChannelWithKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// Required. Gets or sets the Kind of the resource.
func (o ListChannelWithKeysResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o ListChannelWithKeysResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o ListChannelWithKeysResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) string { return v.Name }).(pulumi.StringOutput)
}

// The set of properties specific to bot channel resource
func (o ListChannelWithKeysResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Provisioning state of the resource
func (o ListChannelWithKeysResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// The set of properties specific to bot channel resource
func (o ListChannelWithKeysResultOutput) Resource() pulumi.AnyOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) interface{} { return v.Resource }).(pulumi.AnyOutput)
}

// Channel settings
func (o ListChannelWithKeysResultOutput) Setting() ChannelSettingsResponsePtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *ChannelSettingsResponse { return v.Setting }).(ChannelSettingsResponsePtrOutput)
}

// Gets or sets the SKU of the resource.
func (o ListChannelWithKeysResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o ListChannelWithKeysResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o ListChannelWithKeysResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) string { return v.Type }).(pulumi.StringOutput)
}

// Entity zones
func (o ListChannelWithKeysResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListChannelWithKeysResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListChannelWithKeysResultOutput{})
}
