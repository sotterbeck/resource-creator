import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import App from "./App";
import { HashRouter } from "react-router";
import { ThemeProvider } from "@/components/theme-provider";

const container = document.getElementById("root");

const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <HashRouter>
      <ThemeProvider>
        <App />
      </ThemeProvider>
    </HashRouter>
  </React.StrictMode>,
);
