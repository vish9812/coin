package sqlc

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgres://coinmaster:coinpass@localhost:5432/coin?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
	funcLocation := slog.String("func", "sqlc/TestMain")
	redactedDBSource := strings.ReplaceAll(dbSource, "coinpass", strings.Repeat("*", 5))
	redactedDBSource = strings.ReplaceAll(redactedDBSource, "coinmaster", strings.Repeat("*", 5))
	slog.Info("connect to DB", slog.String("dbSource", redactedDBSource), funcLocation)

	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		slog.Error("unable to connect to database", slog.Any("err", err), funcLocation)
		os.Exit(1)
	}
	defer pool.Close()

	var errPing error
	for i := 0; i < 10; i++ {
		slog.Debug("pinging database...", funcLocation)

		errPing = pool.Ping(context.Background())
		if errPing == nil {
			break
		}
		slog.Warn("failed to ping db", funcLocation)
		time.Sleep(2 * time.Second)
	}

	if errPing != nil {
		slog.Error("unable to ping to database", slog.Any("err", errPing), funcLocation)
		os.Exit(1)
	}

	testQueries = New(pool)

	os.Exit(m.Run())
}
