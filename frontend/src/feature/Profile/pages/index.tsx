import React, { useEffect, useState } from "react";
import { Box, Button, TextField } from "@mui/material";
import clsx from "clsx";

import ArticleList from "../components/ArticleList";
import useStyles from "./styles";

const Profile: React.FC = () => {
  const classes = useStyles();

  const [input, setInput] = useState("");

  useEffect(() => {
    console.log("input", input);
  }, [input]);

  return (
    <Box className={clsx(classes.container)}>
      <Box className={clsx(classes.inputContainer)}>
        <TextField onChange={(event) => setInput(event.target.value)} />
        <Button variant="contained">Button</Button>
      </Box>
      <ArticleList articlesList={["Article1", "Article2", "Article3"]} />
    </Box>
  );
};

export default Profile;
