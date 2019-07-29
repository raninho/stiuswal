package lawsuit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/raninho/stiuswal/test/fixture"
)

func Test_makeURLs(t *testing.T) {
	lawsuitNumber := "0019528-16.2005.8.02.0001"

	urls, err := makeURLs(lawsuitNumber)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(urls) != 2 {
		t.Fatal("len(urls) != 2")
	}

	tjal1Grau := "https://www2.tjal.jus.br/cpopg/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0019528-16.2005&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0019528-16.2005.8.02.0001&dadosConsulta.valorConsulta=&uuidCaptcha="
	tjal2Grau := "https://www2.tjal.jus.br/cposg5/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0019528-16.2005&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0019528-16.2005.8.02.0001&dadosConsulta.valorConsulta=&uuidCaptcha="

	if urls[0] != tjal1Grau {
		t.Fatal("urls[0] != tjal1Grau")
	}

	if urls[1] != tjal2Grau {
		t.Fatal("urls[0] != tjal1Grau")
	}

	lawsuitNumber = "0019528-16.2005.8.12.0001"

	urls, err = makeURLs(lawsuitNumber)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(urls) != 2 {
		t.Fatal("len(urls) != 2")
	}

	tjms1Grau := "https://esaj.tjms.jus.br/cpopg5/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0019528-16.2005&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0019528-16.2005.8.12.0001&dadosConsulta.valorConsulta=&uuidCaptcha="
	tjms2Grau := "https://esaj.tjms.jus.br/cposg5/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0019528-16.2005&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0019528-16.2005.8.12.0001&dadosConsulta.valorConsulta=&uuidCaptcha="

	if urls[0] != tjms1Grau {
		t.Fatal("urls[0] != tjal1Grau")
	}

	if urls[1] != tjms2Grau {
		t.Fatal("urls[0] != tjal1Grau")
	}
}

func TestDoCrawler(t *testing.T) {
	tjalLawsuitNumber := "0710802-55.2018.8.02.0001"

	lsCrawler, err := DoCrawler(tjalLawsuitNumber)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(lsCrawler) != 2 {
		t.Fatal("len(lsCrawler) != 2")
	}

	if lsCrawler["grau1"].Area != "Área: Cível " {
		t.Fatal(`lsCrawler["grau1"].Area == "Área: Cível`, lsCrawler["grau1"].Area)
	}

	if lsCrawler["grau1"].Assunto != "Dano Material" {
		t.Fatal(`lsCrawler["grau1"].Assunto != "Dano Material"`, lsCrawler["grau1"].Assunto)
	}

	if lsCrawler["grau1"].Classe != "Procedimento Ordinário  " {
		t.Fatal(`lsCrawler["grau1"].Classe != "Procedimento Ordinário"`, lsCrawler["grau1"].Classe)
	}

	if lsCrawler["grau1"].Juiz != "José Cícero Alves da Silva" {
		t.Fatal(`lsCrawler["grau1"].Juiz != "José Cícero Alves da Silva"`, lsCrawler["grau1"].Juiz)
	}

	if lsCrawler["grau1"].DataDeDistribuição != "02/05/2018 às 19:01 - Sorteio" {
		t.Fatal(`lsCrawler["grau1"].DataDeDistribuição != "02/05/2018 às 19:01 - Sorteio"`, lsCrawler["grau1"].DataDeDistribuição)
	}

	if lsCrawler["grau1"].ValorDaAção != "R$ 281.178,42" {
		t.Fatal(`lsCrawler["grau1"].ValorDaAção != "R$ 281.178,42"`, lsCrawler["grau1"].ValorDaAção)
	}

	if len(lsCrawler["grau1"].PartesDoProcesso) != 2 {
		t.Fatal(`len(lsCrawler["grau1"].PartesDoProcesso) != 2`)
	}

	if lsCrawler["grau1"].PartesDoProcesso[0] != "Autor: José Carlos Cerqueira Souza Filho  Advogado: Vinicius Faria de Cerqueira      " {
		t.Fatal(`lsCrawler["grau1"].PartesDoProcesso[0] != "Autor: José Carlos Cerqueira Souza Filho  Advogado: Vinicius Faria de Cerqueira      "`, lsCrawler["grau1"].PartesDoProcesso[0])
	}

	if lsCrawler["grau1"].PartesDoProcesso[1] != "Réu: Cony Engenharia Ltda.  Advogado: Marcus Vinicius Cavalcante Lins Filho     Advogado: Thiago Maia Nobre Rocha     Advogado: Orlando de Moura Cavalcante Neto      " {
		t.Fatal(`lsCrawler["grau1"].PartesDoProcesso[1] != "Réu: Cony Engenharia Ltda.  Advogado: Marcus Vinicius Cavalcante Lins Filho     Advogado: Thiago Maia Nobre Rocha     Advogado: Orlando de Moura Cavalcante Neto      "`, lsCrawler["grau1"].PartesDoProcesso[1])
	}

	if len(lsCrawler["grau1"].ListaDasMovimentações) != 37 {
		t.Fatal(`len(lsCrawler["grau1"].ListaDasMovimentações) != 37`, len(lsCrawler["grau1"].ListaDasMovimentações))
	}
}

func MockCrawler() *httptest.Server {
	testServerCrawlerMock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", fixture.HTMLCrawler)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", `{}`)
	}))

	return testServerCrawlerMock
}

func TestCrawler(t *testing.T) {
	testServerCrawler := MockCrawler()
	defer testServerCrawler.Close()

	ls, err := crawler(testServerCrawler.URL)
	if err != nil {
		t.Fatal(err.Error())
	}

	if ls.Area != "Área: Cível " {
		t.Fatal(`lsCrawler["grau1"].Area == "Área: Cível`, ls.Area)
	}

	if ls.Assunto != "Dano Material" {
		t.Fatal(`lsCrawler["grau1"].Assunto != "Dano Material"`, ls.Assunto)
	}

	if ls.Classe != "Procedimento Ordinário  " {
		t.Fatal(`lsCrawler["grau1"].Classe != "Procedimento Ordinário"`, ls.Classe)
	}

	if ls.Juiz != "José Cícero Alves da Silva" {
		t.Fatal(`lsCrawler["grau1"].Juiz != "José Cícero Alves da Silva"`, ls.Juiz)
	}

	if ls.DataDeDistribuição != "02/05/2018 às 19:01 - Sorteio" {
		t.Fatal(`lsCrawler["grau1"].DataDeDistribuição != "02/05/2018 às 19:01 - Sorteio"`, ls.DataDeDistribuição)
	}

	if ls.ValorDaAção != "R$ 281.178,42" {
		t.Fatal(`lsCrawler["grau1"].ValorDaAção != "R$ 281.178,42"`, ls.ValorDaAção)
	}

	if len(ls.PartesDoProcesso) != 2 {
		t.Fatal(`len(lsCrawler["grau1"].PartesDoProcesso) != 2`)
	}

	if ls.PartesDoProcesso[0] != "Autor: José Carlos Cerqueira Souza Filho  Advogado: Vinicius Faria de Cerqueira      " {
		t.Fatal(`lsCrawler["grau1"].PartesDoProcesso[0] != "Autor: José Carlos Cerqueira Souza Filho  Advogado: Vinicius Faria de Cerqueira      "`, ls.PartesDoProcesso[0])
	}

	if ls.PartesDoProcesso[1] != "Réu: Cony Engenharia Ltda.  Advogado: Marcus Vinicius Cavalcante Lins Filho     Advogado: Thiago Maia Nobre Rocha     Advogado: Orlando de Moura Cavalcante Neto      " {
		t.Fatal(`lsCrawler["grau1"].PartesDoProcesso[1] != "Réu: Cony Engenharia Ltda.  Advogado: Marcus Vinicius Cavalcante Lins Filho     Advogado: Thiago Maia Nobre Rocha     Advogado: Orlando de Moura Cavalcante Neto      "`, ls.PartesDoProcesso[1])
	}

	if len(ls.ListaDasMovimentações) != 37 {
		t.Fatal(`len(lsCrawler["grau1"].ListaDasMovimentações) != 37`, len(ls.ListaDasMovimentações))
	}
}
