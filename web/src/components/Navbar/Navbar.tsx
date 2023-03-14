import React, { useState } from 'react'
import NavItem from '../NavItem/NavItem'

import style from '../Navbar/Navbar.module.css'
import LoginButton from '../LoginButton/LoginButton'
import Media from 'react-media'
import Sidebar from '../Sidebar/Sidebar'
import LoadingSpinner from '../LoadingSpinner/LoadingSpinner'

const Navbar = () => {
  const [isLoading, setIsLoading] = useState(false)
  return (
    <>
      <Media query={{ minWidth: 1023 }}>
        <div className={style.navbar}>
          {isLoading && <LoadingSpinner isLoading={isLoading} />}
          <nav>
            <div className={style.logo} data-testid={'logo'}>
              GitCollab
            </div>
            <div className={style.centerNavItems}>
              <NavItem text="projects" link="/projects" />
              <NavItem text="learn" link="/#" />
              <NavItem text="about" link="/#" />
            </div>
            <div className={style.loginBox}>
              <LoginButton setIsLoading={setIsLoading} />
            </div>
          </nav>
        </div>
      </Media>
      <Media query={{ maxWidth: 1024 }}>
        <div className={style.navbarMobile}>
          <Sidebar />
        </div>
      </Media>
    </>
  )
}

export default Navbar
