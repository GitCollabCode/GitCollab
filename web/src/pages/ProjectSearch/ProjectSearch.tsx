import React, { useContext, useEffect, useState } from 'react'
import CheckboxFilter from '../../components/CheckboxFilter/CheckboxFilter'
import SearchBox from '../../components/SearchBox/SearchBox'
import Table from '../../components/Table/Tables'
import styles from './ProjectSearch.module.css'

import Button from '../../components/Button/Button'
import { ModalContextStateContext } from '../../context/modalContext/modalContext'
import { ModalType, ProjectCardType, SearchProjectResponse } from '../../constants/common'
import { UserLoginContext } from '../../context/userLoginContext/userLoginContext'
import ProjectCard from '../../components/ProjectCard/ProjectCard'
import { SEARCH_PROJECT } from '../../constants/endpoints'



const ProjectSearch = () => {
  const { showModal, setModalType } = useContext(ModalContextStateContext)
  const { isLoggedIn } = useContext(UserLoginContext)
  //eslint-disable-next-line
  const [projectCards, setProjectCard] = useState<JSX.Element[] | undefined>(undefined)
  const [topicsFilter, setTopicsFilter] = useState([''])
  const [skillsFilter, setSkillsFilter] = useState([''])
  const [languagesFilter, setLanguagesFilter] = useState([''])
  const [searchFilter, setSearchFilter] = useState('')

  const setSelectedFilterOption = (value: string, filterName: string) => {
    switch (filterName) {
      case 'topics':
        topicsFilter.includes(value)
          ? setTopicsFilter(topicsFilter.splice(topicsFilter.indexOf(value), 1))
          : setTopicsFilter(topicsFilter.concat([value]))
        break
      case 'skills':
        skillsFilter.includes(value)
          ? setSkillsFilter(skillsFilter.splice(skillsFilter.indexOf(value), 1))
          : setSkillsFilter(skillsFilter.concat([value]))
        break

      case 'languages':
        languagesFilter.includes(value)
          ? setLanguagesFilter(
              languagesFilter.splice(languagesFilter.indexOf(value), 1)
            )
          : setLanguagesFilter(languagesFilter.concat([value]))
        break
      case 'search':
        setSearchFilter(value)
        console.log(searchFilter)
        break
    }
  }

  useEffect(() => {
    fetch(process.env.REACT_APP_API_URI + SEARCH_PROJECT, {
      method: 'GET',
    })
      .then((response) => {console.log(response)
        return response.json()})
      .then((data: SearchProjectResponse) => {
        setProjectsCards(data.projects)
        console.log(data)
      })
  }, [])

  const setProjectsCards = (projects:ProjectCardType[]) => {
    if(projects === undefined){
      return 
    }
    //eslint-disable-next-line
    let cards: JSX.Element[] = []
    projects.forEach((element, index) => {
      cards.push(
        <tr key={index}>
          <td>
            <ProjectCard data={element} key={index} />
          </td>
        </tr>
      )
    })
    console.log(cards)
    setProjectCard(cards)
  }


  const skills = ['skill1', 'skill2', 'skills3', 'skills4', 'skill5', 'skills6']
  const languages = ['lang1', 'lang2']
  const topics = ['topic1', 'topic2']

  return (
    <div className={styles.projects}>
      <div className={styles.filterPanel}>
        <SearchBox setFilterOption={setSelectedFilterOption} />
        <CheckboxFilter
          filterName="topics"
          setFilterOption={setSelectedFilterOption}
          values={topics}
        />
        <CheckboxFilter
          filterName="skills"
          setFilterOption={setSelectedFilterOption}
          values={skills}
        />
        <CheckboxFilter
          filterName="languages"
          setFilterOption={setSelectedFilterOption}
          values={languages}
        />
      </div>
      <div className={styles.projectsDisplay}>
        <div className={styles.projectsBox}>
          <div className={styles.projectTitle}>
            <div className={styles.titleBox}>
              <p className={styles.title}>Projects</p>

              {isLoggedIn && (
                <Button
                  type="new"
                  text="New Project"
                  onClick={() => {
                    setModalType(ModalType.NewProjectModal)
                    showModal()
                  }}
                />
              )}
            </div>
            <div className={styles.line} />
          </div>
          <Table>{projectCards}</Table>
        </div>
      </div>
    </div>
  )
}

export default ProjectSearch
