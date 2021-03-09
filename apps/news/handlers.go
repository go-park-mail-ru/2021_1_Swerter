package news

import (
	u "my-motivation/utils"
	"encoding/json"
	"fmt"
	"my-motivation/internal"
	"net/http"
)

func posts(w http.ResponseWriter, r *http.Request) {
	u.SetupCORS(&w)
	post := internal.NewPost(1, "lol", "hi", "path")
	internal.Posts[1] = post
	fmt.Println(internal.Posts)
	jsonValue, _ := json.Marshal(internal.Posts)
	fmt.Println(jsonValue)
	w.Write([]byte(jsonValue))
}
