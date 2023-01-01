import React, { ReactNode, useState } from 'react'
import  ReactDOM  from 'react-dom';
import {useNavigate} from "react-router-dom";
import style from '../Modal/Modal.module.css'
import octocat from  "../../assets/octocat.png"

const LoggedOutModal = () => {    
    return (
        <>
        <div className={style.modalText}>
            <img className={style.modalLogo} src={octocat}/>
            <p className={style.modalTextContent}>
                Sorry! You’re currently logged out, sign in to continue
            </p>
            <button className={style.modalButton}>
                Login with GitHub
            </button>
        </div>
        </>
    ) 
}

export default LoggedOutModal