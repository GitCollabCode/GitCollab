import React from 'react'
import { ContributerType } from '../../constants/common'
import style from './ContributerCard.module.css'

const ContributerCard = ({ contributer }: { contributer: ContributerType }) => {
  return (
    <div className={style.card}>
      <div className={style.content}>
        <p>{contributer.name}</p>
      </div>
    </div>
  )
}

export default ContributerCard
