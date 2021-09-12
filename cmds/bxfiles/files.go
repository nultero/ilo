package bxfiles

var files = []string{
	"tailbox_cache.txt",
	"config.txt",
	"events.txt",
	"ideas.txt",
	"recurrent_reminders.txt",
	"reminders.txt",
	"todos.txt",
	"wishlist.txt",
}

var DefaultConfig = []byte(
	"todoSymbol = ○" +
		"\n" + "icon = ❯➤" +
		"\n" + "# days ahead to check for" +
		"\n" + "days = 3",
)

func All() []string {
	return files
}
func Cache() string {
	return files[0]
}
func Config() string {
	return files[1]
}
func Events() string {
	return files[2]
}
func Ideas() string {
	return files[3]
}
func Recurrents() string {
	return files[4]
}
func Reminders() string {
	return files[5]
}
func Todos() string {
	return files[6]
}
func Wishlist() string {
	return files[7]
}
