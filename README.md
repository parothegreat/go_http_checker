# 🚀 Concurrent HTTP Status Checker

A high-performance command-line interface (CLI) tool written in Go (Golang) for concurrently checking the HTTP status codes of multiple URLs. This project serves as an introductory practice into Go's powerful concurrency model using **Goroutines** and **Channels**.

---

## ✨ Features

* **Concurrent Checking:** Utilizes Goroutines to perform HTTP requests simultaneously, significantly reducing total check time.
* **Status Classification:** Categorizes results into Success (2xx), Client Error (4xx), Server Error (5xx), and Network Errors/Timeouts.
* **Timeout Control:** Implements a 5-second timeout for requests to prevent infinite waiting on unreachable URLs.

---

## ⚙️ Prerequisites

Before running this application, ensure you have the following installed:

* **Go (Golang):** Version 1.18 or higher (required for Go Modules).
* **Git:** For cloning the repository and managing versions.

---

## 🛠️ Installation and Usage

### 1. Clone the Repository

Clone the project to your local machine:

```bash
git clone https://github.com/parothegreat/go-http-checker.git
cd go-http-checker