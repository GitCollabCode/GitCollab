import React, { useEffect, useState } from 'react'
import Media from 'react-media'
import styles from './Profile.module.css'
import { GET_PROFILE, USER_PROJECT } from '../../constants/endpoints'
import { ProfileProjectResponse, profileResponse } from '../../constants/common'
import Table from '../../components/Table/Tables'
import ProjectCard from '../../components/ProjectCard/ProjectCard'

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
    console.log(window.location.href)
    console.log(`Testing with username ${username}`)
    console.log(process.env.REACT_APP_API_URI + GET_PROFILE + username)
    fetch(process.env.REACT_APP_API_URI + GET_PROFILE + username, {
      method: 'GET',
    })
      .then((response) => response.json())
      .then((data: profileResponse) => {
        setProfile(data)
        console.log(data)
      })
  }, [])

  const getLanguages = () => {
    //eslint-disable-next-line
    let languagList: JSX.Element[] = [<></>]
    profile.languages.forEach((element) => {
      languagList.push(<li className={styles.profileLi}>{element}</li>)
    })

    return languagList
  }

  const getSkills = () => {
    //eslint-disable-next-line
    let skills: JSX.Element[] = [<></>]
    profile.skills.forEach((element) => {
      skills.push(<li className={styles.profileLi}>{element}</li>)
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
      .then((response) => response.json())
      .then((data:ProfileProjectResponse ) => {
        console.log(data)
        if(data.projects !==null){
          setProjectsCards(data.projects)
        }
          setIsLoading(false)
          
      })
  }, [])


  const setProjectsCards = (projects:string[]) => {
    //eslint-disable-next-line
    let cards: JSX.Element[] = []
    projects.forEach((element, index) => {
      cards.push(
        <tr key={index}>
          <td>
            <ProjectCard data={{project_name:element, project_description:"",project_owner:"", project_skills:[]}} key={index} />
          </td>
        </tr>
      )
    })
    console.log(cards)
    setProjectCard(cards)
  }
  console.log(isLoading)
  return (
    <div className={styles.container}>
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
            <div className={styles.card}>
              <div className={styles.header}>Skills</div>
              <div className={styles.line}></div>
              <ul>{getSkills()}</ul>
            </div>
            <div className={styles.card}>
              <div className={styles.header}>Languages</div>
              <div className={styles.line}></div>
              <ul>{getLanguages()}</ul>
            </div>
          </div>
          <div className={styles.row}>
            <div className={styles.card}>
              <div className={styles.header}>Projects</div>
              <div className={styles.line}></div>
              <div>
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
