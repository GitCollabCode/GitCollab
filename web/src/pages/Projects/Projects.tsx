import React, { useContext, useEffect, useState } from 'react'
import Button from '../../components/Button/Button'
import ContributerCard from '../../components/ContributerCard/ContributerCard'
import TaskCard from '../../components/TaskCard/TaskCard'
import Table from '../../components/Table/Tables'
import {
  ContributerType,
  ModalType,
  TaskResponse,
} from '../../constants/common'
import style from './Projects.module.css'
import { ModalContextStateContext } from '../../context/modalContext/modalContext'
import { GET_PROJECT, GET_TASKS } from '../../constants/endpoints'

type ProjectResp = {
  project_id:number,
  project_owner_id:string,
  project_owner_username:string,
  project_name:string,
  project_url:string,
  project_skills:string[],
  project_description:string,
}

const initialproject:ProjectResp ={
  project_id:0,
  project_owner_id:"string",
  project_owner_username:"string",
  project_name:"string",
  project_url:"string",
  project_skills:[],
  project_description:"string",
}

const Project = () => {
  const { showModal, setModalType } = useContext(ModalContextStateContext)
  const projectName = window.location.href.split('project/')[1]
  const [project, setProject] = useState<ProjectResp>(initialproject)
  const [tasks, setTasks] = useState<TaskResponse>()
  
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

  const getIssuesList = (tasks: TaskResponse) => {
    //eslint-disable-next-line
    let taskCards: JSX.Element[] = []
    tasks[0].forEach((element) => {
      taskCards.push(<TaskCard task={element} />)
    })

    return taskCards
  }


useEffect(() => {
  //setIsLoading(true)
  console.log(projectName)
  fetch(process.env.REACT_APP_API_URI + GET_PROJECT + projectName +"/", {
    method: 'GET',
  })
    .then((response) => {
      if (response.status >= 400) {
        throw new Error('API failed')
      } else {
        return response.json()
      }
    })
    .then((data: ProjectResp) => {
      console.log(data)
      setProject(data)
      //setIsLoading(false)
    })
    .catch((err) => {
     // setError(true)
    })

    fetch(process.env.REACT_APP_API_URI + GET_PROJECT + projectName + GET_TASKS, {
      method: 'GET',
    })
      .then((response) => {
        if (response.status >= 400) {
          throw new Error('API failed')
        } else {
          return response.json()
        }
      })
      .then((data: TaskResponse) => {
        console.log(data)
        setTasks(data)
        //setIsLoading(false)
      })
      .catch((err) => {
       // setError(true)
      })
}, [projectName])


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
  

  return (
    <div className={style.page}>
      <div className={style.left}>
        <div className={[style.card, style.descriptionCard].join(' ')}>
          <p>
            <u>Description</u>
          </p>
          <p className={style.description}>
            {project.project_description}
          </p>
        </div>
        <div className={[style.card, style.languagesCard].join(' ')}>
          <p>Project Langs</p>
          <div className={style.over}>{getPills(project.project_skills)}</div>
        </div>
        <div className={[style.card, style.skillsCard].join(' ')}>
          <p>Project Skills</p>
          <div className={style.over}>{getPills(project.project_skills)}</div>
        </div>
      </div>

      <div className={style.right}>
        <div className={style.projectInfo}>
          <p>{project.project_name}</p>
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
            <Table>{tasks && getIssuesList(tasks)}</Table>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Project
