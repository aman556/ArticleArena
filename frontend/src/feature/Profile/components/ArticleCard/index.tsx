import React, { FunctionComponent } from "react";
import StarBorderRoundedIcon from "@mui/icons-material/StarBorderRounded";
import clsx from "clsx";

import useStyles from "./styles";
import { Box, Paper, Typography } from "@mui/material";

interface IArticleCard {
  name: string;
  link: string;
}

const ArticleCard: FunctionComponent<IArticleCard> = (props: IArticleCard) => {
  const classes = useStyles();

  return (
    <Paper
      className={clsx(classes.container)}
      variant="outlined"
      elevation={24}
    >
      <Typography className={clsx(classes.articleName)}>
        {props.name}
      </Typography>
      <Box
        className={clsx(classes.additionalInfoContainer)}
        borderBottom={0.5}
        borderColor="grey"
      >
        <StarBorderRoundedIcon />
        <Box className={clsx(classes.additionalInfoTextContainer)}>
          <Typography>Created on -</Typography>
        </Box>
      </Box>
      <Box></Box>
    </Paper>
  );
};

export default ArticleCard;
