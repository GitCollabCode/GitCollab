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

const Table = ({rows, isExpandable, expandableRows}:{rows:any[], isExpandable: boolean, expandableRows: string[]}) => {


    const intitalDataRows: rowType[] = [
        {id : "1", date : "2014-04-18", total : 121.0, status : "Shipped", name : "A", points: 5, percent : 50},
        {id : "2", date : "2014-04-21", total : 121.0, status : "Not Shipped", name : "B", points: 10, percent: 60},
        {id : "3", date : "2014-08-09", total : 121.0, status : "Not Shipped", name : "C", points: 15, percent: 70},
        {id : "4", date : "2014-04-24", total : 121.0, status : "Shipped", name : "D", points: 20, percent : 80},
        {id : "5", date : "2014-04-26", total : 121.0, status : "Shipped", name : "E", points: 25, percent : 90},
    ]
    const initialData = {
        data : rows,
        expandedRows: ["0"]
    };

    const [tableData, setTableData]=useState(initialData);
   


    let allItemRows:any = [];
        
    tableData.data.forEach((item, index) => {
        const perItemRows = renderItem(item, ""+index);
        allItemRows = allItemRows.concat(perItemRows);
    });

   const handleRowClick = (rowId:string)=> {
        const currentExpandedRows = tableData.expandedRows;
        const isRowCurrentlyExpanded = currentExpandedRows.includes(rowId);
        
        const newExpandedRows = isRowCurrentlyExpanded ? 
			currentExpandedRows.filter((id:string) => id !== rowId) : 
			currentExpandedRows.concat(rowId);
        
        setTableData({data:tableData.data, expandedRows:newExpandedRows})
    }
    
    const buildTableRow = (item:any, id:string) : JSX.Element=>{
        const result = Object.keys(item);
        const clickCallback = () => handleRowClick(id);
        let tds:JSX.Element[] = []
        result.forEach(element => {
            tds.push(<td>{item[element]}</td>)
        });
        return <tr onClick={()=>isExpandable && clickCallback} key={"row-data-" + id}>
            {tds}
        </tr>
        
    }

    const renderItem =(item:rowType, id:string):JSX.Element[]=> {
        
        const itemRows = [
			buildTableRow(item, id)
        ];
        
        if(isExpandable && tableData.expandedRows.includes(id)) {
            itemRows.push(
                <tr key={"row-expanded-" + id}>
                    <td>{item.name}</td>
                    <td>{item.points}</td>
                    <td>{item.percent}</td>
                </tr>
            );
        }
        
        return itemRows;    
    }


return(
     <>
            <table>{allItemRows}</table>
    </>
)
}

export default Table;
