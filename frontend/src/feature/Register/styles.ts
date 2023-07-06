import makeStyles from '@mui/styles/makeStyles';
import background from './Images/background.png'

const useStyles = makeStyles(() => ({
  container: {
    padding: '2rem',
    display: 'flex',
    flex: 1,
    flexDirection: 'column',
    backgroundColor: 'none',
    alignItems: 'center',
    position: 'absolute',
    fontSize: '1.5rem'
  },
  username: {
    alignItems: 'center',
    padding: '0.3rem'
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
    padding: '0.3rem',
  },
  userimage: {
    height: 'auto',
    width: '20%',
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
    justifyContent: 'center',
    backgroundColor: 'deepskyblue'
  }
}));

export default useStyles;