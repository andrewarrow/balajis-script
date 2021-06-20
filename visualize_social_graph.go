package main

import (
	"github.com/dgraph-io/badger/v3"
)

func VisualizeSocialGraph(dir string) {
	db, _ := badger.Open(badger.DefaultOptions(dir))
	defer db.Close()
	PrefixFollowerPKIDToFollowedPKID := byte(28)
	PrefixFollowedPKIDToFollowerPKID := byte(29)
	Follower2Followed(db, []byte{PrefixFollowerPKIDToFollowedPKID})
	Followed2Follower(db, []byte{PrefixFollowedPKIDToFollowerPKID})
}

func Followed2Follower(db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()

			//val, _ := nodeIterator.Item().ValueCopy(nil)

			//data := bytes.NewReader(val)
			followed := key[1:34]
			follower := key[34:]
			HandleFollowed2Follower(followed, follower)
		}
		return nil
	})

}
func Follower2Followed(db *badger.DB, dbPrefix []byte) {

	db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		nodeIterator := txn.NewIterator(opts)
		defer nodeIterator.Close()
		prefix := dbPrefix

		for nodeIterator.Seek(prefix); nodeIterator.ValidForPrefix(prefix); nodeIterator.Next() {
			key := nodeIterator.Item().Key()

			//val, _ := nodeIterator.Item().ValueCopy(nil)

			//data := bytes.NewReader(val)
			follower := key[1:34]
			followed := key[34:]
			HandleFollower2Followed(follower, followed)
		}
		return nil
	})

}

type StakeEntryStats struct {
}

const HashSizeBytes = 32

type BlockHash [HashSizeBytes]byte
type StakeEntry struct {
}

type PostEntry struct {
	PostHash                 *BlockHash
	PosterPublicKey          []byte
	ParentStakeID            []byte
	Body                     []byte
	RecloutedPostHash        *BlockHash
	IsQuotedReclout          bool
	CreatorBasisPoints       uint64
	StakeMultipleBasisPoints uint64
	ConfirmationBlockHeight  uint32
	TimestampNanos           uint64
	IsHidden                 bool
	StakeEntry               *StakeEntry
	LikeCount                uint64
	RecloutCount             uint64
	QuoteRecloutCount        uint64
	DiamondCount             uint64
	stakeStats               *StakeEntryStats
	isDeleted                bool
	CommentCount             uint64
	IsPinned                 bool
	PostExtraData            map[string][]byte
}
