# API de Controle de Estoque em Go

Bem-vindo à API de Controle de Estoque em Go! Esta API foi desenvolvida para gerenciar inventários de produtos em estoque de forma eficiente e fácil de usar.

## Funcionalidades Principais

- **Cadastro de Produtos:** Adicione novos produtos ao estoque, incluindo informações como nome, descrição, preço e quantidade disponível.
- **Consulta de Produtos:** Consulte detalhes de produtos individuais, incluindo sua disponibilidade e outras informações relevantes.
- **Atualização de Estoque:** Atualize a quantidade disponível de produtos conforme as vendas ou reposições de estoque.
- **Histórico de Movimentação:** Acompanhe o histórico de movimentação de estoque, incluindo vendas, reposições e ajustes.
- **Autenticação e Autorização:** Acesso seguro à API por meio de autenticação de usuários e controle de permissões.

## Como Usar

1. **Instalação:**
   - Certifique-se de ter o Go instalado em seu sistema.
   - Clone este repositório em seu ambiente local.
   - Na raiz do projeto, inicialize o módulo Go com o comando `go mod init nome-do-seu-modulo`.

2. **Configuração:**
   - Configure as variáveis de ambiente necessárias no arquivo `.env`.

3. **Execução:**
   - Para compilar e executar o servidor, utilize o comando `go run main.go`.
   - Acesse a API em `http://localhost:8080` (ou outra porta configurada).


## Tecnologias Utilizadas

- **Go:** Linguagem de programação utilizada para o desenvolvimento da API.
- **JWT:** Para autenticação e autorização de usuários.

## Contribuindo

Sinta-se à vontade para contribuir com melhorias ou correções para esta API. Basta seguir estas etapas:

1. Faça um fork do projeto.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Faça commit das suas mudanças (`git commit -am 'Adicionando nova feature'`).
4. Faça push para a branch (`git push origin feature/nova-feature`).
5. Abra um Pull Request.


