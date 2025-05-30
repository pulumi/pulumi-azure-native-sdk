// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcehealth

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// Impacted resource for an event.
type EventImpactedResourceResponse struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Additional information.
	Info []KeyValueItemResponse `pulumi:"info"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Impacted resource region name.
	TargetRegion string `pulumi:"targetRegion"`
	// Identity for resource within Microsoft cloud.
	TargetResourceId string `pulumi:"targetResourceId"`
	// Resource type within Microsoft cloud.
	TargetResourceType string `pulumi:"targetResourceType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Impacted resource for an event.
type EventImpactedResourceResponseOutput struct{ *pulumi.OutputState }

func (EventImpactedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventImpactedResourceResponse)(nil)).Elem()
}

func (o EventImpactedResourceResponseOutput) ToEventImpactedResourceResponseOutput() EventImpactedResourceResponseOutput {
	return o
}

func (o EventImpactedResourceResponseOutput) ToEventImpactedResourceResponseOutputWithContext(ctx context.Context) EventImpactedResourceResponseOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o EventImpactedResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

// Additional information.
func (o EventImpactedResourceResponseOutput) Info() KeyValueItemResponseArrayOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) []KeyValueItemResponse { return v.Info }).(KeyValueItemResponseArrayOutput)
}

// The name of the resource
func (o EventImpactedResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EventImpactedResourceResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Impacted resource region name.
func (o EventImpactedResourceResponseOutput) TargetRegion() pulumi.StringOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) string { return v.TargetRegion }).(pulumi.StringOutput)
}

// Identity for resource within Microsoft cloud.
func (o EventImpactedResourceResponseOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) string { return v.TargetResourceId }).(pulumi.StringOutput)
}

// Resource type within Microsoft cloud.
func (o EventImpactedResourceResponseOutput) TargetResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) string { return v.TargetResourceType }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EventImpactedResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventImpactedResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EventImpactedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (EventImpactedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventImpactedResourceResponse)(nil)).Elem()
}

func (o EventImpactedResourceResponseArrayOutput) ToEventImpactedResourceResponseArrayOutput() EventImpactedResourceResponseArrayOutput {
	return o
}

func (o EventImpactedResourceResponseArrayOutput) ToEventImpactedResourceResponseArrayOutputWithContext(ctx context.Context) EventImpactedResourceResponseArrayOutput {
	return o
}

func (o EventImpactedResourceResponseArrayOutput) Index(i pulumi.IntInput) EventImpactedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventImpactedResourceResponse {
		return vs[0].([]EventImpactedResourceResponse)[vs[1].(int)]
	}).(EventImpactedResourceResponseOutput)
}

// Key value tuple.
type KeyValueItemResponse struct {
	// Key of tuple.
	Key string `pulumi:"key"`
	// Value of tuple.
	Value string `pulumi:"value"`
}

// Key value tuple.
type KeyValueItemResponseOutput struct{ *pulumi.OutputState }

func (KeyValueItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyValueItemResponse)(nil)).Elem()
}

func (o KeyValueItemResponseOutput) ToKeyValueItemResponseOutput() KeyValueItemResponseOutput {
	return o
}

func (o KeyValueItemResponseOutput) ToKeyValueItemResponseOutputWithContext(ctx context.Context) KeyValueItemResponseOutput {
	return o
}

// Key of tuple.
func (o KeyValueItemResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValueItemResponse) string { return v.Key }).(pulumi.StringOutput)
}

// Value of tuple.
func (o KeyValueItemResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyValueItemResponse) string { return v.Value }).(pulumi.StringOutput)
}

type KeyValueItemResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyValueItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyValueItemResponse)(nil)).Elem()
}

func (o KeyValueItemResponseArrayOutput) ToKeyValueItemResponseArrayOutput() KeyValueItemResponseArrayOutput {
	return o
}

func (o KeyValueItemResponseArrayOutput) ToKeyValueItemResponseArrayOutputWithContext(ctx context.Context) KeyValueItemResponseArrayOutput {
	return o
}

func (o KeyValueItemResponseArrayOutput) Index(i pulumi.IntInput) KeyValueItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyValueItemResponse {
		return vs[0].([]KeyValueItemResponse)[vs[1].(int)]
	}).(KeyValueItemResponseOutput)
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
	pulumi.RegisterOutputType(EventImpactedResourceResponseOutput{})
	pulumi.RegisterOutputType(EventImpactedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyValueItemResponseOutput{})
	pulumi.RegisterOutputType(KeyValueItemResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
