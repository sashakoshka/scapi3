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

/* ProjectsCountAllResponse contains the amount of projects uploaded to the
 * site.
 */
type ProjectsCountAllResponse struct {
	Count uint64 `json:"count"`
}

/* ProjectResponse holds information about a project.
 */
type ProjectResponse struct {
	ID              uint64       `json:"id"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	Instructions    string       `json:"instructions"`
	Visibility      string       `json:"visibility"`
	Public          bool         `json:"public"`
	CommentsAllowed bool         `json:"comments_allowed"`
	IsPublished     bool         `json:"is_published"`
	Author          UserResponse `json:"author"`
	Image           string       `json:"image"`
	Images struct {
		Size218px string `json:"282x218"`
		Size163px string `json:"216x163"`
		Size200px string `json:"200x200"`
		Size108px string `json:"144x108"`
		Size102px string `json:"135x102"`
		Size80px  string `json:"100x80"`
	} `json:"images"`
	History struct {
		Created  string `json:"created"`
		Modified string `json:"modified"`
		Shared   string `json:"shared"`
	} `json:"history"`
	Stats struct {
		Views     int `json:"views"`
		Loves     int `json:"loves"`
		Favorites int `json:"favorites"`
		Remixes   int `json:"remixes"`
	} `json:"stats"`
	Remix struct {
		Parent uint64 `json:"parent"`
		Root   uint64 `json:"root"`
	} `json:"remix"`
	ProjectToken string `json:"project_token"`
}

/* UserResponse holds information about a user.
 */
type UserResponse struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	ScratchTeam bool   `json:"scratchteam"`
	History struct {
		Joined string `json:"joined"`
	} `json:"history"`
	Profile struct {
		ID uint64 `json:"id"`
		Images struct {
			Size90px string `json:"90x90"`
			Size60px string `json:"60x60"`
			Size55px string `json:"55x55"`
			Size50px string `json:"50x50"`
			Size32px string `json:"32x32"`
		}
		Status  string `json:"status"`
		Bio     string `json:"bio"`
		Country string `json:"country"`
	} `json:"profile"`
}
