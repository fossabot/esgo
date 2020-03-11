package serverless_es_go

import (
	"testing"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
)

var connectionString = testConfig.Postgres.ConnectionString

func TestSaveNewCheckpoint(t *testing.T) {
	cfg := CheckpointConfig{ConnectionString: connectionString, ProjectionName: "test-projection"}
	timestamp := getTimestamp()
	err := SaveCheckpoint(&cfg, 1, timestamp)
	assert.Nil(t, err)
	verifyCheckpoint(t, "test-projection", 1, timestamp)
}

func verifyCheckpoint(t *testing.T, name string, position int, timestamp int64) {
	db, err := sql.Open("postgres", connectionString)
	assert.Nil(t, err)

	var projectionName string
	var checkpointPosition int
	var checkpointTimestamp int64

	err = db.QueryRow("SELECT * FROM checkpoints WHERE name = $1", name).Scan(&projectionName, &checkpointPosition, &checkpointTimestamp)
	assert.Nil(t, err)

	assert.Equal(t, name, projectionName)
	assert.Equal(t, position, checkpointPosition)
	assert.Equal(t, timestamp, checkpointTimestamp)
}