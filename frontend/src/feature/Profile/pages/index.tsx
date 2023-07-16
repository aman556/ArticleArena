import React from 'react';
import UserCard from '../components/UserCard';
import { Box } from '@mui/material';
import useStyles from './styles';
import LogoutButton from '../../Header/LogoutButton';

const Profile: React.FC = () => {
  const styles = useStyles();

  return (
    <Box className={styles.container}>
      <LogoutButton />
      <UserCard
        name={'Kartikay Sharma'}
        address={'San Francisco, CA'}
        designation={'Full Stack Developer'}
      />
    </Box>
  );
};

export default Profile;
