// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes what action to be applied when rule matches
type ActionType string

const (
	ActionTypeAllow    = ActionType("Allow")
	ActionTypeBlock    = ActionType("Block")
	ActionTypeLog      = ActionType("Log")
	ActionTypeRedirect = ActionType("Redirect")
)

// Caching behavior for the requests
type CacheBehavior string

const (
	CacheBehaviorBypassCache  = CacheBehavior("BypassCache")
	CacheBehaviorOverride     = CacheBehavior("Override")
	CacheBehaviorSetIfMissing = CacheBehavior("SetIfMissing")
)

// The level at which the content needs to be cached.
type CacheType string

const (
	CacheTypeAll = CacheType("All")
)

// Describes operator to be matched
type CookiesOperator string

const (
	CookiesOperatorAny                = CookiesOperator("Any")
	CookiesOperatorEqual              = CookiesOperator("Equal")
	CookiesOperatorContains           = CookiesOperator("Contains")
	CookiesOperatorBeginsWith         = CookiesOperator("BeginsWith")
	CookiesOperatorEndsWith           = CookiesOperator("EndsWith")
	CookiesOperatorLessThan           = CookiesOperator("LessThan")
	CookiesOperatorLessThanOrEqual    = CookiesOperator("LessThanOrEqual")
	CookiesOperatorGreaterThan        = CookiesOperator("GreaterThan")
	CookiesOperatorGreaterThanOrEqual = CookiesOperator("GreaterThanOrEqual")
)

// Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
type CustomRuleEnabledState string

const (
	CustomRuleEnabledStateDisabled = CustomRuleEnabledState("Disabled")
	CustomRuleEnabledStateEnabled  = CustomRuleEnabledState("Enabled")
)

// The name of the action for the delivery rule.
type DeliveryRuleAction string

const (
	DeliveryRuleActionCacheExpiration      = DeliveryRuleAction("CacheExpiration")
	DeliveryRuleActionCacheKeyQueryString  = DeliveryRuleAction("CacheKeyQueryString")
	DeliveryRuleActionModifyRequestHeader  = DeliveryRuleAction("ModifyRequestHeader")
	DeliveryRuleActionModifyResponseHeader = DeliveryRuleAction("ModifyResponseHeader")
	DeliveryRuleActionUrlRedirect          = DeliveryRuleAction("UrlRedirect")
	DeliveryRuleActionUrlRewrite           = DeliveryRuleAction("UrlRewrite")
)

// Protocol to use for the redirect. The default value is MatchRequest
type DestinationProtocol string

const (
	DestinationProtocolMatchRequest = DestinationProtocol("MatchRequest")
	DestinationProtocolHttp         = DestinationProtocol("Http")
	DestinationProtocolHttps        = DestinationProtocol("Https")
)

// Action of the geo filter, i.e. allow or block access.
type GeoFilterActions string

const (
	GeoFilterActionsBlock = GeoFilterActions("Block")
	GeoFilterActionsAllow = GeoFilterActions("Allow")
)

func (GeoFilterActions) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoFilterActions)(nil)).Elem()
}

func (e GeoFilterActions) ToGeoFilterActionsOutput() GeoFilterActionsOutput {
	return pulumi.ToOutput(e).(GeoFilterActionsOutput)
}

func (e GeoFilterActions) ToGeoFilterActionsOutputWithContext(ctx context.Context) GeoFilterActionsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GeoFilterActionsOutput)
}

func (e GeoFilterActions) ToGeoFilterActionsPtrOutput() GeoFilterActionsPtrOutput {
	return e.ToGeoFilterActionsPtrOutputWithContext(context.Background())
}

func (e GeoFilterActions) ToGeoFilterActionsPtrOutputWithContext(ctx context.Context) GeoFilterActionsPtrOutput {
	return GeoFilterActions(e).ToGeoFilterActionsOutputWithContext(ctx).ToGeoFilterActionsPtrOutputWithContext(ctx)
}

func (e GeoFilterActions) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoFilterActions) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GeoFilterActions) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GeoFilterActions) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GeoFilterActionsOutput struct{ *pulumi.OutputState }

func (GeoFilterActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GeoFilterActions)(nil)).Elem()
}

func (o GeoFilterActionsOutput) ToGeoFilterActionsOutput() GeoFilterActionsOutput {
	return o
}

func (o GeoFilterActionsOutput) ToGeoFilterActionsOutputWithContext(ctx context.Context) GeoFilterActionsOutput {
	return o
}

func (o GeoFilterActionsOutput) ToGeoFilterActionsPtrOutput() GeoFilterActionsPtrOutput {
	return o.ToGeoFilterActionsPtrOutputWithContext(context.Background())
}

func (o GeoFilterActionsOutput) ToGeoFilterActionsPtrOutputWithContext(ctx context.Context) GeoFilterActionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GeoFilterActions) *GeoFilterActions {
		return &v
	}).(GeoFilterActionsPtrOutput)
}

func (o GeoFilterActionsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GeoFilterActionsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoFilterActions) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GeoFilterActionsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoFilterActionsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GeoFilterActions) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GeoFilterActionsPtrOutput struct{ *pulumi.OutputState }

func (GeoFilterActionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeoFilterActions)(nil)).Elem()
}

func (o GeoFilterActionsPtrOutput) ToGeoFilterActionsPtrOutput() GeoFilterActionsPtrOutput {
	return o
}

func (o GeoFilterActionsPtrOutput) ToGeoFilterActionsPtrOutputWithContext(ctx context.Context) GeoFilterActionsPtrOutput {
	return o
}

func (o GeoFilterActionsPtrOutput) Elem() GeoFilterActionsOutput {
	return o.ApplyT(func(v *GeoFilterActions) GeoFilterActions {
		if v != nil {
			return *v
		}
		var ret GeoFilterActions
		return ret
	}).(GeoFilterActionsOutput)
}

func (o GeoFilterActionsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GeoFilterActionsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GeoFilterActions) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GeoFilterActionsInput is an input type that accepts GeoFilterActionsArgs and GeoFilterActionsOutput values.
// You can construct a concrete instance of `GeoFilterActionsInput` via:
//
//	GeoFilterActionsArgs{...}
type GeoFilterActionsInput interface {
	pulumi.Input

	ToGeoFilterActionsOutput() GeoFilterActionsOutput
	ToGeoFilterActionsOutputWithContext(context.Context) GeoFilterActionsOutput
}

var geoFilterActionsPtrType = reflect.TypeOf((**GeoFilterActions)(nil)).Elem()

type GeoFilterActionsPtrInput interface {
	pulumi.Input

	ToGeoFilterActionsPtrOutput() GeoFilterActionsPtrOutput
	ToGeoFilterActionsPtrOutputWithContext(context.Context) GeoFilterActionsPtrOutput
}

type geoFilterActionsPtr string

func GeoFilterActionsPtr(v string) GeoFilterActionsPtrInput {
	return (*geoFilterActionsPtr)(&v)
}

func (*geoFilterActionsPtr) ElementType() reflect.Type {
	return geoFilterActionsPtrType
}

func (in *geoFilterActionsPtr) ToGeoFilterActionsPtrOutput() GeoFilterActionsPtrOutput {
	return pulumi.ToOutput(in).(GeoFilterActionsPtrOutput)
}

func (in *geoFilterActionsPtr) ToGeoFilterActionsPtrOutputWithContext(ctx context.Context) GeoFilterActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GeoFilterActionsPtrOutput)
}

// Action to perform
type HeaderAction string

const (
	HeaderActionAppend    = HeaderAction("Append")
	HeaderActionOverwrite = HeaderAction("Overwrite")
	HeaderActionDelete    = HeaderAction("Delete")
)

// Describes operator to be matched
type HttpVersionOperator string

const (
	HttpVersionOperatorEqual = HttpVersionOperator("Equal")
)

// Describes operator to be matched
type IsDeviceOperator string

const (
	IsDeviceOperatorEqual = IsDeviceOperator("Equal")
)

// Describes if the managed rule is in enabled or disabled state. Defaults to Disabled if not specified.
type ManagedRuleEnabledState string

const (
	ManagedRuleEnabledStateDisabled = ManagedRuleEnabledState("Disabled")
	ManagedRuleEnabledStateEnabled  = ManagedRuleEnabledState("Enabled")
)

// Match variable to compare against.
type MatchVariable string

const (
	MatchVariableRemoteAddr    = MatchVariable("RemoteAddr")
	MatchVariableSocketAddr    = MatchVariable("SocketAddr")
	MatchVariableRequestMethod = MatchVariable("RequestMethod")
	MatchVariableRequestHeader = MatchVariable("RequestHeader")
	MatchVariableRequestUri    = MatchVariable("RequestUri")
	MatchVariableQueryString   = MatchVariable("QueryString")
	MatchVariableRequestBody   = MatchVariable("RequestBody")
	MatchVariableCookies       = MatchVariable("Cookies")
	MatchVariablePostArgs      = MatchVariable("PostArgs")
)

// Describes operator to be matched
type Operator string

const (
	OperatorAny                = Operator("Any")
	OperatorIPMatch            = Operator("IPMatch")
	OperatorGeoMatch           = Operator("GeoMatch")
	OperatorEqual              = Operator("Equal")
	OperatorContains           = Operator("Contains")
	OperatorLessThan           = Operator("LessThan")
	OperatorGreaterThan        = Operator("GreaterThan")
	OperatorLessThanOrEqual    = Operator("LessThanOrEqual")
	OperatorGreaterThanOrEqual = Operator("GreaterThanOrEqual")
	OperatorBeginsWith         = Operator("BeginsWith")
	OperatorEndsWith           = Operator("EndsWith")
	OperatorRegEx              = Operator("RegEx")
)

// Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
type OptimizationType string

const (
	OptimizationTypeGeneralWebDelivery          = OptimizationType("GeneralWebDelivery")
	OptimizationTypeGeneralMediaStreaming       = OptimizationType("GeneralMediaStreaming")
	OptimizationTypeVideoOnDemandMediaStreaming = OptimizationType("VideoOnDemandMediaStreaming")
	OptimizationTypeLargeFileDownload           = OptimizationType("LargeFileDownload")
	OptimizationTypeDynamicSiteAcceleration     = OptimizationType("DynamicSiteAcceleration")
)

// describes if the policy is in enabled state or disabled state
type PolicyEnabledState string

const (
	PolicyEnabledStateDisabled = PolicyEnabledState("Disabled")
	PolicyEnabledStateEnabled  = PolicyEnabledState("Enabled")
)

// Describes if it is in detection mode or prevention mode at policy level.
type PolicyMode string

const (
	PolicyModePrevention = PolicyMode("Prevention")
	PolicyModeDetection  = PolicyMode("Detection")
)

// Describes operator to be matched
type PostArgsOperator string

const (
	PostArgsOperatorAny                = PostArgsOperator("Any")
	PostArgsOperatorEqual              = PostArgsOperator("Equal")
	PostArgsOperatorContains           = PostArgsOperator("Contains")
	PostArgsOperatorBeginsWith         = PostArgsOperator("BeginsWith")
	PostArgsOperatorEndsWith           = PostArgsOperator("EndsWith")
	PostArgsOperatorLessThan           = PostArgsOperator("LessThan")
	PostArgsOperatorLessThanOrEqual    = PostArgsOperator("LessThanOrEqual")
	PostArgsOperatorGreaterThan        = PostArgsOperator("GreaterThan")
	PostArgsOperatorGreaterThanOrEqual = PostArgsOperator("GreaterThanOrEqual")
)

// Caching behavior for the requests
type QueryStringBehavior string

const (
	QueryStringBehaviorInclude    = QueryStringBehavior("Include")
	QueryStringBehaviorIncludeAll = QueryStringBehavior("IncludeAll")
	QueryStringBehaviorExclude    = QueryStringBehavior("Exclude")
	QueryStringBehaviorExcludeAll = QueryStringBehavior("ExcludeAll")
)

// Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
type QueryStringCachingBehavior string

const (
	QueryStringCachingBehaviorIgnoreQueryString = QueryStringCachingBehavior("IgnoreQueryString")
	QueryStringCachingBehaviorBypassCaching     = QueryStringCachingBehavior("BypassCaching")
	QueryStringCachingBehaviorUseQueryString    = QueryStringCachingBehavior("UseQueryString")
	QueryStringCachingBehaviorNotSet            = QueryStringCachingBehavior("NotSet")
)

func (QueryStringCachingBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryStringCachingBehavior)(nil)).Elem()
}

func (e QueryStringCachingBehavior) ToQueryStringCachingBehaviorOutput() QueryStringCachingBehaviorOutput {
	return pulumi.ToOutput(e).(QueryStringCachingBehaviorOutput)
}

func (e QueryStringCachingBehavior) ToQueryStringCachingBehaviorOutputWithContext(ctx context.Context) QueryStringCachingBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(QueryStringCachingBehaviorOutput)
}

func (e QueryStringCachingBehavior) ToQueryStringCachingBehaviorPtrOutput() QueryStringCachingBehaviorPtrOutput {
	return e.ToQueryStringCachingBehaviorPtrOutputWithContext(context.Background())
}

func (e QueryStringCachingBehavior) ToQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) QueryStringCachingBehaviorPtrOutput {
	return QueryStringCachingBehavior(e).ToQueryStringCachingBehaviorOutputWithContext(ctx).ToQueryStringCachingBehaviorPtrOutputWithContext(ctx)
}

func (e QueryStringCachingBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e QueryStringCachingBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e QueryStringCachingBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e QueryStringCachingBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type QueryStringCachingBehaviorOutput struct{ *pulumi.OutputState }

func (QueryStringCachingBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryStringCachingBehavior)(nil)).Elem()
}

func (o QueryStringCachingBehaviorOutput) ToQueryStringCachingBehaviorOutput() QueryStringCachingBehaviorOutput {
	return o
}

func (o QueryStringCachingBehaviorOutput) ToQueryStringCachingBehaviorOutputWithContext(ctx context.Context) QueryStringCachingBehaviorOutput {
	return o
}

func (o QueryStringCachingBehaviorOutput) ToQueryStringCachingBehaviorPtrOutput() QueryStringCachingBehaviorPtrOutput {
	return o.ToQueryStringCachingBehaviorPtrOutputWithContext(context.Background())
}

func (o QueryStringCachingBehaviorOutput) ToQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) QueryStringCachingBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryStringCachingBehavior) *QueryStringCachingBehavior {
		return &v
	}).(QueryStringCachingBehaviorPtrOutput)
}

func (o QueryStringCachingBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o QueryStringCachingBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QueryStringCachingBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o QueryStringCachingBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QueryStringCachingBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QueryStringCachingBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type QueryStringCachingBehaviorPtrOutput struct{ *pulumi.OutputState }

func (QueryStringCachingBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryStringCachingBehavior)(nil)).Elem()
}

func (o QueryStringCachingBehaviorPtrOutput) ToQueryStringCachingBehaviorPtrOutput() QueryStringCachingBehaviorPtrOutput {
	return o
}

func (o QueryStringCachingBehaviorPtrOutput) ToQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) QueryStringCachingBehaviorPtrOutput {
	return o
}

func (o QueryStringCachingBehaviorPtrOutput) Elem() QueryStringCachingBehaviorOutput {
	return o.ApplyT(func(v *QueryStringCachingBehavior) QueryStringCachingBehavior {
		if v != nil {
			return *v
		}
		var ret QueryStringCachingBehavior
		return ret
	}).(QueryStringCachingBehaviorOutput)
}

func (o QueryStringCachingBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QueryStringCachingBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *QueryStringCachingBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// QueryStringCachingBehaviorInput is an input type that accepts QueryStringCachingBehaviorArgs and QueryStringCachingBehaviorOutput values.
// You can construct a concrete instance of `QueryStringCachingBehaviorInput` via:
//
//	QueryStringCachingBehaviorArgs{...}
type QueryStringCachingBehaviorInput interface {
	pulumi.Input

	ToQueryStringCachingBehaviorOutput() QueryStringCachingBehaviorOutput
	ToQueryStringCachingBehaviorOutputWithContext(context.Context) QueryStringCachingBehaviorOutput
}

var queryStringCachingBehaviorPtrType = reflect.TypeOf((**QueryStringCachingBehavior)(nil)).Elem()

type QueryStringCachingBehaviorPtrInput interface {
	pulumi.Input

	ToQueryStringCachingBehaviorPtrOutput() QueryStringCachingBehaviorPtrOutput
	ToQueryStringCachingBehaviorPtrOutputWithContext(context.Context) QueryStringCachingBehaviorPtrOutput
}

type queryStringCachingBehaviorPtr string

func QueryStringCachingBehaviorPtr(v string) QueryStringCachingBehaviorPtrInput {
	return (*queryStringCachingBehaviorPtr)(&v)
}

func (*queryStringCachingBehaviorPtr) ElementType() reflect.Type {
	return queryStringCachingBehaviorPtrType
}

func (in *queryStringCachingBehaviorPtr) ToQueryStringCachingBehaviorPtrOutput() QueryStringCachingBehaviorPtrOutput {
	return pulumi.ToOutput(in).(QueryStringCachingBehaviorPtrOutput)
}

func (in *queryStringCachingBehaviorPtr) ToQueryStringCachingBehaviorPtrOutputWithContext(ctx context.Context) QueryStringCachingBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(QueryStringCachingBehaviorPtrOutput)
}

// Describes operator to be matched
type QueryStringOperator string

const (
	QueryStringOperatorAny                = QueryStringOperator("Any")
	QueryStringOperatorEqual              = QueryStringOperator("Equal")
	QueryStringOperatorContains           = QueryStringOperator("Contains")
	QueryStringOperatorBeginsWith         = QueryStringOperator("BeginsWith")
	QueryStringOperatorEndsWith           = QueryStringOperator("EndsWith")
	QueryStringOperatorLessThan           = QueryStringOperator("LessThan")
	QueryStringOperatorLessThanOrEqual    = QueryStringOperator("LessThanOrEqual")
	QueryStringOperatorGreaterThan        = QueryStringOperator("GreaterThan")
	QueryStringOperatorGreaterThanOrEqual = QueryStringOperator("GreaterThanOrEqual")
)

// The redirect type the rule will use when redirecting traffic.
type RedirectType string

const (
	RedirectTypeMoved             = RedirectType("Moved")
	RedirectTypeFound             = RedirectType("Found")
	RedirectTypeTemporaryRedirect = RedirectType("TemporaryRedirect")
	RedirectTypePermanentRedirect = RedirectType("PermanentRedirect")
)

// Describes operator to be matched
type RemoteAddressOperator string

const (
	RemoteAddressOperatorAny      = RemoteAddressOperator("Any")
	RemoteAddressOperatorIPMatch  = RemoteAddressOperator("IPMatch")
	RemoteAddressOperatorGeoMatch = RemoteAddressOperator("GeoMatch")
)

// Describes operator to be matched
type RequestBodyOperator string

const (
	RequestBodyOperatorAny                = RequestBodyOperator("Any")
	RequestBodyOperatorEqual              = RequestBodyOperator("Equal")
	RequestBodyOperatorContains           = RequestBodyOperator("Contains")
	RequestBodyOperatorBeginsWith         = RequestBodyOperator("BeginsWith")
	RequestBodyOperatorEndsWith           = RequestBodyOperator("EndsWith")
	RequestBodyOperatorLessThan           = RequestBodyOperator("LessThan")
	RequestBodyOperatorLessThanOrEqual    = RequestBodyOperator("LessThanOrEqual")
	RequestBodyOperatorGreaterThan        = RequestBodyOperator("GreaterThan")
	RequestBodyOperatorGreaterThanOrEqual = RequestBodyOperator("GreaterThanOrEqual")
)

// Describes operator to be matched
type RequestHeaderOperator string

const (
	RequestHeaderOperatorAny                = RequestHeaderOperator("Any")
	RequestHeaderOperatorEqual              = RequestHeaderOperator("Equal")
	RequestHeaderOperatorContains           = RequestHeaderOperator("Contains")
	RequestHeaderOperatorBeginsWith         = RequestHeaderOperator("BeginsWith")
	RequestHeaderOperatorEndsWith           = RequestHeaderOperator("EndsWith")
	RequestHeaderOperatorLessThan           = RequestHeaderOperator("LessThan")
	RequestHeaderOperatorLessThanOrEqual    = RequestHeaderOperator("LessThanOrEqual")
	RequestHeaderOperatorGreaterThan        = RequestHeaderOperator("GreaterThan")
	RequestHeaderOperatorGreaterThanOrEqual = RequestHeaderOperator("GreaterThanOrEqual")
)

// Describes operator to be matched
type RequestMethodOperator string

const (
	RequestMethodOperatorEqual = RequestMethodOperator("Equal")
)

// Describes operator to be matched
type RequestUriOperator string

const (
	RequestUriOperatorAny                = RequestUriOperator("Any")
	RequestUriOperatorEqual              = RequestUriOperator("Equal")
	RequestUriOperatorContains           = RequestUriOperator("Contains")
	RequestUriOperatorBeginsWith         = RequestUriOperator("BeginsWith")
	RequestUriOperatorEndsWith           = RequestUriOperator("EndsWith")
	RequestUriOperatorLessThan           = RequestUriOperator("LessThan")
	RequestUriOperatorLessThanOrEqual    = RequestUriOperator("LessThanOrEqual")
	RequestUriOperatorGreaterThan        = RequestUriOperator("GreaterThan")
	RequestUriOperatorGreaterThanOrEqual = RequestUriOperator("GreaterThanOrEqual")
)

// Name of the pricing tier.
type SkuName string

const (
	SkuName_Standard_Verizon   = SkuName("Standard_Verizon")
	SkuName_Premium_Verizon    = SkuName("Premium_Verizon")
	SkuName_Custom_Verizon     = SkuName("Custom_Verizon")
	SkuName_Standard_Akamai    = SkuName("Standard_Akamai")
	SkuName_Standard_ChinaCdn  = SkuName("Standard_ChinaCdn")
	SkuName_Standard_Microsoft = SkuName("Standard_Microsoft")
	SkuName_Premium_ChinaCdn   = SkuName("Premium_ChinaCdn")
)

// Describes what transforms are applied before matching
type Transform string

const (
	TransformLowercase   = Transform("Lowercase")
	TransformUppercase   = Transform("Uppercase")
	TransformTrim        = Transform("Trim")
	TransformUrlDecode   = Transform("UrlDecode")
	TransformUrlEncode   = Transform("UrlEncode")
	TransformRemoveNulls = Transform("RemoveNulls")
)

// Describes what transforms were applied before matching.
type TransformType string

const (
	TransformTypeLowercase   = TransformType("Lowercase")
	TransformTypeUppercase   = TransformType("Uppercase")
	TransformTypeTrim        = TransformType("Trim")
	TransformTypeUrlDecode   = TransformType("UrlDecode")
	TransformTypeUrlEncode   = TransformType("UrlEncode")
	TransformTypeRemoveNulls = TransformType("RemoveNulls")
)

// Describes operator to be matched
type UrlFileExtensionOperator string

const (
	UrlFileExtensionOperatorAny                = UrlFileExtensionOperator("Any")
	UrlFileExtensionOperatorEqual              = UrlFileExtensionOperator("Equal")
	UrlFileExtensionOperatorContains           = UrlFileExtensionOperator("Contains")
	UrlFileExtensionOperatorBeginsWith         = UrlFileExtensionOperator("BeginsWith")
	UrlFileExtensionOperatorEndsWith           = UrlFileExtensionOperator("EndsWith")
	UrlFileExtensionOperatorLessThan           = UrlFileExtensionOperator("LessThan")
	UrlFileExtensionOperatorLessThanOrEqual    = UrlFileExtensionOperator("LessThanOrEqual")
	UrlFileExtensionOperatorGreaterThan        = UrlFileExtensionOperator("GreaterThan")
	UrlFileExtensionOperatorGreaterThanOrEqual = UrlFileExtensionOperator("GreaterThanOrEqual")
)

// Describes operator to be matched
type UrlFileNameOperator string

const (
	UrlFileNameOperatorAny                = UrlFileNameOperator("Any")
	UrlFileNameOperatorEqual              = UrlFileNameOperator("Equal")
	UrlFileNameOperatorContains           = UrlFileNameOperator("Contains")
	UrlFileNameOperatorBeginsWith         = UrlFileNameOperator("BeginsWith")
	UrlFileNameOperatorEndsWith           = UrlFileNameOperator("EndsWith")
	UrlFileNameOperatorLessThan           = UrlFileNameOperator("LessThan")
	UrlFileNameOperatorLessThanOrEqual    = UrlFileNameOperator("LessThanOrEqual")
	UrlFileNameOperatorGreaterThan        = UrlFileNameOperator("GreaterThan")
	UrlFileNameOperatorGreaterThanOrEqual = UrlFileNameOperator("GreaterThanOrEqual")
)

// Describes operator to be matched
type UrlPathOperator string

const (
	UrlPathOperatorAny                = UrlPathOperator("Any")
	UrlPathOperatorEqual              = UrlPathOperator("Equal")
	UrlPathOperatorContains           = UrlPathOperator("Contains")
	UrlPathOperatorBeginsWith         = UrlPathOperator("BeginsWith")
	UrlPathOperatorEndsWith           = UrlPathOperator("EndsWith")
	UrlPathOperatorLessThan           = UrlPathOperator("LessThan")
	UrlPathOperatorLessThanOrEqual    = UrlPathOperator("LessThanOrEqual")
	UrlPathOperatorGreaterThan        = UrlPathOperator("GreaterThan")
	UrlPathOperatorGreaterThanOrEqual = UrlPathOperator("GreaterThanOrEqual")
	UrlPathOperatorWildcard           = UrlPathOperator("Wildcard")
)

func init() {
	pulumi.RegisterOutputType(GeoFilterActionsOutput{})
	pulumi.RegisterOutputType(GeoFilterActionsPtrOutput{})
	pulumi.RegisterOutputType(QueryStringCachingBehaviorOutput{})
	pulumi.RegisterOutputType(QueryStringCachingBehaviorPtrOutput{})
}