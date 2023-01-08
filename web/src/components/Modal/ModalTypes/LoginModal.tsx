import React  from 'react'

import style from '../Modal.module.css'

//import { ModalContextStateContext } from '../../../context/modalContext/modalContext'

const LoginModal = () => {    
  //const modalContext = useContext(ModalContextStateContext)
  //const [data, setData] = useState({ errorMessage: '', isLoading: false })
  
  return (
    <>
        <div className={style.modalText}>
        <p className={style.modalTextContent}>
            Tell us more about yourself
        </p>
        <div className={style.modalTextUnderline}/>
        <p className={style.modalTextContent}>
            Select a few topics that interest you
        </p>
        <div className={style.skillButtonContainer}>
            <button className={[style.modalText, style.skillButton].join(" ")}>
                Contin
            </button>
            <button className={[style.modalText, style.skillButton].join(" ")}>
                cheese
            </button>
            <button className={[style.modalText, style.skillButton].join(" ")}>
                testing stuff
            </button>
            <button className={[style.modalText, style.skillButton].join(" ")}>
                apple sauce
            </button>
        </div>
    
        <button className={style.modalButton}>
              Continue
          </button>
      </div>
      </>
    ) 
}

export default LoginModal