# cli-todo  💼
cli-todo is an appealing CLI todo applcation and the purpose of it's creation was to understand golang and how projects are structured in golang.

<img width="870" alt="Screenshot 2023-04-30 at 9 35 13 AM" src="https://user-images.githubusercontent.com/62211740/235365155-755df511-372a-4963-8ea3-4ebfb8f1a9d5.png">

## Steps to run the application

`git clone https://github.com/Ashfaq97/cli-todo.git`

`cd cli-todo`

`go build ./cmd/todo`

To add a todo item:
`./todo -add complete COMP610 HW` 

To list an item as completed:
`./todo -complete=1` 

To list all todos:
`./todo -list`

To delete an item from the list:
`./todo -delete=1`
