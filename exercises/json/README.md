# JSON

`encoding/json` turns Go values into JSON (`json.Marshal`) and back (`json.Unmarshal`). Only exported fields are written. Struct tags change how a field is written: `json:"name"` renames it, `json:"email,omitempty"` leaves it out when it is empty, `json:"-"` never writes it, and `json:"at,omitzero"` (Go 1.24) leaves out any zero value, which `omitempty` does not do for structs such as `time.Time`. Go 1.27 adds `encoding/json/v2`, a new package with stricter defaults, in the `go127` topic.

* [Package encoding/json](https://pkg.go.dev/encoding/json)
* [JSON and Go](https://go.dev/blog/json)
* [Go 1.24 Release Notes: omitzero](https://go.dev/doc/go1.24#encodingjsonpkgencodingjson)
