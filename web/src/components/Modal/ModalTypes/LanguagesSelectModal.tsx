import React, { useEffect, useState, useCallback , useContext} from 'react'
import { LanguageListResponse } from '../../../constants/common'
import { UPDATE_LANGUAGES, GET_LANGUAGES } from '../../../constants/endpoints'

import style from '../Modal.module.css'
import { ModalContextStateContext } from '../../../context/modalContext/modalContext'
import LoadingSpinner from '../../LoadingSpinner/LoadingSpinner'

const LanguagesSelectModal = () => {
  const modalContext = useContext(ModalContextStateContext)
  const [isLoading, setIsLoading] = useState(false)

  const initialArray: string[] = []
  const [languageList, setlanguageList] = useState();
  const [addedlanguages, setAddedLanguages] = useState(initialArray);
  
  const handleAddClick = useCallback((id: string, languageType: string) => {
    const el = document.getElementById(id)
    if (el?.classList.contains(style.active)) {
      el?.classList.remove(style.active)
      addedlanguages.splice(addedlanguages.indexOf(languageType), 1)
    } else {
      el?.classList.add(style.active)
      addedlanguages.push(languageType)
    }
    setAddedLanguages(addedlanguages)
  },[addedlanguages])


  useEffect(() => {
    setIsLoading(true)
    fetch(process.env.REACT_APP_API_URI + GET_LANGUAGES, {
      method: 'GET',
    })
      .then((response) => {
        if (response.status !== 200) {
          console.log('fail')
        }
        return response.json()
      })
      .then((data: LanguageListResponse) => {
        let languages: any = []
        data.languages.forEach((element, index) => {
          languages.push(
            <button
              id={index + ''}
              className={[style.modalText, style.skillButton].join(' ')}
              onClick={() => handleAddClick(index + '', element)}
              key={index}
            >
              {element}
            </button>
          )
          setlanguageList(languages)
          setIsLoading(false)
        })
      })
  },[handleAddClick])


  const submitLanguages = () => {
    setIsLoading(true)
    const responseBody = { languages: addedlanguages }
    fetch(process.env.REACT_APP_API_URI + UPDATE_LANGUAGES, {
      method: 'PATCH',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('gitcollab_jwt'),
      },
      body: JSON.stringify(responseBody),
    })
      .then((response) => {
        if (response.status === 200) {
          modalContext.hideModal()
          setIsLoading(false)
        } else {
          console.log('Failed')
        }
      })
      .then((data: any) => {
        console.log(data)
      })
  }

  return (
    <>
    {isLoading && <LoadingSpinner isLoading={isLoading} />}
      <div className={style.modalText}>
        <p className={style.modalTextTitle}>Tell us more about yourself</p>
        <div className={style.modalTextUnderline}/>
        <p className={style.modalTextContent}>
          Select a few programming lanaguages that interest you
        </p>
        <div className={style.skillButtonContainer}>
          <>{languageList}</>
        </div>

        <button className={[style.modalButton, style.skillContinueButton].join(" ")} onClick={() => submitLanguages()}>
          Continue
        </button>
      </div>
    </>
  )
}

export default LanguagesSelectModal
