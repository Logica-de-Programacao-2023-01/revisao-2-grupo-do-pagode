package q1

//Você está trabalhando em um projeto de gerenciamento de uma escola. O sistema precisa armazenar informações sobre os alunos, incluindo seu nome, idade e as matérias em que estão matriculados, juntamente com suas respectivas notas. Você decidiu usar structs e mapas para representar essas informações.
//
//No entanto, você descobriu que existem dois conjuntos de dados diferentes sobre os alunos. Cada conjunto de dados é representado por um mapa. O mapa "studentData1" contém informações sobre as notas dos alunos para a primeira metade do semestre, enquanto o mapa "studentData2" contém informações sobre as notas para a segunda metade do semestre.
//
//Sua tarefa é criar uma função chamada "mergeStudentData" que recebe os mapas "studentData1" e "studentData2" como parâmetros e retorna um novo mapa que contém as informações combinadas dos dois conjuntos de dados.
//
//O objetivo é combinar as informações de cada aluno, preservando o nome e a idade, e atualizando as matérias e notas de acordo com o mapa mais recente. Lembre-se de que um aluno pode estar matriculado em diferentes matérias em cada metade do semestre.

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	// Seu código aqui
	return nil
}
