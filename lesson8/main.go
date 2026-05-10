package main

func main() {

	storage := NewStorage()

	u1 := storage.CreateUser("Ali", 18)
	u2 := storage.CreateUser("Bilol", 20)

	storage.GetUser(u1.ID)

	storage.UpdateUser(u2.ID, "Bilol V2", 25)

	storage.DeleteUser(u1.ID)

	storage.GetUser(u1.ID)
}
