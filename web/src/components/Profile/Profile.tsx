import React, { useEffect, useState } from 'react'
import Media from 'react-media'
import styles from './Profile.module.css'
import { GET_PROFILE } from '../../constants/endpoints'
import { profileResponse } from '../../constants/common'

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
    console.log(`Testing with username ${username}`)
    fetch(process.env.REACT_APP_API_URI + GET_PROFILE + username, {
      method: 'GET',
    })
      .then((response) => response.json())
      .then((data: profileResponse) => {
        setProfile(data)
        console.log(data)
      })
  }, [])

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
      <Media query={{ maxWidth: 1023 }}>
        <div className={styles.tables}>
          <div className={styles.row}>
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
          </div>
          <div className={styles.row}>
            <div className={styles.card}>
              <div>Card 2</div>
              <div className={styles.line}></div>
              <div>Card 2 Body</div>
            </div>
          </div>
        </div>
      </Media>
      <Media query={{ minWidth: 1024 }}>
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
