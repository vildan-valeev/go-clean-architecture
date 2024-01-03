package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MGInsertItem(ctx context.Context, db MongoQuery, user user) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"messenger_id", user.MessengerID}, {"messenger", user.Messenger}}

	update := bson.D{
		{
			"$set",
			bson.D{
				//{"_id", user.ID},
				{"messenger", user.Messenger},
				{"messenger_id", user.MessengerID},
				{"first_name", user.FirstName},
				{"last_name", user.LastName},
				{"user_name", user.UserName},
				{"created", user.Created},
				{"updated", user.Updated},
				{"is_admin", user.IsAdmin},
				{"is_delete", user.IsDelete},
				{"is_bot", user.IsBot},
			},
		},
	}

	result, err := db.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	log.Info().Msgf("Document User inserted %s - %v\n", user.ID, result.UpsertedCount)
	log.Info().Msgf("Document User updatet %s - %v\n", user.ID, result.ModifiedCount)

	return nil
}

func MGInsertCategory(ctx context.Context, db MongoQuery, order domain.Order) error {
	model, err := toModel(order)
	if err != nil {
		return err
	}

	id, err := db.InsertOne(ctx, model)
	log.Info().Msgf("Document Order inserted with ID: %s\n", id)
	if err != nil {
		return err
	}

	return nil
}

type Filter struct {
	MessengerID int64        `bson:"messenger_id,omitempty"`
	Messenger   pb.Messenger `bson:"messenger,omitempty"`
}

func MGGetItem(ctx context.Context, db MongoQuery, id int64, messenger pb.Messenger) (user, error) {
	var u user

	filter := bson.D{{"messenger_id", id}, {"messenger", messenger}}

	err := db.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return user{}, err
	}

	return u, nil
}

func MGGetAllItems(ctx context.Context, db MongoQuery, id primitive.ObjectID) (orders, error) {
	var ors orders

	filter := bson.D{{"user_id", id}, {"is_delete", false}}
	cursor, err := db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &ors)
	if err != nil {
		return nil, err
	}

	return ors, nil
}

func MGUpdateItem(ctx context.Context, db MongoQuery, order order) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"_id", order.ID}}
	update := bson.D{
		{
			"$set",
			bson.D{
				//{"_id", user.ID},
				{"point_a", order.PointA},
				{"user_id", order.UserID},
				{"point_b", order.PointA},
				{"date_start", order.DateStart},
				{"description", order.Description},
				{"amount", order.Amount},
				{"tg_msg_id", order.TgMsgID},
				{"tg_status", order.TgStatus},
				{"vk_msg_id", order.VkMsgID},
				{"vk_status", order.VkStatus},
				{"published", order.Published},
				{"created", order.Created},
				{"updated", order.Updated},
				{"is_delete", order.IsDelete},
			},
		},
	}
	result, err := db.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	log.Info().Msgf("Document Order inserted %s - %v\n", order.ID, result.UpsertedCount)
	log.Info().Msgf("Document Order updated %s - %v\n", order.ID, result.ModifiedCount)

	return nil
}

func MGDeleteItem(ctx context.Context, db MongoQuery, orderID primitive.ObjectID) error {
	opts := options.Update().SetUpsert(false)
	filter := bson.D{{"_id", orderID}}
	update := bson.D{
		{
			"$set",
			bson.D{
				{"is_delete", true},
			},
		},
	}
	result, err := db.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	log.Info().Msgf("Delete Document Order inserted %s - %v\n", orderID, result.UpsertedCount)
	log.Info().Msgf("Delete Document Order updated %s - %v\n", orderID, result.ModifiedCount)

	return nil
}
