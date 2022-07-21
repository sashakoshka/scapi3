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

/* HealthResponse contains data about the health of the Scratch website.
 */
type HealthResponse struct {
	Version string      `json:"version"`
	Uptime  int         `json:"uptime"`
	Load    []float64   `json:"load"`
	SQL struct {
		Main                HealthResponseSQL `json:"main"`
		ProjectComments     HealthResponseSQL `json:"project_comments"`
		GalleryComments     HealthResponseSQL `json:"gallery_comments"`
		UserprofileComments HealthResponseSQL `json:"userprofile_comments"`
		Timestamp           uint64            `json:"timestamp"`
	} `json:"sql"`
	Cache struct {
		Connected bool `json:"connected"`
		Ready     bool `json:"ready"`
	} `json:"cache"`
}

/* HealthResponseSQL represents an database in the site SQL health data.
 */
type HealthResponseSQL struct {
	Primary HealthResponseSQLStatistic `json:"primary"`
	Replica HealthResponseSQLStatistic `json:"replica"`
}

/* HealthResponseSQL represents a database's health data in the site SQL
 * health data.
 */
type HealthResponseSQLStatistic struct {
	SSL             bool `json:"ssl"`
	Destroyed       bool `json:"destroyed"`
	Min             int  `json:"min"`
	Max             int  `json:"max"`
	NumUsed         int  `json:"numUsed"`
	NumFree         int  `json:"numFree"`
	PendingAcquires int  `json:"pendingAcquires"`
	PendingCreates  int  `json:"pendingCreates"`
}

/* NewsResponse represents a list of recent news from the Scratch website.
 */
type NewsResponse []NewsResponseItem

/* NewsResponseItem represents a single news article
 */
type NewsResponseItem struct {
	Id       uint64 `json:"id"`
	Stamp    string `json:"stamp"`
	Headline string `json:"headline"`
	URL      string `json:"url"`
	Image    string `json:"image"`
	Copy     string `json:"copy"`
}
