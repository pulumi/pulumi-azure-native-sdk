// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A lab.
//
// Deprecated: Version 2016-05-15 will be removed in v2 of the provider.
func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	Expand            *string `pulumi:"expand"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}

// A lab.
type LookupLabResult struct {
	ArtifactsStorageAccount       string            `pulumi:"artifactsStorageAccount"`
	CreatedDate                   string            `pulumi:"createdDate"`
	DefaultPremiumStorageAccount  string            `pulumi:"defaultPremiumStorageAccount"`
	DefaultStorageAccount         string            `pulumi:"defaultStorageAccount"`
	Id                            string            `pulumi:"id"`
	LabStorageType                *string           `pulumi:"labStorageType"`
	Location                      *string           `pulumi:"location"`
	Name                          string            `pulumi:"name"`
	PremiumDataDiskStorageAccount string            `pulumi:"premiumDataDiskStorageAccount"`
	PremiumDataDisks              *string           `pulumi:"premiumDataDisks"`
	ProvisioningState             *string           `pulumi:"provisioningState"`
	Tags                          map[string]string `pulumi:"tags"`
	Type                          string            `pulumi:"type"`
	UniqueIdentifier              *string           `pulumi:"uniqueIdentifier"`
	VaultName                     string            `pulumi:"vaultName"`
}

func LookupLabOutput(ctx *pulumi.Context, args LookupLabOutputArgs, opts ...pulumi.InvokeOption) LookupLabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResult, error) {
			args := v.(LookupLabArgs)
			r, err := LookupLab(ctx, &args, opts...)
			var s LookupLabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResultOutput)
}

type LookupLabOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupLabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabArgs)(nil)).Elem()
}

// A lab.
type LookupLabResultOutput struct{ *pulumi.OutputState }

func (LookupLabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResult)(nil)).Elem()
}

func (o LookupLabResultOutput) ToLookupLabResultOutput() LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToLookupLabResultOutputWithContext(ctx context.Context) LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ArtifactsStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.ArtifactsStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) DefaultPremiumStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.DefaultPremiumStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) DefaultStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.DefaultStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) LabStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.LabStorageType }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) PremiumDataDiskStorageAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.PremiumDataDiskStorageAccount }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) PremiumDataDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.PremiumDataDisks }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLabResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLabResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.VaultName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResultOutput{})
}