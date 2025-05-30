// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration Assignment
//
// Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
//
// Other available API versions: 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ConfigurationAssignmentsForResourceGroup struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Properties of the configuration assignment
	Filter ConfigurationAssignmentFilterPropertiesResponsePtrOutput `pulumi:"filter"`
	// Location of the resource
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrOutput `pulumi:"maintenanceConfigurationId"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique resourceId
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationAssignmentsForResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAssignmentsForResourceGroup(ctx *pulumi.Context,
	name string, args *ConfigurationAssignmentsForResourceGroupArgs, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentsForResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maintenance/v20230401:ConfigurationAssignmentsForResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20230901preview:ConfigurationAssignmentsForResourceGroup"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20231001preview:ConfigurationAssignmentsForResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationAssignmentsForResourceGroup
	err := ctx.RegisterResource("azure-native:maintenance:ConfigurationAssignmentsForResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAssignmentsForResourceGroup gets an existing ConfigurationAssignmentsForResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAssignmentsForResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAssignmentsForResourceGroupState, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentsForResourceGroup, error) {
	var resource ConfigurationAssignmentsForResourceGroup
	err := ctx.ReadResource("azure-native:maintenance:ConfigurationAssignmentsForResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAssignmentsForResourceGroup resources.
type configurationAssignmentsForResourceGroupState struct {
}

type ConfigurationAssignmentsForResourceGroupState struct {
}

func (ConfigurationAssignmentsForResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentsForResourceGroupState)(nil)).Elem()
}

type configurationAssignmentsForResourceGroupArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName *string `pulumi:"configurationAssignmentName"`
	// Properties of the configuration assignment
	Filter *ConfigurationAssignmentFilterProperties `pulumi:"filter"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// Resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The unique resourceId
	ResourceId *string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ConfigurationAssignmentsForResourceGroup resource.
type ConfigurationAssignmentsForResourceGroupArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringPtrInput
	// Properties of the configuration assignment
	Filter ConfigurationAssignmentFilterPropertiesPtrInput
	// Location of the resource
	Location pulumi.StringPtrInput
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrInput
	// Resource group name
	ResourceGroupName pulumi.StringInput
	// The unique resourceId
	ResourceId pulumi.StringPtrInput
}

func (ConfigurationAssignmentsForResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentsForResourceGroupArgs)(nil)).Elem()
}

type ConfigurationAssignmentsForResourceGroupInput interface {
	pulumi.Input

	ToConfigurationAssignmentsForResourceGroupOutput() ConfigurationAssignmentsForResourceGroupOutput
	ToConfigurationAssignmentsForResourceGroupOutputWithContext(ctx context.Context) ConfigurationAssignmentsForResourceGroupOutput
}

func (*ConfigurationAssignmentsForResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignmentsForResourceGroup)(nil)).Elem()
}

func (i *ConfigurationAssignmentsForResourceGroup) ToConfigurationAssignmentsForResourceGroupOutput() ConfigurationAssignmentsForResourceGroupOutput {
	return i.ToConfigurationAssignmentsForResourceGroupOutputWithContext(context.Background())
}

func (i *ConfigurationAssignmentsForResourceGroup) ToConfigurationAssignmentsForResourceGroupOutputWithContext(ctx context.Context) ConfigurationAssignmentsForResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAssignmentsForResourceGroupOutput)
}

type ConfigurationAssignmentsForResourceGroupOutput struct{ *pulumi.OutputState }

func (ConfigurationAssignmentsForResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignmentsForResourceGroup)(nil)).Elem()
}

func (o ConfigurationAssignmentsForResourceGroupOutput) ToConfigurationAssignmentsForResourceGroupOutput() ConfigurationAssignmentsForResourceGroupOutput {
	return o
}

func (o ConfigurationAssignmentsForResourceGroupOutput) ToConfigurationAssignmentsForResourceGroupOutputWithContext(ctx context.Context) ConfigurationAssignmentsForResourceGroupOutput {
	return o
}

// The Azure API version of the resource.
func (o ConfigurationAssignmentsForResourceGroupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Properties of the configuration assignment
func (o ConfigurationAssignmentsForResourceGroupOutput) Filter() ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
		return v.Filter
	}).(ConfigurationAssignmentFilterPropertiesResponsePtrOutput)
}

// Location of the resource
func (o ConfigurationAssignmentsForResourceGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The maintenance configuration Id
func (o ConfigurationAssignmentsForResourceGroupOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) pulumi.StringPtrOutput {
		return v.MaintenanceConfigurationId
	}).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o ConfigurationAssignmentsForResourceGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique resourceId
func (o ConfigurationAssignmentsForResourceGroupOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationAssignmentsForResourceGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource
func (o ConfigurationAssignmentsForResourceGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForResourceGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAssignmentsForResourceGroupOutput{})
}
