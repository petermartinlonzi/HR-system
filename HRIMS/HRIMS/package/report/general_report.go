package report

import (
	"encoding/json"
	"training-backend/package/config"
	"training-backend/package/log"
	"training-backend/package/util"

	"fmt"
	"strings"
	"time"
)

//Generate report
// [title] report title
// [data] contains two dimension data, both in string, including serial number, the first row is the table title
// [columnWidth] array of width of each column, number of table columns should be equal to column dimensions of the data
// [isLandscape] defines the page setup, true for landscape and false for potrait

func GeneralReport(mainTitle, title string, data [][]string, columnWidth []float64, fileName string, fontSize int, isLandcape bool) string {
	begin := time.Now()

	d, err := json.Marshal(data)
	if util.CheckError(err) {
		log.Errorf("error getting json data %v", err)
		return ""
	}

	qrs, doi, err := util.GetQRString(d)
	if util.CheckError(err) {
		log.Errorf("error getting qr string and doi %v", err)
		return ""
	}
	//variable initalisation
	//fmt.Printf("Is landscape: %v\n", isLandcape)
	currentPageNumber = 1
	totalPageNumber = 0
	pdf := initPDF(isLandcape)
	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)

	//add fonts
	addFonts(pdf)
	pdf.AddPage()
	pdf.SetX(rightMargin)
	pdf.SetY(topMargin)
	totalRows = float64(len(data)) //initalise the number of rows

	//display header
	title = strings.ToUpper(title)
	mainHeader(pdf, mainTitle, title, qrs, doi)

	xp := leftMargin
	xy := headingMargin

	columnWidth = normaliseTableWidth(columnWidth, availablePageWidth)

	tableHeading := data[0]
	data = data[1:]
	addTable(pdf, xp, xy, 20, 5, columnWidth, tableHeading, data, fontSize)

	reportDir, err := config.ReportDir()
	if util.CheckError(err) {
		log.Errorf("error getting report directory %v", err)
		return ""
	}
	timeFileName := fmt.Sprintf("-%d", time.Now().Unix())
	path := reportDir + fileName + timeFileName + ".pdf"
	pdf.WritePdf(path)
	pdf.Close()

	end := time.Now()
	fmt.Printf("PDF Report generated in %v\n", end.Sub(begin))
	return path
}
