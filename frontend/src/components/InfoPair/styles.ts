import makeStyles from '@mui/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  container: {
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
    overflow: 'hidden',
  },
  title: {
    fontSize: 15,
    fontWeight: 400,
  },
  subtitle: {
    fontSize: 12,
    fontWeight: 400,
  },
}));

export default useStyles;
