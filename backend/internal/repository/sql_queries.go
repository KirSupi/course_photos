package repository

const (
	queryCreateUser = `
		INSERT INTO public.users(login, hashed_password, name, phone)
		VALUES ($1, $2, $3, $4)
		returning id
	`
	queryCreateSession = `
		INSERT INTO public.sessions(user_id)
		VALUES ($1)
		returning id
	`
	queryDeleteSession = `
		DELETE FROM public.sessions
		WHERE id = $1
	`
	queryGetSessionUserId = `
		SELECT user_id
		FROM public.sessions
		WHERE id=$1
	`
	queryGetUserByLogin = `
		SELECT *
		FROM public.users
		WHERE login = $1
	`
	queryGetUserById = `
		SELECT *
		FROM public.users
		WHERE id = $1
	`
	queryCreateStudio = `
		INSERT INTO public.studios(owner_user_id, name, address, description)
		VALUES ($1, $2, $3, $4)
		returning id
	`
	queryDeleteStudio = `
		DELETE FROM public.studios
		WHERE owner_user_id=$1 AND id=$2
	`
	queryUploadPhoto = `
		INSERT INTO public.photos(studio_id, owner_user_id, photo)
		VALUES (0, $1, $2)
		returning id
	`
	queryGetPhoto = `
		SELECT photo
		FROM public.photos
		WHERE id = $1
	`
	queryUpdatePhotosStudioId = `
		UPDATE public.photos
		SET studio_id=$1::BIGINT
		WHERE id = ANY ($2::BIGINT[])
	`
	queryGetStudios = `
		SELECT s.*,
			   (SELECT array_agg(p.id)
				FROM photos p
				WHERE p.studio_id = s.id) AS "photos_ids",
			   u.name                     AS "owner_name",
			   u.phone                    AS "owner_phone"
		FROM public.studios s
				 JOIN public.users u ON s.owner_user_id = u.id
		WHERE s.owner_user_id != $1
	`
	queryGetStudio = `
		SELECT s.*,
			   (SELECT array_agg(p.id)
				FROM photos p
				WHERE p.studio_id = s.id) AS "photos_ids"
		FROM public.studios s 
		WHERE s.id = $1
	`
	queryGetMyStudios = `
		SELECT s.*,
			   (SELECT array_agg(p.id)
				FROM photos p
				WHERE p.studio_id = s.id) AS "photos_ids"
		FROM public.studios s 
		WHERE s.owner_user_id = $1
	`
	queryGetStudioBookedHours = `
		SELECT hours
		FROM public.bookings
		WHERE studio_id = $1::BIGINT
		  AND DATE = $2::DATE
	`
	queryCreateBooking = `
		INSERT INTO public.bookings(user_id, studio_id, date, hours)
		VALUES ($1::BIGINT, $2::BIGINT, $3::DATE, $4::BIGINT)
	`
	queryGetMyBookings = `
		SELECT b.id                         AS "id",
			   b.date                       AS "date",
			   b.hours                      AS "hours",
			   b.created_at                 AS "created_at",
			   s.id                         AS "studio_id",
			   s.name                       AS "studio_name",
			   s.address                    AS "studio_address",
			   s.description                AS "studio_description",
			   COALESCE(array_agg(p.id)
						filter ( WHERE p.id IS NOT NULL ),
						array []::BIGINT[]) AS "studio_photos_ids",
			   u.id                         AS "owner_user_id",
			   u.name                       AS "owner_name",
			   u.phone                      AS "owner_phone"
		FROM public.bookings b
				 JOIN public.studios s ON b.studio_id = s.id
				 JOIN public.users u ON s.owner_user_id = u.id
				 LEFT JOIN public.photos p ON s.id = p.studio_id
		WHERE b.user_id = $1
		  AND b.date >= DATE(NOW()) - INTERVAL '1 day'
		GROUP BY b.id, b.date, b.hours, b.created_at,
				 s.id, s.name, s.address, s.description,
				 u.id, u.name, u.phone
		ORDER BY b.date DESC, b.hours
	`
	queryGetMyStudioBookings = `
		SELECT b.id                         AS "id",
			   b.date                       AS "date",
			   b.hours                      AS "hours",
			   b.created_at                 AS "created_at",
			   u.id                         AS "guest_user_id",
			   u.name                       AS "guest_name",
			   u.phone                      AS "guest_phone"
		FROM public.bookings b
				 JOIN public.users u ON b.user_id = u.id
		WHERE b.studio_id = $1
		  AND b.date >= DATE(NOW()) - INTERVAL '1 day'
		GROUP BY b.id, b.date, b.hours, b.created_at,
				 u.id, u.name, u.phone
		ORDER BY b.date DESC, b.hours
	`
	queryGetBooking = `
		SELECT *
		FROM public.bookings
		WHERE id = $1;
	`
	queryDeleteBooking = `
		DELETE
		FROM public.bookings
		WHERE id = $1;
	`
)
