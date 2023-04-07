import {
  Box,
  Button,
  Container,
  TextField,
  Typography,
  InputLabel,
} from "@mui/material";
import {
  ChangeEventHandler,
  FormEventHandler,
  MouseEventHandler,
  useEffect,
  useState,
} from "react";
import { Message } from "../../../types/Message";
import { Link, useLoaderData, useParams } from "react-router-dom";

function ChatBox() {
  const [message, setMessage] = useState("");
  const [messagesList, setMessagesList] = useState<Message[]>([]);
  const [webSocket, setWebSocket] = useState<WebSocket | null>(null);
  const [userName, setSenderName] = useState<string>("");
  const { chatName } = useParams<string>();

  useEffect(() => {
    fetch(`${process.env.API_URL}/messages/${chatName}`)
      .then((response) => response.json())
      .then((messages) => {
        console.log(messages);
        setMessagesList(messages);
      });
  }, []);

  useEffect(() => {
    setSenderName(localStorage.getItem("userName") || "anonymous");
    if (webSocket) {
      webSocket.onmessage = (event) => {
        const newMessage: Message = JSON.parse(event.data);
        setMessagesList((prevMessages) => [...prevMessages, newMessage]);
      };
    } else {
      const newWebSocket = new WebSocket(
        `${process.env.WEBSOCKET_URL}/chat/${chatName}`
      );
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

  const handleSendMessage: FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault();
    const newMessage = {
      sender: userName,
      content: message,
      sentAt: new Date(),
      chatName: chatName!,
    };
    if (webSocket && message.length > 0) {
      webSocket.send(JSON.stringify(newMessage));
      setMessage("");
    }
  };

  const handleDeleteMessage: MouseEventHandler<HTMLDivElement> = (event) => {
    const messageId = event.currentTarget.id;
    setMessagesList(
      messagesList.filter((message) => `${message.id}` !== messageId)
    );
  };

  return (
    <Container>
      <Link to="/">Home</Link>
      <Typography variant="h3">{chatName}</Typography>
      <Typography variant="body1">Sending messages as {userName}</Typography>
      <Box
        sx={{
          height: "600px",
          border: `2px solid`,
          borderColor: "primary.main",
          borderRadius: "10px",
          margin: "5px 0",
          padding: "10px",
        }}
      >
        {messagesList.map((message) => {
          return (
            <Box
              onClick={handleDeleteMessage}
              key={message.id}
              id={`${message.id}`}
              sx={{
                textAlign: message.sender == userName ? "left" : "right",
              }}
              marginBottom={3}
            >
              <Typography variant="inherit" fontWeight={700}>
                {message.sender}:
              </Typography>
              <Typography variant="inherit" margin={"0 20px"}>
                {message.content}
              </Typography>
            </Box>
          );
        })}
      </Box>
      <form action="" onSubmit={handleSendMessage}>
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
          <Button variant="contained" sx={{ minWidth: "50%" }} type="submit">
            Send
          </Button>
        </Box>
      </form>
    </Container>
  );
}

export default ChatBox;
