package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"encoding/json"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	DB       = "xingcuntian"
	username = "root"
	password = "root"
)

func main() {
	c := influxDBclient()
	//createMetrics(c)
	readWithLimit(c, 1000000)
	meanCPUUsage(c, "us-west")
	countRegion(c, "us-west")

}

func influxDBclient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return c
}

func createMetrics(clnt client.Client) {
	bathCount := 1000000
	rand.Seed(42)
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DB,
		Precision: "s",
	})
	for i := 0; i < bathCount; i++ {
		regions := []string{"us-west", "us-central", "us-north", "us-east"}
		tags := map[string]string{
			"host":   fmt.Sprintf("192.168.%d.%d", rand.Intn(100), rand.Intn(100)),
			"region": regions[rand.Intn(len(regions))],
		}
		value := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"cpu_usage": value,
		}
		pt, err := client.NewPoint("cpu", tags, fields, time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		bp.AddPoint(pt)
	}
	err := clnt.Write(bp)
	if err != nil {
		log.Fatalln(err)
	}
}

func queryDB(clnt client.Client, command string) (res []client.Result, err error) {
	q := client.Query{
		Command:  command,
		Database: DB,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func readWithLimit(clnt client.Client, limit int) {
	q := fmt.Sprintf("select*from %s Limit %d", "cpu", limit)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	for i, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatalln(err)
		}
		val, err := row[1].(json.Number).Float64()
		fmt.Printf("[%2d] %s: %f\n", i, t.Format(time.Stamp), val)
	}
}

func meanCPUUsage(clnt client.Client, region string) {
	q := fmt.Sprintf("select mean(%s) from %s where region = '%s'", "cpu_usage", "cpu", region)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln(err)
	}
	value, err := res[0].Series[0].Values[0][1].(json.Number).Float64()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Mean value of cpu_usage for region '%s':%f\n", region, value)
}

func countRegion(clnt client.Client, region string) {
	q := fmt.Sprintf("select count(%s) from %s where region = '%s'", "cpu_usage", "cpu", region)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln(err)
	}
	count := res[0].Series[0].Values[0][1]
	fmt.Printf("Found a total of %v records for region '%s'\n", count, region)
}
