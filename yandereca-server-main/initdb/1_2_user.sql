create table userdata (
    id text not null UNIQUE,
    name text not null default '',
    email text not null default '',
    token text not null default '',
    refresh_token text not null default '',
    google_uid text not null default '',
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    constraint userdata_pkey PRIMARY KEY(id)
);

insert into userdata(id, name, email, token, refresh_token, google_uid)
values(
    'example-id',
    'example-name',
    'example-email',
    'example-token',
    'example-refresh_token',
    'example-google_uid'
);