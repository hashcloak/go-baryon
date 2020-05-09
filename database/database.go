package database

type PIRDatabase interface {
	// TODO: Complete method declarations

	Init()

	Add()

	Remove()

	Retrieve()

	// OPTIONAL: Only for IT-PIR protocols that work in the expressive query setting
	ExpressiveRetrieve()
}
