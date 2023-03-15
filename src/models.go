package main


type Role string
type Interest string


type User struct {
    Forename  string     `mapstructure:"forename"`
    Surname   string     `mapstructure:"surname"`
    Role      Role       `mapstructure:"role"`
    EntryYear int        `mapstructure:"entryYear"`
    Interests []Interest `mapstructure:"interests"`
	ID        string     `mapstructure:"id"`
}

