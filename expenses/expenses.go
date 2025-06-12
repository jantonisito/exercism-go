package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filtered := []Record{}
	for _, r := range in {
		if predicate(r) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return c == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	filter_by_day := ByDaysPeriod(p)
	filtered_in := Filter(in, filter_by_day)
	for _, r := range filtered_in {
		total += r.Amount
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filter_by_cat := ByCategory(c)
	in_by_cat := Filter(in, filter_by_cat)

	if len(in_by_cat) == 0 {
		return 0, errors.New("unknown category " + c)
	}

	filter_by_day := ByDaysPeriod(p)
	in_by_cat_day := Filter(in_by_cat, filter_by_day)

	total := 0.0
	for _, r := range in_by_cat_day {
		total += r.Amount
	}
	return total, nil
}
