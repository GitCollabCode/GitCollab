import React, { useEffect, useState } from 'react'
import Media from 'react-media'
import styles from './Profile.module.css'
import { GET_PROFILE } from '../../constants/endpoints'
import { profileResponse, ProjectCardType } from '../../constants/common'
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

  const data: ProjectCardType[] = [
    {
      name: 'Project Name',
      description: 'Lorem upsum',
      languages: ['Stuff', 'Other', 'Other'],
      url: '#',
    },
  ]
  const getProjectsRows = () => {
    //eslint-disable-next-line
    let projectCards: JSX.Element[] = []
    data.forEach((element) => {
      projectCards.push(<ProjectCard data={element} />)
    })
    return projectCards
  }

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
                <Table>{getProjectsRows()}</Table>
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
