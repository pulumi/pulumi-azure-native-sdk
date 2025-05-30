// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package communication

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A class representing a Domains resource.
//
// Uses Azure REST API version 2023-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-03-31.
//
// Other available API versions: 2023-03-31, 2023-04-01, 2023-04-01-preview, 2024-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
//
// Note: If `domainManagement` is set to `AzureManaged`, then `domainName` is required.
type Domain struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The location where the Domains resource data is stored at rest.
	DataLocation pulumi.StringOutput `pulumi:"dataLocation"`
	// Describes how a Domains resource is being managed.
	DomainManagement pulumi.StringOutput `pulumi:"domainManagement"`
	// P2 sender domain that is displayed to the email recipients [RFC 5322].
	FromSenderDomain pulumi.StringOutput `pulumi:"fromSenderDomain"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// P1 sender domain that is present on the email envelope [RFC 5321].
	MailFromSenderDomain pulumi.StringOutput `pulumi:"mailFromSenderDomain"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Describes whether user engagement tracking is enabled or disabled.
	UserEngagementTracking pulumi.StringPtrOutput `pulumi:"userEngagementTracking"`
	// List of DnsRecord
	VerificationRecords DomainPropertiesResponseVerificationRecordsOutput `pulumi:"verificationRecords"`
	// List of VerificationStatusRecord
	VerificationStates DomainPropertiesResponseVerificationStatesOutput `pulumi:"verificationStates"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainManagement == nil {
		return nil, errors.New("invalid value for required argument 'DomainManagement'")
	}
	if args.EmailServiceName == nil {
		return nil, errors.New("invalid value for required argument 'EmailServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:communication/v20211001preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20220701preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230301preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230331:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230401:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230401preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20230601preview:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20240901preview:Domain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("azure-native:communication:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:communication:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Describes how a Domains resource is being managed.
	DomainManagement string `pulumi:"domainManagement"`
	// The name of the Domains resource.
	DomainName *string `pulumi:"domainName"`
	// The name of the EmailService resource.
	EmailServiceName string `pulumi:"emailServiceName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Describes whether user engagement tracking is enabled or disabled.
	UserEngagementTracking *string `pulumi:"userEngagementTracking"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Describes how a Domains resource is being managed.
	DomainManagement pulumi.StringInput
	// The name of the Domains resource.
	DomainName pulumi.StringPtrInput
	// The name of the EmailService resource.
	EmailServiceName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Describes whether user engagement tracking is enabled or disabled.
	UserEngagementTracking pulumi.StringPtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// The Azure API version of the resource.
func (o DomainOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The location where the Domains resource data is stored at rest.
func (o DomainOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DataLocation }).(pulumi.StringOutput)
}

// Describes how a Domains resource is being managed.
func (o DomainOutput) DomainManagement() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainManagement }).(pulumi.StringOutput)
}

// P2 sender domain that is displayed to the email recipients [RFC 5322].
func (o DomainOutput) FromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.FromSenderDomain }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o DomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// P1 sender domain that is present on the email envelope [RFC 5321].
func (o DomainOutput) MailFromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.MailFromSenderDomain }).(pulumi.StringOutput)
}

// The name of the resource
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o DomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Domain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Describes whether user engagement tracking is enabled or disabled.
func (o DomainOutput) UserEngagementTracking() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.UserEngagementTracking }).(pulumi.StringPtrOutput)
}

// List of DnsRecord
func (o DomainOutput) VerificationRecords() DomainPropertiesResponseVerificationRecordsOutput {
	return o.ApplyT(func(v *Domain) DomainPropertiesResponseVerificationRecordsOutput { return v.VerificationRecords }).(DomainPropertiesResponseVerificationRecordsOutput)
}

// List of VerificationStatusRecord
func (o DomainOutput) VerificationStates() DomainPropertiesResponseVerificationStatesOutput {
	return o.ApplyT(func(v *Domain) DomainPropertiesResponseVerificationStatesOutput { return v.VerificationStates }).(DomainPropertiesResponseVerificationStatesOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
