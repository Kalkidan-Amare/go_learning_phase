package controllers

import (
	"fmt"
	"library/models"
	"library/services"
)

func RunLibrary() {
	library := services.NewLibrary()

	for {
		fmt.Println("--------------------------------")
		fmt.Println("Library Management System")
		fmt.Println("0. Register")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Println("--------------------------------")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			var id int
			var name string
			fmt.Print("Enter your ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter your name: ")
			fmt.Scan(&name)
			library.Register(models.Member{ID: id, Name: name, BorrowedBooks: []models.Book{}})
		case 1:
			var id int
			var title, author string
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter book author: ")
			fmt.Scan(&author)
			library.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
		case 2:
			var id int
			fmt.Print("Enter book ID to remove: ")
			fmt.Scan(&id)
			library.RemoveBook(id)
		case 3:
			var bookID, memberID int
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			var bookID, memberID int
			fmt.Print("Enter book ID to return: ")
			fmt.Scan(&bookID)
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 6:
			var memberID int
			fmt.Print("Enter member ID: ")
			fmt.Scan(&memberID)
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case 7:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
