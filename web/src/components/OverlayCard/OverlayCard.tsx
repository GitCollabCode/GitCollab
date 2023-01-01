import React, { useState } from 'react'
import {Navigate} from "react-router-dom";
import style from '../OverlayCard/OverlayCard.module.css'
import octocat from  "../../assets/octocat.png"


const OverlayCard = () => {
    const [showModal, setShowModal] = useState(true);

    //const toggleModal = () => {
    //    // set the state of the mogal
    //    setShowModal(!showModal);
    //}

    const onBackgroundClick = () => {
        // return user to home page
        return (
            <>
            {setShowModal(false)}
            <Navigate  to="/" />
            </>
        )
    }
//
    //const onReLoginBtnClick = () => {
    //    // attempt to log the user back in 
//
    //}

    if (showModal) {
        document.body.classList.add('active-modal');
    } else {
        document.body.classList.remove('active-modal');
    }
    
    return (
        <>
            <div className={style.modalContainer} onClick={onBackgroundClick}> 
                <div className={style.overlayCard}>
                    <div className={style.modalText}>
                        <img className={style.modalLogo} src={octocat}/>
                        <p className={style.modalTextContent}>
                            Sorry! You’re currently logged out, sign in to continue
                        </p>
                        <button className={style.modalButton}>
                            Login with GitHub
                        </button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default OverlayCard