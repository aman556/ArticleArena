import makeStyles from "@mui/styles/makeStyles";

const useStyles = makeStyles((theme) => ({
  container: {
    height: 80,
    display: "flex",
    flexDirection: "column",
    backgroundColor: "pink",
    // borderColor: "black",
    padding: 15,
    margin: 4,
  },
  articleName: {
    height: 40,
    width: "100%",
    // backgroundColor: "red",
  },
  additionalInfoContainer: {
    height: 60,
    width: "100%",
    display: "flex",
    flexDirection: "row",
  },
  additionalInfoTextContainer: {
    height: 40,
    borderColor: "red",
    borderWidth: 50,
  },
}));

export default useStyles;
