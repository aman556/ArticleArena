import React, { FunctionComponent } from "react";
import { Box, Button, TextField } from "@mui/material";
import clsx from "clsx";

import useStyles from "./styles";

interface IArticleList {
  articlesList: Array<String>;
}

const ArticleList: FunctionComponent<IArticleList> = (props: IArticleList) => {
  const classes = useStyles();

  return (
    <>
      {props.articlesList.map((item) => (
        <h2>{item}</h2>
      ))}
    </>
  );
};

export default ArticleList;
