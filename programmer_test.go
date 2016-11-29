package people

import (
	"encoding/json"
	"testing"
)

func TestName(t *testing.T) {

	goodDataTests := []struct {
		data  string
		prop  string
		value interface{}
	}{
		{`{"name": "Jon Dancy"}`, "name", "Jon Dancy"},
		{`{"birthDay": "7/5/1986"}`, "birthDay", "7/5/1986"},
		{`{"favoriteLanguage": "go"}`, "favoriteLanguage", "Go"},
		{`{"favoriteLanguage": "Go"}`, "favoriteLanguage", "Go"},
		{`{"favoriteLanguage": "Go"}`, "favoriteLanguage", "Go"},
		{`{"favoriteLanguage": "javascript"}`, "favoriteLanguage", "JavaScript"},
		{`{"favoriteLanguage": "JaVAsCripT"}`, "favoriteLanguage", "JavaScript"},
		{`{"lastScore": 5}`, "lastScore", 5},
		{`{"lastScore": null}`, "lastScore", nil},
	}

	for _, test := range goodDataTests {
		var p Programmer
		var in, out []byte
		var err error

		in = []byte(test.data)
		err = json.Unmarshal(in, &p)
		if err != nil {
			t.Errorf("failed to deserialize:\n%v\nError: %v", test.data, err.Error())
			continue
		}

		out, err = json.Marshal(p)
		if err != nil {
			t.Error("Serialization failed.  This should not happen in any test.")
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

		}
	}

	// t.Run("name is good", func(t *testing.T) {
	// 	in = []byte(`{"name": "Jon Dancy"}`)

	// 	err = json.Unmarshal(in, &p)
	// 	checkDeserialize(t, err, in)

	// 	out, err = json.Marshal(p)
	// 	checkSerialize(t, err, out)
	// })

	// t.Run("name is bad", func(t *testing.T) {
	// 	in = []byte(`{"name": 123}`)

	// 	err = json.Unmarshal(in, &p)
	// 	return
	// })

}

// func checkDeserialize(t *testing.T, err error, in []byte) {
// 	if err != nil {
// 		t.Error(string(in), "should not have failed to deserialize")
// 	}
// }

// func checkSerialize(t *testing.T, err error, out []byte) {
// 	if err != nil {
// 		t.Error(string(out), "should not have failed to serialize")
// 	}
// }
