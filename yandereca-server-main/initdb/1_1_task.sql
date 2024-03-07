create table tasks (
    task_id text UNIQUE,
    name text not null default '',
    description text not null default '',
    is_done boolean default false,
    user_id text not null default '',
    task_status integer not null default 0,
    deadline timestamp with time zone default current_timestamp,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    constraint task_pkey PRIMARY KEY(task_id)
);
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-1',
        'hogehoge1',
        'example task 1',
        false,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-2',
        'hogehoge2',
        'example task 2',
        false,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-3',
        'hogehoge3',
        'example task 3',
        false,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-4',
        'hogehoge4',
        'example task 4',
        false,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-5',
        'hogehoge5',
        'example task 5',
        false,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-6',
        'hogehoge6',
        'example task 6',
        false,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-7',
        'hogehoge7',
        'example task 7',
        true,
        'example-user-id'
    );
insert into tasks(task_id, name, description, is_done, user_id)
values(
        'example-task-id-8',
        'hogehoge8',
        'example task 8',
        true,
        'example-user-id'
    );