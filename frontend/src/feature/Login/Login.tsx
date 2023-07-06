import React from "react"
import useStyles from "./styles";
import { NavLink, useHref } from "react-router-dom";
import userimage from './Images/user.png'

const Login: React.FC =() => {
    const styles = useStyles();

  return (
    <div className={styles.position}>
        <form className={styles.container}>
            <img src={userimage} className={styles.userimage}></img>
            <div className={styles.email}>
                <label>Email </label>
                <input className= {styles.inputfield} type="email" placeholder="Enter your email" autoComplete="off"></input>
            </div>
            <div className={styles.password}>
                <label>Password </label>
                <input className={styles.inputfield} type="password" placeholder="Enter your Password" autoComplete="off" ></input>
            </div>
            <div className={styles.login}>
                <button type="button">Login</button>
            </div>
            <div  className={styles.login}>
                <NavLink to='/register'>
                    <button type="button">SignUp</button>
                </NavLink>
            </div>
        </form>
    </div>
    )
}

export default Login