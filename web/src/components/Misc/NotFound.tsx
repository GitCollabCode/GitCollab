import React from 'react'
import { Link } from 'react-router-dom'
import styles from './NotFound.module.css'
const NotFound = () => {
  return (
    <div className={styles.container}>
      <h1>404 Error</h1>
      <h2>Page not found...</h2>
      <Link to="/">take this link back home</Link>
    </div>
  )
}

export default NotFound
