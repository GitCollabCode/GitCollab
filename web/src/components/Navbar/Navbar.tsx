import React from 'react'
import NavItem from '../NavItem/NavItem'

import style from '../Navbar/Navbar.module.css'
import LoginButton from '../LoginButton/LoginButton'
import Media from 'react-media'
import Sidebar from '../Sidebar/Sidebar'

const Navbar = () => {
  return (
    <>
      <Media query={{ minWidth: 1023 }}>
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
      </Media>
      <Media query={{ maxWidth: 1024 }}>
        <div className={style.navbar}>
          <Sidebar />
        </div>
      </Media>
    </>
  )
}

export default Navbar
