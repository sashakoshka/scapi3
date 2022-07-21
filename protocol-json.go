package scapi3

import "encoding/json"

/* LoginRequest represents a username and password combination that can
 * be encoded into the body of a login request.
 */
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* Marshal converts the login request body into a JSON encoded byte slice.
 */
func (structure LoginRequest) Marshal () (data []byte) {
	data, _ = json.Marshal(structure)
	return
}

/* CommentRequest represents a comment request.
 */
type CommentRequest struct {
	Content     string
	ParentID    uint64
	CommenteeID uint64
}

/* Marshal converts the comment request body into a JSON encoded byte slice.
 */
func (structure CommentRequest) Marshal () (data []byte) {
	realStructure := map[string] any {
		"content":      structure.Content,
		"parent_id":    structure.ParentID,
		"commentee_id": structure.CommenteeID,
	}

	if structure.ParentID == 0 {
		realStructure["parent_id"] = ""
	}

	if structure.CommenteeID == 0 {
		realStructure["commentee_id"] = ""
	}

	data, _ = json.Marshal(realStructure)
	return
}

/* LoginStatusResponse represents a login response.
 */
type LoginStatusResponse struct {
	Username string   `json:"username"`
	Token    string   `json:"token"`
	NumTries int      `json:"num_tries"`
	Success  int      `json:"success"`
	Msg      string   `json:"msg"`
	Messages []string `json:"messages"`
	ID       int      `json:"id"`
}

/* UnmarshalLoginStatusResponse takes in a JSON encoded byte slice and returns
 * unmarshaled login response data.
 */
func UnmarshalLoginStatusResponse (
	data []byte,
) (
	structure LoginStatusResponse,
	err       error,
) {
	array := []LoginStatusResponse { }
	err = json.Unmarshal(data, &array)
	if err != nil { return }
	structure = array[0]
	return
}

/* CommentStatusResponse represents a response to a comment request.
 */
type CommentStatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

/* AccountsCheckUsernameResponse holds information about the availibility of a
 * username.
 */
type AccountsCheckUsernameResponse struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
}
