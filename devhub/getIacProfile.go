// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representation of a IacProfile.
//
// Uses Azure REST API version 2024-05-01-preview.
//
// Other available API versions: 2024-08-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devhub [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupIacProfile(ctx *pulumi.Context, args *LookupIacProfileArgs, opts ...pulumi.InvokeOption) (*LookupIacProfileResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupIacProfileResult
	err := ctx.Invoke("azure-native:devhub:getIacProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIacProfileArgs struct {
	// The name of the IacProfile.
	IacProfileName string `pulumi:"iacProfileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Resource representation of a IacProfile.
type LookupIacProfileResult struct {
	// Determines the authorization status of requests.
	AuthStatus string `pulumi:"authStatus"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Repository Branch Name
	BranchName *string `pulumi:"branchName"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The status of the Pull Request submitted against the users repository.
	PrStatus string `pulumi:"prStatus"`
	// The number associated with the submitted pull request.
	PullNumber int `pulumi:"pullNumber"`
	// Repository Main Branch
	RepositoryMainBranch *string `pulumi:"repositoryMainBranch"`
	// Repository Name
	RepositoryName *string `pulumi:"repositoryName"`
	// Repository Owner
	RepositoryOwner *string                   `pulumi:"repositoryOwner"`
	Stages          []StagePropertiesResponse `pulumi:"stages"`
	// Terraform Storage Account Name
	StorageAccountName *string `pulumi:"storageAccountName"`
	// Terraform Storage Account Resource Group
	StorageAccountResourceGroup *string `pulumi:"storageAccountResourceGroup"`
	// Terraform Storage Account Subscription
	StorageAccountSubscription *string `pulumi:"storageAccountSubscription"`
	// Terraform Container Name
	StorageContainerName *string `pulumi:"storageContainerName"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags      map[string]string               `pulumi:"tags"`
	Templates []IacTemplatePropertiesResponse `pulumi:"templates"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupIacProfileOutput(ctx *pulumi.Context, args LookupIacProfileOutputArgs, opts ...pulumi.InvokeOption) LookupIacProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIacProfileResultOutput, error) {
			args := v.(LookupIacProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:devhub:getIacProfile", args, LookupIacProfileResultOutput{}, options).(LookupIacProfileResultOutput), nil
		}).(LookupIacProfileResultOutput)
}

type LookupIacProfileOutputArgs struct {
	// The name of the IacProfile.
	IacProfileName pulumi.StringInput `pulumi:"iacProfileName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIacProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIacProfileArgs)(nil)).Elem()
}

// Resource representation of a IacProfile.
type LookupIacProfileResultOutput struct{ *pulumi.OutputState }

func (LookupIacProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIacProfileResult)(nil)).Elem()
}

func (o LookupIacProfileResultOutput) ToLookupIacProfileResultOutput() LookupIacProfileResultOutput {
	return o
}

func (o LookupIacProfileResultOutput) ToLookupIacProfileResultOutputWithContext(ctx context.Context) LookupIacProfileResultOutput {
	return o
}

// Determines the authorization status of requests.
func (o LookupIacProfileResultOutput) AuthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.AuthStatus }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o LookupIacProfileResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Repository Branch Name
func (o LookupIacProfileResultOutput) BranchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.BranchName }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o LookupIacProfileResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
func (o LookupIacProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupIacProfileResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupIacProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The status of the Pull Request submitted against the users repository.
func (o LookupIacProfileResultOutput) PrStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.PrStatus }).(pulumi.StringOutput)
}

// The number associated with the submitted pull request.
func (o LookupIacProfileResultOutput) PullNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIacProfileResult) int { return v.PullNumber }).(pulumi.IntOutput)
}

// Repository Main Branch
func (o LookupIacProfileResultOutput) RepositoryMainBranch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.RepositoryMainBranch }).(pulumi.StringPtrOutput)
}

// Repository Name
func (o LookupIacProfileResultOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

// Repository Owner
func (o LookupIacProfileResultOutput) RepositoryOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.RepositoryOwner }).(pulumi.StringPtrOutput)
}

func (o LookupIacProfileResultOutput) Stages() StagePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupIacProfileResult) []StagePropertiesResponse { return v.Stages }).(StagePropertiesResponseArrayOutput)
}

// Terraform Storage Account Name
func (o LookupIacProfileResultOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

// Terraform Storage Account Resource Group
func (o LookupIacProfileResultOutput) StorageAccountResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.StorageAccountResourceGroup }).(pulumi.StringPtrOutput)
}

// Terraform Storage Account Subscription
func (o LookupIacProfileResultOutput) StorageAccountSubscription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.StorageAccountSubscription }).(pulumi.StringPtrOutput)
}

// Terraform Container Name
func (o LookupIacProfileResultOutput) StorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIacProfileResult) *string { return v.StorageContainerName }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupIacProfileResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupIacProfileResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupIacProfileResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIacProfileResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIacProfileResultOutput) Templates() IacTemplatePropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupIacProfileResult) []IacTemplatePropertiesResponse { return v.Templates }).(IacTemplatePropertiesResponseArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupIacProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIacProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIacProfileResultOutput{})
}
