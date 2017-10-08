package common

import(
	"time"
	"math/rand"
)

// add additive noise that is scaled by a given factor
func addAdditiveNoise(scalingFactor float32) float32{
	noiseSource := rand.NewSource(time.Now().UnixNano()) // random seed based on the current time
	return scalingFactor * (2 * rand.New(noiseSource).Float32() - 1) // generate a random number in [-1,1] and rescale it by the given scaling factor
}



