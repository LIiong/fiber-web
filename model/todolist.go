package model

import (
	"context"
	"fiber-web/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

const tableName = "todolist"

type TodoList struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task       string             `json:"task,omitempty"`
	Status     bool               `json:"status,omitempty"`
	CreateTime time.Time          `json:"createTime,omitempty"`
}

//查询所有
func (t *TodoList) TodoLists() ([]*TodoList, error) {
	res, err := db.Mg.Db.Collection(tableName).Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error while fetching technologies:", err.Error())
		return nil, err
	}
	var tech []*TodoList
	err = res.All(context.TODO(), &tech)
	if err != nil {
		log.Println("Error while decoding technologies:", err.Error())
		return nil, err
	}
	return tech, nil
}

//新增
func (t *TodoList) Save() (interface{}, error) {
	t.Status = false
	t.CreateTime = time.Now()
	insertResult, err := db.Mg.Db.Collection(tableName).InsertOne(context.TODO(), t)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

//修改
func (t *TodoList) ChangeStatus(taskId string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	status := t.Status
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": status}}
	result, err := db.Mg.Db.Collection(tableName).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return result.UpsertedID, nil
}

//删除
func (t TodoList) Remove(taskId string) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(taskId)
	filter := bson.M{"_id": id}
	d, err := db.Mg.Db.Collection(tableName).DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return d.DeletedCount, nil
}
