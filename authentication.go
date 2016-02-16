package libyggdrasil

import (
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

// Making client a package-level var for efficient reuse.
var client *http.Client

// Agent is needed as a type for the Request, though it's technically optional.
// Right now we only support MC logins so we're hardcoding this.
const agentName string = "Minecraft"
const agentVersion int = 1

type yggdrasilAgent struct {
	name    string
	version int
}

// Request struct to contain JSON payload values.
type yggdrasilAuthenticationPayload struct {
	username    string
	password    string
	clientToken uuid.UUID
	agent       yggdrasilAgent
}

func init() {
	client = &http.Client{Timeout: time.Duration(10) * time.Second}
}

func Login(username string, password string, clientToken interface{}) (*YggdrasilAuthenticationResponse, error) {
	request := yggdrasilAuthenticationPayload{
		username: username,
		password: password,
		agent:    yggdrasilAgent{name: agentName, version: agentVersion},
	}

	switch clientToken := clientToken.(type) {
	case uuid.UUID:
		request.clientToken = clientToken
	case string:
		parsedToken, tokenErr := uuid.FromString(clientToken)
		if tokenErr != nil {
			panic(tokenErr.Error())
		}
		request.clientToken = parsedToken
	}

	return request.login()
}

func (req yggdrasilAuthenticationPayload) login() (resp *YggdrasilAuthenticationResponse, err *YggdrasilAuthenticationError) {
	return
}
