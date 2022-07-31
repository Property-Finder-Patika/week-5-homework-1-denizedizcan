package funcs

import "sync"

var (
	mu sync.Mutex
)

//proxy
type Nginx struct {
	application     *Application
	maxLicenceCount int
	licenceCount    int
}

//create proxy
func NewNginxServer() *Nginx {
	return &Nginx{
		application:     &Application{},
		maxLicenceCount: 10,
		licenceCount:    0,
	}
}

//proxy handle request
func (n *Nginx) HandleRequest(url, method string) (int, string) {

	// licence count < max licence count
	allowed := n.CheckLicenceCount(url)

	if allowed == "error" {
		//licence count >= max licence count
		return 403, "Not Allowed"
	}
	// logut licence
	n.LogutLicenceCount(url)

	return n.application.HandleRequest(url, method)
}

func (n *Nginx) CheckLicenceCount(url string) string {
	// get once licence
	mu.Lock()
	defer mu.Unlock()
	if n.licenceCount < n.maxLicenceCount && url == "/licence" {
		n.licenceCount += 1
		return "getLicence"
	} else if url == "/licence/loguot" {
		return "pass"
	}
	return "error"
}

func (n *Nginx) LogutLicenceCount(url string) bool {
	// gave back licence
	mu.Lock()
	defer mu.Unlock()
	if n.licenceCount > 0 && url == "/licence/loguot" {
		n.licenceCount -= 1
		return true
	}
	return false
}
