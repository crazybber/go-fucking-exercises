package more

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
)

var t *testing.T

func init() {

	t = new(testing.T)
}

//PShow have a print
func PShow(tp ...interface{}) {
	fmt.Printf("Testing : %v \n", tp)
}

//Show have a t.log
func Show(t *testing.T, tp interface{}) {
	t.Logf("Testing Information : %v", tp)
}

//Expect need expression which will return true
func Expect(expression bool) {
	if !expression {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("Failed!!!!!! ----------------------------------------------------------> expression for [%v]", expression)
		targetString := fmt.Sprintf("Failed!!!!!! in file: %s, line: %d ------> expression for [%v]", file, line, expression)
		fmt.Println(aurora.Red(targetString))
		//fmt.Println(targetString)
	} else {
		t.Logf("OK!!! eq in expression: [%v]", expression)
	}

}

//ExpectEq ..
func ExpectEq(expected interface{}, actual interface{}) {
	AssertEq(expected, actual)
}

//ExpectNotEq ..
func ExpectNotEq(expected interface{}, actual interface{}) {
	AssertNotEq(expected, actual)
}

//AssertEq ..
func AssertEq(expected interface{}, actual interface{}) {

	if !assert.Equal(t, expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		targetString := fmt.Sprintf("Failed!!!!!! in file: %s, line: %d -----> expect:[%v] , actual: [%v] \n", file, line, expected, actual)
		fmt.Println(aurora.Red(targetString))
		//fmt.Println(targetString)
		t.Logf("Failed!!!!!! ---------------------------------------------------------->  expect:[%v] , actual: [%v] \n", expected, actual)
	} else {

		t.Logf("OK!!! expect [%v] eq expression: [%v]", expected, actual)
		targetString := fmt.Sprintf("OK!!! expect [%v] eq expression: [%v] \n", expected, actual)
		fmt.Println(aurora.Green(targetString))
	}
}

//AssertNotEq ..
func AssertNotEq(expected interface{}, actual interface{}) {
	if !assert.NotEqual(t, expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("Failed!!!!!! in file: %s, line: %d-----> expect:[%v] , actual: [%v] \n", file, line, expected, actual)
		t.Logf("Failed!!!!!! in file: %s, line: %d------>  expect:[%v] , actual: [%v] \n", file, line, expected, actual)
	} else {
		t.Logf("OK!!! expect [%v] not eq expression: [%v]", expected, actual)
		fmt.Printf("OK!!! expect [%v] not eq expression: [%v] \n", expected, actual)
	}
}
