package scapi3

import "fmt"
import "net/http"
import "encoding/json"

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
	
	err = json.Unmarshal(body, &structure)
	if err != nil { return }
	return
}

/* GetNews returns information relating to the health of the scratch website.
 */
func GetNews () (structure NewsResponse, err error) {
	response, body, err := Request {
		Path:     "/news",
		Hostname: "api.scratch.mit.edu",
	}.Send()
	if err != nil { return }
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf (
			"cannot get site news (%i)",
			response.StatusCode)
		return
	}
	
	err = json.Unmarshal(body, &structure)
	if err != nil { return }
	return
}
