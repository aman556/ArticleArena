import React from 'react';
import UserCard from '../components/UserCard';
import { Box } from '@mui/material';
import useStyles from './styles';

const Profile: React.FC = () => {
  const styles = useStyles();

  return (
    <Box className={styles.container}>
      <UserCard
        name={'Kartikay Sharma'}
        address={'San Francisco, CA'}
        designation={'Full Stack Developer'}
      />
    </Box>
  );
};

export default Profile;
