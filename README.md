# stiuswal

You must already have the docker installed, ok?

Running the application first time
------------------
```
make dependency-install
make dependency-start
make run-api
make run-worker
```

Running the test first time
------------------
```
make dependency-install
make dependency-start
make dependency-test
make test-all
```

Consuming the API
------------------

Request:
```
curl -X POST \
  http://localhost:8080/lawsuits \
  -H 'Content-Type: application/json' \
  -d '{
	"lawsuit-number": "0710802-55.2018.8.02.0001" 
}'
```

Response:
```
{"order-id":"879358de-413a-4248-8dcf-6614cd1febc1","lawsuit-number":"0710802-55.2018.8.02.0001"}
```

Request:
```
curl -X GET \
  http://localhost:8080/lawsuits/879358de-413a-4248-8dcf-6614cd1febc1 \
  -H 'Content-Type: application/json'
```

Response:
```
{"ID":15,"OrderID":"879358de-413a-4248-8dcf-6614cd1febc1","LawsuitNumber":"0710802-55.2018.8.02.0001","Processed":true,"Criticized":false,"CreatedAt":"2019-07-29T23:30:41.980914Z","UpdatedAt":"2019-07-29T23:30:42.547113Z","Output":{"grau1":{"area":"Área: Cível ","juiz":"José Cícero Alves da Silva","classe":"Procedimento Ordinário  ","assunto":"Dano Material","valor_da_ação":"R$ 281.178,42","partes_do_processo":["Autor: José Carlos Cerqueira Souza Filho  Advogado: Vinicius Faria de Cerqueira      ","Réu: Cony Engenharia Ltda.  Advogado: Marcus Vinicius Cavalcante Lins Filho     Advogado: Thiago Maia Nobre Rocha     Advogado: Orlando de Moura Cavalcante Neto      "],"data_de_distribuição":"02/05/2018 às 19:01 - Sorteio","lista_das_movimentações": {...}},"grau2":{"area":"","juiz":"","classe":"","assunto":"","valor_da_ação":"","partes_do_processo":null,"data_de_distribuição":"","lista_das_movimentações":null}}}
```