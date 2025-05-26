package lasagna

func PreparationTime( Layers []string, Time int ) int {
	if Time == 0 {
		Time = 2
	}
	return len(Layers) * Time

}

func Quantities( Ingredients []string ) (int, float64) {
	noodles := 0
	sauce := 0
	for _, i := range Ingredients {
		if i == "noodles" {
			noodles += 1
		}
		if i == "sauce" {
			sauce += 1
		} 
	}
	return noodles * 50, float64(sauce) * 0.2
}

func AddSecretIngredient( friendsList []string, myList []string ) {
	getsecret := len(friendsList) - 1
	replace := len(myList) - 1 

	myList[replace] = friendsList[getsecret]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {

	var newQuantities []float64
	
	for _, quantity := range quantities {
		newQuantities = append(newQuantities, quantity * float64(portions)/2)
	}
		
	return newQuantities
}
