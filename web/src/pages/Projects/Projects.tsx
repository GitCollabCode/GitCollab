import React, { useState } from 'react'
import CheckboxFilter from '../../components/CheckboxFilter/CheckboxFilter'
import SearchBox from '../../components/SearchBox/SearchBox'
import Table, { rowType } from '../../components/Table/Tables'
import styles from './Projects.module.css'
import { Link } from 'react-router-dom'
import Button from '../../components/Button/Button'

const Projects = () => {
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

  const intitalDataRows: rowType[] = [
    {
      id: '1',
      date: '2014-04-18',
      total: 121.0,
      status: 'Shipped',
      name: 'A',
      points: 5,
      percent: 50,
    },
    {
      id: '2',
      date: '2014-04-21',
      total: 121.0,
      status: 'Not Shipped',
      name: 'B',
      points: 10,
      percent: 60,
    },
    {
      id: '3',
      date: '2014-08-09',
      total: 121.0,
      status: 'Not Shipped',
      name: 'C',
      points: 15,
      percent: 70,
    },
    {
      id: '4',
      date: '2014-04-24',
      total: 121.0,
      status: 'Shipped',
      name: 'D',
      points: 20,
      percent: 80,
    },
    {
      id: '5',
      date: '2014-04-26',
      total: 121.0,
      status: 'Shipped',
      name: 'E',
      points: 25,
      percent: 90,
    },
  ]

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
              <Link to={'/new-project'} className={styles.noLink}>
                <Button type="new" text="New Project" />
              </Link>
            </div>
            <div className={styles.line} />
          </div>
          <Table rows={intitalDataRows} isExpandable={false} />
        </div>
      </div>
    </div>
  )
}

export default Projects
