// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Role Assignments
type RoleAssignment struct {
	pulumi.CustomResourceState

	// The Delegation flag for the role assignment
	CanDelegate pulumi.BoolPtrOutput `pulumi:"canDelegate"`
	// The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
	Condition pulumi.StringPtrOutput `pulumi:"condition"`
	// Version of the condition. Currently accepted value is '2.0'
	ConditionVersion pulumi.StringPtrOutput `pulumi:"conditionVersion"`
	// Id of the user who created the assignment
	CreatedBy pulumi.StringPtrOutput `pulumi:"createdBy"`
	// Time it was created
	CreatedOn pulumi.StringPtrOutput `pulumi:"createdOn"`
	// Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId pulumi.StringPtrOutput `pulumi:"delegatedManagedIdentityResourceId"`
	// Description of role assignment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The role assignment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal ID.
	PrincipalId pulumi.StringPtrOutput `pulumi:"principalId"`
	// The principal type of the assigned principal ID.
	PrincipalType pulumi.StringPtrOutput `pulumi:"principalType"`
	// The role definition ID.
	RoleDefinitionId pulumi.StringPtrOutput `pulumi:"roleDefinitionId"`
	// The role assignment scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// The role assignment type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Id of the user who updated the assignment
	UpdatedBy pulumi.StringPtrOutput `pulumi:"updatedBy"`
	// Time it was updated
	UpdatedOn pulumi.StringPtrOutput `pulumi:"updatedOn"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.RoleDefinitionId == nil {
		return nil, errors.New("invalid value for required argument 'RoleDefinitionId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.PrincipalType == nil {
		args.PrincipalType = pulumi.StringPtr("User")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:authorization/v20150701:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20171001preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180101preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20180901preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200301preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20200801preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20201001preview:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20220401:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:authorization:RoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RoleAssignment
	err := ctx.RegisterResource("azure-native:authorization/v20200401preview:RoleAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssignment gets an existing RoleAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssignmentState, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	var resource RoleAssignment
	err := ctx.ReadResource("azure-native:authorization/v20200401preview:RoleAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssignment resources.
type roleAssignmentState struct {
}

type RoleAssignmentState struct {
}

func (RoleAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentState)(nil)).Elem()
}

type roleAssignmentArgs struct {
	// The delegation flag used for creating a role assignment
	CanDelegate *bool `pulumi:"canDelegate"`
	// The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
	Condition *string `pulumi:"condition"`
	// Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `pulumi:"conditionVersion"`
	// Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `pulumi:"delegatedManagedIdentityResourceId"`
	// Description of role assignment
	Description *string `pulumi:"description"`
	// The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
	PrincipalId string `pulumi:"principalId"`
	// The principal type of the assigned principal ID.
	PrincipalType *string `pulumi:"principalType"`
	// A GUID for the role assignment to create. The name must be unique and different for each role assignment.
	RoleAssignmentName *string `pulumi:"roleAssignmentName"`
	// The role definition ID used in the role assignment.
	RoleDefinitionId string `pulumi:"roleDefinitionId"`
	// The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The delegation flag used for creating a role assignment
	CanDelegate pulumi.BoolPtrInput
	// The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
	Condition pulumi.StringPtrInput
	// Version of the condition. Currently accepted value is '2.0'
	ConditionVersion pulumi.StringPtrInput
	// Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId pulumi.StringPtrInput
	// Description of role assignment
	Description pulumi.StringPtrInput
	// The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
	PrincipalId pulumi.StringInput
	// The principal type of the assigned principal ID.
	PrincipalType pulumi.StringPtrInput
	// A GUID for the role assignment to create. The name must be unique and different for each role assignment.
	RoleAssignmentName pulumi.StringPtrInput
	// The role definition ID used in the role assignment.
	RoleDefinitionId pulumi.StringInput
	// The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
	Scope pulumi.StringInput
}

func (RoleAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssignmentArgs)(nil)).Elem()
}

type RoleAssignmentInput interface {
	pulumi.Input

	ToRoleAssignmentOutput() RoleAssignmentOutput
	ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput
}

func (*RoleAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (i *RoleAssignment) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return i.ToRoleAssignmentOutputWithContext(context.Background())
}

func (i *RoleAssignment) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentOutput)
}

type RoleAssignmentOutput struct{ *pulumi.OutputState }

func (RoleAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAssignment)(nil)).Elem()
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutput() RoleAssignmentOutput {
	return o
}

func (o RoleAssignmentOutput) ToRoleAssignmentOutputWithContext(ctx context.Context) RoleAssignmentOutput {
	return o
}

// The Delegation flag for the role assignment
func (o RoleAssignmentOutput) CanDelegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.BoolPtrOutput { return v.CanDelegate }).(pulumi.BoolPtrOutput)
}

// The conditions on the role assignment. This limits the resources it can be assigned to. e.g.: @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'
func (o RoleAssignmentOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.Condition }).(pulumi.StringPtrOutput)
}

// Version of the condition. Currently accepted value is '2.0'
func (o RoleAssignmentOutput) ConditionVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.ConditionVersion }).(pulumi.StringPtrOutput)
}

// Id of the user who created the assignment
func (o RoleAssignmentOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// Time it was created
func (o RoleAssignmentOutput) CreatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.CreatedOn }).(pulumi.StringPtrOutput)
}

// Id of the delegated managed identity resource
func (o RoleAssignmentOutput) DelegatedManagedIdentityResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.DelegatedManagedIdentityResourceId }).(pulumi.StringPtrOutput)
}

// Description of role assignment
func (o RoleAssignmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The role assignment name.
func (o RoleAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principal ID.
func (o RoleAssignmentOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

// The principal type of the assigned principal ID.
func (o RoleAssignmentOutput) PrincipalType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.PrincipalType }).(pulumi.StringPtrOutput)
}

// The role definition ID.
func (o RoleAssignmentOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

// The role assignment scope.
func (o RoleAssignmentOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// The role assignment type.
func (o RoleAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Id of the user who updated the assignment
func (o RoleAssignmentOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.UpdatedBy }).(pulumi.StringPtrOutput)
}

// Time it was updated
func (o RoleAssignmentOutput) UpdatedOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringPtrOutput { return v.UpdatedOn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
}
