package vscodeanalysis

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// DEFECTS 不具合リスト
type DEFECTS struct {
	Defact []DEFECT `xml:"DEFECT"`
}

// DEFECT 不具合単位
type DEFECT struct {
	Sfa         SFA    `xml:"SFA"`
	DefactCode  uint   `xml:"DEFECTCODE"`
	Description string `xml:"DESCRIPTION"`
	Function    string `xml:"FUNCTION"`
	Decorated   string `xml:"DECORATED"`
	FuncLine    uint   `xml:"FUNCLINE"`
	Path        PATH   `xml:"PATH"`
}

// SFA SFAエレメント
type SFA struct {
	FilePath string `xml:"FILEPATH"`
	FileName string `xml:"FILENAME"`
	Line     uint   `xml:"LINE"`
	Colum    uint   `xml:"COLUMN"`
}

// PATH PATHエレメント
type PATH struct{}

// LoadDefectsData 不具合データを読み込む
func LoadDefectsData(filename string) (DEFECTS, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return DEFECTS{}, err
	}

	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return DEFECTS{}, err
	}

	var defectsData DEFECTS
	err = xml.Unmarshal(xmlData, &defectsData)

	return defectsData, err
}
