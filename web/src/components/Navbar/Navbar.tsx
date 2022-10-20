import React from 'react'
import NavItem from '../NavItem/NavItem'

import style from '../Navbar/Navbar.module.css'
import LoginButton from '../LoginButton/LoginButton'

const Navbar = () => {
  return (
    <div className={style.navbar}>
      <nav>
        <div className={style.logo}>GitCollab</div>
        <div className={style.centerNavItems}>
          <NavItem text="discover" link="/#" />
          <NavItem text="learn" link="/#" />
          <NavItem text="about" link="/#" />
        </div>
        <div className={style.loginBox}>
          <LoginButton />
        </div>
      </nav>
    </div>
  )
}

export default Navbar
