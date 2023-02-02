import React, { useContext, useEffect, useState } from 'react'
import Media from 'react-media'
import styles from './Profile.module.css'
import { GET_PROFILE, USER_PROJECT } from '../../constants/endpoints'
import { ModalType, ProfileProjectResponse, ProjectCardType, profileResponse } from '../../constants/common'
import Table from '../../components/Table/Tables'
import ProjectCard from '../../components/ProjectCard/ProjectCard'
import { ModalContextStateContext } from '../../context/modalContext/modalContext'
import LoadingSpinner from '../../components/LoadingSpinner/LoadingSpinner'

const Profile = () => {
  let initialProfile: profileResponse = {
    username: '',
    gitID: -1,
    email: '',
    avatarUrl: '',
    bio: '',
    languages: [''],
    skills: [''],
  }

  const [profile, setProfile] = useState(initialProfile)
  const [isLoading, setIsLoading] = useState(false)
  //eslint-disable-next-line
  const [projectCards, setProjectCard] = useState<JSX.Element[] | undefined>(undefined)
  
  useEffect(() => {
    const username = window.location.href.split('profile/')[1]
    fetch(process.env.REACT_APP_API_URI + GET_PROFILE + username, {
      method: 'GET',
    })
    .then((response) => {
      if(response.status >=400)
      {
          throw new Error()
      }else{
        return response.json()
    }})
      .then((data: profileResponse) => {
        setProfile(data)
      })
      .catch((err) =>{
       const modal = useContext(ModalContextStateContext)
       modal.setModalType(ModalType.PageNotFoundModal)
       modal.showModal()
      })
  }, [])

  const getLanguages = () => {
    //eslint-disable-next-line
    let languagList: JSX.Element[] = []
    profile.languages.forEach((element, index) => {
      languagList.push(<li key={index} className={styles.profileLi}>{element}</li>)
    })

    return languagList
  }

  const getSkills = () => {
    //eslint-disable-next-line
    let skills: JSX.Element[] = []
    profile.skills.forEach((element, index) => {
      skills.push(<li key={index} className={styles.profileLi}>{element}</li>)
    })

    return skills
  }


  useEffect(() => {
    setIsLoading(true)
    fetch(process.env.REACT_APP_API_URI + USER_PROJECT, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('gitcollab_jwt'),
      },
      body: JSON.stringify({username:window.location.href.split('profile/')[1]}),
      })
      .then((response) => {
        if(response.status >=400)
        {
            throw new Error()
        }else{
          return response.json()
      }})
      .then((data:ProfileProjectResponse) => {
        console.log(data)
        if(data.projects !==null){
          setProjectsCards(data.projects)
        }
          setIsLoading(false)
      })
      .catch((err) =>{
        const modal = useContext(ModalContextStateContext)
        modal.setModalType(ModalType.PageNotFoundModal)
        modal.showModal()
       })
  }, [])


  const setProjectsCards = (projects:ProjectCardType[]) => {
    //eslint-disable-next-line
    let cards: JSX.Element[] = []
    projects.forEach((element, index) => {
      cards.push(
        <tr key={index}>
          <td>
            <ProjectCard data={element} key={index} />
          </td>
        </tr>
      )
    })
    console.log(cards)
    setProjectCard(cards)
  }

  return (
    <div className={styles.container}>
      {isLoading && (<LoadingSpinner isLoading={isLoading}/>)}
      <div className={styles.bio}>
        <div className={styles.circle}>
          <img
            className={styles.image}
            alt="profile"
            src={profile.avatarUrl}
          ></img>
        </div>
        <p className={styles.username}>{profile.username}</p>
        <p className={styles.bioText}>{profile.bio}</p>
      </div>
      <Media query={{ minWidth: 1024 }}>
        <div className={styles.tables}>
          <div className={styles.row}>
            <div className={[styles.card, styles.topCard].join(" ")}>
              <div className={styles.header}>Skills</div>
              <div className={styles.line}></div>
              <ul className={styles.overflow}>{getSkills()}</ul>
            </div>
            <div className={[styles.card, styles.topCard].join(" ")}>
              <div className={styles.header}>Languages</div>
              <div className={styles.line}></div>
              <ul className={styles.overflow}>{getLanguages()}</ul>
            </div>
          </div>
          <div className={styles.row}>
            <div className={[styles.card, styles.projectCard].join(" ")}>
              <div className={styles.header}>Projects</div>
              <div className={styles.line}></div>
              <div className={styles.overflow}>
                <Table>{projectCards}</Table>
              </div>
            </div>
          </div>
        </div>
      </Media>
      <Media query={{ maxWidth: 1023 }}>
        <div className={styles.tables}>
          <div className={styles.card}>
            <div>Card 1</div>
            <div className={styles.line}></div>
            <div>Card 1 Body</div>
          </div>
          <div className={styles.card}>
            <div>Card 2</div>
            <div className={styles.line}></div>
            <div>Card 2 Body</div>
          </div>
          <div className={styles.card}>
            <div>Card 2</div>
            <div className={styles.line}></div>
            <div>Card 2 Body</div>
          </div>
        </div>
      </Media>
    </div>
  )
}

export default Profile
