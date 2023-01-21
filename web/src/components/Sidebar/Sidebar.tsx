import React from 'react'
import { bubble as Menu } from 'react-burger-menu'

import style from './Sidebar.module.css'

const Sidebar = () => {
  return (
    <Menu>
      <div className={style.sidebar}>
        <div>
          <a href="/projects" className={style.link}>
            projects
          </a>
        </div>

        <div>
          <a href="/" className={style.link}>
            learn
          </a>
        </div>
        <div>
          <a href="/" className={style.link}>
            about
          </a>
        </div>
      </div>
    </Menu>
  )
}

export default Sidebar
