package services

import (
	"errors"
	"library/models"
)

type LibraryManager interface{
	Register(name models.Member)
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memeberID int)error
	ReturnBook(bookID int, memeberID int)error
	ListAvailableBooks()[]models.Book
	ListBorrowedBooks(memberId int)[]models.Book
}

type Library struct {
	books map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) Register(member models.Member){
	l.members[member.ID] = member
}

func (l *Library) AddBook(book models.Book){
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int){
	delete(l.books,bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int)error{
	book,exists := l.books[bookID]
	if !exists{
		return errors.New("book not found")
	}
	if book.Status == "Borrowed"{
		return errors.New("book already borrowed")
	}

	member,exists := l.members[memberID]
	if !exists{
		return errors.New("user not found")
	}
	member.BorrowedBooks = append(member.BorrowedBooks, book)

	p := &book
	p.Status = "Borrowed"

	return nil
}

func (l *Library) ReturnBook(bookID,memberID int)error{
	book,exists := l.books[bookID]
	if !exists{
		return errors.New("book not found")
	}

	member,exists := l.members[memberID]
	if !exists{
		return errors.New("user not found")
	}

	p := &book
	p.Status = "Avaliable"

	for i,b := range l.books{
		if b.ID == bookID{
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}

	return nil
}

func (l *Library) ListAvailableBooks()[]models.Book{
	var list []models.Book
	for _,book := range l.books{
		if book.Status == "Avaliable"{
			list = append(list,book)
		}
	}

	return list
}

func (l *Library) ListBorrowedBooks(memberID int)[]models.Book{
	member, exists := l.members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}