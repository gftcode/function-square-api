# Documentação Técnica - Function Square API

Documentação detalhada sobre a arquitetura e funcionamento de cada módulo do projeto.

## 📚 Índice

1. [Arquitetura Geral](#arquitetura-geral)
2. [Módulos](#módulos)
3. [Fluxo de Execução](#fluxo-de-execução)
4. [Componentes Detalhados](#componentes-detalhados)

---

## 🏗️ Arquitetura Geral

O projeto segue uma arquitetura em camadas:

```
┌─────────────────────────────────────┐
│         cmd/app/main.go             │  Entry Point
│      (Orquestração principal)        │
└────────────┬────────────────────────┘
             │
     ┌───────┴────────┐
     │                │
┌────▼──────┐    ┌────▼──────────────┐
│   CLI      │    │    Plotter        │
│(Input)     │    │(Visualization)    │
└────┬──────┘    └────┬───────────────┘
     │                │
     └───────┬────────┘
             │
      ┌──────▼────────┐
      │   UseCase     │  Lógica de Negócio
      │(Calculations) │
      └──────┬────────┘
             │
      ┌──────▼────────┐
      │    Domain     │  Domínio
      │   (Formula)   │
      └───────────────┘
```

---

## 📦 Módulos

### 1. **cmd/app/main.go** - Ponto de Entrada

**Responsabilidade**: Orquestar o fluxo principal da aplicação

```go
func main() {
    equation := cli.NewCli().SetValues()      // 1. Obter entrada
    printer := plotter.NewTerminalPlotter()   // 2. Criar plotter
    printer.Plot(equation)                    // 3. Renderizar
    cli.ExportResults(equation)               // 4. Exportar
}
```

**O que faz:**

- Cria a instância do CLI para entrada de dados
- Recebe os coeficientes (a, b, c)
- Instancia o plotter para visualização
- Chama a exportação de resultados

---

### 2. **internal/infrastructure/entrypoint/cli/** - Interface CLI

#### **cli.go** - Entrada de Dados

**Tipo**: Struct `Output`

**Responsabilidades:**

- Validar e coletar coeficientes do usuário
- Confirmar dados antes de processar
- Criar instância da equação

**Métodos:**

```go
NewCli() *Output                          // Construtor
SetValues() usecase.QuadraticFunction    // Coleta e retorna equação
GetValues() (int, int, int)              // Retorna a, b, c
input()                                  // Input com validação
```

**Fluxo de Input:**

1. Solicita valores de a, b, c
2. Exibe a equação formatada
3. Pede confirmação do usuário
4. Repete se não confirmado

#### **fileCli.go** - Exportação de Resultados

**Função**: `ExportResults(equation usecase.QuadraticFunction) error`

**Responsabilidades:**

- Extrair resultados da equação
- Formatar dados para arquivo
- Escrever no arquivo de saída

**Dados Exportados:**

- Fórmulas da função
- Delta (Δ)
- Raízes (X' e X'')
- Vértices (V)

---

### 3. **internal/usecase/** - Lógica de Negócio

**Tipo**: Interface `QuadraticFunction` e struct `Equation`

**Responsabilidades:**

- Implementar cálculos matemáticos
- Armazenar dados da equação
- Fornecer métodos de acesso aos resultados

**Métodos Principais:**

```go
FindDelta() float64           // Calcula Δ = b² - 4ac
Xrows() (x1, x2 float64)      // Calcula raízes X' e X''
FindVertices() (vx, vy float64) // Calcula vértice (-b/2a, -Δ/4a)
```

**Equação Quadrática:**

- Formula geral: ax² + bx + c = 0
- Delta: Δ = b² - 4ac
- Raízes: x = (-b ± √Δ) / 2a
- Vértice: V = (-b/2a, -Δ/4a)

---

### 4. **internal/domain/formula/** - Domínio de Fórmulas

**Tipo**: Struct `Formula`

**Responsabilidades:**

- Armazenar e fornecer fórmulas matemáticas
- Documentação de cálculos

**Métodos:**

```go
NewFormula() *Formula
GetFormulas() string  // Retorna texto com fórmulas
```

**Conteúdo:**

- Fórmula geral da parábola
- Fórmula do delta
- Fórmula das raízes
- Fórmula do vértice

---

### 5. **internal/infrastructure/plotter/** - Visualização

#### **plotter.go** - Orquestrador de Plotagem

**Tipo**: Struct `TerminalPlotter`

**Responsabilidades:**

- Coordenar a renderização do gráfico
- Gerenciar canvas (tela de desenho)

**Métodos:**

```go
NewTerminalPlotter(width, height int) *TerminalPlotter
Plot(equation usecase.QuadraticFunction)  // Desenha a parábola
```

#### **canvas.go** - Área de Desenho

**Tipo**: Struct `Canvas`

**Responsabilidades:**

- Representar a grade de pontos
- Fornecer interface para desenho

**Métodos:**

```go
NewCanvas(width, height int) *Canvas
SetPoint(x, y int, char string)  // Define um caractere na posição
Print()                          // Exibe no terminal
```

#### **drawing.go** - Primitivas de Desenho

**Funções:**

- Desenhar linhas (eixos X e Y)
- Desenhar pontos da parábola
- Traçar curvas

#### **rendering.go** - Renderização

**Responsabilidades:**

- Converter coordenadas matemáticas para coordenadas de pixel
- Renderizar eixos e marcações
- Aplicar escala ao gráfico

#### **graphic_test.go** - Testes de Visualização

- Valida renderização de gráficos
- Testa casos extremos

---

### 6. **pkg/fileutil/** - Utilitários de Arquivo

**Tipo**: Struct `File`

**Responsabilidades:**

- Gerenciar arquivo de exportação
- Criar diretório automaticamente
- Escrever conteúdo no arquivo

**Métodos:**

```go
NewFile() *File                    // Cria/abre arquivo
AddContent(contents ...any) error  // Adiciona conteúdo
open() (*os.File, error)          // Abre arquivo para escrita
```

**Arquivo de Saída:**

- Local: `cmd/exports/function².txt`
- Formato: Texto formatado
- Append: Adiciona conteúdo ao final

---

## 🔄 Fluxo de Execução

```
1. MAIN
   │
   ├─► CLI.SetValues()
   │   │
   │   ├─► input()
   │   │   ├─► Solicita a, b, c
   │   │   ├─► Exibe equação
   │   │   └─► Confirma dados
   │   │
   │   └─► NewEquation(a, b, c)
   │       └─► Cria instância da equação
   │
   ├─► NewTerminalPlotter(45, 20)
   │   └─► Cria canvas 45x20
   │
   ├─► printer.Plot(equation)
   │   │
   │   ├─► rendering.PrepareCoordinates()
   │   │   └─► Converte valores matemáticos para pixels
   │   │
   │   ├─► drawing.DrawAxes()
   │   │   └─► Desenha eixos X e Y
   │   │
   │   ├─► drawing.DrawParabola()
   │   │   └─► Desenha a curva da parábola
   │   │
   │   └─► canvas.Print()
   │       └─► Exibe no terminal
   │
   └─► ExportResults(equation)
       │
       ├─► fileutil.NewFile()
       │   └─► Cria/abre arquivo de exportação
       │
       ├─► Extrai: delta, raízes, vértices
       │
       └─► fileutil.AddContent(...)
           └─► Escreve no arquivo
```

---

## 💻 Componentes Detalhados

### Entrada de Dados (CLI)

**Fluxo:**

```
User Input → Validação → Struct Output → NewEquation → UseCase
     ↓
  Loop até
 confirmação
```

**Validação:**

- Valores numéricos de a, b, c
- Confirmação do usuário
- Repetição se negado

---

### Cálculos (UseCase)

**Equação**: ax² + bx + c

**Fórmulas Implementadas:**

1. **Delta (Δ)**

   ```
   Δ = b² - 4ac
   ```

2. **Raízes (Bhaskara)**

   ```
   x = (-b ± √Δ) / 2a
   x' = (-b + √Δ) / 2a
   x'' = (-b - √Δ) / 2a
   ```

3. **Vértice**
   ```
   Xv = -b / 2a
   Yv = -Δ / 4a
   ```

---

### Visualização (Plotter)

**Canvas:**

- Dimensões: 45x20 (customizável)
- Caracteres: ASCII padrão
- Escala: Automática baseada no intervalo

**Renderização:**

- Eixo X (horizontal)
- Eixo Y (vertical)
- Pontos da parábola (usando \*ou o)
- Marcações numéricas (opcional)

---

### Exportação (FileUtil)

**Arquivo de Saída:**

```
cmd/exports/function².txt
```

**Formato:**

```
[Fórmulas]
----------------------------------
Delta: {valor}
X' and X'': {x1, x2}
Vertice: {vx, vy}
```

---

## 🧪 Testes

### Arquivos de Teste:

- `internal/domain/formula/formula_test.go`
- `internal/infrastructure/entrypoint/cli/cli_test.go`
- `internal/infrastructure/plotter/graphic_test.go`
- `internal/usecase/quadratic_function_test.go`

**Executar Testes:**

```bash
go test ./...
```

---

## 🔧 Configurações e Constantes

### Plotter

- **Largura padrão**: 45 caracteres
- **Altura padrão**: 20 linhas

### FileUtil

- **Nome arquivo**: `function².txt`
- **Diretório**: `cmd/exports/`
- **Modo**: Append (adiciona ao final)

---

## 📝 Notas de Implementação

1. **Type Safety**: Uso de tipos específicos (int para coeficientes, float64 para cálculos)
2. **Error Handling**: Funções de arquivo retornam erro
3. **Modularização**: Cada responsabilidade em seu módulo
4. **Reutilização**: Canvas e Drawing podem ser usados para outros gráficos
5. **Testabilidade**: Interfaces permitem fácil mock e teste

---

## 🚀 Possíveis Extensões

1. Plotar múltiplas funções simultaneamente
2. Exportar em formatos diferentes (CSV, JSON)
3. Interface gráfica (GUI)
4. Cálculos simbólicos
5. Histórico de equações
6. Integração com gnuplot ou similar

---

**Última atualização**: 2024
**Versão do Projeto**: 1.0.0
