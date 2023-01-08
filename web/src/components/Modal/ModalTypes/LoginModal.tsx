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
          {// fix underline
          }
          <div className={style.modalTextUnderline}></div>
          <p className={style.modalTextContent}>
            Select a few topics that interest you
          </p>
        {
            // add skills
        }
          <button className={style.modalButton}>
              Continue
          </button>
      </div>
      </>
    ) 
}

export default LoginModal