package lawsuit

import (
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

const (
	tjal = "02"
	tjms = "12"

	tjal1Grau = "https://www2.tjal.jus.br/cpopg/"
	tjal2Grau = "https://www2.tjal.jus.br/cposg5/"
	tjms1Grau = "https://esaj.tjms.jus.br/cpopg5/"
	tjms2Grau = "https://esaj.tjms.jus.br/cposg5/"
)

func makeURLs(lawsuitNumber string) ([]string, error) {
	if len(lawsuitNumber) != 25 {
		return []string{}, errors.New("lawsuitNumber is invalid")
	}
	numeroDigitoAnoUnificado := lawsuitNumber[:15]
	foroNumeroUnificado := lawsuitNumber[21:25]
	tj := lawsuitNumber[18:20]

	urls := []string{}
	url1Grau := ""
	url2Grau := ""

	switch tj {

	case tjal:
		url1Grau = tjal1Grau + "search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=" + numeroDigitoAnoUnificado + "&foroNumeroUnificado=" + foroNumeroUnificado + "&dadosConsulta.valorConsultaNuUnificado=" + lawsuitNumber + "&dadosConsulta.valorConsulta=&uuidCaptcha="
		url2Grau = tjal2Grau + "search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=" + numeroDigitoAnoUnificado + "&foroNumeroUnificado=" + foroNumeroUnificado + "&dadosConsulta.valorConsultaNuUnificado=" + lawsuitNumber + "&dadosConsulta.valorConsulta=&uuidCaptcha="

	case tjms:
		url1Grau = tjms1Grau + "search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=" + numeroDigitoAnoUnificado + "&foroNumeroUnificado=" + foroNumeroUnificado + "&dadosConsulta.valorConsultaNuUnificado=" + lawsuitNumber + "&dadosConsulta.valorConsulta=&uuidCaptcha="
		url2Grau = tjms2Grau + "search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=" + numeroDigitoAnoUnificado + "&foroNumeroUnificado=" + foroNumeroUnificado + "&dadosConsulta.valorConsultaNuUnificado=" + lawsuitNumber + "&dadosConsulta.valorConsulta=&uuidCaptcha="

	default:
		return []string{}, errors.New("lawsuitNumber with tj invalid")
	}

	urls = append(urls, url1Grau)
	urls = append(urls, url2Grau)

	return urls, nil
}

type LawSuitCrawler struct {
	Classe                string            `json:"classe"`
	Area                  string            `json:"area"`
	Assunto               string            `json:"assunto"`
	DataDeDistribuição    string            `json:"data_de_distribuição"`
	Juiz                  string            `json:"juiz"`
	ValorDaAção           string            `json:"valor_da_ação"`
	PartesDoProcesso      []string          `json:"partes_do_processo"`
	ListaDasMovimentações map[string]string `json:"lista_das_movimentações"`
}

func DoCrawler(lawsuitNumber string) (map[string]LawSuitCrawler, error) {
	var lss = make(map[string]LawSuitCrawler, 2)

	urls, err := makeURLs(lawsuitNumber)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	var ls1Grau = new(LawSuitCrawler)
	var ls2Grau = new(LawSuitCrawler)

	wg.Add(2)
	go func(url string) {
		ls, _ := crawler(url)
		ls1Grau = ls
		wg.Done()
	}(urls[0])

	go func(url string) {
		ls, _ := crawler(url)
		ls2Grau = ls
		wg.Done()
	}(urls[1])
	wg.Wait()

	lss["grau1"] = *ls1Grau
	lss["grau2"] = *ls2Grau

	return lss, nil
}

func crawler(url string) (*LawSuitCrawler, error) {
	ls := new(LawSuitCrawler)

	c := colly.NewCollector()

	c.OnHTML(`table.secaoFormBody`, func(e *colly.HTMLElement) {
		if e.Attr("id") != "secaoFormConsulta" {
			e.ForEach("td", func(c int, el *colly.HTMLElement) {
				switch c {
				case 4:
					classe := strings.Replace(el.Text, "\n", "", -1)
					classe = strings.Replace(classe, "\t", "", -1)
					classe = strings.Replace(classe, "  ", "", -1)
					ls.Classe = classe
				case 8:
					area := strings.Replace(el.Text, "\n", "", -1)
					area = strings.Replace(area, "\t", "", -1)
					area = strings.Replace(area, "  ", "", -1)
					ls.Area = area
				case 10:
					assunto := strings.Replace(el.Text, "\n", "", -1)
					assunto = strings.Replace(assunto, "\t", "", -1)
					assunto = strings.Replace(assunto, "  ", "", -1)
					ls.Assunto = assunto
				case 14:
					dist := strings.Replace(el.Text, "\n", "", -1)
					dist = strings.Replace(dist, "\t", "", -1)
					dist = strings.Replace(dist, "  ", "", -1)
					ls.DataDeDistribuição = dist
				case 20:
					juiz := strings.Replace(el.Text, "\n", "", -1)
					juiz = strings.Replace(juiz, "\t", "", -1)
					juiz = strings.Replace(juiz, "  ", "", -1)
					ls.Juiz = juiz
				case 22:
					valor := strings.Replace(el.Text, "\n", "", -1)
					valor = strings.Replace(valor, "\t", "", -1)
					valor = strings.Replace(valor, "  ", "", -1)
					ls.ValorDaAção = valor
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
					ls.PartesDoProcesso = append(ls.PartesDoProcesso, "Autor: "+autor)
				case 3:
					reu := strings.Replace(el.Text, "\n", "", -1)
					reu = strings.Replace(reu, "\t", "", -1)
					ls.PartesDoProcesso = append(ls.PartesDoProcesso, "Réu: "+reu)
				}
			})
		}
	})

	c.OnHTML(`tbody#tabelaTodasMovimentacoes`, func(e *colly.HTMLElement) {
		movimento := map[string]string{}

		e.ForEach("tr", func(c int, el *colly.HTMLElement) {
			data := ""
			desc := ""

			el.ForEach("td", func(d int, ell *colly.HTMLElement) {
				switch d {
				case 0:
					dtMov := strings.Replace(ell.Text, "\n", "", -1)
					dtMov = strings.Replace(dtMov, "\t", "", -1)
					data = dtMov

				case 2:
					mov2 := strings.Replace(ell.Text, "\n", "", -1)
					mov2 = strings.Replace(mov2, "\t\t", "", -1)
					desc = mov2
				}

				movimento[strconv.Itoa(c)] = data + " " + desc
			})

		})

		ls.ListaDasMovimentações = movimento
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return ls, nil
}
