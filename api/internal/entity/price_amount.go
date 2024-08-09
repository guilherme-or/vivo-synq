/*
 * Products
 *
 * Retorna os dados de posse de produtos de um cliente.  # Definições Está api retorna os produtos contratados de um cliente. Na Vivo existem várias opções de produtos, dentre eles se destacam, a telefonia móvel 5G, a internet de fibra, a TV por assinatura, produtos de casa inteligente e muitos outros produtos de parceiros, como netflix, amazon prime, que podem ser contratados como um pacote pelo cliente.      https://www.vivo.com.br/
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package entity

// Valor pago.
type PriceAmount struct {
	// Amount value
	Value float32 `json:"value"`
}
