// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Microsoft.AwsConnector resource
//
// Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2024-12-01.
type ElasticLoadBalancingV2LoadBalancer struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties ElasticLoadBalancingV2LoadBalancerPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewElasticLoadBalancingV2LoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewElasticLoadBalancingV2LoadBalancer(ctx *pulumi.Context,
	name string, args *ElasticLoadBalancingV2LoadBalancerArgs, opts ...pulumi.ResourceOption) (*ElasticLoadBalancingV2LoadBalancer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:ElasticLoadBalancingV2LoadBalancer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ElasticLoadBalancingV2LoadBalancer
	err := ctx.RegisterResource("azure-native:awsconnector:ElasticLoadBalancingV2LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElasticLoadBalancingV2LoadBalancer gets an existing ElasticLoadBalancingV2LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElasticLoadBalancingV2LoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElasticLoadBalancingV2LoadBalancerState, opts ...pulumi.ResourceOption) (*ElasticLoadBalancingV2LoadBalancer, error) {
	var resource ElasticLoadBalancingV2LoadBalancer
	err := ctx.ReadResource("azure-native:awsconnector:ElasticLoadBalancingV2LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElasticLoadBalancingV2LoadBalancer resources.
type elasticLoadBalancingV2LoadBalancerState struct {
}

type ElasticLoadBalancingV2LoadBalancerState struct {
}

func (ElasticLoadBalancingV2LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticLoadBalancingV2LoadBalancerState)(nil)).Elem()
}

type elasticLoadBalancingV2LoadBalancerArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of ElasticLoadBalancingV2LoadBalancer
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *ElasticLoadBalancingV2LoadBalancerProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ElasticLoadBalancingV2LoadBalancer resource.
type ElasticLoadBalancingV2LoadBalancerArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of ElasticLoadBalancingV2LoadBalancer
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties ElasticLoadBalancingV2LoadBalancerPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ElasticLoadBalancingV2LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elasticLoadBalancingV2LoadBalancerArgs)(nil)).Elem()
}

type ElasticLoadBalancingV2LoadBalancerInput interface {
	pulumi.Input

	ToElasticLoadBalancingV2LoadBalancerOutput() ElasticLoadBalancingV2LoadBalancerOutput
	ToElasticLoadBalancingV2LoadBalancerOutputWithContext(ctx context.Context) ElasticLoadBalancingV2LoadBalancerOutput
}

func (*ElasticLoadBalancingV2LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticLoadBalancingV2LoadBalancer)(nil)).Elem()
}

func (i *ElasticLoadBalancingV2LoadBalancer) ToElasticLoadBalancingV2LoadBalancerOutput() ElasticLoadBalancingV2LoadBalancerOutput {
	return i.ToElasticLoadBalancingV2LoadBalancerOutputWithContext(context.Background())
}

func (i *ElasticLoadBalancingV2LoadBalancer) ToElasticLoadBalancingV2LoadBalancerOutputWithContext(ctx context.Context) ElasticLoadBalancingV2LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElasticLoadBalancingV2LoadBalancerOutput)
}

type ElasticLoadBalancingV2LoadBalancerOutput struct{ *pulumi.OutputState }

func (ElasticLoadBalancingV2LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElasticLoadBalancingV2LoadBalancer)(nil)).Elem()
}

func (o ElasticLoadBalancingV2LoadBalancerOutput) ToElasticLoadBalancingV2LoadBalancerOutput() ElasticLoadBalancingV2LoadBalancerOutput {
	return o
}

func (o ElasticLoadBalancingV2LoadBalancerOutput) ToElasticLoadBalancingV2LoadBalancerOutputWithContext(ctx context.Context) ElasticLoadBalancingV2LoadBalancerOutput {
	return o
}

// The Azure API version of the resource.
func (o ElasticLoadBalancingV2LoadBalancerOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o ElasticLoadBalancingV2LoadBalancerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ElasticLoadBalancingV2LoadBalancerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o ElasticLoadBalancingV2LoadBalancerOutput) Properties() ElasticLoadBalancingV2LoadBalancerPropertiesResponseOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) ElasticLoadBalancingV2LoadBalancerPropertiesResponseOutput {
		return v.Properties
	}).(ElasticLoadBalancingV2LoadBalancerPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ElasticLoadBalancingV2LoadBalancerOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ElasticLoadBalancingV2LoadBalancerOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ElasticLoadBalancingV2LoadBalancerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ElasticLoadBalancingV2LoadBalancer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ElasticLoadBalancingV2LoadBalancerOutput{})
}
