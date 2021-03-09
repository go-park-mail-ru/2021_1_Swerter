package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"my-motivation/internal"
)

func posts(w http.ResponseWriter, r *http.Request)  {
	post := internal.NewPost(1,"lol","hi","path")
	internal.Posts[1] = post
	fmt.Println(internal.Posts)
	jsonValue, _ := json.Marshal(internal.Posts)
	fmt.Println(jsonValue)
	w.Write([]byte(jsonValue))
}