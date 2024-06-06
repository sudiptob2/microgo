package auth

import (
	"context"
	"database/sql"

	pb "github.com/sudiptob2/microgo/auth/proto"
)

type Auth struct {
	pb.UnimplementedAuthServiceServer
	db *sql.DB
}

func New(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (this *Auth) GetToken(ctx context.Context, credentials *pb.Credentials) (*pb.Token, error) {

	return &pb.Token{}, nil

}

func (this *Auth) ValidateToken(ctx context.Context, token *pb.Token) (*pb.User, error) {

	return &pb.User{}, nil
}
