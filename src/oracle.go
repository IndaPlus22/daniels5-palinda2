// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")
	fmt.Print(prompt)

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard you \n", star)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	// TODO: Answer questions.
	go answer(questions, answers)
	// TODO: Make prophecies.
	go nonsenseFunc(answers)
	// TODO: Print answers.
	go print(answers)
	return questions
}

func nonsenseFunc(answer chan<- string) {
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"I see that you are using a computer",
		"Wow!!! nvm i thought you said something",
		"It is a me mario, uuh I mean Pythia",
		"Nice weather today",
	}

	for {
		time.Sleep(time.Duration(30+rand.Intn(30)) * time.Second)
		answer <- nonsense[rand.Intn(len(nonsense))]
	}

}

func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	//time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	fmt.Print(prompt)
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	standardAsnwers := map[string]string{
		"where":  "You are here",
		"life":   "How should I know?",
		"wisdom": "Only one thing is ever guaranteed, that is that you will definitely not achieve the goal if you don't take the shot.",
		"name":   "Pythia",
		"21":     "Is the sum of 9 + 10",
	}
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, word := range words {
		if standardAsnwers[strings.ToLower(word)] != "" {
			answer <- standardAsnwers[strings.ToLower(word)]
			return
		}
	}
	// Find the longest word.
	longestWord := ""
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"I have shoes",
		"I do not like java",
		"RUST FOR THE WIN",
		"Milk",
	}
	answer <- "In my world " + longestWord + " means " + nonsense[rand.Intn(len(nonsense))]

}

func answer(question <-chan string, answer chan<- string) {

	for quse := range question {
		go prophecy(quse, answer)
	}

}

func print(answer <-chan string) {
	for question := range answer {
		//fmt.Println(star + ": " + question)
		strArr := []rune(question)
		fmt.Print("ANSWER: ")
		for index := range question {
			time.Sleep(time.Duration(10+rand.Intn(100)) * time.Millisecond)
			fmt.Print(string(strArr[index]))
		}

		fmt.Print("\n" + prompt)
	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
