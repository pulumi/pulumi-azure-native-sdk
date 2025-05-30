// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcompute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a license profile in a hybrid machine.
//
// Uses Azure REST API version 2024-07-10. In version 2.x of the Azure Native provider, it used API version 2023-06-20-preview.
//
// Other available API versions: 2023-06-20-preview, 2023-10-03-preview, 2024-03-31-preview, 2024-05-20-preview, 2024-07-31-preview, 2024-09-10-preview, 2024-11-10-preview, 2025-01-13, 2025-02-19-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native hybridcompute [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type LicenseProfile struct {
	pulumi.CustomResourceState

	// The resource id of the license.
	AssignedLicense pulumi.StringPtrOutput `pulumi:"assignedLicense"`
	// The guid id of the license.
	AssignedLicenseImmutableId pulumi.StringOutput `pulumi:"assignedLicenseImmutableId"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The timestamp in UTC when the billing ends.
	BillingEndDate pulumi.StringOutput `pulumi:"billingEndDate"`
	// The timestamp in UTC when the billing starts.
	BillingStartDate pulumi.StringOutput `pulumi:"billingStartDate"`
	// The timestamp in UTC when the user disenrolled the feature.
	DisenrollmentDate pulumi.StringOutput `pulumi:"disenrollmentDate"`
	// The timestamp in UTC when the user enrolls the feature.
	EnrollmentDate pulumi.StringOutput `pulumi:"enrollmentDate"`
	// The errors that were encountered during the feature enrollment or disenrollment.
	Error ErrorDetailResponseOutput `pulumi:"error"`
	// Indicates the eligibility state of Esu.
	EsuEligibility pulumi.StringOutput `pulumi:"esuEligibility"`
	// Indicates whether there is an ESU Key currently active for the machine.
	EsuKeyState pulumi.StringOutput `pulumi:"esuKeyState"`
	// The list of ESU keys.
	EsuKeys EsuKeyResponseArrayOutput `pulumi:"esuKeys"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of product features.
	ProductFeatures ProductFeatureResponseArrayOutput `pulumi:"productFeatures"`
	// Indicates the product type of the license.
	ProductType pulumi.StringPtrOutput `pulumi:"productType"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the Esu servers.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// Specifies if this machine is licensed as part of a Software Assurance agreement.
	SoftwareAssuranceCustomer pulumi.BoolPtrOutput `pulumi:"softwareAssuranceCustomer"`
	// Indicates the subscription status of the product.
	SubscriptionStatus pulumi.StringPtrOutput `pulumi:"subscriptionStatus"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLicenseProfile registers a new resource with the given unique name, arguments, and options.
func NewLicenseProfile(ctx *pulumi.Context,
	name string, args *LicenseProfileArgs, opts ...pulumi.ResourceOption) (*LicenseProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MachineName == nil {
		return nil, errors.New("invalid value for required argument 'MachineName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute/v20230620preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20231003preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240331preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240520preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240710:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240731preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20240910preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20241110preview:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20250113:LicenseProfile"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20250219preview:LicenseProfile"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource LicenseProfile
	err := ctx.RegisterResource("azure-native:hybridcompute:LicenseProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLicenseProfile gets an existing LicenseProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicenseProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LicenseProfileState, opts ...pulumi.ResourceOption) (*LicenseProfile, error) {
	var resource LicenseProfile
	err := ctx.ReadResource("azure-native:hybridcompute:LicenseProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LicenseProfile resources.
type licenseProfileState struct {
}

type LicenseProfileState struct {
}

func (LicenseProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseProfileState)(nil)).Elem()
}

type licenseProfileArgs struct {
	// The resource id of the license.
	AssignedLicense *string `pulumi:"assignedLicense"`
	// The name of the license profile.
	LicenseProfileName *string `pulumi:"licenseProfileName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the hybrid machine.
	MachineName string `pulumi:"machineName"`
	// The list of product features.
	ProductFeatures []ProductFeature `pulumi:"productFeatures"`
	// Indicates the product type of the license.
	ProductType *string `pulumi:"productType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies if this machine is licensed as part of a Software Assurance agreement.
	SoftwareAssuranceCustomer *bool `pulumi:"softwareAssuranceCustomer"`
	// Indicates the subscription status of the product.
	SubscriptionStatus *string `pulumi:"subscriptionStatus"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LicenseProfile resource.
type LicenseProfileArgs struct {
	// The resource id of the license.
	AssignedLicense pulumi.StringPtrInput
	// The name of the license profile.
	LicenseProfileName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the hybrid machine.
	MachineName pulumi.StringInput
	// The list of product features.
	ProductFeatures ProductFeatureArrayInput
	// Indicates the product type of the license.
	ProductType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Specifies if this machine is licensed as part of a Software Assurance agreement.
	SoftwareAssuranceCustomer pulumi.BoolPtrInput
	// Indicates the subscription status of the product.
	SubscriptionStatus pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LicenseProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseProfileArgs)(nil)).Elem()
}

type LicenseProfileInput interface {
	pulumi.Input

	ToLicenseProfileOutput() LicenseProfileOutput
	ToLicenseProfileOutputWithContext(ctx context.Context) LicenseProfileOutput
}

func (*LicenseProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseProfile)(nil)).Elem()
}

func (i *LicenseProfile) ToLicenseProfileOutput() LicenseProfileOutput {
	return i.ToLicenseProfileOutputWithContext(context.Background())
}

func (i *LicenseProfile) ToLicenseProfileOutputWithContext(ctx context.Context) LicenseProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseProfileOutput)
}

type LicenseProfileOutput struct{ *pulumi.OutputState }

func (LicenseProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseProfile)(nil)).Elem()
}

func (o LicenseProfileOutput) ToLicenseProfileOutput() LicenseProfileOutput {
	return o
}

func (o LicenseProfileOutput) ToLicenseProfileOutputWithContext(ctx context.Context) LicenseProfileOutput {
	return o
}

// The resource id of the license.
func (o LicenseProfileOutput) AssignedLicense() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringPtrOutput { return v.AssignedLicense }).(pulumi.StringPtrOutput)
}

// The guid id of the license.
func (o LicenseProfileOutput) AssignedLicenseImmutableId() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.AssignedLicenseImmutableId }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LicenseProfileOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The timestamp in UTC when the billing ends.
func (o LicenseProfileOutput) BillingEndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.BillingEndDate }).(pulumi.StringOutput)
}

// The timestamp in UTC when the billing starts.
func (o LicenseProfileOutput) BillingStartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.BillingStartDate }).(pulumi.StringOutput)
}

// The timestamp in UTC when the user disenrolled the feature.
func (o LicenseProfileOutput) DisenrollmentDate() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.DisenrollmentDate }).(pulumi.StringOutput)
}

// The timestamp in UTC when the user enrolls the feature.
func (o LicenseProfileOutput) EnrollmentDate() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.EnrollmentDate }).(pulumi.StringOutput)
}

// The errors that were encountered during the feature enrollment or disenrollment.
func (o LicenseProfileOutput) Error() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *LicenseProfile) ErrorDetailResponseOutput { return v.Error }).(ErrorDetailResponseOutput)
}

// Indicates the eligibility state of Esu.
func (o LicenseProfileOutput) EsuEligibility() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.EsuEligibility }).(pulumi.StringOutput)
}

// Indicates whether there is an ESU Key currently active for the machine.
func (o LicenseProfileOutput) EsuKeyState() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.EsuKeyState }).(pulumi.StringOutput)
}

// The list of ESU keys.
func (o LicenseProfileOutput) EsuKeys() EsuKeyResponseArrayOutput {
	return o.ApplyT(func(v *LicenseProfile) EsuKeyResponseArrayOutput { return v.EsuKeys }).(EsuKeyResponseArrayOutput)
}

// The geo-location where the resource lives
func (o LicenseProfileOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LicenseProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The list of product features.
func (o LicenseProfileOutput) ProductFeatures() ProductFeatureResponseArrayOutput {
	return o.ApplyT(func(v *LicenseProfile) ProductFeatureResponseArrayOutput { return v.ProductFeatures }).(ProductFeatureResponseArrayOutput)
}

// Indicates the product type of the license.
func (o LicenseProfileOutput) ProductType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringPtrOutput { return v.ProductType }).(pulumi.StringPtrOutput)
}

// The provisioning state, which only appears in the response.
func (o LicenseProfileOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the Esu servers.
func (o LicenseProfileOutput) ServerType() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.ServerType }).(pulumi.StringOutput)
}

// Specifies if this machine is licensed as part of a Software Assurance agreement.
func (o LicenseProfileOutput) SoftwareAssuranceCustomer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.BoolPtrOutput { return v.SoftwareAssuranceCustomer }).(pulumi.BoolPtrOutput)
}

// Indicates the subscription status of the product.
func (o LicenseProfileOutput) SubscriptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringPtrOutput { return v.SubscriptionStatus }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LicenseProfileOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LicenseProfile) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LicenseProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LicenseProfileOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LicenseProfile) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LicenseProfileOutput{})
}
