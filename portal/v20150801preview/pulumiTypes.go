// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A dashboard lens.
type DashboardLens struct {
	// The dashboard len's metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The lens order.
	Order int `pulumi:"order"`
	// The dashboard parts.
	Parts map[string]DashboardParts `pulumi:"parts"`
}

// DashboardLensInput is an input type that accepts DashboardLensArgs and DashboardLensOutput values.
// You can construct a concrete instance of `DashboardLensInput` via:
//
//	DashboardLensArgs{...}
type DashboardLensInput interface {
	pulumi.Input

	ToDashboardLensOutput() DashboardLensOutput
	ToDashboardLensOutputWithContext(context.Context) DashboardLensOutput
}

// A dashboard lens.
type DashboardLensArgs struct {
	// The dashboard len's metadata.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The lens order.
	Order pulumi.IntInput `pulumi:"order"`
	// The dashboard parts.
	Parts DashboardPartsMapInput `pulumi:"parts"`
}

func (DashboardLensArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLens)(nil)).Elem()
}

func (i DashboardLensArgs) ToDashboardLensOutput() DashboardLensOutput {
	return i.ToDashboardLensOutputWithContext(context.Background())
}

func (i DashboardLensArgs) ToDashboardLensOutputWithContext(ctx context.Context) DashboardLensOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensOutput)
}

// DashboardLensMapInput is an input type that accepts DashboardLensMap and DashboardLensMapOutput values.
// You can construct a concrete instance of `DashboardLensMapInput` via:
//
//	DashboardLensMap{ "key": DashboardLensArgs{...} }
type DashboardLensMapInput interface {
	pulumi.Input

	ToDashboardLensMapOutput() DashboardLensMapOutput
	ToDashboardLensMapOutputWithContext(context.Context) DashboardLensMapOutput
}

type DashboardLensMap map[string]DashboardLensInput

func (DashboardLensMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardLens)(nil)).Elem()
}

func (i DashboardLensMap) ToDashboardLensMapOutput() DashboardLensMapOutput {
	return i.ToDashboardLensMapOutputWithContext(context.Background())
}

func (i DashboardLensMap) ToDashboardLensMapOutputWithContext(ctx context.Context) DashboardLensMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardLensMapOutput)
}

// A dashboard lens.
type DashboardLensOutput struct{ *pulumi.OutputState }

func (DashboardLensOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLens)(nil)).Elem()
}

func (o DashboardLensOutput) ToDashboardLensOutput() DashboardLensOutput {
	return o
}

func (o DashboardLensOutput) ToDashboardLensOutputWithContext(ctx context.Context) DashboardLensOutput {
	return o
}

// The dashboard len's metadata.
func (o DashboardLensOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardLens) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The lens order.
func (o DashboardLensOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardLens) int { return v.Order }).(pulumi.IntOutput)
}

// The dashboard parts.
func (o DashboardLensOutput) Parts() DashboardPartsMapOutput {
	return o.ApplyT(func(v DashboardLens) map[string]DashboardParts { return v.Parts }).(DashboardPartsMapOutput)
}

type DashboardLensMapOutput struct{ *pulumi.OutputState }

func (DashboardLensMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardLens)(nil)).Elem()
}

func (o DashboardLensMapOutput) ToDashboardLensMapOutput() DashboardLensMapOutput {
	return o
}

func (o DashboardLensMapOutput) ToDashboardLensMapOutputWithContext(ctx context.Context) DashboardLensMapOutput {
	return o
}

func (o DashboardLensMapOutput) MapIndex(k pulumi.StringInput) DashboardLensOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardLens {
		return vs[0].(map[string]DashboardLens)[vs[1].(string)]
	}).(DashboardLensOutput)
}

// A dashboard lens.
type DashboardLensResponse struct {
	// The dashboard len's metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The lens order.
	Order int `pulumi:"order"`
	// The dashboard parts.
	Parts map[string]DashboardPartsResponse `pulumi:"parts"`
}

// A dashboard lens.
type DashboardLensResponseOutput struct{ *pulumi.OutputState }

func (DashboardLensResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardLensResponse)(nil)).Elem()
}

func (o DashboardLensResponseOutput) ToDashboardLensResponseOutput() DashboardLensResponseOutput {
	return o
}

func (o DashboardLensResponseOutput) ToDashboardLensResponseOutputWithContext(ctx context.Context) DashboardLensResponseOutput {
	return o
}

// The dashboard len's metadata.
func (o DashboardLensResponseOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardLensResponse) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The lens order.
func (o DashboardLensResponseOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardLensResponse) int { return v.Order }).(pulumi.IntOutput)
}

// The dashboard parts.
func (o DashboardLensResponseOutput) Parts() DashboardPartsResponseMapOutput {
	return o.ApplyT(func(v DashboardLensResponse) map[string]DashboardPartsResponse { return v.Parts }).(DashboardPartsResponseMapOutput)
}

type DashboardLensResponseMapOutput struct{ *pulumi.OutputState }

func (DashboardLensResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardLensResponse)(nil)).Elem()
}

func (o DashboardLensResponseMapOutput) ToDashboardLensResponseMapOutput() DashboardLensResponseMapOutput {
	return o
}

func (o DashboardLensResponseMapOutput) ToDashboardLensResponseMapOutputWithContext(ctx context.Context) DashboardLensResponseMapOutput {
	return o
}

func (o DashboardLensResponseMapOutput) MapIndex(k pulumi.StringInput) DashboardLensResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardLensResponse {
		return vs[0].(map[string]DashboardLensResponse)[vs[1].(string)]
	}).(DashboardLensResponseOutput)
}

// A dashboard part.
type DashboardParts struct {
	// A dashboard part metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The dashboard's part position.
	Position DashboardPartsPosition `pulumi:"position"`
}

// DashboardPartsInput is an input type that accepts DashboardPartsArgs and DashboardPartsOutput values.
// You can construct a concrete instance of `DashboardPartsInput` via:
//
//	DashboardPartsArgs{...}
type DashboardPartsInput interface {
	pulumi.Input

	ToDashboardPartsOutput() DashboardPartsOutput
	ToDashboardPartsOutputWithContext(context.Context) DashboardPartsOutput
}

// A dashboard part.
type DashboardPartsArgs struct {
	// A dashboard part metadata.
	Metadata pulumi.Input `pulumi:"metadata"`
	// The dashboard's part position.
	Position DashboardPartsPositionInput `pulumi:"position"`
}

func (DashboardPartsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardParts)(nil)).Elem()
}

func (i DashboardPartsArgs) ToDashboardPartsOutput() DashboardPartsOutput {
	return i.ToDashboardPartsOutputWithContext(context.Background())
}

func (i DashboardPartsArgs) ToDashboardPartsOutputWithContext(ctx context.Context) DashboardPartsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsOutput)
}

// DashboardPartsMapInput is an input type that accepts DashboardPartsMap and DashboardPartsMapOutput values.
// You can construct a concrete instance of `DashboardPartsMapInput` via:
//
//	DashboardPartsMap{ "key": DashboardPartsArgs{...} }
type DashboardPartsMapInput interface {
	pulumi.Input

	ToDashboardPartsMapOutput() DashboardPartsMapOutput
	ToDashboardPartsMapOutputWithContext(context.Context) DashboardPartsMapOutput
}

type DashboardPartsMap map[string]DashboardPartsInput

func (DashboardPartsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardParts)(nil)).Elem()
}

func (i DashboardPartsMap) ToDashboardPartsMapOutput() DashboardPartsMapOutput {
	return i.ToDashboardPartsMapOutputWithContext(context.Background())
}

func (i DashboardPartsMap) ToDashboardPartsMapOutputWithContext(ctx context.Context) DashboardPartsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsMapOutput)
}

// A dashboard part.
type DashboardPartsOutput struct{ *pulumi.OutputState }

func (DashboardPartsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardParts)(nil)).Elem()
}

func (o DashboardPartsOutput) ToDashboardPartsOutput() DashboardPartsOutput {
	return o
}

func (o DashboardPartsOutput) ToDashboardPartsOutputWithContext(ctx context.Context) DashboardPartsOutput {
	return o
}

// A dashboard part metadata.
func (o DashboardPartsOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v DashboardParts) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The dashboard's part position.
func (o DashboardPartsOutput) Position() DashboardPartsPositionOutput {
	return o.ApplyT(func(v DashboardParts) DashboardPartsPosition { return v.Position }).(DashboardPartsPositionOutput)
}

type DashboardPartsMapOutput struct{ *pulumi.OutputState }

func (DashboardPartsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardParts)(nil)).Elem()
}

func (o DashboardPartsMapOutput) ToDashboardPartsMapOutput() DashboardPartsMapOutput {
	return o
}

func (o DashboardPartsMapOutput) ToDashboardPartsMapOutputWithContext(ctx context.Context) DashboardPartsMapOutput {
	return o
}

func (o DashboardPartsMapOutput) MapIndex(k pulumi.StringInput) DashboardPartsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardParts {
		return vs[0].(map[string]DashboardParts)[vs[1].(string)]
	}).(DashboardPartsOutput)
}

// The dashboard's part position.
type DashboardPartsPosition struct {
	// The dashboard's part column span.
	ColSpan int `pulumi:"colSpan"`
	// The dashboard part's metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The dashboard's part row span.
	RowSpan int `pulumi:"rowSpan"`
	// The dashboard's part x coordinate.
	X int `pulumi:"x"`
	// The dashboard's part y coordinate.
	Y int `pulumi:"y"`
}

// DashboardPartsPositionInput is an input type that accepts DashboardPartsPositionArgs and DashboardPartsPositionOutput values.
// You can construct a concrete instance of `DashboardPartsPositionInput` via:
//
//	DashboardPartsPositionArgs{...}
type DashboardPartsPositionInput interface {
	pulumi.Input

	ToDashboardPartsPositionOutput() DashboardPartsPositionOutput
	ToDashboardPartsPositionOutputWithContext(context.Context) DashboardPartsPositionOutput
}

// The dashboard's part position.
type DashboardPartsPositionArgs struct {
	// The dashboard's part column span.
	ColSpan pulumi.IntInput `pulumi:"colSpan"`
	// The dashboard part's metadata.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The dashboard's part row span.
	RowSpan pulumi.IntInput `pulumi:"rowSpan"`
	// The dashboard's part x coordinate.
	X pulumi.IntInput `pulumi:"x"`
	// The dashboard's part y coordinate.
	Y pulumi.IntInput `pulumi:"y"`
}

func (DashboardPartsPositionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsPosition)(nil)).Elem()
}

func (i DashboardPartsPositionArgs) ToDashboardPartsPositionOutput() DashboardPartsPositionOutput {
	return i.ToDashboardPartsPositionOutputWithContext(context.Background())
}

func (i DashboardPartsPositionArgs) ToDashboardPartsPositionOutputWithContext(ctx context.Context) DashboardPartsPositionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardPartsPositionOutput)
}

// The dashboard's part position.
type DashboardPartsPositionOutput struct{ *pulumi.OutputState }

func (DashboardPartsPositionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsPosition)(nil)).Elem()
}

func (o DashboardPartsPositionOutput) ToDashboardPartsPositionOutput() DashboardPartsPositionOutput {
	return o
}

func (o DashboardPartsPositionOutput) ToDashboardPartsPositionOutputWithContext(ctx context.Context) DashboardPartsPositionOutput {
	return o
}

// The dashboard's part column span.
func (o DashboardPartsPositionOutput) ColSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.ColSpan }).(pulumi.IntOutput)
}

// The dashboard part's metadata.
func (o DashboardPartsPositionOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardPartsPosition) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The dashboard's part row span.
func (o DashboardPartsPositionOutput) RowSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.RowSpan }).(pulumi.IntOutput)
}

// The dashboard's part x coordinate.
func (o DashboardPartsPositionOutput) X() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.X }).(pulumi.IntOutput)
}

// The dashboard's part y coordinate.
func (o DashboardPartsPositionOutput) Y() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsPosition) int { return v.Y }).(pulumi.IntOutput)
}

// A dashboard part.
type DashboardPartsResponse struct {
	// A dashboard part metadata.
	Metadata interface{} `pulumi:"metadata"`
	// The dashboard's part position.
	Position DashboardPartsResponsePosition `pulumi:"position"`
}

// A dashboard part.
type DashboardPartsResponseOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponse)(nil)).Elem()
}

func (o DashboardPartsResponseOutput) ToDashboardPartsResponseOutput() DashboardPartsResponseOutput {
	return o
}

func (o DashboardPartsResponseOutput) ToDashboardPartsResponseOutputWithContext(ctx context.Context) DashboardPartsResponseOutput {
	return o
}

// A dashboard part metadata.
func (o DashboardPartsResponseOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v DashboardPartsResponse) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

// The dashboard's part position.
func (o DashboardPartsResponseOutput) Position() DashboardPartsResponsePositionOutput {
	return o.ApplyT(func(v DashboardPartsResponse) DashboardPartsResponsePosition { return v.Position }).(DashboardPartsResponsePositionOutput)
}

type DashboardPartsResponseMapOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DashboardPartsResponse)(nil)).Elem()
}

func (o DashboardPartsResponseMapOutput) ToDashboardPartsResponseMapOutput() DashboardPartsResponseMapOutput {
	return o
}

func (o DashboardPartsResponseMapOutput) ToDashboardPartsResponseMapOutputWithContext(ctx context.Context) DashboardPartsResponseMapOutput {
	return o
}

func (o DashboardPartsResponseMapOutput) MapIndex(k pulumi.StringInput) DashboardPartsResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DashboardPartsResponse {
		return vs[0].(map[string]DashboardPartsResponse)[vs[1].(string)]
	}).(DashboardPartsResponseOutput)
}

// The dashboard's part position.
type DashboardPartsResponsePosition struct {
	// The dashboard's part column span.
	ColSpan int `pulumi:"colSpan"`
	// The dashboard part's metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The dashboard's part row span.
	RowSpan int `pulumi:"rowSpan"`
	// The dashboard's part x coordinate.
	X int `pulumi:"x"`
	// The dashboard's part y coordinate.
	Y int `pulumi:"y"`
}

// The dashboard's part position.
type DashboardPartsResponsePositionOutput struct{ *pulumi.OutputState }

func (DashboardPartsResponsePositionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DashboardPartsResponsePosition)(nil)).Elem()
}

func (o DashboardPartsResponsePositionOutput) ToDashboardPartsResponsePositionOutput() DashboardPartsResponsePositionOutput {
	return o
}

func (o DashboardPartsResponsePositionOutput) ToDashboardPartsResponsePositionOutputWithContext(ctx context.Context) DashboardPartsResponsePositionOutput {
	return o
}

// The dashboard's part column span.
func (o DashboardPartsResponsePositionOutput) ColSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.ColSpan }).(pulumi.IntOutput)
}

// The dashboard part's metadata.
func (o DashboardPartsResponsePositionOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The dashboard's part row span.
func (o DashboardPartsResponsePositionOutput) RowSpan() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.RowSpan }).(pulumi.IntOutput)
}

// The dashboard's part x coordinate.
func (o DashboardPartsResponsePositionOutput) X() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.X }).(pulumi.IntOutput)
}

// The dashboard's part y coordinate.
func (o DashboardPartsResponsePositionOutput) Y() pulumi.IntOutput {
	return o.ApplyT(func(v DashboardPartsResponsePosition) int { return v.Y }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardLensOutput{})
	pulumi.RegisterOutputType(DashboardLensMapOutput{})
	pulumi.RegisterOutputType(DashboardLensResponseOutput{})
	pulumi.RegisterOutputType(DashboardLensResponseMapOutput{})
	pulumi.RegisterOutputType(DashboardPartsOutput{})
	pulumi.RegisterOutputType(DashboardPartsMapOutput{})
	pulumi.RegisterOutputType(DashboardPartsPositionOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponseOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponseMapOutput{})
	pulumi.RegisterOutputType(DashboardPartsResponsePositionOutput{})
}