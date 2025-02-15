package migrations

import (
	"github.com/ledgerwatch/erigon/common/dbutils"
	"github.com/ledgerwatch/erigon/common/etl"
	"github.com/ledgerwatch/erigon/ethdb"
)

var removeCliqueBucket = Migration{
	Name: "remove_clique_bucket",
	Up: func(db ethdb.Database, tmpdir string, progress []byte, CommitProgress etl.LoadCommitHandler) (err error) {

		if exists, err := db.(ethdb.BucketsMigrator).BucketExists(dbutils.CliqueBucket); err != nil {
			return err
		} else if !exists {
			return CommitProgress(db, nil, true)
		}

		if err := db.(ethdb.BucketsMigrator).DropBuckets(dbutils.CliqueBucket); err != nil {
			return err
		}

		return CommitProgress(db, nil, true)
	},
}
