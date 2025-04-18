// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves details of a specific security connector
//
// Uses Azure REST API version 2024-08-01-preview.
//
// Other available API versions: 2021-07-01-preview, 2021-12-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2023-03-01-preview, 2023-10-01-preview, 2024-03-01-preview, 2024-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native security [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupSecurityConnector(ctx *pulumi.Context, args *LookupSecurityConnectorArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityConnectorResult
	err := ctx.Invoke("azure-native:security:getSecurityConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}

// The security connector resource.
type LookupSecurityConnectorResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// The security connector environment data.
	EnvironmentData interface{} `pulumi:"environmentData"`
	// The multi cloud resource's cloud name.
	EnvironmentName *string `pulumi:"environmentName"`
	// Entity tag is used for comparing two or more entities from the same requested resource.
	Etag *string `pulumi:"etag"`
	// The multi cloud resource identifier (account id in case of AWS connector, project number in case of GCP connector).
	HierarchyIdentifier *string `pulumi:"hierarchyIdentifier"`
	// The date on which the trial period will end, if applicable. Trial period exists for 30 days after upgrading to payed offerings.
	HierarchyIdentifierTrialEndDate string `pulumi:"hierarchyIdentifierTrialEndDate"`
	// Resource Id
	Id string `pulumi:"id"`
	// Kind of the resource
	Kind *string `pulumi:"kind"`
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// A collection of offerings for the security connector.
	Offerings []interface{} `pulumi:"offerings"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// A list of key value pairs that describe the resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupSecurityConnectorOutput(ctx *pulumi.Context, args LookupSecurityConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityConnectorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSecurityConnectorResultOutput, error) {
			args := v.(LookupSecurityConnectorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:security:getSecurityConnector", args, LookupSecurityConnectorResultOutput{}, options).(LookupSecurityConnectorResultOutput), nil
		}).(LookupSecurityConnectorResultOutput)
}

type LookupSecurityConnectorOutputArgs struct {
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The security connector name.
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (LookupSecurityConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorArgs)(nil)).Elem()
}

// The security connector resource.
type LookupSecurityConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorResult)(nil)).Elem()
}

func (o LookupSecurityConnectorResultOutput) ToLookupSecurityConnectorResultOutput() LookupSecurityConnectorResultOutput {
	return o
}

func (o LookupSecurityConnectorResultOutput) ToLookupSecurityConnectorResultOutputWithContext(ctx context.Context) LookupSecurityConnectorResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupSecurityConnectorResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The security connector environment data.
func (o LookupSecurityConnectorResultOutput) EnvironmentData() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) interface{} { return v.EnvironmentData }).(pulumi.AnyOutput)
}

// The multi cloud resource's cloud name.
func (o LookupSecurityConnectorResultOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

// Entity tag is used for comparing two or more entities from the same requested resource.
func (o LookupSecurityConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// The multi cloud resource identifier (account id in case of AWS connector, project number in case of GCP connector).
func (o LookupSecurityConnectorResultOutput) HierarchyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.HierarchyIdentifier }).(pulumi.StringPtrOutput)
}

// The date on which the trial period will end, if applicable. Trial period exists for 30 days after upgrading to payed offerings.
func (o LookupSecurityConnectorResultOutput) HierarchyIdentifierTrialEndDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.HierarchyIdentifierTrialEndDate }).(pulumi.StringOutput)
}

// Resource Id
func (o LookupSecurityConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kind of the resource
func (o LookupSecurityConnectorResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Location where the resource is stored
func (o LookupSecurityConnectorResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupSecurityConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of offerings for the security connector.
func (o LookupSecurityConnectorResultOutput) Offerings() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) []interface{} { return v.Offerings }).(pulumi.ArrayOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupSecurityConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// A list of key value pairs that describe the resource.
func (o LookupSecurityConnectorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupSecurityConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityConnectorResultOutput{})
}
