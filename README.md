# Metas-Redes IoT Example

- **Emular um sensor que gera valores aleatórios**
- **Desenvolver comunicação entre dois processos**

## Como funciona

1. **Sensor**: aplicação Go que abre uma conexão TCP na porta 9000 e envia um número aleatório por segundo.
2. **Consumer**: outra aplicação Go que se conecta ao sensor e imprime os dados recebidos.
3. Cada componente roda em um container Docker diferente; a comunicação é feita via rede do `docker-compose`.

## Como testar

```bash
cd c:/Users/Walace/Documents/Metas-Redes

docker compose build

docker compose up
```

- O serviço `sensor` escuta em `:9000`.
- O serviço `consumer` se conecta automaticamente a `sensor:9000` e exibe os valores.

Você pode também entrar em qualquer container e ver o binário:

```bash
docker compose exec sensor sh
# ./sensor # roda o sensor dentro do container
```

## Extensões

- Mudar o tipo de transporte (UDP em vez de TCP).
- Adicionar outro processo que consome e grava em banco de dados.
- Exibir valores em página web ou CLI interativa.

Nenhum framework de mensagens externo, tudo em containers Docker, e comunicação simples ponto a ponto.

## Comandos úteis

| Comando | O que faz |
|---------|-----------|
| `sudo apt install netcat` | Instala o `netcat` no Linux (Debian/Ubuntu). |
| `nc localhost 9000` | Conecta ao sensor a partir do host e exibe os valores recebidos. |
| `nc -l 9001` | Abre um listener TCP na porta 9001 (dentro de um container) para testes. |
| `echo "hello" \| nc sensor 9000` | Envia uma mensagem para o sensor via TCP. |
| `docker-compose exec sensor sh` | Abre um shell dentro do container `sensor`. |
| `docker-compose logs -f` | Exibe em tempo real os logs de todos os serviços. |
   
Criar o docker do sensor maquina 1    
   docker compose up --build sensor
Criar o do sonsumer:
    docker build -t consumer-app ./consumer   # ← criou a imagem
    docker run --rm -e SENSOR_ADDR=...:9000 consumer-app  # ← rodou o container
