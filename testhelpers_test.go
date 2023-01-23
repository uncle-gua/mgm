package mgm_test

import (
	"testing"

	"github.com/uncle-gua/mgm"
	"github.com/uncle-gua/mgm/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDefConnection() {
	util.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("mongodb://root:12345@localhost:27017")),
	)
}

func resetCollection() {
	_, err := mgm.Coll(&Doc{}).DeleteMany(bson.M{})
	_, err2 := mgm.Coll(&Person{}).DeleteMany(bson.M{})

	util.PanicErr(err)
	util.PanicErr(err2)
}

func seed() {
	docs := []interface{}{
		NewDoc("Ali", 24),
		NewDoc("Mehran", 24),
		NewDoc("Reza", 26),
		NewDoc("Omid", 27),
	}
	_, err := mgm.Coll(&Doc{}).InsertMany(docs)

	util.PanicErr(err)
}

func findDoc(t *testing.T) *Doc {
	found := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(found).FindOne(bson.M{}).Decode(found))

	return found
}

type Doc struct {
	mgm.DefaultModel `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func NewDoc(name string, age int) *Doc {
	return &Doc{Name: name, Age: age}
}
