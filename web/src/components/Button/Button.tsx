import React from 'react'
import styles from './Button.module.css'
import { ReactComponent as PlusIcon } from '../../assets/circle-plus-solid.svg'
import { ReactComponent as GithubIcon } from '../../assets/github.svg'
import { ReactComponent as LogoutIcon } from '../../assets/logout.svg'
import { ReactComponent as NextIcon } from '../../assets/next-arrow.svg'
const Button = ({
  type,
  text,
  onClick,
}: {
  type: 'login' | 'new' | 'logout' | 'next'
  text: string
  onClick?: () => void
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
    case 'next':
      icon = <NextIcon />
      break
  }

  return (
    <button
      className={styles.greenButton}
      title={type}
      onClick={() => onClick && onClick()}
    >
      {icon}
      <p className={styles.buttonText}>{text}</p>
    </button>
  )
}

export default Button
