import makeStyles from '@mui/styles/makeStyles';

const useStyles = makeStyles((theme) => ({
  usercard: {
    display: 'flex',
    flexDirection: 'column',
  },
  additionalText: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    width : '100%',
  },
  titleText: {
    fontFamily: 'sans-serif',
    fontSize: 12,
  },
  imageStyles: { border: 1, borderRadius: 250, margin: '5%' },
  cls_button: {
    padding: "5%",
    
  }
}));

export default useStyles;
