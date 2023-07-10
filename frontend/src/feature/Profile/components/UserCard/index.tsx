import React from 'react';
import userDefault from '../../../../constants/userDefault.png';
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
  const imageURL = 'https://images.unsplash.com/photo-1533035350251-aa8b8e208d95';

  return (
    <Box
        component="div"
        sx={{
          position: 'absolute',
          width: '99%',
          height: '98%',
          backgroundImage: `url(${imageURL})`,
          backgroundPosition: 'center',
          backgroundSize: 'cover',
          zIndex: '-1'
        }}
        className={styles.container}
      >
      <img src={userDefault} height={250} width={250} className={styles.imageStyles} alt='Kartikay ki fotu'/>
      <Box className={styles.additionalText}>
        <InfoPair title={props.name} classes={{ title: styles.titleText }} />
        <InfoPair title={props.designation} />
        <InfoPair title={props.address} />
      </Box>
    </Box>
  );
};

export default UserCard;
