// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the Virtual Instance for SAP solutions resource.
type SAPVirtualInstance struct {
	pulumi.CustomResourceState

	// Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
	Configuration pulumi.AnyOutput `pulumi:"configuration"`
	// Defines the environment type - Production/Non Production.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Indicates any errors on the Virtual Instance for SAP solutions resource.
	Errors SAPVirtualInstanceErrorResponseOutput `pulumi:"errors"`
	// Defines the health of SAP Instances.
	Health pulumi.StringOutput `pulumi:"health"`
	// Managed service identity (user assigned identities)
	Identity UserAssignedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Managed resource group configuration
	ManagedResourceGroupConfiguration ManagedRGConfigurationResponsePtrOutput `pulumi:"managedResourceGroupConfiguration"`
	// Specifies the network access configuration for the resources that will be deployed in the Managed Resource Group. The options to choose from are Public and Private. If 'Private' is chosen, the Storage Account service tag should be enabled on the subnets in which the SAP VMs exist. This is required for establishing connectivity between VM extensions and the managed resource group storage account. This setting is currently applicable only to Storage Account. Learn more here https://go.microsoft.com/fwlink/?linkid=2247228
	ManagedResourcesNetworkAccessType pulumi.StringPtrOutput `pulumi:"managedResourcesNetworkAccessType"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines the SAP Product type.
	SapProduct pulumi.StringOutput `pulumi:"sapProduct"`
	// Defines the Virtual Instance for SAP state.
	State pulumi.StringOutput `pulumi:"state"`
	// Defines the SAP Instance status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSAPVirtualInstance registers a new resource with the given unique name, arguments, and options.
func NewSAPVirtualInstance(ctx *pulumi.Context,
	name string, args *SAPVirtualInstanceArgs, opts ...pulumi.ResourceOption) (*SAPVirtualInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapProduct == nil {
		return nil, errors.New("invalid value for required argument 'SapProduct'")
	}
	if args.ManagedResourcesNetworkAccessType == nil {
		args.ManagedResourcesNetworkAccessType = pulumi.StringPtr("Public")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:SAPVirtualInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20221101preview:SAPVirtualInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20230401:SAPVirtualInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20240901:SAPVirtualInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads:SAPVirtualInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SAPVirtualInstance
	err := ctx.RegisterResource("azure-native:workloads/v20231001preview:SAPVirtualInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSAPVirtualInstance gets an existing SAPVirtualInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSAPVirtualInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAPVirtualInstanceState, opts ...pulumi.ResourceOption) (*SAPVirtualInstance, error) {
	var resource SAPVirtualInstance
	err := ctx.ReadResource("azure-native:workloads/v20231001preview:SAPVirtualInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SAPVirtualInstance resources.
type sapvirtualInstanceState struct {
}

type SAPVirtualInstanceState struct {
}

func (SAPVirtualInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapvirtualInstanceState)(nil)).Elem()
}

type sapvirtualInstanceArgs struct {
	// Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
	Configuration interface{} `pulumi:"configuration"`
	// Defines the environment type - Production/Non Production.
	Environment string `pulumi:"environment"`
	// Managed service identity (user assigned identities)
	Identity *UserAssignedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Managed resource group configuration
	ManagedResourceGroupConfiguration *ManagedRGConfiguration `pulumi:"managedResourceGroupConfiguration"`
	// Specifies the network access configuration for the resources that will be deployed in the Managed Resource Group. The options to choose from are Public and Private. If 'Private' is chosen, the Storage Account service tag should be enabled on the subnets in which the SAP VMs exist. This is required for establishing connectivity between VM extensions and the managed resource group storage account. This setting is currently applicable only to Storage Account. Learn more here https://go.microsoft.com/fwlink/?linkid=2247228
	ManagedResourcesNetworkAccessType *string `pulumi:"managedResourcesNetworkAccessType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Defines the SAP Product type.
	SapProduct string `pulumi:"sapProduct"`
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName *string `pulumi:"sapVirtualInstanceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SAPVirtualInstance resource.
type SAPVirtualInstanceArgs struct {
	// Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
	Configuration pulumi.Input
	// Defines the environment type - Production/Non Production.
	Environment pulumi.StringInput
	// Managed service identity (user assigned identities)
	Identity UserAssignedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Managed resource group configuration
	ManagedResourceGroupConfiguration ManagedRGConfigurationPtrInput
	// Specifies the network access configuration for the resources that will be deployed in the Managed Resource Group. The options to choose from are Public and Private. If 'Private' is chosen, the Storage Account service tag should be enabled on the subnets in which the SAP VMs exist. This is required for establishing connectivity between VM extensions and the managed resource group storage account. This setting is currently applicable only to Storage Account. Learn more here https://go.microsoft.com/fwlink/?linkid=2247228
	ManagedResourcesNetworkAccessType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Defines the SAP Product type.
	SapProduct pulumi.StringInput
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SAPVirtualInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapvirtualInstanceArgs)(nil)).Elem()
}

type SAPVirtualInstanceInput interface {
	pulumi.Input

	ToSAPVirtualInstanceOutput() SAPVirtualInstanceOutput
	ToSAPVirtualInstanceOutputWithContext(ctx context.Context) SAPVirtualInstanceOutput
}

func (*SAPVirtualInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPVirtualInstance)(nil)).Elem()
}

func (i *SAPVirtualInstance) ToSAPVirtualInstanceOutput() SAPVirtualInstanceOutput {
	return i.ToSAPVirtualInstanceOutputWithContext(context.Background())
}

func (i *SAPVirtualInstance) ToSAPVirtualInstanceOutputWithContext(ctx context.Context) SAPVirtualInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAPVirtualInstanceOutput)
}

type SAPVirtualInstanceOutput struct{ *pulumi.OutputState }

func (SAPVirtualInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SAPVirtualInstance)(nil)).Elem()
}

func (o SAPVirtualInstanceOutput) ToSAPVirtualInstanceOutput() SAPVirtualInstanceOutput {
	return o
}

func (o SAPVirtualInstanceOutput) ToSAPVirtualInstanceOutputWithContext(ctx context.Context) SAPVirtualInstanceOutput {
	return o
}

// Defines if the SAP system is being created using Azure Center for SAP solutions (ACSS) or if an existing SAP system is being registered with ACSS
func (o SAPVirtualInstanceOutput) Configuration() pulumi.AnyOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.AnyOutput { return v.Configuration }).(pulumi.AnyOutput)
}

// Defines the environment type - Production/Non Production.
func (o SAPVirtualInstanceOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Indicates any errors on the Virtual Instance for SAP solutions resource.
func (o SAPVirtualInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

// Defines the health of SAP Instances.
func (o SAPVirtualInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// Managed service identity (user assigned identities)
func (o SAPVirtualInstanceOutput) Identity() UserAssignedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) UserAssignedServiceIdentityResponsePtrOutput { return v.Identity }).(UserAssignedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o SAPVirtualInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Managed resource group configuration
func (o SAPVirtualInstanceOutput) ManagedResourceGroupConfiguration() ManagedRGConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) ManagedRGConfigurationResponsePtrOutput {
		return v.ManagedResourceGroupConfiguration
	}).(ManagedRGConfigurationResponsePtrOutput)
}

// Specifies the network access configuration for the resources that will be deployed in the Managed Resource Group. The options to choose from are Public and Private. If 'Private' is chosen, the Storage Account service tag should be enabled on the subnets in which the SAP VMs exist. This is required for establishing connectivity between VM extensions and the managed resource group storage account. This setting is currently applicable only to Storage Account. Learn more here https://go.microsoft.com/fwlink/?linkid=2247228
func (o SAPVirtualInstanceOutput) ManagedResourcesNetworkAccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringPtrOutput { return v.ManagedResourcesNetworkAccessType }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SAPVirtualInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o SAPVirtualInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the SAP Product type.
func (o SAPVirtualInstanceOutput) SapProduct() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.SapProduct }).(pulumi.StringOutput)
}

// Defines the Virtual Instance for SAP state.
func (o SAPVirtualInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Defines the SAP Instance status.
func (o SAPVirtualInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SAPVirtualInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SAPVirtualInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SAPVirtualInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SAPVirtualInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SAPVirtualInstanceOutput{})
}
