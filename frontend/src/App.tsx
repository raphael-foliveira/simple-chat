import "./App.css";
import ChatBox from "./components/chatbox/ChatBox";
import { grey } from "@mui/material/colors";
import { ThemeProvider, createTheme, Typography, Box } from "@mui/material";

function App() {
  return (
    <div className="App">
      <Typography align="center" margin={6} variant="h3">
        Local Chat App
      </Typography>
      <ChatBox />
    </div>
  );
}

export default App;
