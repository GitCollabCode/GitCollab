import React from 'react'
import styles from './Button.module.css'
import { ReactComponent as PlusIcon } from '../../assets/circle-plus-solid.svg'
import { ReactComponent as GithubIcon } from '../../assets/github.svg'
import { ReactComponent as LogoutIcon } from '../../assets/logout.svg'

const Button = ({
  type,
  text,
}: {
  type: 'login' | 'new' | 'logout'
  text: string
}) => {
  let icon
  switch (type) {
    case 'login':
      icon = <GithubIcon />
      break
    case 'new':
      icon = <PlusIcon />
      break
    case 'logout':
      icon = <LogoutIcon />
      break
  }

  return (
    <button className={styles.greenButton}>
      {icon}
      <p className={styles.buttonText}>{text}</p>
    </button>
  )
}

export default Button
