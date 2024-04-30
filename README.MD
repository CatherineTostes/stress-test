# Stress Test

Esse projeto tem como objetivo passar argumentos para realizar um teste de estresse na url informada, mandando um x de requisições com x de concorrência.

## Gerando imagem Docker:
```
docker build -t stress-test .
```

## Argumentos:
- url: url que deseja realizar o teste
- requests: quantidade de requisições
- concurrency: quantidade de concorrência

## Rodando a imagem:

Comando a ser enviado:
```
docker run <imagem docker> —url=http://google.com —requests=1000 —concurrency=10
```

Ficaria dessa forma:
```
docker run stress-test -url=http://google.com -requests=100 -concurrency=10
```

Após executar a aplicação passando os arqumentos, ao final da execução será emitido um relatório com as seguintes informações:
- Tempo total gasto na execução
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).