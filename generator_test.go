package main

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/carlmjohnson/csv"
)

func TestHelloName(t *testing.T) {
	alleIdentitaeten := generate(20, "test.test", 5)
	fmt.Println(alleIdentitaeten)
	t.Fatal("work in progress")
}

func ExampleFieldReader_options() {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewFieldReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	for r.Scan() {
		fmt.Println(r.Field("username"))
	}

	if err := r.Err(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// rob
	// ken
	// gri
}
