// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package communication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the Domains resource and its properties.
//
// Uses Azure REST API version 2023-06-01-preview.
//
// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:communication:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	// The name of the Domains resource.
	DomainName string `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName string `pulumi:"emailServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A class representing a Domains resource.
type LookupDomainResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The location where the Domains resource data is stored at rest.
	DataLocation string `pulumi:"dataLocation"`
	// Describes how a Domains resource is being managed.
	DomainManagement string `pulumi:"domainManagement"`
	// P2 sender domain that is displayed to the email recipients [RFC 5322].
	FromSenderDomain string `pulumi:"fromSenderDomain"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// P1 sender domain that is present on the email envelope [RFC 5321].
	MailFromSenderDomain string `pulumi:"mailFromSenderDomain"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Describes whether user engagement tracking is enabled or disabled.
	UserEngagementTracking *string `pulumi:"userEngagementTracking"`
	// List of DnsRecord
	VerificationRecords DomainPropertiesResponseVerificationRecords `pulumi:"verificationRecords"`
	// List of VerificationStatusRecord
	VerificationStates DomainPropertiesResponseVerificationStates `pulumi:"verificationStates"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDomainResultOutput, error) {
			args := v.(LookupDomainArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:communication:getDomain", args, LookupDomainResultOutput{}, options).(LookupDomainResultOutput), nil
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// The name of the Domains resource.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName pulumi.StringInput `pulumi:"emailServiceName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

// A class representing a Domains resource.
type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupDomainResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The location where the Domains resource data is stored at rest.
func (o LookupDomainResultOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DataLocation }).(pulumi.StringOutput)
}

// Describes how a Domains resource is being managed.
func (o LookupDomainResultOutput) DomainManagement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DomainManagement }).(pulumi.StringOutput)
}

// P2 sender domain that is displayed to the email recipients [RFC 5322].
func (o LookupDomainResultOutput) FromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.FromSenderDomain }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

// P1 sender domain that is present on the email envelope [RFC 5321].
func (o LookupDomainResultOutput) MailFromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.MailFromSenderDomain }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o LookupDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

// Describes whether user engagement tracking is enabled or disabled.
func (o LookupDomainResultOutput) UserEngagementTracking() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.UserEngagementTracking }).(pulumi.StringPtrOutput)
}

// List of DnsRecord
func (o LookupDomainResultOutput) VerificationRecords() DomainPropertiesResponseVerificationRecordsOutput {
	return o.ApplyT(func(v LookupDomainResult) DomainPropertiesResponseVerificationRecords { return v.VerificationRecords }).(DomainPropertiesResponseVerificationRecordsOutput)
}

// List of VerificationStatusRecord
func (o LookupDomainResultOutput) VerificationStates() DomainPropertiesResponseVerificationStatesOutput {
	return o.ApplyT(func(v LookupDomainResult) DomainPropertiesResponseVerificationStates { return v.VerificationStates }).(DomainPropertiesResponseVerificationStatesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
