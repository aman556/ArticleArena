import React from "react"
import useStyles from "./styles"
import { NavLink } from "react-router-dom";
import userimage from "./Images/note.png"
import GoogleSignUp from "../Login/GoogleSignUp";

const Register: React.FC =() => {
    const styles= useStyles();
    return (
        <div className={styles.position}>
            <form className={styles.container}>
                <img className={styles.userimage} src={userimage}></img>
                <div className={styles.username}>
                    <label>User Name </label>
                    <input className= {styles.inputfield} type="text" placeholder="Enter User Name" ></input>
                </div>
                <div className={styles.email}>
                    <label>Email </label>
                    <input className= {styles.inputfield} type="email" placeholder="Enter Email"></input>
                </div>
                <div className={styles.password}>
                    <label>Password </label>
                    <input className= {styles.inputfield} type="password" placeholder="Enter Password"></input>
                </div>
                <div className={styles.password}>
                    <label>Confirm Password </label>
                    <input className= {styles.inputfield} type="password" placeholder="Enter Confirm Password"></input>
                </div>
                <div className={styles.login}>
                    <button type="button"> Register</button>
                </div>
                <div className={styles.login}>
                    <NavLink to='/login'>
                        <button type="button">Back To Login</button>
                    </NavLink>
                </div>
                <div>
                    <GoogleSignUp/>
                </div>
            </form>
        </div>
    )
}

export default Register