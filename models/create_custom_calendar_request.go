package models

type CreateCustomCalendarRequest struct {
	Name              string                       `json:"name"`
	CreationMethod    string                       `json:"creation_method"`
	TableReference    CustomCalendarTableReference `json:"table_reference"`
	StartDate         string                       `json:"start_date,omitempty"`
	EndDate           string                       `json:"end_date,omitempty"`
	CalendarType      string                       `json:"calendar_type,omitempty"`
	MonthOffset       string                       `json:"month_offset,omitempty"`
	StartDayOfWeek    string                       `json:"start_day_of_week,omitempty"`
	QuarterNamePrefix string                       `json:"quarter_name_prefix,omitempty"`
	YearNamePrefix    string                       `json:"year_name_prefix,omitempty"`
}
