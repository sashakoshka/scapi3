package scapi3

import "fmt"
import "strconv"
import "net/http"
import "encoding/json"
import "net/http/cookiejar"

/* UserSession represents a login session. It is used to access user-specific
 * parts of the API like cloud variables.
 */
type UserSession struct {
	loaded		bool
	valid		bool
	username	string
	password	string
	id		int
	token		string
	client		http.Client
}

/* CreateUserSession creates and loads a new user session with the specified
 * username and password.
 */
func CreateUserSession (
	username string,
	password string,
) (
	session *UserSession,
	err     error,
) {
	session = &UserSession {
		username: username,
		password: password,
	}
	
	session.client.Jar, err = cookiejar.New(nil)
	if err != nil {
		err = fmt.Errorf("could not create cookie jar: %v", err)
		return
	}
	
	err = session.Login()
	return
}

/* Login connects to the scratch servers and starts the session.
 */
func (session *UserSession) Login () (err error) {
	if session.loaded { return }

	response, body, err := Request {
		Path:	"/accounts/login/",
		Method:	MethodPost,
		
		Body: LoginRequest {
			Username: session.username,
			Password: session.password,
		},
		
		Headers: map[string] string {
			"X-Requested-With": "XMLHttpRequest",
			"Content-Type": "application/json",
		},

		Client: &session.client,
	}.Send()
	
	loginData, err := UnmarshalLoginStatusResponse(body)
	if err != nil {
		return fmt.Errorf (
			"cannot parse server response (%s): %v",
			response.Status, err)
	}
	
	if loginData.Success != 1 {
		return fmt.Errorf (
			"cannot log in (%s): %s",
			response.Status, loginData.Msg)
	}

	session.id     = loginData.ID
	session.loaded = true
	session.valid  = false
	session.token  = loginData.Token

	// urlObj, _ := url.Parse("https://scratch.mit.edu")
	// fmt.Println(session.client.Jar.Cookies(urlObj))
	return
}

/* Verify checks whether a session is valid. It returns whether or not it is
 * valid, and sets the session's valid flag accordingly. If the session is not
 * loaded, this function does nothing.
 */
func (session *UserSession) Verify () (valid bool) {
	if !session.loaded { return }
	
	// TODO: there is a high chance this doesn't work. Find out if it does
	// work, and if it doesn't, create a better solution.
	response, _, err := Request {
		Client: &session.client,
	}.Send()
	valid = err == nil && response.StatusCode == http.StatusOK
	session.valid = valid
	return
}

/* CommentOnProject writes a comment on a project. If parent is nonzero, this
 * comment will be in reply to the comment with that ID. If tagging is not
 * blank, this comment will "@" the user with that ID.
 */
func (session *UserSession) CommentOnProject (
	project uint64,
	parent  uint64,
	tagging string,
	content string,
) (
	err error,
) {
	if !session.loaded { return }
	return session.comment (
		strconv.FormatUint(project, 10),
		"project", parent, tagging, content)
}

/* CommentOnUser writes a comment on a project. If parent is nonzero, this
 * comment will be in reply to the comment with that ID. If tagging is not
 * blank, this comment will "@" the user with that ID.
 */
func (session *UserSession) CommentOnUser (
	user    string,
	parent  uint64,
	tagging string,
	content string,
) (
	err error,
) {
	if !session.loaded { return }
	return session.comment(user, "user", parent, tagging, content)
}

/* CommentOnStudio writes a comment on a project. If parent is nonzero, this
 * comment will be in reply to the comment with that ID. If tagging is not
 * blank, this comment will "@" the user with that ID.
 */
func (session *UserSession) CommentOnStudio (
	studio  string,
	parent  uint64,
	tagging string,
	content string,
) (
	err error,
) {
	if !session.loaded { return }
	return session.comment(studio, "gallery", parent, tagging, content)
}

func (session *UserSession) comment (
	id      string,
	where   string,
	parent  uint64,
	tagging string,
	content string,
) (
	err error,
) {	
	response, body, err := Request {
		Hostname: "api.scratch.mit.edu",
		Path:     "/proxy/comments/" + where + "/" + id,
		Method:   MethodPost,

		Body: CommentRequest {
			Content:   content,
			ParentID:  parent,
			Commentee: tagging,
		},
		
		Headers: map[string] string {
			"X-Requested-With": "XMLHttpRequest",
			"Content-Type": "application/json",
			"Origin": "https://scratch.mit.edu",
		},

		Client: &session.client,
	}.Send()

	if err != nil {
		return fmt.Errorf (
			"cannot comment (%s): %v",
			response.Status, err)
	}

	commentData := CommentStatusResponse { }
	err = json.Unmarshal(body, &commentData)
	if err != nil {
		return fmt.Errorf (
			"cannot parse server response (%s): %v",
			response.Status, err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf (
			"cannot comment (%s): %s: %s",
			response.Status, commentData.Code, commentData.Message)
	}
	return
}
