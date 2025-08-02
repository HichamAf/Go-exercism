package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averageTimePerLayer int) int{
    count :=0
    for i:=0; i<len(layers); i++ {
        count += 1
    }
    if averageTimePerLayer > 0 {
        return count * averageTimePerLayer
    } else {
        return count * 2
    }
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for i:=0; i<len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 1
        } else if layers[i] == "sauce" {
            sauce += 1
        }
    }
    return noodles * 50, sauce * 0.2
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string)  {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, quantity int) []float64 {
    slice2 := []float64{}
    for i:= 0; i<len(slice); i++ {
        slice2 = append(slice2, (slice[i] * (float64(quantity)/2)))
    }
    return slice2
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
