import React, { ReactNode, useState } from 'react'
import  ReactDOM  from 'react-dom';
import {useNavigate} from "react-router-dom";
import style from '../Modal/Modal.module.css'
//import octocat from  "../../assets/octocat.png"
//const modalRoot = document.getElementById('modal-root');

const Modal = ({children}:{children : ReactNode}) => {
    const [displayModal, setdisplayModal] = useState(true);
    const navigate = useNavigate();
    
    const hideModal = () => {
        setdisplayModal(false);
    }

    //const showModal = () => {
    //    setdisplayModal(true);
    //}

    const onBackgroundClick = () => {
        navigate("/")
        hideModal();
    }
    
    return ( displayModal ? 
        ReactDOM.createPortal(
            <>
                <div className={style.modalContainer} > 
                    <div className={style.bg} onClick={onBackgroundClick}></div>
                    <div className={style.overlayCard}>
                        {children}
                    </div>
                </div>,
            </>,
            document.getElementById('modal-root') as Element
        )  : <></>
    ) 
}

export default Modal