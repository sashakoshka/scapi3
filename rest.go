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
	err = RestRequest (
		&structure,
		"/projects/" + strconv.FormatUint(id, 10) + "/remixes")
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
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/projects")
	return
}

/* GetStudioManagers returns a list of all managers of a studio.
 */
func GetStudioManagers (id uint64) (structure StudioManagersResponse, err error) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/managers")
	return
}

/* GetStudioCurators returns a list of all curators of a studio.
 */
func GetStudioCurators (id uint64) (structure StudioCuratorsResponse, err error) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/curators")
	return
}

/* GetStudioActivity returns a list of all curators of a studio. If since is
 * blank, a date limit is not sent.
 */
func GetStudioActivity (
	id uint64,
	// TODO: make this a golang time
	since string,
) (
	structure StudioActivityResponse,
	err error,
) {
	url := "/studios/" + strconv.FormatUint(id, 10) + "/activity"

	if since != "" {
		url += "?dateLimit=" + since
	}
	
	err = RestRequest(&structure, url)
	return
}

/* GetStudioComments returns a list of all comments on a studio.
 */
func GetStudioComments (id uint64) (structure CommentsResponse, err error) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/comments")
	return
}

/* GetStudioComment returns a comment on a studio.
 */
func GetStudioComment (id, commentID uint64) (structure CommentResponse, err error) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) +
		"/comments/" + strconv.FormatUint(commentID, 10))
	return
}

/* GetStudioCommentReplies returns the replies for a comment on a studio.
 */
func GetStudioCommentReplies (id, commentID uint64) (structure CommentsResponse, err error) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) +
		"/comments/" + strconv.FormatUint(commentID, 10) + "/replies")
	return
}
