package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"course_photos/internal/models"
	"course_photos/pkg/dates"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (id int64, err error) {
	err = r.db.GetContext(
		ctx,
		&id,
		queryCreateUser,
		user.Login,
		user.HashedPassword,
		user.Name,
		user.Phone,
	)
	if err != nil {
		return id, err
	}

	return id, nil
}
func (r *Repository) GetUserByLogin(ctx context.Context, login string) (user models.User, err error) {
	err = r.db.GetContext(
		ctx,
		&user,
		queryGetUserByLogin,
		login,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (r *Repository) GetUserById(ctx context.Context, id int64) (user models.User, err error) {
	err = r.db.GetContext(
		ctx,
		&user,
		queryGetUserById,
		id,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (r *Repository) GetSessionUserId(ctx context.Context, sessionId uuid.UUID) (userId int64, err error) {
	err = r.db.GetContext(
		ctx,
		&userId,
		queryGetSessionUserId,
		sessionId,
	)
	if err != nil {
		return userId, err
	}

	return userId, nil
}

func (r *Repository) CreateSession(ctx context.Context, userId int64) (sessionId uuid.UUID, err error) {
	err = r.db.GetContext(
		ctx,
		&sessionId,
		queryCreateSession,
		userId,
	)
	if err != nil {
		return sessionId, err
	}

	return sessionId, nil
}

func (r *Repository) GetStudio(ctx context.Context, id int64) (res models.Studio, err error) {
	err = r.db.GetContext(
		ctx,
		&res,
		queryGetStudio,
		id,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repository) GetStudios(ctx context.Context) (res []models.CatalogItem, err error) {
	res = make([]models.CatalogItem, 0)

	err = r.db.SelectContext(
		ctx,
		&res,
		queryGetStudios,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}
func (r *Repository) GetStudioBookings(ctx context.Context, studioId int64, date dates.Date) (res []int, err error) {
	res = make([]int, 0)

	err = r.db.SelectContext(
		ctx,
		&res,
		queryGetStudioBookings,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repository) GetMyStudios(ctx context.Context, userId int64) (res []models.Studio, err error) {
	res = make([]models.Studio, 0)

	err = r.db.SelectContext(
		ctx,
		&res,
		queryGetMyStudios,
		userId,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *Repository) CreateStudio(ctx context.Context, s models.Studio) error {
	err := r.db.GetContext(
		ctx,
		&s.Id,
		queryCreateStudio,
		s.OwnerUserId,
		s.Name,
		s.Address,
		s.Description,
	)
	if err != nil {
		return err
	}

	if len(s.PhotosIds) != 0 {
		_, err = r.db.ExecContext(
			ctx,
			queryUpdatePhotosStudioId,
			s.Id,
			pq.Int64Array(s.PhotosIds),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
func (r *Repository) DeleteStudio(ctx context.Context, userId int64, studioId int64) error {
	_, err := r.db.ExecContext(
		ctx,
		queryDeleteStudio,
		userId,
		studioId,
	)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repository) UploadPhoto(ctx context.Context, userId int64, photo []byte) (id int64, err error) {
	err = r.db.GetContext(
		ctx,
		&id,
		queryUploadPhoto,
		userId,
		photo,
	)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *Repository) GetPhoto(ctx context.Context, id int64) (res []byte, err error) {
	res = make([]byte, 0)

	err = r.db.GetContext(
		ctx,
		&res,
		queryGetPhoto,
		id,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}
