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

// The object describing the smtp username resource.
//
// Uses Azure REST API version 2024-09-01-preview.
type SmtpUsername struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The application Id for the linked Entra Application.
	EntraApplicationId pulumi.StringOutput `pulumi:"entraApplicationId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant of the linked Entra Application.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The SMTP username. Could be free form or in the email address format.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewSmtpUsername registers a new resource with the given unique name, arguments, and options.
func NewSmtpUsername(ctx *pulumi.Context,
	name string, args *SmtpUsernameArgs, opts ...pulumi.ResourceOption) (*SmtpUsername, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CommunicationServiceName == nil {
		return nil, errors.New("invalid value for required argument 'CommunicationServiceName'")
	}
	if args.EntraApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'EntraApplicationId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:communication/v20240901preview:SmtpUsername"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SmtpUsername
	err := ctx.RegisterResource("azure-native:communication:SmtpUsername", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmtpUsername gets an existing SmtpUsername resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmtpUsername(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmtpUsernameState, opts ...pulumi.ResourceOption) (*SmtpUsername, error) {
	var resource SmtpUsername
	err := ctx.ReadResource("azure-native:communication:SmtpUsername", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmtpUsername resources.
type smtpUsernameState struct {
}

type SmtpUsernameState struct {
}

func (SmtpUsernameState) ElementType() reflect.Type {
	return reflect.TypeOf((*smtpUsernameState)(nil)).Elem()
}

type smtpUsernameArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName string `pulumi:"communicationServiceName"`
	// The application Id for the linked Entra Application.
	EntraApplicationId string `pulumi:"entraApplicationId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SmtpUsernameResource.
	SmtpUsername *string `pulumi:"smtpUsername"`
	// The tenant of the linked Entra Application.
	TenantId string `pulumi:"tenantId"`
	// The SMTP username. Could be free form or in the email address format.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a SmtpUsername resource.
type SmtpUsernameArgs struct {
	// The name of the CommunicationService resource.
	CommunicationServiceName pulumi.StringInput
	// The application Id for the linked Entra Application.
	EntraApplicationId pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the SmtpUsernameResource.
	SmtpUsername pulumi.StringPtrInput
	// The tenant of the linked Entra Application.
	TenantId pulumi.StringInput
	// The SMTP username. Could be free form or in the email address format.
	Username pulumi.StringInput
}

func (SmtpUsernameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smtpUsernameArgs)(nil)).Elem()
}

type SmtpUsernameInput interface {
	pulumi.Input

	ToSmtpUsernameOutput() SmtpUsernameOutput
	ToSmtpUsernameOutputWithContext(ctx context.Context) SmtpUsernameOutput
}

func (*SmtpUsername) ElementType() reflect.Type {
	return reflect.TypeOf((**SmtpUsername)(nil)).Elem()
}

func (i *SmtpUsername) ToSmtpUsernameOutput() SmtpUsernameOutput {
	return i.ToSmtpUsernameOutputWithContext(context.Background())
}

func (i *SmtpUsername) ToSmtpUsernameOutputWithContext(ctx context.Context) SmtpUsernameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmtpUsernameOutput)
}

type SmtpUsernameOutput struct{ *pulumi.OutputState }

func (SmtpUsernameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmtpUsername)(nil)).Elem()
}

func (o SmtpUsernameOutput) ToSmtpUsernameOutput() SmtpUsernameOutput {
	return o
}

func (o SmtpUsernameOutput) ToSmtpUsernameOutputWithContext(ctx context.Context) SmtpUsernameOutput {
	return o
}

// The Azure API version of the resource.
func (o SmtpUsernameOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpUsername) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The application Id for the linked Entra Application.
func (o SmtpUsernameOutput) EntraApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpUsername) pulumi.StringOutput { return v.EntraApplicationId }).(pulumi.StringOutput)
}

// The name of the resource
func (o SmtpUsernameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpUsername) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SmtpUsernameOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SmtpUsername) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant of the linked Entra Application.
func (o SmtpUsernameOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpUsername) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SmtpUsernameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpUsername) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The SMTP username. Could be free form or in the email address format.
func (o SmtpUsernameOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *SmtpUsername) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SmtpUsernameOutput{})
}
