package db

import (
	"context"
	"fmt"

	"github.com/dimishpatriot/rest-art-of-development/internal/logging"
	"github.com/dimishpatriot/rest-art-of-development/internal/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

// Create new user and return UUID
func (d *db) Create(ctx context.Context, user *user.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("can't create new user in DB: %s", err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("cant get ObjectID for new user")
	}
	uuid := oid.Hex()
	d.logger.Infof("[OK] user with UUID<%s> was created", uuid)

	return uuid, nil
}

// Delete user by UUID
func (d *db) Delete(ctx context.Context, uuid string) error {
	oid, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return fmt.Errorf("can't get ObjectID from UUID=%s", err)
	}

	result, err := d.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return fmt.Errorf("can't delete user from DB by UUID<%s>: %s", uuid, err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("can't find user UUID<%s> to delete", uuid)
	}
	d.logger.Infof("[OK] user with UUID<%s> was deleted", uuid)

	return nil
}

// FindOne
func (d *db) FindOne(ctx context.Context, uuid string) (user *user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return nil, fmt.Errorf("can't get ObjectID from UUID=%s: %s", uuid, err)
	}

	if err = d.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("can't find user UUID<%s>: %w", uuid, err)
		}
		return nil, err
	}
	d.logger.Infof("[OK] user UUID<%s> was found", user.ID)

	return user, nil
}

// Update
func (d *db) Update(ctx context.Context, user *user.User) error {
	oid, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	update := bson.M{
		"username":      user.Username,
		"email":         user.Email,
		"password_hash": user.PasswordHash,
	}

	result, err := d.collection.UpdateByID(ctx, oid, update)
	if err != nil {
		return fmt.Errorf("can't update user UUID<%s>: %s", user.ID, err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("can't find user UUID<%s> to update", user.ID)
	}
	d.logger.Infof("[OK] user UUID<%s> was updated", user.ID)

	return nil
}

func NewCollection(database *mongo.Database, collectionName string) user.Storage {
	return &db{
		collection: database.Collection(collectionName),
		logger:     logging.GetLogger(),
	}
}
