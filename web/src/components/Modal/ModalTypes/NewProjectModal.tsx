import React, { useCallback, useContext, useEffect, useState } from 'react'
import {
  ReposResponse,
  SelectType,
  SkillListResponse,
} from '../../../constants/common'
import {
  CREATE_PROJECT,
  GET_SKILLS,
  GET_USER_REPOS,
} from '../../../constants/endpoints'
import styles from '../Modal.module.css'

import octocat from '../../../assets/octocat.png'

import Select from 'react-select'
import { ModalContextStateContext } from '../../../context/modalContext/modalContext'
import LoadingSpinner from '../../LoadingSpinner/LoadingSpinner'

const NewProjectModal = () => {
  const { hideModal } = useContext(ModalContextStateContext)
  const intialRepos: ReposResponse = {
    repos: [''],
  }

  const [repos, setRepos] = useState(intialRepos)
  const [step, setCurrentStep] = useState(0)
  const [selectedRepo, setSelectedRepo] = useState('')
  const [description, setDescription] = useState('')

  console.log(selectedRepo, description)

  const initialArray: string[] = []
  const [skillList, setSkillList] = useState()
  const [addedSkills, setAddedSkills] = useState(initialArray)
  const [isLoading, setIsLoading] = useState(false)

  const handleAddClick = useCallback(
    (id: string, skillType: string) => {
      const el = document.getElementById(id)
      if (el?.classList.contains(styles.active)) {
        el?.classList.remove(styles.active)
        addedSkills.splice(addedSkills.indexOf(skillType), 1)
      } else {
        el?.classList.add(styles.active)
        addedSkills.push(skillType)
      }
      setAddedSkills(addedSkills)
    },
    [addedSkills]
  )

  //Fetch a users public repos
  useEffect(() => {
    setIsLoading(true)
    fetch(process.env.REACT_APP_API_URI + GET_USER_REPOS, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('gitcollab_jwt'),
      },
    })
      .then((response) => response.json())
      .then((data: ReposResponse) => {
        setRepos(data)
        setIsLoading(false)
      })
  }, [])

  //UseEffect to query each piece of data for each step in the form
  useEffect(() => {
    if (step === 1) {
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
                className={[styles.modalText, styles.skillButton].join(' ')}
                onClick={() => handleAddClick(index + '', element)}
                key={index}
              >
                {element}
              </button>
            )
            setSkillList(skills)
          })
        })
    }
  }, [handleAddClick, step])

  //Handle React Select change
  const handleRepoChange = (value: string | undefined) => {
    if (value) {
      setSelectedRepo(value)
    }
  }

  const createProject = () => {
    const requestData = {
      repo_name: selectedRepo,
      skills: addedSkills,
      description: description,
    }

    console.log('posting' + requestData)
    fetch(process.env.REACT_APP_API_URI + CREATE_PROJECT, {
      method: 'POST',
      body: JSON.stringify(requestData),
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('gitcollab_jwt'),
      },
    }).then(() => {
      hideModal()
    })
  }

  //Formats the Repos into usable data
  const getReposSelect = () => {
    let values: SelectType[] = []
    repos.repos.forEach((element) => {
      values.push({ value: element, label: element })
    })
    return values
  }

  //Gets the current step for the submitting a New Project
  //eslint-disable-next-line
  const getCurrentStepsHtml = (step: number): JSX.Element => {
    switch (step) {
      case 2:
        return (
          <>
            <div className={styles.modalText}>
              <p className={styles.modalTextTitle}>
                Tell us more about your project
              </p>
              <div className={styles.modalTextUnderline} />
              <p className={styles.modalTextContent}>
                Please provide a description of your project
              </p>

              <textarea
                className={styles.textArea}
                rows={10}
                onChange={(e) => setDescription(e.target.value)}
              ></textarea>
              <div className={styles.spaceBox}></div>
              <button
                className={[
                  styles.modalButton,
                  styles.skillContinueButton,
                ].join(' ')}
                onClick={() => createProject()}
              >
                Create Project
              </button>
            </div>
          </>
        )

      case 1:
        return (
          <>
            <div className={styles.modalText}>
              <p className={styles.modalTextTitle}>
                Tell us more about your project
              </p>
              <div className={styles.modalTextUnderline} />
              <p className={styles.modalTextContent}>
                Select a few topics that you want in your project
              </p>
              <div className={styles.skillButtonContainer}>
                <>{skillList}</>
              </div>
              <button
                className={[
                  styles.modalButton,
                  styles.skillContinueButton,
                ].join(' ')}
                onClick={() => setCurrentStep(2)}
              >
                Continue
              </button>
            </div>
          </>
        )

      case 0:
        return (
          <>
            {isLoading ? (
              <LoadingSpinner isLoading={isLoading} type="fixed" />
            ) : (
              <div className={styles.newProjectContainer}>
                <img
                  className={styles.modalLogo}
                  src={octocat}
                  alt="Github Cat"
                />
                <p className={styles.modalTextContent}>
                  {repos.repos.length > 0
                    ? 'Please select the project you would like to register from Github'
                    : 'To create a GitCollab Project, please create a GitHub project first'}
                </p>
                {repos.repos.length !== 0 ? (
                  <div className={styles.selectBox}>
                    <Select
                      options={getReposSelect()}
                      onChange={(e) => handleRepoChange(e?.value)}
                    />
                  </div>
                ) : (
                  <></>
                )}
                <div className={styles.spaceBox}></div>
                <button
                  className={[
                    styles.modalButton,
                    styles.skillContinueButton,
                  ].join(' ')}
                  onClick={() => {
                    repos.repos.length > 0 ? setCurrentStep(1) : hideModal()
                  }}
                >
                  {repos.repos.length > 0 ? 'Continue' : 'Close'}
                </button>
              </div>
            )}
          </>
        )
      default:
        return <>error</>
    }
  }

  return getCurrentStepsHtml(step)
}

export default NewProjectModal
