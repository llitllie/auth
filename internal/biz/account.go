package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Account struct {
	ID     int
	Email  string
	Secret string
	Status bool
	Type   string
}

type AccountRepo interface {
	CreateAccount(ctx context.Context, c *Account) (*Account, error)
	UpdateAccount(ctx context.Context, c *Account) (*Account, error)
	GetAccount(ctx context.Context, id int) (*Account, error)
	RemoveAccount(ctx context.Context, id int) (*Account, error)
	ListAccount(ctx context.Context) ([]*Account, error)
}

type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/Account"))}
}

func (uc *AccountUseCase) Create(ctx context.Context, u *Account) (*Account, error) {
	return uc.repo.CreateAccount(ctx, u)
}

func (uc *AccountUseCase) Get(ctx context.Context, id int) (*Account, error) {
	return uc.repo.GetAccount(ctx, id)
}

func (uc *AccountUseCase) Update(ctx context.Context, u *Account) (*Account, error) {
	return uc.repo.UpdateAccount(ctx, u)
}

func (uc *AccountUseCase) Remove(ctx context.Context, id int) (*Account, error) {
	return uc.repo.RemoveAccount(ctx, id)
}

func (uc *AccountUseCase) List(ctx context.Context) ([]*Account, error) {
	return uc.repo.ListAccount(ctx)
}
