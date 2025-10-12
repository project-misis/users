-- name: CreateUser :exec
insert into users (email, first_name, second_name, password_hash)
values ($1, $2, $3, $4);

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
    email = coalesce(sqlc.narg('email'), email)
where id = sqlc.arg('id')
returning *;

