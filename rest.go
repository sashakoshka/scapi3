package scapi3

import "fmt"
import "net/http"
// import "time"
// import "strconv"

/* GetHealth returns information relating to the health of the scratch website.
 */
func GetHealth () (structure HealthResponse, err error) {
	response, body, err := Request {
		Path:     "/health",
		Hostname: "api.scratch.mit.edu",
	}.Send()
	if err != nil { return }
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf (
			"cannot get site health (%i)",
			response.StatusCode)
		return
	}
	
	structure, err = UnmarshalHealthResponse(body)
	if err != nil { return }
	return
}
