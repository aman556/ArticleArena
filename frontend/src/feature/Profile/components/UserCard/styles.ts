import makeStyles from '@mui/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  container: {
    display: 'flex',
    flexDirection: 'column',
    margin: 5,
  },
  additionalText: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    background: 'pink',
  },
  titleText: {
    fontFamily: 'sans-serif',
    fontSize: 12,
  },
}));

export default useStyles;
