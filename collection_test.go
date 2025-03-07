package mgm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uncle-gua/mgm"
	"github.com/uncle-gua/mgm/builder"
	"github.com/uncle-gua/mgm/internal/util"
	"github.com/uncle-gua/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindByIdWithInvalidId(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	require.NotNil(t, mgm.Coll(&Doc{}).FindByID("invalid id", &Doc{}))
}

func TestFindFirst(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	d := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(&Doc{}).First(bson.M{}, d))

	require.NotEqual(t, primitive.ObjectID{}, d.ID)
}

func TestCollection_Create(t *testing.T) {
	setupDefConnection()
	resetCollection()

	doc := NewDoc("Ali", 24)

	util.AssertErrIsNil(t, mgm.Coll(doc).Create(doc))

	// Inserted model's id should not be nil:
	require.NotNil(t, doc.ID, "Expected document having id after insertion, got nil")

	// We should have one document in database that is equal to this doc:
	foundDoc := &Doc{}
	util.AssertErrIsNil(t, mgm.Coll(doc).FindByID(doc.ID, foundDoc))

	require.Equal(t, doc.Name, foundDoc.Name, "expected inserted and retrieved docs be equal, got %v and %v", doc.Name, foundDoc.Name)
	require.Equal(t, doc.Age, foundDoc.Age, "expected inserted and retrieved docs be equal, got %v and %v", doc.Age, foundDoc.Age)
}

func TestCollection_Update(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	found.Name = found.Name + "_extra_val"
	found.Age = found.Age + 4

	util.AssertErrIsNil(t, mgm.Coll(found).Update(found))

	// Find that doc again:
	newFound := findDoc(t)

	if found.ID != newFound.ID {
		panic("two fond document dont have same id!")
	}
	require.Equal(t, found.Name, newFound.Name)
	require.Equal(t, found.Age, newFound.Age)
}

func TestCollection_Delete(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	found := findDoc(t)

	util.AssertErrIsNil(t, mgm.Coll(found).Delete(found))

	// Find that doc again:
	newFound := findDoc(t)

	require.NotEqual(t, found.ID, newFound.ID)
}

func TestCollection_SimpleFind(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	expectedResult := []Doc{}
	gotResult := []Doc{}

	filter := bson.M{"age": bson.M{operator.Gt: 24}}
	err := mgm.Coll(&Doc{}).SimpleFind(&gotResult, filter)

	util.AssertErrIsNil(t, err)

	// Create same aggregation by raw methods
	cur, err := mgm.Coll(&Doc{}).Find(filter)
	util.AssertErrIsNil(t, err)

	ctx, cancel := mgm.Ctx()
	defer cancel()

	util.AssertErrIsNil(t, cur.All(ctx, &expectedResult))

	require.Equal(t, len(expectedResult), len(gotResult))

	// We should have same documents
	for i, expectedDoc := range expectedResult {
		if expectedDoc != gotResult[i] {
			t.Errorf("Expected %v, got %v", expectedDoc, gotResult[i])
		}
	}
}

func TestCollection_SimpleAggregateFirst(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	expectedResult := []Doc{}
	gotResult := Doc{}

	// We dont want to change document.
	group := builder.Group("$_id", nil)

	found, err := mgm.Coll(&Doc{}).SimpleAggregateFirst(&gotResult, group)

	assert.True(t, found)
	util.AssertErrIsNil(t, err)

	// Create same aggregation by raw methods
	cur, err := mgm.Coll(&Doc{}).Aggregate(bson.A{builder.S(group)}, nil)
	util.AssertErrIsNil(t, err)

	ctx, cancel := mgm.Ctx()
	defer cancel()

	util.AssertErrIsNil(t, cur.All(ctx, &expectedResult))
	assert.Equal(t, expectedResult[0], gotResult)
}

func TestCollection_SimpleAggregateFirstFalse(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	var gotResult *Doc
	match := bson.M{operator.Match: bson.M{"user_id": "unknown"}}
	found, err := mgm.Coll(&Doc{}).SimpleAggregateFirst(gotResult, match)

	assert.False(t, found)
	util.AssertErrIsNil(t, err)
	assert.Nil(t, gotResult)
}

func TestCollection_SimpleAggregate(t *testing.T) {
	setupDefConnection()
	resetCollection()
	seed()

	expectedResult := []Doc{}
	gotResult := []Doc{}

	// We dont want to change document.
	group := builder.Group("$_id", nil)

	project := bson.M{operator.Project: bson.M{"age": 0}}

	err := mgm.Coll(&Doc{}).SimpleAggregate(&gotResult, group, project)

	util.AssertErrIsNil(t, err)

	// Create same aggregation by raw methods
	cur, err := mgm.Coll(&Doc{}).Aggregate(bson.A{builder.S(group), project}, nil)
	util.AssertErrIsNil(t, err)

	ctx, cancel := mgm.Ctx()
	defer cancel()

	util.AssertErrIsNil(t, cur.All(ctx, &expectedResult))

	require.Equal(t, len(expectedResult), len(gotResult))

	// We should have same documents
	for i, expectedDoc := range expectedResult {
		if expectedDoc != gotResult[i] {
			t.Errorf("Expected %v, got %v", expectedDoc, gotResult[i])
		}
	}
}
