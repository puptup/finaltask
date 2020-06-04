CREATE TABLE groups(
    group_id Serial primary key,
    title varchar(255)
);
CREATE TABLE tasks(
    task_id serial primary key,
    title varchar(255),
    group_id int REFERENCES groups(group_id) ON DELETE CASCADE
);
CREATE TABLE time_frames(
    task_id int REFERENCES tasks(task_id) ON DELETE CASCADE,
    start_at varchar(255) not null,
    end_at varchar(255) not null 
);
insert into groups(title) values('Hi there'),('Hi everyone');
insert into tasks(title, group_id) values('first', 1),('second',2);