// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automanage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration profile assignment is an association between a VM and automanage profile configuration.
//
// Uses Azure REST API version 2022-05-04. In version 2.x of the Azure Native provider, it used API version 2022-05-04.
//
// Other available API versions: 2021-04-30-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automanage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ConfigurationProfileHCRPAssignment struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfileHCRPAssignment registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfileHCRPAssignment(ctx *pulumi.Context,
	name string, args *ConfigurationProfileHCRPAssignmentArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCRPAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automanage/v20210430preview:ConfigurationProfileHCRPAssignment"),
		},
		{
			Type: pulumi.String("azure-native:automanage/v20220504:ConfigurationProfileHCRPAssignment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConfigurationProfileHCRPAssignment
	err := ctx.RegisterResource("azure-native:automanage:ConfigurationProfileHCRPAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfileHCRPAssignment gets an existing ConfigurationProfileHCRPAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfileHCRPAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfileHCRPAssignmentState, opts ...pulumi.ResourceOption) (*ConfigurationProfileHCRPAssignment, error) {
	var resource ConfigurationProfileHCRPAssignment
	err := ctx.ReadResource("azure-native:automanage:ConfigurationProfileHCRPAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfileHCRPAssignment resources.
type configurationProfileHCRPAssignmentState struct {
}

type ConfigurationProfileHCRPAssignmentState struct {
}

func (ConfigurationProfileHCRPAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCRPAssignmentState)(nil)).Elem()
}

type configurationProfileHCRPAssignmentArgs struct {
	// Name of the configuration profile assignment. Only default is supported.
	ConfigurationProfileAssignmentName *string `pulumi:"configurationProfileAssignmentName"`
	// The name of the Arc machine.
	MachineName string `pulumi:"machineName"`
	// Properties of the configuration profile assignment.
	Properties *ConfigurationProfileAssignmentProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ConfigurationProfileHCRPAssignment resource.
type ConfigurationProfileHCRPAssignmentArgs struct {
	// Name of the configuration profile assignment. Only default is supported.
	ConfigurationProfileAssignmentName pulumi.StringPtrInput
	// The name of the Arc machine.
	MachineName pulumi.StringInput
	// Properties of the configuration profile assignment.
	Properties ConfigurationProfileAssignmentPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ConfigurationProfileHCRPAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfileHCRPAssignmentArgs)(nil)).Elem()
}

type ConfigurationProfileHCRPAssignmentInput interface {
	pulumi.Input

	ToConfigurationProfileHCRPAssignmentOutput() ConfigurationProfileHCRPAssignmentOutput
	ToConfigurationProfileHCRPAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCRPAssignmentOutput
}

func (*ConfigurationProfileHCRPAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCRPAssignment)(nil)).Elem()
}

func (i *ConfigurationProfileHCRPAssignment) ToConfigurationProfileHCRPAssignmentOutput() ConfigurationProfileHCRPAssignmentOutput {
	return i.ToConfigurationProfileHCRPAssignmentOutputWithContext(context.Background())
}

func (i *ConfigurationProfileHCRPAssignment) ToConfigurationProfileHCRPAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCRPAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileHCRPAssignmentOutput)
}

type ConfigurationProfileHCRPAssignmentOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileHCRPAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationProfileHCRPAssignment)(nil)).Elem()
}

func (o ConfigurationProfileHCRPAssignmentOutput) ToConfigurationProfileHCRPAssignmentOutput() ConfigurationProfileHCRPAssignmentOutput {
	return o
}

func (o ConfigurationProfileHCRPAssignmentOutput) ToConfigurationProfileHCRPAssignmentOutputWithContext(ctx context.Context) ConfigurationProfileHCRPAssignmentOutput {
	return o
}

// The Azure API version of the resource.
func (o ConfigurationProfileHCRPAssignmentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Azure resource id. Indicates if this resource is managed by another Azure resource.
func (o ConfigurationProfileHCRPAssignmentOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConfigurationProfileHCRPAssignmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the configuration profile assignment.
func (o ConfigurationProfileHCRPAssignmentOutput) Properties() ConfigurationProfileAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) ConfigurationProfileAssignmentPropertiesResponseOutput {
		return v.Properties
	}).(ConfigurationProfileAssignmentPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConfigurationProfileHCRPAssignmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConfigurationProfileHCRPAssignmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationProfileHCRPAssignment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationProfileHCRPAssignmentOutput{})
}
