# refund-api

## Este é um projeto de api para estudos da linguagem *Golang*. 


Essa Api foi desenvolvida após ver a necessidade de uma empresa de turismo no departamento de reembolso. Este serviço tem as seguntes funcionalidades: 

- Busca de todos registros no banco de dados;
- Busca por um ID de um registro;
- Busca por um registro através do número do bilhete aéreo;
- Busca por registro com base no identificador da agencia;
- Busca por registro com base no identificador da agencia + intervalo de período;
- Criação de um registro de reembolso;
- Atualização de um registro;
- Exclusão através de um ID;

### As libs utilizadas: 

- Postgres, para base de dados;
- Gorm, para persistência na base de dados;
- Gorilla mux, para roteamento dos endpoints;
- Docker, para maquina virtual e criação da base de dados;

## Para execução deste pequeno projeto: 

Verifique se o SDK do Golang está instalado em sua maquina, conforme instruções na pagina oficial [https://go.dev/](https://go.dev/doc/install)

No terminal:

- Será necessário que o docker seja inicializado, com o comando *docker-compose up* .
- Após o docker estar devidamente iniciado, na raiz do projeto executar o comando *go run main.go*.
