import React, { useContext }  from 'react'
import { UPDATE_SKILLS } from '../../../constants/endpoints'
import { ModalContextStateContext } from '../../../context/modalContext/modalContext'
import style from '../Modal.module.css'

const SkillSelectModal = () => {    
    let addedSkills:string[] = [] 
    const modalContext = useContext(ModalContextStateContext)

    const handleAddClick = (id:string, skillType: string) => {
        const el = document.getElementById(id)
        if(el?.classList.contains(style.active)){
            el?.classList.remove(style.active)
            addedSkills.splice(addedSkills.indexOf(skillType), 1)
        }else{
            el?.classList.add(style.active)
            addedSkills.push(skillType)
        }
    }

    const submitSkills = () => {
        const responseBody = {skills:addedSkills}
        fetch(process.env.REACT_APP_API_URI + UPDATE_SKILLS , {
            method: 'PATCH',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                Authorization: 'Bearer ' + localStorage.getItem('gitcollab_jwt'),
            },
            body:JSON.stringify(responseBody),
            
        }).then((response) => {
            if(response.status === 200){
                console.log("It worked")
                // close modal and redirect to their profile
                // TODO REDIRECT SOMEWHERE ELSE
                modalContext.hideModal()
            } else{
                console.log("Failed")
            }
            })
            .then((data: any) => {
                console.log(data)
            })
        }

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

            
            <button id="1" className={[style.modalText, style.skillButton].join(" ")} onClick={()=>handleAddClick("1", "continue")}>
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
    
        <button className={style.modalButton} onClick={()=>submitSkills()}>
              Continue
          </button>
      </div>
      </>
    ) 
}

export default SkillSelectModal