import React, { useEffect, useState } from 'react'
import Media from 'react-media'
import styles from './Profile.module.css'
import { GET_PROFILE } from '../../constants/endpoints'
import { profileResponse } from '../../constants/common'
import Table, { rowType } from '../../components/Table/Tables'

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

  const intitalDataRows: rowType[] = [
    {
      id: '1',
      date: '2014-04-18',
      total: 121.0,
      status: 'Shipped',
      name: 'A',
      points: 5,
      percent: 50,
    },
    {
      id: '2',
      date: '2014-04-21',
      total: 121.0,
      status: 'Not Shipped',
      name: 'B',
      points: 10,
      percent: 60,
    },
    {
      id: '3',
      date: '2014-08-09',
      total: 121.0,
      status: 'Not Shipped',
      name: 'C',
      points: 15,
      percent: 70,
    },
    {
      id: '4',
      date: '2014-04-24',
      total: 121.0,
      status: 'Shipped',
      name: 'D',
      points: 20,
      percent: 80,
    },
    {
      id: '5',
      date: '2014-04-26',
      total: 121.0,
      status: 'Shipped',
      name: 'E',
      points: 25,
      percent: 90,
    },
  ]

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
                <Table rows={intitalDataRows} isExpandable={false} />
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
