package main

import (
    "encoding/json"   // Used for encoding and decoding JSON data.
    "fmt"             // Used for formatted input/output (printing messages, scanning user input).
    "io/ioutil"       // Used to read the response body from the API.
    "net/http"        // Used to make HTTP requests.
    "os"             // Used for interacting with the filesystem (saving jokes in a file).
)

// Joke struct represents the structure of a joke returned by the API.
type Joke struct {
    Setup       string `json:"setup"`     // "setup" field of the joke (first part)
    Punchline   string `json:"punchline"` // "punchline" field of the joke (second part)
}

// fetchJoke fetches a random joke from an external API.
func fetchJoke() (*Joke, error) {
    url := "https://official-joke-api.appspot.com/random_joke" // API endpoint for a random joke.

    // Send a GET request to the API.
    resp, err := http.Get(url)
    if err != nil {
        return nil, err // Return an error if the request fails.
    }
    defer resp.Body.Close() // Ensure the response body is closed after function execution.

    // Read the response body.
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err // Return an error if reading the response body fails.
    }

    // Create a Joke variable to store the parsed JSON data.
    var joke Joke

    // Unmarshal (decode) the JSON response into the joke struct.
    if err := json.Unmarshal(body, &joke); err != nil {
        return nil, err // Return an error if JSON decoding fails.
    }

    return &joke, nil // Return the joke and no error.
}

// saveJokeToFile saves a joke to a text file.
func saveJokeToFile(joke *Joke) error {
    // Open (or create) a file named "jokes.txt" in append mode.
    file, err := os.OpenFile("jokes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err // Return an error if the file can't be opened.
    }
    defer file.Close() // Ensure the file is closed after writing.

    // Format the joke as a string: "Setup - Punchline"
    jokeText := fmt.Sprintf("%s - %s\n", joke.Setup, joke.Punchline)

    // Write the joke to the file.
    _, err = file.WriteString(jokeText)
    return err // Return any error encountered while writing.
}

// main function controls the program's flow.
func main() {
    // Fetch a joke from the API.
    joke, err := fetchJoke()
    if err != nil {
        fmt.Println("Error fetching the joke:", err)
        return // Exit if fetching the joke fails.
    }

    // Display the joke to the user.
    fmt.Println("😂 Joke of the day:")
    fmt.Println(joke.Setup)   // Print the setup (first part of the joke).
    fmt.Println(joke.Punchline) // Print the punchline (second part of the joke).

    // Ask the user if they want to save the joke.
    fmt.Print("Do you want to save this joke? (yes/no): ")
    var response string
    fmt.Scanln(&response) // Read user input.

    // If the user chooses to save the joke, write it to the file.
    if response == "yes" {
        err := saveJokeToFile(joke)
        if err != nil {
            fmt.Println("Error saving the joke:", err)
        } else {
            fmt.Println("✅ Joke saved in jokes.txt!") // Confirm successful save.
        }
    }
}
