package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// NullObjectID -
type NullObjectID struct {
	*primitive.ObjectID
}

// UnmarshalText for QueryParser
func (p *NullObjectID) UnmarshalText(text []byte) error {
	oid, err := primitive.ObjectIDFromHex(string(text))
	if err != nil {
		return err
	}
	p.ObjectID = &oid
	return nil
}
