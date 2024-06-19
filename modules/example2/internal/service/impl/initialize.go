package impl

import (
	"crab/modules/example2/internal/service"
)

func Initialize() {
	service.Api = newImplApi()
}
