package usecase

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"course_photos/internal/models"
	"course_photos/internal/repository"
	"course_photos/pkg/dates"
	"course_photos/pkg/passwords"
)

type UseCase struct {
	r *repository.Repository
}

func New(r *repository.Repository) *UseCase {
	return &UseCase{
		r: r,
	}
}

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

func (uc *UseCase) Register(ctx context.Context, req RegisterRequest) (sessionId uuid.UUID, err error) {
	hashedPassword, err := passwords.HashPassword(req.Password)
	if err != nil {
		return sessionId, err
	}

	id, err := uc.r.CreateUser(ctx, models.User{
		Id:             0,
		Login:          req.Login,
		HashedPassword: hashedPassword,
		Name:           req.Name,
		Phone:          req.Phone,
	})
	if err != nil {
		return sessionId, err
	}

	sessionId, err = uc.r.CreateSession(ctx, id)
	if err != nil {
		return sessionId, err
	}

	return sessionId, nil
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (uc *UseCase) Login(ctx context.Context, req LoginRequest) (sessionId uuid.UUID, err error) {
	user, err := uc.r.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return sessionId, err
	}

	hashedPassword, err := passwords.HashPassword(req.Password)
	if err != nil {
		return sessionId, err
	}

	if user.HashedPassword != hashedPassword {
		return sessionId, errors.New("invalid password")
	}

	sessionId, err = uc.r.CreateSession(ctx, user.Id)
	if err != nil {
		return sessionId, err
	}

	return sessionId, nil
}

func (uc *UseCase) Auth(ctx context.Context, sessionId uuid.UUID) (user models.User, err error) {
	userId, err := uc.r.GetSessionUserId(ctx, sessionId)
	if err != nil {
		return user, err
	}

	user, err = uc.r.GetUserById(ctx, userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

type CreateStudioRequest struct {
	OwnerUserId int64   `json:"owner_user_id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	PhotosIds   []int64 `json:"photos_ids"`
}

func (uc *UseCase) CreateStudio(ctx context.Context, userId int64, req CreateStudioRequest) error {
	return uc.r.CreateStudio(ctx, models.Studio{
		Id:          0,
		OwnerUserId: userId,
		Name:        req.Name,
		Address:     req.Address,
		Description: req.Description,
		PhotosIds:   req.PhotosIds,
	})
}

type UpdateStudioRequest struct {
	Name        *string       `json:"name"`
	Address     *string       `json:"address"`
	Description *string       `json:"description"`
	PhotosIds   pq.Int64Array `json:"photos_ids"`
}

func (uc *UseCase) UpdateStudio(ctx context.Context, userId int64, studioId int64, req UpdateStudioRequest) error {
	return uc.r.UpdateStudio(ctx, userId, studioId)
}
func (uc *UseCase) DeleteStudio(ctx context.Context, userId int64, studioId int64) error {
	return uc.r.DeleteStudio(ctx, userId, studioId)
}
func (uc *UseCase) UploadPhoto(ctx context.Context, userId int64, photo []byte) (id int64, err error) {
	return uc.r.UploadPhoto(ctx, userId, photo)
}

func (uc *UseCase) GetStudios(ctx context.Context) ([]models.CatalogItem, error) {
	return uc.r.GetStudios(ctx)
}
func (uc *UseCase) GetStudioAvailableHours(
	ctx context.Context,
	studioId int64,
	date dates.Date,
) (hours []int, err error) {
	studio, err := uc.r.GetStudio(ctx, studioId)
	if err != nil {
		return hours, err
	}

	working := false

	switch time.Time(date).Weekday() {
	case time.Sunday:
		working = studio.WorksOnSunday
	case time.Monday:
		working = studio.WorksOnMonday
	case time.Tuesday:
		working = studio.WorksOnTuesday
	case time.Wednesday:
		working = studio.WorksOnWednesday
	case time.Thursday:
		working = studio.WorksOnThursday
	case time.Friday:
		working = studio.WorksOnFriday
	case time.Saturday:
		working = studio.WorksOnSaturday
	}

	if !working {
		return hours, errors.New("студия не работает в этот день недели")
	}

	bookedHours, err := uc.r.GetStudioBookings(ctx, studioId, date)
	if err != nil {
		return hours, err
	}

	for i := range 23 {
		if slices.Contains(bookedHours, i) {
			continue
		}

		hours = append(hours, i)
	}

	return hours, nil
}
func (uc *UseCase) GetPhoto(ctx context.Context, id int64) ([]byte, error) {
	return uc.r.GetPhoto(ctx, id)
}

func (uc *UseCase) GetMyStudios(ctx context.Context, userId int64) ([]models.Studio, error) {
	return uc.r.GetMyStudios(ctx, userId)
}
