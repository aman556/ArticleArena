import React, { useEffect, useState } from "react";
import { Box, Button, TextField } from "@mui/material";
import clsx from "clsx";

import ArticleList from "../components/ArticleList";
import { UserArticles } from "../models/Articles";
import useStyles from "./styles";

const Profile: React.FC = () => {
  const classes = useStyles();

  const [input, setInput] = useState("");
  const [articleData, setArticleData] = useState<UserArticles>(
    {} as UserArticles
  );

  useEffect(() => {
    console.log("input", input);
  }, [input]);

  const fetchGFGData = () => {
    fetch("http://localhost:8081/geeksforgeeksUserInfo/".concat(input))
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        setArticleData(data);
      });
  };

  return (
    <Box className={clsx(classes.container)}>
      <Box className={clsx(classes.inputContainer)}>
        <TextField onChange={(event) => setInput(event.target.value)} />
        <Button variant="contained" onClick={fetchGFGData}>
          Get GFG Articles
        </Button>
      </Box>
      <ArticleList articlesList={articleData} />
    </Box>
  );
};

export default Profile;
