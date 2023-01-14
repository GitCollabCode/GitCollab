import React, { useState } from 'react'
import styles from './Table.module.css'

export type rowType = {
    id:string,
    name:string,
    date:string,
    total:number,
    points:number,
    percent:number,
    status:string
}

const Table = ({rows, isExpandable, expandableRows}:{rows:any[], isExpandable: boolean, expandableRows?: string[]}) => {

    const initialData = {
        data : rows,
        expandedRows: ["0"]
    };

    const [tableData, setTableData]=useState(initialData);
   

    const handleRowClick = (rowId:string)=> {
        const currentExpandedRows = tableData.expandedRows;
        const isRowCurrentlyExpanded = currentExpandedRows.includes(rowId);
        
        const newExpandedRows = isRowCurrentlyExpanded ? 
			currentExpandedRows.filter((id:string) => id !== rowId) : 
			currentExpandedRows.concat(rowId);
        
        setTableData({data:tableData.data, expandedRows:newExpandedRows})
    }
    
    //eslint-disable-next-line
    const buildTableRow = (item:any, id:string) : JSX.Element=>{
        const result = Object.keys(item);
        const clickCallback = () => handleRowClick(id);
        //eslint-disable-next-line
        let tds:JSX.Element[] = []
        result.forEach(element => {
            tds.push(<td className={styles.row}>{item[element]}</td>)
        });
        return (<tr onClick={()=>isExpandable && clickCallback} key={"row-data-" + id}>
            {tds}
        </tr>);
    }

//eslint-disable-next-line
    const renderItem =(item:rowType, id:string):JSX.Element[]=> {
        
        const itemRows = [
			buildTableRow(item, id)
        ];
        /*
        if(isExpandable && tableData.expandedRows.includes(id)) {
            itemRows.push(
                <tr key={"row-expanded-" + id}>
                    <td>{item.name}</td>
                    <td>{item.points}</td>
                    <td>{item.percent}</td>
                </tr>
            );
        }
        */
        return itemRows;    
    }

    let allItemRows:any = [];
    //eslint-disable-next-line
    let header:JSX.Element[] = []
    Object.keys(tableData.data[0]).forEach(element => {
            header.push(<th>{element}</th>)
        });
    
    allItemRows = allItemRows.concat(<tr>{header}</tr>);
    tableData.data.forEach((item, index) => {
        const perItemRows = renderItem(item, ""+index);
        allItemRows = allItemRows.concat(perItemRows);
    });

   


return(
     <>
            <table>{allItemRows}</table>
    </>
)
}

export default Table;
