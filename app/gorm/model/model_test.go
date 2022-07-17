package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

func Test_Str(t *testing.T) {
	str := "not found: ☀️   🌡️+47°F 🌬️↖2mph\n"
	fmt.Println(str)
	fmt.Println(strings.Trim(str, "\n"))
}

func Test_JsonConv(t *testing.T) {
	model := Model{
		ID:        1,
		CreatedAt: UnixTime(time.Now()),
		UpdatedAt: UnixTime(time.Now()),
	}

	bytes, _ := json.Marshal(&model)
	fmt.Println(string(bytes))
	modelCopy := Model{}
	json.Unmarshal(bytes, &modelCopy)
	fmt.Println(modelCopy.ID, modelCopy.CreatedAt, modelCopy.UpdatedAt)
}

func Test_Time(t *testing.T) {
	fmt.Println("Time with MicroSeconds: ", time.Now().Format("2006-01-02 15:04:05.000"))

}
