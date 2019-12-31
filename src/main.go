package main

// api a consumir para saber el grado https://api.ssllabs.com/api/v3/analyze?host=google.com
import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
	// "github.com/likexian/whois-parser-go"
)

func main() {

	// Ruta a manejar con Chi
	r := chi.NewRouter()

	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		IndexHandler(w, r)
	})

	r.Get("/prueba", func(w http.ResponseWriter, r *http.Request) {
		searchServer(w, r)
	})
	// Servidor escuchando en el puerto 5000
	http.ListenAndServe(":5000", r)
}

// IndexHandler nos permite manejar la petición a la ruta '/'
// y retornar "hola mundo" como respuesta al cliente.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// devolver un template
	template, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Página no encontrada!")
	} else {
		template.Execute(w, nil)
	}
}

func pruebaHandler(w http.ResponseWriter, r *http.Request) {
	result, err := whois.Whois("google.com")
	if err == nil {
		fmt.Println("resultado")
		fmt.Println(result)
	}

	fmt.Println(err)
}

func searchServer(w http.ResponseWriter, r *http.Request) {
	// devolver un json
	w.Header().Set("Content-Type", "application/json")

	//capturando el valor de la variable "host" por GET
	dominio := r.URL.Query().Get("host")
	fmt.Println("VARIABLE")
	fmt.Println(dominio)
	message := "envio correcto"
	if len(dominio) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		message = "envie host como parametro en la url"
	} else {
		result, err := whois.Whois(dominio)

		//si no hay error se procede a mostrar el resultado
		if err == nil {
			resultParse, _ := whoisparser.Parse(result)

			servidor := models.server{
				country: resultParse.Registrant.Country,
				owner:   resultParse.Registrant.Organization,
			}
			// servidor := server{
			// 	country: resultParse.Registrant.Country,
			// 	owner:   resultParse.Registrant.Organization,
			// }

			fmt.Println("resultado")
			fmt.Println("Host: " + resultParse.Registrant.Name)
			fmt.Println("Country: " + resultParse.Registrant.Country)
			fmt.Println("Owner: " + resultParse.Registrant.Organization)
			//fmt.Println(result)
			// aqui aplicar el parser
			//fmt.Println(result.Registrant.Organization);

			// resp, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=google.com")
			// if err == nil {
			// 	fmt.Println(resp)
			// }
		}

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"message": ` + message + `}`))
	}

}
