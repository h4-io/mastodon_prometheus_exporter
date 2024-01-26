package mastodon

import (
	"encoding/json"
)

func GetGenericStats(instance string) (GenericStats, error) {
	var stats GenericStats
	data, err := GenericHttpGet(instance + "/api/v1/instance")
	if err != nil {
		return stats, err
	}
	err = json.Unmarshal(data, &stats)
	if err != nil {
		return stats, err
	}
	return stats, nil
}

func GetActivities(instance string) ([]Activities, error) {
	var activities []Activities
	data, err := GenericHttpGet(instance + "/api/v1/instance/activity")
	if err != nil {
		return activities, err
	}
	err = json.Unmarshal(data, &activities)
	if err != nil {
		return activities, err
	}
	return activities, nil
}

func GetNodeInfo(instance string) (NodeInfo, error) {
	var nodeInfo NodeInfo
	data, err := GenericHttpGet(instance + "/nodeinfo/2.0")
	if err != nil {
		return nodeInfo, err
	}
	err = json.Unmarshal(data, &nodeInfo)
	if err != nil {
		return nodeInfo, err
	}
	return nodeInfo, nil
}
