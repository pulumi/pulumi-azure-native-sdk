// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The private link scope resource reference.
type PrivateLinkScopedResourceResponse struct {
	// The full resource Id of the private link scope resource.
	ResourceId *string `pulumi:"resourceId"`
	// The private link scope unique Identifier.
	ScopeId *string `pulumi:"scopeId"`
}

// The private link scope resource reference.
type PrivateLinkScopedResourceResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutput() PrivateLinkScopedResourceResponseOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseOutput) ToPrivateLinkScopedResourceResponseOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseOutput {
	return o
}

// The full resource Id of the private link scope resource.
func (o PrivateLinkScopedResourceResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// The private link scope unique Identifier.
func (o PrivateLinkScopedResourceResponseOutput) ScopeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkScopedResourceResponse) *string { return v.ScopeId }).(pulumi.StringPtrOutput)
}

type PrivateLinkScopedResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateLinkScopedResourceResponse)(nil)).Elem()
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutput() PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) ToPrivateLinkScopedResourceResponseArrayOutputWithContext(ctx context.Context) PrivateLinkScopedResourceResponseArrayOutput {
	return o
}

func (o PrivateLinkScopedResourceResponseArrayOutput) Index(i pulumi.IntInput) PrivateLinkScopedResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateLinkScopedResourceResponse {
		return vs[0].([]PrivateLinkScopedResourceResponse)[vs[1].(int)]
	}).(PrivateLinkScopedResourceResponseOutput)
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

// The daily volume cap for ingestion.
type WorkspaceCapping struct {
	// The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `pulumi:"dailyQuotaGb"`
}

// WorkspaceCappingInput is an input type that accepts WorkspaceCappingArgs and WorkspaceCappingOutput values.
// You can construct a concrete instance of `WorkspaceCappingInput` via:
//
//	WorkspaceCappingArgs{...}
type WorkspaceCappingInput interface {
	pulumi.Input

	ToWorkspaceCappingOutput() WorkspaceCappingOutput
	ToWorkspaceCappingOutputWithContext(context.Context) WorkspaceCappingOutput
}

// The daily volume cap for ingestion.
type WorkspaceCappingArgs struct {
	// The workspace daily quota for ingestion.
	DailyQuotaGb pulumi.Float64PtrInput `pulumi:"dailyQuotaGb"`
}

func (WorkspaceCappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCapping)(nil)).Elem()
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingOutput() WorkspaceCappingOutput {
	return i.ToWorkspaceCappingOutputWithContext(context.Background())
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingOutputWithContext(ctx context.Context) WorkspaceCappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingOutput)
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return i.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (i WorkspaceCappingArgs) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingOutput).ToWorkspaceCappingPtrOutputWithContext(ctx)
}

// WorkspaceCappingPtrInput is an input type that accepts WorkspaceCappingArgs, WorkspaceCappingPtr and WorkspaceCappingPtrOutput values.
// You can construct a concrete instance of `WorkspaceCappingPtrInput` via:
//
//	        WorkspaceCappingArgs{...}
//
//	or:
//
//	        nil
type WorkspaceCappingPtrInput interface {
	pulumi.Input

	ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput
	ToWorkspaceCappingPtrOutputWithContext(context.Context) WorkspaceCappingPtrOutput
}

type workspaceCappingPtrType WorkspaceCappingArgs

func WorkspaceCappingPtr(v *WorkspaceCappingArgs) WorkspaceCappingPtrInput {
	return (*workspaceCappingPtrType)(v)
}

func (*workspaceCappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCapping)(nil)).Elem()
}

func (i *workspaceCappingPtrType) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return i.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (i *workspaceCappingPtrType) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCappingPtrOutput)
}

// The daily volume cap for ingestion.
type WorkspaceCappingOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCapping)(nil)).Elem()
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingOutput() WorkspaceCappingOutput {
	return o
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingOutputWithContext(ctx context.Context) WorkspaceCappingOutput {
	return o
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return o.ToWorkspaceCappingPtrOutputWithContext(context.Background())
}

func (o WorkspaceCappingOutput) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceCapping) *WorkspaceCapping {
		return &v
	}).(WorkspaceCappingPtrOutput)
}

// The workspace daily quota for ingestion.
func (o WorkspaceCappingOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkspaceCapping) *float64 { return v.DailyQuotaGb }).(pulumi.Float64PtrOutput)
}

type WorkspaceCappingPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCapping)(nil)).Elem()
}

func (o WorkspaceCappingPtrOutput) ToWorkspaceCappingPtrOutput() WorkspaceCappingPtrOutput {
	return o
}

func (o WorkspaceCappingPtrOutput) ToWorkspaceCappingPtrOutputWithContext(ctx context.Context) WorkspaceCappingPtrOutput {
	return o
}

func (o WorkspaceCappingPtrOutput) Elem() WorkspaceCappingOutput {
	return o.ApplyT(func(v *WorkspaceCapping) WorkspaceCapping {
		if v != nil {
			return *v
		}
		var ret WorkspaceCapping
		return ret
	}).(WorkspaceCappingOutput)
}

// The workspace daily quota for ingestion.
func (o WorkspaceCappingPtrOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkspaceCapping) *float64 {
		if v == nil {
			return nil
		}
		return v.DailyQuotaGb
	}).(pulumi.Float64PtrOutput)
}

// The daily volume cap for ingestion.
type WorkspaceCappingResponse struct {
	// The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `pulumi:"dailyQuotaGb"`
	// The status of data ingestion for this workspace.
	DataIngestionStatus string `pulumi:"dataIngestionStatus"`
	// The time when the quota will be rest.
	QuotaNextResetTime string `pulumi:"quotaNextResetTime"`
}

// The daily volume cap for ingestion.
type WorkspaceCappingResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCappingResponse)(nil)).Elem()
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponseOutput() WorkspaceCappingResponseOutput {
	return o
}

func (o WorkspaceCappingResponseOutput) ToWorkspaceCappingResponseOutputWithContext(ctx context.Context) WorkspaceCappingResponseOutput {
	return o
}

// The workspace daily quota for ingestion.
func (o WorkspaceCappingResponseOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) *float64 { return v.DailyQuotaGb }).(pulumi.Float64PtrOutput)
}

// The status of data ingestion for this workspace.
func (o WorkspaceCappingResponseOutput) DataIngestionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) string { return v.DataIngestionStatus }).(pulumi.StringOutput)
}

// The time when the quota will be rest.
func (o WorkspaceCappingResponseOutput) QuotaNextResetTime() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceCappingResponse) string { return v.QuotaNextResetTime }).(pulumi.StringOutput)
}

type WorkspaceCappingResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceCappingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceCappingResponse)(nil)).Elem()
}

func (o WorkspaceCappingResponsePtrOutput) ToWorkspaceCappingResponsePtrOutput() WorkspaceCappingResponsePtrOutput {
	return o
}

func (o WorkspaceCappingResponsePtrOutput) ToWorkspaceCappingResponsePtrOutputWithContext(ctx context.Context) WorkspaceCappingResponsePtrOutput {
	return o
}

func (o WorkspaceCappingResponsePtrOutput) Elem() WorkspaceCappingResponseOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) WorkspaceCappingResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceCappingResponse
		return ret
	}).(WorkspaceCappingResponseOutput)
}

// The workspace daily quota for ingestion.
func (o WorkspaceCappingResponsePtrOutput) DailyQuotaGb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DailyQuotaGb
	}).(pulumi.Float64PtrOutput)
}

// The status of data ingestion for this workspace.
func (o WorkspaceCappingResponsePtrOutput) DataIngestionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataIngestionStatus
	}).(pulumi.StringPtrOutput)
}

// The time when the quota will be rest.
func (o WorkspaceCappingResponsePtrOutput) QuotaNextResetTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceCappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QuotaNextResetTime
	}).(pulumi.StringPtrOutput)
}

// Workspace features.
type WorkspaceFeatures struct {
	// Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceId *string `pulumi:"clusterResourceId"`
	// Disable Non-AAD based Auth.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Flag that indicate if data should be exported.
	EnableDataExport *bool `pulumi:"enableDataExport"`
	// Flag that indicate which permission to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions *bool `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	// Flag that describes if we want to remove the data after 30 days.
	ImmediatePurgeDataOn30Days *bool `pulumi:"immediatePurgeDataOn30Days"`
}

// WorkspaceFeaturesInput is an input type that accepts WorkspaceFeaturesArgs and WorkspaceFeaturesOutput values.
// You can construct a concrete instance of `WorkspaceFeaturesInput` via:
//
//	WorkspaceFeaturesArgs{...}
type WorkspaceFeaturesInput interface {
	pulumi.Input

	ToWorkspaceFeaturesOutput() WorkspaceFeaturesOutput
	ToWorkspaceFeaturesOutputWithContext(context.Context) WorkspaceFeaturesOutput
}

// Workspace features.
type WorkspaceFeaturesArgs struct {
	// Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceId pulumi.StringPtrInput `pulumi:"clusterResourceId"`
	// Disable Non-AAD based Auth.
	DisableLocalAuth pulumi.BoolPtrInput `pulumi:"disableLocalAuth"`
	// Flag that indicate if data should be exported.
	EnableDataExport pulumi.BoolPtrInput `pulumi:"enableDataExport"`
	// Flag that indicate which permission to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions pulumi.BoolPtrInput `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	// Flag that describes if we want to remove the data after 30 days.
	ImmediatePurgeDataOn30Days pulumi.BoolPtrInput `pulumi:"immediatePurgeDataOn30Days"`
}

func (WorkspaceFeaturesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeatures)(nil)).Elem()
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesOutput() WorkspaceFeaturesOutput {
	return i.ToWorkspaceFeaturesOutputWithContext(context.Background())
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesOutputWithContext(ctx context.Context) WorkspaceFeaturesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesOutput)
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return i.ToWorkspaceFeaturesPtrOutputWithContext(context.Background())
}

func (i WorkspaceFeaturesArgs) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesOutput).ToWorkspaceFeaturesPtrOutputWithContext(ctx)
}

// WorkspaceFeaturesPtrInput is an input type that accepts WorkspaceFeaturesArgs, WorkspaceFeaturesPtr and WorkspaceFeaturesPtrOutput values.
// You can construct a concrete instance of `WorkspaceFeaturesPtrInput` via:
//
//	        WorkspaceFeaturesArgs{...}
//
//	or:
//
//	        nil
type WorkspaceFeaturesPtrInput interface {
	pulumi.Input

	ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput
	ToWorkspaceFeaturesPtrOutputWithContext(context.Context) WorkspaceFeaturesPtrOutput
}

type workspaceFeaturesPtrType WorkspaceFeaturesArgs

func WorkspaceFeaturesPtr(v *WorkspaceFeaturesArgs) WorkspaceFeaturesPtrInput {
	return (*workspaceFeaturesPtrType)(v)
}

func (*workspaceFeaturesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeatures)(nil)).Elem()
}

func (i *workspaceFeaturesPtrType) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return i.ToWorkspaceFeaturesPtrOutputWithContext(context.Background())
}

func (i *workspaceFeaturesPtrType) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceFeaturesPtrOutput)
}

// Workspace features.
type WorkspaceFeaturesOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeatures)(nil)).Elem()
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesOutput() WorkspaceFeaturesOutput {
	return o
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesOutputWithContext(ctx context.Context) WorkspaceFeaturesOutput {
	return o
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return o.ToWorkspaceFeaturesPtrOutputWithContext(context.Background())
}

func (o WorkspaceFeaturesOutput) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceFeatures) *WorkspaceFeatures {
		return &v
	}).(WorkspaceFeaturesPtrOutput)
}

// Dedicated LA cluster resourceId that is linked to the workspaces.
func (o WorkspaceFeaturesOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *string { return v.ClusterResourceId }).(pulumi.StringPtrOutput)
}

// Disable Non-AAD based Auth.
func (o WorkspaceFeaturesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Flag that indicate if data should be exported.
func (o WorkspaceFeaturesOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.EnableDataExport }).(pulumi.BoolPtrOutput)
}

// Flag that indicate which permission to use - resource or workspace or both.
func (o WorkspaceFeaturesOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.EnableLogAccessUsingOnlyResourcePermissions }).(pulumi.BoolPtrOutput)
}

// Flag that describes if we want to remove the data after 30 days.
func (o WorkspaceFeaturesOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeatures) *bool { return v.ImmediatePurgeDataOn30Days }).(pulumi.BoolPtrOutput)
}

type WorkspaceFeaturesPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeatures)(nil)).Elem()
}

func (o WorkspaceFeaturesPtrOutput) ToWorkspaceFeaturesPtrOutput() WorkspaceFeaturesPtrOutput {
	return o
}

func (o WorkspaceFeaturesPtrOutput) ToWorkspaceFeaturesPtrOutputWithContext(ctx context.Context) WorkspaceFeaturesPtrOutput {
	return o
}

func (o WorkspaceFeaturesPtrOutput) Elem() WorkspaceFeaturesOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) WorkspaceFeatures {
		if v != nil {
			return *v
		}
		var ret WorkspaceFeatures
		return ret
	}).(WorkspaceFeaturesOutput)
}

// Dedicated LA cluster resourceId that is linked to the workspaces.
func (o WorkspaceFeaturesPtrOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *string {
		if v == nil {
			return nil
		}
		return v.ClusterResourceId
	}).(pulumi.StringPtrOutput)
}

// Disable Non-AAD based Auth.
func (o WorkspaceFeaturesPtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

// Flag that indicate if data should be exported.
func (o WorkspaceFeaturesPtrOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDataExport
	}).(pulumi.BoolPtrOutput)
}

// Flag that indicate which permission to use - resource or workspace or both.
func (o WorkspaceFeaturesPtrOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.EnableLogAccessUsingOnlyResourcePermissions
	}).(pulumi.BoolPtrOutput)
}

// Flag that describes if we want to remove the data after 30 days.
func (o WorkspaceFeaturesPtrOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeatures) *bool {
		if v == nil {
			return nil
		}
		return v.ImmediatePurgeDataOn30Days
	}).(pulumi.BoolPtrOutput)
}

// Workspace features.
type WorkspaceFeaturesResponse struct {
	// Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceId *string `pulumi:"clusterResourceId"`
	// Disable Non-AAD based Auth.
	DisableLocalAuth *bool `pulumi:"disableLocalAuth"`
	// Flag that indicate if data should be exported.
	EnableDataExport *bool `pulumi:"enableDataExport"`
	// Flag that indicate which permission to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions *bool `pulumi:"enableLogAccessUsingOnlyResourcePermissions"`
	// Flag that describes if we want to remove the data after 30 days.
	ImmediatePurgeDataOn30Days *bool `pulumi:"immediatePurgeDataOn30Days"`
}

// Workspace features.
type WorkspaceFeaturesResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceFeaturesResponse)(nil)).Elem()
}

func (o WorkspaceFeaturesResponseOutput) ToWorkspaceFeaturesResponseOutput() WorkspaceFeaturesResponseOutput {
	return o
}

func (o WorkspaceFeaturesResponseOutput) ToWorkspaceFeaturesResponseOutputWithContext(ctx context.Context) WorkspaceFeaturesResponseOutput {
	return o
}

// Dedicated LA cluster resourceId that is linked to the workspaces.
func (o WorkspaceFeaturesResponseOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *string { return v.ClusterResourceId }).(pulumi.StringPtrOutput)
}

// Disable Non-AAD based Auth.
func (o WorkspaceFeaturesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

// Flag that indicate if data should be exported.
func (o WorkspaceFeaturesResponseOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.EnableDataExport }).(pulumi.BoolPtrOutput)
}

// Flag that indicate which permission to use - resource or workspace or both.
func (o WorkspaceFeaturesResponseOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.EnableLogAccessUsingOnlyResourcePermissions }).(pulumi.BoolPtrOutput)
}

// Flag that describes if we want to remove the data after 30 days.
func (o WorkspaceFeaturesResponseOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WorkspaceFeaturesResponse) *bool { return v.ImmediatePurgeDataOn30Days }).(pulumi.BoolPtrOutput)
}

type WorkspaceFeaturesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceFeaturesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceFeaturesResponse)(nil)).Elem()
}

func (o WorkspaceFeaturesResponsePtrOutput) ToWorkspaceFeaturesResponsePtrOutput() WorkspaceFeaturesResponsePtrOutput {
	return o
}

func (o WorkspaceFeaturesResponsePtrOutput) ToWorkspaceFeaturesResponsePtrOutputWithContext(ctx context.Context) WorkspaceFeaturesResponsePtrOutput {
	return o
}

func (o WorkspaceFeaturesResponsePtrOutput) Elem() WorkspaceFeaturesResponseOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) WorkspaceFeaturesResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceFeaturesResponse
		return ret
	}).(WorkspaceFeaturesResponseOutput)
}

// Dedicated LA cluster resourceId that is linked to the workspaces.
func (o WorkspaceFeaturesResponsePtrOutput) ClusterResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterResourceId
	}).(pulumi.StringPtrOutput)
}

// Disable Non-AAD based Auth.
func (o WorkspaceFeaturesResponsePtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

// Flag that indicate if data should be exported.
func (o WorkspaceFeaturesResponsePtrOutput) EnableDataExport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableDataExport
	}).(pulumi.BoolPtrOutput)
}

// Flag that indicate which permission to use - resource or workspace or both.
func (o WorkspaceFeaturesResponsePtrOutput) EnableLogAccessUsingOnlyResourcePermissions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableLogAccessUsingOnlyResourcePermissions
	}).(pulumi.BoolPtrOutput)
}

// Flag that describes if we want to remove the data after 30 days.
func (o WorkspaceFeaturesResponsePtrOutput) ImmediatePurgeDataOn30Days() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WorkspaceFeaturesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ImmediatePurgeDataOn30Days
	}).(pulumi.BoolPtrOutput)
}

// The SKU (tier) of a workspace.
type WorkspaceSku struct {
	// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
	CapacityReservationLevel *int `pulumi:"capacityReservationLevel"`
	// The name of the SKU.
	Name string `pulumi:"name"`
}

// WorkspaceSkuInput is an input type that accepts WorkspaceSkuArgs and WorkspaceSkuOutput values.
// You can construct a concrete instance of `WorkspaceSkuInput` via:
//
//	WorkspaceSkuArgs{...}
type WorkspaceSkuInput interface {
	pulumi.Input

	ToWorkspaceSkuOutput() WorkspaceSkuOutput
	ToWorkspaceSkuOutputWithContext(context.Context) WorkspaceSkuOutput
}

// The SKU (tier) of a workspace.
type WorkspaceSkuArgs struct {
	// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
	CapacityReservationLevel pulumi.IntPtrInput `pulumi:"capacityReservationLevel"`
	// The name of the SKU.
	Name pulumi.StringInput `pulumi:"name"`
}

func (WorkspaceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSku)(nil)).Elem()
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuOutput() WorkspaceSkuOutput {
	return i.ToWorkspaceSkuOutputWithContext(context.Background())
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuOutputWithContext(ctx context.Context) WorkspaceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuOutput)
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return i.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (i WorkspaceSkuArgs) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuOutput).ToWorkspaceSkuPtrOutputWithContext(ctx)
}

// WorkspaceSkuPtrInput is an input type that accepts WorkspaceSkuArgs, WorkspaceSkuPtr and WorkspaceSkuPtrOutput values.
// You can construct a concrete instance of `WorkspaceSkuPtrInput` via:
//
//	        WorkspaceSkuArgs{...}
//
//	or:
//
//	        nil
type WorkspaceSkuPtrInput interface {
	pulumi.Input

	ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput
	ToWorkspaceSkuPtrOutputWithContext(context.Context) WorkspaceSkuPtrOutput
}

type workspaceSkuPtrType WorkspaceSkuArgs

func WorkspaceSkuPtr(v *WorkspaceSkuArgs) WorkspaceSkuPtrInput {
	return (*workspaceSkuPtrType)(v)
}

func (*workspaceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSku)(nil)).Elem()
}

func (i *workspaceSkuPtrType) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return i.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (i *workspaceSkuPtrType) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceSkuPtrOutput)
}

// The SKU (tier) of a workspace.
type WorkspaceSkuOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSku)(nil)).Elem()
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuOutput() WorkspaceSkuOutput {
	return o
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuOutputWithContext(ctx context.Context) WorkspaceSkuOutput {
	return o
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return o.ToWorkspaceSkuPtrOutputWithContext(context.Background())
}

func (o WorkspaceSkuOutput) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkspaceSku) *WorkspaceSku {
		return &v
	}).(WorkspaceSkuPtrOutput)
}

// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
func (o WorkspaceSkuOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkspaceSku) *int { return v.CapacityReservationLevel }).(pulumi.IntPtrOutput)
}

// The name of the SKU.
func (o WorkspaceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSku) string { return v.Name }).(pulumi.StringOutput)
}

type WorkspaceSkuPtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSku)(nil)).Elem()
}

func (o WorkspaceSkuPtrOutput) ToWorkspaceSkuPtrOutput() WorkspaceSkuPtrOutput {
	return o
}

func (o WorkspaceSkuPtrOutput) ToWorkspaceSkuPtrOutputWithContext(ctx context.Context) WorkspaceSkuPtrOutput {
	return o
}

func (o WorkspaceSkuPtrOutput) Elem() WorkspaceSkuOutput {
	return o.ApplyT(func(v *WorkspaceSku) WorkspaceSku {
		if v != nil {
			return *v
		}
		var ret WorkspaceSku
		return ret
	}).(WorkspaceSkuOutput)
}

// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
func (o WorkspaceSkuPtrOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSku) *int {
		if v == nil {
			return nil
		}
		return v.CapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

// The name of the SKU.
func (o WorkspaceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The SKU (tier) of a workspace.
type WorkspaceSkuResponse struct {
	// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
	CapacityReservationLevel *int `pulumi:"capacityReservationLevel"`
	// The last time when the sku was updated.
	LastSkuUpdate string `pulumi:"lastSkuUpdate"`
	// The name of the SKU.
	Name string `pulumi:"name"`
}

// The SKU (tier) of a workspace.
type WorkspaceSkuResponseOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceSkuResponse)(nil)).Elem()
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponseOutput() WorkspaceSkuResponseOutput {
	return o
}

func (o WorkspaceSkuResponseOutput) ToWorkspaceSkuResponseOutputWithContext(ctx context.Context) WorkspaceSkuResponseOutput {
	return o
}

// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
func (o WorkspaceSkuResponseOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) *int { return v.CapacityReservationLevel }).(pulumi.IntPtrOutput)
}

// The last time when the sku was updated.
func (o WorkspaceSkuResponseOutput) LastSkuUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) string { return v.LastSkuUpdate }).(pulumi.StringOutput)
}

// The name of the SKU.
func (o WorkspaceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspaceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type WorkspaceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkspaceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceSkuResponse)(nil)).Elem()
}

func (o WorkspaceSkuResponsePtrOutput) ToWorkspaceSkuResponsePtrOutput() WorkspaceSkuResponsePtrOutput {
	return o
}

func (o WorkspaceSkuResponsePtrOutput) ToWorkspaceSkuResponsePtrOutputWithContext(ctx context.Context) WorkspaceSkuResponsePtrOutput {
	return o
}

func (o WorkspaceSkuResponsePtrOutput) Elem() WorkspaceSkuResponseOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) WorkspaceSkuResponse {
		if v != nil {
			return *v
		}
		var ret WorkspaceSkuResponse
		return ret
	}).(WorkspaceSkuResponseOutput)
}

// The capacity reservation level in GB for this workspace, when CapacityReservation sku is selected.
func (o WorkspaceSkuResponsePtrOutput) CapacityReservationLevel() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.CapacityReservationLevel
	}).(pulumi.IntPtrOutput)
}

// The last time when the sku was updated.
func (o WorkspaceSkuResponsePtrOutput) LastSkuUpdate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSkuUpdate
	}).(pulumi.StringPtrOutput)
}

// The name of the SKU.
func (o WorkspaceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkScopedResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceCappingResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceFeaturesResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuPtrOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuResponseOutput{})
	pulumi.RegisterOutputType(WorkspaceSkuResponsePtrOutput{})
}
