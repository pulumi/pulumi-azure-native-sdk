// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

// The entity.
type EntityInfoResponse struct {
	// The friendly name of the management group.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified ID for the entity.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id string `pulumi:"id"`
	// The users specific permissions to this item.
	InheritedPermissions *string `pulumi:"inheritedPermissions"`
	// The name of the entity. For example, 00000000-0000-0000-0000-000000000000
	Name string `pulumi:"name"`
	// Number of children is the number of Groups that are exactly one level underneath the current Group.
	NumberOfChildGroups *int `pulumi:"numberOfChildGroups"`
	// Number of children is the number of Groups and Subscriptions that are exactly one level underneath the current Group.
	NumberOfChildren    *int `pulumi:"numberOfChildren"`
	NumberOfDescendants *int `pulumi:"numberOfDescendants"`
	// (Optional) The ID of the parent management group.
	Parent *EntityParentGroupInfoResponse `pulumi:"parent"`
	// The parent display name chain from the root group to the immediate parent
	ParentDisplayNameChain []string `pulumi:"parentDisplayNameChain"`
	// The parent name chain from the root group to the immediate parent
	ParentNameChain []string `pulumi:"parentNameChain"`
	// The users specific permissions to this item.
	Permissions *string `pulumi:"permissions"`
	// The AAD Tenant ID associated with the entity. For example, 00000000-0000-0000-0000-000000000000
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. For example, Microsoft.Management/managementGroups
	Type string `pulumi:"type"`
}

// The entity.
type EntityInfoResponseOutput struct{ *pulumi.OutputState }

func (EntityInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInfoResponse)(nil)).Elem()
}

func (o EntityInfoResponseOutput) ToEntityInfoResponseOutput() EntityInfoResponseOutput {
	return o
}

func (o EntityInfoResponseOutput) ToEntityInfoResponseOutputWithContext(ctx context.Context) EntityInfoResponseOutput {
	return o
}

// The friendly name of the management group.
func (o EntityInfoResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the entity.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o EntityInfoResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EntityInfoResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The users specific permissions to this item.
func (o EntityInfoResponseOutput) InheritedPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *string { return v.InheritedPermissions }).(pulumi.StringPtrOutput)
}

// The name of the entity. For example, 00000000-0000-0000-0000-000000000000
func (o EntityInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EntityInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Number of children is the number of Groups that are exactly one level underneath the current Group.
func (o EntityInfoResponseOutput) NumberOfChildGroups() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *int { return v.NumberOfChildGroups }).(pulumi.IntPtrOutput)
}

// Number of children is the number of Groups and Subscriptions that are exactly one level underneath the current Group.
func (o EntityInfoResponseOutput) NumberOfChildren() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *int { return v.NumberOfChildren }).(pulumi.IntPtrOutput)
}

func (o EntityInfoResponseOutput) NumberOfDescendants() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *int { return v.NumberOfDescendants }).(pulumi.IntPtrOutput)
}

// (Optional) The ID of the parent management group.
func (o EntityInfoResponseOutput) Parent() EntityParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *EntityParentGroupInfoResponse { return v.Parent }).(EntityParentGroupInfoResponsePtrOutput)
}

// The parent display name chain from the root group to the immediate parent
func (o EntityInfoResponseOutput) ParentDisplayNameChain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EntityInfoResponse) []string { return v.ParentDisplayNameChain }).(pulumi.StringArrayOutput)
}

// The parent name chain from the root group to the immediate parent
func (o EntityInfoResponseOutput) ParentNameChain() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EntityInfoResponse) []string { return v.ParentNameChain }).(pulumi.StringArrayOutput)
}

// The users specific permissions to this item.
func (o EntityInfoResponseOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *string { return v.Permissions }).(pulumi.StringPtrOutput)
}

// The AAD Tenant ID associated with the entity. For example, 00000000-0000-0000-0000-000000000000
func (o EntityInfoResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInfoResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. For example, Microsoft.Management/managementGroups
func (o EntityInfoResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EntityInfoResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EntityInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (EntityInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityInfoResponse)(nil)).Elem()
}

func (o EntityInfoResponseArrayOutput) ToEntityInfoResponseArrayOutput() EntityInfoResponseArrayOutput {
	return o
}

func (o EntityInfoResponseArrayOutput) ToEntityInfoResponseArrayOutputWithContext(ctx context.Context) EntityInfoResponseArrayOutput {
	return o
}

func (o EntityInfoResponseArrayOutput) Index(i pulumi.IntInput) EntityInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityInfoResponse {
		return vs[0].([]EntityInfoResponse)[vs[1].(int)]
	}).(EntityInfoResponseOutput)
}

// (Optional) The ID of the parent management group.
type EntityParentGroupInfoResponse struct {
	// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id *string `pulumi:"id"`
}

// (Optional) The ID of the parent management group.
type EntityParentGroupInfoResponseOutput struct{ *pulumi.OutputState }

func (EntityParentGroupInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityParentGroupInfoResponse)(nil)).Elem()
}

func (o EntityParentGroupInfoResponseOutput) ToEntityParentGroupInfoResponseOutput() EntityParentGroupInfoResponseOutput {
	return o
}

func (o EntityParentGroupInfoResponseOutput) ToEntityParentGroupInfoResponseOutputWithContext(ctx context.Context) EntityParentGroupInfoResponseOutput {
	return o
}

// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o EntityParentGroupInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityParentGroupInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type EntityParentGroupInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (EntityParentGroupInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityParentGroupInfoResponse)(nil)).Elem()
}

func (o EntityParentGroupInfoResponsePtrOutput) ToEntityParentGroupInfoResponsePtrOutput() EntityParentGroupInfoResponsePtrOutput {
	return o
}

func (o EntityParentGroupInfoResponsePtrOutput) ToEntityParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) EntityParentGroupInfoResponsePtrOutput {
	return o
}

func (o EntityParentGroupInfoResponsePtrOutput) Elem() EntityParentGroupInfoResponseOutput {
	return o.ApplyT(func(v *EntityParentGroupInfoResponse) EntityParentGroupInfoResponse {
		if v != nil {
			return *v
		}
		var ret EntityParentGroupInfoResponse
		return ret
	}).(EntityParentGroupInfoResponseOutput)
}

// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o EntityParentGroupInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EntityInfoResponseOutput{})
	pulumi.RegisterOutputType(EntityInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(EntityParentGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(EntityParentGroupInfoResponsePtrOutput{})
}
