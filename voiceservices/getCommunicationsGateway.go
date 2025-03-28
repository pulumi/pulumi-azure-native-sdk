// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package voiceservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a CommunicationsGateway
//
// Uses Azure REST API version 2023-04-03.
//
// Other available API versions: 2023-09-01.
func LookupCommunicationsGateway(ctx *pulumi.Context, args *LookupCommunicationsGatewayArgs, opts ...pulumi.InvokeOption) (*LookupCommunicationsGatewayResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupCommunicationsGatewayResult
	err := ctx.Invoke("azure-native:voiceservices:getCommunicationsGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCommunicationsGatewayArgs struct {
	// Unique identifier for this deployment
	CommunicationsGatewayName string `pulumi:"communicationsGatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A CommunicationsGateway resource
type LookupCommunicationsGatewayResult struct {
	// Details of API bridge functionality, if required
	ApiBridge interface{} `pulumi:"apiBridge"`
	// The autogenerated label used as part of the FQDNs for accessing the Communications Gateway
	AutoGeneratedDomainNameLabel string `pulumi:"autoGeneratedDomainNameLabel"`
	// The scope at which the auto-generated domain name can be re-used
	AutoGeneratedDomainNameLabelScope *string `pulumi:"autoGeneratedDomainNameLabelScope"`
	// Voice codecs to support
	Codecs []string `pulumi:"codecs"`
	// How to connect back to the operator network, e.g. MAPS
	Connectivity string `pulumi:"connectivity"`
	// How to handle 911 calls
	E911Type string `pulumi:"e911Type"`
	// A list of dial strings used for emergency calling.
	EmergencyDialStrings []string `pulumi:"emergencyDialStrings"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Whether an integrated Mobile Control Point is in use.
	IntegratedMcpEnabled *bool `pulumi:"integratedMcpEnabled"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Whether an on-premises Mobile Control Point is in use.
	OnPremMcpEnabled *bool `pulumi:"onPremMcpEnabled"`
	// What platforms to support
	Platforms []string `pulumi:"platforms"`
	// Resource provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The regions in which to deploy the resources needed for Teams Calling
	ServiceLocations []ServiceRegionPropertiesResponse `pulumi:"serviceLocations"`
	// The current status of the deployment.
	Status string `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// This number is used in Teams Phone Mobile scenarios for access to the voicemail IVR from the native dialer.
	TeamsVoicemailPilotNumber *string `pulumi:"teamsVoicemailPilotNumber"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupCommunicationsGatewayResult
func (val *LookupCommunicationsGatewayResult) Defaults() *LookupCommunicationsGatewayResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AutoGeneratedDomainNameLabelScope == nil {
		autoGeneratedDomainNameLabelScope_ := "TenantReuse"
		tmp.AutoGeneratedDomainNameLabelScope = &autoGeneratedDomainNameLabelScope_
	}
	if tmp.IntegratedMcpEnabled == nil {
		integratedMcpEnabled_ := false
		tmp.IntegratedMcpEnabled = &integratedMcpEnabled_
	}
	if tmp.OnPremMcpEnabled == nil {
		onPremMcpEnabled_ := false
		tmp.OnPremMcpEnabled = &onPremMcpEnabled_
	}
	return &tmp
}
func LookupCommunicationsGatewayOutput(ctx *pulumi.Context, args LookupCommunicationsGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupCommunicationsGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCommunicationsGatewayResultOutput, error) {
			args := v.(LookupCommunicationsGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:voiceservices:getCommunicationsGateway", args, LookupCommunicationsGatewayResultOutput{}, options).(LookupCommunicationsGatewayResultOutput), nil
		}).(LookupCommunicationsGatewayResultOutput)
}

type LookupCommunicationsGatewayOutputArgs struct {
	// Unique identifier for this deployment
	CommunicationsGatewayName pulumi.StringInput `pulumi:"communicationsGatewayName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCommunicationsGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommunicationsGatewayArgs)(nil)).Elem()
}

// A CommunicationsGateway resource
type LookupCommunicationsGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupCommunicationsGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCommunicationsGatewayResult)(nil)).Elem()
}

func (o LookupCommunicationsGatewayResultOutput) ToLookupCommunicationsGatewayResultOutput() LookupCommunicationsGatewayResultOutput {
	return o
}

func (o LookupCommunicationsGatewayResultOutput) ToLookupCommunicationsGatewayResultOutputWithContext(ctx context.Context) LookupCommunicationsGatewayResultOutput {
	return o
}

// Details of API bridge functionality, if required
func (o LookupCommunicationsGatewayResultOutput) ApiBridge() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) interface{} { return v.ApiBridge }).(pulumi.AnyOutput)
}

// The autogenerated label used as part of the FQDNs for accessing the Communications Gateway
func (o LookupCommunicationsGatewayResultOutput) AutoGeneratedDomainNameLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.AutoGeneratedDomainNameLabel }).(pulumi.StringOutput)
}

// The scope at which the auto-generated domain name can be re-used
func (o LookupCommunicationsGatewayResultOutput) AutoGeneratedDomainNameLabelScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) *string { return v.AutoGeneratedDomainNameLabelScope }).(pulumi.StringPtrOutput)
}

// Voice codecs to support
func (o LookupCommunicationsGatewayResultOutput) Codecs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []string { return v.Codecs }).(pulumi.StringArrayOutput)
}

// How to connect back to the operator network, e.g. MAPS
func (o LookupCommunicationsGatewayResultOutput) Connectivity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Connectivity }).(pulumi.StringOutput)
}

// How to handle 911 calls
func (o LookupCommunicationsGatewayResultOutput) E911Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.E911Type }).(pulumi.StringOutput)
}

// A list of dial strings used for emergency calling.
func (o LookupCommunicationsGatewayResultOutput) EmergencyDialStrings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []string { return v.EmergencyDialStrings }).(pulumi.StringArrayOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupCommunicationsGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o LookupCommunicationsGatewayResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Whether an integrated Mobile Control Point is in use.
func (o LookupCommunicationsGatewayResultOutput) IntegratedMcpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) *bool { return v.IntegratedMcpEnabled }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o LookupCommunicationsGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupCommunicationsGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether an on-premises Mobile Control Point is in use.
func (o LookupCommunicationsGatewayResultOutput) OnPremMcpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) *bool { return v.OnPremMcpEnabled }).(pulumi.BoolPtrOutput)
}

// What platforms to support
func (o LookupCommunicationsGatewayResultOutput) Platforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []string { return v.Platforms }).(pulumi.StringArrayOutput)
}

// Resource provisioning state.
func (o LookupCommunicationsGatewayResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The regions in which to deploy the resources needed for Teams Calling
func (o LookupCommunicationsGatewayResultOutput) ServiceLocations() ServiceRegionPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) []ServiceRegionPropertiesResponse { return v.ServiceLocations }).(ServiceRegionPropertiesResponseArrayOutput)
}

// The current status of the deployment.
func (o LookupCommunicationsGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupCommunicationsGatewayResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupCommunicationsGatewayResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// This number is used in Teams Phone Mobile scenarios for access to the voicemail IVR from the native dialer.
func (o LookupCommunicationsGatewayResultOutput) TeamsVoicemailPilotNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) *string { return v.TeamsVoicemailPilotNumber }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupCommunicationsGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCommunicationsGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCommunicationsGatewayResultOutput{})
}
