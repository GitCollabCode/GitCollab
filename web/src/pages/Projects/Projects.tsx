import React, { useContext } from 'react'
import Button from '../../components/Button/Button'
import ContributerCard from '../../components/ContributerCard/ContributerCard'
import TaskCard from '../../components/TaskCard/TaskCard'
import Table from '../../components/Table/Tables'
import {
  TaskCardType,
  TaskProgress,
  ContributerType,
  ModalType,
} from '../../constants/common'
import style from './Projects.module.css'
import { ModalContextStateContext } from '../../context/modalContext/modalContext'

const Project = () => {
  const { showModal, setModalType } = useContext(ModalContextStateContext)
  const getPills = (data: string[]) => {
    //eslint-disable-next-line
    let pills: JSX.Element[] = []
    data.forEach((element) => {
      pills.push(<div className={style.pill}>{element}</div>)
    })

    return pills
  }

  const getContributers = (contributer: ContributerType[]) => {
    //eslint-disable-next-line
    let contributerCards: JSX.Element[] = []
    contributer.forEach((element) => {
      contributerCards.push(<ContributerCard contributer={element} />)
    })

    return contributerCards
  }

  const getIssuesList = (tasks: TaskCardType[]) => {
    //eslint-disable-next-line
    let taskCards: JSX.Element[] = []
    tasks.forEach((element) => {
      taskCards.push(<TaskCard task={element} />)
    })

    return taskCards
  }

  const langs = ['lang1', 'lang2', 'lang3']
  const skills = ['skills1', 'skills2', 'skills3']
  const contributers: ContributerType[] = [
    {
      name: 'kevin',
      url: 'https://avatars.githubusercontent.com/u/39808977?s=40&v=4',
    },
    {
      name: 'ehvan',
      url: 'https://avatars.githubusercontent.com/u/39808977?s=40&v=4',
    },
    {
      name: 'kevin',
      url: 'https://avatars.githubusercontent.com/u/39808977?s=40&v=4',
    },
    {
      name: 'ehvan',
      url: 'https://avatars.githubusercontent.com/u/39808977?s=40&v=4',
    },
  ]
  const tasks: TaskCardType[] = [
    {
      name: '',
      description: '',
      languages: ['', ''],
      assignedTo: '',
      assignedToImg: '',
      progress: TaskProgress.UnAssigned,
    },
  ]
  console.log(contributers)
  return (
    <div className={style.page}>
      <div className={style.left}>
        <div className={[style.card, style.descriptionCard].join(' ')}>
          <p>
            <u>Description</u>
          </p>
          <p className={style.description}>
            This project is actually so stupid, Majed is cringe
          </p>
        </div>
        <div className={[style.card, style.languagesCard].join(' ')}>
          <p>Project Langs</p>
          <div className={style.over}>{getPills(langs)}</div>
        </div>
        <div className={[style.card, style.skillsCard].join(' ')}>
          <p>Project Skills</p>
          <div className={style.over}>{getPills(skills)}</div>
        </div>
      </div>

      <div className={style.right}>
        <div className={style.projectInfo}>
          <p>The Project</p>
          <div className={style.contributerList}>
            {getContributers(contributers)}
          </div>
        </div>
        <div className={style.tasksDisplay}>
          <div className={style.tasksBox}>
            <div className={style.tasksTitle}>
              <div className={style.titleBox}>
                <p className={style.title}>Tasks</p>
                <Button
                  type="new"
                  text="New Task"
                  onClick={() => {
                    setModalType(ModalType.NewTaskModal)
                    showModal()
                  }}
                />
              </div>
              <div className={style.line} />
            </div>
            <Table>{getIssuesList(tasks)}</Table>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Project
