CREATE TABLE public.users
(
    id              SERIAL CONSTRAINT users_pk
            PRIMARY KEY,
    login           TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    name            TEXT NOT NULL,
    phone           TEXT NOT NULL
);

CREATE UNIQUE INDEX users_login_uindex
    ON public.users (login);

CREATE TABLE public.studios
(
    id                   SERIAL,
    owner_user_id        BIGINT               NOT NULL,
    name                 TEXT                 NOT NULL,
    address              TEXT                 NOT NULL,
    description          TEXT                 NOT NULL,
    available_hours_from INTEGER DEFAULT 0    NOT NULL,
    available_hours_to   INTEGER DEFAULT 23   NOT NULL,
    works_on_monday      BOOLEAN DEFAULT TRUE NOT NULL,
    works_on_tuesday     BOOLEAN DEFAULT TRUE NOT NULL,
    works_on_wednesday   BOOLEAN DEFAULT TRUE NOT NULL,
    works_on_thursday    BOOLEAN DEFAULT TRUE NOT NULL,
    works_on_friday      BOOLEAN DEFAULT TRUE NOT NULL,
    works_on_saturday    BOOLEAN DEFAULT TRUE NOT NULL,
    works_on_sunday      BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE public.photos
(
    id            SERIAL,
    studio_id     BIGINT NOT NULL,
    photo BYTEA NOT NULL,
    owner_user_id BIGINT NOT NULL
);

CREATE TABLE public.bookings
(
    id        SERIAL NOT NULL CONSTRAINT bookings_pk
            PRIMARY KEY,
    user_id   BIGINT NOT NULL,
    studio_id BIGINT NOT NULL,
    date      DATE   NOT NULL,
    hours     BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT bookings_pk_2
        UNIQUE (studio_id, date, hours)
);

CREATE TABLE public.sessions
(
    id UUID DEFAULT gen_random_uuid() NOT NULL
        CONSTRAINT sessions_pk
        PRIMARY KEY,
    user_id BIGINT NOT NULL
);

