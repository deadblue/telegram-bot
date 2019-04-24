/*
This package contains all arguments structs for Telegram Bot API mehtod.

All arguments structs are named with "Args" suffix.

They can be simply constructed by using new() operator,
then developer can call related method to set argument value.

Example:

	args := new(arguments.SendMessageArgs)
	args.ChatId(1)
	args.Text("Hello, world")

Note:

* All arguments structs have Archive() method which will be called
by core code, developer SHOULD NOT call it.

* All arguments structs are NOT goroutine-safety, you SHOULD avoid
manipulating it in mulit-goroutine.

*/
package arguments
