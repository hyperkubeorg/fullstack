# The Guide to a Fullstack Project with Typescript & Go
At the end of this guide, you should have produced a `Typescript/Go` mono repo containing a frontend and backend.
## Frontend
The root directory of the project should be a go module.
```sh
go mod init github.com/hyperkubeorg/fullstack
```

We setup our frontend with `vite`, `react`, and `tailwindcss`
```sh
npm create vite@latest -- ./frontend -t react-ts
cd frontend
npm install -D @types/node react-router-dom tw-animate-css clsx tailwind-merge tailwindcss @tailwindcss/vite lucide
```

Update `frontend/tsconfig.json`.
```json
{
  "files": [],
  "references": [
    {
      "path": "./tsconfig.app.json"
    },
    {
      "path": "./tsconfig.node.json"
    }
  ],
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": [
        "./src/*"
      ]
    }
  }
}
```

Update `frontend/tsconfig.app.json`.
```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": [
        "./src/*"
      ]
    },
    "tsBuildInfoFile": "./node_modules/.tmp/tsconfig.app.tsbuildinfo",
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": [
      "ES2020",
      "DOM",
      "DOM.Iterable"
    ],
    "module": "ESNext",
    "skipLibCheck": true,
    /* Bundler mode */
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "verbatimModuleSyntax": true,
    "moduleDetection": "force",
    "noEmit": true,
    "jsx": "react-jsx",
    /* Linting */
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "erasableSyntaxOnly": true,
    "noFallthroughCasesInSwitch": true,
    "noUncheckedSideEffectImports": true
  },
  "include": [
    "src"
  ]
}
```

Update `frontend/vite.config.ts`.
```typescript
import path from "path"
import tailwindcss from '@tailwindcss/vite'

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  server: {
    proxy: {
      // this is a proxy for the backend during development
      // it will forward requests from /api to the backend server
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        // strip the /api prefix
        // rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
```

Update `frontend/src/App.tsx`.
```typescript
import '@/App.css'
import { Routes, Route, BrowserRouter } from "react-router-dom";
import NotFoundPage from '@/components/pages/404';
import TodoPage from '@/components/pages/todo';

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<TodoPage />} />

        {/* Catch-all 404 Route - MUST REMAIN AT BOTTOM */}
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}
```

Create `frontend/src/components/pages/404.tsx`.
```typescript
export default function NotFoundPage() {
  return (
    <>
      <h1>404 - Page Not Found</h1>
      <p>The page located at <code>{window.location.href}</code> was not found.</p>
    </>
  );
}
```

Create `frontend/src/components/pages/todo.tsx`. This is a placeholder page.
```typescript
import { useState } from 'react'
import reactLogo from '@/assets/react.svg'
import viteLogo from '/vite.svg'

export default function TodoPage() {
  const [ count, setCount ] = useState(0);

  return (
    <>
      <div className="flex items-center justify-center space-x-4 my-4">
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo w-24 h-24" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react w-24 h-24" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}
```

Create `frontend/src/lib/utils.ts`.
```typescript
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
```

Prepend the following to `frontend/src/App.css`.
```css
@tailwind utilities;

@import "tailwindcss";
@import "tw-animate-css";
```

## Backend

Create `main.go`.
```go
package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hyperkubeorg/fullstack/backend"
	"github.com/hyperkubeorg/fullstack/frontend"
)

var BIND_ADDR = func() string {
	if addr := os.Getenv("BIND_ADDR"); addr != "" {
		return addr
	}
	return "127.0.0.1:8080"
}()

var DISABLE_FRONTEND = func() bool {
	value := os.Getenv("DISABLE_FRONTEND")
	if value != "" {
		if v, err := strconv.ParseBool(value); err == nil {
			return v
		}
	}
	return false
}()

func main() {
	r := mux.NewRouter().StrictSlash(true)

	if _, err := backend.AddRoutes(r); err != nil {
		log.Fatalf("Failed to add routes: %v", err)
	}

	if !DISABLE_FRONTEND {
		if _, err := frontend.AddRoutes(r); err != nil {
			log.Fatalf("Failed to add frontend routes: %v", err)
		}
	}

	log.Printf("Starting server on %s", BIND_ADDR)
	log.Fatal(http.ListenAndServe(BIND_ADDR, r))
}
```


Create `backend/routes.go`.
```go
package backend

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const failedToEncodeResponse = `{"status": "FAILURE", "message": "Failed to encode response.", "data": {}}`

func AddRoutes(r *mux.Router) (*mux.Router, error) {
	r.HandleFunc("/api/v1/todo", todoHandler).Methods("GET")

	return r, nil
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	// serialize and send the response
	write(w, http.StatusOK,
		struct {
			Status  string      `json:"status"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Status:  "SUCCESS",
			Message: "This is a placeholder handler.",
			Data: map[string]interface{}{
				"current_time_utc": time.Now().UTC().Format(time.RFC3339),
			},
		},
	)
}

func write(w http.ResponseWriter, status int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, failedToEncodeResponse, http.StatusInternalServerError)
		return
	}
}
```

Update the go modules.
```sh
go mod tidy
```
## Serving the Frontend
This fullstack solution serves the static frontend from the same binary that has the backend code. This glue code might not be necessary in the future, but for now the most important things it does is:
1. Embeds the frontend into the server binary written in Go.
2. Defines which routes the server will distribute `index.html` from in the server binary.

Create `frontend/frontend.go`.
```go
package frontend

import (
	"bytes"
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed dist/**
var dist embed.FS

// TODO: Needs documentation
var Content = func() fs.FS {
	subFS, err := fs.Sub(dist, "dist")
	if err != nil {
		log.Fatalf("Failed to create sub file system (static): %v", err)
	}

	// list all files in views.EmbeddedFS
	err = fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			Manifest = append(Manifest, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to walk embedded filesystem: %v", err)
	}

	return subFS
}()

var Manifest = []string{}

func serveIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	file, err := Content.Open("index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Printf("Error getting file stat for index.html: %v", err)
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
		return
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file content for index.html: %v", err)
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, "index.html", stat.ModTime(), bytes.NewReader(content))
}

// Updates the 404 handler to serve the index.html.
// This must be set after all other routes.
// If this is done before any routes are registered, the 404 handler will be called unexectedly.
func AddRoute404(r *mux.Router) *mux.Router {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Set 404 status in the headers
		serveIndexHandler(w, r)
	})
	return r
}

// Adds frontend routes to the router.
func AddRoutes(r *mux.Router) (*mux.Router, error) {
	// Server production ready UI paths
	// When a new route is added to App.tsx, it should be added here as well
	r.HandleFunc("/", serveIndexHandler).Methods("GET")

	// list all files in views.EmbeddedFS
	err := fs.WalkDir(Content, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// Ignore index.html, it is configured to be served by multiple routes
			if path == "index.html" {
				return nil
			}

			log.Printf("Serving %s", "/"+path)

			r.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
				file, err := Content.Open(path)
				if err != nil {
					log.Printf("Error opening file %s: %v", path, err)
					http.Error(w, "File not found", http.StatusNotFound)
					return
				}
				defer file.Close()

				stat, err := file.Stat()
				if err != nil {
					log.Printf("Error getting file stat for %s: %v", path, err)
					http.Error(w, "Failed to get file info", http.StatusInternalServerError)
					return
				}

				content, err := io.ReadAll(file)
				if err != nil {
					log.Printf("Error reading file content for %s: %v", path, err)
					http.Error(w, "Failed to read file", http.StatusInternalServerError)
					return
				}
				http.ServeContent(w, r, path, stat.ModTime(), bytes.NewReader(content))
			})
		}
		return nil
	})

	if err != nil {
		log.Fatalf("failed to walk embedded filesystem: %v", err)
	}

	return r, err
}
```


## Developer Tooling

Create `Makefile`.
```sh
# Friendly help message.
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build-frontend   Build the frontend"
	@echo "  run-frontend     Run the frontend in development mode"
	@echo "  run-backend      Run the backend"
	@echo "  services-up      Start the service stack"
	@echo "  services-down    Stop the service stack"
	@echo "  compose-up       Start the backend with Docker Compose"
	@echo "  compose-down     Stop the backend with Docker Compose"
	@echo "  help             Show this help message"
	@echo ""


# Build the frontend
build-frontend:
	cd frontend && npm install && npm run build

# Run the frontend in development mode
run-frontend:
	cd frontend && npm run dev

# Run the backend
run-backend: build-frontend
	go mod tidy && DB_AUTO_INITIALIZE_SCHEMA=true DB_AUTO_INITIALIZE_SCHEMA_DROP=true go run .

# Run only the service
services-up:
	docker compose -f docker-compose.deps.yml up

# Tear down the service stack
services-down:
	docker compose -f docker-compose.deps.yml down

# Run the backend with docker compose
compose-up:
	docker compose up --build

# Tear down the backend with docker compose
compose-down:
	docker compose down
```

Create `Dockerfile`.
```dockerfile
## This builds a static frontend.
FROM node:22 AS node-builder
WORKDIR /usr/src/app
COPY ./frontend ./frontend
WORKDIR /usr/src/app/frontend
RUN npm install
RUN npm run build
RUN npm prune --production

# This builds the Go binary with static frontend embedded.
FROM golang:1.24 AS go-builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=node-builder /usr/src/app/frontend/dist ./frontend/dist
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -v -o /usr/local/bin/main .
RUN chmod +x /usr/local/bin/main

# This produces a very tiny image with only the binary and no dependencies.
FROM scratch
COPY --from=go-builder /usr/local/bin/main /usr/local/bin/main
USER 1000:1000
WORKDIR /data
ENTRYPOINT ["/usr/local/bin/main"]
```

Create `docker-compose.deps.yml`.
```yaml
services:
  yugabyte:
    image: "yugabytedb/yugabyte:2.25.1.0-b381"
    ports:
      - "7000:7000"
      - "9000:9000"
      - "15433:15433"
      - "5433:5433"
      - "9042:9042"
    command: bin/yugabyted start --background=false
```

Create `docker-compose.yml`.
```yaml
services:
  yugabyte:
    image: "yugabytedb/yugabyte:2.25.1.0-b381"
    ports:
      - "7000:7000"
      - "9000:9000"
      - "15433:15433"
      - "5433:5433"
      - "9042:9042"
    command: bin/yugabyted start --background=false
  fullstack:
    build:
      context: .
      args:
        BINARY_NAME: "fullstack"
    environment:
      BIND_ADDR: "0.0.0.0:8080"
      DATABASE_URI: "host=yugabyte port=5433 user=yugabyte password= dbname=yugabyte sslmode=disable"
      DB_AUTO_INITIALIZE_SCHEMA: "true"
    ports:
      - "8080:8080"
    command: main
    depends_on:
      - yugabyte
```

## Developer Quickstart
You can build the whole project and run it in a container.
```shell
make compose-up # Control+C to stop
make compose-down # Purges containers and data.
```

The prior method doesn't rely on any cached modules and is slow for iterating.

### Iterating Fast
If you would like to iterate fast, open 3 terminal tabs and a code editor fixed on the root directory of this project.

First terminal, start the data services for development.
```shell
make services-up # Control+C to stop
make services-down # Purges containers and data.

# These are services you might want to access during development.
docker exec -it fullstack-yugabyte-1 ysqlsh -h yugabyte
```

Second terminal, start the backend. You will CTRL+C this and restart it for backend changes.
```shell
make run-backend
```

Third terminal, start the frontend. You will navigate this in your browser. Vite server is configured to proxy `/api` to the backend you have in the last terminal.
```shell
make run-frontend
```

At this point if you're working on a component or part of the frontend, that will auto refresh in the browser. You still have to partly touch the terminal for the backend changes to reflect, but this is guaranteed to be quick to iterate on.

## Points of Interest
Once you have the application stack running, you should open up the following urls.

The frontend is Vite+React. Note, your port number might be different, check your terminal output from `make run-frontend` to verify.
1. http://localhost:5173/
    - This is Vite server.
    - It's only for development purposes.
    - It supports hot reloading.
2. http://localhost:8080/
    - This is the Go backend.
    - The GUI you see here has been compiled.
    - YOU DO NOT GET HOT RELOADING HERE!
3. http://localhost:8080/api/v1/todo
    - This is the example route we have implemented for your viewing pleasure.
4. http://localhost:5173/api/v1/todo
    - The path `/api` for Vite is configured to proxy to the backend running on 8080.
    - This is probably why you never need to really go to the backend in your browser.

## (Un)Conventional Design Choices
This project heavily relies on containerization to ensure the final result is a container. Containers are extremely portable, but the larger the container is, the longer it will take to scale up when in a time of urgency.

Go seems like a really good choice for any containerized backend. It produces very small binaries that don't require any of the linux system to be included in the container. Also it implements concurrency in a way that takes advantage of all the threads on a system.

Dealing with templates in Go is a bit of a hassle. By adopting Typescript into the project with some light integrations into the repository, Go templates can be completely avoided.

A hard design choice made in this example is the use of environment variables. The idea here is to avoid taking on the hassle of managing a CLI tool and narrow all focus into building a service. All tooling should be implemented in the frontend and/or backend.

Everything should be implemented with the expectation it should maintain it's own relevant environment variables close to where they get consumed.

Nothing in this project is expected to be DRY, we just want to write code to balance developer velocity with complexity.
