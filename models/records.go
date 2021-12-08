package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const dateFormat = "2006-01-02"

//RecordFilter is filter parameters for records.
type RecordFilter struct {
	StartDate string `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty" bson:"endDate,omitempty"`
	MinCount  int    `json:"minCount,omitempty" bson:"minCount,omitempty"`
	MaxCount  int    `json:"maxCount,omitempty" bson:"maxCount,omitempty"`
}

// Record is filtered record object.
type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

// RecordsResponse is filtered response.
type RecordsResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []Record `json:"records"`
}

// GetRecordsByFilter is filter func by filter parameter on mongo.
func GetRecordsByFilter(mongoDB *mongo.Collection, filter RecordFilter) *RecordsResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	startDate, _ := time.Parse(dateFormat, filter.StartDate)
	endDate, _ := time.Parse(dateFormat, filter.EndDate)

	filterQuery := []bson.M{
		{"$match": bson.M{"createdAt": bson.M{"$gt": startDate, "$lt": endDate}}},
		{"$project": bson.M{"_id": 1, "key": 1, "value": 1, "createdAt": 1, "totalCount": bson.M{"$sum": "$counts"}}},
		{"$match": bson.M{"totalCount": bson.M{"$gt": filter.MinCount, "$lt": filter.MaxCount}}},
	}

	cur, err := mongoDB.Aggregate(ctx, filterQuery)
	if err != nil {
		return &RecordsResponse{Code: 3, Msg: err.Error()}
	}
	defer cur.Close(ctx)

	rec := new(RecordsResponse)
	if err = cur.All(ctx, &rec.Records); err != nil {
		return &RecordsResponse{Code: 4, Msg: err.Error()}
	}

	return rec
}
