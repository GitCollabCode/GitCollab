import React, { useEffect, useState } from 'react'
import styles from './NewProject.module.css'
import { ReposResponse, SelectType } from '../../constants/common'
import { GET_USER_REPOS } from '../../constants/endpoints'
import Select from 'react-select'

const NewProject = () => {
  const intialRepos: ReposResponse = {
    repos: [''],
  }
  // const [step, setStep] = useState(0)
  const [repos, setRepos] = useState(intialRepos)

  useEffect(() => {
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
        console.log(data)
      })
  }, [])

  const getReposSelect = () => {
    let values: SelectType[] = []
    repos.repos.forEach((element) => {
      values.push({ value: element, label: element })
    })
    return values
  }

  //step1 - getRepos

  //step2 - get repo info

  //step 3 - select skills

  //step 4 - select description

  //step 5 - create project

  //step 6 - redirect to new project page

  return (
    <div className={styles.node}>
      <Select options={getReposSelect()} />
    </div>
  )
}

export default NewProject
