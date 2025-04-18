package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Maksim646/tokens/database/postgresql"
	"github.com/Maksim646/tokens/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type UserRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IUserRepository {
	return &UserRepository{sqalxConn: sqalxConn}
}

func (r *UserRepository) CreateUserByEmail(ctx context.Context, userID string, email string) (string, error) {

	_, err := uuid.Parse(userID)
	if err != nil {
		return "", model.ErrInvalidUserID
	}

	query, params, err := postgresql.Builder.Insert("users").
		Columns(
			"id",
			"email",
		).
		Values(
			userID,
			email,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return "", err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))

	var newUserID string
	err = r.sqalxConn.GetContext(ctx, &newUserID, query, params...)
	return newUserID, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (model.User, error) {
	var user model.User
	query, params, err := postgresql.Builder.Select(
		"id",
		"email",
		"created_at",
	).
		From("users").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return user, err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	if err = r.sqalxConn.GetContext(ctx, &user, query, params...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, model.ErrUserNotFound
		}
	}

	return user, err
}
