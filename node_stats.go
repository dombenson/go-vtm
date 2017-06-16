package stingray

import (
	"encoding/json"
	"net/http"
)

// A NodeStats is a Stingray node.
type NodeStats struct {
	jsonStatsResource `json:"-"`
	NodeStatistics    `json:"statistics"`
}

type NodeStatistics struct {
	BytesIn            int64  `json:"bytes_from_node"`
	BytesInHigh        uint32 `json:"bytes_from_node_hi"`
	BytesInLow         uint32 `json:"bytes_from_node_lo"`
	BytesOut           int64  `json:"bytes_to_node"`
	BytesOutHigh       uint32 `json:"bytes_to_node_hi"`
	BytesOutLow        uint32 `json:"bytes_to_node_lo"`
	CurrentConnections int    `json:"current_conn"`
	CurrentRequests    int    `json:"current_requests"`
	Errors             int    `json:"errors"`
	Failures           int    `json:"failures"`
	NewConnections     int    `json:"new_conn"`
	PooledConnections  int    `json:"pooled_conn"`
	Port               int    `json:"port"`
	MaxResponseTime    int    `json:"response_max"`
	MinResponseTime    int    `json:"response_min"`
	MeanResponseTime   int    `json:"response_mean"`
	State              string `json:"state"`
	TotalConnections   int    `json:"total_conn"`
}

func (r *NodeStats) endpoint() string {
	return "nodes/node"
}

//String will return back the json as a string
func (r *NodeStats) String() string {
	b := r.Bytes()
	return string(b)
}

//Bytes will return back just the bytes
func (r *NodeStats) Bytes() []byte {
	b, _ := jsonMarshal(r)
	return b
}

func (r *NodeStats) decode(data []byte) error {
	return json.Unmarshal(data, &r)
}

func NewNodeStats(name string) *NodeStats {
	r := new(NodeStats)
	r.setName(name)
	return r
}

func (c *Client) GetNodeStats(name string) (*NodeStats, *http.Response, error) {
	r := NewNodeStats(name)

	resp, err := c.Get(r)
	if err != nil {
		return nil, resp, err
	}

	r.BytesOut = int64(r.BytesOutHigh)<<32 + int64(r.BytesOutLow)
	r.BytesIn = int64(r.BytesInHigh)<<32 + int64(r.BytesInLow)

	return r, resp, nil
}
