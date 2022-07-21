package scapi3

import "fmt"
import "strconv"
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

/* GetProjectsCountAll returns the amount of projects that have been uploaded to
 * the site.
 */
func GetProjectsCountAll () (count uint64, err error) {
	response, body, err := Request {
		Path:     "/projects/count/all",
		Hostname: "api.scratch.mit.edu",
	}.Send()
	if err != nil { return }
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf (
			"cannot get projects count (%i)",
			response.StatusCode)
		return
	}

	structure := ProjectsCountAllResponse { }
	err = json.Unmarshal(body, &structure)
	if err != nil { return }
	count = structure.Count
	return
}

/* GetProject returns information about a project.
 */
func GetProject (id uint64) (structure ProjectResponse, err error) {
	response, body, err := Request {
		Path:     "/projects/" + strconv.FormatUint(id, 10),
		Hostname: "api.scratch.mit.edu",
	}.Send()
	if err != nil { return }
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf (
			"cannot get project info (%i)",
			response.StatusCode)
		return
	}

	err = json.Unmarshal(body, &structure)
	if err != nil { return }
	return
}
