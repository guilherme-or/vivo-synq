
# vivo-synq

Projeto para o Vivo FastData Challenge, proposto para os alunos do 4o ano de Sistemas de Informação na FIAP, em 2024.

De acordo com o [diagrama da solução](https://drive.google.com/file/d/1307gwE5-zwkV4eI8Rmtf4EdXXMsroeYc/view?usp=sharing), o projeto 
tem como objetivo resolver o problema de alto desacoplamento que a arquitetura atual da Vivo apresenta.

Para isso, a arquitetura sugerida resolve esse problema, trazendo **Consistência**, **Latência**, **Resiliência** e **Evolução**.

A seguir estão listados os principais **componentes** da solução:

- Camada de acesso
  - Proxy reverso e Cache de requisições (NGINX)
  - API (Go Fiber - FastHTTP Framework)
  - Cache de dados recentes (Redis)
  - Base de dados integrada (MongoDB)
  - Dashboard de monitoramento da arquitetura (Prometheus e Grafana)
- Camada de infraestrutura
  - Consumidor da fila de mudanças (Go confluentic/confluent-kafka-go)
  - Fila de mudanças (Kafka e Zookeeper)
  - Leitor de mudanças nas bases legadas (Debezium Connector)
