package protocol

import (
	"github.com/hunchly/funchly/onionscan/config"
	"github.com/hunchly/funchly/onionscan/report"
)

type Scanner interface {
	ScanProtocol(hiddenService string, onionscanConfig *config.OnionScanConfig, report *report.OnionScanReport)
}
