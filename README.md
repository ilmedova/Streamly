# 🌎 Streamly

A simple Go-based news aggregator that fetches and filters news articles from various sources, such as BBC and The New York Times, and outputs the results to a `news.md` file. The aggregator supports keyword filtering and can be extended with additional news sources via drivers.

### ✨ Features
- 🏛 Fetches news articles from BBC and The New York Times.
- 📄 Outputs results to a `news.md` file.
- 🔍 Supports keyword filtering.
- 🔌 Easily extendable with new drivers in `scraper` folder.

### 📥 Installation
Make sure you have [Go](https://golang.org/doc/install) installed.

Clone the repository:
```sh
git clone https://github.com/yourusername/news-aggregator.git
cd news-aggregator
```

### 🚀 Usage
Run the aggregator:
```sh
go run main.go
```

### 📦 Docker Instructions
Build the Docker image:
```sh
docker build -t streamly .
```

Run the container:
```sh
docker run --rm streamly
```

### 🛠 Extending with New Drivers
You can add new news sources by implementing a new driver in the `scraper/` directory. Each driver should follow the existing structure to ensure compatibility.

### 📜 License
This project is licensed under the MIT License.

### 🤝 Contributing
Contributions are welcome! Feel free to submit a pull request or open an issue to suggest improvements.

### 👤 Author
[Mahri Ilmedova](https://github.com/ilmedova)
