package upload

import (
	"fmt"
	"github.com/jmatsu/dpg/util"
	"github.com/sirupsen/logrus"
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

	out, err := util.StringifyKeysOfReaderMap(parts)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (req Request) Verify() error {
	if req.AppFilePath == "" {
		return fmt.Errorf("app file path must be present")
	}

	if f, err := os.Stat(req.AppFilePath); err != nil {
		return err
	} else if f.Size() == 0 {
		return fmt.Errorf("%s must not be an empty", req.AppFilePath)
	}

	if req.DistributionKey.Valid && req.DistributionKey.String == "" {
		return fmt.Errorf("distribution key must not be empty if specified")
	}

	if req.DistributionKey.Valid && req.DistributionName.String != "" {
		req.DistributionName = null.StringFromPtr(nil)
		logrus.Warnf("ignored the name of the distribution because the both of name and key were present")
	}

	return nil
}
