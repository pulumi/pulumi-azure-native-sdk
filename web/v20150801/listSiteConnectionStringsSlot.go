// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// String dictionary resource
func ListSiteConnectionStringsSlot(ctx *pulumi.Context, args *ListSiteConnectionStringsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteConnectionStringsSlotResult, error) {
	var rv ListSiteConnectionStringsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteConnectionStringsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteConnectionStringsSlotArgs struct {
	// Name of web app
	Name string `pulumi:"name"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot string `pulumi:"slot"`
}

// String dictionary resource
type ListSiteConnectionStringsSlotResult struct {
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Connection strings
	Properties map[string]ConnStringValueTypePairResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

func ListSiteConnectionStringsSlotOutput(ctx *pulumi.Context, args ListSiteConnectionStringsSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteConnectionStringsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteConnectionStringsSlotResult, error) {
			args := v.(ListSiteConnectionStringsSlotArgs)
			r, err := ListSiteConnectionStringsSlot(ctx, &args, opts...)
			var s ListSiteConnectionStringsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteConnectionStringsSlotResultOutput)
}

type ListSiteConnectionStringsSlotOutputArgs struct {
	// Name of web app
	Name pulumi.StringInput `pulumi:"name"`
	// Name of resource group
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of web app slot. If not specified then will default to production slot.
	Slot pulumi.StringInput `pulumi:"slot"`
}

func (ListSiteConnectionStringsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteConnectionStringsSlotArgs)(nil)).Elem()
}

// String dictionary resource
type ListSiteConnectionStringsSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteConnectionStringsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteConnectionStringsSlotResult)(nil)).Elem()
}

func (o ListSiteConnectionStringsSlotResultOutput) ToListSiteConnectionStringsSlotResultOutput() ListSiteConnectionStringsSlotResultOutput {
	return o
}

func (o ListSiteConnectionStringsSlotResultOutput) ToListSiteConnectionStringsSlotResultOutputWithContext(ctx context.Context) ListSiteConnectionStringsSlotResultOutput {
	return o
}

// Resource Id
func (o ListSiteConnectionStringsSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Kind of resource
func (o ListSiteConnectionStringsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource Location
func (o ListSiteConnectionStringsSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

// Resource Name
func (o ListSiteConnectionStringsSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Connection strings
func (o ListSiteConnectionStringsSlotResultOutput) Properties() ConnStringValueTypePairResponseMapOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) map[string]ConnStringValueTypePairResponse {
		return v.Properties
	}).(ConnStringValueTypePairResponseMapOutput)
}

// Resource tags
func (o ListSiteConnectionStringsSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o ListSiteConnectionStringsSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteConnectionStringsSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteConnectionStringsSlotResultOutput{})
}
