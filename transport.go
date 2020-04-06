package shell

import "net/http"

type transport struct {
	token        string
	httptr       http.RoundTripper
	extraHeaders map[string]string
}

func newAuthenticatedTransport(tr http.RoundTripper, token string, extraHeaders map[string]string) *transport {
	return &transport{
		token:        token,
		httptr:       tr,
		extraHeaders: extraHeaders,
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.extraHeaders != nil {
		for k, v := range t.extraHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.httptr.RoundTrip(req)
}
