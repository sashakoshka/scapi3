package scapi3

import "io"
import "bytes"
import "net/http"
import "io/ioutil"

type Method int

const (
	MethodGet = iota
	MethodPost
)

/* Request represents an HTTP request made to the Scratch API.
 */
type Request struct {
	Hostname	string
	Path		string
	Method		Method
	Headers		map[string] string
	Body		ProtocolRequest
	Client		*http.Client
}

/* Send sends the request to the Scratch servers, and returns the response.
 */
func (request Request) Send () (
	response	*http.Response,
	body		[]byte,
	err		error,
) {
	// create request
	var method string
	if request.Method == MethodGet {
		method = "GET"
	} else {
		method = "POST"
	}
	if request.Hostname == "" {
		request.Hostname = "scratch.mit.edu"
	}
	
	var requestBody io.Reader
	if request.Body != nil {
		requestBody = bytes.NewBuffer(request.Body.Marshal())
	}
	var httpRequest *http.Request
	httpRequest, err = http.NewRequest (
		method,
		"https://" + request.Hostname + request.Path,
		requestBody)

	// set request headers
	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}
	httpRequest.Header.Set("X-CSRFToken", "a")
	httpRequest.Header.Set("Referer", "https://scratch.mit.edu")
        httpRequest.Header.Add("Cookie", "scratchcsrftoken=a; scratchlanguage=en;")
	httpRequest.Header.Set("User-Agent", "")

	// perform request
	if request.Client == nil {
		response, err = http.DefaultClient.Do(httpRequest)
	} else {
		response, err = request.Client.Do(httpRequest)
	}
	defer response.Body.Close()
	
	// read response body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil { return }
	return
}
