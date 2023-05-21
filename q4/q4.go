package q4

//Você está construindo um sistema de gerenciamento de estudantes em uma escola. Cada estudante possui um ID único, nome,
//notas em diferentes disciplinas e uma média geral. Você decidiu usar uma struct para representar as informações de cada
//estudante.
//
//Agora, você precisa implementar uma função chamada "updateAverage" que recebe como parâmetro um mapa onde as chaves são
//os IDs dos estudantes e os valores são ponteiros para objetos do tipo Student. A função deve atualizar a média geral de
//cada estudante com base em suas notas atualizadas.
//
//Sua tarefa é escrever a função "updateAverage" e usá-la para atualizar a média geral de cada estudante no mapa
//fornecido.

type Student struct {
	ID      int
	Name    string
	Grades  []float64
	Average float64
}

func UpdateAverage(students map[int]*Student) {
	// Seu código aqui
}
