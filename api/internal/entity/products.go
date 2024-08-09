/*
 * Products
 *
 * Retorna os dados de posse de produtos de um cliente.  # Definições Está api retorna os produtos contratados de um cliente. Na Vivo existem várias opções de produtos, dentre eles se destacam, a telefonia móvel 5G, a internet de fibra, a TV por assinatura, produtos de casa inteligente e muitos outros produtos de parceiros, como netflix, amazon prime, que podem ser contratados como um pacote pelo cliente.      https://www.vivo.com.br/
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package entity

import (
	"time"
)

// Object that models a  product
type Products struct {
	// Indetificador unico do produto.
	Id string `json:"id"`
	// Representa o estado do produto. -
	// **'active'** produto é ativo e pronto para uso. -
	// **'activating'** o produto foi contratado mas está
	// em processo de ativação e criando os recursos necessários
	// para ser utilizado pelo cliente. -
	// **'suspended'** produto foi suspenso por falta de
	// pagamento ou outro problema legal. -
	// **'cancelled'** Produto foi cancelado pelo cliente.
	Status string `json:"status"`
	// Lista de identificadores dos clientes associados ao produto,
	// pode ser uma linha telefônica, um identificador de banda larga,
	// linha fixa, ou em alguns casos até o CPF do cliente.
	Identifiers []string `json:"identifiers"`
	// Data de contratação do produto , no formato ISO-8601.
	StartDate time.Time `json:"start_date"`
	// Data de término da assinatura do produto, no formato ISO-8601.
	EndDate time.Time `json:"end_date,omitempty"`
	// Lista de preços pagos pelo cliente.
	Prices []Price `json:"prices,omitempty"`
	// Array of products objects. Only applies for product bundle
	SubProducts []Product `json:"sub_products,omitempty"`
	// Nome comercial do produto.
	ProductName      string            `json:"product_name"`
	ProductType      *ProductType      `json:"product_type"`
	Descriptions     *[]Description    `json:"descriptions,omitempty"`
	SubscriptionType *SubscriptionType `json:"subscription_type,omitempty"`
	// Tags de identificação do produto.
	Tags []string `json:"tags,omitempty"`
}
