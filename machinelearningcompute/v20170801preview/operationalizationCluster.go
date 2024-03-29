// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170801preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Instance of an Azure ML Operationalization Cluster resource.
type OperationalizationCluster struct {
	pulumi.CustomResourceState

	// AppInsights configuration.
	AppInsights AppInsightsPropertiesResponsePtrOutput `pulumi:"appInsights"`
	// The cluster type.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// Container Registry properties.
	ContainerRegistry ContainerRegistryPropertiesResponsePtrOutput `pulumi:"containerRegistry"`
	// Parameters for the Azure Container Service cluster.
	ContainerService AcsClusterPropertiesResponsePtrOutput `pulumi:"containerService"`
	// The date and time when the cluster was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// The description of the cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Contains global configuration for the web services in the cluster.
	GlobalServiceConfiguration GlobalServiceConfigurationResponsePtrOutput `pulumi:"globalServiceConfiguration"`
	// Specifies the location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The date and time when the cluster was last modified.
	ModifiedOn pulumi.StringOutput `pulumi:"modifiedOn"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of provisioning errors reported by the resource provider.
	ProvisioningErrors ErrorResponseWrapperResponseArrayOutput `pulumi:"provisioningErrors"`
	// The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Storage Account properties.
	StorageAccount StorageAccountPropertiesResponsePtrOutput `pulumi:"storageAccount"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOperationalizationCluster registers a new resource with the given unique name, arguments, and options.
func NewOperationalizationCluster(ctx *pulumi.Context,
	name string, args *OperationalizationClusterArgs, opts ...pulumi.ResourceOption) (*OperationalizationCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ContainerService != nil {
		args.ContainerService = args.ContainerService.ToAcsClusterPropertiesPtrOutput().ApplyT(func(v *AcsClusterProperties) *AcsClusterProperties { return v.Defaults() }).(AcsClusterPropertiesPtrOutput)
	}
	if args.GlobalServiceConfiguration != nil {
		args.GlobalServiceConfiguration = args.GlobalServiceConfiguration.ToGlobalServiceConfigurationPtrOutput().ApplyT(func(v *GlobalServiceConfiguration) *GlobalServiceConfiguration { return v.Defaults() }).(GlobalServiceConfigurationPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningcompute:OperationalizationCluster"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningcompute/v20170601preview:OperationalizationCluster"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OperationalizationCluster
	err := ctx.RegisterResource("azure-native:machinelearningcompute/v20170801preview:OperationalizationCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOperationalizationCluster gets an existing OperationalizationCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOperationalizationCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OperationalizationClusterState, opts ...pulumi.ResourceOption) (*OperationalizationCluster, error) {
	var resource OperationalizationCluster
	err := ctx.ReadResource("azure-native:machinelearningcompute/v20170801preview:OperationalizationCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OperationalizationCluster resources.
type operationalizationClusterState struct {
}

type OperationalizationClusterState struct {
}

func (OperationalizationClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*operationalizationClusterState)(nil)).Elem()
}

type operationalizationClusterArgs struct {
	// AppInsights configuration.
	AppInsights *AppInsightsProperties `pulumi:"appInsights"`
	// The name of the cluster.
	ClusterName *string `pulumi:"clusterName"`
	// The cluster type.
	ClusterType string `pulumi:"clusterType"`
	// Container Registry properties.
	ContainerRegistry *ContainerRegistryProperties `pulumi:"containerRegistry"`
	// Parameters for the Azure Container Service cluster.
	ContainerService *AcsClusterProperties `pulumi:"containerService"`
	// The description of the cluster.
	Description *string `pulumi:"description"`
	// Contains global configuration for the web services in the cluster.
	GlobalServiceConfiguration *GlobalServiceConfiguration `pulumi:"globalServiceConfiguration"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Name of the resource group in which the cluster is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Storage Account properties.
	StorageAccount *StorageAccountProperties `pulumi:"storageAccount"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a OperationalizationCluster resource.
type OperationalizationClusterArgs struct {
	// AppInsights configuration.
	AppInsights AppInsightsPropertiesPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringPtrInput
	// The cluster type.
	ClusterType pulumi.StringInput
	// Container Registry properties.
	ContainerRegistry ContainerRegistryPropertiesPtrInput
	// Parameters for the Azure Container Service cluster.
	ContainerService AcsClusterPropertiesPtrInput
	// The description of the cluster.
	Description pulumi.StringPtrInput
	// Contains global configuration for the web services in the cluster.
	GlobalServiceConfiguration GlobalServiceConfigurationPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Name of the resource group in which the cluster is located.
	ResourceGroupName pulumi.StringInput
	// Storage Account properties.
	StorageAccount StorageAccountPropertiesPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (OperationalizationClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*operationalizationClusterArgs)(nil)).Elem()
}

type OperationalizationClusterInput interface {
	pulumi.Input

	ToOperationalizationClusterOutput() OperationalizationClusterOutput
	ToOperationalizationClusterOutputWithContext(ctx context.Context) OperationalizationClusterOutput
}

func (*OperationalizationCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationalizationCluster)(nil)).Elem()
}

func (i *OperationalizationCluster) ToOperationalizationClusterOutput() OperationalizationClusterOutput {
	return i.ToOperationalizationClusterOutputWithContext(context.Background())
}

func (i *OperationalizationCluster) ToOperationalizationClusterOutputWithContext(ctx context.Context) OperationalizationClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationalizationClusterOutput)
}

type OperationalizationClusterOutput struct{ *pulumi.OutputState }

func (OperationalizationClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationalizationCluster)(nil)).Elem()
}

func (o OperationalizationClusterOutput) ToOperationalizationClusterOutput() OperationalizationClusterOutput {
	return o
}

func (o OperationalizationClusterOutput) ToOperationalizationClusterOutputWithContext(ctx context.Context) OperationalizationClusterOutput {
	return o
}

// AppInsights configuration.
func (o OperationalizationClusterOutput) AppInsights() AppInsightsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) AppInsightsPropertiesResponsePtrOutput { return v.AppInsights }).(AppInsightsPropertiesResponsePtrOutput)
}

// The cluster type.
func (o OperationalizationClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// Container Registry properties.
func (o OperationalizationClusterOutput) ContainerRegistry() ContainerRegistryPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) ContainerRegistryPropertiesResponsePtrOutput {
		return v.ContainerRegistry
	}).(ContainerRegistryPropertiesResponsePtrOutput)
}

// Parameters for the Azure Container Service cluster.
func (o OperationalizationClusterOutput) ContainerService() AcsClusterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) AcsClusterPropertiesResponsePtrOutput { return v.ContainerService }).(AcsClusterPropertiesResponsePtrOutput)
}

// The date and time when the cluster was created.
func (o OperationalizationClusterOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.CreatedOn }).(pulumi.StringOutput)
}

// The description of the cluster.
func (o OperationalizationClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Contains global configuration for the web services in the cluster.
func (o OperationalizationClusterOutput) GlobalServiceConfiguration() GlobalServiceConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) GlobalServiceConfigurationResponsePtrOutput {
		return v.GlobalServiceConfiguration
	}).(GlobalServiceConfigurationResponsePtrOutput)
}

// Specifies the location of the resource.
func (o OperationalizationClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The date and time when the cluster was last modified.
func (o OperationalizationClusterOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.ModifiedOn }).(pulumi.StringOutput)
}

// Specifies the name of the resource.
func (o OperationalizationClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of provisioning errors reported by the resource provider.
func (o OperationalizationClusterOutput) ProvisioningErrors() ErrorResponseWrapperResponseArrayOutput {
	return o.ApplyT(func(v *OperationalizationCluster) ErrorResponseWrapperResponseArrayOutput {
		return v.ProvisioningErrors
	}).(ErrorResponseWrapperResponseArrayOutput)
}

// The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed.
func (o OperationalizationClusterOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Storage Account properties.
func (o OperationalizationClusterOutput) StorageAccount() StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *OperationalizationCluster) StorageAccountPropertiesResponsePtrOutput { return v.StorageAccount }).(StorageAccountPropertiesResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o OperationalizationClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o OperationalizationClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OperationalizationCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OperationalizationClusterOutput{})
}
