# Resource Creator

Resource Creator is a powerful tool designed to streamline the creation of Minecraft Resource Packs. It automates the
tedious parts of resource pack creation, so you can focus on creativity.

![A screenshot of Resource Creator](docs/banner.png)

## Features

### CTM Pattern Texture Generator

- **Automated Processing:** Generate a complete pattern texture from a single source image.
- **Smart Splitting:** Automatically splits the pattern and generates the necessary .properties file.
- **Instant Preview:** Quickly preview your generated pattern within the application.

### Alternate Texture Generator

- Coming Soon: Stay tuned for upcoming features.

> **Note:** Resource Creator is still in development, and some features may not be available yet.

## Development Setup

Resource Creator is built with [Wails](https://wails.app), a Go framework for desktop applications, and features a
modern frontend powered by Vite and React.

### Prerequisites

- [Go](https://golang.org/dl/)
- [Node.js](https://nodejs.org/en/download/)
- [Wails](https://wails.io/docs/gettingstarted/installation)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/sotterbeck/resource-creator.git
   ```

2. Navigate to the project directory:
   ```bash
   cd resource-creator
   ```
3. Run all tests to ensure everything is working correctly:
   ```bash
   go test ./...
   ```

Now you all set up. You can start the application in development mode or build it for production. For more information,
see below.

### Live Development

For rapid development and testing:

1. Navigate to the project directory.
2. Run:

    ```bash
    wails dev
    ```
   This command:
    - Launches a Vite development server with fast hot-reloading for frontend changes.
    - Starts a dev server on http://localhost:34115 to allow direct interaction with your Go backend methods via your
      browser’s devtools.

### Building for Production

To create a production-ready distributable:

```bash
wails build
```

This command compiles the application into a standalone package for redistribution.