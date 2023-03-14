import React, { useState } from "react"

import styles from "./FivePointSelector.module.css"



const FivePointSelector = ({title, onChange}:{title:string, onChange: (num:number)=> void}) => {

  
    const [isActive, setIsActive] = useState<number>(0)

    function changeOption(num : number) {
        setIsActive(num)
        onChange(num)
    }


    return <div className={styles.container}>
        <div className={styles.title}>{title}</div>
        <div className={styles.options}>
        <button className={[styles.option, isActive === 1 && styles.selected].join(" ")} onClick={()=>changeOption(1)}>1</button>
        <button className={[styles.option, isActive === 2 && styles.selected].join(" ")} onClick={()=>changeOption(2)}>2</button>
        <button className={[styles.option, isActive === 3 && styles.selected].join(" ")} onClick={()=>changeOption(3)}>3</button>
        <button className={[styles.option, isActive === 4 && styles.selected].join(" ")} onClick={()=>changeOption(4)}>4</button>
        <button className={[styles.option, isActive === 5 && styles.selected].join(" ")} onClick={()=>changeOption(5)}>5</button>
        </div>
    </div>
}

export default FivePointSelector