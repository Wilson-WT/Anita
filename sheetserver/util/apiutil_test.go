package util

import "testing"
import "reflect"

//// ArrangeData
func TestArrangeData(t *testing.T) {
	// Prepare test data
	reports := [][]interface{}{[]interface{}{"2016/7/11 下午 10:02:46", "henry@happygorgi.com", "AEP", "AEP 2.7 架構討論", "1", "2016/8/4"},
		[]interface{}{"2017/7/11 下午 10:02:46", "timmy@happygorgi.com", "AFU", "AFU 討論", "2", "2017/7/5"},
		[]interface{}{"2017/7/12 下午 10:02:22", "jay@happygorgi.com", "AFU", "AFU 討論", "3", "2017/7/6"}}
	// Test
	resultString := "###### AEP ######\n" +
		"[2016/08]\n" +
		"1. AEP 2.7 架構討論\n" +
		"\n" +
		"###### AFU ######\n" +
		"[2017/07]\n" +
		"1. AFU 討論\n" +
		"2. AFU 討論\n" +
		"\n"
	arrangedData := ArrangeData(reports, "201608-201707")
	// assert
	if resultString != arrangedData {
		t.Log("resultString: \n" + resultString)
		t.Log("arrangedData: \n" + arrangedData)
		t.Error("TestArrangeData Fail")
	} else {
		t.Log("resultString: \n" + resultString)
		t.Log("arrangedData: \n" + arrangedData)
	}
}

//// getRecordsByYearAndMonth
func TestGetRecordsByYearAndMonth(t *testing.T) {
	// Prepare test data
	reports := [][]interface{}{[]interface{}{"2016/7/11 下午 10:02:46", "henry@happygorgi.com", "AEP", "AEP 2.7 架構討論", "1", "2016/7/4"},
		[]interface{}{"2016/5/12 下午 10:02:46", "timmy@happygorgi.com", "AFU", "AFU 討論", "2", "2016/5/5"},
		[]interface{}{"2016/9/13 下午 10:02:22", "jay@happygorgi.com", "AFU", "AFU 討論", "3", "2016/6/6"}}
	// Test
	resultReports := getRecordsByYearAndMonth(reports, 2016, 5)
	// assert
	if !reflect.DeepEqual(resultReports[0], reports[1]) {
		t.Error("Result not correct")
	}
}

func TestGetRecordsByYearAndMonth_invalidDoneDate(t *testing.T) {
	// Prepare test data
	reports := [][]interface{}{[]interface{}{"2016/7/11 下午 10:02:46", "henry@happygorgi.com", "AEP", "AEP 2.7 架構討論", "1", "yyyy/mm/dd"},
		[]interface{}{"2016/5/12 下午 10:02:46", "timmy@happygorgi.com", "AFU", "AFU 討論", "2", "2016/5/5"},
		[]interface{}{"2016/9/13 下午 10:02:22", "jay@happygorgi.com", "AFU", "AFU 討論", "3", "2016/6/6"}}
	// Test
	resultReports := getRecordsByYearAndMonth(reports, 2016, 5)
	// assert
	if !reflect.DeepEqual(resultReports[0], reports[1]) {
		t.Error("Result not correct")
	}
}

//// getReportsByProject
func TestGetReportsByProject(t *testing.T) {
	// Prepare test data
	reports := [][]interface{}{[]interface{}{"2016/7/11 下午 10:02:46", "henry@happygorgi.com", "AEP", "AEP 2.7 架構討論", "1", "2016/7/4"},
		[]interface{}{"2016/7/11 下午 10:02:46", "timmy@happygorgi.com", "AFU", "AFU 討論", "2", "2016/7/5"},
		[]interface{}{"2016/7/12 下午 10:02:22", "jay@happygorgi.com", "AFU", "AFU 討論", "3", "2016/7/6"}}
	// Test
	AEPProjects := getReportsByProject(reports, "AEP")
	AFUProjects := getReportsByProject(reports, "AFU")
	// assert length
	if len(AEPProjects) != 1 && len(AFUProjects) != 2 {
		t.Error("len not correct")
	}
	// assert content
	if !reflect.DeepEqual(AEPProjects[0], reports[0]) {
		t.Error("content not correct")
	} else if !reflect.DeepEqual(AFUProjects[0], reports[1]) {
		t.Error("content not correct")
	} else if !reflect.DeepEqual(AFUProjects[1], reports[2]) {
		t.Error("content not correct")
	}
}

//// getProjectsInReports Test
func TestGetProjectsInReports(t *testing.T) {
	// Prepare test data
	reports := [][]interface{}{[]interface{}{"2016/7/11 下午 10:02:46", "henry.hsieh@happygorgi.com", "AEP", "AEP 2.7 架構討論", "1", "2016/7/4"},
		[]interface{}{"2016/7/11 下午 10:02:46", "henry.hsieh@happygorgi.com", "AFU", "AEP 2.7 架構討論", "1", "2016/7/4"}}
	// call getProjectsInReports
	projects := getProjectsInReports(reports)
	// assert
	if projects[0] == "AEP" && projects[1] == "AFU" {
		t.Log("getProjectsInReports PASS")
	} else {
		t.Error("getProjectsInReports FAIL")
	}
}

func TestGetProjectsInReports_invalidProjectName(t *testing.T) {
	// Prepare test data
	reports := [][]interface{}{[]interface{}{"2016/7/11 下午 10:02:46", "henry.hsieh@happygorgi.com", nil, "AEP 2.7 架構討論", "1", "2016/7/4"},
		[]interface{}{"2016/7/11 下午 10:02:46", "henry.hsieh@happygorgi.com", "AFU", "AEP 2.7 架構討論", "1", "2016/7/4"}}
	// call getProjectsInReports
	projects := getProjectsInReports(reports)
	// assert
	if len(projects) == 1 && projects[0] == "AFU" {
		t.Log("getProjectsInReports PASS")
	} else {
		t.Error("getProjectsInReports FAIL")
	}
}

//// isProjectInArray Test
func TestIsProjectInArray_true(t *testing.T) {
	projects := []string{"project1", "project2", "project3", "project4"}
	isExist := isProjectInArray(projects, "project1")
	if isExist {
		t.Log("TestIsProjectInArray PASS")
	} else {
		t.Error("TestIsProjectInArray FAIL")
	}
}

func TestIsProjectInArray_false(t *testing.T) {
	projects := []string{"project1", "project2", "project3", "project4"}
	isExist := isProjectInArray(projects, "project0")
	if !isExist {
		t.Log("TestIsProjectInArray PASS")
	} else {
		t.Error("TestIsProjectInArray FAIL")
	}
}

func TestIsProjectInArray_empty(t *testing.T) {
	projects := []string{}
	isExist := isProjectInArray(projects, "project1")
	if !isExist {
		t.Log("TestIsProjectInArray PASS")
	} else {
		t.Error("TestIsProjectInArray FAIL")
	}
}

func TestIsDoneDateValid_Valid(t *testing.T) {
	if isDoneDateValid("201607-201608", "2016/7/12") {
		t.Log("isDoneDateValid PASS")
	} else {
		t.Error("isDoneDateValid FAIL")
	}
}

func TestIsDoneDateValid_Invalid(t *testing.T) {
	// period = "201607-201608", done date = "2016/9/15"
	if !isDoneDateValid("201607-201608", "2016/9/15") {
		t.Log("isDoneDateValid PASS")
	} else {
		t.Error("isDoneDateValid FAIL")
	}

	// period = "201607-201608", done date = "2016/7"
	if !isDoneDateValid("201607-201608", "2016/7") {
		t.Log("isDoneDateValid PASS")
	} else {
		t.Error("isDoneDateValid FAIL")
	}
}

func TestIsPeriodValid_Valid(t *testing.T) {
	// period = "201608-201808"
	if isPeriodValid("201608-201808") {
		t.Log("isPeriodValid PASS")
	} else {
		t.Error("isPeriodValid FAIL")
	}

	// period = "201607-201607"
	if isPeriodValid("201607-201607") {
		t.Log("isPeriodValid PASS")
	} else {
		t.Error("isPeriodValid FAIL")
	}

	// period = "201607-201608"
	if isPeriodValid("201607-201608") {
		t.Log("isPeriodValid PASS")
	} else {
		t.Error("isPeriodValid FAIL")
	}
}

func TestIsPeriodValid_Invalid(t *testing.T) {
	// period = "201708-201609"
	if !isPeriodValid("201708-201609") {
		t.Log("isPeriodValid PASS")
	} else {
		t.Error("isPeriodValid FAIL")
	}

	// period = "201608-201607"
	if !isPeriodValid("201608-201607") {
		t.Log("isPeriodValid PASS")
	} else {
		t.Error("isPeriodValid FAIL")
	}
}
