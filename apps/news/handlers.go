package news

import (
	"encoding/json"
	"fmt"
	i "my-motivation/internal"
	u "my-motivation/utils"
	"net/http"
)


func posts(w http.ResponseWriter, r *http.Request) {
	u.SetupCORS(&w)
	jsonValue, _ := json.Marshal(i.Posts)
	fmt.Println(jsonValue)
	w.Write([]byte(jsonValue))
}

func addPost(w http.ResponseWriter, r *http.Request) {
	u.SetupCORS(&w)

	//parse post info

	//make post newPost
	//id := 1
	//i.Posts[id] = newPost
	//

	fmt.Println(i.Posts)
	jsonValue, _ := json.Marshal(i.Posts)
	fmt.Println(jsonValue)
	w.Write([]byte(jsonValue))
}
