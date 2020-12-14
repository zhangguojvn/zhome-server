package Device

import "zhome-server/DataBase/UIConfig"

type Device struct {
	DUID string
	Name string
	UI   UIConfig.UIConfig
	Info string
}
