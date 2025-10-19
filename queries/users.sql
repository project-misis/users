-- name: GetUser :one
select * from users
where id = $1;

-- name: DeleteUser :exec
delete from users where id = $1;

-- name: UpdateUser :one
update users
set 
    firstname = coalesce(sqlc.narg('firstname'), firstname),
    username = coalesce(sqlc.narg('username'), username),
    course = coalesce(sqlc.narg('course'), course),
    faculty = coalesce(sqlc.narg('faculty'), faculty)
where id = sqlc.arg('id')
returning *;

