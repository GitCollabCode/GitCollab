import React from "react";

import style from "./ProjectCard.module.css";
import { ProjectCardType } from "../../constants/common";



const ProjectCard = ({data}:{data:ProjectCardType}) => {
    
    
    const getLanguagePills = () => {
        if(data.project_skills === undefined){
            return <></>
        }
         //eslint-disable-next-line
        let languagePills: JSX.Element[] = []
        data.project_skills.forEach((element,index)=>{
            languagePills.push(<div className={style.pill} key={index}>{element}</div>)
        })
    
        return languagePills
    }

return ( 
    
    <div className={style.card}>
        <a href={process.env.REACT_APP_REDIRECT_URI +"project/"+data.project_name}>
        <div className={style.info}>
            <h3>{data.project_name}</h3> 
           {data.project_description!=="" && <p>{data.project_description}</p>}
        </div>
        </a>
        {data.project_skills.length !==0 && 
        <div className={style.tags}>
            {getLanguagePills()}
        </div>}

    </div>
)

}

export default ProjectCard;