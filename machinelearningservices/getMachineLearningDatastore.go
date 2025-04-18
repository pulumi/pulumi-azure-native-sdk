// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a Datastore by name.
//
// Uses Azure REST API version 2020-05-01-preview.
func LookupMachineLearningDatastore(ctx *pulumi.Context, args *LookupMachineLearningDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningDatastoreResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMachineLearningDatastoreResult
	err := ctx.Invoke("azure-native:machinelearningservices:getMachineLearningDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupMachineLearningDatastoreArgs struct {
	// The Datastore name.
	DatastoreName string `pulumi:"datastoreName"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Machine Learning datastore object wrapped into ARM resource envelope.
type LookupMachineLearningDatastoreResult struct {
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// Datastore properties
	Properties DatastoreResponse `pulumi:"properties"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupMachineLearningDatastoreResult
func (val *LookupMachineLearningDatastoreResult) Defaults() *LookupMachineLearningDatastoreResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
func LookupMachineLearningDatastoreOutput(ctx *pulumi.Context, args LookupMachineLearningDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupMachineLearningDatastoreResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMachineLearningDatastoreResultOutput, error) {
			args := v.(LookupMachineLearningDatastoreArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:machinelearningservices:getMachineLearningDatastore", args, LookupMachineLearningDatastoreResultOutput{}, options).(LookupMachineLearningDatastoreResultOutput), nil
		}).(LookupMachineLearningDatastoreResultOutput)
}

type LookupMachineLearningDatastoreOutputArgs struct {
	// The Datastore name.
	DatastoreName pulumi.StringInput `pulumi:"datastoreName"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMachineLearningDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningDatastoreArgs)(nil)).Elem()
}

// Machine Learning datastore object wrapped into ARM resource envelope.
type LookupMachineLearningDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupMachineLearningDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineLearningDatastoreResult)(nil)).Elem()
}

func (o LookupMachineLearningDatastoreResultOutput) ToLookupMachineLearningDatastoreResultOutput() LookupMachineLearningDatastoreResultOutput {
	return o
}

func (o LookupMachineLearningDatastoreResultOutput) ToLookupMachineLearningDatastoreResultOutputWithContext(ctx context.Context) LookupMachineLearningDatastoreResultOutput {
	return o
}

// The Azure API version of the resource.
func (o LookupMachineLearningDatastoreResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Specifies the resource ID.
func (o LookupMachineLearningDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupMachineLearningDatastoreResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Specifies the location of the resource.
func (o LookupMachineLearningDatastoreResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupMachineLearningDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// Datastore properties
func (o LookupMachineLearningDatastoreResultOutput) Properties() DatastoreResponseOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) DatastoreResponse { return v.Properties }).(DatastoreResponseOutput)
}

// The sku of the workspace.
func (o LookupMachineLearningDatastoreResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupMachineLearningDatastoreResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupMachineLearningDatastoreResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineLearningDatastoreResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineLearningDatastoreResultOutput{})
}
