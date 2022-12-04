CREATE TABLE public.users (
    id bigint NOT NULL,
    username character varying(128) NOT NULL,
    salt character varying(255) NOT NULL,
    password character varying(255),
    email character varying(255),
    confirmed boolean,
    created_at timestamp without time zone NOT NULL
);
