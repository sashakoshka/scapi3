package scapi3

import "fmt"
import "strconv"
import "net/http"
import "encoding/json"

/* RestRequest performs a generic request to the scratch rest API.
 */
func RestRequest [T any](
	structure *T,
	path   string,
	limit  int,
	offset int,
) (
	err error,
) {
	return RestRequestWithQueryString (
		structure, path,
		limit, offset, "")
}

/* RestRequestWithQueryString performs a generic request to the scratch rest
 * API with additional query string parameters.
 */
func RestRequestWithQueryString [T any](
	structure   *T,
	path        string,
	limit       int,
	offset      int,
	queryString string,
) (
	err error,
) {
	path += "?"
	if limit  > 0 { path += fmt.Sprintf("limit=%d&",  limit ) }
	if offset > 0 { path += fmt.Sprintf("offset=%d&", offset) }
	if queryString != "" {
		path += queryString
	}

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
	err = RestRequest(&structure, "/health", 0, 0)
	return
}

/* GetNews returns recent news articles from the scratch website.
 */
func GetNews (limit, offset int) (structure NewsResponse, err error) {
	err = RestRequest(&structure, "/news", limit, offset)
	return
}

/* GetProjectsCountAll returns the amount of projects that have been uploaded to
 * the site.
 */
func GetProjectsCountAll () (count uint64, err error) {
	structure := CountResponse { }
	err = RestRequest(&structure, "/projects/count/all", 0, 0)
	count = structure.Count
	return
}

/* GetProject returns information about a project.
 */
func GetProject (id uint64) (structure ProjectResponse, err error) {
	err = RestRequest (
		&structure, "/projects/" + strconv.FormatUint(id, 10),
		0, 0)
	return
}

/* GetProjectRemixes returns the remixes of a project.
 */
func GetProjectRemixes (
	id     uint64,
	limit  int,
	offset int,
) (
	structure []ProjectResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/projects/" + strconv.FormatUint(id, 10) + "/remixes",
		limit, offset)
	return
}

/* GetStudio returns information about a studio.
 */
func GetStudio (id uint64) (structure StudioResponse, err error) {
	err = RestRequest (
		&structure, "/studios/" + strconv.FormatUint(id, 10),
		0, 0)
	return
}

/* GetStudioProjects returns a list of all projects in a studio.
 */
func GetStudioProjects (
	id     uint64,
	limit  int,
	offset int,
) (
	structure StudioProjectsResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/projects",
		limit, offset)
	return
}

/* GetStudioManagers returns a list of all managers of a studio.
 */
func GetStudioManagers (
	id     uint64,
	limit  int,
	offset int,
) (
	structure []UserResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/managers",
		limit, offset)
	return
}

/* GetStudioCurators returns a list of all curators of a studio.
 */
func GetStudioCurators (
	id     uint64,
	limit  int,
	offset int,
) (
	structure []UserResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/curators",
		limit, offset)
	return
}

/* GetStudioActivity returns a list of all recent activityin s studio. If since
 * is blank, a date limit is not sent.
 */
func GetStudioActivity (
	id uint64,
	// TODO: make this a golang time
	since string,
	limit int,
) (
	structure StudioActivityResponse,
	err error,
) {
	url := "/studios/" + strconv.FormatUint(id, 10) + "/activity"

	if since != "" {
		since = "dateLimit=" + since
	}
	
	err = RestRequestWithQueryString(&structure, url, limit, 0, since)
	return
}

/* GetStudioComments returns a list of all comments on a studio.
 */
func GetStudioComments (
	id     uint64,
	limit  int,
	offset int,
) (
	structure []CommentResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) + "/comments",
		limit, offset)
	return
}

/* GetStudioComment returns a comment on a studio.
 */
func GetStudioComment (
	id        uint64,
	commentID uint64,
) (
	structure CommentResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) +
		"/comments/" + strconv.FormatUint(commentID, 10),
		0, 0)
	return
}

/* GetStudioCommentReplies returns the replies for a comment on a studio.
 */
func GetStudioCommentReplies (
	id        uint64,
	commentID uint64,
	limit     int,
	offset    int,
) (
	structure []CommentResponse,
	err error,
) {
	err = RestRequest (
		&structure,
		"/studios/" + strconv.FormatUint(id, 10) +
		"/comments/" + strconv.FormatUint(commentID, 10) + "/replies",
		limit, offset)
	return
}

/* GetFeatured returns information about front paged projects.
 */
func GetFeatured () (structure FeaturedResponse, err error) {
	err = RestRequest(&structure, "/proxy/featured", 0, 0)
	return
}

/* GetUser returns information about a user.
 */
func GetUser (name string) (structure UserResponse, err error) {
	err = RestRequest(&structure, "/users/" + name, 0, 0)
	return
}

/* GetUserFavorites returns all projects favorited by a user.
 */
func GetUserFavorites (
	name   string,
	limit  int,
	offset int,
) (
	structure []ProjectResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/favorites",
		limit, offset)
	return
}

/* GetUserFollowers returns all followers of a user.
 */
func GetUserFollowers (
	name   string,
	limit  int,
	offset int,
) (
	structure []UserResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/followers",
		limit, offset)
	return
}

/* GetUserFollowing returns all users a user is following.
 */
func GetUserFollowing (
	name   string,
	limit  int,
	offset int,
) (
	structure []UserResponse,
	err error,
) {
	err = RestRequest (&structure, "/users/" + name + "/followers",
		limit, offset)
	return
}

/* GetUserMessageCount returns the amount of messages a user has.
 */
func GetUserMessageCount (name string) (count uint64, err error) {
	structure := CountResponse { }
	err = RestRequest (
		&structure, "/users/" + name + "/messages/count",
		0, 0)
	count = structure.Count
	return
}

/* GetUserProjects returns all projects made by a user.
 */
func GetUserProjects (
	name   string,
	limit  int,
	offset int,
) (
	structure []ProjectResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/projects",
		limit, offset)
	return
}

/* GetUserProject returns a specific project made by a user.
 */
func GetUserProject (
	name string,
	id   uint64,
) (
	structure ProjectResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/projects/" +
		strconv.FormatUint(id, 10),
		0, 0)
	return
}

/* GetUserProject returns a specific project made by a user.
 */
func GetUserProjectStudios (
	name   string,
	id     uint64,
	limit  int,
	offset int,
) (
	structure []StudioResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/projects/" +
		strconv.FormatUint(id, 10) + "/studios",
		limit, offset)
	return
}

/* GetUserProjectComments returns a list of all comments on a user's project.
 */
func GetUserProjectComments (
	name   string,
	id     uint64,
	limit  int,
	offset int,
) (
	structure []CommentResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/projects/" +
		strconv.FormatUint(id, 10) + "/comments",
		limit, offset)
	return
}

/* GetUserProjectComment returns information about a specific comment on a
 * user's project.
 */
func GetUserProjectComment (
	name      string,
	id        uint64,
	commentid uint64,
	limit     int,
	offset    int,
) (
	structure CommentResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/projects/" +
		strconv.FormatUint(id, 10) + "/comments/" +
		strconv.FormatUint(commentid, 10),
		limit, offset)
	return
}

/* GetUserProjectCommentReplies returns replies to a specific comment on a
 * user's project.
 */
func GetUserProjectCommentReplies (
	name      string,
	id        uint64,
	commentid uint64,
	limit     int,
	offset    int,
) (
	structure []CommentResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/projects/" +
		strconv.FormatUint(id, 10) + "/comments/" +
		strconv.FormatUint(commentid, 10) + "/replies",
		limit, offset)
	return
}

/* GetUserStudiosCurate returns the studios that the user is a curator of.
 */
func GetUserStudiosCurate (
	name   string,
	limit  int,
	offset int,
) (
	structure []StudioResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/users/" + name + "/studios/curate",
		limit, offset)
	return
}

/* GetAccountsCheckUsername checks whether a new account with the specified
 * username can be created (that is, it isn't taken).
 */
func GetAccountsCheckUsername (
	name string,
) (
	structure AccountsCheckUsernameResponse,
	err error,
) {
	err = RestRequest (
		&structure, "/accounts/checkusername/" + name,
		0, 0)
	return
}
