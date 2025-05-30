// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The cluster.
//
// Uses Azure REST API version 2024-05-01-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hdinsight [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ClusterPoolCluster struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Cluster profile.
	ClusterProfile ClusterProfileResponseOutput `pulumi:"clusterProfile"`
	// The type of cluster.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// The compute profile.
	ComputeProfile ClusterPoolComputeProfileResponseOutput `pulumi:"computeProfile"`
	// A unique id generated by the RP to identify the resource.
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Business status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewClusterPoolCluster registers a new resource with the given unique name, arguments, and options.
func NewClusterPoolCluster(ctx *pulumi.Context,
	name string, args *ClusterPoolClusterArgs, opts ...pulumi.ResourceOption) (*ClusterPoolCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterPoolName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterPoolName'")
	}
	if args.ClusterProfile == nil {
		return nil, errors.New("invalid value for required argument 'ClusterProfile'")
	}
	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.ComputeProfile == nil {
		return nil, errors.New("invalid value for required argument 'ComputeProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ClusterProfile = args.ClusterProfile.ToClusterProfileOutput().ApplyT(func(v ClusterProfile) ClusterProfile { return *v.Defaults() }).(ClusterProfileOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hdinsight/v20230601preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20230601preview:ClusterPoolCluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20231101preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20231101preview:ClusterPoolCluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20240501preview:Cluster"),
		},
		{
			Type: pulumi.String("azure-native:hdinsight/v20240501preview:ClusterPoolCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ClusterPoolCluster
	err := ctx.RegisterResource("azure-native:hdinsight:ClusterPoolCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPoolCluster gets an existing ClusterPoolCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPoolCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPoolClusterState, opts ...pulumi.ResourceOption) (*ClusterPoolCluster, error) {
	var resource ClusterPoolCluster
	err := ctx.ReadResource("azure-native:hdinsight:ClusterPoolCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPoolCluster resources.
type clusterPoolClusterState struct {
}

type ClusterPoolClusterState struct {
}

func (ClusterPoolClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPoolClusterState)(nil)).Elem()
}

type clusterPoolClusterArgs struct {
	// The name of the HDInsight cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The name of the cluster pool.
	ClusterPoolName string `pulumi:"clusterPoolName"`
	// Cluster profile.
	ClusterProfile ClusterProfile `pulumi:"clusterProfile"`
	// The type of cluster.
	ClusterType string `pulumi:"clusterType"`
	// The compute profile.
	ComputeProfile ClusterPoolComputeProfile `pulumi:"computeProfile"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterPoolCluster resource.
type ClusterPoolClusterArgs struct {
	// The name of the HDInsight cluster.
	ClusterName pulumi.StringPtrInput
	// The name of the cluster pool.
	ClusterPoolName pulumi.StringInput
	// Cluster profile.
	ClusterProfile ClusterProfileInput
	// The type of cluster.
	ClusterType pulumi.StringInput
	// The compute profile.
	ComputeProfile ClusterPoolComputeProfileInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ClusterPoolClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPoolClusterArgs)(nil)).Elem()
}

type ClusterPoolClusterInput interface {
	pulumi.Input

	ToClusterPoolClusterOutput() ClusterPoolClusterOutput
	ToClusterPoolClusterOutputWithContext(ctx context.Context) ClusterPoolClusterOutput
}

func (*ClusterPoolCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPoolCluster)(nil)).Elem()
}

func (i *ClusterPoolCluster) ToClusterPoolClusterOutput() ClusterPoolClusterOutput {
	return i.ToClusterPoolClusterOutputWithContext(context.Background())
}

func (i *ClusterPoolCluster) ToClusterPoolClusterOutputWithContext(ctx context.Context) ClusterPoolClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPoolClusterOutput)
}

type ClusterPoolClusterOutput struct{ *pulumi.OutputState }

func (ClusterPoolClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPoolCluster)(nil)).Elem()
}

func (o ClusterPoolClusterOutput) ToClusterPoolClusterOutput() ClusterPoolClusterOutput {
	return o
}

func (o ClusterPoolClusterOutput) ToClusterPoolClusterOutputWithContext(ctx context.Context) ClusterPoolClusterOutput {
	return o
}

// The Azure API version of the resource.
func (o ClusterPoolClusterOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Cluster profile.
func (o ClusterPoolClusterOutput) ClusterProfile() ClusterProfileResponseOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) ClusterProfileResponseOutput { return v.ClusterProfile }).(ClusterProfileResponseOutput)
}

// The type of cluster.
func (o ClusterPoolClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// The compute profile.
func (o ClusterPoolClusterOutput) ComputeProfile() ClusterPoolComputeProfileResponseOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) ClusterPoolComputeProfileResponseOutput { return v.ComputeProfile }).(ClusterPoolComputeProfileResponseOutput)
}

// A unique id generated by the RP to identify the resource.
func (o ClusterPoolClusterOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ClusterPoolClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ClusterPoolClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o ClusterPoolClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Business status of the resource.
func (o ClusterPoolClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ClusterPoolClusterOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ClusterPoolClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ClusterPoolClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPoolCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterPoolClusterOutput{})
}
