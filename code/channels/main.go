package main

/*
Status checker
Are up & responding to HTTP trafic?
make an http get request to every website in the list.
*/
import (
	"fmt"
	"net/http"
	"time"
	//"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// (C)
	// create a channel the only way to communicate beetween go routines.
	// A channel is typed. This one has string type => Messages in chanel must be string only
	// make() create a value of the inside type
	// Note: c means channel
	c := make(chan string)

	//_ instead of index (because it doesn't matter to us here)
	for _, link := range links {

		// (A)
		//checkLink(link)

		// (B)&(C)
		// go inicial le dice que ejecute esa función en una "brand new child" "go routine" (<=> thread)
		// If we expect that the go routine that is launched by this line of code to be able to use this channel, we have to pass (as parameter) the channel into that function
		go checkLink(link, c)
	}

	// (B)
	// Waits (blocks) until receiving msg from channel (<=> receiving message from chanel is a blocking operation!!. ¡¡¡this is the main concept!!!)
	fmt.Println(<-c)
	fmt.Println(<-c)

	// (C)
	// Receive exactly the number of messages equal to the number urls in list
	// for  {
	// 	fmt.Println(<-c)
	// }

	// (D)
	// Whenever we receive from an url, we send it a new request to attempt receive again from the same url. We do it both if it has been fetched successfully or an error happened.
	// for  { // Infinite loop!
	// go checkLink(<-c, c)
	// }

	// (E)
	//Syntax equivalent to D but clearer for others.
	// Note l means link (<=> url)
	// for l := range c {
	// go checkLink(l, c)
	// }

	// (F)
	// vscode Warns 'loop variable l captured by a Function Literal'. Trying to share variables between different go routines MUST BE AVOIDED. Doing that is bad becase the go code inside function literal executes on another go routine than main routine where l takes value. What happens is that inside function literal always takes the main value of l of the first time that the function literal was invoked. Solution is pass l as parameter
	// In between every successful fetch, we put a very small pause between each checkLink call so that these different websites don't think that we are just trying to flood them with requests
	// for l := range c {
		// We don't call checkLink(link, c) becasuse doing that would pause main go routine
		// Instead we call as new routine a 'function literal' (same as C++ lambda function or javascript anonymous function) that is an unnamed function that we use to wrap some little chunk of code so we can execute it at some point in the future
		// go func() {
			// Inside the Literal function, we can pause our code 5 seconds
			// time.Sleep(5 * time.Second)
			// and then call the checkLink() function
			// checkLink(l, c)
			// The () after the close curly brace of the anonymous function means "call it now"
		// }()
	// }

	// (G)
	// In between every successful fetch, we put a very small pause between each checkLink call so that these different websites don't think that we are just trying to flood them with requests
	for l := range c {
		// We don't call checkLink(link, c) becasuse doing that would pause main go routine
		// Instead we call as new routine a 'function literal' (same as C++ lambda function or javascript anonymous function) that is an unnamed function that we use to wrap some little chunk of code so we can execute it at some point in the future
		go func(link string) {
			// Inside the Literal function, we can pause our code 5 seconds
			time.Sleep(5 * time.Second)
			// and then call the checkLink() function
			checkLink(link, c)
			// The () after the close curly brace of the anonymous function means "call it now"
		}(l)
	}
}

// Take a url & make an HTTP request to it & decide whether or not it is responding to traffic
// (A)
//func checkLink(link string) {
// (B)&(C)
func checkLink(link string, c chan string) {
	// We don't care the Response returned value.
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")

		// (B)&(C)
		// Send (the value of) link to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")

	// (B)&(C)
	c <- link
}

/*
Output:

(A)
http://google.com is up!
http://facebook.com is up!
http://stackoverflow.com is up!
http://golang.org is up!
http://amazon.com is up!

(B) only receives one link, the fastest responsing me.
http://stackoverflow.com is up!
http://stackoverflow.com


*/


/*
Which of the following best describes what a go routine is?
A separate line of code execution that can be used to handle blocking code

What's the purpose of a channel?
For communication between go routines



Take a look at the following program.  Are there any issues with it?
    package main
     
    import (
     "fmt"
    )
     
    func main() {
     greeting := "Hi There!"
     
     go (func() {
         fmt.Println(greeting) 
     })()
    }
Both are true: 
-The greeting variable is referenced from directly in the go routine, which might lead to issues if we eventually start to change the value of greeting
- and the program will likely exit before the fmt.Println function has an opportunity to actually print anything out to the terminal; this might not be the intent of the program




Is there any issue with the following code?
    package main
     
    func main() {
     c := make(chan string)
     c <- []byte("Hi there!")
    }
The channel is expecting values of type string, but we are passing in a value of type byte slice, which is not technically a string



Is there any issue with the following code?
    package main
     
    func main() {
         c := make(chan string)
         c <- "Hi there!"
    }
The syntax of this program is OK, but the program will never exit because it will wait for something to receive the value we're passing into the channel.





	
	


Ignoring whether or not the program will exit correctly, are the following two code snippets equivalent?

Snippet #1
    package main
     
    import "fmt"
     
    func main() {
     c := make(chan string)
     for i := 0; i < 4; i++ {
         go printString("Hello there!", c)
     }
     
     for s := range c {
         fmt.Println(s)
     }
    }
     
    func printString(s string, c chan string) {
     fmt.Println(s)
     c <- "Done printing." 
    }

Snippet #2
    package main
     
    import "fmt"
     
    func main() {
     c := make(chan string)
     
     for i := 0; i < 4; i++ {
         go printString("Hello there!", c)
     }
     
     for {
         fmt.Println(<- c)
     }
    }
     
    func printString(s string, c chan string) {
     fmt.Println(s)
     c <- "Done printing." 
    }
They are the same








*/
