package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/vish9812/coin/internal/common/random"
)

func createRandomUser(t *testing.T) CreateUserRow {
	name := random.String(7)
	userParam := CreateUserParams{
		Email:     name + "@test.com",
		Password:  name,
		FirstName: name,
		LastName:  &name,
	}

	userRow, err := testQueries.CreateUser(context.Background(), userParam)
	require.NoError(t, err)

	require.NotEmpty(t, userRow)
	require.NotZero(t, userRow.ID)
	require.NotZero(t, userRow.CreatedAt)

	require.Equal(t, userParam.FirstName, userRow.FirstName)
	require.Equal(t, userParam.LastName, userRow.LastName)

	return userRow
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	createRandomUser(t)
}

func TestDeleteUser(t *testing.T) {
	t.Parallel()

	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	user1, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, pgx.ErrNoRows)
	require.Empty(t, user1)
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	user := createRandomUser(t)

	user1, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user1)
	require.Equal(t, user.ID, user1.ID)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.FirstName, user1.FirstName)
	require.Equal(t, user.LastName, user1.LastName)
	require.WithinDuration(t, user.CreatedAt.Time, user1.CreatedAt.Time, time.Second)
}

func TestGetUserPassword(t *testing.T) {
	t.Parallel()

	name := random.String(7)
	userParam := CreateUserParams{
		Email:     name + "@test.com",
		Password:  name,
		FirstName: name,
		LastName:  &name,
	}

	user, err := testQueries.CreateUser(context.Background(), userParam)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	pass, err := testQueries.GetUserPassword(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pass)
	require.Equal(t, userParam.Password, pass)
}

func TestListUsers(t *testing.T) {
	t.Parallel()

	n := 5
	for i := 0; i < n; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  3,
		Offset: 2,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Len(t, users, int(arg.Limit))

	for i := 0; i < int(arg.Limit); i++ {
		require.NotEmpty(t, users[i])
	}
}

func TestUpdateUserPassword(t *testing.T) {
	t.Parallel()

	name := random.String(7)
	userParam := CreateUserParams{
		Email:     name + "@test.com",
		Password:  name,
		FirstName: name,
		LastName:  &name,
	}

	user, err := testQueries.CreateUser(context.Background(), userParam)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	pass, err := testQueries.GetUserPassword(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pass)
	require.Equal(t, userParam.Password, pass)

	arg := UpdateUserPasswordParams{
		ID:       user.ID,
		Password: random.String(7),
	}

	err = testQueries.UpdateUserPassword(context.Background(), arg)
	require.NoError(t, err)

	pass, err = testQueries.GetUserPassword(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pass)
	require.Equal(t, arg.Password, pass)
}
