import React from 'react'
import style from './NavItem.module.css'

const NavItem = ({ text, link }: { text: string; link: string }) => {
  return (
    <a className={style.navText} href={link}>
      {text}
    </a>
  )
}

export default NavItem
