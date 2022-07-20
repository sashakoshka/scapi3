package scapi3

import "fmt"
// import "net/url"
import "net/http"
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
		
		Body: ProtocolLoginRequest {
			Username: session.username,
			Password: session.password,
		},
		
		Headers: map[string] string {
			"X-Requested-With": "XMLHttpRequest",
			"Content-Type": "application/json",
		},

		Client: &session.client,
	}.Send()
	
	loginData, err := UnmarshalProtocolLoginResponse(body)
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
