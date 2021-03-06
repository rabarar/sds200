package server

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

func IsPrint(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func getXmlGLTFormatType(gltBuffer []byte) GltXmlType {
	switch {
	case strings.Contains(string(gltBuffer), "FL Index"):
		return GltXmlFL
	case strings.Contains(string(gltBuffer), "SYS Index"):
		return GltXmlSYS
	case strings.Contains(string(gltBuffer), "DEPT Index"):
		return GltXmlDEPT
	case strings.Contains(string(gltBuffer), "SITE"):
		return GltXmlSITE
	case strings.Contains(string(gltBuffer), "TRN_DISCOV"):
		return GltXmlTRN_DISCOV
	case strings.Contains(string(gltBuffer), "CNV_DISCOV"):
		return GltXmlCNV_DISCOV
	case strings.Contains(string(gltBuffer), "FTO"):
		return GltXmlFTO
	case strings.Contains(string(gltBuffer), "UREC_FOLDER"):
		return GltXmlUREC_FOLDER
	case strings.Contains(string(gltBuffer), "CS_BANK"):
		return GltXmlCSBANK
		/*
			GltXmlCFREQ
			GltXmlTGID
			GltXmlSFREQ
			GltXmlAFREQ
			GltXmlATGID
			GltXmlUREC
			GltXmlIREC_FILE
			GltXmlUREC_FILE
		*/
	default:
		return GltXmlUnknown
	}
}

/* func loadValidKeys() map[string]SDSKeyType {
	sdsKeys := make(map[string]SDSKeyType)

	sdsKeys["M"] = KEY_MENU
	sdsKeys["F"] = KEY_F
	sdsKeys["1"] = KEY_1
	sdsKeys["2"] = KEY_2
	sdsKeys["3"] = KEY_3
	sdsKeys["4"] = KEY_4
	sdsKeys["5"] = KEY_5
	sdsKeys["6"] = KEY_6
	sdsKeys["7"] = KEY_7
	sdsKeys["8"] = KEY_8
	sdsKeys["9"] = KEY_9
	sdsKeys["0"] = KEY_0
	sdsKeys["."] = KEY_DOT
	sdsKeys["E"] = KEY_ENTER
	sdsKeys[">"] = KEY_ROT_RIGHT
	sdsKeys["<"] = KEY_ROT_LEFT
	sdsKeys["^"] = KEY_ROT_PUSH
	sdsKeys["V"] = KEY_VOL_PUSH
	sdsKeys["Q"] = KEY_SQL_PUSH
	sdsKeys["Y"] = KEY_REPLAY
	sdsKeys["A"] = KEY_SYSTEM
	sdsKeys["B"] = KEY_DEPT
	sdsKeys["C"] = KEY_CHANNEL
	sdsKeys["Z"] = KEY_ZIP
	sdsKeys["T"] = KEY_SERV
	sdsKeys["R"] = KEY_RANGE
	return sdsKeys
}

func isValidKey(buffer []byte) bool {
	if len(buffer) != VALID_KEY_CMD_LENGTH {
		return false
	}

	_, ok := validKeys[string(buffer[4:5])]
	if !ok && !(string(buffer[6:10]) == KEY_PUSH_CMD || string(buffer[6:10]) == KEY_HOLD_CMD) {
		return false
	}
	return true
} */

func crlfStrip(msg []byte, flags uint) string {
	var replacer []string
	switch {
	case flags&LF == LF:
		replacer = append(replacer, "\r", "\\r")
		fallthrough
	case flags&NL == NL:
		replacer = append(replacer, "\n", "")
	}
	r := strings.NewReplacer(replacer...)
	return r.Replace(string(msg))
}

func IsValidXMLMessage(msgType string, buffer []byte) bool {
	var msg interface{}
	switch msgType {
	case "GSI", "PSI":
		msg = ScannerInfo{}
	case "MSI":
		msg = MsiInfo{}
	case "STS":
	case "GLG":
	case "GLT":
		switch getXmlGLTFormatType(buffer) {
		case GltXmlFL:
			msg = GltFLInfo{}
		case GltXmlSYS:
			msg = GltSysInfo{}
		case GltXmlDEPT:
			msg = GltDeptInfo{}
		case GltXmlSITE:
			msg = GltSiteInfo{}
		case GltXmlFTO:
			msg = GltFto{}
		case GltXmlCSBANK:
			msg = GltCSBank{}
		case GltXmlTRN_DISCOV:
			msg = GltTrnDiscovery{}
		case GltXmlCNV_DISCOV:
			msg = GltCnvDiscovery{}
		case GltXmlUREC_FOLDER:
			msg = GltUrecFolder{}
		default:
			log.Warnln("Unhandled GltXml Type", buffer)
		}
	}

	if len(buffer) < 11 {
		return false
	}

	clean := buffer[11:]

	if xmlErr := xml.Unmarshal(clean, &msg); xmlErr != nil {
		log.Debugln("Error when checking validity of XML message", xmlErr)
		return false
	}

	return true
}

// https://stackoverflow.com/a/52395088/486182
func ScanLinesWithCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\r'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func homepatrolChecksum(cmd string) string {
	cmd = fmt.Sprintf("%s\t", cmd)
	chk := int64(0)
	for _, char := range cmd {
		chk = chk + int64(char)
	}

	return strconv.FormatInt(chk, 10)
}

func homepatrolCommand(args []string) string {
	joined := strings.Join(args, "\t")
	return fmt.Sprintf("%s\t%s\r", joined, homepatrolChecksum(joined))
}
