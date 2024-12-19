package config

var (
	// StringConexaoBanco é a string de conexão com o MySQL
	StringConexaoBanco = "root:admin@/devbook?charset=utf8&parseTime=True&loc=Local"
	// Porta onde a API vai estar rodando
	Porta = "5000"
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {

}