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
	value, response := n.application.HandleRequest(url, method)

	if value == 200 {
		// licence count < max licence count
		allowed := n.CheckLicenceCount(url)

		if !allowed {
			//licence count >= mac licence count
			return 403, "Not Allowed"
		}
	} else if value == 201 {
		// logut licence
		n.LogutLicenceCount(url)
	}
	return value, response
}

func (n *Nginx) CheckLicenceCount(url string) bool {
	// get once licence
	mu.Lock()
	defer mu.Unlock()
	if n.licenceCount >= n.maxLicenceCount {
		return false
	}
	n.licenceCount += 1
	return true
}

func (n *Nginx) LogutLicenceCount(url string) bool {
	// gave back licence
	mu.Lock()
	defer mu.Unlock()
	if n.licenceCount <= 0 {
		return false
	}
	n.licenceCount -= 1
	return true
}
