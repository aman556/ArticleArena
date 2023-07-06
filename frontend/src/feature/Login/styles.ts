import makeStyles from '@mui/styles/makeStyles';
import background from './Images/loginbackground.jpg';

const useStyles = makeStyles(() => ({
  container: {
    padding: '2rem',
    display: 'flex',
    flex: 1,
    flexDirection: 'column',
    backgroundColor: 'none',
    alignItems: 'center',
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    fontSize: '1.5rem',
  },
  email: {
    alignItems: 'center',
    padding: '0.3rem'
  },
  password: {
    alignItems: 'center',
    padding: '0.3rem'
  },
  login: {
    alignItems: 'center',
  },
  userimage: {
    height: 'auto',
    width: '5%',
  },
  inputfield: {
    width: '15rem',
    height: '1.5rem',
  },
  position: {
    backgroundImage: `url(${background})`,
    backgroundSize: 'cover',
    backgroundRepeat: 'no-repeat',
    alignItems: 'center',
    display: 'flex',
    height: '97.6vh',
    justifyContent: 'flex-start',
  }
}));

export default useStyles;