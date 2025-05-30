// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the extended information of the specified manager name.
//
// Uses Azure REST API version 2017-06-01.
func LookupManagerExtendedInfo(ctx *pulumi.Context, args *LookupManagerExtendedInfoArgs, opts ...pulumi.InvokeOption) (*LookupManagerExtendedInfoResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupManagerExtendedInfoResult
	err := ctx.Invoke("azure-native:storsimple:getManagerExtendedInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerExtendedInfoArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The extended info of the manager.
type LookupManagerExtendedInfoResult struct {
	// Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
	Algorithm string `pulumi:"algorithm"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Represents the CEK of the resource.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// Represents the Cert thumbprint that was used to encrypt the CEK.
	EncryptionKeyThumbprint *string `pulumi:"encryptionKeyThumbprint"`
	// The etag of the resource.
	Etag *string `pulumi:"etag"`
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// Represents the CIK of the resource.
	IntegrityKey string `pulumi:"integrityKey"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name string `pulumi:"name"`
	// Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
	PortalCertificateThumbprint *string `pulumi:"portalCertificateThumbprint"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// The version of the extended info being persisted.
	Version *string `pulumi:"version"`
}

func LookupManagerExtendedInfoOutput(ctx *pulumi.Context, args LookupManagerExtendedInfoOutputArgs, opts ...pulumi.InvokeOption) LookupManagerExtendedInfoResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupManagerExtendedInfoResultOutput, error) {
			args := v.(LookupManagerExtendedInfoArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:storsimple:getManagerExtendedInfo", args, LookupManagerExtendedInfoResultOutput{}, options).(LookupManagerExtendedInfoResultOutput), nil
		}).(LookupManagerExtendedInfoResultOutput)
}

type LookupManagerExtendedInfoOutputArgs struct {
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagerExtendedInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerExtendedInfoArgs)(nil)).Elem()
}

// The extended info of the manager.
type LookupManagerExtendedInfoResultOutput struct{ *pulumi.OutputState }

func (LookupManagerExtendedInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerExtendedInfoResult)(nil)).Elem()
}

func (o LookupManagerExtendedInfoResultOutput) ToLookupManagerExtendedInfoResultOutput() LookupManagerExtendedInfoResultOutput {
	return o
}

func (o LookupManagerExtendedInfoResultOutput) ToLookupManagerExtendedInfoResultOutputWithContext(ctx context.Context) LookupManagerExtendedInfoResultOutput {
	return o
}

// Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
func (o LookupManagerExtendedInfoResultOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Algorithm }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupManagerExtendedInfoResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Represents the CEK of the resource.
func (o LookupManagerExtendedInfoResultOutput) EncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.EncryptionKey }).(pulumi.StringPtrOutput)
}

// Represents the Cert thumbprint that was used to encrypt the CEK.
func (o LookupManagerExtendedInfoResultOutput) EncryptionKeyThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.EncryptionKeyThumbprint }).(pulumi.StringPtrOutput)
}

// The etag of the resource.
func (o LookupManagerExtendedInfoResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The path ID that uniquely identifies the object.
func (o LookupManagerExtendedInfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Id }).(pulumi.StringOutput)
}

// Represents the CIK of the resource.
func (o LookupManagerExtendedInfoResultOutput) IntegrityKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.IntegrityKey }).(pulumi.StringOutput)
}

// The Kind of the object. Currently only Series8000 is supported
func (o LookupManagerExtendedInfoResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the object.
func (o LookupManagerExtendedInfoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Name }).(pulumi.StringOutput)
}

// Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
func (o LookupManagerExtendedInfoResultOutput) PortalCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.PortalCertificateThumbprint }).(pulumi.StringPtrOutput)
}

// The hierarchical type of the object.
func (o LookupManagerExtendedInfoResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Type }).(pulumi.StringOutput)
}

// The version of the extended info being persisted.
func (o LookupManagerExtendedInfoResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagerExtendedInfoResultOutput{})
}
