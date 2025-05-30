// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costmanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Scheduled action definition.
//
// Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
//
// Other available API versions: 2022-04-01-preview, 2022-06-01-preview, 2022-10-01, 2023-03-01, 2023-04-01-preview, 2023-07-01-preview, 2023-08-01, 2023-09-01, 2023-11-01, 2024-10-01-preview, 2025-03-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native costmanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ScheduledAction struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Scheduled action name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
	ETag pulumi.StringOutput `pulumi:"eTag"`
	// Destination format of the view data. This is optional.
	FileDestination FileDestinationResponsePtrOutput `pulumi:"fileDestination"`
	// Kind of the scheduled action.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Notification properties based on scheduled action kind.
	Notification NotificationPropertiesResponseOutput `pulumi:"notification"`
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail pulumi.StringPtrOutput `pulumi:"notificationEmail"`
	// Schedule of the scheduled action.
	Schedule SchedulePropertiesResponseOutput `pulumi:"schedule"`
	// For private scheduled action(Create or Update), scope will be empty.<br /> For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Status of the scheduled action.
	Status pulumi.StringOutput `pulumi:"status"`
	// Kind of the scheduled action.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId pulumi.StringOutput `pulumi:"viewId"`
}

// NewScheduledAction registers a new resource with the given unique name, arguments, and options.
func NewScheduledAction(ctx *pulumi.Context,
	name string, args *ScheduledActionArgs, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Notification == nil {
		return nil, errors.New("invalid value for required argument 'Notification'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.ViewId == nil {
		return nil, errors.New("invalid value for required argument 'ViewId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20220401preview:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20220601preview:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20221001:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230301:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230401preview:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230701preview:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230801:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20230901:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20231101:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20240801:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20241001preview:ScheduledAction"),
		},
		{
			Type: pulumi.String("azure-native:costmanagement/v20250301:ScheduledAction"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ScheduledAction
	err := ctx.RegisterResource("azure-native:costmanagement:ScheduledAction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScheduledAction gets an existing ScheduledAction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScheduledAction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduledActionState, opts ...pulumi.ResourceOption) (*ScheduledAction, error) {
	var resource ScheduledAction
	err := ctx.ReadResource("azure-native:costmanagement:ScheduledAction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScheduledAction resources.
type scheduledActionState struct {
}

type ScheduledActionState struct {
}

func (ScheduledActionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionState)(nil)).Elem()
}

type scheduledActionArgs struct {
	// Scheduled action name.
	DisplayName string `pulumi:"displayName"`
	// Destination format of the view data. This is optional.
	FileDestination *FileDestination `pulumi:"fileDestination"`
	// Kind of the scheduled action.
	Kind *string `pulumi:"kind"`
	// Scheduled action name.
	Name *string `pulumi:"name"`
	// Notification properties based on scheduled action kind.
	Notification NotificationProperties `pulumi:"notification"`
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail *string `pulumi:"notificationEmail"`
	// Schedule of the scheduled action.
	Schedule ScheduleProperties `pulumi:"schedule"`
	// For private scheduled action(Create or Update), scope will be empty.<br /> For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope *string `pulumi:"scope"`
	// Status of the scheduled action.
	Status string `pulumi:"status"`
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId string `pulumi:"viewId"`
}

// The set of arguments for constructing a ScheduledAction resource.
type ScheduledActionArgs struct {
	// Scheduled action name.
	DisplayName pulumi.StringInput
	// Destination format of the view data. This is optional.
	FileDestination FileDestinationPtrInput
	// Kind of the scheduled action.
	Kind pulumi.StringPtrInput
	// Scheduled action name.
	Name pulumi.StringPtrInput
	// Notification properties based on scheduled action kind.
	Notification NotificationPropertiesInput
	// Email address of the point of contact that should get the unsubscribe requests and notification emails.
	NotificationEmail pulumi.StringPtrInput
	// Schedule of the scheduled action.
	Schedule SchedulePropertiesInput
	// For private scheduled action(Create or Update), scope will be empty.<br /> For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
	Scope pulumi.StringPtrInput
	// Status of the scheduled action.
	Status pulumi.StringInput
	// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
	ViewId pulumi.StringInput
}

func (ScheduledActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduledActionArgs)(nil)).Elem()
}

type ScheduledActionInput interface {
	pulumi.Input

	ToScheduledActionOutput() ScheduledActionOutput
	ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput
}

func (*ScheduledAction) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil)).Elem()
}

func (i *ScheduledAction) ToScheduledActionOutput() ScheduledActionOutput {
	return i.ToScheduledActionOutputWithContext(context.Background())
}

func (i *ScheduledAction) ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledActionOutput)
}

type ScheduledActionOutput struct{ *pulumi.OutputState }

func (ScheduledActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduledAction)(nil)).Elem()
}

func (o ScheduledActionOutput) ToScheduledActionOutput() ScheduledActionOutput {
	return o
}

func (o ScheduledActionOutput) ToScheduledActionOutputWithContext(ctx context.Context) ScheduledActionOutput {
	return o
}

// The Azure API version of the resource.
func (o ScheduledActionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Scheduled action name.
func (o ScheduledActionOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Etag. For update calls, eTag is optional and can be specified to achieve optimistic concurrency. Fetch the resource's eTag by doing a 'GET' call first and then including the latest eTag as part of the request body or 'If-Match' header while performing the update. For create calls, eTag is not required.
func (o ScheduledActionOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.ETag }).(pulumi.StringOutput)
}

// Destination format of the view data. This is optional.
func (o ScheduledActionOutput) FileDestination() FileDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ScheduledAction) FileDestinationResponsePtrOutput { return v.FileDestination }).(FileDestinationResponsePtrOutput)
}

// Kind of the scheduled action.
func (o ScheduledActionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ScheduledActionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Notification properties based on scheduled action kind.
func (o ScheduledActionOutput) Notification() NotificationPropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledAction) NotificationPropertiesResponseOutput { return v.Notification }).(NotificationPropertiesResponseOutput)
}

// Email address of the point of contact that should get the unsubscribe requests and notification emails.
func (o ScheduledActionOutput) NotificationEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringPtrOutput { return v.NotificationEmail }).(pulumi.StringPtrOutput)
}

// Schedule of the scheduled action.
func (o ScheduledActionOutput) Schedule() SchedulePropertiesResponseOutput {
	return o.ApplyT(func(v *ScheduledAction) SchedulePropertiesResponseOutput { return v.Schedule }).(SchedulePropertiesResponseOutput)
}

// For private scheduled action(Create or Update), scope will be empty.<br /> For shared scheduled action(Create or Update By Scope), Cost Management scope can be 'subscriptions/{subscriptionId}' for subscription scope, 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for BillingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoiceSections/{invoiceSectionId}' for InvoiceSection scope, '/providers/Microsoft.CostManagement/externalBillingAccounts/{externalBillingAccountName}' for ExternalBillingAccount scope, and '/providers/Microsoft.CostManagement/externalSubscriptions/{externalSubscriptionName}' for ExternalSubscription scope.
func (o ScheduledActionOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Status of the scheduled action.
func (o ScheduledActionOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Kind of the scheduled action.
func (o ScheduledActionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ScheduledAction) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ScheduledActionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Cost analysis viewId used for scheduled action. For example, '/providers/Microsoft.CostManagement/views/swaggerExample'
func (o ScheduledActionOutput) ViewId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScheduledAction) pulumi.StringOutput { return v.ViewId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduledActionOutput{})
}
