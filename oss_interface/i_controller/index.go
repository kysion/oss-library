package i_controller

import "github.com/kysion/oss-library/oss_interface"

type iModule interface {
	GetModules() oss_interface.IModules
}
