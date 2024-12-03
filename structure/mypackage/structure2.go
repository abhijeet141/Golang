package mypackage

type Student struct {
	Name   string
	Branch string
	Year   int
}
type Teacher struct {
	Name    string
	Subject string
	Exp     int
	Details Student
}
