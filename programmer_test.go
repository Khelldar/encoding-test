package people

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestProgrammerEncoding(t *testing.T) {
	var testPass, testTotal int
	goodDataTests := []struct {
		data          string
		prop          string
		value         interface{}
		expectedError bool
	}{
		{`{"name": "Jon Dancy"}`, "name", "Jon Dancy", false},
		{`{"birthday": "7/25/1986"}`, "birthday", "7/25/1986", false},
		{`{"favoriteLanguage": "go"}`, "favoriteLanguage", "Go", false},
		{`{"favoriteLanguage": "Go"}`, "favoriteLanguage", "Go", false},
		{`{"favoriteLanguage": "Go"}`, "favoriteLanguage", "Go", false},
		{`{"favoriteLanguage": "javascript"}`, "favoriteLanguage", "JavaScript", false},
		{`{"favoriteLanguage": "JaVAsCripT"}`, "favoriteLanguage", "JavaScript", false},
		{`{"lastScore": 5}`, "lastScore", float64(5), false},
		{`{"lastScore": 0}`, "lastScore", float64(0), false},
		{`{"lastScore": null}`, "lastScore", nil, false},
		{`{"name": false}`, "", nil, true},
		{`{"birthday": "7/25/1986-it was pretty sunny that day"}`, "", nil, true},
		{`{"favoriteLanguage": "pearl"}`, "", nil, true},
		{`{"favoriteLanguage": 0}`, "", nil, true},
		{`{"lastScore": -1}`, "", nil, true},
	}

	for _, test := range goodDataTests {
		testTotal++
		var p Programmer
		var in, out []byte
		var err error

		in = []byte(test.data)
		err = json.Unmarshal(in, &p)

		//got an error when we shouldn't have
		if err != nil && !test.expectedError {
			t.Errorf("failed to deserialize:\n%v\nError: %v", test.data, err.Error())
			continue
		}

		if err == nil && test.expectedError {
			t.Errorf("expected to see an error given this bad input:\n %v", test.data)
			continue
		}

		if err != nil && test.expectedError {
			testPass++
			continue
		}

		out, err = json.Marshal(&p)
		if err != nil {
			t.Error("Serialization failed.  This should not happen in any test. Error: ", err)
			continue
		}

		tester := map[string]interface{}{}
		err = json.Unmarshal(out, &tester)
		if err != nil {
			t.Error("Unable to test deserialization.  It should not look like this:", string(out))
			continue
		}

		if tester[test.prop] != test.value {
			t.Errorf("%v data was modified with input of:\n%v\n'%v' should be '%v' but deserialized into '%v'\nfull dump of deserialized data:\n%v", test.prop, test.data, test.prop, test.value, tester[test.prop], string(out))
			continue
		}
		testPass++
	}

	fmt.Printf("%v/%v good input tests passing.\n", testPass, testTotal)
	if testPass != testTotal {
		fmt.Printf("Failing tests:\n")
	}

}
