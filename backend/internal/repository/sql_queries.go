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
)
