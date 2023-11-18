import { useState } from 'react'
import './App.css'

function Editor() {
    const [problem, setProblem] = useState(null)
    const problemList = [
        { statement: "Problem 1"
        , description: "Solve Problem 1"
        , id: 1
        },
        { statement: "Problem 2"
        , description: "Solve Problem 2"
        , id: 2
        }
    ];
    if (problem == null){
        return (
            <>
                <h1>Select Problem</h1>
                <div>
                    <div>
                        <ProblemList problemList={problemList} setProblem = {setProblem}/>
                    </div>
                </div>
            </>
        )
    } else {
        let problemData;
        for (const each of problemList) {
            if (each.id == problem){
                problemData = each;
                break;
            }
        }
        return (
            <>
                <div
                    style={{ 
                        display: "flex", 
                        flexDirection: "row",
                        height: "100%",
                        margin: "20px",
                    }}
                >
                    <div style={{flex:1}}>
                        <h1>{problemData.statement}</h1>
                        <h3>{problemData.description}</h3>
                    </div>
                    <div style={{flex:1}}>
                        <textarea resize="none" name="Code Editor" id="primary" cols="100" rows="20"></textarea>
                        <button> Submit </button>
                    </div>
                </div>
            </>
        )
    }
}

function ProblemList(props){
    const Problem = props.problemList;
    const setProblem = props.setProblem;
    return (
        <>
            {Problem.map((problem) => { 
                return <List
                    key={problem.id}
                    id={problem.id}
                    name={problem.statement}
                    description={problem.description}
                    setProblem={setProblem}
                />; 
                })
            } 
        </>
    )
}

function List(props) { 
    return ( 
        <button 
            style={{ 
                display: "flex", 
                flexDirection: "column", 
                alignItems: "center", 
                height: "100%", 
                backgroundColor: "#fafafa", 
                margin: "20px", 
                width: "300px",
            }}
            onClick={() => props.setProblem(() => props.id)}
        > 
            {/* <div> */}
                {props.name}
            {/* </div> */}
        </button> 
    ); 
} 

export default Editor
