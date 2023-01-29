import React from "React";

import style from "./ProjectCard.module.css";

type cardData = {
    name:string,
    description:string,
    languages:string[],
    url:string
}

const ProjectCard = () => {
    const data = {
        name:"Project Name",
        description:"Lorem upsum",
        languages:["Stuff" ,"Other", "Other"],
        url:"#",
    }
    
    const getLanguagePills = () => {
         //eslint-disable-next-line
        let languagePills: JSX.Element[] = []
        data.languages.forEach(element=>{
            languagePills.push(<div className={style.pill}>{element}</div>)
        })
    
        return languagePills
    }


return (
    <div className={style.card}>
        <div className={style.info}>
            <h3>{data.name}</h3>
            <p>{data.description}</p>
        </div>
        <div className={style.tags}>
            {getLanguagePills()}
        </div>
        <div className={style.actions}>
            <button>I am a button</button>
        </div>
    </div>
)

}

export default ProjectCard;