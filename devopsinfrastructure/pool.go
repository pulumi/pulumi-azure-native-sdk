// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devopsinfrastructure

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Concrete tracked resource types can be created by aliasing this type using a specific property type.
//
// Uses Azure REST API version 2025-01-21. In version 2.x of the Azure Native provider, it used API version 2023-10-30-preview.
//
// Other available API versions: 2023-10-30-preview, 2023-12-13-preview, 2024-03-26-preview, 2024-04-04-preview, 2024-10-19. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devopsinfrastructure [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Pool struct {
	pulumi.CustomResourceState

	// Defines how the machine will be handled once it executed a job.
	AgentProfile pulumi.AnyOutput `pulumi:"agentProfile"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The resource id of the DevCenter Project the pool belongs to.
	DevCenterProjectResourceId pulumi.StringOutput `pulumi:"devCenterProjectResourceId"`
	// Defines the type of fabric the agent will run on.
	FabricProfile VmssFabricProfileResponseOutput `pulumi:"fabricProfile"`
	// The managed service identities assigned to this resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Defines how many resources can there be created at any given time.
	MaximumConcurrency pulumi.IntOutput `pulumi:"maximumConcurrency"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the organization in which the pool will be used.
	OrganizationProfile pulumi.AnyOutput `pulumi:"organizationProfile"`
	// The status of the current operation.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentProfile == nil {
		return nil, errors.New("invalid value for required argument 'AgentProfile'")
	}
	if args.DevCenterProjectResourceId == nil {
		return nil, errors.New("invalid value for required argument 'DevCenterProjectResourceId'")
	}
	if args.FabricProfile == nil {
		return nil, errors.New("invalid value for required argument 'FabricProfile'")
	}
	if args.MaximumConcurrency == nil {
		return nil, errors.New("invalid value for required argument 'MaximumConcurrency'")
	}
	if args.OrganizationProfile == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devopsinfrastructure/v20231030preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devopsinfrastructure/v20231213preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devopsinfrastructure/v20240326preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devopsinfrastructure/v20240404preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devopsinfrastructure/v20241019:Pool"),
		},
		{
			Type: pulumi.String("azure-native:devopsinfrastructure/v20250121:Pool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Pool
	err := ctx.RegisterResource("azure-native:devopsinfrastructure:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:devopsinfrastructure:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
}

type PoolState struct {
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// Defines how the machine will be handled once it executed a job.
	AgentProfile interface{} `pulumi:"agentProfile"`
	// The resource id of the DevCenter Project the pool belongs to.
	DevCenterProjectResourceId string `pulumi:"devCenterProjectResourceId"`
	// Defines the type of fabric the agent will run on.
	FabricProfile VmssFabricProfile `pulumi:"fabricProfile"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Defines how many resources can there be created at any given time.
	MaximumConcurrency int `pulumi:"maximumConcurrency"`
	// Defines the organization in which the pool will be used.
	OrganizationProfile interface{} `pulumi:"organizationProfile"`
	// Name of the pool. It needs to be globally unique.
	PoolName *string `pulumi:"poolName"`
	// The status of the current operation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// Defines how the machine will be handled once it executed a job.
	AgentProfile pulumi.Input
	// The resource id of the DevCenter Project the pool belongs to.
	DevCenterProjectResourceId pulumi.StringInput
	// Defines the type of fabric the agent will run on.
	FabricProfile VmssFabricProfileInput
	// The managed service identities assigned to this resource.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Defines how many resources can there be created at any given time.
	MaximumConcurrency pulumi.IntInput
	// Defines the organization in which the pool will be used.
	OrganizationProfile pulumi.Input
	// Name of the pool. It needs to be globally unique.
	PoolName pulumi.StringPtrInput
	// The status of the current operation.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

// Defines how the machine will be handled once it executed a job.
func (o PoolOutput) AgentProfile() pulumi.AnyOutput {
	return o.ApplyT(func(v *Pool) pulumi.AnyOutput { return v.AgentProfile }).(pulumi.AnyOutput)
}

// The Azure API version of the resource.
func (o PoolOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The resource id of the DevCenter Project the pool belongs to.
func (o PoolOutput) DevCenterProjectResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.DevCenterProjectResourceId }).(pulumi.StringOutput)
}

// Defines the type of fabric the agent will run on.
func (o PoolOutput) FabricProfile() VmssFabricProfileResponseOutput {
	return o.ApplyT(func(v *Pool) VmssFabricProfileResponseOutput { return v.FabricProfile }).(VmssFabricProfileResponseOutput)
}

// The managed service identities assigned to this resource.
func (o PoolOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Pool) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o PoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Defines how many resources can there be created at any given time.
func (o PoolOutput) MaximumConcurrency() pulumi.IntOutput {
	return o.ApplyT(func(v *Pool) pulumi.IntOutput { return v.MaximumConcurrency }).(pulumi.IntOutput)
}

// The name of the resource
func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the organization in which the pool will be used.
func (o PoolOutput) OrganizationProfile() pulumi.AnyOutput {
	return o.ApplyT(func(v *Pool) pulumi.AnyOutput { return v.OrganizationProfile }).(pulumi.AnyOutput)
}

// The status of the current operation.
func (o PoolOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Pool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
