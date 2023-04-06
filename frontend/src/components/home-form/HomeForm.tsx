import { Box, Button, Container, TextField, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./HomeForm.module.css";

function HomeForm() {
  const [userName, setUserName] = useState("");
  const [chatName, setChatName] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    setUserName(localStorage.getItem("userName") || "");
  }, []);

  return (
    <Container
      sx={{
        maxWidth: "1000px",
        marginTop: "100px",
        display: "flex",
        alignItems: "center",
        flexDirection: "column",
      }}
      className={styles.homeFormContainer}
    >
      <Typography variant="h3" sx={{ marginBottom: 3 }}>
        Welcome
      </Typography>
      <form
        action=""
        onSubmit={(event) => {
          event.preventDefault();
          localStorage.setItem("userName", userName);
          navigate(`/chat/${chatName}`);
        }}
      >
        <Box>
          <TextField
            label="Username"
            onChange={(event) => {
              setUserName(event.target.value);
            }}
            value={userName}
          />
        </Box>
        <Box>
          <TextField
            label="Chat Name"
            onChange={(event) => {
              setChatName(event.target.value);
            }}
            value={chatName}
          />
        </Box>
        <Button variant="contained" type="submit">
          Join
        </Button>
      </form>
    </Container>
  );
}

export default HomeForm;
