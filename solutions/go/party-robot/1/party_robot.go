package partyrobot

import "strconv"

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

func convert(table int) string {
    var table_str string = strconv.Itoa(table)
    for len(table_str) < 3 {
        table_str = "0" + table_str
    }
    return table_str
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return "Welcome to my party, " + name + "!\nYou have been assigned to table " + convert(table) + ". Your table is " + direction + ", exactly " + strconv.FormatFloat(distance, 'f', 1, 64) + " meters from here.\nYou will be sitting next to " + neighbor + "."
}
