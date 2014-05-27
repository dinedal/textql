package inputs

import (
	"github.com/dinedal/textql/inputs"
	"io/ioutil"
	"os"
	"strings"

	"reflect"
	"testing"
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

func csvOpened(contents string) os.File {
	f, _ := ioutil.TempFile("./", "csv")
	f.WriteString(contents)
	f.Seek(0, 0)
	return *f
}

func TestCSVInputFakesHeader(t *testing.T) {
	fp := csvOpened(simple)

	opts := &inputs.CSVInputOptions{
		HasHeader: false,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)
	expected := []string{"c0", "c1", "c2"}

	if !reflect.DeepEqual(input.Header(), expected) {
		t.Errorf("Header() = %v, want %v", input.Header(), expected)
	}

	defer fp.Close()
	defer os.Remove(fp.Name())
}

func TestCSVInputReadsHeader(t *testing.T) {
	fp := csvOpened(simple)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)
	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(input.Header(), expected) {
		t.Errorf("Header() = %v, want %v", input.Header(), expected)
	}
}

func TestCSVInputReadsSimple(t *testing.T) {
	fp := csvOpened(simple)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)
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
	fp := csvOpened(bad)
	defer fp.Close()
	defer os.Remove(fp.Name())

	opts := &inputs.CSVInputOptions{
		HasHeader: true,
		Seperator: ',',
		ReadFrom:  fp,
	}

	input := inputs.NewCSVInput(opts)
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
