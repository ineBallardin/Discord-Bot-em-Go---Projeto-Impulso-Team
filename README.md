# Discord Bot em Go - Projeto Impulso Team

## Sobre o Projeto
Este projeto foi desenvolvido em parceria com a Impulso Team com o objetivo de introduzir conceitos fundamentais de backend usando a linguagem Go. Ele explora padrões de projeto, comunicação entre sistemas, requisições HTTP, APIs RESTful e webhooks de forma prática e divertida, através da criação de um bot para Discord.

Os principais conceitos abordados são:
- **Padrões de Projeto:**
  - **Singleton**: Aplicado para gerenciar o client do Discord e as configurações da aplicação.
  - **Command Pattern**: Utilizado para facilitar a criação e gerenciamento de comandos do bot.
- **Conceitos de APIs e Webhooks:** Envolvendo comunicação com a API do Discord.
- **Requisições HTTP e manipulação de erros:** Seguindo boas práticas de desenvolvimento.

## Tecnologias Utilizadas
- **Linguagem:** Go
- **Frameworks/Bibliotecas:**
  - [Gorilla WebSocket](https://github.com/gorilla/websocket) para conexão com o gateway do Discord.
  - [GoDotEnv](https://github.com/joho/godotenv) para gerenciamento de variáveis de ambiente.
- **API do Discord:** Para integração e eventos.
- **Heroku:** Para hospedagem do bot em produção.

## Série de Lives no YouTube
O desenvolvimento deste projeto foi documentado em uma série de lives no canal da Impulso Team no YouTube:

1. [Criando Bots para Discord com Go: Do Zero à Produção #01 | Introdução ao Go e à API do Discord](https://youtu.be/QQgyFL6kBXU)
2. [Criando Bots para Discord com Go: Interagindo com a API do Discord](https://www.youtube.com/watch?v=XYgBLc2cYJM)
3. [Criando Bots para Discord com Go: Padrões de Projeto Úteis para Bots](https://www.youtube.com/watch?v=GDA1Yr6rwDI)
4. [Criando Bots para Discord com Go: Mantendo o bot online e implementando Comandos e Eventos](https://www.youtube.com/watch?v=p4cjIP6wqqU)
5. [Criando Bots para Discord com Go: Melhorando o Bot com Boas Práticas](https://www.youtube.com/watch?v=DyItWMsoJq4&t)
6. [Criando Bots para Discord com Go: Como Hospedar e Manter o Bot em Produção](https://www.youtube.com/watch?v=o6bDoaCvo3A)

## Como Executar o Projeto Localmente

### Pré-requisitos
- [Go](https://golang.org/) instalado (versão 1.19 ou superior).
- Conta no Discord com permissões para criar aplicativos e bots.
- [Git](https://git-scm.com/) para clonar o repositório.
- Variáveis de ambiente configuradas (descritas no passo **Configuração do Ambiente**).

### Passo a Passo

#### 1. Clone o Repositório
```bash
git clone https://github.com/ineBallardin/discord-go-bot-impulso.git
cd discord-go-bot-impulso
```

#### 2. Configure o Módulo do Go e Instale as Dependências
Inicialize o módulo do Go e baixe as dependências necessárias:
```bash
go mod init discord_bot
go mod tidy
```

#### 3. Crie o Binário Executável
Gere o executável do bot:
```bash
go build -o go-bot
```

#### 4. Configuração do Ambiente
Crie um arquivo `.env` na raiz do projeto e configure as variáveis de ambiente exigidas:
```env
BASE_URL=https://discord.com/api/v10
TOKEN=seu-token-de-bot
CHANNEL_ID=id-do-canal
APPLICATION_ID=id-da-aplicacao
GUILD_ID=id-do-servidor
GATEWAY_URL=wss://gateway.discord.gg/?v=10&encoding=json
```

#### 5. Execute o Bot Localmente
Inicie o bot com o binário criado:
```bash
./go-bot
```
Se tudo estiver configurado corretamente, você verá mensagens no terminal indicando que o bot está online e conectado ao gateway do Discord.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests com melhorias, correções ou novas ideias para o bot.

## Licença
Este projeto é licenciado sob a licença MIT.

