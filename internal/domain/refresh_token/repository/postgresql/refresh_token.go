package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Maksim646/tokens/database/postgresql"
	"github.com/Maksim646/tokens/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/heetch/sqalx"
	"go.uber.org/zap"
)

type RefreshTokenRepository struct {
	sqalxConn sqalx.Node
}

func New(sqalxConn sqalx.Node) model.IRefreshTokenRepository {
	return &RefreshTokenRepository{sqalxConn: sqalxConn}
}

func (r *RefreshTokenRepository) CreateRefreshToken(ctx context.Context, refreshToken model.RefreshToken) error {
	query, params, err := postgresql.Builder.Insert("refresh_tokens").
		Columns(
			"token_hash",
			"id",
			"user_id",
			"expired_at",
		).
		Values(
			refreshToken.TokenHash,
			refreshToken.ID,
			refreshToken.UserID,
			refreshToken.ExpiredAt,
		).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug("CreateRefreshToken SQL", zap.String("query", query), zap.Any("params", params))

	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	if err != nil {
		zap.L().Error("failed to execute query", zap.Error(err))
		return err
	}

	return nil
}

func (r *RefreshTokenRepository) GetRefreshTokenByUserID(ctx context.Context, userID string) (model.RefreshToken, error) {
	query, params, err := postgresql.Builder.
		Select(
			"token_hash",
			"id",
			"user_id",
			"created_at",
			"expired_at",
		).
		From("refresh_tokens").
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return model.RefreshToken{}, err
	}

	zap.L().Debug("GetRefreshTokenByUserID SQL", zap.String("query", query), zap.Any("params", params))

	row := r.sqalxConn.QueryRowxContext(ctx, query, params...)

	var refreshToken model.RefreshToken
	err = row.StructScan(&refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.RefreshToken{}, sql.ErrNoRows
		}
		zap.L().Error("failed to scan refresh token", zap.Error(err))
		return model.RefreshToken{}, err
	}

	return refreshToken, nil
}

func (r *RefreshTokenRepository) DeleteRefreshTokenByID(ctx context.Context, refreshTokenID string) error {
	query, params, err := postgresql.Builder.Delete("refresh_tokens").
		Where(sq.Eq{"refresh_tokens.id": refreshTokenID}).
		ToSql()
	if err != nil {
		return err
	}

	zap.L().Debug(postgresql.BuildQuery(query, params))
	_, err = r.sqalxConn.ExecContext(ctx, query, params...)
	return err
}
