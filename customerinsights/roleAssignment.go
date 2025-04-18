// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Role Assignment resource format.
//
// Uses Azure REST API version 2017-04-26. In version 2.x of the Azure Native provider, it used API version 2017-04-26.
type RoleAssignment struct {
	pulumi.CustomResourceState

	// The name of the metadata object.
	AssignmentName pulumi.StringOutput `pulumi:"assignmentName"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Widget types set for the assignment.
	ConflationPolicies ResourceSetDescriptionResponsePtrOutput `pulumi:"conflationPolicies"`
	// Connectors set for the assignment.
	Connectors ResourceSetDescriptionResponsePtrOutput `pulumi:"connectors"`
	// Localized description for the metadata.
	Description pulumi.StringMapOutput `pulumi:"description"`
	// Localized display names for the metadata.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// Interactions set for the assignment.
	Interactions ResourceSetDescriptionResponsePtrOutput `pulumi:"interactions"`
	// Kpis set for the assignment.
	Kpis ResourceSetDescriptionResponsePtrOutput `pulumi:"kpis"`
	// Links set for the assignment.
	Links ResourceSetDescriptionResponsePtrOutput `pulumi:"links"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The principals being assigned to.
	Principals AssignmentPrincipalResponseArrayOutput `pulumi:"principals"`
	// Profiles set for the assignment.
	Profiles ResourceSetDescriptionResponsePtrOutput `pulumi:"profiles"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The Role assignments set for the relationship links.
	RelationshipLinks ResourceSetDescriptionResponsePtrOutput `pulumi:"relationshipLinks"`
	// The Role assignments set for the relationships.
	Relationships ResourceSetDescriptionResponsePtrOutput `pulumi:"relationships"`
	// Type of roles.
	Role pulumi.StringOutput `pulumi:"role"`
	// The Role assignments set for the assignment.
	RoleAssignments ResourceSetDescriptionResponsePtrOutput `pulumi:"roleAssignments"`
	// Sas Policies set for the assignment.
	SasPolicies ResourceSetDescriptionResponsePtrOutput `pulumi:"sasPolicies"`
	// The Role assignments set for the assignment.
	Segments ResourceSetDescriptionResponsePtrOutput `pulumi:"segments"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Views set for the assignment.
	Views ResourceSetDescriptionResponsePtrOutput `pulumi:"views"`
	// Widget types set for the assignment.
	WidgetTypes ResourceSetDescriptionResponsePtrOutput `pulumi:"widgetTypes"`
}

// NewRoleAssignment registers a new resource with the given unique name, arguments, and options.
func NewRoleAssignment(ctx *pulumi.Context,
	name string, args *RoleAssignmentArgs, opts ...pulumi.ResourceOption) (*RoleAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubName == nil {
		return nil, errors.New("invalid value for required argument 'HubName'")
	}
	if args.Principals == nil {
		return nil, errors.New("invalid value for required argument 'Principals'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:customerinsights/v20170101:RoleAssignment"),
		},
		{
			Type: pulumi.String("azure-native:customerinsights/v20170426:RoleAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RoleAssignment
	err := ctx.RegisterResource("azure-native:customerinsights:RoleAssignment", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:customerinsights:RoleAssignment", name, id, state, &resource, opts...)
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
	// The assignment name
	AssignmentName *string `pulumi:"assignmentName"`
	// Widget types set for the assignment.
	ConflationPolicies *ResourceSetDescription `pulumi:"conflationPolicies"`
	// Connectors set for the assignment.
	Connectors *ResourceSetDescription `pulumi:"connectors"`
	// Localized description for the metadata.
	Description map[string]string `pulumi:"description"`
	// Localized display names for the metadata.
	DisplayName map[string]string `pulumi:"displayName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Interactions set for the assignment.
	Interactions *ResourceSetDescription `pulumi:"interactions"`
	// Kpis set for the assignment.
	Kpis *ResourceSetDescription `pulumi:"kpis"`
	// Links set for the assignment.
	Links *ResourceSetDescription `pulumi:"links"`
	// The principals being assigned to.
	Principals []AssignmentPrincipal `pulumi:"principals"`
	// Profiles set for the assignment.
	Profiles *ResourceSetDescription `pulumi:"profiles"`
	// The Role assignments set for the relationship links.
	RelationshipLinks *ResourceSetDescription `pulumi:"relationshipLinks"`
	// The Role assignments set for the relationships.
	Relationships *ResourceSetDescription `pulumi:"relationships"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Type of roles.
	Role RoleTypes `pulumi:"role"`
	// The Role assignments set for the assignment.
	RoleAssignments *ResourceSetDescription `pulumi:"roleAssignments"`
	// Sas Policies set for the assignment.
	SasPolicies *ResourceSetDescription `pulumi:"sasPolicies"`
	// The Role assignments set for the assignment.
	Segments *ResourceSetDescription `pulumi:"segments"`
	// Views set for the assignment.
	Views *ResourceSetDescription `pulumi:"views"`
	// Widget types set for the assignment.
	WidgetTypes *ResourceSetDescription `pulumi:"widgetTypes"`
}

// The set of arguments for constructing a RoleAssignment resource.
type RoleAssignmentArgs struct {
	// The assignment name
	AssignmentName pulumi.StringPtrInput
	// Widget types set for the assignment.
	ConflationPolicies ResourceSetDescriptionPtrInput
	// Connectors set for the assignment.
	Connectors ResourceSetDescriptionPtrInput
	// Localized description for the metadata.
	Description pulumi.StringMapInput
	// Localized display names for the metadata.
	DisplayName pulumi.StringMapInput
	// The name of the hub.
	HubName pulumi.StringInput
	// Interactions set for the assignment.
	Interactions ResourceSetDescriptionPtrInput
	// Kpis set for the assignment.
	Kpis ResourceSetDescriptionPtrInput
	// Links set for the assignment.
	Links ResourceSetDescriptionPtrInput
	// The principals being assigned to.
	Principals AssignmentPrincipalArrayInput
	// Profiles set for the assignment.
	Profiles ResourceSetDescriptionPtrInput
	// The Role assignments set for the relationship links.
	RelationshipLinks ResourceSetDescriptionPtrInput
	// The Role assignments set for the relationships.
	Relationships ResourceSetDescriptionPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Type of roles.
	Role RoleTypesInput
	// The Role assignments set for the assignment.
	RoleAssignments ResourceSetDescriptionPtrInput
	// Sas Policies set for the assignment.
	SasPolicies ResourceSetDescriptionPtrInput
	// The Role assignments set for the assignment.
	Segments ResourceSetDescriptionPtrInput
	// Views set for the assignment.
	Views ResourceSetDescriptionPtrInput
	// Widget types set for the assignment.
	WidgetTypes ResourceSetDescriptionPtrInput
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

// The name of the metadata object.
func (o RoleAssignmentOutput) AssignmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.AssignmentName }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o RoleAssignmentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Widget types set for the assignment.
func (o RoleAssignmentOutput) ConflationPolicies() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.ConflationPolicies }).(ResourceSetDescriptionResponsePtrOutput)
}

// Connectors set for the assignment.
func (o RoleAssignmentOutput) Connectors() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Connectors }).(ResourceSetDescriptionResponsePtrOutput)
}

// Localized description for the metadata.
func (o RoleAssignmentOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringMapOutput { return v.Description }).(pulumi.StringMapOutput)
}

// Localized display names for the metadata.
func (o RoleAssignmentOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringMapOutput { return v.DisplayName }).(pulumi.StringMapOutput)
}

// Interactions set for the assignment.
func (o RoleAssignmentOutput) Interactions() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Interactions }).(ResourceSetDescriptionResponsePtrOutput)
}

// Kpis set for the assignment.
func (o RoleAssignmentOutput) Kpis() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Kpis }).(ResourceSetDescriptionResponsePtrOutput)
}

// Links set for the assignment.
func (o RoleAssignmentOutput) Links() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Links }).(ResourceSetDescriptionResponsePtrOutput)
}

// Resource name.
func (o RoleAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principals being assigned to.
func (o RoleAssignmentOutput) Principals() AssignmentPrincipalResponseArrayOutput {
	return o.ApplyT(func(v *RoleAssignment) AssignmentPrincipalResponseArrayOutput { return v.Principals }).(AssignmentPrincipalResponseArrayOutput)
}

// Profiles set for the assignment.
func (o RoleAssignmentOutput) Profiles() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Profiles }).(ResourceSetDescriptionResponsePtrOutput)
}

// Provisioning state.
func (o RoleAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The Role assignments set for the relationship links.
func (o RoleAssignmentOutput) RelationshipLinks() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.RelationshipLinks }).(ResourceSetDescriptionResponsePtrOutput)
}

// The Role assignments set for the relationships.
func (o RoleAssignmentOutput) Relationships() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Relationships }).(ResourceSetDescriptionResponsePtrOutput)
}

// Type of roles.
func (o RoleAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The Role assignments set for the assignment.
func (o RoleAssignmentOutput) RoleAssignments() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.RoleAssignments }).(ResourceSetDescriptionResponsePtrOutput)
}

// Sas Policies set for the assignment.
func (o RoleAssignmentOutput) SasPolicies() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.SasPolicies }).(ResourceSetDescriptionResponsePtrOutput)
}

// The Role assignments set for the assignment.
func (o RoleAssignmentOutput) Segments() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Segments }).(ResourceSetDescriptionResponsePtrOutput)
}

// The hub name.
func (o RoleAssignmentOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type.
func (o RoleAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Views set for the assignment.
func (o RoleAssignmentOutput) Views() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.Views }).(ResourceSetDescriptionResponsePtrOutput)
}

// Widget types set for the assignment.
func (o RoleAssignmentOutput) WidgetTypes() ResourceSetDescriptionResponsePtrOutput {
	return o.ApplyT(func(v *RoleAssignment) ResourceSetDescriptionResponsePtrOutput { return v.WidgetTypes }).(ResourceSetDescriptionResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RoleAssignmentOutput{})
}
