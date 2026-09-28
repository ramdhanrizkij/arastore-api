package migrations

import (
	"testing"

	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectMigrations(t *testing.T) {
	goose.SetBaseFS(EmbedMigrations())

	migrations, err := goose.CollectMigrations(".", 0, 1000)
	require.NoError(t, err)
	assert.Len(t, migrations, 25)

	// Verify each migration is parsed and has a valid version
	for i, m := range migrations {
		expectedVersion := int64(i + 1)
		assert.Equal(t, expectedVersion, m.Version, "migration version mismatch for %s", m.Source)
		assert.NotEmpty(t, m.Source)
	}
}
