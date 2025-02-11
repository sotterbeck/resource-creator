# Resource Creator

Resource Creator is a powerful tool designed to streamline the creation of Minecraft Resource Packs. It automates the
tedious parts of resource pack generation, so you can focus on creativity.

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