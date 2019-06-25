package ws

import (
	"net/http"

	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"

	"github.com/gorilla/websocket"

	"fileupload/file"

	"fileupload/db"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Handler handles websocket endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		print(err)
		http.Error(w, "Failed to establish websocket connection", http.StatusBadRequest)
		return
	}

	// successfully established websocket connection
	go _wsreader(ws)
}

func _wsreader(conn *websocket.Conn) {
	defer conn.Close()
	//start, stop := false, false // start -> true : read binary; start&stop = true: FIN
	var filename string
	var size int64
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			print(err)
			conn.WriteMessage(1, []byte("error reading message"))
			break
		}
		// no error

		// switch-case to handle message type
		switch t {
		case 1:
			// string
			// decode json
			var b map[string]interface{}
			err = json.Unmarshal(msg, &b)
			if err != nil {
				fmt.Println(err)
				conn.WriteMessage(1, []byte("error parsing message"))
				break
			}
			// read message
			fmt.Println(b) //debug
			filename = (b["filename"]).(string)
			size = int64((b["size"]).(float64))
			print(filename) //debug
			print(size)     //debug

			break
		case 2:
			// binary
			if len(filename) > 0 && size > 0 {
				// declared filename + size, ready to upload file
				// generate uuid
				uid, err := uuid.NewV4()
				if err != nil {
					fmt.Println(err)
					errMap := map[string]string{"result": "error"}
					errJSON, _ := json.Marshal(errMap)
					conn.WriteMessage(1, []byte(errJSON))
					return
				}

				// generate string uuid as file id
				fileID := uid.String()
				// store fileid in redis
				err = db.SetFile(&filename, &fileID)
				if err != nil {
					fmt.Println(err)
					errMap := map[string]string{"result": "error"}
					errJSON, _ := json.Marshal(errMap)
					conn.WriteMessage(1, []byte(errJSON))
					return
				}

				err = file.UploadS3(msg, &fileID, &size, http.DetectContentType(msg))
				if err != nil {
					fmt.Println(err)
					errMap := map[string]string{"result": "error"}
					errJSON, _ := json.Marshal(errMap)
					conn.WriteMessage(1, []byte(errJSON))
					return
				}
				// success
				successMap := map[string]string{"result": "success", "id": fileID}
				successJSON, _ := json.Marshal(successMap)
				conn.WriteMessage(1, []byte(successJSON))
				return
			}
		default:
			break
		}

		//conn.WriteMessage(t, msg)

	}
}
