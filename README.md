# Revisão

## Instruções

1. Resolva as questões em seus respectivos arquivos. Ex: ``q1/q1.go``
2. Utilize o arquivo `main.go` para testar suas funções.
3. Ao finalizar, não esqueça de fazer o commit e o push.

## Questão 1

Você está trabalhando em um projeto de gerenciamento de uma escola. O sistema precisa armazenar informações sobre os
alunos, incluindo seu nome, idade e as matérias em que estão matriculados, juntamente com suas respectivas notas. Você
decidiu usar structs e mapas para representar essas informações.

No entanto, você descobriu que existem dois conjuntos de dados diferentes sobre os alunos. Cada conjunto de dados é
representado por um mapa. O mapa "studentData1" contém informações sobre as notas dos alunos para a primeira metade do
semestre, enquanto o mapa "studentData2" contém informações sobre as notas para a segunda metade do semestre.

Sua tarefa é criar uma função chamada "mergeStudentData" que recebe os mapas "studentData1" e "studentData2" como
parâmetros e retorna um novo mapa que contém as informações combinadas dos dois conjuntos de dados.

O objetivo é combinar as informações de cada aluno, preservando o nome e a idade, e atualizando as matérias e notas de
acordo com o mapa mais recente. Lembre-se de que um aluno pode estar matriculado em diferentes matérias em cada metade
do semestre.

### Exemplo de entrada

```
studentData1 := map[string]Student{
    "John": Student{
        Name: "John",
        Age: 20,
        Subjects: map[string]float64{
            "Math":  8.5,
            "Science": 7.8,
        },
    },
    "Alice": Student{
        Name: "Alice",
        Age: 22,
        Subjects: map[string]float64{
            "History": 9.2,
        },
    },
}

studentData2 := map[string]Student{
    "John": Student{
        Name: "John",
        Age: 20,
        Subjects: map[string]float64{
            "Science": 8.0,
            "English": 7.5,
        },
    },
    "Bob": Student{
        Name: "Bob",
        Age: 21,
        Subjects: map[string]float64{
            "Math": 6.7,
            "Physics": 8.9,
        },
    },
}

result := mergeStudentData(studentData1, studentData2)
fmt.Println(result)
```

### Exemplo de saída:

```
map[Alice:{Alice 22 map[History:9.2]} Bob:{Bob 21 map[Math:6.7 Physics:8.9]} John:{John 20 map[English:7.5 Math:8.5 Science:8]}]
```

### Explicação:

Nesse exemplo, o aluno "John" está matriculado em três matérias diferentes: "Math", "Science" e "English". As notas
dessas matérias são atualizadas de acordo com o mapa mais recente, que é o mapa "studentData2". O aluno "Alice" está
matriculado em duas matérias diferentes: "History" e "Science". Como o mapa "studentData2" não contém informações sobre
a matéria "History", as informações sobre essa matéria são preservadas.

## Questão 2

Você é um desenvolvedor de software em uma empresa financeira e está trabalhando em um sistema de folha de pagamento.
Cada funcionário possui um ID único, nome, cargo, salário base e um conjunto de bônus mensais. Você decidiu usar uma
struct para representar as informações de cada funcionário.

Agora, você precisa implementar uma função chamada "calculateTotalSalary" que recebe como parâmetro um ponteiro para um
objeto do tipo Employee e calcula o salário total do funcionário, considerando o salário base e a soma dos bônus
mensais. Caso a soma dos bônus mensais seja maior que 1500.0, a titulação do funcionário deve ser atualizada com o
prefixo "Senior".

Se o ponteiro para o objeto do tipo Employee for nulo, a função deve retornar um erro.

### Exemplo de entrada

```
employee := &Employee{
    ID:        1,
    Name:      "John Doe",
    Title:     "Software Engineer",
    BaseSalary: 5000.0,
    Bonuses:   []float64{1000.0, 500.0, 750.0},
}

totalSalary := calculateTotalSalary(employee)
fmt.Println(totalSalary)
fmt.Println(employee.Title)
```

### Exemplo de saída:

```
7250
Senior Software Engineer
```

### Explicação:

Neste exemplo, o salário total do funcionário "John Doe" é calculado somando o salário base de 5000.0 com os bônus
mensais de 1000.0, 500.0 e 750.0, resultando em um salário total de 7250.0. Como a soma dos bônus mensais é maior que o
valor de 1500.0, a titulação do funcionário é atualizada para "Senior".

## Questão 3

Você está desenvolvendo um sistema de gerenciamento de estoque para uma loja. Cada produto possui um código único, nome,
preço unitário e quantidade em estoque. Você decidiu usar uma struct para representar as informações de cada produto.

Agora, você precisa implementar uma função chamada "updateStock" que recebe como parâmetro um ponteiro para um objeto do
tipo Product e um mapa que representa as vendas realizadas, onde as chaves são os códigos dos produtos vendidos e os
valores são as quantidades vendidas. A função deve atualizar a quantidade em estoque do produto com base nas vendas
realizadas. Caso o ponteiro para o objeto do tipo Product seja nulo, a função deve retornar um erro. Se a quantidade
resultante for menor que zero, a função deve retornar um erro e a quantidade original do produto deve ser mantida.

Sua tarefa é escrever a função "updateStock" e usá-la para atualizar a quantidade em estoque dos produtos vendidos.

### Exemplo de entrada

```
product := &Product{
    Code:     "P001",
    Name:     "Camiseta",
    Price:    29.99,
    Quantity: 100,
}

sales := map[string]int{
    "P001": 20,
    "P002": 10,
}

updateStock(product, sales)
fmt.Println(product.Quantity)
```

### Exemplo de saída:

```
80
```

### Explicação:

Neste exemplo, o produto "Camiseta" tem uma quantidade inicial em estoque de 100. Com base no mapa sales, que indica que
20 unidades do produto "P001" foram vendidas, a função updateStock atualiza a quantidade em estoque para 80.

## Questão 4

Você está construindo um sistema de gerenciamento de estudantes em uma escola. Cada estudante possui um ID único, nome,
notas em diferentes disciplinas e uma média geral. Você decidiu usar uma struct para representar as informações de cada
estudante.

Agora, você precisa implementar uma função chamada "updateAverage" que recebe como parâmetro um mapa onde as chaves são
os IDs dos estudantes e os valores são ponteiros para objetos do tipo Student. A função deve atualizar a média geral de
cada estudante com base em suas notas atualizadas.

Sua tarefa é escrever a função "updateAverage" e usá-la para atualizar a média geral de cada estudante no mapa
fornecido.

### Exemplo de entrada

```
students := map[int]*Student{
    1: &Student{
        ID:       1,
        Name:     "John Doe",
        Grades:   []float64{8.5, 7.0, 9.2},
        Average:  0.0,
    },
    2: &Student{
        ID:       2,
        Name:     "Jane Smith",
        Grades:   []float64{6.8, 7.5, 8.0},
        Average:  0.0,
    },
}

updateAverage(students)
fmt.Println(students[1].Average)
fmt.Println(students[2].Average)
```

### Exemplo de saída:

```
8.2333333
7.4333333
```

### Explicação:

Após a chamada da função updateAverage, a média geral de cada estudante no mapa students deve ser atualizada com base
nas notas presentes em Grades. A média geral do estudante com ID 1 é 8.2333333 e a média geral do estudante com ID 2 é
7.4333333.

## Questão 5

Você está desenvolvendo um sistema de gerenciamento de vendas para uma loja. Cada venda possui um código único, data,
valor total e uma lista de produtos vendidos. Cada produto tem um código, nome e preço unitário. Você decidiu usar
structs para representar as informações de cada venda e de cada produto.

Agora, você precisa implementar uma função chamada "calculateTotalSales" que recebe como parâmetro um mapa onde as
chaves são os códigos das vendas e os valores são ponteiros para objetos do tipo Sale. A função deve calcular o valor
total de todas as vendas.

Sua tarefa é escrever a função "calculateTotalSales" e usá-la para calcular o valor total de todas as vendas no mapa
fornecido.

### Exemplo de entrada

```
sale1 := &Sale{
    Code: "S001",
    Date: "2023-05-01",
    Products: []*Product{
        &Product{Code: "P001", Name: "Product 1", Price: 50.0},
        &Product{Code: "P002", Name: "Product 2", Price: 25.0},
        &Product{Code: "P003", Name: "Product 3", Price: 25.0},
    },
}

sale2 := &Sale{
    Code: "S002",
    Date: "2023-05-05",
    Products: []*Product{
        &Product{Code: "P004", Name: "Product 4", Price: 25.0},
        &Product{Code: "P005", Name: "Product 5", Price: 50.0},
    },
}

sales := map[string]*Sale{
    "S001": sale1,
    "S002": sale2,
}

total := calculateTotalSales(sales)
fmt.Println(total)
```

### Exemplo de saída:

```
250.0
```

### Explicação:

Neste exemplo, temos duas vendas no mapa sales, cada uma com seu valor total e uma lista de produtos vendidos. A função
calculateTotalSales percorre as vendas e calcula o valor total de todas as vendas, retornando o resultado.

## Questão Bônus

Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela. Caso não
exista cidade de destino, retorne um erro.

### Exemplo de entrada

```
caminhos = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
```

### Exemplo de saída:

```
Sao Paulo
```

### Explicação:

Começando em "London", você vai para "New York", depois para "Lima" e depois para "Sao Paulo". Nenhuma cidade tem um
caminho de saída para a cidade de destino, então o destino é "Sao Paulo".
