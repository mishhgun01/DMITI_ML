package main

import (
	"DMITI_ML/pkg"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		log.Print(err.Error())
	}
	if file == nil {
		file, err = os.Create("output.txt")
		defer file.Close()
		if err != nil {
			log.Fatal("unresolved error:", err)
		}
	}
	rand.Seed(time.Now().UnixNano())
	var A, B pkg.Player
	r := 0.5
	for i := 0; i < 100; i++ {
		A.RandomizePicker(r)
		B.RandomizePicker(r)
		pkg.Winner(&A, &B)
	}
	file.WriteString("2nd experiment:\n")
	var sA = ""
	for _, v := range A.TableScore {
		sA += strconv.Itoa(v)
	}
	file.WriteString("A: " + sA + "\n")
	var sB = ""
	for _, v := range B.TableScore {
		sB += strconv.Itoa(v)
	}
	teoretically := 2*((1-r)*r) - 3*(r*(1-r)) - 1*((1-r)*r) + 2*(r*(1-r))
	result := float64(A.Score) / 100
	quad := math.Sqrt(teoretically*teoretically + result*result)
	dispersion := math.Pow(teoretically-result, 2) / 2
	file.WriteString("B: " + sB + "\n")
	file.WriteString("1st Won/Lost: " + strconv.Itoa(A.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(A.Score)/100) + "\n")
	file.WriteString("2nd Won/Lost: " + strconv.Itoa(B.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(B.Score)/100) + "\n")
	file.WriteString("Theoretical result for A: " + fmt.Sprintf("%.5f", teoretically) + "\n")
	file.WriteString("Standart deviation: " + fmt.Sprintf("%f", quad) + "\n")
	file.WriteString("Variance: " + fmt.Sprintf("%f", dispersion) + "\n")
	file.WriteString("3rd experiment:\n")
	rB := 0.25
	A, B = pkg.Player{}, pkg.Player{}
	for i := 0; i < 100; i++ {
		A.RandomizePicker(r)
		B.RandomizePicker(rB)
		pkg.Winner(&A, &B)
	}

	sA = ""
	for _, v := range A.TableScore {
		sA += strconv.Itoa(v)
	}
	file.WriteString("A: " + sA + "\n")
	sB = ""
	for _, v := range B.TableScore {
		sB += strconv.Itoa(v)
	}
	teoretically = 2*((1-r)*rB) - 3*(r*(1-rB)) - 1*((1-r)*rB) + 2*(r*(1-rB))
	result = float64(A.Score) / 100
	quad = math.Sqrt(teoretically*teoretically + result*result)
	dispersion = math.Pow(teoretically-result, 2) / 2
	file.WriteString("B: " + sB + "\n")
	file.WriteString("1st Won/Lost: " + strconv.Itoa(A.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(A.Score)/100) + "\n")
	file.WriteString("2nd Won/Lost: " + strconv.Itoa(B.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(B.Score)/100) + "\n")
	file.WriteString("Theoretical result for A: " + fmt.Sprintf("%.5f", teoretically) + "\n")
	file.WriteString("Standart deviation: " + fmt.Sprintf("%f", quad) + "\n")
	file.WriteString("Variance: " + fmt.Sprintf("%f", dispersion) + "\n")
	file.WriteString("4.1 - reinforcement learning:\n")
	rB = 0.5
	A, B = pkg.Player{}, pkg.Player{}
	ballsA := []int{10, 10}
	for i := 0; i < 100; i++ {
		r = float64(ballsA[0]) / float64(ballsA[0]+ballsA[1])
		A.RandomizePicker(r)
		B.RandomizePicker(rB)
		flag, score := pkg.Winner(&A, &B)
		if flag {
			ballsA[score[1]] += score[0]
		}
	}
	sA = ""
	for _, v := range A.TableScore {
		sA += strconv.Itoa(v)
	}
	file.WriteString("A: " + sA + "\n")
	sB = ""
	for _, v := range B.TableScore {
		sB += strconv.Itoa(v)
	}
	teoretically = 2*((1-r)*rB) - 3*(r*(1-rB)) - 1*((1-r)*rB) + 2*(r*(1-rB))
	result = float64(A.Score) / 100
	quad = math.Sqrt(teoretically*teoretically + result*result)
	dispersion = math.Pow(teoretically-result, 2) / 2
	file.WriteString("B: " + sB + "\n")
	file.WriteString("1st Won/Lost: " + strconv.Itoa(A.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(A.Score)/100) + "\n")
	file.WriteString("2nd Won/Lost: " + strconv.Itoa(B.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(B.Score)/100) + "\n")
	file.WriteString("Theoretical result for A: " + fmt.Sprintf("%.5f", teoretically) + "\n")
	file.WriteString("Standart deviation: " + fmt.Sprintf("%f", quad) + "\n")
	file.WriteString("Variance: " + fmt.Sprintf("%f", dispersion) + "\n")
	file.WriteString("4.2 - learning with punishment:\n")
	rB = 0.5
	A, B = pkg.Player{}, pkg.Player{}
	ballsA = []int{10, 10}
	for i := 0; i < 100; i++ {
		r = float64(ballsA[0]) / float64(ballsA[0]+ballsA[1])
		A.RandomizePicker(r)
		B.RandomizePicker(rB)
		flag, score := pkg.Winner(&A, &B)
		if !flag {
			ballsA[score[1]] += score[0]
		}
	}
	sA = ""
	for _, v := range A.TableScore {
		sA += strconv.Itoa(v)
	}
	file.WriteString("A: " + sA + "\n")
	sB = ""
	for _, v := range B.TableScore {
		sB += strconv.Itoa(v)
	}
	teoretically = 2*((1-r)*rB) - 3*(r*(1-rB)) - 1*((1-r)*rB) + 2*(r*(1-rB))
	result = float64(A.Score) / 100
	quad = math.Sqrt(teoretically*teoretically + result*result)
	dispersion = math.Pow(teoretically-result, 2) / 2
	file.WriteString("B: " + sB + "\n")
	file.WriteString("1st Won/Lost: " + strconv.Itoa(A.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(A.Score)/100) + "\n")
	file.WriteString("2nd Won/Lost: " + strconv.Itoa(B.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(B.Score)/100) + "\n")
	file.WriteString("Theoretical result for A: " + fmt.Sprintf("%.5f", teoretically) + "\n")
	file.WriteString("Standart deviation: " + fmt.Sprintf("%f", quad) + "\n")
	file.WriteString("Variance: " + fmt.Sprintf("%f", dispersion) + "\n")
	file.WriteString("4.3 - learning both players with reinforcement:\n")

	A, B = pkg.Player{}, pkg.Player{}
	ballsA = []int{10, 10}
	ballsB := []int{10, 10}
	for i := 0; i < 100; i++ {
		rB = float64(ballsB[0]) / float64(ballsB[0]+ballsB[1])
		r = float64(ballsA[0]) / float64(ballsA[0]+ballsA[1])
		A.RandomizePicker(r)
		B.RandomizePicker(rB)
		flag, score := pkg.Winner(&A, &B)
		if flag {
			ballsA[score[1]] += score[0]
		} else {
			ballsB[score[1]] -= score[0]
		}
	}
	sA = ""
	for _, v := range A.TableScore {
		sA += strconv.Itoa(v)
	}
	file.WriteString("A: " + sA + "\n")
	sB = ""
	for _, v := range B.TableScore {
		sB += strconv.Itoa(v)
	}
	teoretically = 2*((1-r)*rB) - 3*(r*(1-rB)) - 1*((1-r)*rB) + 2*(r*(1-rB))
	result = float64(A.Score) / 100
	quad = math.Sqrt(teoretically*teoretically + result*result)
	dispersion = math.Pow(teoretically-result, 2) / 2
	file.WriteString("B: " + sB + "\n")
	file.WriteString("1st Won/Lost: " + strconv.Itoa(A.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(A.Score)/100) + "\n")
	file.WriteString("2nd Won/Lost: " + strconv.Itoa(B.Score) + "\n" + "Middle value: " + fmt.Sprintf("%.2f", float64(B.Score)/100) + "\n")
	file.WriteString("Theoretical result for A: " + fmt.Sprintf("%.5f", teoretically) + "\n")
	file.WriteString("Standart deviation: " + fmt.Sprintf("%f", quad) + "\n")
	file.WriteString("Variance: " + fmt.Sprintf("%f", dispersion) + "\n")
}
