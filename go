package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func getDifficulty() (low, high, maxAttempts int) {
    fmt.Println("\nChoose difficulty:")
    fmt.Println("1. Easy (1-50, 15 attempts)")
    fmt.Println("2. Medium (1-100, 10 attempts)")
    fmt.Println("3. Hard (1-200, 8 attempts)")
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            return 1, 50, 15
        case "2":
            return 1, 100, 10
        case "3":
            return 1, 200, 8
        default:
            fmt.Println("Invalid choice. Enter 1, 2, or 3.")
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    play()
}

func play() {
    low, high, maxAttempts := getDifficulty()
    secret := rand.Intn(high-low+1) + low
    attempts := 0
    best := 0
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf("\nI'm thinking of a number between %d and %d. You have %d attempts.\n", low, high, maxAttempts)
    for attempts < maxAttempts {
        fmt.Printf("Attempt %d/%d: ", attempts+1, maxAttempts)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        guess, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Please enter a valid number.")
            continue
        }
        if guess < low || guess > high {
            fmt.Printf("Number must be between %d and %d.\n", low, high)
            continue
        }
        attempts++
        if guess == secret {
            fmt.Printf("Congratulations! You guessed it in %d attempts.\n", attempts)
            if best == 0 || attempts < best {
                best = attempts
                fmt.Printf("New best score: %d attempts!\n", best)
            }
            break
        } else if guess < secret {
            fmt.Println("Too low!")
        } else {
            fmt.Println("Too high!")
        }
    }
    if attempts == maxAttempts {
        fmt.Printf("Sorry, you've used all attempts. The number was %d.\n", secret)
    }
    fmt.Printf("Best score this session: ")
    if best == 0 { fmt.Println("N/A") } else { fmt.Println(best) }
    fmt.Print("Play again? (y/n): ")
    again, _ := reader.ReadString('\n')
    again = strings.TrimSpace(strings.ToLower(again))
    if again == "y" {
        play()
    } else {
        fmt.Println("Thanks for playing!")
    }
}
