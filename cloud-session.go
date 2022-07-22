package scapi3

import "fmt"
import "strings"
import "strconv"
import "net/url"
import "net/http"
import "github.com/gorilla/websocket"

/* CloudMethod represents the method of a cloud session message. It can be
 * either handshake, or set. Handshake initialtes a cloud connection, and set
 * sets a variable.
 */
type CloudMethod string

const (
	CloudMethodHandshake CloudMethod = "handshake"
	CloudMethodSet       CloudMethod = "set"
)

/* CloudMessage represents a single message sent over a cloud session.
 */
type CloudMessage struct {
	Method   CloudMethod `json:"method"`
	Name     string      `json:"name"`
	Value    string      `json:"value"`
	Variable *CloudVariable
}

/* CloudVariable represents the value of a cloud variable. It has a string and
 * float value that are automatically kept in sync with one another.
 */
type CloudVariable struct {
	// TODO: add reference to parent session so mutator methods can send
	// set methods to server when variable is set.
	name        string
	stringValue string
	floatValue  float64
}

/* CloudSession represents a cloud variable session.
 */
type CloudSession struct {
	userSession *UserSession
	projectID   uint64
	connection  *websocket.Conn
	variables   map[string] *CloudVariable
}

/* CreateCloudSession creates a new cloud session for the specified user in the
 * specified project.
 */
func CreateCloudSession (
	userSession *UserSession,
	projectID   uint64,
) (
	session *CloudSession,
	err error,
) {
	session = &CloudSession {
		userSession: userSession,
		projectID:   projectID,
	}

	header := http.Header { }
	header.Add("Cookie", userSession.sessionID)
	header.Set("User-Agent", "")
	header.Set("Origin", "https://scratch.mit.edu")

	cloudUrl := url.URL {
		Scheme: "ws",
		Host:   "clouddata.scratch.mit.edu",
		Path:   "/",
	}

	session.connection, _, err = websocket.DefaultDialer.Dial (
		cloudUrl.String(),
		header)
	if err != nil { return}

	err = session.sendHandshake()
	return
}

/* ReadMessage reads a single message from the Scratch site, and returns it.
 */
func (session *CloudSession) ReadMessage () (message CloudMessage, err error) {
	err = session.connection.ReadJSON(message)
	if err != nil { return }

	switch message.Method {
	case CloudMethodSet:
		variable := session.GetVariable(message.Name)
		if variable == nil {
			variable = &CloudVariable { }
			variable.SetString(message.Value)
			session.variables[message.Name] = variable
		}

		message.Variable = variable
		break
	}
	
	return
}

/* GetVariable returns the cloud variable object of name string from the cloud
 * session.
 */
func (session *CloudSession) GetVariable (name string) (variable *CloudVariable) {
	if !strings.HasPrefix(name, "☁ ") {
		name = addCloudSymbol(name)
	}
	return session.variables[name]
}

/* sendHandshake sends login information to the Scratch servers, officially
 * initiating the cloud session.
 */
func (session *CloudSession) sendHandshake () (err error) {
	return session.send(CloudMethodHandshake, nil)
}

/* send sends a single message over the cloud session.
 */
func (session *CloudSession) send (
	method CloudMethod,
	data map[string] any,
) (
	err error,
) {
	if data == nil {
		data = make(map[string] any)
	}

	data["method"]     = method
	data["user"]       = session.userSession.username
	data["project_id"] = session.projectID

	fmt.Println(data)
	
	return session.connection.WriteJSON(data)
}

/* Close cleanly ends the cloud session.
 */
func (session *CloudSession) Close () (err error) {
	err = session.connection.WriteMessage (
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	session.connection.Close()
	return
}

/* SetString sets the variable value using a string. The float value is parsed
 * from this string.
 */
func (variable *CloudVariable) SetString (value string) {
	if value == variable.stringValue { return }
	variable.stringValue = value
	variable.floatValue, _ = strconv.ParseFloat(value, 64)
}

/* String returns the variable value as a string.
 */
func (variable *CloudVariable) String () (value string) {
	return variable.stringValue
}

/* SetFloat sets the variable value using a float. The string value is formatted
 * from this float.
 */
func (variable *CloudVariable) SetFloat (value float64) {
	if value == variable.floatValue { return }
	variable.floatValue  = value
	variable.stringValue = strconv.FormatFloat(value, 'f', 5, 64)
	
}

/* Float returns the variable value as a float.
 */
func (variable *CloudVariable) Float () (value float64) {
	return variable.floatValue
}

/* addCloudSymbol prepends the cloud unicode symbol to a variable name.
 */
func addCloudSymbol (variableName string) (cloudVariableName string) {
	return "☁ " + variableName
}
