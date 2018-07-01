package upload

import (
	"bytes"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/util"
	"gopkg.in/guregu/null.v3"
	"io"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	AppFilePath          string
	AppVisibility        appVisibility
	SuppressNotification bool
	ShortMessage         null.String
	DistributionKey      null.String
	DistributionName     null.String
	ReleaseNote          null.String
}

type appVisibility string

const (
	publicApp  appVisibility = "public"
	privateApp appVisibility = "private"
)

func FormatVisibility(s string) appVisibility {
	switch s {
	case "public":
		return publicApp
	default:
		return privateApp
	}
}

func BoolToVisibility(b bool) appVisibility {
	if b {
		return publicApp
	} else {
		return privateApp
	}
}

type Key string

const (
	keyFile                 Key = "file"
	keyShortMessage         Key = "message"
	keyDistributionKey      Key = "distribution_key"
	keyDistributionName     Key = "distribution_name"
	keyReleaseNote          Key = "release_note"
	keySuppressNotification Key = "disable_notify"
	keyVisibility           Key = "visibility"
)

func (req Request) MultiPartForm(authority api.Authority) (data bytes.Buffer, contentType string, err error) {
	reqIOMap, err := req.ioReaderMap()

	if err != nil {
		return data, "", err
	}

	ioMap, err := authority.Merge(*reqIOMap)

	if err != nil {
		return data, "", err
	}

	data, contentType, err = util.Buffering(ioMap)

	if err != nil {
		return data, "", err
	}

	return data, contentType, nil
}

func (req Request) ioReaderMap() (*map[string]io.Reader, error) {
	fd, err := os.Open(req.AppFilePath)

	if err != nil {
		return nil, err
	}

	parts := map[Key]io.Reader{
		keyFile:                 fd,
		keyVisibility:           strings.NewReader(string(req.AppVisibility)),
		keySuppressNotification: strings.NewReader(strconv.FormatBool(req.SuppressNotification)),
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
