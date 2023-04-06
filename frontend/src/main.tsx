import { createTheme, ThemeProvider } from "@mui/material";
import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import ChatBox from "./components/chatbox/ChatBox";
import HomeForm from "./components/home-form/HomeForm";
import "./index.css";

const theme = createTheme({
  palette: {
    mode: "light",
    background: {
      default: "#fd8877",
    },
  },
});

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomeForm />} />
          <Route path="/chat/:chatName" element={<ChatBox />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  </React.StrictMode>
);
