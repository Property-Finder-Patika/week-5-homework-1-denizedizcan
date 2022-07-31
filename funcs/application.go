package funcs

type Application struct {
}

//Aplication requests control
func (a *Application) HandleRequest(url, method string) (int, string) {
	if url == "/licence" && method == "GET" {
		return 200, "Ok"
	} else if url == "/licence/loguot" && method == "POST" {
		return 201, "Succesful Logout"
	}
	return 404, "Not Ok"
}
