# FactUp

"Little facts about anything" social media

## Escopo

Rede social de pequenos fatos sobre qualquer coisa, onde é possível postar apenas 3 fatos diários por usuário.

## Tecnologias utilizadas

Planejamento:
- [DrawDB](https://drawdb.app/) - Diagrama do banco de dados.
- [Figma](https://figma.com/) - Wireframing e design da tela.

Infraestrutura:

- Arquitetura - Monolítico.
- [Docker](https://www.docker.com/) - Conteinerizar a aplicação, para fácil manuseio e padronização de ambiente.
- [Docker Compose](https://docs.docker.com/compose/) - Orquestração de Contêineres.
- [Nginx](https://nginx.org/) - Servidor HTTP e Proxy Reverso, servirá para manter tudo sob apenas um domínio/subdomínio (factup.me, aquisição vemos depois).
- [Hostinger](https://www.hostinger.com.br/) - VMs naquele precinho.

Back-end:

- [Hoppscotch](https://hoppscotch.io/) - Testes de API
- [PostgreSQL](https://www.postgresql.org/) - Banco de dados relacional mais utilizado ultimamente.
- [Golang-Migrate](https://github.com/golang-migrate/migrate) - Migrações de banco de dados.
- [Go](https://go.dev/) - Linguagem focada em infraestrutura, está em ascensão para desenvolvimento web.
- [Gin](https://gin-gonic.com/) (framework web) - Framework Web mais utilizado no ecossistema Go.
- [Testify](https://github.com/stretchr/testify) - Testes unitários melhorados com Go.
- [JWT](https://jwt.io) - Autenticação e autorização.
- [OAuth2](https://oauth.net/2/) - Login por outras aplicações, como Google e GitHub ([Goth](https://github.com/markbates/goth)).
- [sqlc](https://sqlc.dev/) - Gerador de código Go a partir de SQL.

Front-end:

- SPA - Preferimos usar o formato SPA (single-page application) ao invés de SSR (server-side rendering) pois algumas coisas não funcionariam no formato SSR.
- [TypeScript](https://www.typescriptlang.org/) - Linguagem de programação baseada no JavaScript, mas com tipos.
- [Vite](https://vite.dev/) - Mais rápido app bundler no mercado.
- [React](https://react.dev/) - Library para interfaces web de usuário.
- [Radix UI](https://www.radix-ui.com/primitives) - Usado para os componentes mais elaborados.
- [TailwindCSS](https://tailwindcss.com/) - Forma mais rápida e eficiente de criar CSS.
- [React Router](https://reactrouter.com/) - Roteador de páginas para React.
- [Tanstack Query](https://tanstack.com/query/latest) - Ferramenta de data-fetching e caching no client-side.

Front-end (mobile):
- Avaliar a criação de um front-end mobile em Flutter.

## Requisitos funcionais

- RF001: O sistema deve permitir o login do usuário pelo provedor Instagram.
- RF002: O sistema deve permitir o login do usuário pelo provedor Facebook.
- RF003: O sistema deve permitir o login do usuário pelo provedor Google.
- RF004: O sistema deve permitir o login do usuário pelo provedor GitHub.
- RF005: O sistema deve permitir o login do usuário pelo provedor Discord.
- RF006: O sistema deve permitir a publicação de apenas 3 posts diariamente por usuário.
- RF007: O sistema deve permitir a publicação de fatos.
- RF008: O sistema deve permitir a publicação de ditados populares.
- RF009: O sistema deve permitir cadastro de texto em um fato.
- RF010: O sistema deve permitir cadastro de uma imagem em um fato.
- RF011: O sistema deve permitir cadastro de fonte em um fato.
- RF012: O sistema deve permitir cadastro de texto em um ditado popular.
- RF013: O sistema deve permitir a curtida (voto positivo) do fato.
- RF014: O sistema deve permitir o "it's fake!" (voto negativo) do fato.
- RF015: O sistema deve permitir a curtida (voto positivo) do ditado popular.
- RF016: O sistema deve permitir o "it's joking!" (voto negativo) do ditado popular.
- RF017: O sistema deve deletar o fato caso ele possua 8 votos negativos.
- RF018: O sistema deve deletar o ditado popular caso ele possua 8 votos negativos.
- RF019: O sistema deve permitir o cadastro de Display Name do usuário.
- RF020: O sistema deve permitir o usuário não cadastrado apenas visualizar os fatos.

## Requisitos não funcionais

- RNF001: O sistema deve usar Go para o back-end.
- RNF002: O sistema deve usar React para o front-end.
- RNF003: O sistema deve responder rápido (até 5 segundos).
- RNF004: O sistema deve possuir layout responsivo.
- RNF005: O sistema deve usar TailwindCSS para a estilização.

## Diagramas C4

### Nível 1

![nivel 1](docs/c4_level_1.png)

### Nível 2

![nivel 2](docs/c4_level_2.png)

### Nível 3

![nivel 3](docs/c4_level_3.png)

## Comandos recorrentes

- [gocyclo](https://github.com/fzipp/gocyclo) - `gocyclo .` - complexidade ciclomática
- [cloc](https://github.com/AlDanial/cloc) - `cloc .` - contagem de linhas de código
