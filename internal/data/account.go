package data

import (
	"context"

	"github.com/llitllie/auth/internal/biz"
	"github.com/llitllie/auth/internal/data/ent/account"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/account")),
	}
}

func (r *accountRepo) CreateAccount(ctx context.Context, b *biz.Account) (*biz.Account, error) {
	acc, err := r.data.db.Account.
		Create().
		SetEmail(b.Email).
		SetSecret(b.Secret).
		SetType(account.Type(b.Type)).
		SetStatus(b.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		ID:     acc.ID,
		Email:  acc.Email,
		Secret: "",
		Status: acc.Status,
		Type:   acc.Type.String(),
	}, nil
}

func (r *accountRepo) UpdateAccount(ctx context.Context, b *biz.Account) (*biz.Account, error) {
	acc, err := r.data.db.Account.
		UpdateOneID(b.ID).
		SetEmail(b.Email).
		SetType(account.Type(b.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		ID:     acc.ID,
		Email:  acc.Email,
		Secret: "",
		Status: acc.Status,
		Type:   acc.Type.String(),
	}, nil
}

func (r *accountRepo) GetAccount(ctx context.Context, id int) (*biz.Account, error) {
	acc, err := r.data.db.Account.Query().Where(account.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		ID:     acc.ID,
		Email:  acc.Email,
		Secret: "",
		Status: acc.Status,
		Type:   acc.Type.String(),
	}, nil
}

func (r *accountRepo) RemoveAccount(ctx context.Context, id int) (*biz.Account, error) {
	err := r.data.db.Account.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *accountRepo) ListAccount(ctx context.Context) ([]*biz.Account, error) {
	accs, err := r.data.db.Account.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	bizs := make([]*biz.Account, 0)
	for _, acc := range accs {
		bizs = append(bizs, &biz.Account{
			ID:     acc.ID,
			Email:  acc.Email,
			Secret: "",
			Status: acc.Status,
			Type:   acc.Type.String(),
		})
	}
	return bizs, nil
}
