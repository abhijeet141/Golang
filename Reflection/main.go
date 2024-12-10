package main

type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}
type resWrapper struct {
	response
}

// var j1 = `{
// "item":"album",
// "album":{"title":"abc"}
// }`
// var j2 = `{
// "item":"song",
// "song":{"title":"acd","artist":"alan"}
// }`

// func main() {
// 	var resp1, resp2 resWrapper
// 	var err error
// 	if err = json.Unmarshal([]byte(j1), &resp1); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp1.response)
// 	if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(resp2.response)
// }
// func (r *resWrapper) UnmarshalJSON(b []byte) (err error) {
// 	var raw map[string]interface{} //empty interface can take value of any type
// 	err = json.Unmarshal(b, &r.response)
// 	fmt.Println(raw)
// 	return err
// }
