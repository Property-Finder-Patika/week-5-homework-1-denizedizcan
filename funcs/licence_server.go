package funcs

type server interface {
	handleRequest(string, string) (int, string)
}
