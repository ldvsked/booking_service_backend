create table account(
    id uuid primary key,
    email text not null, 
    role text not null сheck(role in ('user', 'admin')), 
    created_at timestamptz default now()
);

create table room(
    id uuid primary key, 
    name string not null, 
    description string, 
    capacity int, 
    created_at timestamptz default now()
);

create table schedule(
    id uuid primary key, 
    room_id uuid not null, 
    days_of_week []int not null, ---?  как добавить ограничения что мин 1 и макс 7?
    start_time text not null ---паттерн надо добавлять? , 
    end_time text not null
)