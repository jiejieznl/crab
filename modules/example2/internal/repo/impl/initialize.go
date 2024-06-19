package impl

import "crab/modules/example2/internal/repo"

func Initialize() {
	repo.Api = newImplApi()
}
