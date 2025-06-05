package clase6

/*
Ejercicio 2: Registro de estudiantes (con archivo .txt)
Descripción:
El archivo estudiantes.txt contiene líneas con este formato:


Juan,85
María,92
Pedro,70

Objetivo:
Leer el archivo y calcular:

Promedio general

Estudiante con nota más alta

Cuántos aprobaron (> 60)

Si una línea está mal formateada, usar panic y atraparlo con recover.

Bonus:
Implementar tests para las funciones de promedio y máximo.
Usar defer para cerrar el archivo y/o capturar errores.
*/
