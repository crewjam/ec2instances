package ec2instances

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLive() ([]Instance, error) {
	resp, err := http.Get("http://www.ec2instances.info/instances.json")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	instances := []Instance{}
	if err := json.NewDecoder(resp.Body).Decode(&instances); err != nil {
		return nil, err
	}

	return instances, nil
}

func Get() []Instance {
	rv := []Instance{}
	if err := json.Unmarshal([]byte(data), &rv); err != nil {
		panic(err)
	}
	return rv
}

func GetInstance(instanceType string) *Instance {
	for _, i := range Get() {
		if i.InstanceType == instanceType {
			return &i
		}
	}
	return nil
}
