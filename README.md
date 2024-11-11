## Assignment Description

### Management Problem

Managers of a technology startup are concerned about their "monthly burn" (how fast they are spending investment capital). Some are intent on cutting software development costs.

A recent MIT lecture suggested that computers (in particular, large language models) could replace programmers in the not-to-distant future: 

Large Language Models and the End of Programming by Matt WelshLinks to an external site.

One manager is making a case for using AI-assisted programming today, even before the firm introduces its first product:

Why are we paying junior Go programmers a base salary of $120 thousand a year?  Why are we paying senior software engineers more than $175 thousand? From what I read about Go code generators and GitHub Copilot, we should be able to reduce development costs substantially, maybe by as much as 50 percent. We can use Go to generate code. We can pair every junior programmer with a robot coder. We can insist that every programmer use GitHub Copilot. Let's implement automation in the software development process whenever and wherever possible. We either cut personnel costs or go for another round of venture capital funding. I say we cut costs.

Concerned that the programmers and software engineers within the firm cannot be trusted to provide an objective evaluation of the possibilities of automated programming, AI-assisted programming, or AI-generated code, the firm's management hires an outside consultant to evaluate the possibilities.

### Assignment Requirements 

In working with code generation or AI-assistant tools for this assignment, ensure that you use free versions. You are not being asked to subscribe to any services that have a fee.

Suppose you are the outside consultant hired by the startup. Here are the suggested steps in carrying out this assignment:

* Return to an earlier assignment in the course: ***Go for Statistics***. Recall that this assignment drew on data for the Anscombe Quartet.
* **AI-assisted programming**. Looking over your repository for Go for Statistics, consider the possibilities of revising your original code with the assistance of a product like GitHub Copilot (perhaps by using the GitHub Copilot extension to VS Code). Do you see ways that Copilot can help with the programming assignment? GitHub Copilot provides suggestions for numerous languages and a wide variety of frameworks, and it works especially well for Go, Python, JavaScript, TypeScript, Ruby, C# and C++, as well as SQL. See documentation for GitHub Copilot in VS CodeLinks to an external site..
* **AI-generated code**. Select an LLM-based service or agent such as ChatGPT. Pick a free plan. (Do not pay for services in completing this assignment.) Indicate which service you selected, providing its web address or uniform resource locator (URL) See if you can get your selected LLM agent to generate Go code for Go for Statistics. Provide a step-by-step review of your work. Try various prompts. Build a conversation with the LLM agent. Tailor the conversation to the specific needs of the programming assignment. In your repository for this assignment, include a plain text file of your complete conversation with the LLM agent.
* Drawing on the methods above, share as much actual Go code as possible, recognizing that the code may be incomplete or invalid. In the README.md file, refer to the repository file(s) for the code and describe how it was obtained.
* As part of the README.md file, summarize your experiences with each of the three methods: automated code generation, AI-assisted programming, and AI-generated code. What are the possibilities for reducing the programming workload in the startup? What is your recommendation to the firm regarding its staffing of programmers and software engineers? 
* As part of the README.md file, also list the training materials that you used, providing the URLs and noting their usefulness in completing the assignment. Your recommendations will be noted and could influence future versions of this assignment.

## AI-Assisted Programming

## AI-Generated Code

### Chosen LLM: Chat-GPT 4
During this assignment, it was decided to use Chat-GPT for code generation and testing since it is generally regarded as the [most popular LLM at the time of creating this](https://www.techradar.com/computing/artificial-intelligence/best-llms). Especially with the release of GPT 4 and 4-o, the AI is able to not only give more consise responses, but also handle mathematical logic significantly more gracefully.

Prompts and code were fed to the LLM directly from the [official website](https://chatgpt.com/). 

### AI-Generated Code
The following prompt was used to have Chat GPT generate code to output the results of the anscombe quartet.

 *Are you familiar with the Stats Package for Go by montanaflynn? Utilizing the package, can you write program that outputs the results of linear regression from the Anscombe quartet.*

````go
package main

import (
	"fmt"
	"log"

	"github.com/montanaflynn/stats"
)

// Define the Anscombe quartet datasets
var anscombeQuartet = []struct {
	x []float64
	y []float64
}{
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	},
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
	},
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	},
	{
		x: []float64{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 19},
		y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 5.56, 7.91, 6.89, 12.50},
	},
}

func main() {
	// Loop over each dataset in Anscombe's quartet
	for i, dataset := range anscombeQuartet {
		// Calculate the linear regression
		slope, intercept, err := stats.LinearRegression(dataset.x, dataset.y)
		if err != nil {
			log.Fatalf("Error calculating linear regression for dataset %d: %v", i+1, err)
		}

		// Output the results
		fmt.Printf("Dataset %d:\n", i+1)
		fmt.Printf("Slope: %.4f\n", slope)
		fmt.Printf("Intercept: %.4f\n\n", intercept)
	}
}

````

This code did not compile and the errors printed were:

````
.\main.go:37:28: assignment mismatch: 3 variables but stats.LinearRegression returns 2 values
.\main.go:37:62: too many arguments in call to stats.LinearRegression
        have ([]float64, []float64)
        want (stats.Series)
````
This error message was fed back to the LLM and the revised code below was generated:

````go
package main

import (
	"fmt"
	"log"

	"github.com/montanaflynn/stats"
)

// Define the Anscombe quartet datasets
var anscombeQuartet = []struct {
	x []float64
	y []float64
}{
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	},
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.10, 6.13, 3.10, 9.13, 7.26, 4.74},
	},
	{
		x: []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		y: []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	},
	{
		x: []float64{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 19},
		y: []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 5.56, 7.91, 6.89, 12.50},
	},
}

func main() {
	// Loop over each dataset in Anscombe's quartet
	for i, dataset := range anscombeQuartet {
		// Create a stats.Series from the x and y values
		var series stats.Series
		for j := 0; j < len(dataset.x); j++ {
			series = append(series, stats.Coordinate{X: dataset.x[j], Y: dataset.y[j]})
		}

		// Calculate the linear regression
		line, err := stats.LinearRegression(series)
		if err != nil {
			log.Fatalf("Error calculating linear regression for dataset %d: %v", i+1, err)
		}

		// Output the results
		fmt.Printf("Dataset %d:\n", i+1)
		fmt.Printf("Slope: %.4f\n", line.Slope)
		fmt.Printf("Intercept: %.4f\n\n", line.Intercept)
	}
}
````

This, however, did not compile either. GPT made the incorrect assumption that the returned value of the *LinearRegression* function contained was a type struct that contained an attribute of Slope and intercept. However, this is not how the function returns values; instead it converts the provided x,y coordinates into the coordinates of the regression line. At this point, providing GPT with the error messages was not producing any working code. This information was passed to the LLM through the following prompt:

*The LinearRegression function returns values as a stats.Series struct, not a line struct*

While the LLM model was able to correctly determine that they it would have to manually calculate the slope and intercept, it still attemtped to provide *stats.LinearRegression* with too many return arguments. At this point, it does not seem that although the LLM does have training data on the general overview of how to use the stats package, it does not have enough to correctly understand how the function returns values. 

The most likely cause of this interesting behavior is that the model may have been built on older versions of the stats package that returned just the slope and intercept, and not a corrdinate series. Ultimately, I was not able to get the LLM to correct it's mistakes or correct the code through prompt engineering.

The entire conversation can be found on this [link](https://chatgpt.com/share/672e52ae-d3e0-8008-8dea-bf95c3c61c7a) or in [GPT Conversation.txt](https://github.com/miguswong/WongMSDS437Assignment7/blob/main/GPT%20Conversation.txt)