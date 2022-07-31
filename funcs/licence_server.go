package funcs

//server interface
type Server interface {
	HandleRequest(string, string) (int, string)
}
