from cieloApi3 import *

import json

# Configure o ambiente
environment = Environment(sandbox=True)

# Configure seu merchant, para gerar acesse: https://cadastrosandbox.cieloecommerce.cielo.com.br/
merchant = Merchant('cab433b8-f168-48cc-ae0a-807d509e76a2', 'UHYVBLKJTQDBXJTKQVLDQQTNNSGQXWAZQNCIRMDG')

# Crie uma instância de Sale informando o ID do pagamento
sale = Sale('123')

# Crie uma instância de Customer informando o nome do cliente
sale.customer = Customer('Fulano de Tal')

# Crie uma instância de Credit Card utilizando os dados de teste
# esses dados estão disponíveis no manual de integração
credit_card = CreditCard('123', 'Visa')
credit_card.expiration_date = '12/2018'
credit_card.card_number = '0000000000000001'
credit_card.holder = 'Fulano de Tal'

# Crie uma instância de Payment informando o valor do pagamento
sale.payment = Payment(15700)
sale.payment.credit_card = credit_card

# Cria instância do controlador do ecommerce
cielo_ecommerce = CieloEcommerce(merchant, environment)

# Criar a venda e imprime o retorno
response_create_sale = cielo_ecommerce.create_sale(sale)
print (response_create_sale)

# Com a venda criada na Cielo, já temos o ID do pagamento, TID e demais
# dados retornados pela Cielo
payment_id = sale.payment.payment_id

# Com o ID do pagamento, podemos fazer sua captura,
# se ela não tiver sido capturada ainda
response_capture_sale = cielo_ecommerce.capture_sale(payment_id, 15700, 0)

# E também podemos fazer seu cancelamento, se for o caso
response_cancel_sale = cielo_ecommerce.cancel_sale(payment_id, 15700)