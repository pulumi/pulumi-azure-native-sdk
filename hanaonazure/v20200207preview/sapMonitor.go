// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200207preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
type SapMonitor struct {
	pulumi.CustomResourceState

	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics pulumi.BoolPtrOutput `pulumi:"enableCustomerAnalytics"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmId pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceArmId"`
	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceId pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey pulumi.StringPtrOutput `pulumi:"logAnalyticsWorkspaceSharedKey"`
	// The name of the resource group the SAP Monitor resources get deployed into.
	ManagedResourceGroupName pulumi.StringOutput `pulumi:"managedResourceGroupName"`
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet pulumi.StringPtrOutput `pulumi:"monitorSubnet"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of provisioning of the HanaInstance
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The version of the payload running in the Collector VM
	SapMonitorCollectorVersion pulumi.StringOutput `pulumi:"sapMonitorCollectorVersion"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSapMonitor registers a new resource with the given unique name, arguments, and options.
func NewSapMonitor(ctx *pulumi.Context,
	name string, args *SapMonitorArgs, opts ...pulumi.ResourceOption) (*SapMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hanaonazure:SapMonitor"),
		},
	})
	opts = append(opts, aliases)
	var resource SapMonitor
	err := ctx.RegisterResource("azure-native:hanaonazure/v20200207preview:SapMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSapMonitor gets an existing SapMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSapMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SapMonitorState, opts ...pulumi.ResourceOption) (*SapMonitor, error) {
	var resource SapMonitor
	err := ctx.ReadResource("azure-native:hanaonazure/v20200207preview:SapMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SapMonitor resources.
type sapMonitorState struct {
}

type SapMonitorState struct {
}

func (SapMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*sapMonitorState)(nil)).Elem()
}

type sapMonitorArgs struct {
	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics *bool `pulumi:"enableCustomerAnalytics"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmId *string `pulumi:"logAnalyticsWorkspaceArmId"`
	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey *string `pulumi:"logAnalyticsWorkspaceSharedKey"`
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet *string `pulumi:"monitorSubnet"`
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SAP monitor resource.
	SapMonitorName *string `pulumi:"sapMonitorName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SapMonitor resource.
type SapMonitorArgs struct {
	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmId pulumi.StringPtrInput
	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey pulumi.StringPtrInput
	// The subnet which the SAP monitor will be deployed in
	MonitorSubnet pulumi.StringPtrInput
	// Name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Name of the SAP monitor resource.
	SapMonitorName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SapMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sapMonitorArgs)(nil)).Elem()
}

type SapMonitorInput interface {
	pulumi.Input

	ToSapMonitorOutput() SapMonitorOutput
	ToSapMonitorOutputWithContext(ctx context.Context) SapMonitorOutput
}

func (*SapMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**SapMonitor)(nil)).Elem()
}

func (i *SapMonitor) ToSapMonitorOutput() SapMonitorOutput {
	return i.ToSapMonitorOutputWithContext(context.Background())
}

func (i *SapMonitor) ToSapMonitorOutputWithContext(ctx context.Context) SapMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SapMonitorOutput)
}

type SapMonitorOutput struct{ *pulumi.OutputState }

func (SapMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SapMonitor)(nil)).Elem()
}

func (o SapMonitorOutput) ToSapMonitorOutput() SapMonitorOutput {
	return o
}

func (o SapMonitorOutput) ToSapMonitorOutputWithContext(ctx context.Context) SapMonitorOutput {
	return o
}

// The value indicating whether to send analytics to Microsoft
func (o SapMonitorOutput) EnableCustomerAnalytics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.BoolPtrOutput { return v.EnableCustomerAnalytics }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o SapMonitorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The ARM ID of the Log Analytics Workspace that is used for monitoring
func (o SapMonitorOutput) LogAnalyticsWorkspaceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceArmId }).(pulumi.StringPtrOutput)
}

// The workspace ID of the log analytics workspace to be used for monitoring
func (o SapMonitorOutput) LogAnalyticsWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceId }).(pulumi.StringPtrOutput)
}

// The shared key of the log analytics workspace that is used for monitoring
func (o SapMonitorOutput) LogAnalyticsWorkspaceSharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.LogAnalyticsWorkspaceSharedKey }).(pulumi.StringPtrOutput)
}

// The name of the resource group the SAP Monitor resources get deployed into.
func (o SapMonitorOutput) ManagedResourceGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.ManagedResourceGroupName }).(pulumi.StringOutput)
}

// The subnet which the SAP monitor will be deployed in
func (o SapMonitorOutput) MonitorSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringPtrOutput { return v.MonitorSubnet }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SapMonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of provisioning of the HanaInstance
func (o SapMonitorOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The version of the payload running in the Collector VM
func (o SapMonitorOutput) SapMonitorCollectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.SapMonitorCollectorVersion }).(pulumi.StringOutput)
}

// Resource tags.
func (o SapMonitorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SapMonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SapMonitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SapMonitorOutput{})
}
