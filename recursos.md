scopo:recurso:ação

scopo = usuario, organização etc
	user-{}
recurso = funcionalidade (ex: membros, cobrancas, etc)
ação = CRUD + especificas

constantes
{userId} - usuario id (deve ser modificado de acordo com o id do usuario)
{organizationId} - organização id (deve ser modificado de acordo com o id da organização)
* - tudo

# Padrão
- Por padrão existem algumas roles, elas devem ser criadas na criação da organização


# criação
- usuario especifica os recursos que quer ter acesso, 