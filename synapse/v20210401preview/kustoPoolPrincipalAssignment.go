// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a cluster principal assignment.
type KustoPoolPrincipalAssignment struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// The principal name
	PrincipalName pulumi.StringOutput `pulumi:"principalName"`
	// Principal type.
	PrincipalType pulumi.StringOutput `pulumi:"principalType"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Cluster principal role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id of the principal
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The tenant name of the principal
	TenantName pulumi.StringOutput `pulumi:"tenantName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKustoPoolPrincipalAssignment registers a new resource with the given unique name, arguments, and options.
func NewKustoPoolPrincipalAssignment(ctx *pulumi.Context,
	name string, args *KustoPoolPrincipalAssignmentArgs, opts ...pulumi.ResourceOption) (*KustoPoolPrincipalAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.PrincipalType == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:KustoPoolPrincipalAssignment"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:KustoPoolPrincipalAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolPrincipalAssignment
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:KustoPoolPrincipalAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKustoPoolPrincipalAssignment gets an existing KustoPoolPrincipalAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKustoPoolPrincipalAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolPrincipalAssignmentState, opts ...pulumi.ResourceOption) (*KustoPoolPrincipalAssignment, error) {
	var resource KustoPoolPrincipalAssignment
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:KustoPoolPrincipalAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KustoPoolPrincipalAssignment resources.
type kustoPoolPrincipalAssignmentState struct {
}

type KustoPoolPrincipalAssignmentState struct {
}

func (KustoPoolPrincipalAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolPrincipalAssignmentState)(nil)).Elem()
}

type kustoPoolPrincipalAssignmentArgs struct {
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName *string `pulumi:"principalAssignmentName"`
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId string `pulumi:"principalId"`
	// Principal type.
	PrincipalType string `pulumi:"principalType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cluster principal role.
	Role string `pulumi:"role"`
	// The tenant id of the principal
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a KustoPoolPrincipalAssignment resource.
type KustoPoolPrincipalAssignmentArgs struct {
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput
	// The name of the Kusto principalAssignment.
	PrincipalAssignmentName pulumi.StringPtrInput
	// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
	PrincipalId pulumi.StringInput
	// Principal type.
	PrincipalType pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Cluster principal role.
	Role pulumi.StringInput
	// The tenant id of the principal
	TenantId pulumi.StringPtrInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (KustoPoolPrincipalAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolPrincipalAssignmentArgs)(nil)).Elem()
}

type KustoPoolPrincipalAssignmentInput interface {
	pulumi.Input

	ToKustoPoolPrincipalAssignmentOutput() KustoPoolPrincipalAssignmentOutput
	ToKustoPoolPrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolPrincipalAssignmentOutput
}

func (*KustoPoolPrincipalAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolPrincipalAssignment)(nil)).Elem()
}

func (i *KustoPoolPrincipalAssignment) ToKustoPoolPrincipalAssignmentOutput() KustoPoolPrincipalAssignmentOutput {
	return i.ToKustoPoolPrincipalAssignmentOutputWithContext(context.Background())
}

func (i *KustoPoolPrincipalAssignment) ToKustoPoolPrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolPrincipalAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolPrincipalAssignmentOutput)
}

type KustoPoolPrincipalAssignmentOutput struct{ *pulumi.OutputState }

func (KustoPoolPrincipalAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolPrincipalAssignment)(nil)).Elem()
}

func (o KustoPoolPrincipalAssignmentOutput) ToKustoPoolPrincipalAssignmentOutput() KustoPoolPrincipalAssignmentOutput {
	return o
}

func (o KustoPoolPrincipalAssignmentOutput) ToKustoPoolPrincipalAssignmentOutputWithContext(ctx context.Context) KustoPoolPrincipalAssignmentOutput {
	return o
}

// The name of the resource
func (o KustoPoolPrincipalAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The principal ID assigned to the cluster principal. It can be a user email, application ID, or security group name.
func (o KustoPoolPrincipalAssignmentOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// The principal name
func (o KustoPoolPrincipalAssignmentOutput) PrincipalName() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.PrincipalName }).(pulumi.StringOutput)
}

// Principal type.
func (o KustoPoolPrincipalAssignmentOutput) PrincipalType() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.PrincipalType }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o KustoPoolPrincipalAssignmentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Cluster principal role.
func (o KustoPoolPrincipalAssignmentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o KustoPoolPrincipalAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id of the principal
func (o KustoPoolPrincipalAssignmentOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The tenant name of the principal
func (o KustoPoolPrincipalAssignmentOutput) TenantName() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.TenantName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o KustoPoolPrincipalAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolPrincipalAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolPrincipalAssignmentOutput{})
}