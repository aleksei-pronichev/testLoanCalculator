package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

const incorrectParameters = "Incorrect parameters"

func main() {
	var loanAmount, monthlyPayment, numberOfMonths, interest float64
	var interestType string
	initFlagArguments(&loanAmount, &monthlyPayment, &numberOfMonths, &interest, &interestType)

	if len(os.Args) < 5 {
		fmt.Println(incorrectParameters)
		return
	}
	if loanAmount < 0 || monthlyPayment < 0 || numberOfMonths < 0 || interest < 0 {
		fmt.Println(incorrectParameters)
		return
	}

	if interestType == "annuity" {
		annuityLoan(loanAmount, monthlyPayment, numberOfMonths, interest)
	} else if interestType == "diff" {
		diffLoan(loanAmount, monthlyPayment, numberOfMonths, interest)
	} else {
		fmt.Println(incorrectParameters)
	}
}

func annuityLoan(loanAmount float64, monthlyPayment float64, numberOfMonths float64, interest float64) {
	if interest == 0 {
		fmt.Println(incorrectParameters)
		return
	}

	if loanAmount == 0 {
		calculateLoan(monthlyPayment, numberOfMonths, interest)
	} else if monthlyPayment == 0 {
		calculateMonthlyPayment(loanAmount, numberOfMonths, interest)
	} else if numberOfMonths == 0 {
		calculateNumberOfMonth(loanAmount, monthlyPayment, interest)
	}
}

func diffLoan(amount float64, payment float64, months float64, interest float64) {
	if payment != 0 {
		fmt.Println(incorrectParameters)
		return
	}

	var totalPayment float64
	for i := 1; i <= int(months); i++ {
		monthlyPayment := amount/months + interest/1200*(amount-(amount*(float64(i)-1)/months))
		monthlyPayment = math.Ceil(monthlyPayment)
		totalPayment += monthlyPayment
		fmt.Printf("Month %d: payment is %.0f\n", i, monthlyPayment)
	}
	fmt.Printf("\nOverpayment = %.0f", totalPayment-amount)
}

func initFlagArguments(loanAmount, monthlyPayment, numberOfMonths, interest *float64, interestType *string) {
	flag.Float64Var(loanAmount, "principal", 0, "loan amount")
	flag.Float64Var(monthlyPayment, "payment", 0, "monthly payment")
	flag.Float64Var(numberOfMonths, "periods", 0, "number of months")
	flag.Float64Var(interest, "interest", 0, "interest rate")
	flag.StringVar(interestType, "type", "", "type of loan")
	flag.Parse()
}

func calculateLoan(payment float64, months float64, interest float64) {
	loan := payment * (math.Pow(1+interest/1200, months) - 1) / (interest / 1200 * math.Pow(1+interest/1200, months))
	loan = math.Round(loan)
	overPayment := payment*months - loan
	fmt.Printf("Your loan principal = %.0f!\nOverpayment = %.0f", loan, overPayment)

}

func calculateMonthlyPayment(amount float64, months float64, interest float64) {
	var monthlyPayment float64
	i := interest / (12 * 100)
	monthlyPayment = amount * (i * math.Pow(1+i, months)) / (math.Pow(1+i, months) - 1)
	monthlyPayment = math.Ceil(monthlyPayment)
	fmt.Printf("Your monthly payment = %.0f!", monthlyPayment)
}

func calculateNumberOfMonth(amount float64, payment float64, interest float64) {
	var months = calculateNumberOfPayments(payment, amount, interest)
	years := months / 12
	months = months % 12
	overPayment := payment*float64(months) + payment*float64(years)*12 - amount

	var monthString string = ""
	if months == 1 {
		monthString = "1 month"
	} else if months != 0 {
		monthString = fmt.Sprintf("%d months", months)
	}
	var yearString string = ""
	if years == 1 {
		yearString = "1 year"
	} else if years > 0 {
		yearString = fmt.Sprintf("%d years", years)
	}
	fmt.Printf("It will take %s %s to repay the loan", yearString, monthString)
	fmt.Printf("\nOverpayment = %.0f", overPayment)
}

func calculateNumberOfPayments(payment, principal, interestRate float64) int {
	i := interestRate / (12 * 100) // Convert annual interest rate to monthly and to a decimal

	n := math.Log(payment/(payment-i*principal)) / math.Log(1+i)

	return int(math.Ceil(n)) // Round up to the next whole number
}
