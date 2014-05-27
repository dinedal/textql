package inputs

type Input interface {
	ReadRecord() []string
	Header() []string
}
