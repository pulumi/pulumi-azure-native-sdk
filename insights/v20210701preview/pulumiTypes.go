// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties that define the scope private link mode settings.
type AccessModeSettings struct {
	// List of exclusions that override the default access mode settings for specific private endpoint connections.
	Exclusions []AccessModeSettingsExclusion `pulumi:"exclusions"`
	// Specifies the default access mode of ingestion through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
	IngestionAccessMode string `pulumi:"ingestionAccessMode"`
	// Specifies the default access mode of queries through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
	QueryAccessMode string `pulumi:"queryAccessMode"`
}

// AccessModeSettingsInput is an input type that accepts AccessModeSettingsArgs and AccessModeSettingsOutput values.
// You can construct a concrete instance of `AccessModeSettingsInput` via:
//
//	AccessModeSettingsArgs{...}
type AccessModeSettingsInput interface {
	pulumi.Input

	ToAccessModeSettingsOutput() AccessModeSettingsOutput
	ToAccessModeSettingsOutputWithContext(context.Context) AccessModeSettingsOutput
}

// Properties that define the scope private link mode settings.
type AccessModeSettingsArgs struct {
	// List of exclusions that override the default access mode settings for specific private endpoint connections.
	Exclusions AccessModeSettingsExclusionArrayInput `pulumi:"exclusions"`
	// Specifies the default access mode of ingestion through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
	IngestionAccessMode pulumi.StringInput `pulumi:"ingestionAccessMode"`
	// Specifies the default access mode of queries through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
	QueryAccessMode pulumi.StringInput `pulumi:"queryAccessMode"`
}

func (AccessModeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettings)(nil)).Elem()
}

func (i AccessModeSettingsArgs) ToAccessModeSettingsOutput() AccessModeSettingsOutput {
	return i.ToAccessModeSettingsOutputWithContext(context.Background())
}

func (i AccessModeSettingsArgs) ToAccessModeSettingsOutputWithContext(ctx context.Context) AccessModeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsOutput)
}

// Properties that define the scope private link mode settings.
type AccessModeSettingsOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettings)(nil)).Elem()
}

func (o AccessModeSettingsOutput) ToAccessModeSettingsOutput() AccessModeSettingsOutput {
	return o
}

func (o AccessModeSettingsOutput) ToAccessModeSettingsOutputWithContext(ctx context.Context) AccessModeSettingsOutput {
	return o
}

// List of exclusions that override the default access mode settings for specific private endpoint connections.
func (o AccessModeSettingsOutput) Exclusions() AccessModeSettingsExclusionArrayOutput {
	return o.ApplyT(func(v AccessModeSettings) []AccessModeSettingsExclusion { return v.Exclusions }).(AccessModeSettingsExclusionArrayOutput)
}

// Specifies the default access mode of ingestion through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
func (o AccessModeSettingsOutput) IngestionAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettings) string { return v.IngestionAccessMode }).(pulumi.StringOutput)
}

// Specifies the default access mode of queries through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
func (o AccessModeSettingsOutput) QueryAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettings) string { return v.QueryAccessMode }).(pulumi.StringOutput)
}

// Properties that define the scope private link mode settings exclusion item. This setting applies to a specific private endpoint connection and overrides the default settings for that private endpoint connection.
type AccessModeSettingsExclusion struct {
	// Specifies the access mode of ingestion through the specified private endpoint connection in the exclusion.
	IngestionAccessMode *string `pulumi:"ingestionAccessMode"`
	// The private endpoint connection name associated to the private endpoint on which we want to apply the specific access mode settings.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Specifies the access mode of queries through the specified private endpoint connection in the exclusion.
	QueryAccessMode *string `pulumi:"queryAccessMode"`
}

// AccessModeSettingsExclusionInput is an input type that accepts AccessModeSettingsExclusionArgs and AccessModeSettingsExclusionOutput values.
// You can construct a concrete instance of `AccessModeSettingsExclusionInput` via:
//
//	AccessModeSettingsExclusionArgs{...}
type AccessModeSettingsExclusionInput interface {
	pulumi.Input

	ToAccessModeSettingsExclusionOutput() AccessModeSettingsExclusionOutput
	ToAccessModeSettingsExclusionOutputWithContext(context.Context) AccessModeSettingsExclusionOutput
}

// Properties that define the scope private link mode settings exclusion item. This setting applies to a specific private endpoint connection and overrides the default settings for that private endpoint connection.
type AccessModeSettingsExclusionArgs struct {
	// Specifies the access mode of ingestion through the specified private endpoint connection in the exclusion.
	IngestionAccessMode pulumi.StringPtrInput `pulumi:"ingestionAccessMode"`
	// The private endpoint connection name associated to the private endpoint on which we want to apply the specific access mode settings.
	PrivateEndpointConnectionName pulumi.StringPtrInput `pulumi:"privateEndpointConnectionName"`
	// Specifies the access mode of queries through the specified private endpoint connection in the exclusion.
	QueryAccessMode pulumi.StringPtrInput `pulumi:"queryAccessMode"`
}

func (AccessModeSettingsExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusion)(nil)).Elem()
}

func (i AccessModeSettingsExclusionArgs) ToAccessModeSettingsExclusionOutput() AccessModeSettingsExclusionOutput {
	return i.ToAccessModeSettingsExclusionOutputWithContext(context.Background())
}

func (i AccessModeSettingsExclusionArgs) ToAccessModeSettingsExclusionOutputWithContext(ctx context.Context) AccessModeSettingsExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsExclusionOutput)
}

// AccessModeSettingsExclusionArrayInput is an input type that accepts AccessModeSettingsExclusionArray and AccessModeSettingsExclusionArrayOutput values.
// You can construct a concrete instance of `AccessModeSettingsExclusionArrayInput` via:
//
//	AccessModeSettingsExclusionArray{ AccessModeSettingsExclusionArgs{...} }
type AccessModeSettingsExclusionArrayInput interface {
	pulumi.Input

	ToAccessModeSettingsExclusionArrayOutput() AccessModeSettingsExclusionArrayOutput
	ToAccessModeSettingsExclusionArrayOutputWithContext(context.Context) AccessModeSettingsExclusionArrayOutput
}

type AccessModeSettingsExclusionArray []AccessModeSettingsExclusionInput

func (AccessModeSettingsExclusionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusion)(nil)).Elem()
}

func (i AccessModeSettingsExclusionArray) ToAccessModeSettingsExclusionArrayOutput() AccessModeSettingsExclusionArrayOutput {
	return i.ToAccessModeSettingsExclusionArrayOutputWithContext(context.Background())
}

func (i AccessModeSettingsExclusionArray) ToAccessModeSettingsExclusionArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessModeSettingsExclusionArrayOutput)
}

// Properties that define the scope private link mode settings exclusion item. This setting applies to a specific private endpoint connection and overrides the default settings for that private endpoint connection.
type AccessModeSettingsExclusionOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusion)(nil)).Elem()
}

func (o AccessModeSettingsExclusionOutput) ToAccessModeSettingsExclusionOutput() AccessModeSettingsExclusionOutput {
	return o
}

func (o AccessModeSettingsExclusionOutput) ToAccessModeSettingsExclusionOutputWithContext(ctx context.Context) AccessModeSettingsExclusionOutput {
	return o
}

// Specifies the access mode of ingestion through the specified private endpoint connection in the exclusion.
func (o AccessModeSettingsExclusionOutput) IngestionAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusion) *string { return v.IngestionAccessMode }).(pulumi.StringPtrOutput)
}

// The private endpoint connection name associated to the private endpoint on which we want to apply the specific access mode settings.
func (o AccessModeSettingsExclusionOutput) PrivateEndpointConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusion) *string { return v.PrivateEndpointConnectionName }).(pulumi.StringPtrOutput)
}

// Specifies the access mode of queries through the specified private endpoint connection in the exclusion.
func (o AccessModeSettingsExclusionOutput) QueryAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusion) *string { return v.QueryAccessMode }).(pulumi.StringPtrOutput)
}

type AccessModeSettingsExclusionArrayOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusion)(nil)).Elem()
}

func (o AccessModeSettingsExclusionArrayOutput) ToAccessModeSettingsExclusionArrayOutput() AccessModeSettingsExclusionArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionArrayOutput) ToAccessModeSettingsExclusionArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionArrayOutput) Index(i pulumi.IntInput) AccessModeSettingsExclusionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessModeSettingsExclusion {
		return vs[0].([]AccessModeSettingsExclusion)[vs[1].(int)]
	}).(AccessModeSettingsExclusionOutput)
}

// Properties that define the scope private link mode settings exclusion item. This setting applies to a specific private endpoint connection and overrides the default settings for that private endpoint connection.
type AccessModeSettingsExclusionResponse struct {
	// Specifies the access mode of ingestion through the specified private endpoint connection in the exclusion.
	IngestionAccessMode *string `pulumi:"ingestionAccessMode"`
	// The private endpoint connection name associated to the private endpoint on which we want to apply the specific access mode settings.
	PrivateEndpointConnectionName *string `pulumi:"privateEndpointConnectionName"`
	// Specifies the access mode of queries through the specified private endpoint connection in the exclusion.
	QueryAccessMode *string `pulumi:"queryAccessMode"`
}

// Properties that define the scope private link mode settings exclusion item. This setting applies to a specific private endpoint connection and overrides the default settings for that private endpoint connection.
type AccessModeSettingsExclusionResponseOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsExclusionResponse)(nil)).Elem()
}

func (o AccessModeSettingsExclusionResponseOutput) ToAccessModeSettingsExclusionResponseOutput() AccessModeSettingsExclusionResponseOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseOutput) ToAccessModeSettingsExclusionResponseOutputWithContext(ctx context.Context) AccessModeSettingsExclusionResponseOutput {
	return o
}

// Specifies the access mode of ingestion through the specified private endpoint connection in the exclusion.
func (o AccessModeSettingsExclusionResponseOutput) IngestionAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusionResponse) *string { return v.IngestionAccessMode }).(pulumi.StringPtrOutput)
}

// The private endpoint connection name associated to the private endpoint on which we want to apply the specific access mode settings.
func (o AccessModeSettingsExclusionResponseOutput) PrivateEndpointConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusionResponse) *string { return v.PrivateEndpointConnectionName }).(pulumi.StringPtrOutput)
}

// Specifies the access mode of queries through the specified private endpoint connection in the exclusion.
func (o AccessModeSettingsExclusionResponseOutput) QueryAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessModeSettingsExclusionResponse) *string { return v.QueryAccessMode }).(pulumi.StringPtrOutput)
}

type AccessModeSettingsExclusionResponseArrayOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsExclusionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AccessModeSettingsExclusionResponse)(nil)).Elem()
}

func (o AccessModeSettingsExclusionResponseArrayOutput) ToAccessModeSettingsExclusionResponseArrayOutput() AccessModeSettingsExclusionResponseArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseArrayOutput) ToAccessModeSettingsExclusionResponseArrayOutputWithContext(ctx context.Context) AccessModeSettingsExclusionResponseArrayOutput {
	return o
}

func (o AccessModeSettingsExclusionResponseArrayOutput) Index(i pulumi.IntInput) AccessModeSettingsExclusionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AccessModeSettingsExclusionResponse {
		return vs[0].([]AccessModeSettingsExclusionResponse)[vs[1].(int)]
	}).(AccessModeSettingsExclusionResponseOutput)
}

// Properties that define the scope private link mode settings.
type AccessModeSettingsResponse struct {
	// List of exclusions that override the default access mode settings for specific private endpoint connections.
	Exclusions []AccessModeSettingsExclusionResponse `pulumi:"exclusions"`
	// Specifies the default access mode of ingestion through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
	IngestionAccessMode string `pulumi:"ingestionAccessMode"`
	// Specifies the default access mode of queries through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
	QueryAccessMode string `pulumi:"queryAccessMode"`
}

// Properties that define the scope private link mode settings.
type AccessModeSettingsResponseOutput struct{ *pulumi.OutputState }

func (AccessModeSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessModeSettingsResponse)(nil)).Elem()
}

func (o AccessModeSettingsResponseOutput) ToAccessModeSettingsResponseOutput() AccessModeSettingsResponseOutput {
	return o
}

func (o AccessModeSettingsResponseOutput) ToAccessModeSettingsResponseOutputWithContext(ctx context.Context) AccessModeSettingsResponseOutput {
	return o
}

// List of exclusions that override the default access mode settings for specific private endpoint connections.
func (o AccessModeSettingsResponseOutput) Exclusions() AccessModeSettingsExclusionResponseArrayOutput {
	return o.ApplyT(func(v AccessModeSettingsResponse) []AccessModeSettingsExclusionResponse { return v.Exclusions }).(AccessModeSettingsExclusionResponseArrayOutput)
}

// Specifies the default access mode of ingestion through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
func (o AccessModeSettingsResponseOutput) IngestionAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettingsResponse) string { return v.IngestionAccessMode }).(pulumi.StringOutput)
}

// Specifies the default access mode of queries through associated private endpoints in scope. If not specified default value is 'Open'. You can override this default setting for a specific private endpoint connection by adding an exclusion in the 'exclusions' array.
func (o AccessModeSettingsResponseOutput) QueryAccessMode() pulumi.StringOutput {
	return o.ApplyT(func(v AccessModeSettingsResponse) string { return v.QueryAccessMode }).(pulumi.StringOutput)
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnectionResponse struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// The Private Endpoint Connection resource.
type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The resource of private end point.
func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

// The provisioning state of the private endpoint connection resource.
func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

// The Private Endpoint resource.
type PrivateEndpointResponse struct {
	// The ARM identifier for Private Endpoint
	Id string `pulumi:"id"`
}

// The Private Endpoint resource.
type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

// The ARM identifier for Private Endpoint
func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

// The ARM identifier for Private Endpoint
func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `pulumi:"actionsRequired"`
	// The reason for approval/rejection of the connection.
	Description *string `pulumi:"description"`
	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *string `pulumi:"status"`
}

// PrivateLinkServiceConnectionStateInput is an input type that accepts PrivateLinkServiceConnectionStateArgs and PrivateLinkServiceConnectionStateOutput values.
// You can construct a concrete instance of `PrivateLinkServiceConnectionStateInput` via:
//
//	PrivateLinkServiceConnectionStateArgs{...}
type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionStateArgs struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	// The reason for approval/rejection of the connection.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

// A message indicating if changes on the service provider require any updates on the consumer.
func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

// The reason for approval/rejection of the connection.
func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionStateResponse struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `pulumi:"actionsRequired"`
	// The reason for approval/rejection of the connection.
	Description *string `pulumi:"description"`
	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *string `pulumi:"status"`
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

// A message indicating if changes on the service provider require any updates on the consumer.
func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

// The reason for approval/rejection of the connection.
func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(AccessModeSettingsOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionArrayOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionResponseOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsExclusionResponseArrayOutput{})
	pulumi.RegisterOutputType(AccessModeSettingsResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}