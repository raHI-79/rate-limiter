# Go Rate Limiter
A simple Rate Limiting project built with **Go** for the backend and **HTML, CSS, and JavaScript** for the frontend.

## Features
* Limits requests from each IP address
* Maximum **5 requests per minute**
* Returns **429 Too Many Requests** after the limit is reached
* Simple and easy-to-understand Go backend
* Simple frontend for testing
* Automatically allows requests again after 1 minute

## Project Structure
```text
rate-limiter/
│
├── main.go
└── index.html
```

## How It Works
The server checks the IP address of each client.

```text
Client
   ↓
Request
   ↓
Go Server
   ↓
Check IP + Request Count
   ↓
 ┌───────────────┐
 │ 5 requests?   │
 └───────┬───────┘
         │
    ┌────┴────┐
    ↓         ↓
 Allowed    Blocked
    ↓         ↓
Response     429
```

### Request Limit
```text
1 IP Address
     ↓
5 Requests/Minute
     ↓
6th Request
     ↓
429 Too Many Requests
```

After 1 minute, the request counter is reset.
## Technologies Used
* **Go** — Backend
* **HTML** — Frontend structure
* **CSS** — Frontend design
* **JavaScript** — Sending requests to the Go server

## How to Run
### 1. Clone the repository

```bash
git clone https://github.com/USERNAME/rate-limiter-go.git
```

Replace `USERNAME` with your GitHub username.
### 2. Run the Go server
```bash
go run main.go
```

The server will start at:
```text
http://localhost:8080
```

### 3. Open the Frontend
Open `index.html` in a web browser.

## Testing
Click the **Send Request** button multiple times.
```text
Request 1 → Accepted
Request 2 → Accepted
Request 3 → Accepted
Request 4 → Accepted
Request 5 → Accepted
Request 6 → 429 Too Many Requests
```

Wait for 1 minute and the requests will be allowed again.
## Purpose
This project demonstrates the basic concept of **Rate Limiting**, which helps prevent excessive requests from overwhelming a server.

## Author
**MD. RATUL ISLAM**
