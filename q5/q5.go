package q5

//Você está desenvolvendo um sistema de gerenciamento de vendas para uma loja. Cada venda possui um código único, data,
//valor total e uma lista de produtos vendidos. Cada produto tem um código, nome e preço unitário. Você decidiu usar
//structs para representar as informações de cada venda e de cada produto.
//
//Agora, você precisa implementar uma função chamada "calculateTotalSales" que recebe como parâmetro um mapa onde as
//chaves são os códigos das vendas e os valores são ponteiros para objetos do tipo Sale. A função deve calcular o valor
//total de todas as vendas.
//
//Sua tarefa é escrever a função "calculateTotalSales" e usá-la para calcular o valor total de todas as vendas no mapa
//fornecido.

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Sale struct {
	Code     string
	Products []*Product
}

func CalculateTotalSales(sales map[string]*Sale) float64 {
	// Seu código aqui
	return 0
}
