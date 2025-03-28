// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20241201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The variable column.
type PolicyVariableColumn struct {
	// The name of this policy variable column.
	ColumnName string `pulumi:"columnName"`
}

// PolicyVariableColumnInput is an input type that accepts PolicyVariableColumnArgs and PolicyVariableColumnOutput values.
// You can construct a concrete instance of `PolicyVariableColumnInput` via:
//
//	PolicyVariableColumnArgs{...}
type PolicyVariableColumnInput interface {
	pulumi.Input

	ToPolicyVariableColumnOutput() PolicyVariableColumnOutput
	ToPolicyVariableColumnOutputWithContext(context.Context) PolicyVariableColumnOutput
}

// The variable column.
type PolicyVariableColumnArgs struct {
	// The name of this policy variable column.
	ColumnName pulumi.StringInput `pulumi:"columnName"`
}

func (PolicyVariableColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableColumn)(nil)).Elem()
}

func (i PolicyVariableColumnArgs) ToPolicyVariableColumnOutput() PolicyVariableColumnOutput {
	return i.ToPolicyVariableColumnOutputWithContext(context.Background())
}

func (i PolicyVariableColumnArgs) ToPolicyVariableColumnOutputWithContext(ctx context.Context) PolicyVariableColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableColumnOutput)
}

// PolicyVariableColumnArrayInput is an input type that accepts PolicyVariableColumnArray and PolicyVariableColumnArrayOutput values.
// You can construct a concrete instance of `PolicyVariableColumnArrayInput` via:
//
//	PolicyVariableColumnArray{ PolicyVariableColumnArgs{...} }
type PolicyVariableColumnArrayInput interface {
	pulumi.Input

	ToPolicyVariableColumnArrayOutput() PolicyVariableColumnArrayOutput
	ToPolicyVariableColumnArrayOutputWithContext(context.Context) PolicyVariableColumnArrayOutput
}

type PolicyVariableColumnArray []PolicyVariableColumnInput

func (PolicyVariableColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableColumn)(nil)).Elem()
}

func (i PolicyVariableColumnArray) ToPolicyVariableColumnArrayOutput() PolicyVariableColumnArrayOutput {
	return i.ToPolicyVariableColumnArrayOutputWithContext(context.Background())
}

func (i PolicyVariableColumnArray) ToPolicyVariableColumnArrayOutputWithContext(ctx context.Context) PolicyVariableColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableColumnArrayOutput)
}

// The variable column.
type PolicyVariableColumnOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableColumn)(nil)).Elem()
}

func (o PolicyVariableColumnOutput) ToPolicyVariableColumnOutput() PolicyVariableColumnOutput {
	return o
}

func (o PolicyVariableColumnOutput) ToPolicyVariableColumnOutputWithContext(ctx context.Context) PolicyVariableColumnOutput {
	return o
}

// The name of this policy variable column.
func (o PolicyVariableColumnOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableColumn) string { return v.ColumnName }).(pulumi.StringOutput)
}

type PolicyVariableColumnArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableColumn)(nil)).Elem()
}

func (o PolicyVariableColumnArrayOutput) ToPolicyVariableColumnArrayOutput() PolicyVariableColumnArrayOutput {
	return o
}

func (o PolicyVariableColumnArrayOutput) ToPolicyVariableColumnArrayOutputWithContext(ctx context.Context) PolicyVariableColumnArrayOutput {
	return o
}

func (o PolicyVariableColumnArrayOutput) Index(i pulumi.IntInput) PolicyVariableColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableColumn {
		return vs[0].([]PolicyVariableColumn)[vs[1].(int)]
	}).(PolicyVariableColumnOutput)
}

// The variable column.
type PolicyVariableColumnResponse struct {
	// The name of this policy variable column.
	ColumnName string `pulumi:"columnName"`
}

// The variable column.
type PolicyVariableColumnResponseOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableColumnResponse)(nil)).Elem()
}

func (o PolicyVariableColumnResponseOutput) ToPolicyVariableColumnResponseOutput() PolicyVariableColumnResponseOutput {
	return o
}

func (o PolicyVariableColumnResponseOutput) ToPolicyVariableColumnResponseOutputWithContext(ctx context.Context) PolicyVariableColumnResponseOutput {
	return o
}

// The name of this policy variable column.
func (o PolicyVariableColumnResponseOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableColumnResponse) string { return v.ColumnName }).(pulumi.StringOutput)
}

type PolicyVariableColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableColumnResponse)(nil)).Elem()
}

func (o PolicyVariableColumnResponseArrayOutput) ToPolicyVariableColumnResponseArrayOutput() PolicyVariableColumnResponseArrayOutput {
	return o
}

func (o PolicyVariableColumnResponseArrayOutput) ToPolicyVariableColumnResponseArrayOutputWithContext(ctx context.Context) PolicyVariableColumnResponseArrayOutput {
	return o
}

func (o PolicyVariableColumnResponseArrayOutput) Index(i pulumi.IntInput) PolicyVariableColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableColumnResponse {
		return vs[0].([]PolicyVariableColumnResponse)[vs[1].(int)]
	}).(PolicyVariableColumnResponseOutput)
}

// The name value tuple for this variable value column.
type PolicyVariableValueColumnValue struct {
	// Column name for the variable value
	ColumnName string `pulumi:"columnName"`
	// Column value for the variable value; this can be an integer, double, boolean, null or a string.
	ColumnValue interface{} `pulumi:"columnValue"`
}

// PolicyVariableValueColumnValueInput is an input type that accepts PolicyVariableValueColumnValueArgs and PolicyVariableValueColumnValueOutput values.
// You can construct a concrete instance of `PolicyVariableValueColumnValueInput` via:
//
//	PolicyVariableValueColumnValueArgs{...}
type PolicyVariableValueColumnValueInput interface {
	pulumi.Input

	ToPolicyVariableValueColumnValueOutput() PolicyVariableValueColumnValueOutput
	ToPolicyVariableValueColumnValueOutputWithContext(context.Context) PolicyVariableValueColumnValueOutput
}

// The name value tuple for this variable value column.
type PolicyVariableValueColumnValueArgs struct {
	// Column name for the variable value
	ColumnName pulumi.StringInput `pulumi:"columnName"`
	// Column value for the variable value; this can be an integer, double, boolean, null or a string.
	ColumnValue pulumi.Input `pulumi:"columnValue"`
}

func (PolicyVariableValueColumnValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableValueColumnValue)(nil)).Elem()
}

func (i PolicyVariableValueColumnValueArgs) ToPolicyVariableValueColumnValueOutput() PolicyVariableValueColumnValueOutput {
	return i.ToPolicyVariableValueColumnValueOutputWithContext(context.Background())
}

func (i PolicyVariableValueColumnValueArgs) ToPolicyVariableValueColumnValueOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableValueColumnValueOutput)
}

// PolicyVariableValueColumnValueArrayInput is an input type that accepts PolicyVariableValueColumnValueArray and PolicyVariableValueColumnValueArrayOutput values.
// You can construct a concrete instance of `PolicyVariableValueColumnValueArrayInput` via:
//
//	PolicyVariableValueColumnValueArray{ PolicyVariableValueColumnValueArgs{...} }
type PolicyVariableValueColumnValueArrayInput interface {
	pulumi.Input

	ToPolicyVariableValueColumnValueArrayOutput() PolicyVariableValueColumnValueArrayOutput
	ToPolicyVariableValueColumnValueArrayOutputWithContext(context.Context) PolicyVariableValueColumnValueArrayOutput
}

type PolicyVariableValueColumnValueArray []PolicyVariableValueColumnValueInput

func (PolicyVariableValueColumnValueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableValueColumnValue)(nil)).Elem()
}

func (i PolicyVariableValueColumnValueArray) ToPolicyVariableValueColumnValueArrayOutput() PolicyVariableValueColumnValueArrayOutput {
	return i.ToPolicyVariableValueColumnValueArrayOutputWithContext(context.Background())
}

func (i PolicyVariableValueColumnValueArray) ToPolicyVariableValueColumnValueArrayOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyVariableValueColumnValueArrayOutput)
}

// The name value tuple for this variable value column.
type PolicyVariableValueColumnValueOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableValueColumnValue)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueOutput) ToPolicyVariableValueColumnValueOutput() PolicyVariableValueColumnValueOutput {
	return o
}

func (o PolicyVariableValueColumnValueOutput) ToPolicyVariableValueColumnValueOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueOutput {
	return o
}

// Column name for the variable value
func (o PolicyVariableValueColumnValueOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValue) string { return v.ColumnName }).(pulumi.StringOutput)
}

// Column value for the variable value; this can be an integer, double, boolean, null or a string.
func (o PolicyVariableValueColumnValueOutput) ColumnValue() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValue) interface{} { return v.ColumnValue }).(pulumi.AnyOutput)
}

type PolicyVariableValueColumnValueArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableValueColumnValue)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueArrayOutput) ToPolicyVariableValueColumnValueArrayOutput() PolicyVariableValueColumnValueArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueArrayOutput) ToPolicyVariableValueColumnValueArrayOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueArrayOutput) Index(i pulumi.IntInput) PolicyVariableValueColumnValueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableValueColumnValue {
		return vs[0].([]PolicyVariableValueColumnValue)[vs[1].(int)]
	}).(PolicyVariableValueColumnValueOutput)
}

// The name value tuple for this variable value column.
type PolicyVariableValueColumnValueResponse struct {
	// Column name for the variable value
	ColumnName string `pulumi:"columnName"`
	// Column value for the variable value; this can be an integer, double, boolean, null or a string.
	ColumnValue interface{} `pulumi:"columnValue"`
}

// The name value tuple for this variable value column.
type PolicyVariableValueColumnValueResponseOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyVariableValueColumnValueResponse)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueResponseOutput) ToPolicyVariableValueColumnValueResponseOutput() PolicyVariableValueColumnValueResponseOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseOutput) ToPolicyVariableValueColumnValueResponseOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueResponseOutput {
	return o
}

// Column name for the variable value
func (o PolicyVariableValueColumnValueResponseOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValueResponse) string { return v.ColumnName }).(pulumi.StringOutput)
}

// Column value for the variable value; this can be an integer, double, boolean, null or a string.
func (o PolicyVariableValueColumnValueResponseOutput) ColumnValue() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyVariableValueColumnValueResponse) interface{} { return v.ColumnValue }).(pulumi.AnyOutput)
}

type PolicyVariableValueColumnValueResponseArrayOutput struct{ *pulumi.OutputState }

func (PolicyVariableValueColumnValueResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyVariableValueColumnValueResponse)(nil)).Elem()
}

func (o PolicyVariableValueColumnValueResponseArrayOutput) ToPolicyVariableValueColumnValueResponseArrayOutput() PolicyVariableValueColumnValueResponseArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseArrayOutput) ToPolicyVariableValueColumnValueResponseArrayOutputWithContext(ctx context.Context) PolicyVariableValueColumnValueResponseArrayOutput {
	return o
}

func (o PolicyVariableValueColumnValueResponseArrayOutput) Index(i pulumi.IntInput) PolicyVariableValueColumnValueResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyVariableValueColumnValueResponse {
		return vs[0].([]PolicyVariableValueColumnValueResponse)[vs[1].(int)]
	}).(PolicyVariableValueColumnValueResponseOutput)
}

// The resource selector to filter policies by resource properties.
type ResourceSelector struct {
	// The name of the resource selector.
	Name *string `pulumi:"name"`
	// The list of the selector expressions.
	Selectors []Selector `pulumi:"selectors"`
}

// ResourceSelectorInput is an input type that accepts ResourceSelectorArgs and ResourceSelectorOutput values.
// You can construct a concrete instance of `ResourceSelectorInput` via:
//
//	ResourceSelectorArgs{...}
type ResourceSelectorInput interface {
	pulumi.Input

	ToResourceSelectorOutput() ResourceSelectorOutput
	ToResourceSelectorOutputWithContext(context.Context) ResourceSelectorOutput
}

// The resource selector to filter policies by resource properties.
type ResourceSelectorArgs struct {
	// The name of the resource selector.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The list of the selector expressions.
	Selectors SelectorArrayInput `pulumi:"selectors"`
}

func (ResourceSelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelector)(nil)).Elem()
}

func (i ResourceSelectorArgs) ToResourceSelectorOutput() ResourceSelectorOutput {
	return i.ToResourceSelectorOutputWithContext(context.Background())
}

func (i ResourceSelectorArgs) ToResourceSelectorOutputWithContext(ctx context.Context) ResourceSelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSelectorOutput)
}

// ResourceSelectorArrayInput is an input type that accepts ResourceSelectorArray and ResourceSelectorArrayOutput values.
// You can construct a concrete instance of `ResourceSelectorArrayInput` via:
//
//	ResourceSelectorArray{ ResourceSelectorArgs{...} }
type ResourceSelectorArrayInput interface {
	pulumi.Input

	ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput
	ToResourceSelectorArrayOutputWithContext(context.Context) ResourceSelectorArrayOutput
}

type ResourceSelectorArray []ResourceSelectorInput

func (ResourceSelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelector)(nil)).Elem()
}

func (i ResourceSelectorArray) ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput {
	return i.ToResourceSelectorArrayOutputWithContext(context.Background())
}

func (i ResourceSelectorArray) ToResourceSelectorArrayOutputWithContext(ctx context.Context) ResourceSelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSelectorArrayOutput)
}

// The resource selector to filter policies by resource properties.
type ResourceSelectorOutput struct{ *pulumi.OutputState }

func (ResourceSelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelector)(nil)).Elem()
}

func (o ResourceSelectorOutput) ToResourceSelectorOutput() ResourceSelectorOutput {
	return o
}

func (o ResourceSelectorOutput) ToResourceSelectorOutputWithContext(ctx context.Context) ResourceSelectorOutput {
	return o
}

// The name of the resource selector.
func (o ResourceSelectorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSelector) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The list of the selector expressions.
func (o ResourceSelectorOutput) Selectors() SelectorArrayOutput {
	return o.ApplyT(func(v ResourceSelector) []Selector { return v.Selectors }).(SelectorArrayOutput)
}

type ResourceSelectorArrayOutput struct{ *pulumi.OutputState }

func (ResourceSelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelector)(nil)).Elem()
}

func (o ResourceSelectorArrayOutput) ToResourceSelectorArrayOutput() ResourceSelectorArrayOutput {
	return o
}

func (o ResourceSelectorArrayOutput) ToResourceSelectorArrayOutputWithContext(ctx context.Context) ResourceSelectorArrayOutput {
	return o
}

func (o ResourceSelectorArrayOutput) Index(i pulumi.IntInput) ResourceSelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceSelector {
		return vs[0].([]ResourceSelector)[vs[1].(int)]
	}).(ResourceSelectorOutput)
}

// The resource selector to filter policies by resource properties.
type ResourceSelectorResponse struct {
	// The name of the resource selector.
	Name *string `pulumi:"name"`
	// The list of the selector expressions.
	Selectors []SelectorResponse `pulumi:"selectors"`
}

// The resource selector to filter policies by resource properties.
type ResourceSelectorResponseOutput struct{ *pulumi.OutputState }

func (ResourceSelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSelectorResponse)(nil)).Elem()
}

func (o ResourceSelectorResponseOutput) ToResourceSelectorResponseOutput() ResourceSelectorResponseOutput {
	return o
}

func (o ResourceSelectorResponseOutput) ToResourceSelectorResponseOutputWithContext(ctx context.Context) ResourceSelectorResponseOutput {
	return o
}

// The name of the resource selector.
func (o ResourceSelectorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSelectorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The list of the selector expressions.
func (o ResourceSelectorResponseOutput) Selectors() SelectorResponseArrayOutput {
	return o.ApplyT(func(v ResourceSelectorResponse) []SelectorResponse { return v.Selectors }).(SelectorResponseArrayOutput)
}

type ResourceSelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceSelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceSelectorResponse)(nil)).Elem()
}

func (o ResourceSelectorResponseArrayOutput) ToResourceSelectorResponseArrayOutput() ResourceSelectorResponseArrayOutput {
	return o
}

func (o ResourceSelectorResponseArrayOutput) ToResourceSelectorResponseArrayOutputWithContext(ctx context.Context) ResourceSelectorResponseArrayOutput {
	return o
}

func (o ResourceSelectorResponseArrayOutput) Index(i pulumi.IntInput) ResourceSelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceSelectorResponse {
		return vs[0].([]ResourceSelectorResponse)[vs[1].(int)]
	}).(ResourceSelectorResponseOutput)
}

// The selector expression.
type Selector struct {
	// The list of values to filter in.
	In []string `pulumi:"in"`
	// The selector kind.
	Kind *string `pulumi:"kind"`
	// The list of values to filter out.
	NotIn []string `pulumi:"notIn"`
}

// SelectorInput is an input type that accepts SelectorArgs and SelectorOutput values.
// You can construct a concrete instance of `SelectorInput` via:
//
//	SelectorArgs{...}
type SelectorInput interface {
	pulumi.Input

	ToSelectorOutput() SelectorOutput
	ToSelectorOutputWithContext(context.Context) SelectorOutput
}

// The selector expression.
type SelectorArgs struct {
	// The list of values to filter in.
	In pulumi.StringArrayInput `pulumi:"in"`
	// The selector kind.
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// The list of values to filter out.
	NotIn pulumi.StringArrayInput `pulumi:"notIn"`
}

func (SelectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (i SelectorArgs) ToSelectorOutput() SelectorOutput {
	return i.ToSelectorOutputWithContext(context.Background())
}

func (i SelectorArgs) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorOutput)
}

// SelectorArrayInput is an input type that accepts SelectorArray and SelectorArrayOutput values.
// You can construct a concrete instance of `SelectorArrayInput` via:
//
//	SelectorArray{ SelectorArgs{...} }
type SelectorArrayInput interface {
	pulumi.Input

	ToSelectorArrayOutput() SelectorArrayOutput
	ToSelectorArrayOutputWithContext(context.Context) SelectorArrayOutput
}

type SelectorArray []SelectorInput

func (SelectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (i SelectorArray) ToSelectorArrayOutput() SelectorArrayOutput {
	return i.ToSelectorArrayOutputWithContext(context.Background())
}

func (i SelectorArray) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelectorArrayOutput)
}

// The selector expression.
type SelectorOutput struct{ *pulumi.OutputState }

func (SelectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Selector)(nil)).Elem()
}

func (o SelectorOutput) ToSelectorOutput() SelectorOutput {
	return o
}

func (o SelectorOutput) ToSelectorOutputWithContext(ctx context.Context) SelectorOutput {
	return o
}

// The list of values to filter in.
func (o SelectorOutput) In() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Selector) []string { return v.In }).(pulumi.StringArrayOutput)
}

// The selector kind.
func (o SelectorOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Selector) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The list of values to filter out.
func (o SelectorOutput) NotIn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Selector) []string { return v.NotIn }).(pulumi.StringArrayOutput)
}

type SelectorArrayOutput struct{ *pulumi.OutputState }

func (SelectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Selector)(nil)).Elem()
}

func (o SelectorArrayOutput) ToSelectorArrayOutput() SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) ToSelectorArrayOutputWithContext(ctx context.Context) SelectorArrayOutput {
	return o
}

func (o SelectorArrayOutput) Index(i pulumi.IntInput) SelectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Selector {
		return vs[0].([]Selector)[vs[1].(int)]
	}).(SelectorOutput)
}

// The selector expression.
type SelectorResponse struct {
	// The list of values to filter in.
	In []string `pulumi:"in"`
	// The selector kind.
	Kind *string `pulumi:"kind"`
	// The list of values to filter out.
	NotIn []string `pulumi:"notIn"`
}

// The selector expression.
type SelectorResponseOutput struct{ *pulumi.OutputState }

func (SelectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseOutput) ToSelectorResponseOutput() SelectorResponseOutput {
	return o
}

func (o SelectorResponseOutput) ToSelectorResponseOutputWithContext(ctx context.Context) SelectorResponseOutput {
	return o
}

// The list of values to filter in.
func (o SelectorResponseOutput) In() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []string { return v.In }).(pulumi.StringArrayOutput)
}

// The selector kind.
func (o SelectorResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SelectorResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The list of values to filter out.
func (o SelectorResponseOutput) NotIn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SelectorResponse) []string { return v.NotIn }).(pulumi.StringArrayOutput)
}

type SelectorResponseArrayOutput struct{ *pulumi.OutputState }

func (SelectorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelectorResponse)(nil)).Elem()
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutput() SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) ToSelectorResponseArrayOutputWithContext(ctx context.Context) SelectorResponseArrayOutput {
	return o
}

func (o SelectorResponseArrayOutput) Index(i pulumi.IntInput) SelectorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelectorResponse {
		return vs[0].([]SelectorResponse)[vs[1].(int)]
	}).(SelectorResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyVariableColumnOutput{})
	pulumi.RegisterOutputType(PolicyVariableColumnArrayOutput{})
	pulumi.RegisterOutputType(PolicyVariableColumnResponseOutput{})
	pulumi.RegisterOutputType(PolicyVariableColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueArrayOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueResponseOutput{})
	pulumi.RegisterOutputType(PolicyVariableValueColumnValueResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceSelectorOutput{})
	pulumi.RegisterOutputType(ResourceSelectorArrayOutput{})
	pulumi.RegisterOutputType(ResourceSelectorResponseOutput{})
	pulumi.RegisterOutputType(ResourceSelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(SelectorOutput{})
	pulumi.RegisterOutputType(SelectorArrayOutput{})
	pulumi.RegisterOutputType(SelectorResponseOutput{})
	pulumi.RegisterOutputType(SelectorResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
