<H1 align="center"> Discord Bot em Go - Projeto Impulso Team </H1>
<p align="center">
  <img src="https://img.shields.io/github/license/ineBallardin/discord-go-bot-impulso" alt="License">
  <img src="https://img.shields.io/github/repo-size/ineBallardin/discord-go-bot-impulso" alt="Repo Size">
  <img src="https://img.shields.io/github/last-commit/ineBallardin/discord-go-bot-impulso" alt="Last Commit">
  <img src="https://img.shields.io/github/issues/ineBallardin/discord-go-bot-impulso" alt="Issues">
  <img src="https://img.shields.io/github/forks/ineBallardin/discord-go-bot-impulso" alt="Forks">
  <img src="https://img.shields.io/github/stars/ineBallardin/discord-go-bot-impulso" alt="Stars">
  <img src="https://img.shields.io/badge/Go-1.19%2B-blue" alt="Go Version">
  <img src="https://img.shields.io/badge/Discord-API%20v10-blueviolet" alt="Discord API Version">
  <img src="https://img.shields.io/badge/Status-Em%20Desenvolvimento-yellow" alt="Status">
</p>

## **Sumário**

1. [Sobre o Projeto](#sobre-o-projeto)
2. [Tecnologias Utilizadas](#tecnologias-utilizadas)
3. [Série de Lives no YouTube](#série-de-lives-no-youtube)
4. [Como Executar o Projeto Localmente](#como-executar-o-projeto-localmente)
   - [Pré-requisitos](#pré-requisitos)
   - [Passo a Passo](#passo-a-passo)
5. [Instruções de Deploy no Google Cloud](#instruções-de-deploy-no-google-cloud)
   - [1. Configurar uma Conta no Google Cloud](#1)
   - [2. Criar a Máquina Virtual (VM)](#2)
   - [3. Conectar-se à VM via SSH](#3)
   - [4. Instalar o Docker](#4)
   - [5. Instalar o Git e Clonar o Repositório do Bot](#5)
   - [6. Configurar Variáveis de Ambiente](#6)
   - [7. Compilar e Rodar o Bot](#7)
   - [8. Configurar Alertas de Orçamento](#8)
6. [Contribuição](#contribuição)
7. [Licença](#licença)

---
## Sobre o Projeto
Este projeto foi desenvolvido em parceria com a Impulso Team com o objetivo de introduzir conceitos fundamentais de backend usando a linguagem Go. Ele explora padrões de projeto, comunicação entre sistemas, requisições HTTP, APIs RESTful e webhooks de forma prática e divertida, através da criação de um bot para Discord.

Os principais conceitos abordados são:
- **Padrões de Projeto:**
  - **Singleton**: Aplicado para gerenciar o client do Discord e as configurações da aplicação.
  - **Command Pattern**: Utilizado para facilitar a criação e gerenciamento de comandos do bot.
- **Conceitos de APIs e Webhooks:** Envolvendo comunicação com a API do Discord.
- **Requisições HTTP e manipulação de erros:** Seguindo boas práticas de desenvolvimento.
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>

## Tecnologias Utilizadas

- **Linguagem:** [Go](https://go.dev)
- **Frameworks e Bibliotecas:**
  - [Gorilla WebSocket](https://github.com/gorilla/websocket): Conexão com o gateway do Discord.
  - [GoDotEnv](https://github.com/joho/godotenv): Gerenciamento de variáveis de ambiente.
- **Integrações:** [API do Discord](https://discord.com/developers/docs/reference): Eventos e comunicação com a plataforma.
- **Infraestrutura:** [Google Cloud - Compute Engine](https://console.cloud.google.com/compute): Hospedagem em produção.

## Série de Lives no YouTube
O desenvolvimento deste projeto foi documentado em uma série de lives no canal da Impulso Team no YouTube:

1. [Criando Bots para Discord com Go: Do Zero à Produção #01 | Introdução ao Go e à API do Discord](https://youtu.be/QQgyFL6kBXU)
2. [Criando Bots para Discord com Go: Interagindo com a API do Discord](https://www.youtube.com/watch?v=XYgBLc2cYJM)
3. [Criando Bots para Discord com Go: Padrões de Projeto Úteis para Bots](https://www.youtube.com/watch?v=GDA1Yr6rwDI)
4. [Criando Bots para Discord com Go: Mantendo o bot online e implementando Comandos e Eventos](https://www.youtube.com/watch?v=p4cjIP6wqqU)
5. [Criando Bots para Discord com Go: Melhorando o Bot com Boas Práticas](https://www.youtube.com/watch?v=DyItWMsoJq4&t)
6. [Criando Bots para Discord com Go: Como Hospedar e Manter o Bot em Produção](https://www.youtube.com/watch?v=o6bDoaCvo3A)
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>

## Como Executar o Projeto Localmente

### Pré-requisitos
<table>
  <tr>
    <th>Pré-requisito</th>
    <th>Descrição</th>
    <th>Link</th>
  </tr>
  <tr>
    <td>Go</td>
    <td>Versão 1.19 ou superior</td>
    <td><a href="https://golang.org/">Download</a></td>
  </tr>
  <tr>
    <td>Git</td>
    <td>Para clonar o repositório</td>
    <td><a href="https://git-scm.com/">Download</a></td>
  </tr>
  <tr>
    <td>Discord Account</td>
    <td>Conta no Discord para criar aplicativos e bots</td>
    <td><a href="https://discord.com/developers/applications">Discord Developer Portal</a></td>
  </tr>
    <tr>
    <td>Docker</td>
    <td>Para executar o bot em um container</td>
    <td><a href="https://www.docker.com/">Download</a></td>
  </tr>
  <tr>
    <td>Docker Compose</td>
    <td>Para gerenciar o container do bot</td>
    <td><a href="https://docs.docker.com/compose/">Download</a></td>
  </tr>
</table>

> Você pode consultar o manual de instalação do Docker específico para seu SO [clicando aqui](https://docs.docker.com/engine/install/)

### Passo a Passo

#### 1. Clone o Repositório
```bash
git clone https://github.com/ineBallardin/discord-go-bot-impulso.git
cd discord-go-bot-impulso
```

#### 2. Configuração do Ambiente
Crie um arquivo `.env` na raiz do projeto e configure as variáveis de ambiente exigidas:
```env
BASE_URL=https://discord.com/api/v10
TOKEN=seu-token-de-bot
CHANNEL_ID=id-do-canal
APPLICATION_ID=id-da-aplicacao
GUILD_ID=id-do-servidor
GATEWAY_URL=wss://gateway.discord.gg/?v=10&encoding=json
```
> [Clique aqui](https://discord.com/developers/docs) para acessar a  documentação oficial do Discord para:
> ↪ [Criar um app no Discord](https://discord.com/developers/docs/quick-start/getting-started#step-1-creating-an-app) e setar as configurações iniciais ou se preferir siga os passos abaixo:
1. `TOKEN do Bot`:
   - Acesse o [Discord Developer Portal](https://discord.com/developers/applications).
   - Crie ou selecione seu aplicativo.
   - Vá para a aba **Bot** e clique em **Reset Token** para gerar o token.

2. `CHANNEL_ID`:
   - Ative o **modo desenvolvedor** no Discord: Configurações > Avançado > Modo Desenvolvedor.
   - Clique com o botão direito no canal desejado e selecione **Copiar ID**.

3. `APPLICATION_ID`:
   - No [Discord Developer Portal](https://discord.com/developers/applications), acesse seu aplicativo.
   - O **Application ID** está na seção "General Information".

4. `GUILD_ID` (ID do Servidor):
   - No Discord, clique com o botão direito no servidor (guild) desejado e selecione **Copiar ID** (com o Modo Desenvolvedor ativado).

#### 5. Execute o Bot Localmente
- Construa e execute o container Docker:
```bash
docker compose build
docker compose up -d
```

<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>

## Instruções de Deploy no Google Cloud

Este guia descreve o passo a passo para fazer o deploy do **bot de Discord** em uma máquina virtual (VM) do **Google Cloud Platform (GCP)**.

<details id="1"> <summary style="font-size: 1.15em; font-weight: bold;">1. Configurar uma Conta no Google Cloud</summary>

- Acesse o [Google Cloud Console](https://console.cloud.google.com/welcome?hl=pt).
- Crie uma conta ou faça login com sua conta existente.
- Adicione uma forma de pagamento (necessária para ativar o Free Tier, mas você não será cobrado se seguir os limites gratuitos).
> **O Google Cloud oferece um Free Tier com limites gratuitos, como 1 VM e2-micro e 30 GB de armazenamento padrão. Confira os detalhes completos [neste link oficial](https://cloud.google.com/free/docs/free-cloud-features).**
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

<details id="2"> <summary style="font-size: 1.15em; font-weight: bold;">2. Criar a Máquina Virtual (VM)</summary>

- No Google Cloud Console, acesse **Compute Engine** > **Instâncias de VM**.
- Clique em **"Criar Instância"**.
- Configure a instância da seguinte maneira:
  - **Nome:** `discord-go-bot-impulso`
  - **Região:** `us-west1` (Oregon) ou outra região compatível com o Free Tier.
  - **Série:** `E2`
  - **Tipo de máquina:** `e2-micro` (1 vCPU, 0.6 GB de memória — compatível com o Free Tier).
- Na seção **SO e Armazenamento**:
  - **Sistema operacional:** Ubuntu 20.04 LTS.
  - **Tipo de disco:** Disco permanente padrão.
  - **Tamanho do disco:** `10 GB` (dentro do limite do Free Tier).
- Clique em **Criar** e aguarde a inicialização da VM.
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>


<details id="3"> <summary style="font-size: 1.15em; font-weight: bold;">3. Conectar-se à VM via SSH</summary>

- No [Google Cloud Console](https://console.cloud.google.com/welcome?hl=pt), acesse a página **Compute Engine** > **Instâncias de VM**.
- Na lista de instâncias, localize a sua VM `discord-go-bot-impulso`.
- Clique em **"Conectar"** > **"Abrir em uma nova janela de terminal SSH"**.
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

<details id="4"> <summary style="font-size: 1.15em; font-weight: bold;">4. Instalar o Docker e docker-compose na VM</summary>

> Você encontra instruções de instalação do Docker para Debian [clicando aqui](https://docs.docker.com/engine/install/debian/)

### **1. Instale Dependências Necessárias**
- Execute:
```bash
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg
```

---

### **2. Adicione a Chave GPG do Docker para Debian**

```bash
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg
```

---

### **3. Configure o Repositório Correto para Debian**
- Adicione o repositório **Debian** e não Ubuntu:

```bash
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

---

### **4. Atualize os Pacotes e Instale o Docker**
- Atualize a lista de pacotes e instale o Docker:
```bash
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

---

### **5. Adicione o Usuário ao Grupo Docker**
- Para evitar usar `sudo` ao rodar comandos do Docker, adicione seu usuário ao grupo `docker`:
```bash
sudo usermod -aG docker $USER
```

- Reinicie a sessão para que a mudança tenha efeito. Basta fechar o SSH e se conectar novamente.

---

### **6. Verifique a Instalação**
- Teste se o Docker está funcionando corretamente:

```bash
docker --version
docker compose --version
```
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

<details id="5"> <summary style="font-size: 1.15em; font-weight: bold;">5. Instalar o Git e Clonar o Repositório do Bot</summary>

- Instale o [Git](https://git-scm.com/downloads):
```bash
sudo apt install -y git
```
- Clone o repositório do bot:
```bash
git clone https://github.com/ineBallardin/discord-go-bot-impulso.git
cd discord-go-bot-impulso
```
- Baixe as dependências:
```bash
go mod tidy
```
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

<details id="6"> <summary style="font-size: 1.15em; font-weight: bold;">6. Configurar Variáveis de Ambiente</summary>

- Crie o arquivo `.env`:
```bash
nano .env
```
- Adicione as variáveis necessárias:
```env
BASE_URL=https://discord.com/api/v10
TOKEN=seu-token-do-bot
APPLICATION_ID=seu-application-id
GUILD_ID=seu-guild-id
GATEWAY_URL=wss://gateway.discord.gg/?v=10&encoding=json
```
- Salve (`Ctrl + O`, `Enter`) e feche o editor (`Ctrl + X`).
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

<details id="7"> <summary style="font-size: 1.15em; font-weight: bold;">7. Subir o Container</summary>

```bash
docker compose build
docker compose up -d
```
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

<details id="8"> <summary style="font-size: 1.15em; font-weight: bold;"><strong>8. Configurar Alertas de Orçamento</strong></summary>

Para evitar cobranças, configure alertas no Google Cloud:
- Acesse **Billing** > **Orçamentos e Alertas**.
- Crie um orçamento com o valor **US$ 0.00**.
- Configure notificações por e-mail.
Para mais informações sobre o Free Tier do Google Cloud, confira o [link oficial](https://cloud.google.com/free/docs/free-cloud-features).
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
</details>

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests com melhorias, correções ou novas ideias para o bot.

## Licença
Este projeto é licenciado sob a licença MIT.
<div align="right">
  <a href="#discord-bot-em-go---projeto-impulso-team" style="
    display: inline-block;
    padding: 6px 12px;
    font-size: 14px;
    font-weight: 600;
    color: #ffffff;
    background-color: #2ea44f;
    border-radius: 6px;
    text-decoration: none;
    box-shadow: 0 1px 2px rgba(27, 31, 35, 0.1);
  ">⬆ Voltar ao Topo</a>
</div>
