// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eventgrid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get properties of a partner destination.
//
// Uses Azure REST API version 2024-12-15-preview.
//
// Other available API versions: 2023-06-01-preview, 2023-12-15-preview, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native eventgrid [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupPartnerDestination(ctx *pulumi.Context, args *LookupPartnerDestinationArgs, opts ...pulumi.InvokeOption) (*LookupPartnerDestinationResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPartnerDestinationResult
	err := ctx.Invoke("azure-native:eventgrid:getPartnerDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerDestinationArgs struct {
	// Name of the partner destination.
	PartnerDestinationName string `pulumi:"partnerDestinationName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Event Grid Partner Destination.
type LookupPartnerDestinationResult struct {
	// Activation state of the partner destination.
	ActivationState *string `pulumi:"activationState"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Endpoint Base URL of the partner destination
	EndpointBaseUrl *string `pulumi:"endpointBaseUrl"`
	// Endpoint context associated with this partner destination.
	EndpointServiceContext *string `pulumi:"endpointServiceContext"`
	// Expiration time of the partner destination. If this timer expires and the partner destination was never activated,
	// the partner destination and corresponding channel are deleted.
	ExpirationTimeIfNotActivatedUtc *string `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Fully qualified identifier of the resource.
	Id string `pulumi:"id"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Context or helpful message that can be used during the approval process.
	MessageForActivation *string `pulumi:"messageForActivation"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// The immutable Id of the corresponding partner registration.
	PartnerRegistrationImmutableId *string `pulumi:"partnerRegistrationImmutableId"`
	// Provisioning state of the partner destination.
	ProvisioningState string `pulumi:"provisioningState"`
	// The system metadata relating to the Event Grid resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource.
	Type string `pulumi:"type"`
}

func LookupPartnerDestinationOutput(ctx *pulumi.Context, args LookupPartnerDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupPartnerDestinationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPartnerDestinationResultOutput, error) {
			args := v.(LookupPartnerDestinationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:eventgrid:getPartnerDestination", args, LookupPartnerDestinationResultOutput{}, options).(LookupPartnerDestinationResultOutput), nil
		}).(LookupPartnerDestinationResultOutput)
}

type LookupPartnerDestinationOutputArgs struct {
	// Name of the partner destination.
	PartnerDestinationName pulumi.StringInput `pulumi:"partnerDestinationName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPartnerDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerDestinationArgs)(nil)).Elem()
}

// Event Grid Partner Destination.
type LookupPartnerDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupPartnerDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPartnerDestinationResult)(nil)).Elem()
}

func (o LookupPartnerDestinationResultOutput) ToLookupPartnerDestinationResultOutput() LookupPartnerDestinationResultOutput {
	return o
}

func (o LookupPartnerDestinationResultOutput) ToLookupPartnerDestinationResultOutputWithContext(ctx context.Context) LookupPartnerDestinationResultOutput {
	return o
}

// Activation state of the partner destination.
func (o LookupPartnerDestinationResultOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.ActivationState }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o LookupPartnerDestinationResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Endpoint Base URL of the partner destination
func (o LookupPartnerDestinationResultOutput) EndpointBaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.EndpointBaseUrl }).(pulumi.StringPtrOutput)
}

// Endpoint context associated with this partner destination.
func (o LookupPartnerDestinationResultOutput) EndpointServiceContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.EndpointServiceContext }).(pulumi.StringPtrOutput)
}

// Expiration time of the partner destination. If this timer expires and the partner destination was never activated,
// the partner destination and corresponding channel are deleted.
func (o LookupPartnerDestinationResultOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

// Fully qualified identifier of the resource.
func (o LookupPartnerDestinationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Location of the resource.
func (o LookupPartnerDestinationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Location }).(pulumi.StringOutput)
}

// Context or helpful message that can be used during the approval process.
func (o LookupPartnerDestinationResultOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupPartnerDestinationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Name }).(pulumi.StringOutput)
}

// The immutable Id of the corresponding partner registration.
func (o LookupPartnerDestinationResultOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) *string { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

// Provisioning state of the partner destination.
func (o LookupPartnerDestinationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The system metadata relating to the Event Grid resource.
func (o LookupPartnerDestinationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Tags of the resource.
func (o LookupPartnerDestinationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the resource.
func (o LookupPartnerDestinationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPartnerDestinationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPartnerDestinationResultOutput{})
}
