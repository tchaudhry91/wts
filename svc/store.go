package svc

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/tchaudhry91/wts-service/wts"
)

// ScriptRecord is a model for storing scripts with contexual information
type ScriptRecord struct {
	User   string      `json:"user,omitempty"`
	Public bool        `json:"public,omitempty"`
	Script *wts.Script `json:"script,omitempty"`
}

var dbError = errors.New("Failed on a database operation")
var notFound = errors.New("Record not found")

// Store is an interface that allows repository actions for Scripts
type Store interface {
	Put(ctx context.Context, user string, record *ScriptRecord) error
	Get(ctx context.Context, user string, name string) (*ScriptRecord, error)
}

// BoltStore is a single file database meant for single service setups
type BoltStore struct {
	DB *bolt.DB
}

// Put stores the record in the given bolt db
func (s *BoltStore) Put(ctx context.Context, user string, record *ScriptRecord) error {
	// Start a write transaction
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		buc, err := tx.CreateBucketIfNotExists([]byte("scripts"))
		if err != nil {
			return err
		}

		// Store the JSON marshalled record
		key := "user" + "/" + record.Script.Name
		val, err := json.Marshal(record)
		if err != nil {
			return err
		}
		if err := buc.Put([]byte(key), val); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return errors.Wrap(dbError, err.Error())
	}
	return nil
}

// Get retreives a record from the given bolt db
func (s *BoltStore) Get(ctx context.Context, user string, name string) (*ScriptRecord, error) {
	record := &ScriptRecord{}
	// Read the value for the given key
	key := user + "/" + name
	if err := s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("scripts"))
		if b == nil {
			return notFound
		}
		data := b.Get([]byte(key))
		if data == nil {
			return notFound
		}
		if errIn := json.Unmarshal(data, record); errIn != nil {
			return errors.Wrap(dbError, errIn.Error())
		}
		return nil
	}); err != nil {
		return record, err
	}
	return record, nil
}

// Returns a new Bolt based Store implementation for WTS
func NewBoltStore(fname string) (*BoltStore, error) {
	store := &BoltStore{}
	db, err := bolt.Open(fname, 0600, nil)
	if err != nil {
		return store, err
	}
	store.DB = db
	return store, nil
}

// WTSFirestore provides a store implementation via Google Firestore
type WTSFirestore struct {
	Client *firestore.Client
}

// Put stores the script record in the database
func (s *WTSFirestore) Put(ctx context.Context, user string, record ScriptRecord) error {
	doc := s.Client.Doc(fmt.Sprintf("%s/%s", user, record.Script.Name))
	_, err := doc.Create(ctx, record)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the script record from the database
func (s *WTSFirestore) Get(ctx context.Context, user, name string) (record *ScriptRecord, err error) {
	doc := s.Client.Doc(fmt.Sprintf("%s/%s", user, name))
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return record, err
	}
	err = docSnap.DataTo(record)
	if err != nil {
		return record, err
	}
	return record, nil
}

// New WTSFirestore constructs a new firestore backed Store
func NewWTSFirestore(projectID string) (*WTSFirestore, error) {
	fs := WTSFirestore{}
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return &fs, err
	}
	fs.Client = client
	return &fs, err
}
