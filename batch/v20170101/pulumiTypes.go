// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An application package which represents a particular version of an application.
type ApplicationPackageResponse struct {
	// The format of the application package, if the package is active.
	Format string `pulumi:"format"`
	// The ID of the application.
	Id string `pulumi:"id"`
	// The time at which the package was last activated, if the package is active.
	LastActivationTime string `pulumi:"lastActivationTime"`
	// The current state of the application package.
	State string `pulumi:"state"`
	// The storage URL at which the application package is stored.
	StorageUrl string `pulumi:"storageUrl"`
	// The UTC time at which the storage URL will expire.
	StorageUrlExpiry string `pulumi:"storageUrlExpiry"`
	// The version of the application package.
	Version string `pulumi:"version"`
}

// An application package which represents a particular version of an application.
type ApplicationPackageResponseOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutput() ApplicationPackageResponseOutput {
	return o
}

func (o ApplicationPackageResponseOutput) ToApplicationPackageResponseOutputWithContext(ctx context.Context) ApplicationPackageResponseOutput {
	return o
}

// The format of the application package, if the package is active.
func (o ApplicationPackageResponseOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.Format }).(pulumi.StringOutput)
}

// The ID of the application.
func (o ApplicationPackageResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The time at which the package was last activated, if the package is active.
func (o ApplicationPackageResponseOutput) LastActivationTime() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.LastActivationTime }).(pulumi.StringOutput)
}

// The current state of the application package.
func (o ApplicationPackageResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.State }).(pulumi.StringOutput)
}

// The storage URL at which the application package is stored.
func (o ApplicationPackageResponseOutput) StorageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.StorageUrl }).(pulumi.StringOutput)
}

// The UTC time at which the storage URL will expire.
func (o ApplicationPackageResponseOutput) StorageUrlExpiry() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.StorageUrlExpiry }).(pulumi.StringOutput)
}

// The version of the application package.
func (o ApplicationPackageResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ApplicationPackageResponse) string { return v.Version }).(pulumi.StringOutput)
}

type ApplicationPackageResponseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationPackageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationPackageResponse)(nil)).Elem()
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutput() ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) ToApplicationPackageResponseArrayOutputWithContext(ctx context.Context) ApplicationPackageResponseArrayOutput {
	return o
}

func (o ApplicationPackageResponseArrayOutput) Index(i pulumi.IntInput) ApplicationPackageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationPackageResponse {
		return vs[0].([]ApplicationPackageResponse)[vs[1].(int)]
	}).(ApplicationPackageResponseOutput)
}

// The properties related to auto storage account.
type AutoStorageBaseProperties struct {
	// The resource ID of the storage account to be used for auto storage account.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// AutoStorageBasePropertiesInput is an input type that accepts AutoStorageBasePropertiesArgs and AutoStorageBasePropertiesOutput values.
// You can construct a concrete instance of `AutoStorageBasePropertiesInput` via:
//
//	AutoStorageBasePropertiesArgs{...}
type AutoStorageBasePropertiesInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput
	ToAutoStorageBasePropertiesOutputWithContext(context.Context) AutoStorageBasePropertiesOutput
}

// The properties related to auto storage account.
type AutoStorageBasePropertiesArgs struct {
	// The resource ID of the storage account to be used for auto storage account.
	StorageAccountId pulumi.StringInput `pulumi:"storageAccountId"`
}

func (AutoStorageBasePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return i.ToAutoStorageBasePropertiesOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput)
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i AutoStorageBasePropertiesArgs) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesOutput).ToAutoStorageBasePropertiesPtrOutputWithContext(ctx)
}

// AutoStorageBasePropertiesPtrInput is an input type that accepts AutoStorageBasePropertiesArgs, AutoStorageBasePropertiesPtr and AutoStorageBasePropertiesPtrOutput values.
// You can construct a concrete instance of `AutoStorageBasePropertiesPtrInput` via:
//
//	        AutoStorageBasePropertiesArgs{...}
//
//	or:
//
//	        nil
type AutoStorageBasePropertiesPtrInput interface {
	pulumi.Input

	ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput
	ToAutoStorageBasePropertiesPtrOutputWithContext(context.Context) AutoStorageBasePropertiesPtrOutput
}

type autoStorageBasePropertiesPtrType AutoStorageBasePropertiesArgs

func AutoStorageBasePropertiesPtr(v *AutoStorageBasePropertiesArgs) AutoStorageBasePropertiesPtrInput {
	return (*autoStorageBasePropertiesPtrType)(v)
}

func (*autoStorageBasePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return i.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (i *autoStorageBasePropertiesPtrType) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoStorageBasePropertiesPtrOutput)
}

// The properties related to auto storage account.
type AutoStorageBasePropertiesOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutput() AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesOutputWithContext(ctx context.Context) AutoStorageBasePropertiesOutput {
	return o
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o.ToAutoStorageBasePropertiesPtrOutputWithContext(context.Background())
}

func (o AutoStorageBasePropertiesOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoStorageBaseProperties) *AutoStorageBaseProperties {
		return &v
	}).(AutoStorageBasePropertiesPtrOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStorageBasePropertiesOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStorageBaseProperties) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

type AutoStorageBasePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AutoStorageBasePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoStorageBaseProperties)(nil)).Elem()
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutput() AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) ToAutoStorageBasePropertiesPtrOutputWithContext(ctx context.Context) AutoStorageBasePropertiesPtrOutput {
	return o
}

func (o AutoStorageBasePropertiesPtrOutput) Elem() AutoStorageBasePropertiesOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) AutoStorageBaseProperties {
		if v != nil {
			return *v
		}
		var ret AutoStorageBaseProperties
		return ret
	}).(AutoStorageBasePropertiesOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStorageBasePropertiesPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoStorageBaseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

// Contains information about the auto storage account associated with a Batch account.
type AutoStoragePropertiesResponse struct {
	// The UTC time at which storage keys were last synchronized with the Batch account.
	LastKeySync string `pulumi:"lastKeySync"`
	// The resource ID of the storage account to be used for auto storage account.
	StorageAccountId string `pulumi:"storageAccountId"`
}

// Contains information about the auto storage account associated with a Batch account.
type AutoStoragePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutoStoragePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoStoragePropertiesResponse)(nil)).Elem()
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutput() AutoStoragePropertiesResponseOutput {
	return o
}

func (o AutoStoragePropertiesResponseOutput) ToAutoStoragePropertiesResponseOutputWithContext(ctx context.Context) AutoStoragePropertiesResponseOutput {
	return o
}

// The UTC time at which storage keys were last synchronized with the Batch account.
func (o AutoStoragePropertiesResponseOutput) LastKeySync() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.LastKeySync }).(pulumi.StringOutput)
}

// The resource ID of the storage account to be used for auto storage account.
func (o AutoStoragePropertiesResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v AutoStoragePropertiesResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReference struct {
	// The resource ID of the Azure key vault associated with the Batch account.
	Id string `pulumi:"id"`
	// The Url of the Azure key vault associated with the Batch account.
	Url string `pulumi:"url"`
}

// KeyVaultReferenceInput is an input type that accepts KeyVaultReferenceArgs and KeyVaultReferenceOutput values.
// You can construct a concrete instance of `KeyVaultReferenceInput` via:
//
//	KeyVaultReferenceArgs{...}
type KeyVaultReferenceInput interface {
	pulumi.Input

	ToKeyVaultReferenceOutput() KeyVaultReferenceOutput
	ToKeyVaultReferenceOutputWithContext(context.Context) KeyVaultReferenceOutput
}

// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReferenceArgs struct {
	// The resource ID of the Azure key vault associated with the Batch account.
	Id pulumi.StringInput `pulumi:"id"`
	// The Url of the Azure key vault associated with the Batch account.
	Url pulumi.StringInput `pulumi:"url"`
}

func (KeyVaultReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return i.ToKeyVaultReferenceOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput)
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultReferenceArgs) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferenceOutput).ToKeyVaultReferencePtrOutputWithContext(ctx)
}

// KeyVaultReferencePtrInput is an input type that accepts KeyVaultReferenceArgs, KeyVaultReferencePtr and KeyVaultReferencePtrOutput values.
// You can construct a concrete instance of `KeyVaultReferencePtrInput` via:
//
//	        KeyVaultReferenceArgs{...}
//
//	or:
//
//	        nil
type KeyVaultReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput
	ToKeyVaultReferencePtrOutputWithContext(context.Context) KeyVaultReferencePtrOutput
}

type keyVaultReferencePtrType KeyVaultReferenceArgs

func KeyVaultReferencePtr(v *KeyVaultReferenceArgs) KeyVaultReferencePtrInput {
	return (*keyVaultReferencePtrType)(v)
}

func (*keyVaultReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return i.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultReferencePtrType) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultReferencePtrOutput)
}

// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutput() KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferenceOutputWithContext(ctx context.Context) KeyVaultReferenceOutput {
	return o
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o.ToKeyVaultReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultReferenceOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultReference) *KeyVaultReference {
		return &v
	}).(KeyVaultReferencePtrOutput)
}

// The resource ID of the Azure key vault associated with the Batch account.
func (o KeyVaultReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Id }).(pulumi.StringOutput)
}

// The Url of the Azure key vault associated with the Batch account.
func (o KeyVaultReferenceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReference) string { return v.Url }).(pulumi.StringOutput)
}

type KeyVaultReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultReference)(nil)).Elem()
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutput() KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) ToKeyVaultReferencePtrOutputWithContext(ctx context.Context) KeyVaultReferencePtrOutput {
	return o
}

func (o KeyVaultReferencePtrOutput) Elem() KeyVaultReferenceOutput {
	return o.ApplyT(func(v *KeyVaultReference) KeyVaultReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultReference
		return ret
	}).(KeyVaultReferenceOutput)
}

// The resource ID of the Azure key vault associated with the Batch account.
func (o KeyVaultReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// The Url of the Azure key vault associated with the Batch account.
func (o KeyVaultReferencePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultReference) *string {
		if v == nil {
			return nil
		}
		return &v.Url
	}).(pulumi.StringPtrOutput)
}

// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReferenceResponse struct {
	// The resource ID of the Azure key vault associated with the Batch account.
	Id string `pulumi:"id"`
	// The Url of the Azure key vault associated with the Batch account.
	Url string `pulumi:"url"`
}

// Identifies the Azure key vault associated with a Batch account.
type KeyVaultReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultReferenceResponse)(nil)).Elem()
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutput() KeyVaultReferenceResponseOutput {
	return o
}

func (o KeyVaultReferenceResponseOutput) ToKeyVaultReferenceResponseOutputWithContext(ctx context.Context) KeyVaultReferenceResponseOutput {
	return o
}

// The resource ID of the Azure key vault associated with the Batch account.
func (o KeyVaultReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The Url of the Azure key vault associated with the Batch account.
func (o KeyVaultReferenceResponseOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultReferenceResponse) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationPackageResponseOutput{})
	pulumi.RegisterOutputType(ApplicationPackageResponseArrayOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesOutput{})
	pulumi.RegisterOutputType(AutoStorageBasePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AutoStoragePropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultReferenceResponseOutput{})
}