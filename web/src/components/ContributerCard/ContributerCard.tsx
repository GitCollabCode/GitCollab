import React from 'react'
import { ContributerType } from '../../constants/common'
import style from './ContributerCard.module.css'

const ContributerCard = ({ contributer }: { contributer: ContributerType }) => {
  return (
    <div className={style.card}>
      <div className={style.content}>
        <p>{contributer.name}</p>
        <img
          className={style.profilePic}
          src={contributer.url}
          alt={'Profile'}
        />
      </div>
    </div>
  )
}

export default ContributerCard
