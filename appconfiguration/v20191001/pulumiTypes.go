// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An API key used for authenticating with a configuration store endpoint.
type ApiKeyResponse struct {
	// A connection string that can be used by supporting clients for authentication.
	ConnectionString string `pulumi:"connectionString"`
	// The key ID.
	Id string `pulumi:"id"`
	// The last time any of the key's properties were modified.
	LastModified string `pulumi:"lastModified"`
	// A name for the key describing its usage.
	Name string `pulumi:"name"`
	// Whether this key can only be used for read operations.
	ReadOnly bool `pulumi:"readOnly"`
	// The value of the key that is used for authentication purposes.
	Value string `pulumi:"value"`
}

// An API key used for authenticating with a configuration store endpoint.
type ApiKeyResponseOutput struct{ *pulumi.OutputState }

func (ApiKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKeyResponse)(nil)).Elem()
}

func (o ApiKeyResponseOutput) ToApiKeyResponseOutput() ApiKeyResponseOutput {
	return o
}

func (o ApiKeyResponseOutput) ToApiKeyResponseOutputWithContext(ctx context.Context) ApiKeyResponseOutput {
	return o
}

// A connection string that can be used by supporting clients for authentication.
func (o ApiKeyResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

// The key ID.
func (o ApiKeyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The last time any of the key's properties were modified.
func (o ApiKeyResponseOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.LastModified }).(pulumi.StringOutput)
}

// A name for the key describing its usage.
func (o ApiKeyResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Whether this key can only be used for read operations.
func (o ApiKeyResponseOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v ApiKeyResponse) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

// The value of the key that is used for authentication purposes.
func (o ApiKeyResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ApiKeyResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ApiKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiKeyResponse)(nil)).Elem()
}

func (o ApiKeyResponseArrayOutput) ToApiKeyResponseArrayOutput() ApiKeyResponseArrayOutput {
	return o
}

func (o ApiKeyResponseArrayOutput) ToApiKeyResponseArrayOutputWithContext(ctx context.Context) ApiKeyResponseArrayOutput {
	return o
}

func (o ApiKeyResponseArrayOutput) Index(i pulumi.IntInput) ApiKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiKeyResponse {
		return vs[0].([]ApiKeyResponse)[vs[1].(int)]
	}).(ApiKeyResponseOutput)
}

type ResourceIdentity struct {
	// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
	Type *string `pulumi:"type"`
	// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}

// ResourceIdentityInput is an input type that accepts ResourceIdentityArgs and ResourceIdentityOutput values.
// You can construct a concrete instance of `ResourceIdentityInput` via:
//
//	ResourceIdentityArgs{...}
type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities pulumi.MapInput `pulumi:"userAssignedIdentities"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}

// ResourceIdentityPtrInput is an input type that accepts ResourceIdentityArgs, ResourceIdentityPtr and ResourceIdentityPtrOutput values.
// You can construct a concrete instance of `ResourceIdentityPtrInput` via:
//
//	        ResourceIdentityArgs{...}
//
//	or:
//
//	        nil
type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
func (o ResourceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ResourceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
func (o ResourceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ResourceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ResourceIdentityResponse struct {
	// The principal id of the identity. This property will only be provided for a system-assigned identity.
	PrincipalId string `pulumi:"principalId"`
	// The tenant id associated with the resource's identity. This property will only be provided for a system-assigned identity.
	TenantId string `pulumi:"tenantId"`
	// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
	Type *string `pulumi:"type"`
	// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]UserIdentityResponse `pulumi:"userAssignedIdentities"`
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

// The principal id of the identity. This property will only be provided for a system-assigned identity.
func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

// The tenant id associated with the resource's identity. This property will only be provided for a system-assigned identity.
func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
func (o ResourceIdentityResponseOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) map[string]UserIdentityResponse { return v.UserAssignedIdentities }).(UserIdentityResponseMapOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

// The principal id of the identity. This property will only be provided for a system-assigned identity.
func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

// The tenant id associated with the resource's identity. This property will only be provided for a system-assigned identity.
func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
func (o ResourceIdentityResponsePtrOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) map[string]UserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityResponseMapOutput)
}

// Describes a configuration store SKU.
type Sku struct {
	// The SKU name of the configuration store.
	Name string `pulumi:"name"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//	SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// Describes a configuration store SKU.
type SkuArgs struct {
	// The SKU name of the configuration store.
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

// Describes a configuration store SKU.
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

// The SKU name of the configuration store.
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

// Describes a configuration store SKU.
type SkuResponse struct {
	// The SKU name of the configuration store.
	Name string `pulumi:"name"`
}

// Describes a configuration store SKU.
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

// The SKU name of the configuration store.
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type UserIdentityResponse struct {
	// The client ID of the user-assigned identity.
	ClientId string `pulumi:"clientId"`
	// The principal ID of the user-assigned identity.
	PrincipalId string `pulumi:"principalId"`
}

type UserIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return o
}

// The client ID of the user-assigned identity.
func (o UserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

// The principal ID of the user-assigned identity.
func (o UserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutput() UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutputWithContext(ctx context.Context) UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityResponse {
		return vs[0].(map[string]UserIdentityResponse)[vs[1].(string)]
	}).(UserIdentityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiKeyResponseOutput{})
	pulumi.RegisterOutputType(ApiKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseMapOutput{})
}