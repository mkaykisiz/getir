package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const RecordsCollectionName = "records"

//RecordFilter is filter parameters for records.
type RecordFilter struct {
	StartDate string `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty" bson:"endDate,omitempty"`
	MinCount  int    `json:"minCount,omitempty" bson:"minCount,omitempty"`
	MaxCount  int    `json:"maxCount,omitempty" bson:"maxCount,omitempty"`
}

// Record is filtered record object.
type Record struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}

// RecordsResponse is filtered response.
type RecordsResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []Record `json:"records"`
}

// GetRecordsByFilter is filter func by filter parameter on mongo.
func GetRecordsByFilter(mongoDB *mongo.Database, filter RecordFilter) *RecordsResponse {
	mongoFilter, err := bson.Marshal(filter)
	if err != nil {
		return &RecordsResponse{Code: 2, Msg: "Wrong Filter Error"}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := mongoDB.Collection(RecordsCollectionName).Aggregate(ctx, mongoFilter)
	if err != nil {
		return &RecordsResponse{Code: 3, Msg: "Collection Error"}

	}
	defer cur.Close(ctx)

	rec := new(RecordsResponse)
	if err = cur.All(ctx, rec.Records); err != nil {
		return &RecordsResponse{Code: 4, Msg: "Cursor Error"}
	}
	return rec
}
