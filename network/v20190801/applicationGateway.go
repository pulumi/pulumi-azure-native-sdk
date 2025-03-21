// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Application gateway resource.
type ApplicationGateway struct {
	pulumi.CustomResourceState

	// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	AuthenticationCertificates ApplicationGatewayAuthenticationCertificateResponseArrayOutput `pulumi:"authenticationCertificates"`
	// Autoscale Configuration.
	AutoscaleConfiguration ApplicationGatewayAutoscaleConfigurationResponsePtrOutput `pulumi:"autoscaleConfiguration"`
	// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendAddressPools ApplicationGatewayBackendAddressPoolResponseArrayOutput `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsResponseArrayOutput `pulumi:"backendHttpSettingsCollection"`
	// Custom error configurations of the application gateway resource.
	CustomErrorConfigurations ApplicationGatewayCustomErrorResponseArrayOutput `pulumi:"customErrorConfigurations"`
	// Whether FIPS is enabled on the application gateway resource.
	EnableFips pulumi.BoolPtrOutput `pulumi:"enableFips"`
	// Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 pulumi.BoolPtrOutput `pulumi:"enableHttp2"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Reference of the FirewallPolicy resource.
	FirewallPolicy SubResourceResponsePtrOutput `pulumi:"firewallPolicy"`
	// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendIPConfigurations ApplicationGatewayFrontendIPConfigurationResponseArrayOutput `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendPorts ApplicationGatewayFrontendPortResponseArrayOutput `pulumi:"frontendPorts"`
	// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	GatewayIPConfigurations ApplicationGatewayIPConfigurationResponseArrayOutput `pulumi:"gatewayIPConfigurations"`
	// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	HttpListeners ApplicationGatewayHttpListenerResponseArrayOutput `pulumi:"httpListeners"`
	// The identity of the application gateway, if configured.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Operational state of the application gateway resource.
	OperationalState pulumi.StringOutput `pulumi:"operationalState"`
	// Probes of the application gateway resource.
	Probes ApplicationGatewayProbeResponseArrayOutput `pulumi:"probes"`
	// The provisioning state of the application gateway resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	RedirectConfigurations ApplicationGatewayRedirectConfigurationResponseArrayOutput `pulumi:"redirectConfigurations"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleResponseArrayOutput `pulumi:"requestRoutingRules"`
	// The resource GUID property of the application gateway resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// Rewrite rules for the application gateway resource.
	RewriteRuleSets ApplicationGatewayRewriteRuleSetResponseArrayOutput `pulumi:"rewriteRuleSets"`
	// SKU of the application gateway resource.
	Sku ApplicationGatewaySkuResponsePtrOutput `pulumi:"sku"`
	// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslCertificates ApplicationGatewaySslCertificateResponseArrayOutput `pulumi:"sslCertificates"`
	// SSL policy of the application gateway resource.
	SslPolicy ApplicationGatewaySslPolicyResponsePtrOutput `pulumi:"sslPolicy"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedRootCertificates ApplicationGatewayTrustedRootCertificateResponseArrayOutput `pulumi:"trustedRootCertificates"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	UrlPathMaps ApplicationGatewayUrlPathMapResponseArrayOutput `pulumi:"urlPathMaps"`
	// Web application firewall configuration.
	WebApplicationFirewallConfiguration ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput `pulumi:"webApplicationFirewallConfiguration"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewApplicationGateway registers a new resource with the given unique name, arguments, and options.
func NewApplicationGateway(ctx *pulumi.Context,
	name string, args *ApplicationGatewayArgs, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20221101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20230901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20231101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-native:network:ApplicationGateway"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApplicationGateway
	err := ctx.RegisterResource("azure-native:network/v20190801:ApplicationGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGateway gets an existing ApplicationGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayState, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	var resource ApplicationGateway
	err := ctx.ReadResource("azure-native:network/v20190801:ApplicationGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGateway resources.
type applicationGatewayState struct {
}

type ApplicationGatewayState struct {
}

func (ApplicationGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayState)(nil)).Elem()
}

type applicationGatewayArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName *string `pulumi:"applicationGatewayName"`
	// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	AuthenticationCertificates []ApplicationGatewayAuthenticationCertificate `pulumi:"authenticationCertificates"`
	// Autoscale Configuration.
	AutoscaleConfiguration *ApplicationGatewayAutoscaleConfiguration `pulumi:"autoscaleConfiguration"`
	// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettings `pulumi:"backendHttpSettingsCollection"`
	// Custom error configurations of the application gateway resource.
	CustomErrorConfigurations []ApplicationGatewayCustomError `pulumi:"customErrorConfigurations"`
	// Whether FIPS is enabled on the application gateway resource.
	EnableFips *bool `pulumi:"enableFips"`
	// Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 *bool `pulumi:"enableHttp2"`
	// Reference of the FirewallPolicy resource.
	FirewallPolicy *SubResource `pulumi:"firewallPolicy"`
	// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendIPConfigurations []ApplicationGatewayFrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendPorts []ApplicationGatewayFrontendPort `pulumi:"frontendPorts"`
	// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	GatewayIPConfigurations []ApplicationGatewayIPConfiguration `pulumi:"gatewayIPConfigurations"`
	// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	HttpListeners []ApplicationGatewayHttpListener `pulumi:"httpListeners"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The identity of the application gateway, if configured.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Probes of the application gateway resource.
	Probes []ApplicationGatewayProbe `pulumi:"probes"`
	// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	RedirectConfigurations []ApplicationGatewayRedirectConfiguration `pulumi:"redirectConfigurations"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRule `pulumi:"requestRoutingRules"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the application gateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Rewrite rules for the application gateway resource.
	RewriteRuleSets []ApplicationGatewayRewriteRuleSet `pulumi:"rewriteRuleSets"`
	// SKU of the application gateway resource.
	Sku *ApplicationGatewaySku `pulumi:"sku"`
	// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslCertificates []ApplicationGatewaySslCertificate `pulumi:"sslCertificates"`
	// SSL policy of the application gateway resource.
	SslPolicy *ApplicationGatewaySslPolicy `pulumi:"sslPolicy"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedRootCertificates []ApplicationGatewayTrustedRootCertificate `pulumi:"trustedRootCertificates"`
	// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	UrlPathMaps []ApplicationGatewayUrlPathMap `pulumi:"urlPathMaps"`
	// Web application firewall configuration.
	WebApplicationFirewallConfiguration *ApplicationGatewayWebApplicationFirewallConfiguration `pulumi:"webApplicationFirewallConfiguration"`
	// A list of availability zones denoting where the resource needs to come from.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a ApplicationGateway resource.
type ApplicationGatewayArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName pulumi.StringPtrInput
	// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	AuthenticationCertificates ApplicationGatewayAuthenticationCertificateArrayInput
	// Autoscale Configuration.
	AutoscaleConfiguration ApplicationGatewayAutoscaleConfigurationPtrInput
	// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput
	// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsArrayInput
	// Custom error configurations of the application gateway resource.
	CustomErrorConfigurations ApplicationGatewayCustomErrorArrayInput
	// Whether FIPS is enabled on the application gateway resource.
	EnableFips pulumi.BoolPtrInput
	// Whether HTTP2 is enabled on the application gateway resource.
	EnableHttp2 pulumi.BoolPtrInput
	// Reference of the FirewallPolicy resource.
	FirewallPolicy SubResourcePtrInput
	// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendIPConfigurations ApplicationGatewayFrontendIPConfigurationArrayInput
	// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	FrontendPorts ApplicationGatewayFrontendPortArrayInput
	// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	GatewayIPConfigurations ApplicationGatewayIPConfigurationArrayInput
	// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	HttpListeners ApplicationGatewayHttpListenerArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The identity of the application gateway, if configured.
	Identity ManagedServiceIdentityPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Probes of the application gateway resource.
	Probes ApplicationGatewayProbeArrayInput
	// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	RedirectConfigurations ApplicationGatewayRedirectConfigurationArrayInput
	// Request routing rules of the application gateway resource.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the application gateway resource.
	ResourceGuid pulumi.StringPtrInput
	// Rewrite rules for the application gateway resource.
	RewriteRuleSets ApplicationGatewayRewriteRuleSetArrayInput
	// SKU of the application gateway resource.
	Sku ApplicationGatewaySkuPtrInput
	// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	SslCertificates ApplicationGatewaySslCertificateArrayInput
	// SSL policy of the application gateway resource.
	SslPolicy ApplicationGatewaySslPolicyPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	TrustedRootCertificates ApplicationGatewayTrustedRootCertificateArrayInput
	// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
	UrlPathMaps ApplicationGatewayUrlPathMapArrayInput
	// Web application firewall configuration.
	WebApplicationFirewallConfiguration ApplicationGatewayWebApplicationFirewallConfigurationPtrInput
	// A list of availability zones denoting where the resource needs to come from.
	Zones pulumi.StringArrayInput
}

func (ApplicationGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayArgs)(nil)).Elem()
}

type ApplicationGatewayInput interface {
	pulumi.Input

	ToApplicationGatewayOutput() ApplicationGatewayOutput
	ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput
}

func (*ApplicationGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGateway)(nil)).Elem()
}

func (i *ApplicationGateway) ToApplicationGatewayOutput() ApplicationGatewayOutput {
	return i.ToApplicationGatewayOutputWithContext(context.Background())
}

func (i *ApplicationGateway) ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayOutput)
}

type ApplicationGatewayOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGateway)(nil)).Elem()
}

func (o ApplicationGatewayOutput) ToApplicationGatewayOutput() ApplicationGatewayOutput {
	return o
}

func (o ApplicationGatewayOutput) ToApplicationGatewayOutputWithContext(ctx context.Context) ApplicationGatewayOutput {
	return o
}

// Authentication certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) AuthenticationCertificates() ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayAuthenticationCertificateResponseArrayOutput {
		return v.AuthenticationCertificates
	}).(ApplicationGatewayAuthenticationCertificateResponseArrayOutput)
}

// Autoscale Configuration.
func (o ApplicationGatewayOutput) AutoscaleConfiguration() ApplicationGatewayAutoscaleConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayAutoscaleConfigurationResponsePtrOutput {
		return v.AutoscaleConfiguration
	}).(ApplicationGatewayAutoscaleConfigurationResponsePtrOutput)
}

// Backend address pool of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) BackendAddressPools() ApplicationGatewayBackendAddressPoolResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayBackendAddressPoolResponseArrayOutput {
		return v.BackendAddressPools
	}).(ApplicationGatewayBackendAddressPoolResponseArrayOutput)
}

// Backend http settings of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) BackendHttpSettingsCollection() ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayBackendHttpSettingsResponseArrayOutput {
		return v.BackendHttpSettingsCollection
	}).(ApplicationGatewayBackendHttpSettingsResponseArrayOutput)
}

// Custom error configurations of the application gateway resource.
func (o ApplicationGatewayOutput) CustomErrorConfigurations() ApplicationGatewayCustomErrorResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayCustomErrorResponseArrayOutput {
		return v.CustomErrorConfigurations
	}).(ApplicationGatewayCustomErrorResponseArrayOutput)
}

// Whether FIPS is enabled on the application gateway resource.
func (o ApplicationGatewayOutput) EnableFips() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.BoolPtrOutput { return v.EnableFips }).(pulumi.BoolPtrOutput)
}

// Whether HTTP2 is enabled on the application gateway resource.
func (o ApplicationGatewayOutput) EnableHttp2() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.BoolPtrOutput { return v.EnableHttp2 }).(pulumi.BoolPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o ApplicationGatewayOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// Reference of the FirewallPolicy resource.
func (o ApplicationGatewayOutput) FirewallPolicy() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) SubResourceResponsePtrOutput { return v.FirewallPolicy }).(SubResourceResponsePtrOutput)
}

// Frontend IP addresses of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) FrontendIPConfigurations() ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayFrontendIPConfigurationResponseArrayOutput {
		return v.FrontendIPConfigurations
	}).(ApplicationGatewayFrontendIPConfigurationResponseArrayOutput)
}

// Frontend ports of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) FrontendPorts() ApplicationGatewayFrontendPortResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayFrontendPortResponseArrayOutput { return v.FrontendPorts }).(ApplicationGatewayFrontendPortResponseArrayOutput)
}

// Subnets of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) GatewayIPConfigurations() ApplicationGatewayIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayIPConfigurationResponseArrayOutput {
		return v.GatewayIPConfigurations
	}).(ApplicationGatewayIPConfigurationResponseArrayOutput)
}

// Http listeners of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) HttpListeners() ApplicationGatewayHttpListenerResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayHttpListenerResponseArrayOutput { return v.HttpListeners }).(ApplicationGatewayHttpListenerResponseArrayOutput)
}

// The identity of the application gateway, if configured.
func (o ApplicationGatewayOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Resource location.
func (o ApplicationGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ApplicationGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Operational state of the application gateway resource.
func (o ApplicationGatewayOutput) OperationalState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.OperationalState }).(pulumi.StringOutput)
}

// Probes of the application gateway resource.
func (o ApplicationGatewayOutput) Probes() ApplicationGatewayProbeResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayProbeResponseArrayOutput { return v.Probes }).(ApplicationGatewayProbeResponseArrayOutput)
}

// The provisioning state of the application gateway resource.
func (o ApplicationGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Redirect configurations of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) RedirectConfigurations() ApplicationGatewayRedirectConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRedirectConfigurationResponseArrayOutput {
		return v.RedirectConfigurations
	}).(ApplicationGatewayRedirectConfigurationResponseArrayOutput)
}

// Request routing rules of the application gateway resource.
func (o ApplicationGatewayOutput) RequestRoutingRules() ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRequestRoutingRuleResponseArrayOutput {
		return v.RequestRoutingRules
	}).(ApplicationGatewayRequestRoutingRuleResponseArrayOutput)
}

// The resource GUID property of the application gateway resource.
func (o ApplicationGatewayOutput) ResourceGuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringPtrOutput { return v.ResourceGuid }).(pulumi.StringPtrOutput)
}

// Rewrite rules for the application gateway resource.
func (o ApplicationGatewayOutput) RewriteRuleSets() ApplicationGatewayRewriteRuleSetResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayRewriteRuleSetResponseArrayOutput {
		return v.RewriteRuleSets
	}).(ApplicationGatewayRewriteRuleSetResponseArrayOutput)
}

// SKU of the application gateway resource.
func (o ApplicationGatewayOutput) Sku() ApplicationGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySkuResponsePtrOutput { return v.Sku }).(ApplicationGatewaySkuResponsePtrOutput)
}

// SSL certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) SslCertificates() ApplicationGatewaySslCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySslCertificateResponseArrayOutput {
		return v.SslCertificates
	}).(ApplicationGatewaySslCertificateResponseArrayOutput)
}

// SSL policy of the application gateway resource.
func (o ApplicationGatewayOutput) SslPolicy() ApplicationGatewaySslPolicyResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewaySslPolicyResponsePtrOutput { return v.SslPolicy }).(ApplicationGatewaySslPolicyResponsePtrOutput)
}

// Resource tags.
func (o ApplicationGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Trusted Root certificates of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) TrustedRootCertificates() ApplicationGatewayTrustedRootCertificateResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayTrustedRootCertificateResponseArrayOutput {
		return v.TrustedRootCertificates
	}).(ApplicationGatewayTrustedRootCertificateResponseArrayOutput)
}

// Resource type.
func (o ApplicationGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// URL path map of the application gateway resource. For default limits, see [Application Gateway limits](https://docs.microsoft.com/azure/azure-subscription-service-limits#application-gateway-limits).
func (o ApplicationGatewayOutput) UrlPathMaps() ApplicationGatewayUrlPathMapResponseArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayUrlPathMapResponseArrayOutput { return v.UrlPathMaps }).(ApplicationGatewayUrlPathMapResponseArrayOutput)
}

// Web application firewall configuration.
func (o ApplicationGatewayOutput) WebApplicationFirewallConfiguration() ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGateway) ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput {
		return v.WebApplicationFirewallConfiguration
	}).(ApplicationGatewayWebApplicationFirewallConfigurationResponsePtrOutput)
}

// A list of availability zones denoting where the resource needs to come from.
func (o ApplicationGatewayOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApplicationGateway) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGatewayOutput{})
}
