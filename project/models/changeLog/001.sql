CREATE TABLE groups(
    group_id Serial primary key,
    title varchar(255) not null
);
CREATE TABLE tasks(
    task_id serial primary key,
    title varchar(255) not null,
    group_id int REFERENCES groups(group_id) ON DELETE CASCADE
);
CREATE TABLE time_frames(
    task_id int REFERENCES tasks(task_id) ON DELETE CASCADE,
    start_at timestamp without time zone,
    end_at timestamp without time zone 
);
