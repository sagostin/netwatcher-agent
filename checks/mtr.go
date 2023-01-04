package checks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

type MtrResult struct {
	StartTimestamp time.Time `json:"start_timestamp"bson:"start_timestamp"`
	StopTimestamp  time.Time `json:"stop_timestamp"bson:"stop_timestamp"`
	Report         struct {
		Mtr struct {
			Src        string `json:"src"bson:"src"`
			Dst        string `json:"dst"bson:"dst"`
			Tos        int    `json:"tos"bson:"tos"`
			Tests      int    `json:"tests"bson:"tests"`
			Psize      string `json:"psize"bson:"psize"`
			Bitpattern string `json:"bitpattern"bson:"bitpattern"`
		} `json:"mtr"`
		Hubs []struct {
			Count int     `json:"count"bson:"count"`
			Host  string  `json:"host"bson:"host"`
			Loss  float64 `json:"Loss%"bson:"lossp"`
			Snt   int     `json:"Snt"bson:"snt"`
			Last  float64 `json:"Last"bson:"last"`
			Avg   float64 `json:"Avg"bson:"avg"`
			Best  float64 `json:"Best"bson:"best"`
			Wrst  float64 `json:"Wrst"bson:"wrst"`
			StDev float64 `json:"StDev"bson:"st_dev"`
		} `json:"hubs"bson:"hubs"`
	} `json:"report"bson:"report"`
}

/*type MtrMetrics struct {
	Address  string `json:"address"bson:"address"`
	FQDN     string `bson:"fqdn"json:"fqdn"`
	Sent     int    `json:"sent"bson:"sent"`
	Received int    `json:"received"bson:"received"`
	Last     string `bson:"last"json:"last"`
	Avg      string `bson:"avg"json:"avg"`
	Best     string `bson:"best"json:"best"`
	Worst    string `bson:"worst"json:"worst"`
}*/

func (r *MtrResult) Check(cd *CheckData) error {
	osDetect := runtime.GOOS
	r.StartTimestamp = time.Now()

	var cmd *exec.Cmd
	switch osDetect {
	case "windows":
		break
	case "darwin":
		// mtr needs to be installed manually currently
		args := []string{"-c", "./lib/mtr_darwin " + cd.Target + " --json"}
		cmd = exec.CommandContext(context.TODO(), "/bin/bash", args...)
		break
	case "linux":

		break
	default:
		log.Fatalf("Unknown OS")
		panic("TODO")
	}

	output, err := cmd.Output()
	fmt.Printf("%s\n", output)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &r.Report)
	if err != nil {
		return err
	}

	r.StopTimestamp = time.Now()
	cd.Result = r

	return nil
}

func mtrNumDashCheck(str string) int {
	if str == "-" {
		return 0
	}
	return ConvHandleStrInt(str)
}
