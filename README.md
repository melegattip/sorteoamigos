# Sorteoamigos

API Para sortear nombres para amigo invisible.

Contiene 2 metodos principales para:
1) /people : cargar los nombres de las personas
2) /grid : muestra el sorteo de cada nombre

# Ejemplo cURLs:
# 1) 
curl --location --request POST 'http://localhost:8080/people' \
--header 'Content-Type: application/json' \
--data-raw '[
	{"name": "Alice"},
  {"name": "Bob"}, 
  {"name": "Pablo"}
]'
## Response : 
{
    "message": "3 personas añadidas"
}

# 2) 
curl --location --request GET 'http://localhost:8080/grid'
## Response:
[
    {
        "Entrega": "Pablo",
        "Recibe": "Bob"
    },
    {
        "Entrega": "Bob",
        "Recibe": "Alice"
    },
    {
        "Entrega": "Alice",
        "Recibe": "Pablo"
    }
]
