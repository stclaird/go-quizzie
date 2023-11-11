package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func Open(dbPath string) (*badger.DB, error) {
	//Open Badger Database
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		return nil, err
	}
	return db, err
}

func Close(db *badger.DB ) {
	defer db.Close()
}

func InsertOneItem(question Question, db *badger.DB) (err error) {
	//Insert a Question

	keyStr := question.Qid

	var b bytes.Buffer
    e := gob.NewEncoder(&b)

    if err := e.Encode(question); err != nil {
      log.Println(err)
	  return err
    }

    err = db.Update(func(txn *badger.Txn) error {
      err := txn.Set([]byte(keyStr), b.Bytes())
      return err
    })

    if err != nil {
      fmt.Printf("ERROR saving to badger db : %s\n", err)
    }

	return err
}

//retrieve all items in the database
func GetAllItems(db *badger.DB) ([]Question, error) {
  var questions []Question
  defer db.Close()

  err := db.View(func(txn *badger.Txn) error {
      opts := badger.DefaultIteratorOptions
      opts.PrefetchSize = 10
      it := txn.NewIterator(opts)
      defer it.Close()
      for it.Rewind(); it.Valid(); it.Next() {
        item := it.Item()
        err := item.Value(func(v []byte) error {
        var questionDecode Question
        d := gob.NewDecoder(bytes.NewReader(v))
        if err := d.Decode(&questionDecode); err != nil {
          panic(err)
        }
        questions = append(questions, questionDecode)
        return nil
        })
        if err != nil {
        return err
        }
      }
      return nil
      })
      return questions, err
}

//retrieve the questions with a key prefix starting with <prefix>
func GetItemsbyPrefix(prefix string, db *badger.DB) ([]QuestionNoAnswer, error) {
	var questions []QuestionNoAnswer
  defer db.Close()

  err :=  db.View(func(txn *badger.Txn) error {
      it := txn.NewIterator(badger.DefaultIteratorOptions)
      defer it.Close()
      prefix := []byte(prefix)
      for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
        item := it.Item()
        err := item.Value(func(v []byte) error {
          var questionDecode QuestionNoAnswer
          d := gob.NewDecoder(bytes.NewReader(v))
          if err := d.Decode(&questionDecode); err != nil {
            panic(err)
          }
          questions = append(questions, questionDecode)
          return nil
        })
        if err != nil {
          return err
        }
      }
      return nil
    })
    return questions, err
}

func GetItem(itemKey string, db *badger.DB) ( Question, error){
  var question Question
  defer db.Close()

  err :=  db.View(func(txn *badger.Txn) error {
    prefix := []byte(itemKey)
    item, err := txn.Get(prefix)
    if err != nil {
      return fmt.Errorf("getting value: %w", err)
     }

     valCopy, err := item.ValueCopy(nil)
     if err != nil {
      return fmt.Errorf("copying value: %w", err)
     }

     d := gob.NewDecoder(bytes.NewReader(valCopy))
     if err := d.Decode(&question); err != nil {
       panic(err)
     }

    return nil
  })

  return question, err
}