package pkg

func CalculateVAT(subTotal,vatRate float64)(float64,float64){
	vatAmount := subTotal * vatRate/100
	total:=vatAmount+subTotal
	return vatAmount,total
}