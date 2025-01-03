package models

import (
	"time"

	"github.com/lib/pq"

	"course_photos/pkg/dates"
)

type (
	User struct {
		Id             int64  `json:"id" db:"id"`
		Login          string `json:"login" db:"login"`
		HashedPassword string `json:"-" db:"hashed_password"`
		Name           string `json:"name" db:"name"`
		Phone          string `json:"phone" db:"phone"`
	}

	Studio struct {
		Id          int64         `json:"id" db:"id"`
		OwnerUserId int64         `json:"owner_user_id" db:"owner_user_id"`
		Name        string        `json:"name" db:"name"`
		Address     string        `json:"address" db:"address"`
		Description string        `json:"description" db:"description"`
		PhotosIds   pq.Int64Array `json:"photos_ids" db:"photos_ids"`

		AvailableHoursFrom int `json:"available_hours_from" db:"available_hours_from"`
		AvailableHoursTo   int `json:"available_hours_to" db:"available_hours_to"`

		WorksOnMonday    bool `json:"works_on_monday" db:"works_on_monday"`
		WorksOnTuesday   bool `json:"works_on_tuesday" db:"works_on_tuesday"`
		WorksOnWednesday bool `json:"works_on_wednesday" db:"works_on_wednesday"`
		WorksOnThursday  bool `json:"works_on_thursday" db:"works_on_thursday"`
		WorksOnFriday    bool `json:"works_on_friday" db:"works_on_friday"`
		WorksOnSaturday  bool `json:"works_on_saturday" db:"works_on_saturday"`
		WorksOnSunday    bool `json:"works_on_sunday" db:"works_on_sunday"`
	}

	CatalogItem struct {
		Id          int64         `json:"id" db:"id"`
		OwnerUserId int64         `json:"owner_user_id" db:"owner_user_id"`
		OwnerName   string        `json:"owner_name" db:"owner_name"`
		OwnerPhone  string        `json:"owner_phone" db:"owner_phone"`
		Name        string        `json:"name" db:"name"`
		Address     string        `json:"address" db:"address"`
		Description string        `json:"description" db:"description"`
		PhotosIds   pq.Int64Array `json:"photos_ids" db:"photos_ids"`

		AvailableHoursFrom int `json:"available_hours_from" db:"available_hours_from"`
		AvailableHoursTo   int `json:"available_hours_to" db:"available_hours_to"`

		WorksOnMonday    bool `json:"works_on_monday" db:"works_on_monday"`
		WorksOnTuesday   bool `json:"works_on_tuesday" db:"works_on_tuesday"`
		WorksOnWednesday bool `json:"works_on_wednesday" db:"works_on_wednesday"`
		WorksOnThursday  bool `json:"works_on_thursday" db:"works_on_thursday"`
		WorksOnFriday    bool `json:"works_on_friday" db:"works_on_friday"`
		WorksOnSaturday  bool `json:"works_on_saturday" db:"works_on_saturday"`
		WorksOnSunday    bool `json:"works_on_sunday" db:"works_on_sunday"`
	}

	Booking struct {
		Id        int64      `json:"id" db:"id"`
		UserId    int64      `json:"user_id" db:"user_id"`
		StudioId  int64      `json:"studio_id" db:"studio_id"`
		Date      dates.Date `json:"date" db:"date"`
		Hours     int        `json:"hours" db:"hours"`
		CreatedAt time.Time  `json:"created_at" db:"created_at"`
	}
	MyBookingsItem struct {
		Id        int64      `json:"id" db:"id"`
		Date      dates.Date `json:"date" db:"date"`
		Hours     int        `json:"hours" db:"hours"`
		CreatedAt time.Time  `json:"created_at" db:"created_at"`

		StudioId          int64         `json:"studio_id" db:"studio_id"`
		StudioName        string        `json:"studio_name" db:"studio_name"`
		StudioAddress     string        `json:"studio_address" db:"studio_address"`
		StudioDescription string        `json:"studio_description" db:"studio_description"`
		StudioPhotosIds   pq.Int64Array `json:"studio_photos_ids" db:"studio_photos_ids"`

		OwnerUserId int64  `json:"owner_user_id" db:"owner_user_id"`
		OwnerName   string `json:"owner_name" db:"owner_name"`
		OwnerPhone  string `json:"owner_phone" db:"owner_phone"`
	}
	StudioBookingsItem struct {
		Id        int64      `json:"id" db:"id"`
		Date      dates.Date `json:"date" db:"date"`
		Hours     int        `json:"hours" db:"hours"`
		CreatedAt time.Time  `json:"created_at" db:"created_at"`

		GuestUserId int64  `json:"guest_user_id" db:"guest_user_id"`
		GuestName   string `json:"guest_name" db:"guest_name"`
		GuestPhone  string `json:"guest_phone" db:"guest_phone"`
	}
)
