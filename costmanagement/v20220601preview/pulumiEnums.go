// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601preview

// Days of Week.
type DaysOfWeek string

const (
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
)

// Destination of the view data. Currently only csv format is supported.
type FileFormat string

const (
	FileFormatCsv = FileFormat("Csv")
)

// Frequency of the schedule.
type ScheduleFrequency string

const (
	// Cost analysis data will be emailed every day.
	ScheduleFrequencyDaily = ScheduleFrequency("Daily")
	// Cost analysis data will be emailed every week.
	ScheduleFrequencyWeekly = ScheduleFrequency("Weekly")
	// Cost analysis data will be emailed every month.
	ScheduleFrequencyMonthly = ScheduleFrequency("Monthly")
)

// Kind of the scheduled action.
type ScheduledActionKind string

const (
	// Cost analysis data will be emailed.
	ScheduledActionKindEmail = ScheduledActionKind("Email")
	// Cost anomaly information will be emailed. Available only on subscription scope at daily frequency. If no anomaly is detected on the resource, an email won't be sent.
	ScheduledActionKindInsightAlert = ScheduledActionKind("InsightAlert")
)

// Status of the scheduled action.
type ScheduledActionStatus string

const (
	// Scheduled action is saved but will not be executed.
	ScheduledActionStatusDisabled = ScheduledActionStatus("Disabled")
	// Scheduled action is saved and will be executed.
	ScheduledActionStatusEnabled = ScheduledActionStatus("Enabled")
)

// Weeks of month.
type WeeksOfMonth string

const (
	WeeksOfMonthFirst  = WeeksOfMonth("First")
	WeeksOfMonthSecond = WeeksOfMonth("Second")
	WeeksOfMonthThird  = WeeksOfMonth("Third")
	WeeksOfMonthFourth = WeeksOfMonth("Fourth")
	WeeksOfMonthLast   = WeeksOfMonth("Last")
)

func init() {
}
