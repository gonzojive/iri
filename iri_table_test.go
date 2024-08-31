package iri

import (
	"encoding/json"
	"fmt"
	"testing"

	_ "embed"
)

//go:embed testdata/iris.json
var irisJSON string

var suite *IRITestSuite = must(loadFromJSON[IRITestSuite]([]byte(irisJSON)))

func TestIRI_AgainstSuite(t *testing.T) {
	for _, group := range suite.Suite.Groups {
		switch group.Name {
		case "anchor":
			{

				for _, ex := range group.Examples {
					t.Run(fmt.Sprintf("%s/%s", group.Name, ex.ID), func(t *testing.T) {
						parsed, err := Parse(ex.URL)
						if err != nil {
							t.Fatal(err)
						}
						t.Log("got parsed url $s", parsed)
					})
				}
			}
		}

	}
}

// IRI test suite struct represents the structure of the test data loaded from JSON.
// This data originates from the https://github.com/cweb/iri-tests repository.
type IRITestSuite struct {
	Suite struct {
		Description string `json:"desc"` // Description of the overall test suite
		Groups      []struct {
			Name        string    `json:"name"` // Name of the test group
			Link        string    `json:"link"` // Optional link for further information
			Description string    `json:"desc"` // Description of the test group
			Examples    []Example `json:"test"`
		} `json:"group"`
	} `json:"tests"`
}

// Example represents a single test case in the IRI test suite.
//
// Fields are populated based on the specific test case.
//
// If the test case involves a relative URL, `Base` and `Rel` will be set, with
// `ExpectRel` as the expected output.
//
// For absolute URLs, `URL` is set along with expected components (`ExpectURL`, `ExpectScheme`, etc.).
//
// Original source: https://github.com/cweb/iri-tests
type Example struct {
	// ID of the test case.
	ID string `json:"id"`

	// Name of the test case.
	Name string `json:"name"`

	// Base URL for the test, if applicable.
	Base string `json:"base,omitempty"`

	// Input URL for the test.
	URL string `json:"url,omitempty"`

	// Relative URL for the test, if applicable.
	Rel string `json:"rel,omitempty"`

	// Expected output URL after parsing.
	ExpectURL string `json:"expect_url,omitempty"`

	// Expected relative URL after parsing.
	ExpectRel string `json:"expect_rel,omitempty"`

	// Expected scheme component of the URL.
	ExpectScheme string `json:"expect_scheme,omitempty"`

	// Expected host component of the URL.
	ExpectHost string `json:"expect_host,omitempty"`

	// Expected port component of the URL.
	ExpectPort string `json:"expect_port,omitempty"`

	// Expected path component of the URL.
	ExpectPath string `json:"expect_path,omitempty"`

	// Expected query component of the URL.
	ExpectQuery string `json:"expect_query,omitempty"`

	// Expected fragment component of the URL.
	ExpectFragment string `json:"expect_fragment,omitempty"`
}

// loadFromJSON parses the provided JSON text into an instance of T.
func loadFromJSON[T any](data []byte) (*T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	return &t, err
}

// must is a helper function that takes a value and an error.
// If the error is non-nil, must panics. Otherwise, it returns the value.
func must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
