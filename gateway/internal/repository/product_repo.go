package repository

type ProductRepository interface {
	Find(userID string) ([]byte, error)           // Busca os produtos de um usuário por seu ID e retorna o objeto JSON em bytes
	FindInCache(userID string) ([]byte, error)    // Busca os produtos de um usuário por seu ID no cache e retorna o objeto JSON em bytes
	SaveInCache(userID string, data []byte) error // Salva os produtos de um usuário no cache
}
