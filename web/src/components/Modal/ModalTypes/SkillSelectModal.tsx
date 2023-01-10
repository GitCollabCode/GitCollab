import React, { useEffect, useState } from 'react'
import { SkillListResponse } from '../../../constants/common'
import { UPDATE_SKILLS, GET_SKILLS } from '../../../constants/endpoints'
import { ModalContextStateContext } from '../../../context/modalContext/modalContext'

import style from '../Modal.module.css'

//import { ModalContextStateContext } from '../../../context/modalContext/modalContext'

const LoginModal = () => {
  //const modalContext = useContext(ModalContextStateContext)
  //const [data, setData] = useState({ errorMessage: '', isLoading: false })

  const initialArray: string[] = []
  const [skillList, setSkillList] = useState();
  const [addedSkills, setAddedSkills] = useState(initialArray);
  
  
  useEffect(() => {
    fetch(process.env.REACT_APP_API_URI + GET_SKILLS, {
      method: 'GET',
    })
      .then((response) => {
        if (response.status !== 200) {
          console.log('fail')
        }
        return response.json()
      })
      .then((data: SkillListResponse) => {
        let skills: any = []
        data.skills.forEach((element, index) => {
          skills.push(
            <button
              id={index + ''}
              className={[style.modalText, style.skillButton].join(' ')}
              onClick={() => handleAddClick(index + '', element)}
              key={index}
            >
              {element}
            </button>
          )
          setSkillList(skills)
        })
      })
  }, [])

  

  const handleAddClick = (id: string, skillType: string) => {
    console.log(addedSkills)
    const el = document.getElementById(id)
    if (el?.classList.contains(style.active)) {
      el?.classList.remove(style.active)
      addedSkills.splice(addedSkills.indexOf(skillType), 1)
    } else {
      el?.classList.add(style.active)
      addedSkills.push(skillType)
    }
    setAddedSkills(addedSkills)
  }

  const submitSkills = () => {
    const responseBody = { skills: addedSkills }
    console.log(responseBody)
    fetch(process.env.REACT_APP_API_URI + UPDATE_SKILLS, {
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
          console.log('It worked')
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
      <div className={style.modalText}>
        <p className={style.modalTextTitle}>Tell us more about yourself</p>
        <div className={style.modalTextUnderline}/>
        <p className={style.modalTextContent}>
          Select a few topics that interest you
        </p>
        <div className={style.skillButtonContainer}>
          <>{skillList}</>
        </div>

        <button className={[style.modalButton, style.skillContinueButton].join(" ")} onClick={() => submitSkills()}>
          Continue
        </button>
      </div>
    </>
  )
}

export default LoginModal
