package db_test

import (
	"context"
	"testing"

	"github.com/dimishpatriot/rest-art-of-development/internal/config"
	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
	"github.com/dimishpatriot/rest-art-of-development/internal/user"
	"github.com/dimishpatriot/rest-art-of-development/internal/user/db"
)

var (
	logger     *logging.Logger
	cfg        *config.Config
	ctx        context.Context
	collection user.Storage
)

func TestMain(m *testing.M) {
	logger = logging.GetLogger()
	ctx = context.Background()

	cfg = &config.Config{}
	cfg.Storage = config.Storage{
		Host:       "127.0.0.1",
		Port:       "27017",
		Database:   "rest-art",
		Collection: "users",
		Username:   "",
		Password:   "",
	}

	database := db.Connect(ctx, cfg, logger)
	collection = db.NewCollection(database, cfg.Storage.Collection)

	m.Run()
}

func Test_db_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *user.User
	}
	tests := []struct {
		name       string
		collection user.Storage
		args       args
		wantErr    bool
	}{
		{
			name: "correct user data", collection: collection,
			args:    args{ctx: ctx, user: &user.User{Username: "New User", PasswordHash: "123qwe", Email: "email@example.com"}},
			wantErr: false,
		},
		{
			name: "empty user data", collection: collection,
			args:    args{ctx: ctx, user: &user.User{Username: "", PasswordHash: "", Email: ""}},
			wantErr: true,
		},
		{
			name: "empty user name", collection: collection,
			args: args{ctx: ctx, user: &user.User{Username: "", PasswordHash: "123qwe", Email: "email@example.com"}}, wantErr: true,
		},
		{
			name: "empty user password hash", collection: collection,
			args: args{ctx: ctx, user: &user.User{Username: "New User", PasswordHash: "", Email: "email@example.com"}}, wantErr: true,
		},
		{
			name: "empty user email", collection: collection,
			args: args{ctx: ctx, user: &user.User{Username: "New User", PasswordHash: "123qwe", Email: ""}}, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.collection.Create(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("db.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// // TODO: test
// 	// 1
// 	database := db.Connect(ctx, cfg, logger)
// 	collection := db.NewCollection(database, cfg.Storage.Collection)

// 	// 2
// 	uuid, err := collection.Create(ctx, &user.User{
// 		Username:     "Pop",
// 		PasswordHash: "1234",
// 		Email:        "example@example.com",
// 	})
// 	if err != nil {
// 		logger.Fatal(err)
// 	}

// 	// 3
// 	if _, err = collection.FindOne(ctx, uuid); err != nil {
// 		logger.Fatal(err)
// 	}

// 	// 4
// 	newUserData := &user.User{
// 		ID:           uuid,
// 		Username:     "UPD_USERNAME",
// 		PasswordHash: "UPD_HASH",
// 		Email:        "upd@upd.me",
// 	}
// 	if err = collection.Update(ctx, newUserData); err != nil {
// 		logger.Fatal(err)
// 	}
// 	if _, err := collection.FindOne(ctx, uuid); err != nil {
// 		logger.Fatal(err)
// 	}

// 	// 5
// 	if err = collection.Delete(ctx, uuid); err != nil {
// 		logger.Fatal(err)
// 	}
// 	if _, err = collection.FindOne(ctx, uuid); err == nil {
// 		logger.Fatal(err)
// 	}
// 	// ---
