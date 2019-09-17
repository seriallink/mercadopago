package mp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createPreference(t *testing.T) (*Preference, error) {
	t.Helper()

	var items []Item
	items = append(items, Item{
		Title:      "Teste de produto",
		Quantity:   1,
		UnitPrice:  10.50,
		CurrencyId: BrazilianReal,
	})
	preference := &Preference{
		Items: items,
	}
	err := client.CreatePreference(preference)
	return preference, err
}

func TestCreatePreference(t *testing.T) {
	preference, err := createPreference(t)
	assert.NoError(t, err)
	assert.NotEmpty(t, preference.PreferenceId)
}

func TestGetPreference(t *testing.T) {
	pref, err := createPreference(t)
	if err != nil {
		t.FailNow()
	}

	preference := &Preference{
		PreferenceId: pref.PreferenceId,
	}
	err = client.GetPreference(preference)
	assert.NoError(t, err)
	assert.NotEmpty(t, preference.PreferenceId)
	assert.Equal(t, pref.PreferenceId, preference.PreferenceId)
	assert.Len(t, preference.Items, 1)
	assert.Equal(t, pref.Items[0].Title, preference.Items[0].Title)
	assert.Equal(t, pref.Items[0].UnitPrice, preference.Items[0].UnitPrice)
}

func TestSavePreference(t *testing.T) {
	pref, err := createPreference(t)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, "", pref.NotificationUrl)

	preference := &Preference{
		PreferenceId:    pref.PreferenceId,
		NotificationUrl: "https://example.com/link-de-notificacao-alterado",
	}
	err = client.SavePreference(preference)
	assert.NoError(t, err)
	assert.NotEmpty(t, preference.PreferenceId)
	assert.Equal(t, "https://example.com/link-de-notificacao-alterado", preference.NotificationUrl)
}

func ExampleCreatePreference() {
	client := NewClient("", "", "", accessToken)

	preference := &Preference{
		Items: []Item{
			{
				Title:     "Meu produto",
				Quantity:  1,
				UnitPrice: 75.56,
			},
		},
	}
	err := client.CreatePreference(preference)
	if err != nil {
		panic(err)
	}

	fmt.Println(preference.InitPoint)
}

func ExampleCreatePreferenceAdvanced() {
	client := NewClient("", "", "", accessToken)

	backUrls := make(map[string]string)
	backUrls["success"] = "https://example.com/compra-sucesso"
	backUrls["pending"] = "https://example.com/compra-em-andamento"
	backUrls["failure"] = "https://example.com/compra-falhou"

	preference := &Preference{
		Items: []Item{
			{
				Title:     "Meu produto 1",
				Quantity:  1,
				UnitPrice: 75.56,
			},
			{
				Title:     "Meu produto 2",
				Quantity:  5,
				UnitPrice: 12.34,
			},
		},
		BackUrls:          backUrls,
		NotificationUrl:   "https://example.com/minha-url-de-notificacao",
		AdditionalInfo:    "Informações adicionais",
		ExternalReference: "Minha referência externa",
		PaymentMethods: &PaymentMethods{
			DefaultPaymentMethodId: "bolbradesco",
			ExcludedPaymentTypes: []ExcludedPaymentTypes{
				{
					Id: CreditCard,
				},
				{
					Id: DebitCard,
				},
				{
					Id: PrepaidCard,
				},
				{
					Id: ATM,
				},
				{
					Id: BankTransfer,
				},
			},
		},
	}
	err := client.CreatePreference(preference)
	if err != nil {
		panic(err)
	}

	fmt.Println(preference.InitPoint)
}
