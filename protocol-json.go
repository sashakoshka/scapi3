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
