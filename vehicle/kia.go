package vehicle

import (
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/vehicle/bluelink"
)

// Kia is an api.Vehicle implementation
type Kia struct {
	*embed
	*bluelink.API
}

func init() {
	registry.Add("kia", NewKiaFromConfig)
}

// NewKiaFromConfig creates a new Vehicle
func NewKiaFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	cc := struct {
		Title          string
		Capacity       int64
		User, Password string
		PIN            string
		Cache          time.Duration
	}{}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	settings := bluelink.Config{
		URI:               "https://prd.eu-ccapi.kia.com:8080",
		TokenAuth:         "ZmRjODVjMDAtMGEyZi00YzY0LWJjYjQtMmNmYjE1MDA3MzBhOnNlY3JldA==",
		CCSPServiceID:     "fdc85c00-0a2f-4c64-bcb4-2cfb1500730a",
		CCSPApplicationID: "693a33fa-c117-43f2-ae3b-61a02d24f417",
		DeviceID:          "/api/v1/spa/notifications/register",
		Lang:              "/api/v1/user/language",
		Login:             "/api/v1/user/signin",
		AccessToken:       "/api/v1/user/oauth2/token",
		Vehicles:          "/api/v1/spa/vehicles",
		SendPIN:           "/api/v1/user/pin",
		GetStatus:         "/api/v2/spa/vehicles/",
	}

	log := util.NewLogger("kia")
	api, err := bluelink.New(log, cc.User, cc.Password, cc.PIN, cc.Cache, settings)
	if err != nil {
		return nil, err
	}

	v := &Kia{
		embed: &embed{cc.Title, cc.Capacity},
		API:   api,
	}

	return v, nil
}