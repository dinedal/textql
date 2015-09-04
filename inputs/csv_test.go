package inputs

import (
	"os"
	"strings"

	"reflect"
	"testing"

	"github.com/dinedal/textql/test_util"
)

var (
	simple = `a,b,c
1,2,3
4,5,6`

	bad = `a,b,c
1,2,
4,5,6
7,8


9,,10
11,12,13,14
"foo,bar","boo,\"far",","
'foo,bar','"','"'
"test
",multi-line
`
)

func TestCSVInputFakesHeader(t *testing.T) {
	fp := test_util.OpenFileFromString(simple)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &CSVInputOptions{
		HasHeader: false,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input, _ := NewCSVInput(opts)
	expected := []string{"c0", "c1", "c2"}

	if !reflect.DeepEqual(input.Header(), expected) {
		t.Errorf("Header() = %v, want %v", input.Header(), expected)
	}
}

func TestCSVInputReadsHeader(t *testing.T) {
	fp := test_util.OpenFileFromString(simple)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input, _ := NewCSVInput(opts)
	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(input.Header(), expected) {
		t.Errorf("Header() = %v, want %v", input.Header(), expected)
	}
}

func TestCSVInputReadsSimple(t *testing.T) {
	fp := test_util.OpenFileFromString(simple)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input, _ := NewCSVInput(opts)
	expected := make([][]string, len(strings.Split(simple, "\n"))-1)
	expected[0] = []string{"1", "2", "3"}
	expected[1] = []string{"4", "5", "6"}

	for counter := 0; counter < len(expected); counter++ {
		row := input.ReadRecord()
		if !reflect.DeepEqual(row, expected[counter]) {
			t.Errorf("ReadRecord() = %v, want %v", row, expected[counter])
		}
	}
}

func TestCSVInputReadsBad(t *testing.T) {
	fp := test_util.OpenFileFromString(bad)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input, _ := NewCSVInput(opts)
	expected := make([][]string, len(strings.Split(bad, "\n"))-1)
	expected[0] = []string{"1", "2", ""}
	expected[1] = []string{"4", "5", "6"}
	expected[2] = []string{"7", "8", ""}
	expected[3] = []string{"9", "", "10"}
	expected[4] = []string{"11", "12", "13", "14"}
	expected[5] = []string{"foo,bar", `boo,\"far`, ","}
	expected[6] = []string{`'foo`, `bar'`, `'"'`, `'"'`}
	expected[7] = []string{"test\n", "multi-line", ""}

	for counter := 0; counter < len(expected); counter++ {
		row := input.ReadRecord()
		if !reflect.DeepEqual(row, expected[counter]) {
			t.Errorf("ReadRecord() = %v, want %v", row, expected[counter])
		}
	}
}

func TestCSVInputHasAName(t *testing.T) {
	fp := test_util.OpenFileFromString(simple)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input, _ := NewCSVInput(opts)
	expected := fp.Name()

	if !reflect.DeepEqual(input.Name(), expected) {
		t.Errorf("Name() = %v, want %v", input.Name(), expected)
	}
}
