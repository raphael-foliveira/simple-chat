import { Box, Button, Container, TextField, Typography, InputLabel } from "@mui/material";
import { ChangeEventHandler, MouseEventHandler, useEffect, useState } from "react";
import { Message } from "../../../types/Message";

function ChatBox() {
  const [message, setMessage] = useState("");
  const [messagesList, setMessagesList] = useState<Message[]>([]);
  const [webSocket, setWebSocket] = useState<WebSocket | null>(null);
  const [senderName, setSenderName] = useState<string>("");

  useEffect(() => {
    if (webSocket) {
      webSocket.onmessage = (event) => {
        const newMessage: Message = JSON.parse(event.data);
        setMessagesList((prevMessages) => [...prevMessages, newMessage]);
      };
    } else {
      const newWebSocket = new WebSocket("ws://localhost:8000/chat");
      setWebSocket(newWebSocket);
    }
    return () => {
      if (webSocket) {
        webSocket.close();
      }
    };

  }, [webSocket]);

  const handleChange: ChangeEventHandler<
    HTMLTextAreaElement | HTMLInputElement
  > = (event) => {
    setMessage(event.target.value);
  };

  const handleSendMessage: MouseEventHandler<HTMLButtonElement> = (event) => {
    const newMessage: Message = {
      id: `${Math.floor(Math.random() * 1000)}`,
      sender: senderName,
      content: message,
      sentAt: new Date(),
    };
    if (webSocket && message.length > 0) {
      webSocket.send(JSON.stringify(newMessage));
      setMessage("");
    }
  };

  const handleDeleteMessage: MouseEventHandler<HTMLDivElement> = (event) => {
    const messageId = event.currentTarget.id;
    setMessagesList(messagesList.filter((message) => message.id !== messageId));
  };

  return (
    <Container>
      <TextField label="Name" onChange={(e) => setSenderName(e.target.value)} />
      <Box
        sx={{
          height: "600px",
          border: `2px solid`,
          borderColor: 'primary.main',
          borderRadius: "10px",
          margin: "5px 0",
          padding: "10px",
        }}
      >
        {messagesList.map((message) => {
          return (
            <Typography
              onClick={handleDeleteMessage}
              key={message.id}
              id={message.id}
              sx={{ textAlign: message.sender == senderName ? "left" : "right" }}
              variant="body1"
              marginBottom={3}
            >
              <Typography variant="inherit" fontWeight={700}>{message.sender}:</Typography>
              <Typography variant="inherit" margin={"0 20px"}>{message.content}</Typography>
            </Typography>
          );
        })}
      </Box>
      <Box
        sx={{
          display: "flex",
          flexWrap: "wrap",
          justifyContent: "center",
        }}
      >
        <TextField
          id="outlined-basic"
          label="Message"
          value={message}
          onChange={handleChange}
          sx={{
            width: "100%",
            marginBottom: "5px",
          }}
        />
        <Button
          variant="contained"
          onClick={handleSendMessage}
          sx={{ minWidth: "50%" }}
        >
          Send
        </Button>
      </Box>
    </Container>
  );
}

export default ChatBox;
