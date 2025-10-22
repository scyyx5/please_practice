
package main

import (
	_gostdlib_os "os"
	_gostdlib_strings "strings"
	_gostdlib_testing "testing"
	_gostdlib_testdeps "testing/internal/testdeps"


	greetings_test "greetings_test_lib"

)

var tests = []_gostdlib_testing.InternalTest{

	{"TestGreeting", greetings_test.TestGreeting},

}
var examples = []_gostdlib_testing.InternalExample{

}

var benchmarks = []_gostdlib_testing.InternalBenchmark{

}

var fuzzTargets = []_gostdlib_testing.InternalFuzzTarget{

}

var testDeps = _gostdlib_testdeps.TestDeps{}

func internalMain() int {
	args := []string{_gostdlib_os.Args[0], "-test.v"}


    testVar := _gostdlib_os.Getenv("TESTS")
    if testVar != "" {
		testVar = _gostdlib_strings.ReplaceAll(testVar, " ", "|")
		args = append(args, "-test.run", testVar)
    }
    _gostdlib_os.Args = append(args, _gostdlib_os.Args[1:]...)
	m := _gostdlib_testing.MainStart(testDeps, tests, nil, fuzzTargets, examples)



	return m.Run()

}

func main() {
	_gostdlib_os.Exit(internalMain())
}
