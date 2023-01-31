import React from 'react'
import { ProjectCardType } from '../../constants/common'

import style from './ProjectCard.module.css'

const ProjectCard = ({ data }: { data: ProjectCardType }) => {
  const getLanguagePills = () => {
    //eslint-disable-next-line
    let languagePills: JSX.Element[] = []
    data.languages.forEach((element) => {
      languagePills.push(<div className={style.pill}>{element}</div>)
    })

    return languagePills
  }

  return (
    <a href={data.url}>
      <div className={style.card}>
        <div className={style.info}>
          <h3>{data.name}</h3>

          <p>{data.description}</p>
        </div>
        <div className={style.tags}>{getLanguagePills()}</div>
      </div>
    </a>
  )
}

export default ProjectCard
