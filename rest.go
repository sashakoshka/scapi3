package scapi3

import "fmt"
import "strconv"
import "net/http"
import "encoding/json"

/* RestRequest performs a generic request to the scratch rest API.
 */
func RestRequest [T any](structure *T, path string) (err error) {
	response, body, err := Request {
		Path:     path,
		Hostname: "api.scratch.mit.edu",
	}.Send()
	if err != nil { return }
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf (
			"cannot get %s (%i)",
			path,
			response.StatusCode)
		return
	}
	
	err = json.Unmarshal(body, &structure)
	if err != nil { return }
	return
}

/* GetHealth returns information relating to the health of the scratch website.
 */
func GetHealth () (structure HealthResponse, err error) {
	err = RestRequest(&structure, "/health")
	return
}

/* GetNews returns recent news articles from the scratch website.
 */
func GetNews () (structure NewsResponse, err error) {
	err = RestRequest(&structure, "/news")
	return
}

/* GetProjectsCountAll returns the amount of projects that have been uploaded to
 * the site.
 */
func GetProjectsCountAll () (count uint64, err error) {
	structure := ProjectsCountAllResponse { }
	err = RestRequest(&structure, "/projects/count/all")
	count = structure.Count
	return
}

/* GetProject returns information about a project.
 */
func GetProject (id uint64) (structure ProjectResponse, err error) {
	err = RestRequest(&structure, "/projects/" + strconv.FormatUint(id, 10))
	return
}

/* GetProjectRemixes returns the remixes of a project.
 */
func GetProjectRemixes (id uint64) (structure RemixesResponse, err error) {
	err = RestRequest(&structure, "/projects/" + strconv.FormatUint(id, 10) + "/remixes")
	return
}

/* GetStudio returns information about a studio.
 */
func GetStudio (id uint64) (structure StudioResponse, err error) {
	err = RestRequest(&structure, "/studios/" + strconv.FormatUint(id, 10))
	return
}

/* GetStudioProjects returns a list of all projects in a studio.
 */
func GetStudioProjects (id uint64) (structure StudioProjectsResponse, err error) {
	err = RestRequest(&structure, "/studios/" + strconv.FormatUint(id, 10) + "/projects")
	return
}
