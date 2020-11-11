package handlers

import(
	"fmt"
	"os"
	"encoding/csv"
	model "github.com/viniciusveu/data-integration-challenge/models"
)

var Companies []model.UpdateCompanyInput


//Função que lê o arquivo CSV e cria os slices com os dados
func ReadCSV(csvName string) error  {

	csvFile, err := os.Open(csvName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//No caso de ser o arquivo 1 os dados são adicionados no banco com ocampo site vazio ("")
	if len(csvLines[1]) == 2 {

		for index, line := range csvLines {
			if index == 0 {
				continue
			}

			company := model.Company{Name: line[0], Zip: line[1], Site: ""}
			model.DB.Create(&company)

		}
	} else { //No caso do arquivo 2 é apenas criado o slice, pois a atualização será feita no handler

		for index, line := range csvLines {
			if index == 0 {
				continue
			}

			Companies = append(Companies, model.UpdateCompanyInput{Name: line[0], Zip: line[1], Site: line[2]})
		}
	}

	return err
}
