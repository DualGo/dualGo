package engine

import "log"

func RaisedError(err error){
	if(err != nil){
		log.Fatal(err)
	}
}
