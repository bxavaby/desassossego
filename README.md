**Desassossego** is a minimal TUI designed to immerse readers in the fragmented world of *O Livro do Desassossego* (*The Book of Disquiet*).

<br>

## **Demo**
![Desassossego Demo](assets/ldd.gif)  
*⊹ check out the suaveness ⊹*

<br>

## Features
- **Fragment Navigation**: Enter a fragment number (1-433) to read any part of the book.
- **Keyboard Shortcuts**:
  - **`n`**: Next fragment.
  - **`p`**: Previous fragment.
  - **`b`**: Back to fragment picker.
  - **`q`**: Quit program.
- **Local JSON Storage**: The fragments are stored in `frags/ldd.json`, for quick access.

<br>

## Installation

- Make sure [Go](https://golang.org/) is installed on your machine.

1. **Clone repo**:
   ```bash
   git clone https://github.com/bxavaby/desassossego.git
   cd desassossego
   ```
2. **Edit the absolute path to ldd.json in senon.go if you plan to run the program from a different directory:**
   ```bash
   fragments, err := ui.LoadFragments(
   		"frags/ldd.json",   // relative
   		"~/frags/ldd.json", // replace with your absolute path
   	)
   ```
3. **Build it**:
   ```bash
   go build -o desassossego
   ```
4. **Move the binary**:
   ```bash
   sudo mv desassossego /usr/local/bin/
   ```
5. **Run it**:
   ```bash
   desassossego
   ```
   
<br>

## Future Plans
- **SSH book server**: Transform "Desassossego" into an SSH-accessible book server.
- **Expand catalogue**: Use it as a template for other suitable books.

<br>

#

_Inspired by *O Livro do Desassossego* (*The Book of Disquiet*), Fernando Pessoa. This project is licensed under the MIT License._
