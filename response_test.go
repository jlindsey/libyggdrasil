package libyggdrasil

import (
	"github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestResponse(t *testing.T) {
	resp := YggdrasilAuthenticationResponse{
		ClientToken:       uuid.NewV4(),
		AccessToken:       "abc123",
		AvailableProfiles: make([]yggdrasilProfile, 0),
	}

	Convey("With no available profiles", t, func() {
		Convey("Should have no Minecraft license", func() {
			So(resp.HasMinecraftLicense(), ShouldBeFalse)
		})
	})

	Convey("With available profiles", t, func() {
		resp.AvailableProfiles = append(resp.AvailableProfiles, yggdrasilProfile{})
		Convey("Should have a Minecraft license", func() {
			So(resp.HasMinecraftLicense(), ShouldBeTrue)
		})
	})
}
