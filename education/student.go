// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package education

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Student details.
//
// Uses Azure REST API version 2021-12-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-12-01-preview.
type Student struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Student Budget
	Budget AmountResponseOutput `pulumi:"budget"`
	// Date student was added to the lab
	EffectiveDate pulumi.StringOutput `pulumi:"effectiveDate"`
	// Student Email
	Email pulumi.StringOutput `pulumi:"email"`
	// Date this student is set to expire from the lab.
	ExpirationDate pulumi.StringOutput `pulumi:"expirationDate"`
	// First Name
	FirstName pulumi.StringOutput `pulumi:"firstName"`
	// Last Name
	LastName pulumi.StringOutput `pulumi:"lastName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Student Role
	Role pulumi.StringOutput `pulumi:"role"`
	// Student Lab Status
	Status pulumi.StringOutput `pulumi:"status"`
	// Subscription alias
	SubscriptionAlias pulumi.StringPtrOutput `pulumi:"subscriptionAlias"`
	// Subscription Id
	SubscriptionId pulumi.StringOutput `pulumi:"subscriptionId"`
	// subscription invite last sent date
	SubscriptionInviteLastSentDate pulumi.StringPtrOutput `pulumi:"subscriptionInviteLastSentDate"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStudent registers a new resource with the given unique name, arguments, and options.
func NewStudent(ctx *pulumi.Context,
	name string, args *StudentArgs, opts ...pulumi.ResourceOption) (*Student, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountName == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountName'")
	}
	if args.BillingProfileName == nil {
		return nil, errors.New("invalid value for required argument 'BillingProfileName'")
	}
	if args.Budget == nil {
		return nil, errors.New("invalid value for required argument 'Budget'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.ExpirationDate == nil {
		return nil, errors.New("invalid value for required argument 'ExpirationDate'")
	}
	if args.FirstName == nil {
		return nil, errors.New("invalid value for required argument 'FirstName'")
	}
	if args.InvoiceSectionName == nil {
		return nil, errors.New("invalid value for required argument 'InvoiceSectionName'")
	}
	if args.LastName == nil {
		return nil, errors.New("invalid value for required argument 'LastName'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:education/v20211201preview:Student"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Student
	err := ctx.RegisterResource("azure-native:education:Student", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStudent gets an existing Student resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStudent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudentState, opts ...pulumi.ResourceOption) (*Student, error) {
	var resource Student
	err := ctx.ReadResource("azure-native:education:Student", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Student resources.
type studentState struct {
}

type StudentState struct {
}

func (StudentState) ElementType() reflect.Type {
	return reflect.TypeOf((*studentState)(nil)).Elem()
}

type studentArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName string `pulumi:"billingAccountName"`
	// The ID that uniquely identifies a billing profile.
	BillingProfileName string `pulumi:"billingProfileName"`
	// Student Budget
	Budget Amount `pulumi:"budget"`
	// Student Email
	Email string `pulumi:"email"`
	// Date this student is set to expire from the lab.
	ExpirationDate string `pulumi:"expirationDate"`
	// First Name
	FirstName string `pulumi:"firstName"`
	// The ID that uniquely identifies an invoice section.
	InvoiceSectionName string `pulumi:"invoiceSectionName"`
	// Last Name
	LastName string `pulumi:"lastName"`
	// Student Role
	Role string `pulumi:"role"`
	// Student alias.
	StudentAlias *string `pulumi:"studentAlias"`
	// Subscription alias
	SubscriptionAlias *string `pulumi:"subscriptionAlias"`
	// subscription invite last sent date
	SubscriptionInviteLastSentDate *string `pulumi:"subscriptionInviteLastSentDate"`
}

// The set of arguments for constructing a Student resource.
type StudentArgs struct {
	// The ID that uniquely identifies a billing account.
	BillingAccountName pulumi.StringInput
	// The ID that uniquely identifies a billing profile.
	BillingProfileName pulumi.StringInput
	// Student Budget
	Budget AmountInput
	// Student Email
	Email pulumi.StringInput
	// Date this student is set to expire from the lab.
	ExpirationDate pulumi.StringInput
	// First Name
	FirstName pulumi.StringInput
	// The ID that uniquely identifies an invoice section.
	InvoiceSectionName pulumi.StringInput
	// Last Name
	LastName pulumi.StringInput
	// Student Role
	Role pulumi.StringInput
	// Student alias.
	StudentAlias pulumi.StringPtrInput
	// Subscription alias
	SubscriptionAlias pulumi.StringPtrInput
	// subscription invite last sent date
	SubscriptionInviteLastSentDate pulumi.StringPtrInput
}

func (StudentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studentArgs)(nil)).Elem()
}

type StudentInput interface {
	pulumi.Input

	ToStudentOutput() StudentOutput
	ToStudentOutputWithContext(ctx context.Context) StudentOutput
}

func (*Student) ElementType() reflect.Type {
	return reflect.TypeOf((**Student)(nil)).Elem()
}

func (i *Student) ToStudentOutput() StudentOutput {
	return i.ToStudentOutputWithContext(context.Background())
}

func (i *Student) ToStudentOutputWithContext(ctx context.Context) StudentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudentOutput)
}

type StudentOutput struct{ *pulumi.OutputState }

func (StudentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Student)(nil)).Elem()
}

func (o StudentOutput) ToStudentOutput() StudentOutput {
	return o
}

func (o StudentOutput) ToStudentOutputWithContext(ctx context.Context) StudentOutput {
	return o
}

// The Azure API version of the resource.
func (o StudentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Student Budget
func (o StudentOutput) Budget() AmountResponseOutput {
	return o.ApplyT(func(v *Student) AmountResponseOutput { return v.Budget }).(AmountResponseOutput)
}

// Date student was added to the lab
func (o StudentOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.EffectiveDate }).(pulumi.StringOutput)
}

// Student Email
func (o StudentOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Date this student is set to expire from the lab.
func (o StudentOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.ExpirationDate }).(pulumi.StringOutput)
}

// First Name
func (o StudentOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.FirstName }).(pulumi.StringOutput)
}

// Last Name
func (o StudentOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.LastName }).(pulumi.StringOutput)
}

// The name of the resource
func (o StudentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Student Role
func (o StudentOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Student Lab Status
func (o StudentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Subscription alias
func (o StudentOutput) SubscriptionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Student) pulumi.StringPtrOutput { return v.SubscriptionAlias }).(pulumi.StringPtrOutput)
}

// Subscription Id
func (o StudentOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

// subscription invite last sent date
func (o StudentOutput) SubscriptionInviteLastSentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Student) pulumi.StringPtrOutput { return v.SubscriptionInviteLastSentDate }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o StudentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Student) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o StudentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Student) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StudentOutput{})
}
