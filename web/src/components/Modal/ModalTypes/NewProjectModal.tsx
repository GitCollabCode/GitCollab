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
    repos: [],
  }
  const initialArray: string[] = []

  const [step, setCurrentStep] = useState(0)

  const [repos, setRepos] = useState(intialRepos) //List of repos from user-repos
  const [selectedRepo, setSelectedRepo] = useState('') //The selected repo
  const [description, setDescription] = useState('') //The description of the project
  const [error, setError] = useState(false) //For when an API failed

  const [skillList, setSkillList] = useState() //The total list of skills
  const [addedSkills, setAddedSkills] = useState(initialArray) //Skills that the user selected
  const [isLoading, setIsLoading] = useState(false) //For when api is loading

  //This is for when adding a skill to the array when clicked
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
      .then((response) => {
        if (response.status >= 400) {
          throw new Error('API failed')
        } else {
          return response.json()
        }
      })
      .then((data: ReposResponse) => {
        setRepos(data)
        setIsLoading(false)
      })
      .catch((err) => {
        setError(true)
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

  //Handle Selecting a repo when using react select repos
  const handleRepoChange = (value: string | undefined) => {
    if (value) {
      setSelectedRepo(value)
    }
  }

  //Function to create a new project
  const createProject = () => {
    const requestData = {
      repo_name: selectedRepo,
      skills: addedSkills,
      description: description,
    }

    fetch(process.env.REACT_APP_API_URI + CREATE_PROJECT, {
      method: 'POST',
      body: JSON.stringify(requestData),
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('gitcollab_jwt'),
      },
    })
      .then((response) => {
        if (response.status >= 400) {
          throw new Error()
        }
        hideModal()
        window.location.reload()
      })
      .catch((err) => {
        setError(true)
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

  //Renders the correct html for which step the user is on
  //eslint-disable-next-line
  const getCurrentStepsHtml = (step: number): JSX.Element => {
    if (error) {
      return (
        <>
          <div className={styles.modalText}>
            <p className={styles.modalTextTitle}>
              There was an error in the request
            </p>
            <div className={styles.modalTextUnderline} />
            <p className={styles.modalTextContent}>Please Try Again Later</p>

            <div className={styles.spaceBox}></div>
            <button
              className={[styles.modalButton, styles.skillContinueButton].join(
                ' '
              )}
              onClick={() => hideModal()}
            >
              Close
            </button>
          </div>
        </>
      )
    }
    switch (step) {
      case 0: // Selecting a repo
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
                      onChange={(e: any) => handleRepoChange(e?.value)}
                    />
                  </div>
                ) : (
                  <></>
                )}
                <div className={styles.spaceBox}></div>
                <button
                  disabled={selectedRepo === '' ? true : false}
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
      case 1: //Adding skills list to project
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
      case 2: // Adding a description to a project
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
      default:
        return <></>
    }
  }

  return getCurrentStepsHtml(step)
}

export default NewProjectModal
