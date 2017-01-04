package job

import (
	"testing"
	"time"
)

func TestCalculateWorkingHour_sameMonthTopOfHour(t *testing.T) {
	// Start
	startDate := time.Date(2016, time.November, 15, 9, 0, 0, 0, time.Local)
	// End
	endDate := time.Date(2016, time.November, 30, 18, 0, 0, 0, time.Local)
	// Expect
	// (16 day * 9 ) - (4 day * 9) = 108
	expect := "108"
	actual := calculateWorkingHour(startDate, endDate)

	// Assert
	t.Log("E: " + expect)
	t.Log("A: " + actual)
	if expect != actual {
		t.Error("TestCalculateWorkingHour fail")
	}
}

func TestCalculateWorkingHour_crossMonth(t *testing.T) {
	// Start
	startDate := time.Date(2016, time.November, 29, 9, 0, 0, 0, time.Local)
	// End
	endDate := time.Date(2016, time.December, 1, 18, 0, 0, 0, time.Local)
	// Expect
	expect := "27"
	actual := calculateWorkingHour(startDate, endDate)

	// Assert
	t.Log("E: " + expect)
	t.Log("A: " + actual)
	if expect != actual {
		t.Error("TestCalculateWorkingHour fail")
	}
}

func TestCalculateWorkingHour_crossYear(t *testing.T) {
	// Start
	startDate := time.Date(2016, time.December, 28, 9, 0, 0, 0, time.Local)
	// End
	endDate := time.Date(2017, time.January, 2, 18, 0, 0, 0, time.Local)
	// Expect
	expect := "36"
	actual := calculateWorkingHour(startDate, endDate)

	// Assert
	t.Log("E: " + expect)
	t.Log("A: " + actual)
	if expect != actual {
		t.Error("TestCalculateWorkingHour fail")
	}
}

func TestCalculateWorkingHour_sameDay(t *testing.T) {
	// Start
	startDate := time.Date(2016, time.December, 28, 9, 0, 0, 0, time.Local)
	// End
	endDate := time.Date(2016, time.December, 28, 18, 0, 0, 0, time.Local)
	// Expect
	expect := "9"
	actual := calculateWorkingHour(startDate, endDate)

	// Assert
	t.Log("E: " + expect)
	t.Log("A: " + actual)
	if expect != actual {
		t.Error("TestCalculateWorkingHour fail")
	}
}

func TestCalculateWorkingHour_startDateHourLargerThanEndDateHour(t *testing.T) {
	// Start
	startDate := time.Date(2016, time.November, 15, 17, 0, 0, 0, time.Local)
	// End
	endDate := time.Date(2016, time.November, 30, 10, 0, 0, 0, time.Local)
	// Expect
	expect := "92"
	actual := calculateWorkingHour(startDate, endDate)

	// Assert
	t.Log("E: " + expect)
	t.Log("A: " + actual)
	if expect != actual {
		t.Error("TestCalculateWorkingHour fail")
	}
}

func TestCalculateWorkingHour_float(t *testing.T) {
	// Start
	startDate := time.Date(2016, time.December, 28, 9, 35, 0, 0, time.Local)
	// End
	endDate := time.Date(2017, time.January, 2, 13, 47, 0, 0, time.Local)
	// Expect
	expect := "31.2"
	actual := calculateWorkingHour(startDate, endDate)

	// Assert
	t.Log("E: " + expect)
	t.Log("A: " + actual)
	if expect != actual {
		t.Error("TestCalculateWorkingHour fail")
	}
}
