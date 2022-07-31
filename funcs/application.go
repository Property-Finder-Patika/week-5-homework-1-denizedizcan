package funcs

type Application struct {
}

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == "/licence" && method == "GET" {
		return 200, "Ok"
	}

	return 404, "Not Ok"
}
