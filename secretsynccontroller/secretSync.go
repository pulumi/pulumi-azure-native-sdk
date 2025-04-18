// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretsynccontroller

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The SecretSync resource.
//
// Uses Azure REST API version 2024-08-21-preview. In version 2.x of the Azure Native provider, it used API version 2024-08-21-preview.
type SecretSync struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The complex type of the extended location.
	ExtendedLocation AzureResourceManagerCommonTypesExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// ForceSynchronization can be used to force the secret synchronization. The secret synchronization is triggered by changing the value in this field. This field is not used to resolve synchronization conflicts.
	ForceSynchronization pulumi.StringPtrOutput `pulumi:"forceSynchronization"`
	// Type specifies the type of the Kubernetes secret object, e.g. "Opaque" or"kubernetes.io/tls". The controller must have permission to create secrets of the specified type.
	KubernetesSecretType pulumi.StringOutput `pulumi:"kubernetesSecretType"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of SecretObjectData that maps secret data from the external secret provider to the Kubernetes secret. Each entry specifies the source secret in the external provider and the corresponding key in the Kubernetes secret.
	ObjectSecretMapping KubernetesSecretObjectMappingResponseArrayOutput `pulumi:"objectSecretMapping"`
	// Provisioning state of the SecretSync instance.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SecretProviderClassName specifies the name of the SecretProviderClass resource, which contains the information needed to access the cloud provider secret store.
	SecretProviderClassName pulumi.StringOutput `pulumi:"secretProviderClassName"`
	// ServiceAccountName specifies the name of the service account used to access the cloud provider secret store. The audience field in the service account token must be passed as parameter in the controller configuration. The audience is used when requesting a token from the API server for the service account; the supported audiences are defined by each provider.
	ServiceAccountName pulumi.StringOutput `pulumi:"serviceAccountName"`
	// SecretSyncStatus defines the observed state of the secret synchronization process.
	Status SecretSyncStatusResponseOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecretSync registers a new resource with the given unique name, arguments, and options.
func NewSecretSync(ctx *pulumi.Context,
	name string, args *SecretSyncArgs, opts ...pulumi.ResourceOption) (*SecretSync, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KubernetesSecretType == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesSecretType'")
	}
	if args.ObjectSecretMapping == nil {
		return nil, errors.New("invalid value for required argument 'ObjectSecretMapping'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SecretProviderClassName == nil {
		return nil, errors.New("invalid value for required argument 'SecretProviderClassName'")
	}
	if args.ServiceAccountName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:secretsynccontroller/v20240821preview:SecretSync"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecretSync
	err := ctx.RegisterResource("azure-native:secretsynccontroller:SecretSync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretSync gets an existing SecretSync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretSyncState, opts ...pulumi.ResourceOption) (*SecretSync, error) {
	var resource SecretSync
	err := ctx.ReadResource("azure-native:secretsynccontroller:SecretSync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretSync resources.
type secretSyncState struct {
}

type SecretSyncState struct {
}

func (SecretSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretSyncState)(nil)).Elem()
}

type secretSyncArgs struct {
	// The complex type of the extended location.
	ExtendedLocation *AzureResourceManagerCommonTypesExtendedLocation `pulumi:"extendedLocation"`
	// ForceSynchronization can be used to force the secret synchronization. The secret synchronization is triggered by changing the value in this field. This field is not used to resolve synchronization conflicts.
	ForceSynchronization *string `pulumi:"forceSynchronization"`
	// Type specifies the type of the Kubernetes secret object, e.g. "Opaque" or"kubernetes.io/tls". The controller must have permission to create secrets of the specified type.
	KubernetesSecretType string `pulumi:"kubernetesSecretType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// An array of SecretObjectData that maps secret data from the external secret provider to the Kubernetes secret. Each entry specifies the source secret in the external provider and the corresponding key in the Kubernetes secret.
	ObjectSecretMapping []KubernetesSecretObjectMapping `pulumi:"objectSecretMapping"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SecretProviderClassName specifies the name of the SecretProviderClass resource, which contains the information needed to access the cloud provider secret store.
	SecretProviderClassName string `pulumi:"secretProviderClassName"`
	// The name of the SecretSync
	SecretSyncName *string `pulumi:"secretSyncName"`
	// ServiceAccountName specifies the name of the service account used to access the cloud provider secret store. The audience field in the service account token must be passed as parameter in the controller configuration. The audience is used when requesting a token from the API server for the service account; the supported audiences are defined by each provider.
	ServiceAccountName string `pulumi:"serviceAccountName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SecretSync resource.
type SecretSyncArgs struct {
	// The complex type of the extended location.
	ExtendedLocation AzureResourceManagerCommonTypesExtendedLocationPtrInput
	// ForceSynchronization can be used to force the secret synchronization. The secret synchronization is triggered by changing the value in this field. This field is not used to resolve synchronization conflicts.
	ForceSynchronization pulumi.StringPtrInput
	// Type specifies the type of the Kubernetes secret object, e.g. "Opaque" or"kubernetes.io/tls". The controller must have permission to create secrets of the specified type.
	KubernetesSecretType pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// An array of SecretObjectData that maps secret data from the external secret provider to the Kubernetes secret. Each entry specifies the source secret in the external provider and the corresponding key in the Kubernetes secret.
	ObjectSecretMapping KubernetesSecretObjectMappingArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// SecretProviderClassName specifies the name of the SecretProviderClass resource, which contains the information needed to access the cloud provider secret store.
	SecretProviderClassName pulumi.StringInput
	// The name of the SecretSync
	SecretSyncName pulumi.StringPtrInput
	// ServiceAccountName specifies the name of the service account used to access the cloud provider secret store. The audience field in the service account token must be passed as parameter in the controller configuration. The audience is used when requesting a token from the API server for the service account; the supported audiences are defined by each provider.
	ServiceAccountName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SecretSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretSyncArgs)(nil)).Elem()
}

type SecretSyncInput interface {
	pulumi.Input

	ToSecretSyncOutput() SecretSyncOutput
	ToSecretSyncOutputWithContext(ctx context.Context) SecretSyncOutput
}

func (*SecretSync) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretSync)(nil)).Elem()
}

func (i *SecretSync) ToSecretSyncOutput() SecretSyncOutput {
	return i.ToSecretSyncOutputWithContext(context.Background())
}

func (i *SecretSync) ToSecretSyncOutputWithContext(ctx context.Context) SecretSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretSyncOutput)
}

type SecretSyncOutput struct{ *pulumi.OutputState }

func (SecretSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretSync)(nil)).Elem()
}

func (o SecretSyncOutput) ToSecretSyncOutput() SecretSyncOutput {
	return o
}

func (o SecretSyncOutput) ToSecretSyncOutputWithContext(ctx context.Context) SecretSyncOutput {
	return o
}

// The Azure API version of the resource.
func (o SecretSyncOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The complex type of the extended location.
func (o SecretSyncOutput) ExtendedLocation() AzureResourceManagerCommonTypesExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *SecretSync) AzureResourceManagerCommonTypesExtendedLocationResponsePtrOutput {
		return v.ExtendedLocation
	}).(AzureResourceManagerCommonTypesExtendedLocationResponsePtrOutput)
}

// ForceSynchronization can be used to force the secret synchronization. The secret synchronization is triggered by changing the value in this field. This field is not used to resolve synchronization conflicts.
func (o SecretSyncOutput) ForceSynchronization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringPtrOutput { return v.ForceSynchronization }).(pulumi.StringPtrOutput)
}

// Type specifies the type of the Kubernetes secret object, e.g. "Opaque" or"kubernetes.io/tls". The controller must have permission to create secrets of the specified type.
func (o SecretSyncOutput) KubernetesSecretType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.KubernetesSecretType }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o SecretSyncOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SecretSyncOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of SecretObjectData that maps secret data from the external secret provider to the Kubernetes secret. Each entry specifies the source secret in the external provider and the corresponding key in the Kubernetes secret.
func (o SecretSyncOutput) ObjectSecretMapping() KubernetesSecretObjectMappingResponseArrayOutput {
	return o.ApplyT(func(v *SecretSync) KubernetesSecretObjectMappingResponseArrayOutput { return v.ObjectSecretMapping }).(KubernetesSecretObjectMappingResponseArrayOutput)
}

// Provisioning state of the SecretSync instance.
func (o SecretSyncOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// SecretProviderClassName specifies the name of the SecretProviderClass resource, which contains the information needed to access the cloud provider secret store.
func (o SecretSyncOutput) SecretProviderClassName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.SecretProviderClassName }).(pulumi.StringOutput)
}

// ServiceAccountName specifies the name of the service account used to access the cloud provider secret store. The audience field in the service account token must be passed as parameter in the controller configuration. The audience is used when requesting a token from the API server for the service account; the supported audiences are defined by each provider.
func (o SecretSyncOutput) ServiceAccountName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.ServiceAccountName }).(pulumi.StringOutput)
}

// SecretSyncStatus defines the observed state of the secret synchronization process.
func (o SecretSyncOutput) Status() SecretSyncStatusResponseOutput {
	return o.ApplyT(func(v *SecretSync) SecretSyncStatusResponseOutput { return v.Status }).(SecretSyncStatusResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SecretSyncOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecretSync) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SecretSyncOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SecretSyncOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretSync) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretSyncOutput{})
}
