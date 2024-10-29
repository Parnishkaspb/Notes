# 📁 File Handling Pet Project

> A Go project focusing on mastering file operations, logging, and clean architecture principles. 🚀

## ✨ Overview
This project is designed to enhance your skills in working with files, utilizing Go's built-in logging, implementing **Clean Architecture**, and covering the entire codebase with unit tests. It's a step-by-step exploration of file handling in Go, aimed at building a solid foundation.

### 🎯 Key Features
- ✅ **File Writing:** Save structured data to files.
- ✅ **File Reading:** Load data from files seamlessly.
- ✅ **Logging:** Use Go's built-in `log` package for better insights.
- ✅ **Clean Architecture:** A clear separation of concerns following Clean Architecture principles.
- ✅ **Unit Testing:** Ensures robust and reliable code with 100% test coverage.

## 🏗️ Architecture
This project follows **Clean Architecture** principles, which promotes separation of concerns and high maintainability. The architecture consists of:
- **Entities:** Core business models and logic.
- **Use Cases:** Application logic handling file operations.
- **Interfaces:** Controllers, user interfaces, and gateways.
- **Infrastructure:** Implementation details like file I/O and logging.

## 📦 Installation

1. **Clone the repository:**
    ```bash
    git clone git@github.com:Parnishkaspb/Notes.git
    cd Notes
    ```

2. **Install dependencies:**
   Make sure you have Go installed (version 1.18+). Then run:
    ```bash
    go mod tidy
    ```

## 🚀 Usage
### 1️⃣ Writing Data to a File
Use the `SaveData` function to save structured data to a file:
```go
import (
    "Project2/internal/entities"
    "Project2/internal/infrastructure/file"
)

func main() {
    notes := entities.Notes{
        Note: map[int]interface{}{
            1: "First Note",
            2: "Second Note",
        },
    }
    
    fileHandler := file.NewWorkWithFileImpl()
    success, err := fileHandler.SaveData("example.txt", notes)
    if err != nil {
        log.Fatal("Failed to save data:", err)
    }
    if success {
        log.Println("Data saved successfully!")
    }
}
```
### 2️⃣ Reading Data from a File
Use the `LoadData` function to read structured data from a file
```go
import (
    "Project2/internal/infrastructure/file"
)

func main() {
   fileHandler := file.NewWorkWithFileImpl()
   notes, err := fileHandler.LoadData("example.txt")
   if err != nil {
   log.Fatal("Failed to load data:", err)
   }
   log.Printf("Loaded notes: %v", notes.Note)
}
```

### 🧪 Testing
The project uses the `testify` framework for comprehensive unit testing.

```bash
    go test ./... -v
```

### Happy Coding! ✨
