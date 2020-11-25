package main

import (
    "fmt"
    "net/rpc"
    "math"
    "strconv"
)

func calificionMateria(client *rpc.Client) {
    var nombre string
    fmt.Print("Ingresa el nombre del alumno: ")
    fmt.Scanln(&nombre)
    var materia string
    fmt.Print("Ingresa el nombre de la materia: ")
    fmt.Scanln(&materia)
    var cal float64
    fmt.Print("Ingresa la calificion: ")
    fmt.Scanln(&cal)
    var strings []string
    strings = append(strings, nombre)
    strings = append(strings, materia)
    strings = append(strings, strconv.FormatFloat(cal, 'f', 6, 64))

    var reply string
    client.Call("SERVER.CalificacionMateria", strings, &reply)
    fmt.Println(reply)
}

func promedioAlumno(client *rpc.Client) {
    var nombre string
    fmt.Print("Ingresa el nombre del alumno: ")
    fmt.Scanln(&nombre)

    var prom float64
    client.Call("SERVER.PromedioAlumno", nombre, &prom)
    if !math.IsNaN(prom) {
        fmt.Println("Promedio: ", prom)
    } else {
        fmt.Println("No se pudo obtener el promedio.")
    }
}

func promedioGeneral(client *rpc.Client) {
    var prom float64
    client.Call("SERVER.PromedioGeneral", "", &prom)

    if !math.IsNaN(prom) {
        fmt.Println("Promedio general: ", prom)
    } else {
        fmt.Println("No se pudo obtener el promedio.")
    }
}

func promedioMateria(client *rpc.Client) {
    var nombre string
    fmt.Print("Ingresa el nombre de la materia: ")
    fmt.Scanln(&nombre)

    var prom float64
    client.Call("SERVER.PromedioMateria", nombre, &prom)

    if !math.IsNaN(prom) {
        fmt.Println("Promedio de la materia: ", prom)
    } else {
       fmt.Println("No se pudo obtener el promedio.")
    }
}

func main() {
    client, err := rpc.Dial("tcp", "127.0.0.1:4040")

    if err != nil {
        fmt.Println("Error en conexion: ", err)
    }

    var op int64
    for {
        fmt.Println("1) Agregar la calificacion de un alumno por materia")
        fmt.Println("2) Obtener el promedio del alumno")
        fmt.Println("3) Obtener el promedio de todos los alumnos")
        fmt.Println("4) Obtener el promedio por materia")
        fmt.Println("5) Salir")
        fmt.Scanln(&op)

        switch op {
        case 1:
            calificionMateria(client)
        case 2:
            promedioAlumno(client)
        case 3:
            promedioGeneral(client)
        case 4:
            promedioMateria(client)
        case 5:
            return
        }
    }
}
