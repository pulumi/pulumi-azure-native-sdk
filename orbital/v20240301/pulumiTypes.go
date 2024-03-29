// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// A reference to global communications site.
type EdgeSitesPropertiesGlobalCommunicationsSite struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// EdgeSitesPropertiesGlobalCommunicationsSiteInput is an input type that accepts EdgeSitesPropertiesGlobalCommunicationsSiteArgs and EdgeSitesPropertiesGlobalCommunicationsSiteOutput values.
// You can construct a concrete instance of `EdgeSitesPropertiesGlobalCommunicationsSiteInput` via:
//
//	EdgeSitesPropertiesGlobalCommunicationsSiteArgs{...}
type EdgeSitesPropertiesGlobalCommunicationsSiteInput interface {
	pulumi.Input

	ToEdgeSitesPropertiesGlobalCommunicationsSiteOutput() EdgeSitesPropertiesGlobalCommunicationsSiteOutput
	ToEdgeSitesPropertiesGlobalCommunicationsSiteOutputWithContext(context.Context) EdgeSitesPropertiesGlobalCommunicationsSiteOutput
}

// A reference to global communications site.
type EdgeSitesPropertiesGlobalCommunicationsSiteArgs struct {
	// Resource ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (EdgeSitesPropertiesGlobalCommunicationsSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeSitesPropertiesGlobalCommunicationsSite)(nil)).Elem()
}

func (i EdgeSitesPropertiesGlobalCommunicationsSiteArgs) ToEdgeSitesPropertiesGlobalCommunicationsSiteOutput() EdgeSitesPropertiesGlobalCommunicationsSiteOutput {
	return i.ToEdgeSitesPropertiesGlobalCommunicationsSiteOutputWithContext(context.Background())
}

func (i EdgeSitesPropertiesGlobalCommunicationsSiteArgs) ToEdgeSitesPropertiesGlobalCommunicationsSiteOutputWithContext(ctx context.Context) EdgeSitesPropertiesGlobalCommunicationsSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgeSitesPropertiesGlobalCommunicationsSiteOutput)
}

// A reference to global communications site.
type EdgeSitesPropertiesGlobalCommunicationsSiteOutput struct{ *pulumi.OutputState }

func (EdgeSitesPropertiesGlobalCommunicationsSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeSitesPropertiesGlobalCommunicationsSite)(nil)).Elem()
}

func (o EdgeSitesPropertiesGlobalCommunicationsSiteOutput) ToEdgeSitesPropertiesGlobalCommunicationsSiteOutput() EdgeSitesPropertiesGlobalCommunicationsSiteOutput {
	return o
}

func (o EdgeSitesPropertiesGlobalCommunicationsSiteOutput) ToEdgeSitesPropertiesGlobalCommunicationsSiteOutputWithContext(ctx context.Context) EdgeSitesPropertiesGlobalCommunicationsSiteOutput {
	return o
}

// Resource ID.
func (o EdgeSitesPropertiesGlobalCommunicationsSiteOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EdgeSitesPropertiesGlobalCommunicationsSite) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to global communications site.
type EdgeSitesPropertiesResponseGlobalCommunicationsSite struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// A reference to global communications site.
type EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput struct{ *pulumi.OutputState }

func (EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EdgeSitesPropertiesResponseGlobalCommunicationsSite)(nil)).Elem()
}

func (o EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput) ToEdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput() EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput {
	return o
}

func (o EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput) ToEdgeSitesPropertiesResponseGlobalCommunicationsSiteOutputWithContext(ctx context.Context) EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput {
	return o
}

// Resource ID.
func (o EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EdgeSitesPropertiesResponseGlobalCommunicationsSite) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to global communications site.
type GroundStationsPropertiesGlobalCommunicationsSite struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// GroundStationsPropertiesGlobalCommunicationsSiteInput is an input type that accepts GroundStationsPropertiesGlobalCommunicationsSiteArgs and GroundStationsPropertiesGlobalCommunicationsSiteOutput values.
// You can construct a concrete instance of `GroundStationsPropertiesGlobalCommunicationsSiteInput` via:
//
//	GroundStationsPropertiesGlobalCommunicationsSiteArgs{...}
type GroundStationsPropertiesGlobalCommunicationsSiteInput interface {
	pulumi.Input

	ToGroundStationsPropertiesGlobalCommunicationsSiteOutput() GroundStationsPropertiesGlobalCommunicationsSiteOutput
	ToGroundStationsPropertiesGlobalCommunicationsSiteOutputWithContext(context.Context) GroundStationsPropertiesGlobalCommunicationsSiteOutput
}

// A reference to global communications site.
type GroundStationsPropertiesGlobalCommunicationsSiteArgs struct {
	// Resource ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (GroundStationsPropertiesGlobalCommunicationsSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroundStationsPropertiesGlobalCommunicationsSite)(nil)).Elem()
}

func (i GroundStationsPropertiesGlobalCommunicationsSiteArgs) ToGroundStationsPropertiesGlobalCommunicationsSiteOutput() GroundStationsPropertiesGlobalCommunicationsSiteOutput {
	return i.ToGroundStationsPropertiesGlobalCommunicationsSiteOutputWithContext(context.Background())
}

func (i GroundStationsPropertiesGlobalCommunicationsSiteArgs) ToGroundStationsPropertiesGlobalCommunicationsSiteOutputWithContext(ctx context.Context) GroundStationsPropertiesGlobalCommunicationsSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroundStationsPropertiesGlobalCommunicationsSiteOutput)
}

// A reference to global communications site.
type GroundStationsPropertiesGlobalCommunicationsSiteOutput struct{ *pulumi.OutputState }

func (GroundStationsPropertiesGlobalCommunicationsSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroundStationsPropertiesGlobalCommunicationsSite)(nil)).Elem()
}

func (o GroundStationsPropertiesGlobalCommunicationsSiteOutput) ToGroundStationsPropertiesGlobalCommunicationsSiteOutput() GroundStationsPropertiesGlobalCommunicationsSiteOutput {
	return o
}

func (o GroundStationsPropertiesGlobalCommunicationsSiteOutput) ToGroundStationsPropertiesGlobalCommunicationsSiteOutputWithContext(ctx context.Context) GroundStationsPropertiesGlobalCommunicationsSiteOutput {
	return o
}

// Resource ID.
func (o GroundStationsPropertiesGlobalCommunicationsSiteOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GroundStationsPropertiesGlobalCommunicationsSite) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to global communications site.
type GroundStationsPropertiesResponseGlobalCommunicationsSite struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// A reference to global communications site.
type GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput struct{ *pulumi.OutputState }

func (GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroundStationsPropertiesResponseGlobalCommunicationsSite)(nil)).Elem()
}

func (o GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput) ToGroundStationsPropertiesResponseGlobalCommunicationsSiteOutput() GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput {
	return o
}

func (o GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput) ToGroundStationsPropertiesResponseGlobalCommunicationsSiteOutputWithContext(ctx context.Context) GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput {
	return o
}

// Resource ID.
func (o GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GroundStationsPropertiesResponseGlobalCommunicationsSite) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
type L2ConnectionsPropertiesEdgeSite struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// L2ConnectionsPropertiesEdgeSiteInput is an input type that accepts L2ConnectionsPropertiesEdgeSiteArgs and L2ConnectionsPropertiesEdgeSiteOutput values.
// You can construct a concrete instance of `L2ConnectionsPropertiesEdgeSiteInput` via:
//
//	L2ConnectionsPropertiesEdgeSiteArgs{...}
type L2ConnectionsPropertiesEdgeSiteInput interface {
	pulumi.Input

	ToL2ConnectionsPropertiesEdgeSiteOutput() L2ConnectionsPropertiesEdgeSiteOutput
	ToL2ConnectionsPropertiesEdgeSiteOutputWithContext(context.Context) L2ConnectionsPropertiesEdgeSiteOutput
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
type L2ConnectionsPropertiesEdgeSiteArgs struct {
	// Resource ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (L2ConnectionsPropertiesEdgeSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*L2ConnectionsPropertiesEdgeSite)(nil)).Elem()
}

func (i L2ConnectionsPropertiesEdgeSiteArgs) ToL2ConnectionsPropertiesEdgeSiteOutput() L2ConnectionsPropertiesEdgeSiteOutput {
	return i.ToL2ConnectionsPropertiesEdgeSiteOutputWithContext(context.Background())
}

func (i L2ConnectionsPropertiesEdgeSiteArgs) ToL2ConnectionsPropertiesEdgeSiteOutputWithContext(ctx context.Context) L2ConnectionsPropertiesEdgeSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2ConnectionsPropertiesEdgeSiteOutput)
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
type L2ConnectionsPropertiesEdgeSiteOutput struct{ *pulumi.OutputState }

func (L2ConnectionsPropertiesEdgeSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*L2ConnectionsPropertiesEdgeSite)(nil)).Elem()
}

func (o L2ConnectionsPropertiesEdgeSiteOutput) ToL2ConnectionsPropertiesEdgeSiteOutput() L2ConnectionsPropertiesEdgeSiteOutput {
	return o
}

func (o L2ConnectionsPropertiesEdgeSiteOutput) ToL2ConnectionsPropertiesEdgeSiteOutputWithContext(ctx context.Context) L2ConnectionsPropertiesEdgeSiteOutput {
	return o
}

// Resource ID.
func (o L2ConnectionsPropertiesEdgeSiteOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v L2ConnectionsPropertiesEdgeSite) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
type L2ConnectionsPropertiesGroundStation struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// L2ConnectionsPropertiesGroundStationInput is an input type that accepts L2ConnectionsPropertiesGroundStationArgs and L2ConnectionsPropertiesGroundStationOutput values.
// You can construct a concrete instance of `L2ConnectionsPropertiesGroundStationInput` via:
//
//	L2ConnectionsPropertiesGroundStationArgs{...}
type L2ConnectionsPropertiesGroundStationInput interface {
	pulumi.Input

	ToL2ConnectionsPropertiesGroundStationOutput() L2ConnectionsPropertiesGroundStationOutput
	ToL2ConnectionsPropertiesGroundStationOutputWithContext(context.Context) L2ConnectionsPropertiesGroundStationOutput
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
type L2ConnectionsPropertiesGroundStationArgs struct {
	// Resource ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (L2ConnectionsPropertiesGroundStationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*L2ConnectionsPropertiesGroundStation)(nil)).Elem()
}

func (i L2ConnectionsPropertiesGroundStationArgs) ToL2ConnectionsPropertiesGroundStationOutput() L2ConnectionsPropertiesGroundStationOutput {
	return i.ToL2ConnectionsPropertiesGroundStationOutputWithContext(context.Background())
}

func (i L2ConnectionsPropertiesGroundStationArgs) ToL2ConnectionsPropertiesGroundStationOutputWithContext(ctx context.Context) L2ConnectionsPropertiesGroundStationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L2ConnectionsPropertiesGroundStationOutput)
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
type L2ConnectionsPropertiesGroundStationOutput struct{ *pulumi.OutputState }

func (L2ConnectionsPropertiesGroundStationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*L2ConnectionsPropertiesGroundStation)(nil)).Elem()
}

func (o L2ConnectionsPropertiesGroundStationOutput) ToL2ConnectionsPropertiesGroundStationOutput() L2ConnectionsPropertiesGroundStationOutput {
	return o
}

func (o L2ConnectionsPropertiesGroundStationOutput) ToL2ConnectionsPropertiesGroundStationOutputWithContext(ctx context.Context) L2ConnectionsPropertiesGroundStationOutput {
	return o
}

// Resource ID.
func (o L2ConnectionsPropertiesGroundStationOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v L2ConnectionsPropertiesGroundStation) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
type L2ConnectionsPropertiesResponseEdgeSite struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// A reference to an Microsoft.Orbital/edgeSites resource to route traffic for.
type L2ConnectionsPropertiesResponseEdgeSiteOutput struct{ *pulumi.OutputState }

func (L2ConnectionsPropertiesResponseEdgeSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*L2ConnectionsPropertiesResponseEdgeSite)(nil)).Elem()
}

func (o L2ConnectionsPropertiesResponseEdgeSiteOutput) ToL2ConnectionsPropertiesResponseEdgeSiteOutput() L2ConnectionsPropertiesResponseEdgeSiteOutput {
	return o
}

func (o L2ConnectionsPropertiesResponseEdgeSiteOutput) ToL2ConnectionsPropertiesResponseEdgeSiteOutputWithContext(ctx context.Context) L2ConnectionsPropertiesResponseEdgeSiteOutput {
	return o
}

// Resource ID.
func (o L2ConnectionsPropertiesResponseEdgeSiteOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v L2ConnectionsPropertiesResponseEdgeSite) string { return v.Id }).(pulumi.StringOutput)
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
type L2ConnectionsPropertiesResponseGroundStation struct {
	// Resource ID.
	Id string `pulumi:"id"`
}

// A reference to an Microsoft.Orbital/groundStations resource to route traffic for.
type L2ConnectionsPropertiesResponseGroundStationOutput struct{ *pulumi.OutputState }

func (L2ConnectionsPropertiesResponseGroundStationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*L2ConnectionsPropertiesResponseGroundStation)(nil)).Elem()
}

func (o L2ConnectionsPropertiesResponseGroundStationOutput) ToL2ConnectionsPropertiesResponseGroundStationOutput() L2ConnectionsPropertiesResponseGroundStationOutput {
	return o
}

func (o L2ConnectionsPropertiesResponseGroundStationOutput) ToL2ConnectionsPropertiesResponseGroundStationOutputWithContext(ctx context.Context) L2ConnectionsPropertiesResponseGroundStationOutput {
	return o
}

// Resource ID.
func (o L2ConnectionsPropertiesResponseGroundStationOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v L2ConnectionsPropertiesResponseGroundStation) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIdListResultResponseValue struct {
	// The Azure Resource ID.
	Id *string `pulumi:"id"`
}

type ResourceIdListResultResponseValueOutput struct{ *pulumi.OutputState }

func (ResourceIdListResultResponseValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdListResultResponseValue)(nil)).Elem()
}

func (o ResourceIdListResultResponseValueOutput) ToResourceIdListResultResponseValueOutput() ResourceIdListResultResponseValueOutput {
	return o
}

func (o ResourceIdListResultResponseValueOutput) ToResourceIdListResultResponseValueOutputWithContext(ctx context.Context) ResourceIdListResultResponseValueOutput {
	return o
}

// The Azure Resource ID.
func (o ResourceIdListResultResponseValueOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdListResultResponseValue) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceIdListResultResponseValueArrayOutput struct{ *pulumi.OutputState }

func (ResourceIdListResultResponseValueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceIdListResultResponseValue)(nil)).Elem()
}

func (o ResourceIdListResultResponseValueArrayOutput) ToResourceIdListResultResponseValueArrayOutput() ResourceIdListResultResponseValueArrayOutput {
	return o
}

func (o ResourceIdListResultResponseValueArrayOutput) ToResourceIdListResultResponseValueArrayOutputWithContext(ctx context.Context) ResourceIdListResultResponseValueArrayOutput {
	return o
}

func (o ResourceIdListResultResponseValueArrayOutput) Index(i pulumi.IntInput) ResourceIdListResultResponseValueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceIdListResultResponseValue {
		return vs[0].([]ResourceIdListResultResponseValue)[vs[1].(int)]
	}).(ResourceIdListResultResponseValueOutput)
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
	pulumi.RegisterOutputType(EdgeSitesPropertiesGlobalCommunicationsSiteOutput{})
	pulumi.RegisterOutputType(EdgeSitesPropertiesResponseGlobalCommunicationsSiteOutput{})
	pulumi.RegisterOutputType(GroundStationsPropertiesGlobalCommunicationsSiteOutput{})
	pulumi.RegisterOutputType(GroundStationsPropertiesResponseGlobalCommunicationsSiteOutput{})
	pulumi.RegisterOutputType(L2ConnectionsPropertiesEdgeSiteOutput{})
	pulumi.RegisterOutputType(L2ConnectionsPropertiesGroundStationOutput{})
	pulumi.RegisterOutputType(L2ConnectionsPropertiesResponseEdgeSiteOutput{})
	pulumi.RegisterOutputType(L2ConnectionsPropertiesResponseGroundStationOutput{})
	pulumi.RegisterOutputType(ResourceIdListResultResponseValueOutput{})
	pulumi.RegisterOutputType(ResourceIdListResultResponseValueArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
