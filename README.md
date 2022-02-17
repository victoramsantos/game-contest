# game-contest

Jogo RPG simplificado em Go.

## Cenário

O projeto consiste na criação dos personagens e na execução das lutas.

Os personagens podem ser do tipo Warrior, Thief e Mage, onde cada um possui suas atributos específicos.

Em uma luta, é decidido aleatoriamente o atacante e o oponente. Após isso, eles batalham até a morte...

## APIs 

Abaixo é listado as APIs disponíveis na aplicação. Também está disponível uma [coleção do Postman](assets/postman-collection/game-contest.postman_collection.json) para facilitar o consumo delas.

### GET /character

Retorna uma lista dos personsagens, informando seus nomes, classe e estado de vida.

Exemplo:
Request:
```bash
curl --location --request GET 'http://localhost:8080/character'
```
Respose:
```json
[
    {
        "name": "ster_verstapen",
        "class": "Mage",
        "is_alive": true
    },
    {
        "name": "john_verstapen",
        "class": "Thief",
        "is_alive": false
    }
]

```

### POST /character

Api para criação de um novo personagem

Exemplo:
Request:
```bash
curl --location --request POST 'http://localhost:8080/character' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_name": "John",
    "class_name": "Warrior"
}'
```
Response:
```json
{
    "name": "John",
    "class": {
        "name": "Warrior",
        "attributes": {
            "life": {
                "name": "Life",
                "value": 20
            },
            "strength": {
                "name": "Strength",
                "value": 10
            },
            "skill": {
                "name": "Skill",
                "value": 5
            },
            "intelligence": {
                "name": "Intelligence",
                "value": 5
            }
        },
        "attack_improvement": [
            {
                "target_attribute": {
                    "name": "Strength",
                    "value": 10
                },
                "improvement": 80
            },
            {
                "target_attribute": {
                    "name": "Skill",
                    "value": 5
                },
                "improvement": 20
            }
        ],
        "velocity_improvement": [
            {
                "target_attribute": {
                    "name": "Skill",
                    "value": 5
                },
                "improvement": 60
            },
            {
                "target_attribute": {
                    "name": "Intelligence",
                    "value": 5
                },
                "improvement": 20
            }
        ]
    },
    "is_alive": true
}
```

### GET /character/:name

API que retorna as informações detalhadas de um dado personagem.

Exemplo:
Request:
```bash
curl --location --request GET 'http://localhost:8080/character/John'
```
Response:
```json
{
    "name": "John",
    "class": "Warrior",
    "life": "Life",
    "strength": "Strength",
    "skill": "Skill",
    "intelligence": "Intelligence",
    "attack": "80% de Strength+20% de Skill",
    "velocity": "60% de Skill+20% de Intelligence"
}
```


### POST /game/start

API que inicializa a batalha entre dois personagens.

Exemplo:
Request:
```bash
curl --location --request POST 'http://localhost:8080/game/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "character_a": "John",
    "character_b": "Mary"
}'
```
Response:
```json
[
    "John (3) foi mais veloz que o Mary (1), e irá começar!",
    "John atacou Mary com 0 de dano, Mary com 12 pontos de vida restantes",
    "Mary atacou John com 16 de dano, John com 4 pontos de vida restantes",
    "John atacou Mary com 5 de dano, Mary com 7 pontos de vida restantes",
    "Mary atacou John com 8 de dano, John com 0 pontos de vida restantes",
    "Mary venceu a batalha! Mary ainda tem 7 pontos de vida restantes!"
]
```

### GET /metrics

Retorna as métricas da aplicação no formato do Prometheus

### GET /health

Api para validar a saúde da aplicação
Exemplo:
Request:
```bash
curl --location --request GET 'http://localhost:8080/health'
```
Response:
```json
{
{
    "Status": true
}
```

## Testes Unitários

Para executar os testes da aplicação execute o comando abaixo na raiz do projeto:
```bash
go test -v -cover -covermode=atomic ./... 
```

## Teste Local

Para executar a aplicação local, foi criado dentro da pasta [assets/localrun](assets/localrun/) o arquivo [setup.sh](assets/localrun/setup.sh), o qual configura o ambiente e executa a imagem docker da aplicação. A configuração do ambiente é feita executando a aplicação e as ferramentas de observbailidade [Prometheus](https://prometheus.io/) e [Grafana](https://grafana.com/) também via docker.

Além de configurar o ambiente, o `setup.sh` também executa o arquivo [bootstraper.go](assets/bootstraper/bootstraper.go). Este arquivo verifica se a aplicação já está de pé e cria alguns personagens. Após a criação destes personagens, é iniciado uma série de lutas entre eles.

Para executar o arquivo `setup.sh`, execute o comando dentro da pasta `assets/localrun/`
```bash
bash setup.sh
```

Para remover os componentes criados, execute na mesma pasta:
```bash
bash down.sh
```

Nas [configurações da aplicação](config/local.yaml) é possível ajustar qual a porta a aplicação será exposta e as configurações de simulação. Estas configurações adicionam alguns `sleeps` nas chamadas aos controllers da aplicação. Isso é interessante para simular um cenário produtivo da aplicação e visualizar como o dashboard da aplicação se comporta.

O dashboard da aplicação pode ser encontrado (após executar o `setup.sh`) em [http://localhost:3000/d/yt-D57f7k/general-dashboard](http://localhost:3000/d/yt-D57f7k/general-dashboard?orgId=1&from=now-5m&to=now&refresh=5s). Também está disponível a interface web do prometheus em [http://localhost:9090/graph](http://localhost:9090/graph) para consumir as métricas da aplicação que estão expostas em [http://localhost:8080/metrics](http://localhost:8080/metrics).

## Requisitos
  - go 1.17.6 
  - docker 20.10.12

## Débitos técnicos
  - Melhorar camada de dados: utilizar hashmap ao invés de slices para persistir os personagens
  - Finalizar testes unitários: adicionar os testes unitários para cenários de dupla chamada do mesmo método (ex.: validar a criação de um mesmo personagem)
  - Melhorar o handler das exceptions: Criar exceptions customizadas e mais significativas (ex.: UserNotFoundException, UserAlreadyExistsException)
  - Melhorar os status code de resposta para erros: Melhorar quando retornar erro de usuário (4xx) e quando usar error de servidor (5xx)
  - Criar swagger para documentar as apis da aplicação: Isso ajudaria a melhorar a documentação da APIs da aplicação (ex.: adicionar os possíveis erros de retorno das chamadas)
