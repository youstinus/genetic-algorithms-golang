package main

// MutationRate is the rate of mutation
/*var MutationRate = 0.005

// PopSize is the size of the population
var PopSize = 500

// Letters need to be chosen from
var Letters = []rune(" aąbcčdeęėfghiįyjklmnoprsštuųūvzž")

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	target := []rune("aš dar visiškai žalias ir nežinau ar man dar ilgesnį rašyt sakinį")
	population := createPopulation(target)

	found := false
	generation := 0
	for !found {
		generation++
		bestIndividual := getBest(population)
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestIndividual.DNA), bestIndividual.Fitness)

		if string(bestIndividual.DNA) == string(target) {
			found = true
		} else {
			maxFitness := bestIndividual.Fitness
			pool := createPool(population, target, maxFitness)
			population = naturalSelection(pool, population, target)
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime taken: %s\n", elapsed)
}

// create the breeding pool that creates the next generation
func createPool(population []Individual, target []rune, maxFitness float64) (pool []Individual) {
	pool = make([]Individual, 0)
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// perform natural selection to create the next generation
func naturalSelection(pool []Individual, population []Individual, target []rune) []Individual {
	next := make([]Individual, len(population))

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := crossover(a, b)
		child.mutate()
		child.calcFitness(target)

		next[i] = child
	}
	return next
}

// crosses over 2 Individuals
func crossover(d1 Individual, d2 Individual) Individual {
	child := Individual{
		DNA:     make([]rune, len(d1.DNA)),
		Fitness: 0,
	}
	mid := rand.Intn(len(d1.DNA))
	for i := 0; i < len(d1.DNA); i++ {
		if i > mid {
			child.DNA[i] = d1.DNA[i]
		} else {
			child.DNA[i] = d2.DNA[i]
		}

	}
	return child
}

// mutate the Individual
func (d *Individual) mutate() {
	for i := 0; i < len(d.DNA); i++ {
		if rand.Float64() < MutationRate {
			d.DNA[i] = rune(Letters[rand.Intn(33)])
		}
	}
}

// Get the best individual
func getBest(population []Individual) Individual {
	best := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > best {
			index = i
			best = population[i].Fitness
		}
	}
	return population[index]
}*/
