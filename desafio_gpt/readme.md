Legal! Vamos lá, vou te passar um desafio para consolidar os conceitos de ponteiros em Go. O objetivo será trabalhar com structs e ponteiros para manipular dados de forma eficiente.

### Desafio: Sistema de Gerenciamento de Livros

Imagina que você está criando um sistema simples de gerenciamento de livros. Cada livro possui um título, um autor e um número de páginas. O objetivo é usar ponteiros para atualizar os detalhes dos livros e garantir que as modificações sejam refletidas corretamente.

#### Parte 1: Definir a estrutura `Livro`

Crie a struct `Livro` com os seguintes campos:
- `Título` (string)
- `Autor` (string)
- `Páginas` (int)

#### Parte 2: Função para Atualizar o Título

Crie uma função `atualizarTitulo` que recebe um ponteiro para um livro e o novo título como parâmetros. Ela deve alterar o título do livro.

#### Parte 3: Função para Atualizar o Autor

Crie uma função `atualizarAutor` que recebe um ponteiro para o livro e o novo autor como parâmetros. Ela deve alterar o autor do livro.

#### Parte 4: Função para Incrementar o Número de Páginas

Crie uma função `adicionarPaginas` que recebe um ponteiro para o livro e um número de páginas a ser adicionado. A função deve incrementar o número de páginas do livro.

#### Parte 5: Mostrar os Detalhes do Livro

Crie uma função `mostrarDetalhes` que recebe um ponteiro para o livro e imprime seus detalhes no formato:

```
Título: [Título]
Autor: [Autor]
Páginas: [Número de Páginas]
```

#### Parte 6: Teste no `main`

No `main`, crie um livro e utilize as funções criadas para atualizar seu título, autor e número de páginas, mostrando os detalhes após cada modificação.

---

Esse desafio te força a manipular structs e ponteiros de maneira prática, além de forçar você a entender como as alterações com ponteiros afetam os dados. Quando terminar, se quiser, posso te ajudar a revisar o código e discutir o que foi feito.