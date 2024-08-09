/*
 * Products
 *
 * Retorna os dados de posse de produtos de um cliente.  # Definições Está api retorna os produtos contratados de um cliente. Na Vivo existem várias opções de produtos, dentre eles se destacam, a telefonia móvel 5G, a internet de fibra, a TV por assinatura, produtos de casa inteligente e muitos outros produtos de parceiros, como netflix, amazon prime, que podem ser contratados como um pacote pelo cliente.      https://www.vivo.com.br/
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package entity

// SubscriptionType : Tipo de assinatura
// vinculada no produto (prepaid, postpaid, etc)
type SubscriptionType string

// List of SubscriptionType
const (
	PREPAID  SubscriptionType = "prepaid"
	POSTPAID SubscriptionType = "postpaid"
	CONTROL  SubscriptionType = "control"
)
