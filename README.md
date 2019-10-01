# MercadoPagoSDK

SDK para o gateway de pagamento [MercadoPago](https://www.mercadopago.com.br/developers).

##### Endpoints implementados
- [ ] [Advanced Payments](https://www.mercadopago.com.br/developers/pt/reference/advanced_payments/_advanced_payments/post/)
- [x] [Cartões](https://www.mercadopago.com.br/developers/pt/reference/cards/_customers_customer_id_cards/get/)
- [ ] [Estornos](https://www.mercadopago.com.br/developers/pt/reference/chargebacks/_chargebacks_id/get/)
- [x] [Clientes](https://www.mercadopago.com.br/developers/pt/reference/customers/_customers/post/)
- [x] [Tipos de documentos](https://www.mercadopago.com.br/developers/pt/reference/identification_types/_identification_types/get/)
- [x] [Notas fiscais](https://www.mercadopago.com.br/developers/pt/reference/invoices/_invoices_id/get/)
- [ ] [Ordens](https://www.mercadopago.com.br/developers/pt/reference/merchant_orders/_merchant_orders/post/)
- [x] [Meios de pagamentos](https://www.mercadopago.com.br/developers/pt/reference/payment_methods/_payment_methods/get/)
- [x] [Pagamentos](https://www.mercadopago.com.br/developers/pt/reference/payments/_payments/post/)
- [x] [Preferências](https://www.mercadopago.com.br/developers/pt/reference/preferences/_checkout_preferences/post/)

### Utilizando

Você pode consultar a documentação [clicando aqui](https://godoc.org/github.com/seriallink/mercadopago).

Também é possível utilizar o exemplo abaixo como ponto inicial.

```go
import (
	"github.com/seriallink/mercadopago"
)

const clientId = ""
const clientSecret = ""
const publicKey = ""
const accessToken = "ACCESS_TOKEN"

func main() {
	client := mp.NewClient(clientId, clientSecret, publicKey, accessToken)

	pms, err := client.GetPaymentMethods()
	if err != nil {
		fmt.Println(err)
	}

	for _, pm := range pms {
		fmt.Println(pm.Name, pm.PaymentId)
	}
}
```

*Observações: Em muitos casos é necessário apenas o `accessToken`, podendo ser definido o `clientId` e `clientSecret` strings vazias, caso não possua um `accessToken` pode conseguir o ele utilizando o método `client.Authorize`. [Para saber mais consulte o código de uma das SDK oficiais](https://github.com/mercadopago/dx-php/blob/6ab8a49cd259e9e8eae0119b6f8dbe9b1a5aed56/src/MercadoPago/Config.php#L116).*





