package service

import (
	"context"

	pb "github.com/llitllie/auth/api/auth/v1"
	"github.com/llitllie/auth/internal/biz"
)

func (s *AuthService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	a := &biz.Account{
		Email:  req.Email,
		Secret: req.Secret,
		Type:   req.Type,
		Status: req.Status,
	}
	x, err := s.account.Create(ctx, a)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountReply{
		Account: &pb.Account{
			Id:     int64(x.ID),
			Email:  x.Email,
			Type:   x.Type,
			Status: x.Status,
		},
	}, nil
}
func (s *AuthService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountReply, error) {
	a, err := s.account.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	b := &biz.Account{
		ID:    a.ID,
		Email: a.Email,
		Type:  a.Type,
	}
	if req.Email != "" {
		b.Email = req.Email
	}
	if req.Type != "" {
		b.Type = req.Type
	}
	x, err := s.account.Update(ctx, b)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountReply{
		Account: &pb.Account{
			Id:     int64(x.ID),
			Email:  x.Email,
			Type:   x.Type,
			Status: x.Status,
		},
	}, nil
}
func (s *AuthService) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountReply, error) {
	_, err := s.account.Remove(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccountReply{}, nil
}
func (s *AuthService) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountReply, error) {
	a, err := s.account.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountReply{
		Account: &pb.Account{
			Id:     int64(a.ID),
			Email:  a.Email,
			Type:   a.Type,
			Status: a.Status,
		},
	}, nil
}
func (s *AuthService) ListAccount(ctx context.Context, req *pb.ListAccountRequest) (*pb.ListAccountReply, error) {
	as, err := s.account.List(ctx)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListAccountReply{}
	for _, a := range as {
		reply.Results = append(reply.Results, &pb.Account{
			Id:     int64(a.ID),
			Email:  a.Email,
			Type:   a.Type,
			Status: a.Status,
		})
	}
	return reply, nil
}
func (s *AuthService) CreateToken(ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	return &pb.CreateTokenReply{}, nil
}
func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	return &pb.RefreshTokenReply{}, nil
}
func (s *AuthService) RevokeToken(ctx context.Context, req *pb.RevokeTokenRequest) (*pb.RevokeTokenReply, error) {
	return &pb.RevokeTokenReply{}, nil
}
