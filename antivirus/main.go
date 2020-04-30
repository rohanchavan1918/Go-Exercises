package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Antivirus implementation in Go using Virus total APi
func clearScreen() {
	// CLear the screen
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func banner() {
	fmt.Println("[+] ANTIVIRUS [+]")
	fmt.Println("USAGE :- av.exe -T dir/file/full ")
}

// Md5Hash returns the md5 of the given file
func Md5Hash(filePath string) (string, error) {
	// fmt.Println("calculating MD5 for ->", filePath)
	var returnMD5String string
	file, _ := os.Open(filePath)
	// if err != nil {
	// 	return returnMD5String, err
	// }
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

// SingleFileMode is for single file
func SingleFileMode(upath string) bool {
	var result bool
	mdhash, _ := Md5Hash(upath)
	if DetectionEngines(mdhash) == false {
		// Means it is not malicious
		result = false
	} else {
		// Means malicious
		result = true
	}
	return result
}

// DetectionEngines returns a boolean value wthr the file is malicious or not.
// True for Malicious / false for safe
func DetectionEngines(md5hash string) bool {
	// Struct which is made out of response
	var malicious bool = false
	type APIinfo struct {
		Data struct {
			Attributes struct {
				Authentihash string `json:"authentihash"`
				CreationDate int    `json:"creation_date"`
				Exiftool     struct {
					CharacterSet             string `json:"CharacterSet"`
					CodeSize                 int    `json:"CodeSize"`
					Comments                 string `json:"Comments"`
					CompanyName              string `json:"CompanyName"`
					EntryPoint               string `json:"EntryPoint"`
					FileDescription          string `json:"FileDescription"`
					FileFlagsMask            string `json:"FileFlagsMask"`
					FileOS                   string `json:"FileOS"`
					FileSubtype              string `json:"FileSubtype"`
					FileType                 string `json:"FileType"`
					FileTypeExtension        string `json:"FileTypeExtension"`
					FileVersion              string `json:"FileVersion"`
					FileVersionNumber        string `json:"FileVersionNumber"`
					ImageFileCharacteristics string `json:"ImageFileCharacteristics"`
					ImageVersion             string `json:"ImageVersion"`
					InitializedDataSize      int    `json:"InitializedDataSize"`
					InternalName             string `json:"InternalName"`
					LanguageCode             string `json:"LanguageCode"`
					LegalCopyright           string `json:"LegalCopyright"`
					LinkerVersion            string `json:"LinkerVersion"`
					MIMEType                 string `json:"MIMEType"`
					MachineType              string `json:"MachineType"`
					OSVersion                string `json:"OSVersion"`
					ObjectFileType           string `json:"ObjectFileType"`
					OriginalFileName         string `json:"OriginalFileName"`
					PEType                   string `json:"PEType"`
					ProductName              string `json:"ProductName"`
					ProductVersion           string `json:"ProductVersion"`
					ProductVersionNumber     string `json:"ProductVersionNumber"`
					Subsystem                string `json:"Subsystem"`
					SubsystemVersion         string `json:"SubsystemVersion"`
					TimeStamp                string `json:"TimeStamp"`
					UninitializedDataSize    int    `json:"UninitializedDataSize"`
				} `json:"exiftool"`
				FirstSubmissionDate int `json:"first_submission_date"`
				LastAnalysisDate    int `json:"last_analysis_date"`
				LastAnalysisResults struct {
					ALYac struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"ALYac"`
					APEX struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"APEX"`
					AVG struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"AVG"`
					Acronis struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Acronis"`
					AdAware struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Ad-Aware"`
					AegisLab struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"AegisLab"`
					AhnLabV3 struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"AhnLab-V3"`
					Alibaba struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Alibaba"`
					AntiyAVL struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Antiy-AVL"`
					Arcabit struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Arcabit"`
					Avast struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Avast"`
					AvastMobile struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Avast-Mobile"`
					Avira struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Avira"`
					Baidu struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Baidu"`
					BitDefender struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"BitDefender"`
					BitDefenderTheta struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"BitDefenderTheta"`
					Bkav struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Bkav"`
					CATQuickHeal struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"CAT-QuickHeal"`
					CMC struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"CMC"`
					ClamAV struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"ClamAV"`
					Comodo struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Comodo"`
					CrowdStrike struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"CrowdStrike"`
					Cybereason struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Cybereason"`
					Cylance struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Cylance"`
					Cyren struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Cyren"`
					DrWeb struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"DrWeb"`
					ESETNOD32 struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"ESET-NOD32"`
					Emsisoft struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Emsisoft"`
					Endgame struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Endgame"`
					FProt struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"F-Prot"`
					FSecure struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"F-Secure"`
					FireEye struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"FireEye"`
					Fortinet struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Fortinet"`
					GData struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"GData"`
					Ikarus struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Ikarus"`
					Invincea struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Invincea"`
					Jiangmin struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Jiangmin"`
					K7AntiVirus struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"K7AntiVirus"`
					K7GW struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"K7GW"`
					Kaspersky struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Kaspersky"`
					Kingsoft struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Kingsoft"`
					MAX struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"MAX"`
					Malwarebytes struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Malwarebytes"`
					MaxSecure struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"MaxSecure"`
					McAfee struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"McAfee"`
					McAfeeGWEdition struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"McAfee-GW-Edition"`
					MicroWorldEScan struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"MicroWorld-eScan"`
					Microsoft struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Microsoft"`
					NANOAntivirus struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"NANO-Antivirus"`
					Paloalto struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Paloalto"`
					Panda struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Panda"`
					Qihoo360 struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Qihoo-360"`
					Rising struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Rising"`
					SUPERAntiSpyware struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"SUPERAntiSpyware"`
					Sangfor struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Sangfor"`
					SentinelOne struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"SentinelOne"`
					Sophos struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Sophos"`
					Symantec struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Symantec"`
					SymantecMobileInsight struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"SymantecMobileInsight"`
					TACHYON struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"TACHYON"`
					Tencent struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Tencent"`
					Trapmine struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Trapmine"`
					TrendMicro struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"TrendMicro"`
					TrendMicroHouseCall struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"TrendMicro-HouseCall"`
					VBA32 struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"VBA32"`
					VIPRE struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"VIPRE"`
					ViRobot struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"ViRobot"`
					Webroot struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Webroot"`
					Yandex struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Yandex"`
					Zillya struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Zillya"`
					ZoneAlarm struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"ZoneAlarm"`
					Zoner struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion string      `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"Zoner"`
					EGambit struct {
						Category      string      `json:"category"`
						EngineName    string      `json:"engine_name"`
						EngineUpdate  string      `json:"engine_update"`
						EngineVersion interface{} `json:"engine_version"`
						Method        string      `json:"method"`
						Result        interface{} `json:"result"`
					} `json:"eGambit"`
				} `json:"last_analysis_results"`
				LastAnalysisStats struct {
					ConfirmedTimeout int `json:"confirmed-timeout"`
					Failure          int `json:"failure"`
					Harmless         int `json:"harmless"`
					Malicious        int `json:"malicious"`
					Suspicious       int `json:"suspicious"`
					Timeout          int `json:"timeout"`
					TypeUnsupported  int `json:"type-unsupported"`
					Undetected       int `json:"undetected"`
				} `json:"last_analysis_stats"`
				LastModificationDate int      `json:"last_modification_date"`
				LastSubmissionDate   int      `json:"last_submission_date"`
				Magic                string   `json:"magic"`
				Md5                  string   `json:"md5"`
				MeaningfulName       string   `json:"meaningful_name"`
				Names                []string `json:"names"`
				PeInfo               struct {
					EntryPoint int    `json:"entry_point"`
					Imphash    string `json:"imphash"`
					Imports    struct {
						ADVAPI32Dll []string `json:"ADVAPI32.dll"`
						COMCTL32Dll []string `json:"COMCTL32.dll"`
						GDI32Dll    []string `json:"GDI32.dll"`
						KERNEL32Dll []string `json:"KERNEL32.dll"`
						MFC42UDll   []string `json:"MFC42u.dll"`
						OLEAUT32Dll []string `json:"OLEAUT32.dll"`
						SHELL32Dll  []string `json:"SHELL32.dll"`
						USER32Dll   []string `json:"USER32.dll"`
						WINMMDll    []string `json:"WINMM.dll"`
						WS232Dll    []string `json:"WS2_32.dll"`
						Comdlg32Dll []string `json:"comdlg32.dll"`
						Msvcp60Dll  []string `json:"msvcp60.dll"`
						MsvcrtDll   []string `json:"msvcrt.dll"`
						Ole32Dll    []string `json:"ole32.dll"`
					} `json:"imports"`
					MachineType int `json:"machine_type"`
					Overlay     struct {
						Chi2     float64 `json:"chi2"`
						Entropy  float64 `json:"entropy"`
						Filetype string  `json:"filetype"`
						Md5      string  `json:"md5"`
						Offset   int     `json:"offset"`
						Size     int     `json:"size"`
					} `json:"overlay"`
					ResourceDetails []struct {
						Chi2     float64 `json:"chi2"`
						Entropy  float64 `json:"entropy"`
						Filetype string  `json:"filetype"`
						Lang     string  `json:"lang"`
						Sha256   string  `json:"sha256"`
						Type     string  `json:"type"`
					} `json:"resource_details"`
					ResourceLangs struct {
						ENGLISHUS int `json:"ENGLISH US"`
						GERMAN    int `json:"GERMAN"`
					} `json:"resource_langs"`
					ResourceTypes struct {
						RTBITMAP    int `json:"RT_BITMAP"`
						RTDIALOG    int `json:"RT_DIALOG"`
						RTGROUPICON int `json:"RT_GROUP_ICON"`
						RTICON      int `json:"RT_ICON"`
						RTMANIFEST  int `json:"RT_MANIFEST"`
						RTMENU      int `json:"RT_MENU"`
						RTSTRING    int `json:"RT_STRING"`
						RTVERSION   int `json:"RT_VERSION"`
						Struct240   int `json:"Struct(240)"`
						TEXTINCLUDE int `json:"TEXTINCLUDE"`
					} `json:"resource_types"`
					Sections []struct {
						Entropy        float64 `json:"entropy"`
						Md5            string  `json:"md5"`
						Name           string  `json:"name"`
						RawSize        int     `json:"raw_size"`
						VirtualAddress int     `json:"virtual_address"`
						VirtualSize    int     `json:"virtual_size"`
					} `json:"sections"`
					Timestamp int `json:"timestamp"`
				} `json:"pe_info"`
				Reputation    int    `json:"reputation"`
				Sha1          string `json:"sha1"`
				Sha256        string `json:"sha256"`
				SignatureInfo struct {
					Comments              string `json:"comments"`
					Copyright             string `json:"copyright"`
					CounterSigners        string `json:"counter signers"`
					CounterSignersDetails []struct {
						Algorithm    string `json:"algorithm"`
						CertIssuer   string `json:"cert issuer"`
						Name         string `json:"name"`
						SerialNumber string `json:"serial number"`
						Status       string `json:"status"`
						Thumbprint   string `json:"thumbprint"`
						ValidFrom    string `json:"valid from"`
						ValidTo      string `json:"valid to"`
						ValidUsage   string `json:"valid usage"`
					} `json:"counter signers details"`
					Description    string `json:"description"`
					FileVersion    string `json:"file version"`
					InternalName   string `json:"internal name"`
					OriginalName   string `json:"original name"`
					Product        string `json:"product"`
					Signers        string `json:"signers"`
					SignersDetails []struct {
						Algorithm    string `json:"algorithm"`
						CertIssuer   string `json:"cert issuer"`
						Name         string `json:"name"`
						SerialNumber string `json:"serial number"`
						Status       string `json:"status"`
						Thumbprint   string `json:"thumbprint"`
						ValidFrom    string `json:"valid from"`
						ValidTo      string `json:"valid to"`
						ValidUsage   string `json:"valid usage"`
					} `json:"signers details"`
					SigningDate string `json:"signing date"`
					Verified    string `json:"verified"`
					X509        []struct {
						Algorithm    string `json:"algorithm"`
						CertIssuer   string `json:"cert issuer"`
						Name         string `json:"name"`
						SerialNumber string `json:"serial number"`
						Thumbprint   string `json:"thumbprint"`
						ValidFrom    string `json:"valid from"`
						ValidTo      string `json:"valid to"`
						ValidUsage   string `json:"valid_usage"`
					} `json:"x509"`
				} `json:"signature_info"`
				Size           int      `json:"size"`
				Ssdeep         string   `json:"ssdeep"`
				Tags           []string `json:"tags"`
				TimesSubmitted int      `json:"times_submitted"`
				TotalVotes     struct {
					Harmless  int `json:"harmless"`
					Malicious int `json:"malicious"`
				} `json:"total_votes"`
				Trid []struct {
					FileType    string  `json:"file_type"`
					Probability float64 `json:"probability"`
				} `json:"trid"`
				TypeDescription string `json:"type_description"`
				TypeTag         string `json:"type_tag"`
				UniqueSources   int    `json:"unique_sources"`
				Vhash           string `json:"vhash"`
			} `json:"attributes"`
			ID    string `json:"id"`
			Links struct {
				Self string `json:"self"`
			} `json:"links"`
			Type string `json:"type"`
		} `json:"data"`
	}

	const baseurl string = "https://www.virustotal.com/api/v3/files/"
	const APIKey string = "4abb5d292096c73bc65be078d80601ada83196afd29ffcf3e179bda3da3ea27b"
	finalurl := baseurl + md5hash
	// fmt.Println(finalurl)

	// Create a HTTP CLient
	client := &http.Client{}
	req, err := http.NewRequest("GET", finalurl, nil)
	req.Header.Set("x-apikey", APIKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyjson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := APIinfo{}
	err = json.Unmarshal(bodyjson, &jsonParser)

	// THis works incomplete ...
	avresult := make(map[string]string)
	avresult["ALYac"] = jsonParser.Data.Attributes.LastAnalysisResults.ALYac.Category
	avresult["APEX"] = jsonParser.Data.Attributes.LastAnalysisResults.APEX.Category
	avresult["AVG"] = jsonParser.Data.Attributes.LastAnalysisResults.AVG.Category
	avresult["Acronis"] = jsonParser.Data.Attributes.LastAnalysisResults.Acronis.Category
	// avresult["Ad-Aware"] = jsonParser.Data.Attributes.LastAnalysisResults.Ad - Aware.Category
	avresult["AegisLab"] = jsonParser.Data.Attributes.LastAnalysisResults.AegisLab.Category
	// avresult["AhnLab-V3"] = jsonParser.Data.Attributes.LastAnalysisResults.AhnLab - V3.Category
	avresult["Alibaba"] = jsonParser.Data.Attributes.LastAnalysisResults.Alibaba.Category
	// avresult["Antiy-AVL"] = jsonParser.Data.Attributes.LastAnalysisResults.Antiy - AVL.Category
	avresult["Arcabit"] = jsonParser.Data.Attributes.LastAnalysisResults.Arcabit.Category
	avresult["Avast"] = jsonParser.Data.Attributes.LastAnalysisResults.Avast.Category
	// avresult["Avast-Mobile"] = jsonParser.Data.Attributes.LastAnalysisResults.Avast - Mobile.Category
	avresult["Avira"] = jsonParser.Data.Attributes.LastAnalysisResults.Avira.Category
	avresult["Baidu"] = jsonParser.Data.Attributes.LastAnalysisResults.Baidu.Category
	avresult["BitDefender"] = jsonParser.Data.Attributes.LastAnalysisResults.BitDefender.Category
	avresult["BitDefenderTheta"] = jsonParser.Data.Attributes.LastAnalysisResults.BitDefenderTheta.Category
	avresult["Bkav"] = jsonParser.Data.Attributes.LastAnalysisResults.Bkav.Category
	// avresult["CAT-QuickHeal"] = jsonParser.Data.Attributes.LastAnalysisResults.CAT - QuickHeal.Category
	avresult["CMC"] = jsonParser.Data.Attributes.LastAnalysisResults.CMC.Category
	avresult["ClamAV"] = jsonParser.Data.Attributes.LastAnalysisResults.ClamAV.Category
	avresult["Comodo"] = jsonParser.Data.Attributes.LastAnalysisResults.Comodo.Category
	avresult["CrowdStrike"] = jsonParser.Data.Attributes.LastAnalysisResults.CrowdStrike.Category
	avresult["Cybereason"] = jsonParser.Data.Attributes.LastAnalysisResults.Cybereason.Category
	avresult["Cylance"] = jsonParser.Data.Attributes.LastAnalysisResults.Cylance.Category
	avresult["Cyren"] = jsonParser.Data.Attributes.LastAnalysisResults.Cyren.Category
	avresult["DrWeb"] = jsonParser.Data.Attributes.LastAnalysisResults.DrWeb.Category
	// avresult["ESET-NOD32"] = jsonParser.Data.Attributes.LastAnalysisResults.ESET - NOD32.Category
	avresult["Emsisoft"] = jsonParser.Data.Attributes.LastAnalysisResults.Emsisoft.Category
	avresult["Endgame"] = jsonParser.Data.Attributes.LastAnalysisResults.Endgame.Category
	// avresult["F-Prot"] = jsonParser.Data.Attributes.LastAnalysisResults.F - Prot.Category
	// avresult["F-Secure"] = jsonParser.Data.Attributes.LastAnalysisResults.F - Secure.Category
	avresult["FireEye"] = jsonParser.Data.Attributes.LastAnalysisResults.FireEye.Category
	avresult["Fortinet"] = jsonParser.Data.Attributes.LastAnalysisResults.Fortinet.Category
	avresult["GData"] = jsonParser.Data.Attributes.LastAnalysisResults.GData.Category
	avresult["Ikarus"] = jsonParser.Data.Attributes.LastAnalysisResults.Ikarus.Category
	avresult["Invincea"] = jsonParser.Data.Attributes.LastAnalysisResults.Invincea.Category
	avresult["Jiangmin"] = jsonParser.Data.Attributes.LastAnalysisResults.Jiangmin.Category
	avresult["K7AntiVirus"] = jsonParser.Data.Attributes.LastAnalysisResults.K7AntiVirus.Category
	avresult["K7GW"] = jsonParser.Data.Attributes.LastAnalysisResults.K7GW.Category
	avresult["Kaspersky"] = jsonParser.Data.Attributes.LastAnalysisResults.Kaspersky.Category
	avresult["Kingsoft"] = jsonParser.Data.Attributes.LastAnalysisResults.Kingsoft.Category
	avresult["MAX"] = jsonParser.Data.Attributes.LastAnalysisResults.MAX.Category
	avresult["Malwarebytes"] = jsonParser.Data.Attributes.LastAnalysisResults.Malwarebytes.Category
	avresult["MaxSecure"] = jsonParser.Data.Attributes.LastAnalysisResults.MaxSecure.Category
	avresult["McAfee"] = jsonParser.Data.Attributes.LastAnalysisResults.McAfee.Category
	// avresult["McAfee-GW-Edition"] = jsonParser.Data.Attributes.LastAnalysisResults.McAfee - GW - Edition.Category
	// avresult["MicroWorld-eScan"] = jsonParser.Data.Attributes.LastAnalysisResults.MicroWorld - eScan.Category
	avresult["Microsoft"] = jsonParser.Data.Attributes.LastAnalysisResults.Microsoft.Category
	// avresult["NANO-Antivirus"] = jsonParser.Data.Attributes.LastAnalysisResults.NANO - Antivirus.Category
	avresult["Paloalto"] = jsonParser.Data.Attributes.LastAnalysisResults.Paloalto.Category
	avresult["Panda"] = jsonParser.Data.Attributes.LastAnalysisResults.Panda.Category
	// avresult["Qihoo-360"] = jsonParser.Data.Attributes.LastAnalysisResults.Qihoo-360.Category
	avresult["Rising"] = jsonParser.Data.Attributes.LastAnalysisResults.Rising.Category
	avresult["SUPERAntiSpyware"] = jsonParser.Data.Attributes.LastAnalysisResults.SUPERAntiSpyware.Category
	avresult["Sangfor"] = jsonParser.Data.Attributes.LastAnalysisResults.Sangfor.Category
	avresult["SentinelOne"] = jsonParser.Data.Attributes.LastAnalysisResults.SentinelOne.Category
	avresult["Sophos"] = jsonParser.Data.Attributes.LastAnalysisResults.Sophos.Category
	avresult["Symantec"] = jsonParser.Data.Attributes.LastAnalysisResults.Symantec.Category
	avresult["SymantecMobileInsight"] = jsonParser.Data.Attributes.LastAnalysisResults.SymantecMobileInsight.Category
	avresult["TACHYON"] = jsonParser.Data.Attributes.LastAnalysisResults.TACHYON.Category
	avresult["Tencent"] = jsonParser.Data.Attributes.LastAnalysisResults.Tencent.Category
	avresult["Trapmine"] = jsonParser.Data.Attributes.LastAnalysisResults.Trapmine.Category
	avresult["TrendMicro"] = jsonParser.Data.Attributes.LastAnalysisResults.TrendMicro.Category
	// avresult["TrendMicro-HouseCall"] = jsonParser.Data.Attributes.LastAnalysisResults.TrendMicro - HouseCall.Category
	avresult["VBA32"] = jsonParser.Data.Attributes.LastAnalysisResults.VBA32.Category
	avresult["VIPRE"] = jsonParser.Data.Attributes.LastAnalysisResults.VIPRE.Category
	avresult["ViRobot"] = jsonParser.Data.Attributes.LastAnalysisResults.ViRobot.Category
	avresult["Webroot"] = jsonParser.Data.Attributes.LastAnalysisResults.Webroot.Category
	avresult["Yandex"] = jsonParser.Data.Attributes.LastAnalysisResults.Yandex.Category
	avresult["Zillya"] = jsonParser.Data.Attributes.LastAnalysisResults.Zillya.Category
	avresult["ZoneAlarm"] = jsonParser.Data.Attributes.LastAnalysisResults.ZoneAlarm.Category
	avresult["Zoner"] = jsonParser.Data.Attributes.LastAnalysisResults.Zoner.Category
	avresult["eGambit"] = jsonParser.Data.Attributes.LastAnalysisResults.EGambit.Category
	// AV_dict := []string{"ALYac", "APEX", "AVG", "Acronis", "Ad-Aware", "AegisLab", "AhnLab-V3", "Alibaba", "Antiy-AVL", "Arcabit", "Avast", "Avast-Mobile", "Avira", "Baidu", "BitDefender", "BitDefenderTheta", "Bkav", "CAT-QuickHeal", "CMC", "ClamAV", "Comodo", "CrowdStrike", "Cybereason", "Cylance", "Cyren", "DrWeb", "ESET-NOD32", "Emsisoft", "Endgame", "F-Prot", "F-Secure", "FireEye", "Fortinet", "GData", "Ikarus", "Invincea", "Jiangmin", "K7AntiVirus", "K7GW", "Kaspersky", "Kingsoft", "MAX", "Malwarebytes", "MaxSecure", "McAfee", "McAfee-GW-Edition", "MicroWorld-eScan", "Microsoft", "NANO-Antivirus", "Paloalto", "Panda", "Qihoo-360", "Rising", "SUPERAntiSpyware", "Sangfor", "SentinelOne", "Sophos", "Symantec", "SymantecMobileInsight", "TACHYON", "Tencent", "Trapmine", "TrendMicro", "TrendMicro-HouseCall", "VBA32", "VIPRE", "ViRobot", "Webroot", "Yandex", "Zillya", "ZoneAlarm", "Zoner", "eGambit"}

	for i := range avresult {
		if avresult[i] != "undetected" && avresult[i] != "type-unsupported" {
			malicious = true
		}
	}
	return malicious
}

// VirusTotal Hit
//
func main() {
	clearScreen()
	// Get the type of scan the user wants to do ...
	// entire dir, single file, or entire computer
	banner()
	scantype := flag.String("type", "dir", "Enter the type of scan")
	upath := flag.String("path", "/", "Enter the file path")
	flag.Parse()
	switch *scantype {
	case "dir":
		fmt.Println("dir")
	case "file":
		res := SingleFileMode(*upath)
		switch res {
		case true:
			fmt.Println("File is Malicious/Suspicious")
		case false:
			fmt.Println("File is safe")
		}
	}
}
