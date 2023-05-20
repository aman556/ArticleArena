import React, { FunctionComponent } from "react";

import ArticleCard from "../ArticleCard";
import { UserArticles } from "../../models/Articles";
import useStyles from "./styles";

interface IArticleList {
  articlesList: UserArticles;
}

const ArticleList: FunctionComponent<IArticleList> = (props: IArticleList) => {
  const classes = useStyles();

  return (
    <>
      {props.articlesList?.ArticleData?.map((item) => (
        <ArticleCard name={item.ArticleTitle} link={item.ArtilceLink} />
      ))}
    </>
  );
};

export default ArticleList;
