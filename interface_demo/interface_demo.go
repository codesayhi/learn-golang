package interface_demo

type Animal interface {
	Run()
	Speak()
}
type Reader interface {
	Read(string)
}

type Writer interface {
	Write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}
