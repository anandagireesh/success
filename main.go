package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func passion() string {

	message := ""

	fmt.Println("Do you have passion with your Goal (Y/N):")
	passionInput := scanner()

	if "Y" == strings.ToUpper(passionInput) || "YES" == strings.ToUpper(passionInput) {

		plan := pLan()
		message = plan

	} else if "N" == strings.ToUpper(passionInput) || "NO" == strings.ToUpper(passionInput) {

		message = "Set a goal that you have passioned"

	} else {
		fmt.Println("Wrong input")

	}

	return message

}

func pLan() string {
	message := ""
	fmt.Println("Do you have any plan to achive your Goal (Y/N):")
	plan := scanner()
	if "Y" == strings.ToUpper(plan) || "YES" == strings.ToUpper(plan) {

		hardWork := hardWork()
		message = hardWork

	} else if "N" == strings.ToUpper(plan) || "NO" == strings.ToUpper(plan) {

		message = "Set a goal in life. work for that goal to achive"

	} else {
		fmt.Println("Wrong input")

	}
	return message
}

func hardWork() string {

	message := ""
	fmt.Println("are you working very hard to achive your goal (Y/N):")
	hardwork := scanner()
	if "Y" == strings.ToUpper(hardwork) || "YES" == strings.ToUpper(hardwork) {

		message = "Hope for the best \n"
		fmt.Println("Did your plan Fails (Y/N):")
		planFails := scanner()

		if "Y" == strings.ToUpper(planFails) || "YES" == strings.ToUpper(planFails) {

			message = message + "Do Not loose hope \n Change your plan \n Work for that plan to achieve goal \n"

		} else if "N" == strings.ToUpper(planFails) || "NO" == strings.ToUpper(planFails) {

			message = message + "You will get success one day \n"

		} else {
			fmt.Println("Wrong input")
		}

	} else if "N" == strings.ToUpper(hardwork) || "NO" == strings.ToUpper(hardwork) {

		message = "You have to work hard to achieve your goal"

	} else {
		fmt.Println("Wrong input")

	}
	return message
}

func noGoal() string {

	return "Set a goal in life. work for that goal to achive"

}

func main() {

	fmt.Println("type Y/N")
	fmt.Println("Do you have any goal (Y/N):")
	Goal := scanner()

	if "Y" == strings.ToUpper(Goal) || "YES" == strings.ToUpper(Goal) {
		result := passion()
		fmt.Println(result)

	} else if "N" == strings.ToUpper(Goal) || "NO" == strings.ToUpper(Goal) {
		result := noGoal()
		fmt.Println(result)

	} else {

		fmt.Println("Wrong input")
	}
	
	fmt.Println("successfully working")

}
