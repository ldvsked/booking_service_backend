create table account(
    id uuid primary key,
    email text not null, 
    role text not null check(role in ('user', 'admin')), 
    created_at timestamptz default now()
);

create table room(
    id uuid primary key, 
    name text not null, 
    description text, 
    capacity integer, 
    created_at timestamptz default now()
);

create table schedule(
    id uuid primary key, 
    room_id uuid unique not null references room(id) on delete cascade, 
    days_of_week integer[] not null check((1 <= all(days_of_week)) and (7 >= all(days_of_week))), 
    start_time text not null, 
    end_time text not null
); 

create table slot(
    id uuid primary key, 
    room_id uuid not null references room(id), 
    start_time timestamptz not null, 
    end_time timestamptz not null
);

create table booking(
    id uuid primary key, 
    slot_id uuid not null references slot(id), 
    user_id uuid not null references account(id), 
    status text not null check(status in('active', 'cancelled')), 
    conference_link text, 
    created_at timestamptz default now()
);

