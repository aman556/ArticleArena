import makeStyles from '@mui/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  container: {
    display: 'flex',
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    height: 800,
  },
  additionalText: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    background: 'white',
    marginTop: '5%',
    width : '20%',
  },
  titleText: {
    fontFamily: 'sans-serif',
    fontSize: 12,
  },
  imageStyles: { border: 1, borderRadius: 250, marginTop: '5%' },
}));

export default useStyles;
