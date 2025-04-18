// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a record set.
//
// Uses Azure REST API version 2023-07-01-preview.
//
// Other available API versions: 2015-05-04-preview, 2016-04-01, 2017-09-01, 2017-10-01, 2018-03-01-preview, 2018-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dns [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupRecordSet(ctx *pulumi.Context, args *LookupRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupRecordSetResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupRecordSetResult
	err := ctx.Invoke("azure-native:dns:getRecordSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRecordSetArgs struct {
	// The type of DNS record in this record set.
	RecordType string `pulumi:"recordType"`
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName string `pulumi:"relativeRecordSetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName string `pulumi:"zoneName"`
}

// Describes a DNS record set (a collection of DNS records with the same name and type).
type LookupRecordSetResult struct {
	// The list of A records in the record set.
	ARecords []ARecordResponse `pulumi:"aRecords"`
	// The list of AAAA records in the record set.
	AaaaRecords []AaaaRecordResponse `pulumi:"aaaaRecords"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The list of CAA records in the record set.
	CaaRecords []CaaRecordResponse `pulumi:"caaRecords"`
	// The CNAME record in the  record set.
	CnameRecord *CnameRecordResponse `pulumi:"cnameRecord"`
	// The list of DS records in the record set.
	DsRecords []DsRecordResponse `pulumi:"dsRecords"`
	// The etag of the record set.
	Etag *string `pulumi:"etag"`
	// Fully qualified domain name of the record set.
	Fqdn string `pulumi:"fqdn"`
	// The ID of the record set.
	Id string `pulumi:"id"`
	// The metadata attached to the record set.
	Metadata map[string]string `pulumi:"metadata"`
	// The list of MX records in the record set.
	MxRecords []MxRecordResponse `pulumi:"mxRecords"`
	// The name of the record set.
	Name string `pulumi:"name"`
	// The list of NAPTR records in the record set.
	NaptrRecords []NaptrRecordResponse `pulumi:"naptrRecords"`
	// The list of NS records in the record set.
	NsRecords []NsRecordResponse `pulumi:"nsRecords"`
	// provisioning State of the record set.
	ProvisioningState string `pulumi:"provisioningState"`
	// The list of PTR records in the record set.
	PtrRecords []PtrRecordResponse `pulumi:"ptrRecords"`
	// The SOA record in the record set.
	SoaRecord *SoaRecordResponse `pulumi:"soaRecord"`
	// The list of SRV records in the record set.
	SrvRecords []SrvRecordResponse `pulumi:"srvRecords"`
	// A reference to an azure resource from where the dns resource value is taken.
	TargetResource *SubResourceResponse `pulumi:"targetResource"`
	// The list of TLSA records in the record set.
	TlsaRecords []TlsaRecordResponse `pulumi:"tlsaRecords"`
	// A reference to an azure traffic manager profile resource from where the dns resource value is taken.
	TrafficManagementProfile *SubResourceResponse `pulumi:"trafficManagementProfile"`
	// The TTL (time-to-live) of the records in the record set.
	Ttl *float64 `pulumi:"ttl"`
	// The list of TXT records in the record set.
	TxtRecords []TxtRecordResponse `pulumi:"txtRecords"`
	// The type of the record set.
	Type string `pulumi:"type"`
}

func LookupRecordSetOutput(ctx *pulumi.Context, args LookupRecordSetOutputArgs, opts ...pulumi.InvokeOption) LookupRecordSetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRecordSetResultOutput, error) {
			args := v.(LookupRecordSetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:dns:getRecordSet", args, LookupRecordSetResultOutput{}, options).(LookupRecordSetResultOutput), nil
		}).(LookupRecordSetResultOutput)
}

type LookupRecordSetOutputArgs struct {
	// The type of DNS record in this record set.
	RecordType pulumi.StringInput `pulumi:"recordType"`
	// The name of the record set, relative to the name of the zone.
	RelativeRecordSetName pulumi.StringInput `pulumi:"relativeRecordSetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the DNS zone (without a terminating dot).
	ZoneName pulumi.StringInput `pulumi:"zoneName"`
}

func (LookupRecordSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordSetArgs)(nil)).Elem()
}

// Describes a DNS record set (a collection of DNS records with the same name and type).
type LookupRecordSetResultOutput struct{ *pulumi.OutputState }

func (LookupRecordSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRecordSetResult)(nil)).Elem()
}

func (o LookupRecordSetResultOutput) ToLookupRecordSetResultOutput() LookupRecordSetResultOutput {
	return o
}

func (o LookupRecordSetResultOutput) ToLookupRecordSetResultOutputWithContext(ctx context.Context) LookupRecordSetResultOutput {
	return o
}

// The list of A records in the record set.
func (o LookupRecordSetResultOutput) ARecords() ARecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []ARecordResponse { return v.ARecords }).(ARecordResponseArrayOutput)
}

// The list of AAAA records in the record set.
func (o LookupRecordSetResultOutput) AaaaRecords() AaaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []AaaaRecordResponse { return v.AaaaRecords }).(AaaaRecordResponseArrayOutput)
}

// The Azure API version of the resource.
func (o LookupRecordSetResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The list of CAA records in the record set.
func (o LookupRecordSetResultOutput) CaaRecords() CaaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []CaaRecordResponse { return v.CaaRecords }).(CaaRecordResponseArrayOutput)
}

// The CNAME record in the  record set.
func (o LookupRecordSetResultOutput) CnameRecord() CnameRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *CnameRecordResponse { return v.CnameRecord }).(CnameRecordResponsePtrOutput)
}

// The list of DS records in the record set.
func (o LookupRecordSetResultOutput) DsRecords() DsRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []DsRecordResponse { return v.DsRecords }).(DsRecordResponseArrayOutput)
}

// The etag of the record set.
func (o LookupRecordSetResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Fully qualified domain name of the record set.
func (o LookupRecordSetResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// The ID of the record set.
func (o LookupRecordSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Id }).(pulumi.StringOutput)
}

// The metadata attached to the record set.
func (o LookupRecordSetResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRecordSetResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The list of MX records in the record set.
func (o LookupRecordSetResultOutput) MxRecords() MxRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []MxRecordResponse { return v.MxRecords }).(MxRecordResponseArrayOutput)
}

// The name of the record set.
func (o LookupRecordSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of NAPTR records in the record set.
func (o LookupRecordSetResultOutput) NaptrRecords() NaptrRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []NaptrRecordResponse { return v.NaptrRecords }).(NaptrRecordResponseArrayOutput)
}

// The list of NS records in the record set.
func (o LookupRecordSetResultOutput) NsRecords() NsRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []NsRecordResponse { return v.NsRecords }).(NsRecordResponseArrayOutput)
}

// provisioning State of the record set.
func (o LookupRecordSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The list of PTR records in the record set.
func (o LookupRecordSetResultOutput) PtrRecords() PtrRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []PtrRecordResponse { return v.PtrRecords }).(PtrRecordResponseArrayOutput)
}

// The SOA record in the record set.
func (o LookupRecordSetResultOutput) SoaRecord() SoaRecordResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *SoaRecordResponse { return v.SoaRecord }).(SoaRecordResponsePtrOutput)
}

// The list of SRV records in the record set.
func (o LookupRecordSetResultOutput) SrvRecords() SrvRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []SrvRecordResponse { return v.SrvRecords }).(SrvRecordResponseArrayOutput)
}

// A reference to an azure resource from where the dns resource value is taken.
func (o LookupRecordSetResultOutput) TargetResource() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *SubResourceResponse { return v.TargetResource }).(SubResourceResponsePtrOutput)
}

// The list of TLSA records in the record set.
func (o LookupRecordSetResultOutput) TlsaRecords() TlsaRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []TlsaRecordResponse { return v.TlsaRecords }).(TlsaRecordResponseArrayOutput)
}

// A reference to an azure traffic manager profile resource from where the dns resource value is taken.
func (o LookupRecordSetResultOutput) TrafficManagementProfile() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *SubResourceResponse { return v.TrafficManagementProfile }).(SubResourceResponsePtrOutput)
}

// The TTL (time-to-live) of the records in the record set.
func (o LookupRecordSetResultOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRecordSetResult) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

// The list of TXT records in the record set.
func (o LookupRecordSetResultOutput) TxtRecords() TxtRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupRecordSetResult) []TxtRecordResponse { return v.TxtRecords }).(TxtRecordResponseArrayOutput)
}

// The type of the record set.
func (o LookupRecordSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRecordSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRecordSetResultOutput{})
}
