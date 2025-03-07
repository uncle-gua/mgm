package mgm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uncle-gua/mgm/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPrepareInvalidId(t *testing.T) {
	d := &Doc{}

	_, err := d.PrepareID("test")
	require.Error(t, err, "Expected get error on invalid id value")
}

func TestPrepareId(t *testing.T) {
	d := &Doc{}

	hexId := "5df7fb2b1fff9ee374b6bd2a"
	val, err := d.PrepareID(hexId)
	id, _ := primitive.ObjectIDFromHex(hexId)
	require.Equal(t, val.(primitive.ObjectID), id)
	util.AssertErrIsNil(t, err)
}
