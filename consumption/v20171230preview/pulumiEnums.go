// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171230preview

// The category of the budget, whether the budget tracks cost or something else.
type CategoryType string

const (
	CategoryTypeCost = CategoryType("Cost")
)

// The comparison operator.
type OperatorType string

const (
	OperatorTypeEqualTo              = OperatorType("EqualTo")
	OperatorTypeGreaterThan          = OperatorType("GreaterThan")
	OperatorTypeGreaterThanOrEqualTo = OperatorType("GreaterThanOrEqualTo")
)

// The time covered by a budget. Tracking of the amount will be reset based on the time grain.
type TimeGrainType string

const (
	TimeGrainTypeMonthly   = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually  = TimeGrainType("Annually")
)

func init() {
}
