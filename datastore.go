package main

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/willf/bloom"

	"gopkg.in/redis.v4"
)

// Bloom Filter
var filter = bloom.NewWithEstimates(c.BloomSize, 1e-3) // Configurable in environment var.

// Record is the main struct that is passed from applications to AuthTables as JSON.
// Applications send us these, and AuthTables responds with `OK`s or `BAD`.
type Record struct {
	UID string `json:"uid"`
	IP  string `json:"ip"`
	Mid string `json:"mid"`
}

// Marshaler returns the input data jsonified.
func (r Record) Marshaler() []byte {

	json, err := json.Marshal(r)
	if err != nil {
		log.Error("Issue marshal'ing json")
	}
	fmt.Println(string(json))

	return json
}

// RecordHashes is a struct ready for use in the bloom filter or redis.
type RecordHashes struct {
	uid    []byte
	uidMID []byte
	uidIP  []byte
	uidALL []byte
	ipMID  []byte
	midIP  []byte
}

// Take us online to Redis
var client = redis.NewClient(&redis.Options{
	Addr:     c.Host + ":" + c.Port,
	Password: c.Password, // no password set
	DB:       0,          // use default DB
})
