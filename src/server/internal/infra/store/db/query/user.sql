-- name: GetUser :one
select id, email, first_name, last_name, created_at from "user"
where id = $1 limit 1;

-- name: GetUserPassword :one
select "password" from "user"
where id = $1 limit 1;

-- name: ListUsers :many
select * from "user"
order by first_name, last_name
limit $1 offset $2;

-- name: CreateUser :one
insert into "user" (
    email, "password", first_name, last_name
) values (
    $1, $2, $3, $4
)
returning id, email, first_name, last_name, created_at;

-- name: DeleteUser :exec
delete from "user"
where id = $1;

-- name: UpdateUserPassword :exec
update "user"
    set "password" = $2
where id = $1;