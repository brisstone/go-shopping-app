// create a reusable struct to describe our API’s response. 

package responses

// creates a UserResponse struct with Status, Message, and 
// Data property to represent the API response type.

// PS: json:"status", json:"message", and json:"data" are known as struct tags. 
// Struct tags allow us to attach meta-information to corresponding struct properties. 
// In other words, we use them to reformat the JSON response returned by the API.
type UserResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// func handlePage(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	var message Message
// 	err := json.NewDecoder(request.Body).Decode(&message)
// 	if err != nil {
// 			return
// 	}
// 	err = json.NewEncoder(writer).Encode(message)
// 	if err != nil {
// 			return
// 	}
// }