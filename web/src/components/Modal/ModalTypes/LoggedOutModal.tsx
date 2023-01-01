import React, { ReactNode, useState } from 'react'
import  ReactDOM  from 'react-dom';
import {useNavigate} from "react-router-dom";
import style from '../Modal/Modal.module.css'
//import octocat from  "../../assets/octocat.png"
//const modalRoot = document.getElementById('modal-root');

const LoggedOutModal = ({children}:{children : ReactNode}) => {
    const [displayModal, setdisplayModal] = useState(true);
    
    return (
        <>
        </>
    ) 
}

export default LoggedOutModal