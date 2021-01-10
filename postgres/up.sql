DROP TABLE IF EXISTS public.booking;
DROP TABLE IF EXISTS public.hotelroom;

CREATE TABLE public.booking
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    roomid integer NOT NULL,
    startdate date NOT NULL,
    finishdate date NOT NULL,
    CONSTRAINT booking_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
        )
    TABLESPACE pg_default;

ALTER TABLE public.booking
    OWNER to boba;

CREATE TABLE public.hotelroom
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    description character varying COLLATE pg_catalog."default",
    price integer,
    dateadding date,
    CONSTRAINT hotelnumber_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
        )
    TABLESPACE pg_default;

ALTER TABLE public.hotelroom
    OWNER to boba;