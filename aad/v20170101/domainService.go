// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Domain service.
//
// Deprecated: Version 2017-01-01 will be removed in v2 of the provider.
type DomainService struct {
	pulumi.CustomResourceState

	// Deployment Id
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// List of Domain Controller IP Address
	DomainControllerIpAddress pulumi.StringArrayOutput `pulumi:"domainControllerIpAddress"`
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsResponsePtrOutput `pulumi:"domainSecuritySettings"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrOutput `pulumi:"filteredSync"`
	// List of Domain Health Alerts
	HealthAlerts HealthAlertResponseArrayOutput `pulumi:"healthAlerts"`
	// Last domain evaluation run DateTime
	HealthLastEvaluated pulumi.StringOutput `pulumi:"healthLastEvaluated"`
	// List of Domain Health Monitors
	HealthMonitors HealthMonitorResponseArrayOutput `pulumi:"healthMonitors"`
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsResponsePtrOutput `pulumi:"ldapsSettings"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification Settings
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	// the current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Status of Domain Service instance
	ServiceStatus pulumi.StringOutput `pulumi:"serviceStatus"`
	// The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure Active Directory tenant id
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Virtual network site id
	VnetSiteId pulumi.StringOutput `pulumi:"vnetSiteId"`
}

// NewDomainService registers a new resource with the given unique name, arguments, and options.
func NewDomainService(ctx *pulumi.Context,
	name string, args *DomainServiceArgs, opts ...pulumi.ResourceOption) (*DomainService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:aad:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20170601:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20200101:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210301:DomainService"),
		},
		{
			Type: pulumi.String("azure-native:aad/v20210501:DomainService"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainService
	err := ctx.RegisterResource("azure-native:aad/v20170101:DomainService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainService gets an existing DomainService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainServiceState, opts ...pulumi.ResourceOption) (*DomainService, error) {
	var resource DomainService
	err := ctx.ReadResource("azure-native:aad/v20170101:DomainService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainService resources.
type domainServiceState struct {
}

type DomainServiceState struct {
}

func (DomainServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceState)(nil)).Elem()
}

type domainServiceArgs struct {
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName *string `pulumi:"domainName"`
	// DomainSecurity Settings
	DomainSecuritySettings *DomainSecuritySettings `pulumi:"domainSecuritySettings"`
	// The name of the domain service.
	DomainServiceName *string `pulumi:"domainServiceName"`
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync *string `pulumi:"filteredSync"`
	// Secure LDAP Settings
	LdapsSettings *LdapsSettings `pulumi:"ldapsSettings"`
	// Resource location
	Location *string `pulumi:"location"`
	// Notification Settings
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
	SubnetId *string `pulumi:"subnetId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DomainService resource.
type DomainServiceArgs struct {
	// The name of the Azure domain that the user would like to deploy Domain Services to.
	DomainName pulumi.StringPtrInput
	// DomainSecurity Settings
	DomainSecuritySettings DomainSecuritySettingsPtrInput
	// The name of the domain service.
	DomainServiceName pulumi.StringPtrInput
	// Enabled or Disabled flag to turn on Group-based filtered sync
	FilteredSync pulumi.StringPtrInput
	// Secure LDAP Settings
	LdapsSettings LdapsSettingsPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Notification Settings
	NotificationSettings NotificationSettingsPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
	SubnetId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DomainServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainServiceArgs)(nil)).Elem()
}

type DomainServiceInput interface {
	pulumi.Input

	ToDomainServiceOutput() DomainServiceOutput
	ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput
}

func (*DomainService) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainService)(nil)).Elem()
}

func (i *DomainService) ToDomainServiceOutput() DomainServiceOutput {
	return i.ToDomainServiceOutputWithContext(context.Background())
}

func (i *DomainService) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainServiceOutput)
}

type DomainServiceOutput struct{ *pulumi.OutputState }

func (DomainServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainService)(nil)).Elem()
}

func (o DomainServiceOutput) ToDomainServiceOutput() DomainServiceOutput {
	return o
}

func (o DomainServiceOutput) ToDomainServiceOutputWithContext(ctx context.Context) DomainServiceOutput {
	return o
}

// Deployment Id
func (o DomainServiceOutput) DeploymentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.DeploymentId }).(pulumi.StringOutput)
}

// List of Domain Controller IP Address
func (o DomainServiceOutput) DomainControllerIpAddress() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringArrayOutput { return v.DomainControllerIpAddress }).(pulumi.StringArrayOutput)
}

// The name of the Azure domain that the user would like to deploy Domain Services to.
func (o DomainServiceOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// DomainSecurity Settings
func (o DomainServiceOutput) DomainSecuritySettings() DomainSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) DomainSecuritySettingsResponsePtrOutput { return v.DomainSecuritySettings }).(DomainSecuritySettingsResponsePtrOutput)
}

// Resource etag
func (o DomainServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Enabled or Disabled flag to turn on Group-based filtered sync
func (o DomainServiceOutput) FilteredSync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.FilteredSync }).(pulumi.StringPtrOutput)
}

// List of Domain Health Alerts
func (o DomainServiceOutput) HealthAlerts() HealthAlertResponseArrayOutput {
	return o.ApplyT(func(v *DomainService) HealthAlertResponseArrayOutput { return v.HealthAlerts }).(HealthAlertResponseArrayOutput)
}

// Last domain evaluation run DateTime
func (o DomainServiceOutput) HealthLastEvaluated() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.HealthLastEvaluated }).(pulumi.StringOutput)
}

// List of Domain Health Monitors
func (o DomainServiceOutput) HealthMonitors() HealthMonitorResponseArrayOutput {
	return o.ApplyT(func(v *DomainService) HealthMonitorResponseArrayOutput { return v.HealthMonitors }).(HealthMonitorResponseArrayOutput)
}

// Secure LDAP Settings
func (o DomainServiceOutput) LdapsSettings() LdapsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) LdapsSettingsResponsePtrOutput { return v.LdapsSettings }).(LdapsSettingsResponsePtrOutput)
}

// Resource location
func (o DomainServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o DomainServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notification Settings
func (o DomainServiceOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *DomainService) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

// the current deployment or provisioning state, which only appears in the response.
func (o DomainServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Status of Domain Service instance
func (o DomainServiceOutput) ServiceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.ServiceStatus }).(pulumi.StringOutput)
}

// The name of the virtual network that Domain Services will be deployed on. The id of the subnet that Domain Services will be deployed on. /virtualNetwork/vnetName/subnets/subnetName.
func (o DomainServiceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Resource tags
func (o DomainServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure Active Directory tenant id
func (o DomainServiceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Resource type
func (o DomainServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Virtual network site id
func (o DomainServiceOutput) VnetSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainService) pulumi.StringOutput { return v.VnetSiteId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainServiceOutput{})
}