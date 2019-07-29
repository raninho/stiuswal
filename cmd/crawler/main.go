package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type LawSuitCrawler struct {
	Classe string
	Area string
	Assunto string
	DataDeDistribuição string
	Juiz string
	ValorDaAção string
	PartesDoProcesso string
	ListaDasMovimentações string
}

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www2.tjal.jus.br"))

	c.OnHTML(`table.secaoFormBody`, func(e *colly.HTMLElement) {
		if e.Attr("id") != "secaoFormConsulta" {
			e.ForEach("td", func(c int, el *colly.HTMLElement) {
				switch c {
				case 4:
					classe := strings.Replace(el.Text, "\n", "", -1)
					classe = strings.Replace(classe, "\t", "", -1)
					classe = strings.Replace(classe, "  ", "", -1)
					//fmt.Println("Classe:", classe)
				case 8:
					area := strings.Replace(el.Text, "\n", "", -1)
					area = strings.Replace(area, "\t", "", -1)
					area = strings.Replace(area, "  ", "", -1)
					//fmt.Println("Area:", area)
				case 10:
					assunto := strings.Replace(el.Text, "\n", "", -1)
					assunto = strings.Replace(assunto, "\t", "", -1)
					assunto = strings.Replace(assunto, "  ", "", -1)
					//fmt.Println("Assunto:", assunto)
				case 14:
					dist := strings.Replace(el.Text, "\n", "", -1)
					dist = strings.Replace(dist, "\t", "", -1)
					dist = strings.Replace(dist, "  ", "", -1)
					//fmt.Println("Distribuição:", dist)
				case 20:
					juiz := strings.Replace(el.Text, "\n", "", -1)
					juiz = strings.Replace(juiz, "\t", "", -1)
					juiz = strings.Replace(juiz, "  ", "", -1)
					//fmt.Println("Juiz:", juiz)
				case 22:
					valor := strings.Replace(el.Text, "\n", "", -1)
					valor = strings.Replace(valor, "\t", "", -1)
					valor = strings.Replace(valor, "  ", "", -1)
					//fmt.Println("Valor da ação:", valor)
				}
			})
		}
	})

	c.OnHTML(`table#tablePartesPrincipais`, func(e *colly.HTMLElement) {
		if e.Attr("class") != "mensagemExibindo" {
			e.ForEach("td", func(c int, el *colly.HTMLElement) {
				switch c {
				case 1:
					autor := strings.Replace(el.Text, "\n", "", -1)
					autor = strings.Replace(autor, "\t", "", -1)
					//fmt.Println("autor:", autor)
				case 3:
					re := strings.Replace(el.Text, "\n", "", -1)
					re = strings.Replace(re, "\t", "", -1)
					//fmt.Println("ré:", re)
				}
			})
		}
	})

	c.OnHTML(`tbody#tabelaTodasMovimentacoes`, func(e *colly.HTMLElement) {
			e.ForEach("tr", func(c int, el *colly.HTMLElement) {
				el.ForEach("td", func(d int, ell *colly.HTMLElement) {
					switch d {
					case 0:
						dtMov := strings.Replace(ell.Text, "\n", "", -1)
						dtMov = strings.Replace(dtMov, "\t", "", -1)
						fmt.Println(d, dtMov)
					/*case 1:
						mov1 := strings.Replace(ell.Text, "\n", "", -1)
						mov1 = strings.Replace(mov1, "\t", "", -1)
						fmt.Println(d, mov1)*/
					case 2:
						mov2 := strings.Replace(ell.Text, "\n", "", -1)
						mov2 = strings.Replace(mov2, "\t\t", "", -1)
						fmt.Println(d, mov2)
						//fmt.Println(d, ell.ChildText("span"))
					}
					//mov := strings.Replace(ell.Text, "\n\n", "", -1)
					//mov = strings.Replace(ell.Text, "\t", "", -1)
					//for i := 0; i < len(mov); i++ {
					//	fmt.Printf("%x ", mov[i])
					//}
					//fmt.Println(c, mov)
				})

			})
	})


	c.Visit("https://www2.tjal.jus.br/cpopg/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0710802-55.2018&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0710802-55.2018.8.02.0001&dadosConsulta.valorConsulta=&uuidCaptcha=")

}