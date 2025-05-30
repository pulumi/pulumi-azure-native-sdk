// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkcloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2025-02-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.
//
// Other available API versions: 2024-07-01, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native networkcloud [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type AgentPool struct {
	pulumi.CustomResourceState

	// The administrator credentials to be used for the nodes in this agent pool.
	AdministratorConfiguration AdministratorConfigurationResponsePtrOutput `pulumi:"administratorConfiguration"`
	// The configurations that will be applied to each agent in this agent pool.
	AgentOptions AgentOptionsResponsePtrOutput `pulumi:"agentOptions"`
	// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
	AttachedNetworkConfiguration AttachedNetworkConfigurationResponsePtrOutput `pulumi:"attachedNetworkConfiguration"`
	// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The number of virtual machines that use this configuration.
	Count pulumi.Float64Output `pulumi:"count"`
	// The current status of the agent pool.
	DetailedStatus pulumi.StringOutput `pulumi:"detailedStatus"`
	// The descriptive message about the current detailed status.
	DetailedStatusMessage pulumi.StringOutput `pulumi:"detailedStatusMessage"`
	// Resource ETag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The Kubernetes version running in this agent pool.
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The labels applied to the nodes in this agent pool.
	Labels KubernetesLabelResponseArrayOutput `pulumi:"labels"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the agent pool.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The taints applied to the nodes in this agent pool.
	Taints KubernetesLabelResponseArrayOutput `pulumi:"taints"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The configuration of the agent pool.
	UpgradeSettings AgentPoolUpgradeSettingsResponsePtrOutput `pulumi:"upgradeSettings"`
	// The name of the VM SKU that determines the size of resources allocated for node VMs.
	VmSkuName pulumi.StringOutput `pulumi:"vmSkuName"`
}

// NewAgentPool registers a new resource with the given unique name, arguments, and options.
func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Count == nil {
		return nil, errors.New("invalid value for required argument 'Count'")
	}
	if args.KubernetesClusterName == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesClusterName'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmSkuName == nil {
		return nil, errors.New("invalid value for required argument 'VmSkuName'")
	}
	if args.AgentOptions != nil {
		args.AgentOptions = args.AgentOptions.ToAgentOptionsPtrOutput().ApplyT(func(v *AgentOptions) *AgentOptions { return v.Defaults() }).(AgentOptionsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:networkcloud/v20230701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20231001preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20240601preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20240701:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20241001preview:AgentPool"),
		},
		{
			Type: pulumi.String("azure-native:networkcloud/v20250201:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AgentPool
	err := ctx.RegisterResource("azure-native:networkcloud:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentPool gets an existing AgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azure-native:networkcloud:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentPool resources.
type agentPoolState struct {
}

type AgentPoolState struct {
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	// The administrator credentials to be used for the nodes in this agent pool.
	AdministratorConfiguration *AdministratorConfiguration `pulumi:"administratorConfiguration"`
	// The configurations that will be applied to each agent in this agent pool.
	AgentOptions *AgentOptions `pulumi:"agentOptions"`
	// The name of the Kubernetes cluster agent pool.
	AgentPoolName *string `pulumi:"agentPoolName"`
	// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
	AttachedNetworkConfiguration *AttachedNetworkConfiguration `pulumi:"attachedNetworkConfiguration"`
	// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The number of virtual machines that use this configuration.
	Count float64 `pulumi:"count"`
	// The extended location of the cluster associated with the resource.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The name of the Kubernetes cluster.
	KubernetesClusterName string `pulumi:"kubernetesClusterName"`
	// The labels applied to the nodes in this agent pool.
	Labels []KubernetesLabel `pulumi:"labels"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
	Mode string `pulumi:"mode"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The taints applied to the nodes in this agent pool.
	Taints []KubernetesLabel `pulumi:"taints"`
	// The configuration of the agent pool.
	UpgradeSettings *AgentPoolUpgradeSettings `pulumi:"upgradeSettings"`
	// The name of the VM SKU that determines the size of resources allocated for node VMs.
	VmSkuName string `pulumi:"vmSkuName"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// The administrator credentials to be used for the nodes in this agent pool.
	AdministratorConfiguration AdministratorConfigurationPtrInput
	// The configurations that will be applied to each agent in this agent pool.
	AgentOptions AgentOptionsPtrInput
	// The name of the Kubernetes cluster agent pool.
	AgentPoolName pulumi.StringPtrInput
	// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
	AttachedNetworkConfiguration AttachedNetworkConfigurationPtrInput
	// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
	AvailabilityZones pulumi.StringArrayInput
	// The number of virtual machines that use this configuration.
	Count pulumi.Float64Input
	// The extended location of the cluster associated with the resource.
	ExtendedLocation ExtendedLocationPtrInput
	// The name of the Kubernetes cluster.
	KubernetesClusterName pulumi.StringInput
	// The labels applied to the nodes in this agent pool.
	Labels KubernetesLabelArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
	Mode pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The taints applied to the nodes in this agent pool.
	Taints KubernetesLabelArrayInput
	// The configuration of the agent pool.
	UpgradeSettings AgentPoolUpgradeSettingsPtrInput
	// The name of the VM SKU that determines the size of resources allocated for node VMs.
	VmSkuName pulumi.StringInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}

type AgentPoolInput interface {
	pulumi.Input

	ToAgentPoolOutput() AgentPoolOutput
	ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput
}

func (*AgentPool) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (i *AgentPool) ToAgentPoolOutput() AgentPoolOutput {
	return i.ToAgentPoolOutputWithContext(context.Background())
}

func (i *AgentPool) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgentPoolOutput)
}

type AgentPoolOutput struct{ *pulumi.OutputState }

func (AgentPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgentPool)(nil)).Elem()
}

func (o AgentPoolOutput) ToAgentPoolOutput() AgentPoolOutput {
	return o
}

func (o AgentPoolOutput) ToAgentPoolOutputWithContext(ctx context.Context) AgentPoolOutput {
	return o
}

// The administrator credentials to be used for the nodes in this agent pool.
func (o AgentPoolOutput) AdministratorConfiguration() AdministratorConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AdministratorConfigurationResponsePtrOutput { return v.AdministratorConfiguration }).(AdministratorConfigurationResponsePtrOutput)
}

// The configurations that will be applied to each agent in this agent pool.
func (o AgentPoolOutput) AgentOptions() AgentOptionsResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentOptionsResponsePtrOutput { return v.AgentOptions }).(AgentOptionsResponsePtrOutput)
}

// The configuration of networks being attached to the agent pool for use by the workloads that run on this Kubernetes cluster.
func (o AgentPoolOutput) AttachedNetworkConfiguration() AttachedNetworkConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AttachedNetworkConfigurationResponsePtrOutput {
		return v.AttachedNetworkConfiguration
	}).(AttachedNetworkConfigurationResponsePtrOutput)
}

// The list of availability zones of the Network Cloud cluster used for the provisioning of nodes in this agent pool. If not specified, all availability zones will be used.
func (o AgentPoolOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringArrayOutput { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o AgentPoolOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The number of virtual machines that use this configuration.
func (o AgentPoolOutput) Count() pulumi.Float64Output {
	return o.ApplyT(func(v *AgentPool) pulumi.Float64Output { return v.Count }).(pulumi.Float64Output)
}

// The current status of the agent pool.
func (o AgentPoolOutput) DetailedStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.DetailedStatus }).(pulumi.StringOutput)
}

// The descriptive message about the current detailed status.
func (o AgentPoolOutput) DetailedStatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.DetailedStatusMessage }).(pulumi.StringOutput)
}

// Resource ETag.
func (o AgentPoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The extended location of the cluster associated with the resource.
func (o AgentPoolOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The Kubernetes version running in this agent pool.
func (o AgentPoolOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The labels applied to the nodes in this agent pool.
func (o AgentPoolOutput) Labels() KubernetesLabelResponseArrayOutput {
	return o.ApplyT(func(v *AgentPool) KubernetesLabelResponseArrayOutput { return v.Labels }).(KubernetesLabelResponseArrayOutput)
}

// The geo-location where the resource lives
func (o AgentPoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The selection of how this agent pool is utilized, either as a system pool or a user pool. System pools run the features and critical services for the Kubernetes Cluster, while user pools are dedicated to user workloads. Every Kubernetes cluster must contain at least one system node pool with at least one node.
func (o AgentPoolOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// The name of the resource
func (o AgentPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the agent pool.
func (o AgentPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AgentPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AgentPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o AgentPoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The taints applied to the nodes in this agent pool.
func (o AgentPoolOutput) Taints() KubernetesLabelResponseArrayOutput {
	return o.ApplyT(func(v *AgentPool) KubernetesLabelResponseArrayOutput { return v.Taints }).(KubernetesLabelResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AgentPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The configuration of the agent pool.
func (o AgentPoolOutput) UpgradeSettings() AgentPoolUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *AgentPool) AgentPoolUpgradeSettingsResponsePtrOutput { return v.UpgradeSettings }).(AgentPoolUpgradeSettingsResponsePtrOutput)
}

// The name of the VM SKU that determines the size of resources allocated for node VMs.
func (o AgentPoolOutput) VmSkuName() pulumi.StringOutput {
	return o.ApplyT(func(v *AgentPool) pulumi.StringOutput { return v.VmSkuName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPoolOutput{})
}
