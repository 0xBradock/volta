package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ace-astro/volta/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntryByID(t *testing.T) {
	e1 := createRandomEntry(t)
	e2, err := testQueries.GetEntryByID(context.Background(), e1.ID)

	require.NoError(t, err)
	require.Equal(t, e1.ID, e2.ID)
	require.Equal(t, e1.Domain, e2.Domain)
	require.Equal(t, e1.Login, e2.Login)
	require.Equal(t, e1.Password, e2.Password)
	require.Equal(t, e1.Meta.String, e2.Meta.String)
	require.Equal(t, e1.Created, e2.Created)
	require.Equal(t, e1.Updated, e2.Updated)
}

func TestGetEntriesByDomain(t *testing.T) {
	e1 := createRandomEntry(t)
	e2, err := testQueries.GetEntriesByDomain(context.Background(), e1.Domain)

	require.NoError(t, err)
	require.Equal(t, e1.ID, e2[0].ID)
	require.Equal(t, e1.Domain, e2[0].Domain)
	require.Equal(t, e1.Login, e2[0].Login)
	require.Equal(t, e1.Password, e2[0].Password)
	require.Equal(t, e1.Meta.String, e2[0].Meta.String)
	require.Equal(t, e1.Created, e2[0].Created)
	require.Equal(t, e1.Updated, e2[0].Updated)
}

func TestUpdateEntry(t *testing.T) {
	e1 := createRandomEntry(t)

	newMeta := util.RandomMeta(2)
	newPass := util.RandomPass()
	newLogin := util.RandomLogin()

	e2, err := testQueries.UpdateEntry(
		context.Background(),
		UpdateEntryParams{
			ID:       e1.ID,
			Login:    sql.NullString{String: newLogin, Valid: true},
			Password: newPass,
			Meta:     sql.NullString{String: newMeta, Valid: true},
		},
	)

	require.NoError(t, err)
	require.Equal(t, e1.ID, e2.ID)
	require.Equal(t, e1.Domain, e2.Domain)
	require.Equal(t, e1.Created, e2.Created)

	require.NotEqual(t, e1.Login, e2.Login)
	require.NotEqual(t, e2.Login, newLogin)

	require.NotEqual(t, e1.Password, e2.Password)
	require.Equal(t, e2.Password, newPass)

	require.NotEqual(t, e1.Meta.String, e2.Meta.String)
	require.Equal(t, e2.Meta.String, newMeta)

	require.NotEqual(t, e1.Updated, e2.Updated)
}

func TestDeleteEntry(t *testing.T) {
	e1 := createRandomEntry(t)
	e2, err := testQueries.DeleteEntry(context.Background(), e1.ID)

	require.NoError(t, err)
	require.Equal(t, e1.ID, e2.ID)
	require.Equal(t, e1.Domain, e2.Domain)
	require.Equal(t, e1.Login, e2.Login)
	require.Equal(t, e1.Password, e2.Password)
	require.Equal(t, e1.Meta.String, e2.Meta.String)
	require.Equal(t, e1.Created, e2.Created)
	require.Equal(t, e1.Updated, e2.Updated)

	e3, err := testQueries.GetEntryByID(context.Background(), e1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, e3)
}

func TestGetEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := GetEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.GetEntries(context.Background(), arg)

	require.NoError(t, err)

	require.Len(t, entries, 5)

  for _, entry := range entries {
    require.NotEmpty(t, entry)
  }
}

func createRandomEntry(t *testing.T) Entry {
	args := CreateEntryParams{
		ID:       uuid.New(),
		Domain:   util.RandomDomain(),
		Login:    sql.NullString{String: util.RandomLogin(), Valid: true},
		Password: util.RandomString(8),
		Meta:     sql.NullString{String: util.RandomMeta(5), Valid: true},
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, entry.ID, args.ID)
	require.Equal(t, entry.Domain, args.Domain)
	require.Equal(t, entry.Login.String, args.Login.String)
	require.Equal(t, entry.Password, args.Password)
	require.Equal(t, entry.Meta.String, args.Meta.String)

	return entry
}
