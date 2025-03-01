# FactUp

"Little facts about anything" social media

## Escopo

Rede social de pequenos fatos sobre qualquer coisa, onde é possível postar apenas 3 fatos diários por usuário.

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
- RF015: O sistema deve deletar o fato caso ele possua 8 votos negativos.
- RF016: O sistema deve permitir o cadastro de Display Name do usuário.
- RF017: O sistema deve permitir o usuário não cadastrado apenas visualizar os fatos.

## Requisitos não funcionais

- RNF001: O sistema deve usar Go para o back-end.
- RNF002: O sistema deve usar React para o front-end.
- RNF003: O sistema deve responder rápido.
- RNF004: O sistema deve possuir layout responsivo.
- RNF005: O sistema deve usar TailwindCSS para a estilização.

## Atividades

- (BACK) CRUD da tabela de usuários, será possível alterar Display Name.
- (BACK) CRUD (apenas Create, Read e Delete) da tabela de fatos.
- (FRONT) Criar página de login.
- (FRONT) Criar página com scroll infinito estilo tiktok (de 3 em 3 fatos) para os fatos.
- (FRONT) Criar página de editar Display Name.
- (FRONT) Criar página de criar fato.
