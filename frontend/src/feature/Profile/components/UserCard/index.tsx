import React from 'react';
import userDefault from './../../../../constants/userDefault.png';
import { Box } from '@mui/material';
import useStyles from './styles';
import InfoPair from '../../../../components/InfoPair';

export interface IUserCard {
  name: string;
  designation: string;
  address: string;
}

const UserCard: React.FC<IUserCard> = (props) => {
  const styles = useStyles();

  return (
    <Box className={styles.container}>
      <img src={userDefault} height={500} width={500} />
      <Box className={styles.additionalText}>
        <InfoPair title={props.name} classes={{ title: styles.titleText }} />
        <InfoPair title={props.designation} />
        <InfoPair title={props.address} />
      </Box>
    </Box>
  );
};

export default UserCard;
