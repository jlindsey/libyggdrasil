package libyggdrasil

import (
	"github.com/satori/go.uuid"
)

type yggdrasilProfile struct {
	Id     string
	Name   string
	Legacy bool
}

type YggdrasilAuthenticationResponse struct {
	ClientToken       uuid.UUID
	AccessToken       string
	AvailableProfiles []yggdrasilProfile
	SelectedProfile   yggdrasilProfile
}

func (resp *YggdrasilAuthenticationResponse) HasMinecraftLicense() bool {
	return len(resp.AvailableProfiles) > 0
}
