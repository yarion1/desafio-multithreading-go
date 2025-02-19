# Desafio de Busca de CEP com Multithreading e APIs

Este projeto é uma solução para um desafio que envolve a busca de informações de endereço a partir de um CEP utilizando duas APIs distintas: **ViaCEP** e **BrasilAPI**. O objetivo é fazer requisições simultâneas às duas APIs e retornar o resultado da que responder mais rapidamente, descartando a resposta mais lenta.

## Requisitos do Desafio

- Fazer requisições simultâneas para as seguintes APIs:
  - **ViaCEP**: [https://viacep.com.br/ws/{cep}/json/](https://viacep.com.br/ws/{cep}/json/)
  - **BrasilAPI**: [https://brasilapi.com.br/api/cep/v1/{cep}](https://brasilapi.com.br/api/cep/v1/{cep})
- Retornar o resultado da API que responder mais rápido.
- Exibir no terminal os dados do endereço e qual API forneceu a resposta.
- Limitar o tempo de resposta a **1 segundo**. Caso nenhuma API responda dentro desse tempo, exibir uma mensagem de erro.

