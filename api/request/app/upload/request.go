package upload

import (
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	AppFilePath        string
	AppVisible         bool
	EnableNotification bool
	ShortMessage       null.String
	DistributionKey    null.String
	DistributionName   null.String
	ReleaseNote        null.String
}

func boolToVisibility(b bool) string {
	if b {
		return "public"
	} else {
		return "private"
	}
}

type Key string

const (
	keyFile                  Key = "file"
	keyShortMessage          Key = "message"
	keyDistributionKey       Key = "distribution_key"
	keyDistributionName      Key = "distribution_name"
	keyReleaseNote           Key = "release_note"
	keyNotEnableNotification Key = "disable_notify"
	keyVisibility            Key = "visibility"
)

func (req Request) IoReaderMap() (*map[string]io.Reader, error) {
	fd, err := os.Open(req.AppFilePath)

	if err != nil {
		return nil, err
	}

	parts := map[Key]io.Reader{
		keyFile:                  fd,
		keyVisibility:            strings.NewReader(boolToVisibility(req.AppVisible)),
		keyNotEnableNotification: strings.NewReader(strconv.FormatBool(!req.EnableNotification)), // inverse!
	}

	if message := req.ShortMessage; message.Valid {
		parts[keyShortMessage] = strings.NewReader(message.String)
	}

	if key := req.DistributionKey; key.Valid {
		parts[keyDistributionKey] = strings.NewReader(key.String)
	}

	if name := req.DistributionName; name.Valid {
		parts[keyDistributionName] = strings.NewReader(name.String)
	}

	if note := req.ReleaseNote; note.Valid {
		parts[keyReleaseNote] = strings.NewReader(note.String)
	}

	out, err := util.StringifyKeys(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}
