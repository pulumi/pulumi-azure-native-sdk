// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Define the SAP Application Server Instance resource.
//
// Uses Azure REST API version 2024-09-01.
type SapApplicationServerInstance struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Application server instance dispatcher status.
	DispatcherStatus pulumi.StringOutput `pulumi:"dispatcherStatus"`
	// Defines the Application Instance errors.
	Errors SAPVirtualInstanceErrorResponseOutput `pulumi:"errors"`
	// Application server instance gateway Port.
	GatewayPort pulumi.Float64Output `pulumi:"gatewayPort"`
	// Defines the health of SAP Instances.
	Health pulumi.StringOutput `pulumi:"health"`
	// Application server instance SAP hostname.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Application server instance ICM HTTP Port.
	IcmHttpPort pulumi.Float64Output `pulumi:"icmHttpPort"`
	// Application server instance ICM HTTPS Port.
	IcmHttpsPort pulumi.Float64Output `pulumi:"icmHttpsPort"`
	// Application server Instance Number.
	InstanceNo pulumi.StringOutput `pulumi:"instanceNo"`
	// Application server instance SAP IP Address.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Application server instance SAP Kernel Patch level.
	KernelPatch pulumi.StringOutput `pulumi:"kernelPatch"`
	// Application server instance SAP Kernel Version.
	KernelVersion pulumi.StringOutput `pulumi:"kernelVersion"`
	// The Load Balancer details such as LoadBalancer ID attached to Application Server Virtual Machines
	LoadBalancerDetails LoadBalancerDetailsResponseOutput `pulumi:"loadBalancerDetails"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the provisioning states.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Defines the SAP Instance status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Application server Subnet.
	Subnet pulumi.StringOutput `pulumi:"subnet"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of virtual machines.
	VmDetails ApplicationServerVmDetailsResponseArrayOutput `pulumi:"vmDetails"`
}

// NewSapApplicationServerInstance registers a new resource with the given unique name, arguments, and options.
func NewSapApplicationServerInstance(ctx *pulumi.Context,
	name string, args *SapApplicationServerInstanceArgs, opts ...pulumi.ResourceOption) (*SapApplicationServerInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SapVirtualInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'SapVirtualInstanceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:SAPApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:SapApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20221101preview:SapApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20230401:SAPApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20230401:SapApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20231001preview:SAPApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20231001preview:SapApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20240901:SapApplicationServerInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads:SAPApplicationServerInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SapApplicationServerInstance
	err := ctx.RegisterResource("azure-native:workloads:SapApplicationServerInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSapApplicationServerInstance gets an existing SapApplicationServerInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSapApplicationServerInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SapApplicationServerInstanceState, opts ...pulumi.ResourceOption) (*SapApplicationServerInstance, error) {
	var resource SapApplicationServerInstance
	err := ctx.ReadResource("azure-native:workloads:SapApplicationServerInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SapApplicationServerInstance resources.
type sapApplicationServerInstanceState struct {
}

type SapApplicationServerInstanceState struct {
}

func (SapApplicationServerInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapApplicationServerInstanceState)(nil)).Elem()
}

type sapApplicationServerInstanceArgs struct {
	// The name of SAP Application Server instance resource.
	ApplicationInstanceName *string `pulumi:"applicationInstanceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName string `pulumi:"sapVirtualInstanceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SapApplicationServerInstance resource.
type SapApplicationServerInstanceArgs struct {
	// The name of SAP Application Server instance resource.
	ApplicationInstanceName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Virtual Instances for SAP solutions resource
	SapVirtualInstanceName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SapApplicationServerInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapApplicationServerInstanceArgs)(nil)).Elem()
}

type SapApplicationServerInstanceInput interface {
	pulumi.Input

	ToSapApplicationServerInstanceOutput() SapApplicationServerInstanceOutput
	ToSapApplicationServerInstanceOutputWithContext(ctx context.Context) SapApplicationServerInstanceOutput
}

func (*SapApplicationServerInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**SapApplicationServerInstance)(nil)).Elem()
}

func (i *SapApplicationServerInstance) ToSapApplicationServerInstanceOutput() SapApplicationServerInstanceOutput {
	return i.ToSapApplicationServerInstanceOutputWithContext(context.Background())
}

func (i *SapApplicationServerInstance) ToSapApplicationServerInstanceOutputWithContext(ctx context.Context) SapApplicationServerInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SapApplicationServerInstanceOutput)
}

type SapApplicationServerInstanceOutput struct{ *pulumi.OutputState }

func (SapApplicationServerInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SapApplicationServerInstance)(nil)).Elem()
}

func (o SapApplicationServerInstanceOutput) ToSapApplicationServerInstanceOutput() SapApplicationServerInstanceOutput {
	return o
}

func (o SapApplicationServerInstanceOutput) ToSapApplicationServerInstanceOutputWithContext(ctx context.Context) SapApplicationServerInstanceOutput {
	return o
}

// The Azure API version of the resource.
func (o SapApplicationServerInstanceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Application server instance dispatcher status.
func (o SapApplicationServerInstanceOutput) DispatcherStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.DispatcherStatus }).(pulumi.StringOutput)
}

// Defines the Application Instance errors.
func (o SapApplicationServerInstanceOutput) Errors() SAPVirtualInstanceErrorResponseOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) SAPVirtualInstanceErrorResponseOutput { return v.Errors }).(SAPVirtualInstanceErrorResponseOutput)
}

// Application server instance gateway Port.
func (o SapApplicationServerInstanceOutput) GatewayPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.Float64Output { return v.GatewayPort }).(pulumi.Float64Output)
}

// Defines the health of SAP Instances.
func (o SapApplicationServerInstanceOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// Application server instance SAP hostname.
func (o SapApplicationServerInstanceOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Application server instance ICM HTTP Port.
func (o SapApplicationServerInstanceOutput) IcmHttpPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.Float64Output { return v.IcmHttpPort }).(pulumi.Float64Output)
}

// Application server instance ICM HTTPS Port.
func (o SapApplicationServerInstanceOutput) IcmHttpsPort() pulumi.Float64Output {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.Float64Output { return v.IcmHttpsPort }).(pulumi.Float64Output)
}

// Application server Instance Number.
func (o SapApplicationServerInstanceOutput) InstanceNo() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.InstanceNo }).(pulumi.StringOutput)
}

// Application server instance SAP IP Address.
func (o SapApplicationServerInstanceOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Application server instance SAP Kernel Patch level.
func (o SapApplicationServerInstanceOutput) KernelPatch() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.KernelPatch }).(pulumi.StringOutput)
}

// Application server instance SAP Kernel Version.
func (o SapApplicationServerInstanceOutput) KernelVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.KernelVersion }).(pulumi.StringOutput)
}

// The Load Balancer details such as LoadBalancer ID attached to Application Server Virtual Machines
func (o SapApplicationServerInstanceOutput) LoadBalancerDetails() LoadBalancerDetailsResponseOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) LoadBalancerDetailsResponseOutput { return v.LoadBalancerDetails }).(LoadBalancerDetailsResponseOutput)
}

// The geo-location where the resource lives
func (o SapApplicationServerInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SapApplicationServerInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the provisioning states.
func (o SapApplicationServerInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Defines the SAP Instance status.
func (o SapApplicationServerInstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Application server Subnet.
func (o SapApplicationServerInstanceOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Subnet }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SapApplicationServerInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SapApplicationServerInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SapApplicationServerInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The list of virtual machines.
func (o SapApplicationServerInstanceOutput) VmDetails() ApplicationServerVmDetailsResponseArrayOutput {
	return o.ApplyT(func(v *SapApplicationServerInstance) ApplicationServerVmDetailsResponseArrayOutput {
		return v.VmDetails
	}).(ApplicationServerVmDetailsResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SapApplicationServerInstanceOutput{})
}
