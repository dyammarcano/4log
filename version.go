// this file version.go was generated by build.go

package log4go

func GetVersion() *Version {
	return &Version{
		Version:    "v0.0.1-dev",
		CommitHash: "749b236ca43aa74d063564269771d3dfe06ce7ad",
		Date:       "2023-11-27T03:23:08Z",
		UUID:       "94e0ef6e-8936-4da3-8bae-9c299e3f5bce",
	}
}

type Version struct {
	Version    string `json:"version"`
	CommitHash string `json:"commitHash"`
	Date       string `json:"date"`
	UUID       string `json:"id"`
}
