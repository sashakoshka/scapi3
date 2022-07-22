package scapi3

import "io"
import "bytes"
import "net/http"
import "io/ioutil"
// import "net/http/httputil"

type Method int

const (
	MethodGet = iota
	MethodPost
	MethodOptions
)

/* Request represents an HTTP request made to the Scratch API.
 */
type Request struct {
	Hostname	string
	Path		string
	Method		Method
	Headers		map[string] string
	Body		RequestBody
	SessionID	string
}

/* RequestBody represents a structure that can be marshalled into a byte
 * slice.
 */
type RequestBody interface {
	Marshal () (data []byte)
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
	} else if request.Method == MethodPost {
		method = "POST"
	} else {
		method = "OPTIONS"
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
	if err != nil { return }

	// set request headers
	httpRequest.Header.Set("X-CSRFToken", "a")
	httpRequest.Header.Set("Referer", "https://scratch.mit.edu")
	httpRequest.Header.Set("User-Agent", "")
	for key, value := range request.Headers {
		httpRequest.Header.Set(key, value)
	}
	cookies := "scratchcsrftoken=a; scratchlanguage=en;"
	if request.SessionID != "" {
		cookies += " scratchsessionsid=\"" + request.SessionID + "\";"
	}
        httpRequest.Header.Add("Cookie", cookies)

	// dump, _ := httputil.DumpRequestOut(httpRequest, false)
	// println(string(dump))

	// perform request
	response, err = http.DefaultClient.Do(httpRequest)
	defer response.Body.Close()
	
	// dump, _ = httputil.DumpResponse(response, false)
	// println(string(dump))
	
	// read response body
	body, err = ioutil.ReadAll(response.Body)
	if err != nil { return }
	return
}
