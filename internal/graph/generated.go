// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type AvailableVersion struct {
	Version string `json:"version"`
	Branch  string `json:"branch"`
	Date    string `json:"date"`
}

type DeferredUpdate struct {
	Channel string `json:"channel"`
	Version string `json:"version"`
	Logfile string `json:"logfile"`
}

type Project struct {
	Namespace string   `json:"namespace"`
	Locations []string `json:"locations"`
}

type StateVersion struct {
	License  string `json:"license"`
	Version  string `json:"version"`
	Branch   string `json:"branch"`
	Revision string `json:"revision"`
	Date     string `json:"date"`
}

type Version struct {
	State *StateVersion `json:"state"`
}
