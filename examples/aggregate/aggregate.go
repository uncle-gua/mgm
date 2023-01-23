package aggregate

import (
	"github.com/uncle-uga/mgm"
	"github.com/uncle-uga/mgm/builder"
	"github.com/uncle-uga/mgm/field"
	"go.mongodb.org/mongo-driver/bson"
)

func seed() {
	author := newAuthor("Mehran")
	_ = mgm.Coll(author).Create(author)

	book := newBook("Test", 124, author.ID)
	_ = mgm.Coll(book).Create(book)

}

func delSeededData() {
	_, _ = mgm.Coll(&book{}).DeleteMany(bson.M{})
	_, _ = mgm.Coll(&author{}).DeleteMany(bson.M{})
}

func lookup() error {
	seed()

	defer delSeededData()

	// Author model's collection
	authorColl := mgm.Coll(&author{})

	pipeline := bson.A{
		builder.S(builder.Lookup(authorColl.Name(), "author_id", field.ID, "author")),
	}

	ctx := mgm.Ctx()
	cur, err := mgm.Coll(&book{}).AggregateWithCtx(ctx, pipeline)

	if err != nil {
		return err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return err
		}

		// do something with result....
		//fmt.Printf("%+v\n", result)
	}

	return nil
}
