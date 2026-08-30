# Function Square API

Uma aplicação em Go para análise e visualização de funções quadráticas (parábolas). O projeto calcula as raízes, o delta, os vértices e gera um gráfico ASCII no terminal, exportando os resultados para um arquivo.

## 🎯 Funcionalidades

- **Entrada interativa**: Receba os coeficientes (a, b, c) de uma função quadrática via CLI
- **Cálculos matemáticos**: Delta, raízes (X' e X'') e vértices (V)
- **Visualização**: Gráfico ASCII renderizado no terminal
- **Exportação**: Resultados salvos em arquivo `function².txt`
- **Validação**: Confirmação dos dados antes do processamento

## 📋 Requisitos

- Go 1.26.5 ou superior
- Terminal com suporte a caracteres especiais

## 🚀 Como Usar

### Compilação

```bash
cd cmd/app
go build -o main.exe
```

Ou use a tarefa de compilação do VS Code:

```bash
gcc main.go -o main.exe -g
```

### Execução

```bash
./main.exe
```

A aplicação solicitará os valores dos coeficientes:

```
Digite o valor de a (x²): 1
Digite o valor de b (x): -5
Digite o valor de c (termo constante): 6

Sua equação: 1x² + (-5x) + (6)
Os dados estão corretos? (true/false): true
```

### Saída

O programa exibirá:

1. Um gráfico da parábola no terminal
2. Salvará em `cmd/exports/function².txt`:
   - Fórmulas da função
   - Delta
   - Raízes (X' e X'')
   - Vértices (V)

## 📁 Estrutura do Projeto

```
function-square-api/
├── cmd/
│   └── app/
│       └── main.go              # Ponto de entrada
├── internal/
│   ├── domain/
│   │   └── formula/             # Lógica de fórmulas
│   ├── infrastructure/
│   │   ├── entrypoint/
│   │   │   └── cli/             # Entrada CLI e exportação
│   │   └── plotter/             # Renderização de gráficos
│   └── usecase/                 # Lógica de negócio
├── pkg/
│   └── fileutil/                # Utilitários de arquivo
├── go.mod
└── README.md
```

## 💡 Exemplos de Entrada

### Exemplo 1: Raízes diferentes

```
a = 1, b = -5, c = 6
Equação: x² - 5x + 6
Raízes: 2 e 3
```

### Exemplo 2: Raiz dupla

```
a = 1, b = -4, c = 4
Equação: x² - 4x + 4
Raízes: 2 (dupla)
```

### Exemplo 3: Sem raízes reais

```
a = 1, b = 0, c = 1
Equação: x² + 1
Delta negativo
```

## 👨‍💻 Desenvolvimento

Para mais detalhes sobre a arquitetura e funcionamento de cada módulo, consulte [DOCUMENTATION.md](DOCUMENTATION.md).

---

**Autor**: GFTCode  
**Versão**: 1.0.0
