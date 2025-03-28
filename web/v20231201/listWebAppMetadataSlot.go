// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Gets the metadata of an app.
func ListWebAppMetadataSlot(ctx *pulumi.Context, args *ListWebAppMetadataSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppMetadataSlotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWebAppMetadataSlotResult
	err := ctx.Invoke("azure-native:web/v20231201:listWebAppMetadataSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppMetadataSlotArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the metadata for the production slot.
	Slot string `pulumi:"slot"`
}

// String dictionary resource.
type ListWebAppMetadataSlotResult struct {
	// Resource Id.
	Id string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}

func ListWebAppMetadataSlotOutput(ctx *pulumi.Context, args ListWebAppMetadataSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppMetadataSlotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (ListWebAppMetadataSlotResultOutput, error) {
			args := v.(ListWebAppMetadataSlotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web/v20231201:listWebAppMetadataSlot", args, ListWebAppMetadataSlotResultOutput{}, options).(ListWebAppMetadataSlotResultOutput), nil
		}).(ListWebAppMetadataSlotResultOutput)
}

type ListWebAppMetadataSlotOutputArgs struct {
	// Name of the app.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will get the metadata for the production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppMetadataSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppMetadataSlotArgs)(nil)).Elem()
}

// String dictionary resource.
type ListWebAppMetadataSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppMetadataSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppMetadataSlotResult)(nil)).Elem()
}

func (o ListWebAppMetadataSlotResultOutput) ToListWebAppMetadataSlotResultOutput() ListWebAppMetadataSlotResultOutput {
	return o
}

func (o ListWebAppMetadataSlotResultOutput) ToListWebAppMetadataSlotResultOutputWithContext(ctx context.Context) ListWebAppMetadataSlotResultOutput {
	return o
}

// Resource Id.
func (o ListWebAppMetadataSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of resource.
func (o ListWebAppMetadataSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o ListWebAppMetadataSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Settings.
func (o ListWebAppMetadataSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

// Resource type.
func (o ListWebAppMetadataSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppMetadataSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppMetadataSlotResultOutput{})
}
