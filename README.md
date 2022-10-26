#A Simples Web Server with GoLang

At Trebas (Computer Hardware & Software Environment), the challenger proposed by Ghazal G.Fard starts with a WebServer in GoLang.

Before commenting on the code, we need to understand two differences between python and GoLang.

1) To "ignore" a line in Python we use # and in Golang we use // Some other languages like C, use #, and others, like PHP, use // or /* */.

2) In python I don't need to define the main() function.

We need to keep in mind that we need to develop clean code. For example, is a good practice to use MVC (Model, View, Controller) techniques.

On the code, we need to import three packages: "fmt", "log", "net/http"

The fmt package we use to give a response. Can be on screen, or in this case, a web response. For the log package, we are just writing a log. The net/http we use to start a webserver.

There are a lot of concepts that we can present in the future.
