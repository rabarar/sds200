package main

import "encoding/xml"

type ScannerInfo struct {
	XMLName     xml.Name `xml:"ScannerInfo"`
	Text        string   `xml:",chardata"`
	Mode        string   `xml:"Mode,attr"`
	VScreen     string   `xml:"V_Screen,attr"`
	MonitorList struct {
		Text      string `xml:",chardata"`
		Name      string `xml:"Name,attr"`
		Index     string `xml:"Index,attr"`
		ListType  string `xml:"ListType,attr"`
		QKey      string `xml:"Q_Key,attr"`
		NTag      string `xml:"N_Tag,attr"`
		DBCounter string `xml:"DB_Counter,attr"`
	} `xml:"MonitorList"`
	System struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"Name,attr"`
		Index      string `xml:"Index,attr"`
		Avoid      string `xml:"Avoid,attr"`
		SystemType string `xml:"SystemType,attr"`
		QKey       string `xml:"Q_Key,attr"`
		NTag       string `xml:"N_Tag,attr"`
		Hold       string `xml:"Hold,attr"`
	} `xml:"System"`
	Department struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"Name,attr"`
		Index string `xml:"Index,attr"`
		Avoid string `xml:"Avoid,attr"`
		QKey  string `xml:"Q_Key,attr"`
		Hold  string `xml:"Hold,attr"`
	} `xml:"Department"`
	TGID struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"Name,attr"`
		Index   string `xml:"Index,attr"`
		Avoid   string `xml:"Avoid,attr"`
		TGID    string `xml:"TGID,attr"`
		SetSlot string `xml:"SetSlot,attr"`
		RecSlot string `xml:"RecSlot,attr"`
		NTag    string `xml:"N_Tag,attr"`
		Hold    string `xml:"Hold,attr"`
		SvcType string `xml:"SvcType,attr"`
		PCh     string `xml:"P_Ch,attr"`
		LVL     string `xml:"LVL,attr"`
	} `xml:"TGID"`
	UnitID struct {
		Text string `xml:",chardata"`
		Name string `xml:"Name,attr"`
		UID  string `xml:"U_Id,attr"`
	} `xml:"UnitID"`
	Site struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"Name,attr"`
		Index string `xml:"Index,attr"`
		Avoid string `xml:"Avoid,attr"`
		QKey  string `xml:"Q_Key,attr"`
		Hold  string `xml:"Hold,attr"`
		Mod   string `xml:"Mod,attr"`
	} `xml:"Site"`
	SiteFrequency struct {
		Text string `xml:",chardata"`
		Freq string `xml:"Freq,attr"`
		IFX  string `xml:"IFX,attr"`
		SAS  string `xml:"SAS,attr"`
		SAD  string `xml:"SAD,attr"`
	} `xml:"SiteFrequency"`
	DualWatch struct {
		Text string `xml:",chardata"`
		PRI  string `xml:"PRI,attr"`
		CC   string `xml:"CC,attr"`
		WX   string `xml:"WX,attr"`
	} `xml:"DualWatch"`
	TrunkingDiscovery struct {
		Text       string `xml:",chardata"`
		SystemName string `xml:"SystemName,attr"`
		SiteName   string `xml:"SiteName,attr"`
		TGID       string `xml:"TGID,attr"`
		TgidName   string `xml:"TgidName,attr"`
		SAD        string `xml:"SAD,attr"`
		RecSlot    string `xml:"RecSlot,attr"`
		PastTime   string `xml:"PastTime,attr"`
		HitCount   string `xml:"HitCount,attr"`
		UID        string `xml:"U_Id,attr"`
	} `xml:"TrunkingDiscovery"`
	Property struct {
		Text      string `xml:",chardata"`
		F         string `xml:"F,attr"`
		VOL       string `xml:"VOL,attr"`
		SQL       string `xml:"SQL,attr"`
		Sig       string `xml:"Sig,attr"`
		Att       string `xml:"Att,attr"`
		Rec       string `xml:"Rec,attr"`
		KeyLock   string `xml:"KeyLock,attr"`
		P25Status string `xml:"P25Status,attr"`
		Mute      string `xml:"Mute,attr"`
		Backlight string `xml:"Backlight,attr"`
		ALed      string `xml:"A_Led,attr"`
		Dir       string `xml:"Dir,attr"`
		Rssi      string `xml:"Rssi,attr"`
	} `xml:"Property"`
	ViewDescription struct {
		Text      string `xml:",chardata"`
		PlainText []struct {
			Chardata string `xml:",chardata"`
			AttrText string `xml:"Text,attr"`
		} `xml:"PlainText"`
		PopupScreen struct {
			Chardata string `xml:",chardata"`
			AttrText string `xml:"Text,attr"`
			Button   struct {
				Chardata string `xml:",chardata"`
				AttrText string `xml:"Text,attr"`
				KeyCode  string `xml:"KeyCode,attr"`
			} `xml:"Button"`
		} `xml:"PopupScreen"`
	} `xml:"ViewDescription"`
}

type GltFLInfo struct {
	XMLName xml.Name `xml:"GLT"`
	Text    string   `xml:",chardata"`
	FL      []struct {
		Text    string `xml:",chardata"`
		Index   string `xml:"Index,attr"`
		Name    string `xml:"Name,attr"`
		Monitor string `xml:"Monitor,attr"`
		QKey    string `xml:"Q_Key,attr"`
		NTag    string `xml:"N_Tag,attr"`
	} `xml:"FL"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

type GltSysInfo struct {
	XMLName xml.Name `xml:"GLT"`
	Text    string   `xml:",chardata"`
	SYS     []struct {
		Text    string `xml:",chardata"`
		Index   string `xml:"Index,attr"`
		TrunkId string `xml:"TrunkId,attr"`
		Name    string `xml:"Name,attr"`
		Avoid   string `xml:"Avoid,attr"`
		Type    string `xml:"Type,attr"`
		QKey    string `xml:"Q_Key,attr"`
		NTag    string `xml:"N_Tag,attr"`
	} `xml:"SYS"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

type GltDeptInfo struct {
	XMLName xml.Name `xml:"GLT"`
	Text    string   `xml:",chardata"`
	DEPT    []struct {
		Text     string `xml:",chardata"`
		Index    string `xml:"Index,attr"`
		TGroupId string `xml:"TGroupId,attr"`
		Name     string `xml:"Name,attr"`
		Avoid    string `xml:"Avoid,attr"`
		QKey     string `xml:"Q_Key,attr"`
	} `xml:"DEPT"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

type GltSiteInfo struct {
	XMLName xml.Name `xml:"GLT"`
	Text    string   `xml:",chardata"`
	SITE    []struct {
		Text   string `xml:",chardata"`
		Index  string `xml:"Index,attr"`
		SiteId string `xml:"SiteId,attr"`
		Name   string `xml:"Name,attr"`
		Avoid  string `xml:"Avoid,attr"`
		QKey   string `xml:"Q_Key,attr"`
	} `xml:"SITE"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

// GLT CFREQ

// GLT TGID

// GLT AFREQ

// GLT ATGID

type GltFto struct {
	XMLName xml.Name `xml:"GLT"`
	Text    string   `xml:",chardata"`
	FTO     []struct {
		Text  string `xml:",chardata"`
		Index string `xml:"Index,attr"`
		Freq  string `xml:"Freq,attr"`
		Mod   string `xml:"Mod,attr"`
		Name  string `xml:"Name,attr"`
		ToneA string `xml:"ToneA,attr"`
		ToneB string `xml:"ToneB,attr"`
	} `xml:"FTO"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

type GltCSBank struct {
	XMLName xml.Name `xml:"GLT"`
	Text    string   `xml:",chardata"`
	CSBANK  []struct {
		Text  string `xml:",chardata"`
		Index string `xml:"Index,attr"`
		Name  string `xml:"Name,attr"`
		Lower string `xml:"Lower,attr"`
		Upper string `xml:"Upper,attr"`
		Mod   string `xml:"Mod,attr"`
		Step  string `xml:"Step,attr"`
	} `xml:"CS_BANK"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

// GLT,UREC

// GLT,IREC_FILE

// GLT,UREC_FILE,[folder_index]

type GltUrecFolder struct {
	XMLName    xml.Name `xml:"GLT"`
	Text       string   `xml:",chardata"`
	URECFOLDER []struct {
		Text  string `xml:",chardata"`
		Index string `xml:"Index,attr"`
		Name  string `xml:"Name,attr"`
	} `xml:"UREC_FOLDER"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

type GltTrnDiscovery struct {
	XMLName   xml.Name `xml:"GLT"`
	Text      string   `xml:",chardata"`
	TRNDISCOV []struct {
		Text         string `xml:",chardata"`
		Name         string `xml:"Name,attr"`
		Delay        string `xml:"Delay,attr"`
		Logging      string `xml:"Logging,attr"`
		Duration     string `xml:"Duration,attr"`
		CompareDB    string `xml:"CompareDB,attr"`
		SystemName   string `xml:"SystemName,attr"`
		SystemType   string `xml:"SystemType,attr"`
		SiteName     string `xml:"SiteName,attr"`
		TimeOutTimer string `xml:"TimeOutTimer,attr"`
		AutoStore    string `xml:"AutoStore,attr"`
	} `xml:"TRN_DISCOV"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

// GLT,CNV_DISCOV
type GltCnvDiscovery struct {
	XMLName   xml.Name `xml:"GLT"`
	Text      string   `xml:",chardata"`
	CNVDISCOV []struct {
		Text         string `xml:",chardata"`
		Name         string `xml:"Name,attr"`
		Lower        string `xml:"Lower,attr"`
		Upper        string `xml:"Upper,attr"`
		Mod          string `xml:"Mod,attr"`
		Step         string `xml:"Step,attr"`
		Delay        string `xml:"Delay,attr"`
		Logging      string `xml:"Logging,attr"`
		CompareDB    string `xml:"CompareDB,attr"`
		Duration     string `xml:"Duration,attr"`
		TimeOutTimer string `xml:"TimeOutTimer,attr"`
		AutoStore    string `xml:"AutoStore,attr"`
	} `xml:"CNV_DISCOV"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}

type MsiInfo struct {
	XMLName  xml.Name `xml:"MSI"`
	Text     string   `xml:",chardata"`
	Name     string   `xml:"Name,attr"`
	Index    string   `xml:"Index,attr"`
	MenuType string   `xml:"MenuType,attr"`
	Value    string   `xml:"Value,attr"`
	Selected string   `xml:"Selected,attr"`
	MenuItem []struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"Name,attr"`
		Index string `xml:"Index,attr"`
	} `xml:"MenuItem"`
	Footer struct {
		Text string `xml:",chardata"`
		No   string `xml:"No,attr"`
		EOT  string `xml:"EOT,attr"`
	} `xml:"Footer"`
}