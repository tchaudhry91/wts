package svc

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/tchaudhry91/wts-service/wts"
)

// ScriptRecord is a model for storing scripts with contexual information
type ScriptRecord struct {
	ID     string
	User   string
	Public bool
	Script *wts.Script
}

// Store is an interface that allows repository actions for Scripts
type Store interface {
	Put(ctx context.Context, user string, script *wts.Script) error
	Get(ctx context.Context, user string, name string) (*wts.Script, error)
}

// WTSFirestore provides a store implementation via Google Firestore
type WTSFirestore struct {
	Client *firestore.Client
}

// Put stores the script record in the database
func (s *WTSFirestore) Put(ctx context.Context, user string, script *wts.Script) error {
	doc := s.Client.Doc(fmt.Sprintf("%s/%s", user, script.Name))
	_, err := doc.Create(ctx, script)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the script record from the database
func (s *WTSFirestore) Get(ctx context.Context, user, name string) (script *wts.Script, err error) {
	doc := s.Client.Doc(fmt.Sprintf("%s/%s", user, name))
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return script, err
	}
	err = docSnap.DataTo(script)
	if err != nil {
		return script, err
	}
	return script, nil
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
