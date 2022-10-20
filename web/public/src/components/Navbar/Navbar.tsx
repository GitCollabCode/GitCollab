import React from 'react'
import NavItem from '../NavItem/NavItem'

import style from '../Navbar/Navbar.module.css'
import Login from '../Login/Login'

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
          <Login />
        </div>
      </nav>
    </div>
  )
}

export default Navbar
