# 😂 Joke Generator

This is a simple **Joke Generator** written in Go. It fetches a random joke from the **Official Joke API** and displays it in the console. Additionally, users can save jokes to a file for later enjoyment.

## 🚀 Features
- Fetches a **random joke** from an external API.
- Displays the joke in the console.
- Allows the user to **save** the joke in a `jokes.txt` file.

## 🛠️ Installation & Usage

### Prerequisites
Make sure you have **Go** installed on your machine.

### Clone the Repository
```bash
git clone https://github.com/mpvz7/joke-generator.git
cd joke-generator
```

### Run the Program
```bash
go run main.go
```

### Example Output
```
😂 Joke of the day:
Why don't scientists trust atoms?
Because they make up everything!

Do you want to save this joke? (yes/no)
✅ Joke saved in jokes.txt!
```

## 📂 File Structure
```
/joke-generator
│── main.go   # Main Go program
│── jokes.txt # (Optional) Stored jokes
│── README.md # Project documentation
```

## 🔧 How It Works
1. Sends an HTTP request to the **Official Joke API**.
2. Parses the received JSON data.
3. Displays the joke (setup + punchline) in the terminal.
4. Asks the user if they want to save the joke.
5. If confirmed, appends the joke to `jokes.txt`.

## 🌟 Future Improvements
- Allow fetching multiple jokes at once.
- Add joke categories (e.g., programming, general, dad jokes).
- Create a simple web UI to display jokes.

---
💡 **Made with Go - Have fun coding!** 🚀

