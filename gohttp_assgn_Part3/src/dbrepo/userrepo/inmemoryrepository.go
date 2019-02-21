package userrepo


import (
	"domain"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"github.com/satori/go.uuid"
	//"fmt"
	"log"
)

//MongoRepository mongodb repo
type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "restaurant"

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

//Find a Restaurant
func (r *MongoRepository) Get(id string) (*domain.Restaurant, error) {
	
	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	uid,err := uuid.FromString(id) 
	if err !=nil{
		log.Fatal(err)
	}
	err = coll.Find(bson.M{"_id": uid}).One(&result)
	switch err {
	case nil:
		return &result, nil
	case mgo.ErrNotFound:
		return nil, domain.ErrNotFound
	default:
		return nil, err
	}
}

//Store a Restaurant record has an issue with id while inserting values
func (r *MongoRepository) Store(b *domain.Restaurant) (uuid.UUID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	_, err := coll.UpsertId(b.DBID, b)
	if err != nil {
		return b.DBID, err
	}
	return b.DBID, nil
}

//Get all records from the collection
func (r *MongoRepository) GetAll()([] *domain.Restaurant,error){
	session :=r.mongoSession.Clone()
	defer session.Close()
	var values []*domain.Restaurant
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(nil).All(&values)
	return values,err
}

//Find the all records having given name as a pattern in the name field 
func (r *MongoRepository)FindByName(name string)([] *domain.Restaurant,error){
	session := r.mongoSession.Clone()
	defer session.Close()
	var values [] *domain.Restaurant
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"name":bson.RegEx{Pattern:name,Options:"i"}}).All(&values)
	return values,err
	
}

//Delete all records with the given id
func (r MongoRepository)Delete(id string)(error){
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	uid,err := uuid.FromString(id) 
	if err !=nil{
		log.Fatal(err)
	}
	err = coll.Remove(bson.M{"_id":uid})
	return err	
}

//FindByTypeOfFood retrieves the list of restaurnts with given foodtype else error 
func (r *MongoRepository)FindByTypeOfFood(foodType string)([] *domain.Restaurant,error){
	session := r.mongoSession.Clone()
	defer session.Close()
	var values [] *domain.Restaurant
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"typeOfFood":foodType}).All(&values)
	return values,err
	
}


//FindByTypeOfPostCode returns the list of restaurants with given postcode else error
func (r *MongoRepository)FindByTypeOfPostCode(postCode string)([] *domain.Restaurant,error){
	session := r.mongoSession.Clone()
	defer session.Close()
	var values [] *domain.Restaurant
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode":postCode}).All(&values)
	return values,err
	
}

//Search searchs the pattern in all string fields of domain Restaurant 
func (r *MongoRepository)Search(query string) ([]*domain.Restaurant, error){
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	var values []*domain.Restaurant
	err := coll.Find(bson.M{ "$or":[]bson.M{
			{"name": bson.RegEx{Pattern :query,Options:"i" }},
			{"address": bson.RegEx{Pattern :query,Options:"i" }},
			{"addressLine2": bson.RegEx{Pattern :query,Options:"i" }},
			{"url": bson.RegEx{Pattern :query,Options:"i" }},
			{"outcode": bson.RegEx{Pattern :query,Options:"i" }},
			{"postcode": bson.RegEx{Pattern :query,Options:"i" }},
			{"typeOfFood": bson.RegEx{Pattern :query,Options:"i" }}}}).All(&values)
			
	return values,err	
	
}
