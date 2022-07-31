package funcs

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	licenceCount      int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 10,
		licenceCount:      0,
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkLicenceCount(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkLicenceCount(url string) bool {

	if n.licenceCount > n.maxAllowedRequest {
		return false
	}
	n.licenceCount += 1
	return true
}
