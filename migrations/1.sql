create table users
(
    id            serial primary key,
    email         varchar(50) unique                     not null,
    settings      json,
    created_at    timestamp with time zone default now() not null,
    updated_at    timestamp,
    last_login_at timestamp,
    banned        bool
);

-- task types: 1 - algo, 2 - behavioural, 3 - system design

create table tasks
(
    id            serial primary key,
    text_markdown text                                   not null,
    type          smallint,
    created_at    timestamp with time zone default now() not null,
    updated_at    timestamp with time zone default now() not null
);

-- task status: 1 - not tried, 2 - tried, 3 - solved

create table user_task
(
    user_id    int,
    task_id    int,
    status     smallint,
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);

-- should_start_at: 10-00, 12-00, 14-00, 16-00, 18-00, 20-00, 22-00
-- schedule status: 1 - planned, 2 - started, 3 - finished, 4 - canceled
create table schedule
(
    user_id         int,
    task_id         int,
    should_start_at timestamp,
    status          smallint,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone default now() not null
);
