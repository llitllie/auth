package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/llitllie/auth/api/auth/v1"
	"github.com/llitllie/auth/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthService)

type AuthService struct {
	pb.UnimplementedAuthServer

	account *biz.AccountUseCase
	log     *log.Helper
}

func NewAuthService(account *biz.AccountUseCase, logger log.Logger) *AuthService {
	return &AuthService{
		account: account,
		log:     log.NewHelper(log.With(logger, "module", "service/auth")),
	}
}
