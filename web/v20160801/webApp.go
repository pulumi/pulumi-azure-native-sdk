// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A web app, a mobile app backend, or an API app.
type WebApp struct {
	pulumi.CustomResourceState

	// Management information availability state for the app.
	AvailabilityState pulumi.StringOutput `pulumi:"availabilityState"`
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled pulumi.BoolPtrOutput `pulumi:"clientAffinityEnabled"`
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled pulumi.BoolPtrOutput `pulumi:"clientCertEnabled"`
	// Size of the function container.
	ContainerSize pulumi.IntPtrOutput `pulumi:"containerSize"`
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota pulumi.IntPtrOutput `pulumi:"dailyMemoryTimeQuota"`
	// Default hostname of the app. Read-only.
	DefaultHostName pulumi.StringOutput `pulumi:"defaultHostName"`
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
	// the app is not served on those hostnames.
	EnabledHostNames pulumi.StringArrayOutput `pulumi:"enabledHostNames"`
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates HostNameSslStateResponseArrayOutput `pulumi:"hostNameSslStates"`
	// Hostnames associated with the app.
	HostNames pulumi.StringArrayOutput `pulumi:"hostNames"`
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled pulumi.BoolPtrOutput `pulumi:"hostNamesDisabled"`
	// App Service Environment to use for the app.
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// Managed service identity.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// <code>true</code> if the app is a default container; otherwise, <code>false</code>.
	IsDefaultContainer pulumi.BoolOutput `pulumi:"isDefaultContainer"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Last time the app was modified, in UTC. Read-only.
	LastModifiedTimeUtc pulumi.StringOutput `pulumi:"lastModifiedTimeUtc"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of workers.
	// This only applies to Functions container.
	MaxNumberOfWorkers pulumi.IntOutput `pulumi:"maxNumberOfWorkers"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
	OutboundIpAddresses pulumi.StringOutput `pulumi:"outboundIpAddresses"`
	// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants. Read-only.
	PossibleOutboundIpAddresses pulumi.StringOutput `pulumi:"possibleOutboundIpAddresses"`
	// Name of the repository site.
	RepositorySiteName pulumi.StringOutput `pulumi:"repositorySiteName"`
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved pulumi.BoolPtrOutput `pulumi:"reserved"`
	// Name of the resource group the app belongs to. Read-only.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped pulumi.BoolPtrOutput `pulumi:"scmSiteAlsoStopped"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId pulumi.StringPtrOutput `pulumi:"serverFarmId"`
	// Configuration of the app.
	SiteConfig SiteConfigResponsePtrOutput `pulumi:"siteConfig"`
	// Status of the last deployment slot swap operation.
	SlotSwapStatus SlotSwapStatusResponseOutput `pulumi:"slotSwapStatus"`
	// Current state of the app.
	State pulumi.StringOutput `pulumi:"state"`
	// App suspended till in case memory-time quota is exceeded.
	SuspendedTill pulumi.StringOutput `pulumi:"suspendedTill"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies which deployment slot this app will swap into. Read-only.
	TargetSwapSlot pulumi.StringOutput `pulumi:"targetSwapSlot"`
	// Azure Traffic Manager hostnames associated with the app. Read-only.
	TrafficManagerHostNames pulumi.StringArrayOutput `pulumi:"trafficManagerHostNames"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// State indicating whether the app has exceeded its quota usage. Read-only.
	UsageState pulumi.StringOutput `pulumi:"usageState"`
}

// NewWebApp registers a new resource with the given unique name, arguments, and options.
func NewWebApp(ctx *pulumi.Context,
	name string, args *WebAppArgs, opts ...pulumi.ResourceOption) (*WebApp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Reserved == nil {
		args.Reserved = pulumi.BoolPtr(false)
	}
	if args.ScmSiteAlsoStopped == nil {
		args.ScmSiteAlsoStopped = pulumi.BoolPtr(false)
	}
	if args.SiteConfig != nil {
		args.SiteConfig = args.SiteConfig.ToSiteConfigPtrOutput().ApplyT(func(v *SiteConfig) *SiteConfig { return v.Defaults() }).(SiteConfigPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web/v20150801:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220901:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20230101:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20231201:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web/v20240401:WebApp"),
		},
		{
			Type: pulumi.String("azure-native:web:WebApp"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WebApp
	err := ctx.RegisterResource("azure-native:web/v20160801:WebApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebApp gets an existing WebApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppState, opts ...pulumi.ResourceOption) (*WebApp, error) {
	var resource WebApp
	err := ctx.ReadResource("azure-native:web/v20160801:WebApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebApp resources.
type webAppState struct {
}

type WebAppState struct {
}

func (WebAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppState)(nil)).Elem()
}

type webAppArgs struct {
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// If specified during app creation, the app is cloned from a source app.
	CloningInfo *CloningInfo `pulumi:"cloningInfo"`
	// Size of the function container.
	ContainerSize *int `pulumi:"containerSize"`
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled *bool `pulumi:"enabled"`
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates []HostNameSslState `pulumi:"hostNameSslStates"`
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled *bool `pulumi:"hostNamesDisabled"`
	// App Service Environment to use for the app.
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Managed service identity.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
	Name *string `pulumi:"name"`
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved *bool `pulumi:"reserved"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped *bool `pulumi:"scmSiteAlsoStopped"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `pulumi:"serverFarmId"`
	// Configuration of the app.
	SiteConfig *SiteConfig `pulumi:"siteConfig"`
	// If specified during app creation, the app is created from a previous snapshot.
	SnapshotInfo *SnapshotRecoveryRequest `pulumi:"snapshotInfo"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a WebApp resource.
type WebAppArgs struct {
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled pulumi.BoolPtrInput
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled pulumi.BoolPtrInput
	// If specified during app creation, the app is cloned from a source app.
	CloningInfo CloningInfoPtrInput
	// Size of the function container.
	ContainerSize pulumi.IntPtrInput
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota pulumi.IntPtrInput
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled pulumi.BoolPtrInput
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates HostNameSslStateArrayInput
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled pulumi.BoolPtrInput
	// App Service Environment to use for the app.
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly pulumi.BoolPtrInput
	// Managed service identity.
	Identity ManagedServiceIdentityPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
	Name pulumi.StringPtrInput
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved pulumi.BoolPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped pulumi.BoolPtrInput
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId pulumi.StringPtrInput
	// Configuration of the app.
	SiteConfig SiteConfigPtrInput
	// If specified during app creation, the app is created from a previous snapshot.
	SnapshotInfo SnapshotRecoveryRequestPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (WebAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppArgs)(nil)).Elem()
}

type WebAppInput interface {
	pulumi.Input

	ToWebAppOutput() WebAppOutput
	ToWebAppOutputWithContext(ctx context.Context) WebAppOutput
}

func (*WebApp) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil)).Elem()
}

func (i *WebApp) ToWebAppOutput() WebAppOutput {
	return i.ToWebAppOutputWithContext(context.Background())
}

func (i *WebApp) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppOutput)
}

type WebAppOutput struct{ *pulumi.OutputState }

func (WebAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebApp)(nil)).Elem()
}

func (o WebAppOutput) ToWebAppOutput() WebAppOutput {
	return o
}

func (o WebAppOutput) ToWebAppOutputWithContext(ctx context.Context) WebAppOutput {
	return o
}

// Management information availability state for the app.
func (o WebAppOutput) AvailabilityState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.AvailabilityState }).(pulumi.StringOutput)
}

// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
func (o WebAppOutput) ClientAffinityEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.ClientAffinityEnabled }).(pulumi.BoolPtrOutput)
}

// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
func (o WebAppOutput) ClientCertEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.ClientCertEnabled }).(pulumi.BoolPtrOutput)
}

// Size of the function container.
func (o WebAppOutput) ContainerSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.IntPtrOutput { return v.ContainerSize }).(pulumi.IntPtrOutput)
}

// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
func (o WebAppOutput) DailyMemoryTimeQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.IntPtrOutput { return v.DailyMemoryTimeQuota }).(pulumi.IntPtrOutput)
}

// Default hostname of the app. Read-only.
func (o WebAppOutput) DefaultHostName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.DefaultHostName }).(pulumi.StringOutput)
}

// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
func (o WebAppOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
// the app is not served on those hostnames.
func (o WebAppOutput) EnabledHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringArrayOutput { return v.EnabledHostNames }).(pulumi.StringArrayOutput)
}

// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
func (o WebAppOutput) HostNameSslStates() HostNameSslStateResponseArrayOutput {
	return o.ApplyT(func(v *WebApp) HostNameSslStateResponseArrayOutput { return v.HostNameSslStates }).(HostNameSslStateResponseArrayOutput)
}

// Hostnames associated with the app.
func (o WebAppOutput) HostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringArrayOutput { return v.HostNames }).(pulumi.StringArrayOutput)
}

// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
//
//	If <code>true</code>, the app is only accessible via API management process.
func (o WebAppOutput) HostNamesDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.HostNamesDisabled }).(pulumi.BoolPtrOutput)
}

// App Service Environment to use for the app.
func (o WebAppOutput) HostingEnvironmentProfile() HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) HostingEnvironmentProfileResponsePtrOutput { return v.HostingEnvironmentProfile }).(HostingEnvironmentProfileResponsePtrOutput)
}

// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
// http requests
func (o WebAppOutput) HttpsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.HttpsOnly }).(pulumi.BoolPtrOutput)
}

// Managed service identity.
func (o WebAppOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// <code>true</code> if the app is a default container; otherwise, <code>false</code>.
func (o WebAppOutput) IsDefaultContainer() pulumi.BoolOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolOutput { return v.IsDefaultContainer }).(pulumi.BoolOutput)
}

// Kind of resource.
func (o WebAppOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// Last time the app was modified, in UTC. Read-only.
func (o WebAppOutput) LastModifiedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.LastModifiedTimeUtc }).(pulumi.StringOutput)
}

// Resource Location.
func (o WebAppOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maximum number of workers.
// This only applies to Functions container.
func (o WebAppOutput) MaxNumberOfWorkers() pulumi.IntOutput {
	return o.ApplyT(func(v *WebApp) pulumi.IntOutput { return v.MaxNumberOfWorkers }).(pulumi.IntOutput)
}

// Resource Name.
func (o WebAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
func (o WebAppOutput) OutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.OutboundIpAddresses }).(pulumi.StringOutput)
}

// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants. Read-only.
func (o WebAppOutput) PossibleOutboundIpAddresses() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.PossibleOutboundIpAddresses }).(pulumi.StringOutput)
}

// Name of the repository site.
func (o WebAppOutput) RepositorySiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.RepositorySiteName }).(pulumi.StringOutput)
}

// <code>true</code> if reserved; otherwise, <code>false</code>.
func (o WebAppOutput) Reserved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.Reserved }).(pulumi.BoolPtrOutput)
}

// Name of the resource group the app belongs to. Read-only.
func (o WebAppOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.ResourceGroup }).(pulumi.StringOutput)
}

// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
func (o WebAppOutput) ScmSiteAlsoStopped() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.BoolPtrOutput { return v.ScmSiteAlsoStopped }).(pulumi.BoolPtrOutput)
}

// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
func (o WebAppOutput) ServerFarmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringPtrOutput { return v.ServerFarmId }).(pulumi.StringPtrOutput)
}

// Configuration of the app.
func (o WebAppOutput) SiteConfig() SiteConfigResponsePtrOutput {
	return o.ApplyT(func(v *WebApp) SiteConfigResponsePtrOutput { return v.SiteConfig }).(SiteConfigResponsePtrOutput)
}

// Status of the last deployment slot swap operation.
func (o WebAppOutput) SlotSwapStatus() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v *WebApp) SlotSwapStatusResponseOutput { return v.SlotSwapStatus }).(SlotSwapStatusResponseOutput)
}

// Current state of the app.
func (o WebAppOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// App suspended till in case memory-time quota is exceeded.
func (o WebAppOutput) SuspendedTill() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.SuspendedTill }).(pulumi.StringOutput)
}

// Resource tags.
func (o WebAppOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies which deployment slot this app will swap into. Read-only.
func (o WebAppOutput) TargetSwapSlot() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.TargetSwapSlot }).(pulumi.StringOutput)
}

// Azure Traffic Manager hostnames associated with the app. Read-only.
func (o WebAppOutput) TrafficManagerHostNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringArrayOutput { return v.TrafficManagerHostNames }).(pulumi.StringArrayOutput)
}

// Resource type.
func (o WebAppOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// State indicating whether the app has exceeded its quota usage. Read-only.
func (o WebAppOutput) UsageState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebApp) pulumi.StringOutput { return v.UsageState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppOutput{})
}
