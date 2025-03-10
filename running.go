package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func runAPP(username string) {
	var userChoice string

	fmt.Printf("Hello %s, welcome to our Book application!\n", username)
	fmt.Println("Click `h` for help menu!")
	fmt.Println("--------------------------------------------------")

	for {
		fmt.Print("Enter your command: ")
		fmt.Scan(&userChoice)
		fmt.Println("--------------------------------------------------")

		switch userChoice {
		case "h":
			helpMenu()
		case "b":
			showAllBooks()
		case "a":
			addBook()
		case "d":
			deleteBook()
		case "q":
			quitApp()
			return
		case "e":
			editBook()
		case "rd":
			rankDescending()
		case "ra":
			rankAscending()
		default:
			fmt.Println("This command doesn't exist.")
		}
	}
}

func helpMenu() {
	fmt.Println("Welcome to the help menu.")
	fmt.Println("Command `h` : Shows the help menu.")
	fmt.Println("Command `b` : Shows all the books you have read so far.")
	fmt.Println("Command `a`: Add a new book you have read.")
	fmt.Println("Command `d`: Delete a book you have read.")
	fmt.Println("Command `e`: Edit an existing book you have read.")
	fmt.Println("Command `rd`: Rank books from highest to lowest rated.")
	fmt.Println("Command `ra`: Rank books from lowest to highest rated.")
	fmt.Println("Command `q`: Quit the application.")
	fmt.Println("--------------------------------------------------")
}

type Book struct {
	Title  string
	Author string
	Pages  int
	Rating int
}

var books = []Book{}

func showAllBooks() {
	if len(books) == 0 {
		fmt.Println("No books added yet.")
		fmt.Println("--------------------------------------------------")
		return
	}

	fmt.Println("All Books you have read so far:")
	for _, book := range books {
		fmt.Println("Title: ", book.Title)
		fmt.Println("Author: ", book.Author)
		fmt.Println("Pages: ", book.Pages)
		fmt.Println("Rating: ", book.Rating)
		fmt.Println("------------------------------")
	}
}

func addBook() {
	var title, author string
	var pages, rating int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Add a new book you have read.")

	fmt.Print("Enter book title: ")
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)
	if title == "" {
		fmt.Println("Title cannot be empty.")
		return
	}

	fmt.Print("Enter book author: ")
	author, _ = reader.ReadString('\n')
	author = strings.TrimSpace(author)
	if author == "" {
		fmt.Println("Author cannot be empty.")
		return
	}

	for {
		fmt.Print("Enter amount of book pages (1-2500): ")
		fmt.Scan(&pages)

		if pages >= 1 && pages <= 2500 {
			break
		} else {
			fmt.Println("Please enter a valid number of pages between 1 and 2500.")
		}
	}

	for {
		fmt.Print("Enter rating of book (1-10): ")
		fmt.Scan(&rating)

		if rating >= 1 && rating <= 10 {
			break
		} else {
			fmt.Println("Rating must be between 1 and 10. Please try again.")
		}
	}

	newBook := Book{
		Title:  title,
		Author: author,
		Pages:  pages,
		Rating: rating,
	}

	books = append(books, newBook)
	fmt.Println("New book added successfully!")
	fmt.Println("--------------------------------------------------")
}

func deleteBook() {
	if len(books) == 0 {
		fmt.Println("No books available to delete.")
		fmt.Println("--------------------------------------------------")
		return
	}

	var bookTitle string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the book title you want to delete: ")
	bookTitle, _ = reader.ReadString('\n')
	bookTitle = strings.TrimSpace(bookTitle)

	newBooks := []Book{}
	found := false

	for _, book := range books {
		if strings.EqualFold(book.Title, bookTitle) {
			found = true
			fmt.Println("Deleting the book...")
			continue
		}
		newBooks = append(newBooks, book)
	}

	if !found {
		fmt.Println("Book not found.")
	} else {
		fmt.Println("Book deleted successfully.")
		books = newBooks
	}
	fmt.Println("--------------------------------------------------")
}

func quitApp() {
	fmt.Println("Exit application...")
	os.Exit(0)
}

func editBook() {
	if len(books) == 0 {
		fmt.Println("No books available to edit.")
		fmt.Println("--------------------------------------------------")
		return
	}

	var bookName string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the book name you want to edit: ")
	bookName, _ = reader.ReadString('\n')
	bookName = strings.TrimSpace(bookName) // Trim spaces and newline

	found := false

	for i, book := range books {
		if strings.EqualFold(book.Title, bookName) {
			found = true

			// Ask for new details
			fmt.Println("Editing book:", book.Title)

			fmt.Print("Enter new title (or press Enter to keep it the same): ")
			newTitle, _ := reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)
			if newTitle != "" {
				books[i].Title = newTitle
			}

			fmt.Print("Enter new author (or press Enter to keep it the same): ")
			newAuthor, _ := reader.ReadString('\n')
			newAuthor = strings.TrimSpace(newAuthor)
			if newAuthor != "" {
				books[i].Author = newAuthor
			}

			fmt.Print("Enter new number of pages (or press Enter to keep it the same): ")
			var newPages string
			newPages, _ = reader.ReadString('\n')
			newPages = strings.TrimSpace(newPages)
			if newPages != "" {
				pages, err := strconv.Atoi(newPages)
				if err == nil {
					books[i].Pages = pages
				} else {
					fmt.Println("Invalid input for pages. Keeping the old value.")
				}
			}

			fmt.Print("Enter new rating (or press Enter to keep it the same): ")
			var newRating string
			newRating, _ = reader.ReadString('\n')
			newRating = strings.TrimSpace(newRating)
			if newRating != "" {
				rating, err := strconv.Atoi(newRating)
				if err == nil {
					books[i].Rating = rating
				} else {
					fmt.Println("Invalid input for rating. Keeping the old value.")
				}
			}

			fmt.Println("Book updated successfully!")
			break
		}
	}

	if !found {
		fmt.Println("Book not found.")
	}
	fmt.Println("--------------------------------------------------")
}

func rankDescending() {
	if len(books) == 0 {
		fmt.Println("No books added yet.")
		fmt.Println("-------------------------------------------------")
		return
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].Rating > books[j].Rating
	})

	fmt.Println("All Books you have read so far ranked from highest to lowest rating:")

	for i, book := range books {
		fmt.Printf("%d. Title: %s, Author: %s, Pages: %d, Rating: %d\n", i+1, book.Title, book.Author, book.Pages, book.Rating)
	}

}

func rankAscending() {
	if len(books) == 0 {
		fmt.Println("No books added yet.")
		fmt.Println("-------------------------------------------------")
		return
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].Rating < books[j].Rating
	})

	fmt.Println("All Books you have read so far ranked from lowest to highest rating:")

	for i, book := range books {
		fmt.Printf("%d. Title: %s, Author: %s, Pages: %d, Rating: %d\n", i+1, book.Title, book.Author, book.Pages, book.Rating)
	}

}
