// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
func LookupAFDCustomDomain(ctx *pulumi.Context, args *LookupAFDCustomDomainArgs, opts ...pulumi.InvokeOption) (*LookupAFDCustomDomainResult, error) {
	var rv LookupAFDCustomDomainResult
	err := ctx.Invoke("azure-native:cdn/v20210601:getAFDCustomDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAFDCustomDomainArgs struct {
	CustomDomainName  string `pulumi:"customDomainName"`
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type LookupAFDCustomDomainResult struct {
	AzureDnsZone                       *ResourceReferenceResponse         `pulumi:"azureDnsZone"`
	DeploymentStatus                   string                             `pulumi:"deploymentStatus"`
	DomainValidationState              string                             `pulumi:"domainValidationState"`
	HostName                           string                             `pulumi:"hostName"`
	Id                                 string                             `pulumi:"id"`
	Name                               string                             `pulumi:"name"`
	PreValidatedCustomDomainResourceId *ResourceReferenceResponse         `pulumi:"preValidatedCustomDomainResourceId"`
	ProfileName                        string                             `pulumi:"profileName"`
	ProvisioningState                  string                             `pulumi:"provisioningState"`
	SystemData                         SystemDataResponse                 `pulumi:"systemData"`
	TlsSettings                        *AFDDomainHttpsParametersResponse  `pulumi:"tlsSettings"`
	Type                               string                             `pulumi:"type"`
	ValidationProperties               DomainValidationPropertiesResponse `pulumi:"validationProperties"`
}

func LookupAFDCustomDomainOutput(ctx *pulumi.Context, args LookupAFDCustomDomainOutputArgs, opts ...pulumi.InvokeOption) LookupAFDCustomDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAFDCustomDomainResult, error) {
			args := v.(LookupAFDCustomDomainArgs)
			r, err := LookupAFDCustomDomain(ctx, &args, opts...)
			var s LookupAFDCustomDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAFDCustomDomainResultOutput)
}

type LookupAFDCustomDomainOutputArgs struct {
	CustomDomainName  pulumi.StringInput `pulumi:"customDomainName"`
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAFDCustomDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDCustomDomainArgs)(nil)).Elem()
}

// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
type LookupAFDCustomDomainResultOutput struct{ *pulumi.OutputState }

func (LookupAFDCustomDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAFDCustomDomainResult)(nil)).Elem()
}

func (o LookupAFDCustomDomainResultOutput) ToLookupAFDCustomDomainResultOutput() LookupAFDCustomDomainResultOutput {
	return o
}

func (o LookupAFDCustomDomainResultOutput) ToLookupAFDCustomDomainResultOutputWithContext(ctx context.Context) LookupAFDCustomDomainResultOutput {
	return o
}

func (o LookupAFDCustomDomainResultOutput) AzureDnsZone() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) *ResourceReferenceResponse { return v.AzureDnsZone }).(ResourceReferenceResponsePtrOutput)
}

func (o LookupAFDCustomDomainResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) DomainValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.DomainValidationState }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) PreValidatedCustomDomainResourceId() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) *ResourceReferenceResponse {
		return v.PreValidatedCustomDomainResourceId
	}).(ResourceReferenceResponsePtrOutput)
}

func (o LookupAFDCustomDomainResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAFDCustomDomainResultOutput) TlsSettings() AFDDomainHttpsParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) *AFDDomainHttpsParametersResponse { return v.TlsSettings }).(AFDDomainHttpsParametersResponsePtrOutput)
}

func (o LookupAFDCustomDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAFDCustomDomainResultOutput) ValidationProperties() DomainValidationPropertiesResponseOutput {
	return o.ApplyT(func(v LookupAFDCustomDomainResult) DomainValidationPropertiesResponse { return v.ValidationProperties }).(DomainValidationPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAFDCustomDomainResultOutput{})
}