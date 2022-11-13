import React from 'react'
import Media from 'react-media'
import styles from './Profile.module.css'
const Profile = () => {
  return (
    <div className={styles.container}>
      <div className={styles.bio}>
        <p>User profile</p>
        <div className={styles.circle}>
        </div>
        <strong>user Name</strong>
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
      <Media query={{ minWidth: 1024  }}>
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
