package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//InitHook initializes the webhook collection.
func (db *DBInfo) InitHook() {
	session, err := mgo.Dial(db.DBurl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"timestamp"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DBname).C(db.HookCollection).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

//AddHook adds webhook to collection
func (db *DBInfo) AddHook(s Webhook) error {
	session, err := mgo.Dial(db.DBurl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DBname).C(db.HookCollection).Insert(s)

	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

//GetHook returns one webhook
func (db *DBInfo) GetHook(keyID int64) (Webhook, error) {
	session, err := mgo.Dial(db.DBurl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	track := Webhook{}

	err = session.DB(db.DBname).C(db.HookCollection).Find(bson.M{"timestamp": keyID}).One(&track)

	return track, err
}

//GetAllHooks returns slice with all Hooks
func (db *DBInfo) GetAllHooks() []Webhook {
	session, err := mgo.Dial(db.DBurl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var all []Webhook

	err = session.DB(db.DBname).C(db.HookCollection).Find(bson.M{}).All(&all)
	if err != nil {
		return []Webhook{}
	}

	return all
}

//Delete a webhook from the database
func (db *DBInfo) DeleteHook(id int64) error {
    session, err := mgo.Dial(db.DBurl)
    if err != nil {
        panic(err)
    }
    defer session.Close()

    err = session.DB(db.DBname).C(db.HookCollection).Remove(bson.M{"timestamp": id})

    return err
}
