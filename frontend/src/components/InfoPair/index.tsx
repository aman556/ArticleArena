import React, { FunctionComponent } from 'react';
import { Box, Typography } from '@mui/material';
import clsx from 'clsx';

import useStyles from './styles';

type ClassType = {
  container?: string;
  title?: string;
  subtitle?: string;
};

interface IInfoPair {
  title?: string;
  subtitle?: string;
  classes?: ClassType;
}

const InfoPair: FunctionComponent<IInfoPair> = (props) => {
  const classes = useStyles();

  const getTitle = () => (
    <Typography className={clsx(classes.title, props.classes?.title)} noWrap>
      {props.title ?? '-'}
    </Typography>
  );

  const getSubtitle = () => (
    <Typography
      className={clsx(classes.subtitle, props.classes?.subtitle)}
      noWrap
    >
      {props.subtitle ?? ''}
    </Typography>
  );

  return (
    <Box className={clsx(classes.container, props.classes?.container)}>
      {props.title && getTitle()}
      {props.subtitle && getSubtitle()}
    </Box>
  );
};

export default InfoPair;
