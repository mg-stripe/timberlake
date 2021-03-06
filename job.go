package main

import "time"

type jobConf struct {
	Conf conf   `json:"conf"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Avoid adding additional exported fields
// as event streaming can overwhelm clients
type job struct {
	Details            jobDetail `json:"details"`
	Counters           []counter `json:"counters"`
	conf               conf
	Tasks              tasks `json:"tasks"`
	running            bool
	partial            bool
	updated            time.Time
	Cluster            string `json:"cluster"`
	ResourceManagerURL string `json:"resourceManagerUrl"`
	JobHistoryURL      string `json:"jobHistoryUrl"`

	// http://docs.cascading.org/cascading/1.2/javadoc/cascading/flow/Flow.html
	FlowID *string `json:"flowID"`
}

type jobDetail struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	User       string `json:"user"`
	State      string `json:"state"`
	StartTime  int64  `json:"startTime"`
	FinishTime int64  `json:"finishTime"`

	MapsTotal     int     `json:"mapsTotal"`
	MapProgress   float32 `json:"mapProgress"`
	MapsCompleted int     `json:"mapsCompleted"`
	MapsPending   int     `json:"mapsPending"`
	MapsRunning   int     `json:"mapsRunning"`
	MapsFailed    int     `json:"failedMapAttempts"`
	MapsKilled    int     `json:"killedMapAttempts"`
	MapsTotalTime int64   `json:"mapsTotalTime"`

	ReducesTotal     int     `json:"reducesTotal"`
	ReduceProgress   float32 `json:"reduceProgress"`
	ReducesCompleted int     `json:"reducesCompleted"`
	ReducesPending   int     `json:"reducesPending"`
	ReducesRunning   int     `json:"reducesRunning"`
	ReducesFailed    int     `json:"failedReduceAttempts"`
	ReducesKilled    int     `json:"killedReduceAttempts"`
	ReducesTotalTime int64   `json:"reducesTotalTime"`
}

type jobDetails []jobDetail

func (ds jobDetails) Len() int {
	return len(ds)
}

func (ds jobDetails) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}

func (ds jobDetails) Less(i, j int) bool {
	return ds[i].FinishTime < ds[j].FinishTime
}

type counter struct {
	Name   string `json:"name"`
	Total  int    `json:"total"`
	Map    int    `json:"map"`
	Reduce int    `json:"reduce"`
}

type appsDetailList struct {
	App []jobDetail `json:"app"`
}

type appsResp struct {
	Apps appsDetailList `json:"apps"`
}

type jobsDetailList struct {
	Job []jobDetail `json:"job"`
}

type jobsResp struct {
	Jobs jobsDetailList `json:"jobs"`
	Job  jobDetail      `json:"job"`
}

type confResp struct {
	Conf struct {
		Property []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"property"`
	} `json:"conf"`
}

type countersResp struct {
	JobCounters struct {
		CounterGroups []struct {
			Name     string `json:"counterGroupName"`
			Counters []struct {
				Name   string `json:"name"`
				Total  int    `json:"totalCounterValue"`
				Map    int    `json:"mapCounterValue"`
				Reduce int    `json:"reduceCounterValue"`
			} `json:"counter"`
		} `json:"counterGroup"`
	} `json:"jobCounters"`
}

type tasksResp struct {
	Tasks struct {
		Task []struct {
			StartTime  int64  `json:"startTime"`
			FinishTime int64  `json:"finishTime"`
			Type       string `json:"type"`
			State      string `json:"state"`
		} `json:"task"`
	} `json:"tasks"`
}

type clusterMetricsResp struct {
	Metrics struct {
		Containers int `json:"containersAllocated"`
	} `json:"clusterMetrics"`
}
