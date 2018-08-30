package main

import (
	"fmt"
	"github.com/dreamldx/bb/dal"
	"github.com/dreamldx/bb/http"
	"github.com/globalsign/mgo"
)

func main() {
	dal.LoadArticle()

	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		fmt.Printf("Connnect to mongo error due to %v\n", err)
	}

	defer session.Close()



	http.StartHTTPServer(8321)
}
