package main

import (
    "fmt"
    "net"
    "net/rpc"
    "strconv"
)

type SERVER int
var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)

func (a *SERVER) GetMaterias(empty string, mat *map[string]map[string]float64) error {
    *mat = materias
    return nil
}

func (a *SERVER) SetMaterias(act *map[string]map[string]float64, mat *map[string]map[string]float64) error {
    materias = *act
    return nil
}

func (a *SERVER) GetAlumnos(empty string, alu *map[string]map[string]float64) error {
    *alu = alumnos
    return nil
}

func (a *SERVER) SetAlumnos(act *map[string]map[string]float64, alu *map[string]map[string]float64) error {
    alumnos = *alu
    return nil
}

func (a *SERVER) CalificacionMateria(datos []string, reply *string) error {
    nombre := datos[0]
    materia := datos[1]
    cal, _ := strconv.ParseFloat(datos[2], 64)
    alumno := make(map[string]float64)

    if _, ok := alumnos[nombre][materia]; ok {
        *reply = "Error, alumno ya calificado"
    } else {
        alumno[nombre] = cal
        if _, ok := materias[materia] ; ok{
            materias[materia][nombre] = cal
        } else {
            materias[materia] = make(map[string]float64)
            materias[materia] = alumno
        }

        if _, ok := alumnos[nombre] ; ok{
            alumnos[nombre][materia] = cal
        } else {
            materiaM := make(map[string]float64)
            alumnos[nombre] = make(map[string]float64)
            materiaM[materia] = cal
            alumnos[nombre] = materiaM
        }


        *reply = "Se califico al alumno de manera correcta"
    }

    ImprimeMaps()
    return nil
}

func ImprimeMaps() {
    fmt.Println("\n\tMATERIAS:\n", materias)
    fmt.Println("\n\n\tAlumnos:\n",alumnos)
}

func (a *SERVER) PromedioAlumno(alumno string, prom *float64) error {
    var cont float64
    var totalCal float64

    for key, element := range alumnos {
        if key == alumno {
            for _, cal := range element {
                totalCal += cal
                cont++
            }
            break
        }
    }
    *prom = totalCal/cont

    ImprimeMaps()
    return nil
}

func (a *SERVER) PromedioGeneral(empty string, prom *float64) error {
    var cont float64
    var contAlu float64
    var totalAluCal float64
    var totalProms float64

    for _, element := range alumnos {
        for _, cal := range element {
            totalAluCal += cal
            cont++
        }
        totalProms += (totalAluCal/cont)
        totalAluCal = 0
        cont = 0
        contAlu++
    }
    *prom = totalProms/contAlu

    ImprimeMaps()
    return nil
}

func (a *SERVER) PromedioMateria(materia string, prom *float64) error {
    var cont float64
    var totalCal float64

    for key, element := range materias {
        if key == materia {
            for _, cal := range element {
                totalCal += cal
                cont++
            }
            break
        }
    }
    *prom = totalCal/cont

    ImprimeMaps()
    return nil
}

func servidor() {
    rpc.Register(new(SERVER))
    ln, err := net.Listen("tcp", ":4040")
    if err != nil {
        fmt.Println(err)
    }
    for {
        c, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go rpc.ServeConn(c)
    }
}

func main() {
    go servidor()

    var input string
    fmt.Scanln(&input)
}
