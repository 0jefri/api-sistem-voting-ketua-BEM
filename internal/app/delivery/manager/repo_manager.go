package manager

import (
	"github.com/api-voting/internal/app/repository"
)

type RepoManager interface {
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(m.infraManager.Conn())
}
