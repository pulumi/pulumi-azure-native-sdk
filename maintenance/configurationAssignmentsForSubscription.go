// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package maintenance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration Assignment
//
// Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
//
// Other available API versions: 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ConfigurationAssignmentsForSubscription struct {
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

// NewConfigurationAssignmentsForSubscription registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAssignmentsForSubscription(ctx *pulumi.Context,
	name string, args *ConfigurationAssignmentsForSubscriptionArgs, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentsForSubscription, error) {
	if args == nil {
		args = &ConfigurationAssignmentsForSubscriptionArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:maintenance/v20230401:ConfigurationAssignmentsForSubscription"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20230901preview:ConfigurationAssignmentsForSubscription"),
		},
		{
			Type: pulumi.String("azure-native:maintenance/v20231001preview:ConfigurationAssignmentsForSubscription"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationAssignmentsForSubscription
	err := ctx.RegisterResource("azure-native:maintenance:ConfigurationAssignmentsForSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAssignmentsForSubscription gets an existing ConfigurationAssignmentsForSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAssignmentsForSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAssignmentsForSubscriptionState, opts ...pulumi.ResourceOption) (*ConfigurationAssignmentsForSubscription, error) {
	var resource ConfigurationAssignmentsForSubscription
	err := ctx.ReadResource("azure-native:maintenance:ConfigurationAssignmentsForSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAssignmentsForSubscription resources.
type configurationAssignmentsForSubscriptionState struct {
}

type ConfigurationAssignmentsForSubscriptionState struct {
}

func (ConfigurationAssignmentsForSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentsForSubscriptionState)(nil)).Elem()
}

type configurationAssignmentsForSubscriptionArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName *string `pulumi:"configurationAssignmentName"`
	// Properties of the configuration assignment
	Filter *ConfigurationAssignmentFilterProperties `pulumi:"filter"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// The maintenance configuration Id
	MaintenanceConfigurationId *string `pulumi:"maintenanceConfigurationId"`
	// The unique resourceId
	ResourceId *string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ConfigurationAssignmentsForSubscription resource.
type ConfigurationAssignmentsForSubscriptionArgs struct {
	// Configuration assignment name
	ConfigurationAssignmentName pulumi.StringPtrInput
	// Properties of the configuration assignment
	Filter ConfigurationAssignmentFilterPropertiesPtrInput
	// Location of the resource
	Location pulumi.StringPtrInput
	// The maintenance configuration Id
	MaintenanceConfigurationId pulumi.StringPtrInput
	// The unique resourceId
	ResourceId pulumi.StringPtrInput
}

func (ConfigurationAssignmentsForSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAssignmentsForSubscriptionArgs)(nil)).Elem()
}

type ConfigurationAssignmentsForSubscriptionInput interface {
	pulumi.Input

	ToConfigurationAssignmentsForSubscriptionOutput() ConfigurationAssignmentsForSubscriptionOutput
	ToConfigurationAssignmentsForSubscriptionOutputWithContext(ctx context.Context) ConfigurationAssignmentsForSubscriptionOutput
}

func (*ConfigurationAssignmentsForSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignmentsForSubscription)(nil)).Elem()
}

func (i *ConfigurationAssignmentsForSubscription) ToConfigurationAssignmentsForSubscriptionOutput() ConfigurationAssignmentsForSubscriptionOutput {
	return i.ToConfigurationAssignmentsForSubscriptionOutputWithContext(context.Background())
}

func (i *ConfigurationAssignmentsForSubscription) ToConfigurationAssignmentsForSubscriptionOutputWithContext(ctx context.Context) ConfigurationAssignmentsForSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAssignmentsForSubscriptionOutput)
}

type ConfigurationAssignmentsForSubscriptionOutput struct{ *pulumi.OutputState }

func (ConfigurationAssignmentsForSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationAssignmentsForSubscription)(nil)).Elem()
}

func (o ConfigurationAssignmentsForSubscriptionOutput) ToConfigurationAssignmentsForSubscriptionOutput() ConfigurationAssignmentsForSubscriptionOutput {
	return o
}

func (o ConfigurationAssignmentsForSubscriptionOutput) ToConfigurationAssignmentsForSubscriptionOutputWithContext(ctx context.Context) ConfigurationAssignmentsForSubscriptionOutput {
	return o
}

// The Azure API version of the resource.
func (o ConfigurationAssignmentsForSubscriptionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Properties of the configuration assignment
func (o ConfigurationAssignmentsForSubscriptionOutput) Filter() ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) ConfigurationAssignmentFilterPropertiesResponsePtrOutput {
		return v.Filter
	}).(ConfigurationAssignmentFilterPropertiesResponsePtrOutput)
}

// Location of the resource
func (o ConfigurationAssignmentsForSubscriptionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The maintenance configuration Id
func (o ConfigurationAssignmentsForSubscriptionOutput) MaintenanceConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) pulumi.StringPtrOutput {
		return v.MaintenanceConfigurationId
	}).(pulumi.StringPtrOutput)
}

// Name of the resource
func (o ConfigurationAssignmentsForSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The unique resourceId
func (o ConfigurationAssignmentsForSubscriptionOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) pulumi.StringPtrOutput { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationAssignmentsForSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the resource
func (o ConfigurationAssignmentsForSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationAssignmentsForSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAssignmentsForSubscriptionOutput{})
}
