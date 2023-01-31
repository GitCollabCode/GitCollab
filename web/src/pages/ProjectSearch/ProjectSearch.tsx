import React, { useContext, useState } from 'react'
import CheckboxFilter from '../../components/CheckboxFilter/CheckboxFilter'
import SearchBox from '../../components/SearchBox/SearchBox'
import Table from '../../components/Table/Tables'
import styles from './ProjectSearch.module.css'

import Button from '../../components/Button/Button'
import { ModalContextStateContext } from '../../context/modalContext/modalContext'
import { ModalType, ProjectCardType } from '../../constants/common'
import { UserLoginContext } from '../../context/userLoginContext/userLoginContext'
import ProjectCard from '../../components/ProjectCard/ProjectCard'

const ProjectSearch = () => {
  const { showModal, setModalType } = useContext(ModalContextStateContext)
  const { isLoggedIn } = useContext(UserLoginContext)
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

  const cardData: ProjectCardType[] = [
    {
      name: 'Project Name',
      description: 'Lorem upsum',
      languages: ['Stuff', 'Other', 'Other'],
      url: '#',
    },
    {
      name: 'Project Name',
      description: 'Lorem upsum',
      languages: ['Stuff', 'Other', 'Other', 'Super long pill'],
      url: '#',
    },
  ]

  const getProjectsCards = () => {
    //eslint-disable-next-line
    let cards: JSX.Element[] = []
    cardData.forEach((element, index) => {
      cards.push(
        <tr key={index}>
          <td>
            <ProjectCard data={element} key={index} />
          </td>
        </tr>
      )
    })
    return cards
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
          <Table>{getProjectsCards()}</Table>
        </div>
      </div>
    </div>
  )
}

export default ProjectSearch
