# FactUp

"Little facts about anything" social media

## Escopo

Rede social de pequenos fatos sobre qualquer coisa, onde é possível postar apenas 3 fatos diários por usuário.

## Requisitos funcionais

- RF001: O sistema deve permitir o login do usuário pelo provedor Google.
- RF002: O sistema deve permitir o login do usuário pelo provedor GitHub.
- RF003: O sistema deve permitir o login do usuário pelo provedor Discord.
- RF004: O sistema deve permitir a publicação de apenas 3 fatos diariamente por usuário.
- RF005: O sistema deve permitir o "upvote" (voto positivo) do fato.
- RF006: O sistema deve permitir o "downvote" (voto negativo) do fato.
- RF007: O sistema deve deletar o fato caso a contagem de votos esteja em 8 negativo.
- RF008: O sistema deve permitir o cadastro de Display Name do usuário.
- RF009: O sistema deve permitir o usuário não cadastrado apenas visualizar os fatos.

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
- (FRONT) Criar página com scroll infinito para os fatos.
- (FRONT) Criar página de editar Display Name.
- (FRONT) Criar página de criar fato.
