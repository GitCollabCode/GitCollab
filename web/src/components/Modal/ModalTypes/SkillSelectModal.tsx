import React, { useEffect } from 'react'
import { SkillListResponse } from '../../../constants/common'
import { UPDATE_SKILLS, GET_SKILLS } from '../../../constants/endpoints'

import style from '../Modal.module.css'

//import { ModalContextStateContext } from '../../../context/modalContext/modalContext'

const LoginModal = () => {
  //const modalContext = useContext(ModalContextStateContext)
  //const [data, setData] = useState({ errorMessage: '', isLoading: false })

  let idCounter = 0
  let skillList
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
        skillList = []
        data.skills.forEach((element) => {
          skillList.push(
            <button
              id={idCounter + ''}
              className={[style.modalText, style.skillButton].join(' ')}
              onClick={() => handleAddClick(idCounter + '', element)}
            >
              {element}
            </button>
          )
          idCounter++
        })
      })
  })

  let addedSkills: string[] = []

  const handleAddClick = (id: string, skillType: string) => {
    const el = document.getElementById(id)
    if (el?.classList.contains(style.active)) {
      el?.classList.remove(style.active)
      addedSkills.splice(addedSkills.indexOf(skillType), 1)
    } else {
      el?.classList.add(style.active)
      addedSkills.push(skillType)
    }
  }

  const submitSkills = () => {
    const responseBody = { skills: addedSkills }
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
        <p className={style.modalTextContent}>Tell us more about yourself</p>
        <div className={style.modalTextUnderline} />
        <p className={style.modalTextContent}>
          Select a few topics that interest you
        </p>
        <div className={style.skillButtonContainer}>
          <>{skillList}</>
        </div>

        <button className={style.modalButton} onClick={() => submitSkills()}>
          Continue
        </button>
      </div>
    </>
  )
}

export default LoginModal
