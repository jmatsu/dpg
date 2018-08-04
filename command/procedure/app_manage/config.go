package app_manage

type config struct {
	nameAssignStrategy nameAssignStrategy `json:nameAssignStrategy`
	nameExtractRegexp  *string            `json:nameExtractRegexp`
}

type nameAssignStrategy string

const (
	branchName nameAssignStrategy = "branchName"
	envVar                        = "envVar"
	option                        = "option"
)
