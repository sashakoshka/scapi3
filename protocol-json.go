package scapi3

import "encoding/json"

/* ProtocolRequest represents a structure that can be marshalled into a byte
 * slice.
 */
type ProtocolRequest interface {
	Marshal () (data []byte)
}

/* ProtocolLoginRequest represents a username and password combination that can
 * be encoded into the body of a login request.
 */
type ProtocolLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* Marshal converts the login request body into a JSON encoded byte slice.
 */
func (structure ProtocolLoginRequest) Marshal () (data []byte) {
	data, _ = json.Marshal(structure)
	return
}

/* ProtocolCommentRequest represents a comment request.
 */
type ProtocolCommentRequest struct {
	Content     string
	ParentID    uint64
	CommenteeID uint64
}

/* Marshal converts the comment request body into a JSON encoded byte slice.
 */
func (structure ProtocolCommentRequest) Marshal () (data []byte) {
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

/* ProtocolLoginResponse represents a login response.
 */
type ProtocolLoginResponse struct {
	Username string   `json:"username"`
	Token    string   `json:"token"`
	NumTries int      `json:"num_tries"`
	Success  int      `json:"success"`
	Msg      string   `json:"msg"`
	Messages []string `json:"messages"`
	ID       int      `json:"id"`
}

/* UnmarshalLoginResponse takes in a JSON encoded byte slice and returns
 * unmarshaled login response data.
 */
func UnmarshalLoginResponse (
	data []byte,
) (
	structure ProtocolLoginResponse,
	err       error,
) {
	array := []ProtocolLoginResponse { }
	err = json.Unmarshal(data, &array)
	if err != nil { return }
	structure = array[0]
	return
}

/* ProtocolCommentResponse represents a response to a comment request.
 */
type ProtocolCommentResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
