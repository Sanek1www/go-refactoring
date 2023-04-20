package container

import (
	"github.com/golobby/container/v3"
	"log"
	"refactoring/cmd/config"
	userUC "refactoring/internal/app/usecase/user"
	"refactoring/internal/pkg/repository/user"
)

func InitContainer(c *config.Config) {
	container.Singleton(func() user.ContractRepository {
		repo, err := user.MakeJsonRepository(c.JsonStorageLink)
		if err != nil {
			log.Fatal(err)
		}

		return repo
	})

	container.Singleton(func(repo user.ContractRepository) *userUC.Usecase {
		return userUC.New(repo)
	})
}
