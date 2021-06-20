package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/dgraph-io/badger/v3"
)

var Lookup = map[string]string{}

func HandleFollower2Followed(db *badger.DB, follower, followed []byte) {
	pub58er := base58.Encode(follower)
	if Lookup[pub58er] == "" {
		Lookup[pub58er] = LookupUsername(db, follower)
	}
	pub58ed := base58.Encode(followed)
	if Lookup[pub58ed] == "" {
		Lookup[pub58ed] = LookupUsername(db, followed)
	}
	fmt.Println("HandleFollower2Followed", Lookup[pub58er], Lookup[pub58ed])
}
func HandleFollowed2Follower(db *badger.DB, followed, follower []byte) {
	pub58ed := base58.Encode(followed)
	if Lookup[pub58ed] == "" {
		Lookup[pub58ed] = LookupUsername(db, followed)
	}
	pub58er := base58.Encode(follower)
	if Lookup[pub58er] == "" {
		Lookup[pub58er] = LookupUsername(db, follower)
	}
	fmt.Println("HandleFollowed2Follower", Lookup[pub58ed], Lookup[pub58er])
}

func LookupUsername(db *badger.DB, pkid []byte) string {

	username := ""
	err := db.View(func(txn *badger.Txn) error {

		key := append([]byte{23}, pkid...)
		profile, err := txn.Get(key)

		if err != nil {
			return err
		}
		profile.Value(func(valBytes []byte) error {
			profile := &ProfileEntry{}
			gob.NewDecoder(bytes.NewReader(valBytes)).Decode(profile)
			username = string(profile.Username)
			return nil
		})

		return nil
	})

	if err == nil {
		return username
	}

	return "404"
}
