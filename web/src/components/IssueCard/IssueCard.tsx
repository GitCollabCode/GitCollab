import React from 'react'
import { IssueCardType } from '../../constants/common'
import style from './IssueCard.module.css'

const IssueCard = ({ issue }: { issue: IssueCardType }) => {
  console.log(issue)
  return (
    <>
      <p className={style.style}>{issue.name}</p>
      <p>{issue.description}</p>
      <p>{issue.assignedTo}</p>
      <p>{issue.assignedToImg}</p>
      <p>{issue.languages}</p>
      <p>{issue.progress}</p>
    </>
  )
}

export default IssueCard
