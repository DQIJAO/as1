package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

type Member struct {
	Title         string `json:"slug" bson:"Title"`
	Profile_image string `json:"Profile_image" bson:"image"`
}
type data struct {
	Members []Member `json:"members"`
}

func main() {

	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	resp, err := http.Get("https://gist.githubusercontent.com/DQIJAO/e14c64ea610688e70228a9fb8c649b2c/raw/6cccd444c1ef65411aa3662b112634996b837414/bnk48.json")

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	var data data

	err = json.Unmarshal(respByte, &data)

	if err != nil {
		fmt.Println(err)

		return
	}
	session.SetMode(mgo.Monotonic, true)
	d := session.DB("gonews").C("news")
	for i := 0; i < 29; i++ {
		doc := Member{
			//ID:    bson.NewObjectId(),
			Title:         data.Members[i].Title,
			Profile_image: data.Members[i].Profile_image,
		}
		err = d.Insert(doc)
		if err != nil {
			panic(err)
		}
	}
	//fmt.Println(data.Members[0])
}
