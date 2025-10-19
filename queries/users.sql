-- name: CreateUser :one
insert into users (phone_number, first_name, second_name)
values ($1, $2, $3)
returning *;

-- name: GetUser :one
select * from users
where id = $1;

-- name: DeleteUser :exec
delete from users where id = $1;

-- name: UpdateUser :one
update users
set 
    first_name = coalesce(sqlc.narg('first_name'), first_name),
    second_name = coalesce(sqlc.narg('second_name'), second_name),
    phone_number = coalesce(sqlc.narg('phone_number'), phone_number)
where id = sqlc.arg('id')
returning *;

