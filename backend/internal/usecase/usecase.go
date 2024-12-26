package usecase

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"

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

	if !passwords.CheckPasswordHash(req.Password, user.HashedPassword) {
		return sessionId, errors.New("invalid password")
	}

	sessionId, err = uc.r.CreateSession(ctx, user.Id)
	if err != nil {
		return sessionId, err
	}

	return sessionId, nil
}

func (uc *UseCase) Logout(ctx context.Context, sessionId uuid.UUID) error {
	return uc.r.DeleteSession(ctx, sessionId)
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

func (uc *UseCase) DeleteMyStudio(ctx context.Context, userId int64, studioId int64) error {
	s, err := uc.r.GetStudio(ctx, studioId)
	if err != nil {
		return err
	}

	if s.OwnerUserId != userId {
		return errors.New("invalid studio")
	}

	err = uc.r.DeleteStudio(ctx, userId, studioId)
	if err != nil {
		return err
	}

	return nil
}
func (uc *UseCase) GetMyStudioBookings(ctx context.Context, userId int64, studioId int64) (res []models.StudioBookingsItem, err error) {
	s, err := uc.r.GetStudio(ctx, studioId)
	if err != nil {
		return res, err
	}

	if s.OwnerUserId != userId {
		return res, errors.New("invalid studio")
	}

	res, err = uc.r.GetMyStudioBookings(ctx, userId)
	if err != nil {
		return res, err
	}

	return res, nil
}
func (uc *UseCase) UploadPhoto(ctx context.Context, userId int64, photo []byte) (id int64, err error) {
	return uc.r.UploadPhoto(ctx, userId, photo)
}

func (uc *UseCase) GetStudios(ctx context.Context, userId int64) ([]models.CatalogItem, error) {
	return uc.r.GetStudios(ctx, userId)
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

	bookedHours, err := uc.r.GetStudioBookedHours(ctx, studioId, date)
	if err != nil {
		return hours, err
	}

	for i := range 24 {
		if slices.Contains(bookedHours, i) {
			continue
		}

		hours = append(hours, i)
	}

	return hours, nil
}

type CreateBookingsRequest struct {
	StudioId int64      `params:"id"`
	Date     dates.Date `json:"date"`
	Hours    []int      `json:"hours"`
}

func (uc *UseCase) CreateBookings(ctx context.Context, userId int64, req CreateBookingsRequest) error {
	hours, err := uc.GetStudioAvailableHours(ctx, req.StudioId, req.Date)
	if err != nil {
		return err
	}

	for _, h := range req.Hours {
		if !slices.Contains(hours, h) {
			return errors.New("unavailable hour")
		}
	}

	for _, h := range req.Hours {
		err = uc.r.CreateBooking(ctx, models.Booking{
			Id:        0,
			UserId:    userId,
			StudioId:  req.StudioId,
			Date:      req.Date,
			Hours:     h,
			CreatedAt: time.Now(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
func (uc *UseCase) GetPhoto(ctx context.Context, id int64) ([]byte, error) {
	return uc.r.GetPhoto(ctx, id)
}
func (uc *UseCase) GetMyStudios(ctx context.Context, userId int64) ([]models.Studio, error) {
	return uc.r.GetMyStudios(ctx, userId)
}
func (uc *UseCase) GetMyBookings(ctx context.Context, userId int64) ([]models.MyBookingsItem, error) {
	return uc.r.GetMyBookings(ctx, userId)
}
func (uc *UseCase) DeleteMyBooking(ctx context.Context, userId int64, bookingId int64) error {
	b, err := uc.r.GetBooking(ctx, bookingId)
	if err != nil {
		return err
	}

	if b.UserId != userId {
		return errors.New("forbidden")
	}

	err = uc.r.DeleteBooking(ctx, bookingId)
	if err != nil {
		return err
	}

	return nil
}
