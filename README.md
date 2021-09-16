# Seminario_Golang

Una vez creada la carpeta , archivos y leer el enunciado nos disponemos a analizar el problema
Luego se crea la carpeta model y el archivo characters donde se va a crear el objeto Result y su comportamiento
Se crea el constructor
Luego se escribe la funcion CrearREsult que recibe una cadena y la transforma en la cadena solicitada
A su vez se crean las funciones:
            lon (que analiza si la cantidad de caracteres ingresados se corresponden con el valor numerico que se ingreso
            compruebaTipo (analiza si los dos primeros caracteres son TX o NN
Las comprobaciones anteriormente detalladas, se tienen en cuenta en la funcion CrearResult y , en caso que no se cumplan, se devuelve un objeto Result vacio.

Posteriormente, se comienza con la prueba unitaria.
Para esto, se crea un archivo TestCrearResult y se agrega el codigo de test que se encuentra en el enunciado.
Al llegar a este punto, da un error, porque no se encuentra la funcion ParseString. Luego de un tiempo, y de un detallado analisis, 
nos damos cuenta que no es ParseString, sino CrearResult, la funcion a la que se debe llamar.
una vez descubierto esto se ingresa por consola  go test ./... -count=1 -cover para determinar la cobertura del testing sobre el codigo, y se observa que el cover
es de un 90,5% 
Finalmente se ingresa por consola go test ./... -coverprofile=out.test y go tool cover -html out.test -o out.html para generar el archivo de salida
del test con extension .html, para poder visualizar que codigo es el que se analiza en el test y cual no.
Viendo que el main no es tenido en cuenta en el test, pero si el resto del codigo, y el programa cumple su funcion, se da por finalizado el trabajo practico.



